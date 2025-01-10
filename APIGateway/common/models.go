package common

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"jwt"`
}
