package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"keep-bill/db"
	"keep-bill/pkg/e"
	"keep-bill/utils"
	"log"
	"net/http"
)

type register struct {
	// 注册
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Register(c *gin.Context) {
	/**
	用户注册
	Username 用户名
	Password 密码
	*/

	var (
		code int
		reg  register
	)

	_ = c.BindJSON(&reg)
	valid := validation.Validation{}
	valid.Required(reg.Username, "username").Message("用户名不能为空")
	valid.Required(reg.Password, "password").Message("密码不能为空")

	if valid.HasErrors() {
		code = e.InvalidParams

		// 打印错误日志
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	} else {
		params := map[string]interface{}{
			"username": reg.Username,
		}

		if db.ExistUser(params) {
			code = e.RepeatUser
		} else {
			user := db.User{
				Username: reg.Username,
				Password: utils.Md5(reg.Password),
			}

			if db.CreateUser(user) {
				code = e.Success
			} else {
				code = e.ErrorDataBase
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"user": reg,
	})
}
