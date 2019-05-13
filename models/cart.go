package models

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
	"cart/manager"
)

const (
	// 存储以加入购物车时间倒叙的商品列表
	SKU_SET_KEY = "skuSet:{%s}:{%s}"
	// 存储购物车中商品的数量，选中状态等信息
	SKU_ITEM_KEY = "skuItem:{%s}:{%s}:{%s}"

	CHOOSE	= "choose"
	SKU_NUM	= "skuNum"
)

type Cart struct {
	UserId string			// 用户ID
	ShopId string			// 店铺ID
	SkuIdList []string		// 购物车中skuId列表
	SkuList []SkuItem		// 购物车中sku详情列表
}

type SkuItem struct {
	SkuId string	`json:"skuId"`
	SkuNum int		`json:"skuNum"`
	Choose bool		`json:"choose"`
}

// 获取购物车
func GetCart(userId, shopId string) *Cart {
	return &Cart{userId, shopId, nil, nil}
}

// 获取redis中购物车商品列表存储key
func (this *Cart) getSkuSetKey() string {
	return fmt.Sprintf(SKU_SET_KEY, this.UserId, this.ShopId)
}

// 获取redis中购物车商品详情存储key
func (this *Cart) getSkuItemKey(skuId string) string {
	return fmt.Sprintf(SKU_ITEM_KEY, this.UserId, this.ShopId, skuId)
}

// 获取redis中购物车某个商品详情(数量、选中状态等)
func (this *Cart) getSkuItem(skuId string) map[string]string {
	c := manager.RedisClient.Get()
	defer c.Close()
	skuItem, _ := redis.StringMap(c.Do("hgetall", this.getSkuItemKey(skuId)))
	return skuItem
}

// 获取以加入购物车时间倒叙的skuId列表
func (this *Cart) getSkuIdList() []string {
	if this.SkuIdList == nil {
		c := manager.RedisClient.Get()
		defer c.Close()
		this.SkuIdList, _ = redis.Strings(c.Do("zrevrange", this.getSkuSetKey(), 0, -1))
	}
	return this.SkuIdList
}

// 获取以加入购物车时间倒叙的sku详情列表
func (this *Cart) getSkuList() []SkuItem {
	if this.SkuList == nil {
		this.SkuList = []SkuItem{}
		skuIdList := this.getSkuIdList()
		for _, skuId := range skuIdList {
			// 获取每个sku的数量, 选中情况
			skuItem := this.getSkuItem(skuId)
			skuNum, _ := strconv.Atoi(skuItem[SKU_NUM])
			choose, _ := strconv.ParseBool(skuItem[CHOOSE])
			// todo: 调用商品系统获取商品其他属性，比如售价、快照等等
			this.SkuList = append(this.SkuList, SkuItem{skuId, skuNum, choose})
		}
	}
	return this.SkuList
}

// 获取购物车中商品总数
func (this *Cart) getTotalNum() int {
	skuList := this.getSkuList()
	num := 0
	for _, skuItem := range skuList {
		num += skuItem.SkuNum
	}
	return num
}

// 获取购物车中选中商品总数
func (this *Cart) getChooseNum() int {
	skuList := this.getSkuList()
	num := 0
	for _, skuItem := range skuList {
		if skuItem.Choose {
			num += skuItem.SkuNum
		}
	}
	return num
}

// 购物车内商品汇总信息
func (this *Cart) getSummary() map[string]int {
	return map[string]int{
		"totalNum": this.getTotalNum(),
		"chooseNum": this.getChooseNum(),
	}
}

// 判断商品是否在购物车中
func (this *Cart) isSkuInCart(skuId string) bool {
	c := manager.RedisClient.Get()
	defer c.Close()
	if _, ok := redis.Int(c.Do("zscore", this.getSkuSetKey(), skuId)); ok != nil {
		return false
	}
	return true
}

// 购物车详情
func (this *Cart) Summary() map[string]interface{} {
	return map[string]interface{}{
		"cart": this.getSkuList(),
		"summary": this.getSummary(),
	}
}

// 增加num个skuId商品
func (this *Cart) Incr(skuId string, num int) {
	c := manager.RedisClient.Get()
	defer c.Close()
	if !this.isSkuInCart(skuId) {
		c.Do("zadd", this.getSkuSetKey(), time.Now().Unix(), skuId)
		// 默认选择
		c.Do("hset", this.getSkuItemKey(skuId), CHOOSE, 1)
	}
	c.Do("hincrby", this.getSkuItemKey(skuId), SKU_NUM, num)
}

// 减少num个skuId商品
func (this *Cart) Decr(skuId string, num int) {
	c := manager.RedisClient.Get()
	defer c.Close()
	if this.isSkuInCart(skuId) {
		// 将减少操作和删除操作区分，最少减少到1
		if skuNum, _ := redis.Int(c.Do("hincrby", this.getSkuItemKey(skuId), SKU_NUM, -num)); skuNum < 1 {
			c.Do("hset", this.getSkuItemKey(skuId), SKU_NUM, 1)
		}
	}
}

// 删除某个sku
func (this *Cart) Drop(skuId string) {
	c := manager.RedisClient.Get()
	defer c.Close()
	if this.isSkuInCart(skuId) {
		c.Do("zrem", this.getSkuSetKey(), []string{skuId})
		c.Do("hdel", this.getSkuItemKey(skuId), []string{SKU_NUM, CHOOSE})
	}
}

// 设置某个商品的属性(数量、选中等)
func (this *Cart) Set(skuId string, num int, choose bool) {
	c := manager.RedisClient.Get()
	defer c.Close()
	if !this.isSkuInCart(skuId) {
		c.Do("zadd", this.getSkuSetKey(), time.Now().Unix(), skuId)
	}
	c.Do("hset", this.getSkuItemKey(skuId), SKU_NUM, num)
	c.Do("hset", this.getSkuItemKey(skuId), CHOOSE, choose)
}




