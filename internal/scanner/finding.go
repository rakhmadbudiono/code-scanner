package scanner

type Finding struct {
	Type     string `json:"type"`
	RuleID   string `json:"ruleId"`
	Location `json:"location"`
	Metadata `json:"metadata"`
}

type Location struct {
	Path      string `json:"path"`
	Positions `json:"positions"`
}

type Positions struct {
	Begin `json:"begin"`
}

type Begin struct {
	Line int `json:"line"`
}

type Metadata struct {
	Description string       `json:"description"`
	Severity    RuleSeverity `json:"severity"`
}

func NewFinding(path string, line int, rule Rule) *Finding {
	return &Finding{
		Type:   rule.Type,
		RuleID: rule.ID,
		Location: Location{
			Path: path,
			Positions: Positions{
				Begin: Begin{
					Line: line,
				},
			},
		},
		Metadata: Metadata{
			Description: rule.Description,
			Severity:    rule.Severity,
		},
	}
}
