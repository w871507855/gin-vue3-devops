package request

type UserLoginReq struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

type UserRegisterReq struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}
