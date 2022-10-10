package service

import (
	"errors"
	"fmt"
	"server/common"
	"server/helper"
	"server/models"
	"server/models/request"
)

// GetUserInfo
func UserInfoService(uuid int64) (user *models.User, err error) {
	// for k, v := range c.Request.Header {
	// 	println("头部信息", k, v)
	// }
	//token := c.Request.Header.Get("X-Access-Token")
	//data := new(models.User)
	//info, err := helper.AnalyseToken(token)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": -1,
	//		"mgs":  "token解析失败",
	//	})
	//}
	fmt.Println("uuid", uuid)
	err = common.DB.Where("user_id = ?", uuid).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	fmt.Println("用户详情", user)
	//user.Username = info.Name
	//user.UserID = info.UserID
	var role models.Role
	role.Name = "admin"
	user.Roles = []models.Role{
		role,
	}
	return user, nil
}

// // GetUserDetail
// // @Tags 公共方法
// // @Summary 用户详情
// // @Param id query string false "请输入用户id"
// // @Success 200 {string} json "{"code": "200", "data": ""}"
// // @Router /user-detail [get]
// func GetUserDetail(c *gin.Context) {
// 	id := c.Query("id")
// 	if id == "" {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": -1,
// 			"msg":  "用户id不能为空",
// 		})
// 		return
// 	}
// 	data := new(models.User)
// 	err := common.DB.Omit("password").Where("id = ?", id).Find(&data).Error
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": -1,
// 			"msg":  "get user detail by id:" + id + "error:" + err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 200,
// 		"data": data,
// 	})
// }

// // Login
// // @Tags 公共方法
// // @Summary 用户登录
// // @Param username formData  string false "请输入用户名"
// // @Param password formData  string false "请输入用户密码"
// // @Success 200 {string} json "{"code": "200", "data": ""}"
// // @Router /login [post]
// func Login(c *gin.Context) {
// 	// username := c.PostForm("username")
// 	// password := c.PostForm("password")
// 	var u request.UserReq
// 	err := c.ShouldBindJSON(&u)
// 	// 判断是否为空
// 	// if err != nil {
// 	// 	errs, ok := err.(validator.ValidationErrors)
// 	// 	if !ok {
// 	// 		zap.L().Info("Login", zap.String("user", u.UserName))
// 	// 		c.JSON(http.StatusOK, gin.H{
// 	// 			"msg": err.Error(),
// 	// 		})
// 	// 		return
// 	// 	}
// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"msg": helper.RemoveTopStruct(errs.Translate(helper.Trans)),
// 	// 	})
// 	// 	return
// 	// }
// 	helper.Verify(err, c)

// 	// 密码转成md5
// 	password := helper.GetMd5(u.PassWord)
// 	// 查询数据库中是否存在
// 	data := new(models.User)
// 	err = common.DB.Where("username = ? AND password = ?", u.UserName, password).First(&data).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			// c.JSON(http.StatusOK, gin.H{
// 			// 	"code": -1,
// 			// 	"msg":  "用户名或密码错误",
// 			// })
// 			response.FailWithMessage("用户名或密码错误", c)
// 			return

// 		}
// 		// c.JSON(http.StatusOK, gin.H{
// 		// 	"code": -1,
// 		// 	"msg":  "Get User Error:" + err.Error(),
// 		// })
// 		response.FailWithMessage("Get User Error:"+err.Error(), c)
// 		return
// 	}

// 	token, err := helper.GenerateToken(data.UUID, data.Username)
// 	if err != nil {
// 		// c.JSON(http.StatusOK, gin.H{
// 		// 	"code": -1,
// 		// 	"msg":  "GenerateToken Error:" + err.Error(),
// 		// })
// 		response.FailWithMessage("Get User Error:"+err.Error(), c)
// 		return
// 	}

// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"code": 200,
// 	// 	"data": map[string]interface{}{
// 	// 		"accessToken": token,
// 	// 	},
// 	// })
// 	response.OkWithData(map[string]interface{}{
// 		"accessToken": token,
// 	}, c)
// }

// // 用户注册
// // Register
// // @Tags 公共方法
// // @Summary 用户注册
// // @Param username formData  string false "请输入用户名"
// // @Param password formData  string false "请输入用户密码"
// // @Success 200 {string} json "{"code": "200", "data": ""}"
// // @Router /register [post]
// func Register(c *gin.Context) {
// 	var u request.UserRegisterReq
// 	err := c.ShouldBindJSON(&u)
// 	if err != nil {
// 		errs, ok := err.(validator.ValidationErrors)
// 		if !ok {
// 			zap.L().Error("Register", zap.String("user", u.UserName))
// 			c.JSON(http.StatusOK, gin.H{
// 				"msg": err.Error(),
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"msg": helper.RemoveTopStruct(errs.Translate(helper.Trans)),
// 		})
// 		return
// 	}

// 	// 数据插入
// 	data := &models.User{
// 		Username: u.UserName,
// 		Password: helper.GetMd5(u.PassWord),
// 		UUID:     helper.GetUUID(),
// 	}
// 	err = common.DB.Create(data).Error
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": -1,
// 			"msg":  "create user Error:" + err.Error(),
// 		})
// 		return
// 	}

// 	// 生成token
// 	token, err := helper.GenerateToken(data.UUID, data.Username)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": -1,
// 			"msg":  "GenerateToken Error:" + err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 200,
// 		"data": map[string]interface{}{
// 			"token": token,
// 		},
// 	})
// }

func UserRegisterService(u *request.UserRegisterReq) (err error) {
	var user models.User
	// 判断用户存不存在
	err = common.DB.Where("username = ?", u.UserName).First(&user).Error
	if err == nil {
		return errors.New("用户已存在")
	}
	// 生成UID
	userID := helper.GenID()
	// 保存进数据库
	user.Username = u.UserName
	user.Password = helper.GetMd5(u.PassWord)
	user.UserID = userID
	err = common.DB.Create(&user).Error
	if err != nil {
		return errors.New("用户注册失败")
	}
	return
}

func UserLoginService(u *request.UserLoginReq) (user *models.User, err error) {
	//var user models.User
	// 判断用户存不存在
	err = common.DB.Where("username = ?", u.UserName).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if helper.GetMd5(u.PassWord) != user.Password {
		return nil, errors.New("密码不正确")
	}
	return user, err
}

func UserDeleteService(id string) error {
	// 判断用户存不存在
	var user *models.User
	err := common.DB.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return errors.New("用户不存在")
	}
	common.DB.Delete(&user)
	return nil
}
