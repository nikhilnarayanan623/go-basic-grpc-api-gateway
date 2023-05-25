package domain

type Response struct {
	StatusCode uint32
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
	Data       interface{} `json:"data"`
}

type SignupRequest struct {
	FristName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type SignupRsponse struct {
	StatusCode uint32 `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	StatusCode  uint32 `json:"status_code"`
	Message     string `json:"message"`
	Error       string `json:"error"`
	AccessToken string `json:"access_token"`
}

type ValidatTokenRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
}

type ValidateTokenResponse struct {
	StatusCode uint32 `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
	UserID     uint32 `json:"user_id"`
}
