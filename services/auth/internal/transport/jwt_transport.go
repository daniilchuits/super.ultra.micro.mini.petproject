package transport

type JWT struct {
	Token string `json:"jwt_token"`
}

func NewJWTTransport(token string) JWT {
	return JWT{Token: token}
}
