package v1

import (
	"strings"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/conf"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/ec"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/reply"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/service"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/util"
	"github.com/gin-gonic/gin"
)

// OnCustomerRegister handle customer's register request
func OnCustomerRegister(ctx *gin.Context) {
	cus := new(model.Customer)
	if err := ctx.ShouldBind(cus); err != nil {
		log.Errorf("fail to register customer. %v", err)
		reply.Error(ctx, err)
		return
	}
	if err := service.CreateCustomer(cus); err != nil {
		log.Errorf("fail to register customer. %v", err)
		reply.Error(ctx, err)
		return
	}
	if cus.ID != 0 {
		token, err := util.GenToken(cus.ID)
		if err != nil {
			log.Errorf("customer successfully registered, but fail to generate token. %v", err)
			reply.Error(ctx, err)
			return
		}
		reply.OK(ctx, &reply.UserAuth{
			Token: token,
			Name:  cus.Name,
		})
		return
	}
	reply.OK(ctx, nil)
}

// OnCustomerAuth handle customer's login request
func OnCustomerAuth(ctx *gin.Context) {
	cus := new(model.Customer)
	if err := ctx.ShouldBind(cus); err != nil {
		log.Errorf("fail to auth customer. %v", err)
		reply.Error(ctx, err)
		return
	}
	if cus.Email == "" ||
		!strings.ContainsAny(cus.Email, "@") ||
		cus.Password == "" ||
		len(cus.Password) < 6 {
		log.Errorf("fail to auth cusomter. %v", ec.ErrInvalidFormat)
		reply.Error(ctx, ec.ErrInvalidFormat)
		return
	}
	authCus, err := service.CustomerAuth(cus.Email, cus.Password)
	if err != nil {
		log.Errorf("fail to auth customer. %v", err)
		reply.Error(ctx, err)
		return
	}

	token, err := util.GenToken(authCus.ID)
	if err != nil {
		log.Errorf("fail to auth customer. %v", err)
		reply.Error(ctx, err)
		return
	}
	reply.OK(ctx, &reply.UserAuth{
		Token: token,
		Name:  authCus.Name,
	})
}

// OnCustomerGetProfile handle customer get profile
func OnCustomerGetProfile(ctx *gin.Context) {
	var (
		id  uint
		err error
	)
	id = ctx.GetUint(conf.StrUserID)
	if id == 0 {
		log.Errorf("fail to get customer profile. %v", err)
		reply.Error(ctx, ec.ErrUnauthorized)
	}

	cus, err := service.FindCustomerByID(id)
	if err != nil {
		log.Errorf("fail to get customer profile. %v", err)
		reply.Error(ctx, err)
		return
	}
	cus.Password = ""
	reply.OK(ctx, cus)
}

// OnCustomerEditProfile handler customer edit their profile
func OnCustomerEditProfile(ctx *gin.Context) {
	cus := new(model.Customer)
	if err := ctx.ShouldBind(cus); err != nil {
		log.Errorf("fail to edit customer profile. %v", err)
		reply.Error(ctx, err)
		return
	}
	if err := service.UpdateCustomer(cus); err != nil {
		log.Errorf("fail to edit customer profile. %v", err)
		reply.Error(ctx, err)
		return
	}
	reply.OK(ctx, nil)
}
