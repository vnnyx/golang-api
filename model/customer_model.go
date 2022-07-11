package model

type CustomerCreateRequest struct {
	Id       uint32 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type CustomerUpdateRequest struct {
	Id       uint32 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type CustomerResponse struct {
	Id       uint32 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}
