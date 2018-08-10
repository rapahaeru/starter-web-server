package api

//Token structure
type Token struct {
	Token string `json:"token"`
	Owner string `json:"owner"`
}

// GetToken brings a token list to authenticate the api
func GetToken() []Token {
	data := []Token{
		{"alecrimHash#!", "cael"},
		{"miGueLit0", "Miguelito"},
	}
	return data
}
