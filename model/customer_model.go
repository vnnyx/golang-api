package model

type CustomerCreateRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Gender   string `json:"gender,omitempty"`
}

type CustomerUpdateRequest struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Gender   string `json:"gender,omitempty"`
}

type CustomerResponse struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Gender   string `json:"gender,omitempty"`
}
