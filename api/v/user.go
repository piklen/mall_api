package v

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	p "github.com/piklen/pb/user"
	"io"
	"mall_api/grpcclient"
	"mall_api/pkg/log"
	"mall_api/pkg/util"
	"net/http"
	"strconv"
	"time"
)

type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行验证
}
type SendEmailService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	// OperationType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}
type ValidEmailService struct {
}
type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func UserRegister(c *gin.Context) {
	var userRegister UserService
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, "gin框架数据绑定失败！！！") //绑定不成功返回错误
		log.LogrusObj.Infoln(err)
		return
	}
	//创建一个user的grpc
	client := grpcclient.GetUserClient()
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	nickname := c.PostForm("nick_name")
	username := c.PostForm("user_name")
	password := c.PostForm("password")
	key := c.PostForm("key")
	r, err := client.RegisterUser(ctx, &p.UserRegisterRequest{
		NickName: nickname,
		UserName: username,
		Password: password,
		Key:      key,
	})
	if err != nil {
		fmt.Println("grpc调用失败！！！")
		log.LogrusObj.Infoln(err)
		return
	}
	fmt.Println("连接成功！！！")
	c.JSON(http.StatusOK, r)
}
func UserLogin(c *gin.Context) {
	var userRegister UserService
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gin框架数据绑定失败！！！"})
		log.LogrusObj.Infoln(err)
		return
	}
	client := grpcclient.GetUserClient()
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	// 从表单中获取用户名和密码
	username := c.PostForm("user_name")
	password := c.PostForm("password")
	r, err := client.UserLogin(ctx, &p.UserRegisterRequest{
		UserName: username,
		Password: password,
	})
	if err != nil {
		fmt.Println("grpc调用失败！！！")
		log.LogrusObj.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, r)
}
func UpdateNickName(c *gin.Context) {
	var UpdateNickNameService UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&UpdateNickNameService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gin框架数据绑定失败！！！"})
		log.LogrusObj.Infoln(err)
		return
	}
	client := grpcclient.GetUserClient()
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	nickname := c.PostForm("nick_name")
	r, err := client.UpdateNickName(ctx, &p.UpdateNickNameRequest{
		UserId:   strconv.Itoa(int(claims.ID)),
		NickName: nickname,
	})
	if err != nil {
		fmt.Println("grpc调用失败！！！")
		log.LogrusObj.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, r)
}

// UploadAvatar 上传头像
func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	content, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "上传头像文件编译失败！！！",
		})
		return
	}
	fileSize := fileHeader.Size
	uploadAvatarService := UserService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatarService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gin框架数据绑定失败！！！"})
		log.LogrusObj.Infoln(err)
		return
	}
	client := grpcclient.GetUserClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	res, err := client.UploadAvatar(ctx, &p.UploadAvatarRequest{
		UserId:   strconv.Itoa(int(claims.ID)),
		FileData: content,
		FileSize: fileSize,
	})
	if err != nil {
		fmt.Println("grpc调用失败！！！")
		log.LogrusObj.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// SendEmail 发送邮件
func SendEmail(c *gin.Context) {
	var sendEmailService SendEmailService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmailService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gin框架数据绑定失败！！！"})
		log.LogrusObj.Infoln(err)
		return
	}
	client := grpcclient.GetUserClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	operationType := c.PostForm("operation_type")
	email := c.PostForm("email")
	password := c.PostForm("password")
	res, err := client.SendEmail(ctx, &p.SendEmailRequest{
		UserId:        strconv.Itoa(int(claims.ID)),
		OperationType: operationType,
		Email:         email,
		Password:      password,
	})
	if err != nil {
		fmt.Println("grpc调用失败！！！")
		log.LogrusObj.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
func ValidEmail(c *gin.Context) {
	var validEmailService ValidEmailService
	if err := c.ShouldBind(validEmailService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gin框架数据绑定失败！！！"})
		log.LogrusObj.Infoln(err)
		return
	}
	client := grpcclient.GetUserClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	res, err := client.ValidEmail(ctx, &p.ValidEmailRequest{
		Token: c.GetHeader("Authorization"),
	})
	if err != nil {
		fmt.Println("grpc调用失败！！！")
		log.LogrusObj.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
func ShowMoney(c *gin.Context) {
	showMoneyService := ShowMoneyService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoneyService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gin框架数据绑定失败！！！"})
		log.LogrusObj.Infoln(err)
		return
	}
	client := grpcclient.GetUserClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	key := c.PostForm("key")
	res, err := client.ShowMoney(ctx, &p.ShowMoneyRequest{
		UserId: strconv.Itoa(int(claims.ID)),
		Key:    key,
	})
	if err != nil {
		fmt.Println("grpc调用失败！！！")
		log.LogrusObj.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
