package models

type Ranking struct {
	Username string `json: "username"`
	Amount   int    `json: "amount"`
}
