package api

type ApiReponse struct {
	Version string        `json:"version"`
	Items   []interface{} `json:"items"`
	SentAt  int64         `json:"sent_at"`
}
