package sonar

type Sonar struct {
	Issues []Issue `json:"issues"`
}

type Phpmd struct {
	Version     string `json:"version"`
	PackageName string `json:"package"`
	Files       []File `json:"files"`
}

type File struct {
	Name       string      `json:"file"`
	Violations []violation `json:"violations"`
}

type Issue struct {
	EngineId        string          `json:"engineId"`
	RuleId          string          `json:"ruleId"`
	Typ             string          `json:"type"`
	Severity        string          `json:"severity"`
	PrimaryLocation PrimaryLocation `json:"primaryLocation"`
	effortMinutes   int
}

type PrimaryLocation struct {
	Message   string    `json:"message"`
	FilePath  string    `json:"filePath"`
	TextRange TextRange `json:"textRange"`
}

type TextRange struct {
	StartLine   int `json:"startLine"`
	EndLine     int `json:"endLine"`
	StartColumn int `json:"startColumn,omitempty"`
	EndColumn   int `json:"endColumn,omitempty"`
}

type violation struct {
	BeginLine       int    `json:"beginLine"`
	EndLine         int    `json:"endLine"`
	PackageName     string `json:"package"`
	Function        string `json:"function"`
	Class           string `json:"class"`
	Method          string `json:"method"`
	Description     string `json:"description"`
	Rule            string `json:"rule"`
	RuleSet         string `json:"ruleSet"`
	ExternalInfoUrl string `json:"externalInfoUrl"`
	Priority        int    `json:"priority"`
}
