package dto

//LoginDTO is a model that used by client when POST from /login url
type LoginDto struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

//ChangePasswordDto is a model that used by client when POST from /login url
type ChangePasswordDto struct {
	Id          string `json:"id" form:"id" binding:"required"`
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

type RecoverPasswordDto struct {
	Id          string `json:"id" form:"id"`
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
}

type ChangePhoneDto struct {
	Id    string `json:"id" form:"id"`
	Phone string `json:"phone" form:"phone"`
}

type EmailVerificationDto struct {
	Id string `json:"id" form:"id"`
}
