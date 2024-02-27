package v

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/piklen/pb/user"
	"mall_api/grpcclient"
	"mall_api/pkg/log"
	"mall_api/pkg/util"
	"net/http"
	"strconv"
	"time"
)

type AddressService struct {
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

// CreateAddress 创建收货地址
func CreateAddress(c *gin.Context) {
	addressService := AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addressService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gin框架数据绑定失败！！！"})
		log.LogrusObj.Infoln(err)
		return
	}
	client := grpcclient.GetUserClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	res, err := client.UserCreateAddress(ctx, &pb.UserCreateAddressRequest{
		UserId:  strconv.Itoa(int(claim.ID)),
		Name:    c.PostForm("name"),
		Phone:   c.PostForm("phone"),
		Address: c.PostForm("address"),
	})
	if err != nil {
		fmt.Println("grpc调用失败！！！")
		log.LogrusObj.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

//// GetAddress 展示某个收货地址
//func GetAddress(c *gin.Context) {
//	addressService := service.AddressService{}
//	res := addressService.Show(c.Request.Context(), c.Query("id"))
//	c.JSON(http.StatusOK, res)
//}
//
//// ListAddress 展示收货地址
//func ListAddress(c *gin.Context) {
//	addressService := service.AddressService{}
//	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
//	if err := c.ShouldBind(&addressService); err == nil {
//		res := addressService.List(c.Request.Context(), claim.ID)
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		log.LogrusObj.Infoln(err)
//	}
//}
//
//// UpdateAddress 修改收货地址
//func UpdateAddress(c *gin.Context) {
//	addressService := service.AddressService{}
//	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
//	if err := c.ShouldBind(&addressService); err == nil {
//		res := addressService.Update(c.Request.Context(), claim.ID, c.Param("id"))
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		log.LogrusObj.Infoln(err)
//	}
//}
//
//// DeleteAddress 删除收获地址
//func DeleteAddress(c *gin.Context) {
//	addressService := service.AddressService{}
//	if err := c.ShouldBind(&addressService); err == nil {
//		res := addressService.Delete(c.Request.Context(), c.Param("id"))
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		log.LogrusObj.Infoln(err)
//	}
//}
