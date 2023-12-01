package model

type User struct {
	ID        int64  `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	BotID     int64  `json:"bot_id" db:"bot_id"`
	BotName   string `json:"bot_name" db:"bot_name"`
}
