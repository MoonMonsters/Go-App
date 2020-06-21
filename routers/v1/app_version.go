package v1

import (
	"Go-App/pkg/e"
	"Go-App/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetAppVersionTest(c *gin.Context) {
	util.ResponseWithJson(e.SUCCESS, "返回数据成功", c)
}
