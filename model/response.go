package model

type DetailedUserResponse struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	CreatedAt     string `json:"created_at"`
	EmailVerified bool   `json:"email_verified"`
}

type Success struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessResponse struct {
	Success Success `json:"success"`
}

type AuthData struct {
	User  DetailedUserResponse `json:"user"`
	Token string               `json:"token"`
}
