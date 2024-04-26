package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"k8s-demo/service"
	"net/http"
)

var AllRes allRes

type allRes struct{}

func(*allRes) GetAllNum(ctx *gin.Context) {
	params := new(struct {
		Cluster     string `form:"cluster"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	data, errs := service.AllRes.GetAllNum(client)
	if len(errs) > 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": errs,
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "获取资源数量成功",
		"data": data,
	})
}