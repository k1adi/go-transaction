package model

type Auth struct {
	Id       string `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	Role     string `json:"role,omitempty"`
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}
