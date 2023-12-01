package request

type SendMessageRequest struct {
	Flow       string                 `json:"flow"`
	Parameters map[string]interface{} `json:"parameters"`
}
