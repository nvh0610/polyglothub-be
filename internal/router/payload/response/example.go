package response

type ExampleResponse struct {
	Id           int    `json:"id"`
	VocabularyId int    `json:"vocabulary_id"`
	Sentence     string `json:"sentence"`
	Meaning      string `json:"meaning"`
}
