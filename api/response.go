package api

type ApiResponse struct {
	Version string        `json:"version"`
	Items   []interface{} `json:"items"`
	SentAt  int64         `json:"sent_at"`
}
