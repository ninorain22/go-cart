// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"cart/controllers"
	"github.com/astaxie/beego/context"
	"cart/business"
	"cart/enumerations"
	"github.com/astaxie/beego/validation"
	)

var ValidateToken = func(ctx *context.Context) {
	userId := ctx.Input.Query("userId")
	shopId := ctx.Input.Query("shopId")
	token := ctx.Input.Query("token")

	valid := validation.Validation{}
	valid.Required(userId, "userId")
	valid.Required(shopId, "shopId")
	valid.Required(token, "token")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			ctx.Output.JSON(controllers.Response{
				enumerations.PARAMS_REQUIRED,
				err.Key + " " + err.Message,
				nil},
			true,
			true)
		}
	}
	// 验证token，校验登录态
	if business.ValidateToken(userId, shopId, token) == false {
		ctx.Output.JSON(controllers.Response{enumerations.TOKEN_INVALID,"token invalid", nil}, true, true)
	}
}

func init() {
	beego.Include(&controllers.CartController{})
	beego.InsertFilter("/v1/*", beego.BeforeRouter, ValidateToken)
}
