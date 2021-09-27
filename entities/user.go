package entities

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	First    string `json:"first"`
	Last     string `json:"last"`
}
