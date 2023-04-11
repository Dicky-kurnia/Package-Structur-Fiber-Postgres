package model

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type JwtPayload struct {
	SalesId  uint   `json:"sales_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type TokenDetails struct {
	AccessToken string
	AtExpires   int64
	RtExpires   int64
}

type GetUserProfileResponse struct {
	SalesId  int32  `json:"sales_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type ChangePasswordRequest struct {
	Username           string
	OldPassword        string `json:"old_password" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required,gte=8,lte=100"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required,gte=8,lte=100"`
}
