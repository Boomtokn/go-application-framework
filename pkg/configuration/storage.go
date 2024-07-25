package configuration

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/snyk/go-application-framework/internal/utils"
)

//go:generate $GOPATH/bin/mockgen -source=storage.go -destination ../mocks/config_storage.go -package mocks -self_package github.com/snyk/go-application-framework/pkg/configuration/

type Storage interface {
	Set(key string, value any) error
}

type EmptyStorage struct{}

func (e *EmptyStorage) Set(_ string, _ any) error {
	return nil
}

// keyDeleted is a marker value which, when set, causes a key to be deleted from
// stored configuration.
var keyDeleted = struct{}{}

type JsonStorage struct {
	path   string
	config Configuration
}

type JsonOption func(*JsonStorage)

func WithConfiguration(c Configuration) JsonOption {
	return func(storage *JsonStorage) {
		storage.config = c
	}
}

func NewJsonStorage(path string, options ...JsonOption) *JsonStorage {
	storage := &JsonStorage{
		path: path,
	}

	for _, opt := range options {
		opt(storage)
	}

	return storage
}

// This function deals with the fact that not every key can or shall be written to the config. Keys that belong to
// Environment Variables need to be matched to their alternative names in the config.
// For example "SNYK_TOKEN" in the config file would be "api"
// The logic should in the future be moved closer to the configuration as it might be needed there as well.
func (s *JsonStorage) getNonEnvVarKey(key string) string {
	if s.config == nil {
		return ""
	}

	keys := []string{key}
	keys = append(keys, s.config.GetAlternativeKeys(key)...)
	for _, k := range keys {
		if s.config.GetKeyType(k) != EnvVarKeyType {
			return k
		}
	}

	return ""
}

func (s *JsonStorage) Set(key string, value any) error {
	// Check if path to file exists
	err := os.MkdirAll(filepath.Dir(s.path), utils.FILEPERM_755)
	if err != nil {
		return err
	}

	fileBytes, err := os.ReadFile(s.path)
	if err != nil {
		const emptyJson = "{}"
		fileBytes = []byte(emptyJson)
	}

	config := make(map[string]any)
	err = json.Unmarshal(fileBytes, &config)
	if err != nil {
		return err
	}

	if tmpKey := s.getNonEnvVarKey(key); len(tmpKey) > 0 {
		key = tmpKey
	}

	if _, ok := value.(struct{}); ok {
		// See implementation of Configuration.Unset; when marker value is set,
		// key is deleted from config before writing.
		delete(config, key)
	} else {
		config[key] = value
	}
	configJson, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.path, configJson, utils.FILEPERM_666)

	return err
}
