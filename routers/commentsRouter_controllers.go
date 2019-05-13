package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["cart/controllers:CartController"] = append(beego.GlobalControllerRouter["cart/controllers:CartController"],
        beego.ControllerComments{
            Method: "Decr",
            Router: `/v1/cart/decr`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cart/controllers:CartController"] = append(beego.GlobalControllerRouter["cart/controllers:CartController"],
        beego.ControllerComments{
            Method: "Drop",
            Router: `/v1/cart/drop`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cart/controllers:CartController"] = append(beego.GlobalControllerRouter["cart/controllers:CartController"],
        beego.ControllerComments{
            Method: "Incr",
            Router: `/v1/cart/incr`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cart/controllers:CartController"] = append(beego.GlobalControllerRouter["cart/controllers:CartController"],
        beego.ControllerComments{
            Method: "Pull",
            Router: `/v1/cart/pull`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cart/controllers:CartController"] = append(beego.GlobalControllerRouter["cart/controllers:CartController"],
        beego.ControllerComments{
            Method: "Set",
            Router: `/v1/cart/set`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
