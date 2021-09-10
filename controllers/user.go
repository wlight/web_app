package controllers

import (
	"net/http"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SignUpHandler(c *gin.Context) {
	// 获取参数和校验
	var p  = new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}


		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 手动校验参数不能为空
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	// 	zap.L().Error("SignUp with invalid param")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg":"请求参数有误",
	// 	})
	// 	return
	// }

	
	// 注册逻辑
	logic.SignUp(p)

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
