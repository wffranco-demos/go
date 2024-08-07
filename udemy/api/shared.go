package api

import "time"

const base = "https://reqres.in/api"

type UserData struct {
	User User `json:"data"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type UserCreated struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Job       string    `json:"job"`
	CreatedAt time.Time `json:"createdAt"`
}
