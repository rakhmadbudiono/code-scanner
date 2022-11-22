package scanner

import "strings"

type RuleSeverity string

const (
	High   RuleSeverity = "HIGH"
	Medium RuleSeverity = "MEDIUM"
	Low    RuleSeverity = "LOW"
)

type Rule struct {
	ID          string
	Type        string
	Severity    RuleSeverity
	Description string
	Checker     func(string) bool
}

func GetRules() []Rule {
	return []Rule{
		{
			ID:          "G400",
			Type:        "sast",
			Severity:    High,
			Description: "private key exposed",
			Checker: func(line string) bool {
				return strings.Contains(line, "private_key")
			},
		},
		{
			ID:          "G401",
			Type:        "sast",
			Severity:    High,
			Description: "public key exposed",
			Checker: func(line string) bool {
				return strings.Contains(line, "public_key")
			},
		},
	}
}
