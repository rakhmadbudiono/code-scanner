package scanner

type Finding struct {
	Type   string
	RuleID string
	Location
	Metadata
}

type Location struct {
	Path string
	Positions
}

type Positions struct {
	Begin
}

type Begin struct {
	Line int
}

type Metadata struct {
	Description string
	Severity    RuleSeverity
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
