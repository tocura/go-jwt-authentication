package response

type Token struct {
	Token string `json:"token"`
}

func MapToTokenResponse(token string) *Token {
	return &Token{
		Token: token,
	}
}
