package v1

import (
	"github.com/gin-gonic/gin"
)

//登录 注册接口
func UserStore(c *gin.Context) {
	mobile := c.PostForm("mobile")
	vCode := c.PostForm("code")

	// util.ResponseWithJson(e.SUCCESS, gin.H{
	// 	"Mobile": mobile,
	// 	"Code":   vCode,
	// }, c)

	//对请求对参数进行验证
	validate := validation.validation{}
	validate.Required(mobile, "Mobile").Message("手机号有误")
	validate.Length(vCode, 4, "Code").Message("验证码格式不正确")
}
