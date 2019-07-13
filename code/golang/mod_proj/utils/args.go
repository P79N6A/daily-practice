package utils

type CreateUserSpec struct {
	Name        string `form:"name" json:"name" xml:"name" binding:"required"`
	Password    string `form:"password" json:"password" xml:"password" binding:"required"`
	Description string `form:"description" json:"description" xml:"description" binding:"required"`
	Avatar      string `form:"avatar" json:"avatar" xml:"avatar" binding:"required"`
}

type UpdateUserSpec struct {
	Name        string `form:"name" json:"name" xml:"name" binding:"required"`
	Password    string `form:"password" json:"password" xml:"password" binding:"required"`
	Description string `form:"description" json:"description" xml:"description" binding:"required"`
	Avatar      string `form:"avatar" json:"avatar" xml:"avatar" binding:"required"`
}
