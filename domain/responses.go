package domain

type SaveSuccesResponse struct {
	Alias      string `json:"alias"`
	Url        string `json:"url"`
	Original   string `json:"original_url"`
	Statistics struct {
		TimeTaken string `json:"time_taken"`
	} `json:"statistics"`
}

type ErrorResponse struct {
	Alias       string `json:"alias,omitempty"`
	ERR_Code    string `json:"err_code"`
	Description string `json:"description"`
}
