package endpoint

type CreateRequest struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Pwd      string `json:"pwd" validate:"required"`
}
