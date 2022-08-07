package domain

type ErrorResponse struct {
	Alias       string `json:"alias,omitempty"`
	ERR_Code    string `json:"err_code"`
	Description string `json:"description"`
}
