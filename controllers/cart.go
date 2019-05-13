package controllers

import (
	"encoding/json"
	"cart/models"
)

type CartController struct {
	BaseController
}


// @router /v1/cart/pull [get]
func (this *CartController) Pull() {
	userId := this.GetString("userId")
	shopId := this.GetString("shopId")
	cart := models.GetCart(userId, shopId)
	this.Data["json"] = cart.Summary()
	this.ServeJSON()
}

// @router /v1/cart/incr [post]
func (this *CartController) Incr() {
	userId := this.GetString("userId")
	shopId := this.GetString("shopId")
	var i models.SkuItem
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &i); err != nil {
		this.Data["json"] = err.Error()
		this.StopRun()
	}
	cart := models.GetCart(userId, shopId)
	cart.Incr(i.SkuId, i.SkuNum)
	this.Data["json"] = cart.Summary()
	this.ServeJSON()
}

// @router /v1/cart/decr [post]
func (this *CartController) Decr() {
	userId := this.GetString("userId")
	shopId := this.GetString("shopId")
	var i models.SkuItem
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &i); err != nil {
		this.Data["json"] = err.Error()
		this.StopRun()
	}
	cart := models.GetCart(userId, shopId)
	cart.Decr(i.SkuId, i.SkuNum)
	this.Data["json"] = cart.Summary()
	this.ServeJSON()
}

// @router /v1/cart/drop [post]
func (this *CartController) Drop() {
	userId := this.GetString("userId")
	shopId := this.GetString("shopId")
	var i models.SkuItem
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &i); err != nil {
		this.Data["json"] = err.Error()
		this.StopRun()
	}
	cart := models.GetCart(userId, shopId)
	cart.Drop(i.SkuId)
	this.Data["json"] = cart.Summary()
	this.ServeJSON()
}

// @router /v1/cart/set [post]
func (this *CartController) Set() {
	userId := this.GetString("userId")
	shopId := this.GetString("shopId")
	var i models.SkuItem
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &i); err != nil {
		this.Data["json"] = err.Error()
		this.StopRun()
	}
	cart := models.GetCart(userId, shopId)
	cart.Drop(i.SkuId)
	this.Data["json"] = cart.Summary()
	this.ServeJSON()
}