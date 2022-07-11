package model

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	UserId      uint32 `json:"user_id,omitempty"`
	Username    string `json:"username,omitempty"`
	Email       string `json:"email,omitempty"`
}

type JwtPayload struct {
	UserId     uint32
	Username   string
	Email      string
	AccessUuid string
}

type TokenDetails struct {
	AccessToken string
	AccessUuid  string
	AtExpires   int64
}
