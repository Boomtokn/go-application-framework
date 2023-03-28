package app

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/snyk/go-application-framework/internal/api"
	"github.com/snyk/go-application-framework/internal/constants"
	"github.com/snyk/go-application-framework/internal/utils"
	"github.com/snyk/go-application-framework/pkg/configuration"
	localworkflows "github.com/snyk/go-application-framework/pkg/local_workflows"
	"github.com/snyk/go-application-framework/pkg/networking"
	"github.com/snyk/go-application-framework/pkg/workflow"
	"github.com/snyk/go-httpauth/pkg/httpauth"
)

// initConfiguration initializes the configuration with initial values.
func initConfiguration(config configuration.Configuration, apiClient api.ApiClient, logger *log.Logger) {
	dir, _ := utils.SnykCacheDir()

	config.AddDefaultValue(configuration.FF_OAUTH_AUTH_FLOW_ENABLED, configuration.StandardDefaultValueFunction(false))
	config.AddDefaultValue(configuration.ANALYTICS_DISABLED, configuration.StandardDefaultValueFunction(false))
	config.AddDefaultValue(configuration.WORKFLOW_USE_STDIO, configuration.StandardDefaultValueFunction(false))
	config.AddDefaultValue(configuration.PROXY_AUTHENTICATION_MECHANISM, configuration.StandardDefaultValueFunction(httpauth.StringFromAuthenticationMechanism(httpauth.AnyAuth)))
	config.AddDefaultValue(configuration.DEBUG_FORMAT, configuration.StandardDefaultValueFunction(log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix|log.LUTC))
	config.AddDefaultValue(configuration.CACHE_PATH, configuration.StandardDefaultValueFunction(dir))
	config.AddDefaultValue(configuration.AUTHENTICATION_SUBDOMAINS, configuration.StandardDefaultValueFunction([]string{"deeproxy"}))

	config.AddDefaultValue(configuration.API_URL, func(existingValue any) any {
		urlString := constants.SNYK_DEFAULT_API_URL

		if existingValue != nil {
			if temp, ok := existingValue.(string); ok {
				urlString = temp
			}
		}

		apiString, _ := api.GetCanonicalApiUrlFromString(urlString)
		return apiString
	})

	config.AddDefaultValue(configuration.WEB_APP_URL, func(existingValue any) any {
		canonicalApiUrl := config.GetString(configuration.API_URL)
		appUrl, _ := api.DeriveAppUrl(canonicalApiUrl)
		return appUrl
	})

	config.AddDefaultValue(configuration.ORGANIZATION, func(existingValue any) any {
		client := networking.NewNetworkAccess(config).GetHttpClient()
		url := config.GetString(configuration.API_URL)
		apiClient.Init(url, client)
		if existingValue != nil && len(existingValue.(string)) > 0 {
			orgId := existingValue.(string)
			_, err := uuid.Parse(orgId)
			isSlugName := err != nil
			if isSlugName {
				orgId, err = apiClient.GetOrgIdFromSlug(existingValue.(string))
				if err == nil {
					return orgId
				}
			} else {
				return orgId
			}
		}

		orgId, _ := apiClient.GetDefaultOrgId()

		return orgId
	})

	config.AddDefaultValue(configuration.FF_OAUTH_AUTH_FLOW_ENABLED, func(existingValue any) any {
		alternativeBearerKeys := config.GetAlternativeKeys(configuration.AUTHENTICATION_BEARER_TOKEN)
		alternativeAuthKeys := config.GetAlternativeKeys(configuration.AUTHENTICATION_TOKEN)
		alternativeKeys := append(alternativeBearerKeys, alternativeAuthKeys...)

		for _, key := range alternativeKeys {
			hasPrefix := strings.HasPrefix(key, "snyk_")
			if hasPrefix {
				formattedKey := strings.ToUpper(key)
				_, ok := os.LookupEnv(formattedKey)
				if ok {
					logger.Printf("Found environment variable %s, disabling OAuth flow", formattedKey)
					return false
				}
			}
		}
		return existingValue
	})

	config.AddAlternativeKeys(configuration.AUTHENTICATION_TOKEN, []string{"snyk_token", "snyk_cfg_api", "api"})
	config.AddAlternativeKeys(configuration.AUTHENTICATION_BEARER_TOKEN, []string{"snyk_oauth_token", "snyk_docker_token"})
	config.AddAlternativeKeys(configuration.API_URL, []string{"endpoint"})
}

// CreateAppEngine creates a new workflow engine.
func CreateAppEngine() workflow.Engine {
	discardLogger := log.New(io.Discard, "", 0)
	return CreateAppEngineWithLogger(discardLogger)
}

// App engine with logger injected
func CreateAppEngineWithLogger(logger *log.Logger) workflow.Engine {
	config := configuration.New()
	apiClient := api.NewApiInstance()

	initConfiguration(config, apiClient, logger)

	engine := workflow.NewWorkFlowEngine(config)

	engine.AddExtensionInitializer(localworkflows.Init)

	return engine
}
