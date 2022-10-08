package endpoint

type GenerateTokenRequest struct {
	Id uint64 `json:"id"`
}

type ValidateTokenRequest struct {
	Token string `json:"token"`
}
