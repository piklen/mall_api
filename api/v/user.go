package v

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	p "github.com/piklen/pb"
	"google.golang.org/grpc"
	"mall_api/pkg/log"
	"net/http"
	"time"
)

type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行验证
}

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
)

func UserRegister(c *gin.Context) {
	var userRegister UserService
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, "绑定失败！！！") //绑定不成功返回错误
		log.LogrusObj.Infoln(err)
		return
	}
	flag.Parse()
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败！！！")
		log.LogrusObj.Infoln(err)
		return
	}
	defer conn.Close()
	client := p.NewUserServiceClient(conn)
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	nickname := c.PostForm("nick_name")
	username := c.PostForm("user_name")
	password := c.PostForm("password")
	key := c.PostForm("key")
	r, err := client.UpdateName(ctx, &p.UserRegister{
		NickName: nickname,
		UserName: username,
		Password: password,
		Key:      key,
	})
	if err != nil {
		fmt.Println("调用失败！！！")
		log.LogrusObj.Infoln(err)
		return
	}
	fmt.Println("连接成功！！！")
	c.JSON(http.StatusOK, r)
}

//func BotchUserRegister(c *gin.Context) {
//	var userRegister service.BatchUsersService
//	if err := c.ShouldBind(&userRegister); err == nil {
//		res := userRegister.BatchRegister(c.Request.Context())
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err)) //绑定不成功返回错误
//		log.LogrusObj.Infoln(err)
//	}
//}
//
//// UserLogin 用户登陆接口
//func UserLogin(c *gin.Context) {
//	var userLogin service.UserService
//	if err := c.ShouldBind(&userLogin); err == nil {
//		res := userLogin.Login(c.Request.Context())
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		log.LogrusObj.Infoln(err)
//	}
//}
//func UserUpdate(c *gin.Context) {
//	var userUpdateService service.UserService
//	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
//	if err := c.ShouldBind(&userUpdateService); err == nil {
//		res := userUpdateService.Update(c.Request.Context(), claims.ID)
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		log.LogrusObj.Infoln(err)
//	}
//}
//
//// UploadAvatar 上传头像
//func UploadAvatar(c *gin.Context) {
//	file, fileHeader, _ := c.Request.FormFile("file")
//	fileSize := fileHeader.Size
//	uploadAvatarService := service.UserService{}
//	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
//	if err := c.ShouldBind(&uploadAvatarService); err == nil {
//		res := uploadAvatarService.PostAvatar(c.Request.Context(), chaim.ID, file, fileSize)
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		log.LogrusObj.Infoln(err)
//	}
//}
//
//// SendEmail 发送邮件
//func SendEmail(c *gin.Context) {
//	var sendEmailService service.SendEmailService
//	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
//	if err := c.ShouldBind(&sendEmailService); err == nil {
//		res := sendEmailService.Send(c.Request.Context(), chaim.ID)
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		log.LogrusObj.Infoln(err)
//	}
//}
//func ValidEmail(c *gin.Context) {
//	var validEmailService service.ValidEmailService
//	if err := c.ShouldBind(validEmailService); err == nil {
//		res := validEmailService.Valid(c.Request.Context(), c.GetHeader("Authorization"))
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		log.LogrusObj.Infoln(err)
//	}
//}
//func ShowMoney(c *gin.Context) {
//	showMoneyService := service.ShowMoneyService{}
//	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
//	if err := c.ShouldBind(&showMoneyService); err == nil {
//		res := showMoneyService.Show(c.Request.Context(), claim.ID)
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		log.LogrusObj.Infoln(err)
//	}
//}

//
//func UserRegisterHandler() gin.HandlerFunc {
//
//	return func(ctx *gin.Context) {
//		userDao := dao.NewUserDao(ctx.Request.Context())
//		var req types.UserRegisterReq
//		if err := ctx.ShouldBind(&req); err != nil {
//			return
//		}
//		user := &model.User{
//			NickName: req.NickName,
//			UserName: req.UserName,
//			Status:   "active",
//			Money:    "10000",
//		}
//		userDao.CreateUser(user)
//	}
//}
