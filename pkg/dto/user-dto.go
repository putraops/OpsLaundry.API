package dto

type RegisterDto struct {
	FirstName string `json:"first_name" form:"first_name" binding:"required"`
	LastName  string `json:"last_name" form:"last_name"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password" form:"password" binding:"required"`
	UserType  int    `json:"user_type" form:"user_type" binding:"required"`
	IsAdmin   bool   `json:"is_admin" form:"is_admin;default:false"`
}

// type ApplicationUserDescriptionDto struct {
// 	Id          string `json:"id" form:"id" binding:"required"`
// 	Title       string `json:"title" form:"title"`
// 	Description string `json:"description" form:"description"`
// 	UpdatedBy   string
// }

// type ApplicationUserUpdateDto struct {
// 	Id       string `json:"id" form:"id"`
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required,email"`
// 	Password string `json:"password,omitempty" form:"password,omitempty"`
// }
