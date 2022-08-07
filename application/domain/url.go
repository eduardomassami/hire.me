package domain

type URL struct {
	Alias          string `json:"alias"`
	URL            string `json:"url"`
	RetrievalCount int    `json:"retrievalCount"`
}
