package v1

import (
	"net/http"
	"server/helper"
	"server/models/request"
	"server/models/response"

	"server/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

func UserRegisterAPI(c *gin.Context) {
	var u request.UserRegisterReq
	err := c.ShouldBindJSON(&u)
	// 判断是否为空
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Info("Login", zap.String("user", u.UserName))
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": helper.RemoveTopStruct(errs.Translate(helper.Trans)),
		})
		return
	}
	err = service.UserRegisterService(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("用户注册成功", c)
	}

}

func LoginAPI(c *gin.Context) {
	var u request.UserLoginReq
	err := c.ShouldBindJSON(&u)
	// 判断是否为空
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Info("Login", zap.String("user", u.UserName))
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": helper.RemoveTopStruct(errs.Translate(helper.Trans)),
		})
		return
	}
	// 判断是否存在
	user, err := service.UserLoginService(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		//	生成token
		token, _ := helper.GenerateToken(user.UserID, user.Username)
		response.OkWithDetailed(map[string]interface{}{
			"token": token,
		}, "登录成功", c)
	}
}

func UserDeleteAPI(c *gin.Context) {
	id := c.Query("id")
	err := service.UserDeleteService(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("用户删除成功", c)
	}
}

func UserInfoAPI(c *gin.Context) {

	claims, _ := c.Get("claims")

	user, err := service.UserInfoService(claims.(*helper.UserClaims).UUID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(user, "用户查询成功", c)
	}
}
