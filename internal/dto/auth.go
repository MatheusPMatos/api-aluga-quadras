package dto

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken string `json:"access_token"`
}
