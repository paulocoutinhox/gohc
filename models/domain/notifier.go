package domain

type Notifier struct {
	ID     string                 `json:"id"`
	Plugin string                 `json:"plugin"`
	Params map[string]interface{} `json:"params"`
}
