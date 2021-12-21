package dto

//LoginDTO é um modelo que o cliente usa quando faz um POST de / login url
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}