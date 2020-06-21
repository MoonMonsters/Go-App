package v1

import (
	"Go-App/models"
	"Go-App/pkg/e"
	"Go-App/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UserStore(c *gin.Context) {
	mobile := c.PostForm("mobile")
	vCode := c.PostForm("code")

	validate := validation.Validation{}
	validate.Required(mobile, "Mobile").Message("手机号码有误")
	validate.Length(vCode, 4, "Code").Message("验证码格式不正确")
	// 校验失败
	if isOk := checkValidation(&validate, c); isOk == false {
		return
	}

	user, err := models.FindUserByMobile(mobile)
	if gorm.IsRecordNotFoundError(err) {
		user, err = models.CreateUser(mobile)
		if err != nil {
			util.ResponseWithJson(e.ERROR, "数据库操作失误", c)
			return
		}
	} else {
		if err != nil {
			util.ResponseWithJson(e.ERROR, "数据库操作失误", c)
			return
		}
	}

	token, err := util.GeneratorToken(user.ID, user.Mobile)
	if err != nil {
		util.ResponseWithJson(e.ERROR, "创建token失败", c)
		return
	}

	util.ResponseWithJson(e.SUCCESS, gin.H{
		"User":  user,
		"Token": token,
	}, c)
}

/**
验证提交数据的有效性
*/
func checkValidation(valid *validation.Validation, c *gin.Context) bool {
	if valid.HasErrors() {
		var errs []string
		for _, err := range valid.Errors {
			errs = append(errs, err.Message)
		}
		util.ResponseWithJson(e.INVALID_PARAMS, errs, c)
		return false
	}

	return true
}
