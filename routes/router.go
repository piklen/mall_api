package routes

import (
	"github.com/gin-gonic/gin"
	api "mall_api/api/v"
	"mall_api/middleware"
	"net/http"
)

// 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, "success")
	})
	//处理跨域
	r.Use(middleware.Cors())
	//静态页面返回
	r.StaticFS("/static", http.Dir("./static"))
	v := r.Group("api/v")
	{
		//用户基本操作
		v.POST("user/register", api.UserRegister)
		v.POST("user/login", api.UserLogin)
		authed := v.Group("/") //需要登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.POST("user/update", api.UpdateNickName)   //更新昵称
			authed.POST("user/avatar", api.UploadAvatar)     // 上传头像
			authed.POST("user/sending-email", api.SendEmail) //发送邮件
			authed.POST("user/valid-email", api.ValidEmail)  //邮箱变更修改绑定等
			// 显示金额
			authed.POST("money", api.ShowMoney)
			////商品操作
			//authed.POST("create", api.CreateProduct)
			////收藏夹操作
			//authed.POST("favorites/create", api.CreateFavorite)
			//authed.GET("favorites/show", api.ShowFavorites)
			//authed.POST("favorites/delete", api.DeleteFavorite)
			//// 收获地址操作
			//authed.POST("addresses/create", api.CreateAddress) //创建用户地址
			//authed.GET("addresses/get", api.GetAddress)        //获取某个id的地址
			//authed.GET("addresses/list", api.ListAddress)      //展示全部地址
			//authed.POST("addresses/update", api.UpdateAddress) //更新用户某一个地址id的地址
			//authed.POST("addresses/del", api.DeleteAddress)    //删除某一地址id的地址
			//
			//// 购物车
			//authed.POST("carts/create", api.CreateCart)
			//authed.GET("carts/show", api.ShowCarts)
			//authed.POST("carts/update", api.UpdateCart) //修改的主要是数量
			//authed.POST("carts/del", api.DeleteCart)    //购物车商品删除
			//
			//// 订单操作
			//authed.POST("orders/create", api.CreateOrder) //创建订单信息
			//authed.GET("orders/list", api.ListOrders)     //查询全部的订单
			//authed.GET("orders/show", api.ShowOrder)      //查询某一个订单
			//authed.POST("orders/del", api.DeleteOrder)
			//
			////支付操作
			//authed.POST("pay", api.OrderPay)
			//
			////秒杀专场
			//authed.POST("import_skill_goods", api.ImportSkillGoods) //导入商品内容进去MySQL
			//authed.POST("init_skill_goods", api.InitSkillGoods)
			//authed.POST("skill_goods", api.SkillGoods)
			//authed.POST("skill_goods/mysql", api.SkillGoodsWithMySQL)
			//authed.POST("skill_goods/redis", api.SkillGoodsWithRedis)
		}
	}
	return r
}
