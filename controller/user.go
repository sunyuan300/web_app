package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"web_app/logic"
	"web_app/models"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.参数获取和参数校验
	p := new(models.ParamSign)
	// ShouldBindJSON 只能用于校验字段类型和格式对不对
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误,直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		// 判断 err 是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg:": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg:": removeTopStruct(errs.Translate(trans)), // 翻译错误
		})
		return
	}
	// 手动对请求参数进行详细的业务规则校验
	//if len(p.UserName) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg:": "请求参数有误",
	//	})
	//	return
	//}

	// 2.业务处理
	logic.SignUp(p)
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg:": "success",
	})
}
