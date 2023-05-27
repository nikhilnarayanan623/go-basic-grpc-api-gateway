package domain

type Response struct {
	StatusCode uint32
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
	Data       interface{} `json:"data"`
}

type SignupRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// type SignupRsponse struct {
// 	UserID uint32 `json:"user_id"`
// }

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type ValidateTokenRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
}

type ValidateTokenResponse struct {
	UserID uint32 `json:"user_id"`
}
