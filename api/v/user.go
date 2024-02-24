package v

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	p "github.com/piklen/pb/user"
	"mall_api/grpcclient"
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
	//// 解析命令行标志
	//flag.Parse()
	//// 连接到server端，此处禁用安全传输
	//conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	//if err != nil {
	//	fmt.Println("grpc连接失败！！！")
	//	log.LogrusObj.Infoln(err)
	//	return
	//}
	//defer conn.Close()
	//client := p.NewUserServiceClient(conn)
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
	fmt.Println("连接成功！！！")
	c.JSON(http.StatusOK, r)
}
