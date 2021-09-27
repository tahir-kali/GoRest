package entities

type Transaction struct {
	User    string `json:"user"`
	Type    string `json:"type"`
	Balance int    `json:"balance"`
	Date    string `json:"date"`
}
