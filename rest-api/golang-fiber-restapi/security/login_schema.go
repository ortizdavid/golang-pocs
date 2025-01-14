package security

type LoginRequest struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

type LoginResponse struct {
	Token string `json:"Token"`
}