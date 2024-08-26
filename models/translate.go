package models

type Translate struct {
	BaseModel

	Phrase string `json:"phrase"`
}

type TranslateResult struct {
	BaseModel

	Translation []string `json:"translation"`
}
