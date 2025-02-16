package presenters

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/snyk/code-client-go/sarif"

	sarif_utils "github.com/snyk/go-application-framework/internal/utils/sarif"
)

type Finding struct {
	ID         string
	Severity   string
	Title      string
	Ignored    bool
	Properties []FindingProperty
}

type FindingProperty struct {
	Label string
	Value string
}

type SarifPresenter struct {
	ShowIgnored      bool
	Input            sarif.SarifDocument
	OrgName          string
	TestPath         string
	SeverityMinLevel string
}

type PresenterOption func(*SarifPresenter)

func WithIgnored(showIgnored bool) PresenterOption {
	return func(p *SarifPresenter) {
		p.ShowIgnored = showIgnored
	}
}

func WithOrgName(orgName string) PresenterOption {
	return func(p *SarifPresenter) {
		p.OrgName = orgName
	}
}

func WithTestPath(testPath string) PresenterOption {
	return func(p *SarifPresenter) {
		p.TestPath = testPath
	}
}

func WithSeverityThershold(severityMinLevel string) PresenterOption {
	return func(p *SarifPresenter) {
		p.SeverityMinLevel = severityMinLevel
	}
}

func SarifTestResults(sarifDocument sarif.SarifDocument, options ...PresenterOption) *SarifPresenter {
	p := &SarifPresenter{
		ShowIgnored:      false,
		Input:            sarifDocument,
		OrgName:          "",
		TestPath:         "",
		SeverityMinLevel: "",
	}

	for _, option := range options {
		option(p)
	}

	return p
}

func FilterFindingsBySeverity(findings []Finding, minLevel string, severityOrder []string) []Finding {
	var filteredFindings []Finding

	filteredSeverityASC := FilterSeverityASC(severityOrder, minLevel)
	for _, finding := range findings {
		if slices.Contains(filteredSeverityASC, finding.Severity) {
			filteredFindings = append(filteredFindings, finding)
		}
	}
	return filteredFindings
}

func (p *SarifPresenter) Render() (string, error) {
	summaryData := sarif_utils.CreateCodeSummary(&p.Input, p.TestPath)
	findings :=
		SortFindings(convertSarifToFindingsList(p.Input), summaryData.SeverityOrderAsc)
	summaryOutput, err := RenderSummary(summaryData, p.OrgName, p.TestPath, p.SeverityMinLevel)

	if err != nil {
		return "", err
	}

	// Filter findings based on severity
	findings = FilterFindingsBySeverity(findings, p.SeverityMinLevel, summaryData.SeverityOrderAsc)

	str := strings.Join([]string{
		"",
		renderBold(fmt.Sprintf("Testing %s ...", p.TestPath)),
		RenderFindings(findings, p.ShowIgnored, len(p.SeverityMinLevel) > 0),
		summaryOutput,
		getFinalTip(p.ShowIgnored),
		"",
	}, "\n")

	return str, nil
}

func convertSarifToFindingsList(input sarif.SarifDocument) []Finding {
	var findings []Finding
	for _, run := range input.Runs {
		for _, result := range run.Results {
			var rule sarif.Rule
			for _, ruleItem := range run.Tool.Driver.Rules {
				if ruleItem.ID == result.RuleID {
					rule = ruleItem
					break
				}
			}

			severity := sarif_utils.SarifLevelToSeverity(result.Level)

			var title string
			if rule.ShortDescription.Text != "" {
				title = rule.ShortDescription.Text
			} else {
				title = rule.Name
			}

			location := sarif.Location{}
			if len(result.Locations) > 0 {
				location = result.Locations[0]
			}

			isIgnored := len(result.Suppressions) > 0

			var properties []FindingProperty

			properties = append(properties, FindingProperty{
				Label: "Path",
				Value: fmt.Sprintf("%s, line %d",
					location.PhysicalLocation.ArtifactLocation.URI,
					location.PhysicalLocation.Region.StartLine,
				),
			})

			properties = append(properties, FindingProperty{
				Label: "Info",
				Value: result.Message.Text,
			})

			properties = append(properties, FindingProperty{
				Label: "",
				Value: "",
			})

			for _, suppression := range result.Suppressions {
				expiration := ""
				if suppression.Properties.Expiration != nil {
					expiration = *suppression.Properties.Expiration
				}

				properties = append(properties, FindingProperty{
					Label: "Expiration",
					Value: expiration,
				})

				properties = append(properties, FindingProperty{
					Label: "Category",
					Value: strings.Replace(string(suppression.Properties.Category), "wont-fix", "Won't fix", 1),
				})

				// TODO: Verify date formatting
				s, err := time.Parse(time.RFC3339, suppression.Properties.IgnoredOn)

				if err == nil {
					properties = append(properties, FindingProperty{
						Label: "Ignored on",
						Value: s.Format("January 02, 2006"),
					})
				}

				properties = append(properties, FindingProperty{
					Label: "Ignored by",
					Value: suppression.Properties.IgnoredBy.Name,
				})

				properties = append(properties, FindingProperty{
					Label: "Reason",
					Value: suppression.Justification,
				})
			}

			findings = append(findings, Finding{
				ID:         result.RuleID,
				Severity:   severity,
				Title:      title,
				Ignored:    isIgnored,
				Properties: properties,
			})
		}
	}
	return findings
}

func getFinalTip(showIgnored bool) string {
	tip := "To view ignored issues, use the --include-ignores option."
	if showIgnored {
		return ""
	}

	return RenderTip(tip)
}

func SortFindings(findings []Finding, order []string) []Finding {
	result := make([]Finding, 0, len(findings))

	result = append(result, findings...)

	slices.SortFunc(result, func(a, b Finding) int {
		if a.Severity != b.Severity {
			return slices.Index(order, a.Severity) - slices.Index(order, b.Severity)
		}

		return 0
	})

	return result
}
