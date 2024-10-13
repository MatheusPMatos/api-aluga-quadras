package dto

type Auth struct {
	Email    string
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken string `json:"access_token"`
}
