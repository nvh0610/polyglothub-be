package response

type VocabularyResponse struct {
	ID        int               `json:"id"`
	Word      string            `json:"word"`
	Meaning   string            `json:"meaning"`
	IPA       string            `json:"ipa"`
	Type      string            `json:"type"`
	Url       string            `json:"url"`
	Examples  []ExampleResponse `json:"examples"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}
