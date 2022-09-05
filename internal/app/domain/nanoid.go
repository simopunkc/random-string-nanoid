package domain

type NanoidResult struct {
	Value string `json:"value,omitempty"`
}

type NanoidConfig struct {
	Prefix         string `json:"prefix,omitempty"`
	LimitCharacter int    `json:"limit,omitempty"`
}
