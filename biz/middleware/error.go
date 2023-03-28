package middleware

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
)

func ErrorHandler(ctx context.Context, c *app.RequestContext) {
	c.Next(ctx)
	if len(c.Errors) == 0 {
		return
	}
	hertzErr := c.Errors[0]
	err := hertzErr.Unwrap()
	log.Printf("%+v", err)
	err = errors.Unwrap(err)
	c.JSON(consts.StatusBadRequest, utils.H{
		"code":    consts.StatusBadRequest,
		"message": err.Error(),
		"data":    "",
	})
}
