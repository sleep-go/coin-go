package base

type Msg struct {
	Success    bool   `json:"success"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}
