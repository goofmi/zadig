package login

import (
	"github.com/gin-gonic/gin"
	"github.com/koderover/zadig/pkg/microservice/user/core/service/login"
	internalhandler "github.com/koderover/zadig/pkg/shared/handler"
)

func InternalLogin(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()
	args := &login.LoginArgs{}
	if err := c.ShouldBindJSON(args); err != nil {
		ctx.Err = err
		return
	}
	ctx.Resp, ctx.Err = login.InternalLogin(args, ctx.Logger)
}
