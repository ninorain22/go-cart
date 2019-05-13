package controllers

import (
	"github.com/astaxie/beego"
	)

type BaseController struct {
	beego.Controller
}

type Response struct {
	Code int					`json:"code"`
	Msg string					`json:"msg"`
	Data interface{}			`json:"data"`
}

func (this *BaseController) ServeJSON(encoding ...bool) {
	var (
		hasIndent   = beego.BConfig.RunMode != beego.PROD
		hasEncoding = len(encoding) > 0 && encoding[0]
	)

	response := Response{0, "ok", this.Data["json"]}
	this.Ctx.Output.JSON(response, hasIndent, hasEncoding)
}
