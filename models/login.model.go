package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserId int    `json:"userId"`
	RoleId int    `json:"roleId"`
}
