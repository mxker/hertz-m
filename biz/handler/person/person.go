package person

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/requestid"
	"github.com/pkg/errors"
)

// GetInfo .
func GetInfo(ctx context.Context, c *app.RequestContext) {
	// 获取query参数
	name := c.Query("name")
	// 获取路径参数
	age := c.Param("age")
	c.JSON(consts.StatusOK, utils.H{
		"data": &map[string]interface{}{
			"name": name,
			"age":  age,
		},
	})
}

type Userinfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// PostInfo .
func PostInfo(ctx context.Context, c *app.RequestContext) {
	body, err := c.Body()
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    consts.StatusBadRequest,
			"message": "request body is err",
		})
	}
	var input Userinfo
	if err := json.Unmarshal(body, &input); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"code":    consts.StatusBadRequest,
			"message": "json unmarshal err",
		})
	}

	c.JSON(consts.StatusOK, utils.H{
		"code": consts.StatusOK,
		"data": &map[string]interface{}{
			"name": input.Name,
			"age":  input.Age,
		},
	})
}

type BindReq struct {
	Name string `query:"name",json:"name"`
	Age  int    `path:"age"json:"age"`
	Sex  string `json:"sex"`
}

func BindInfo(ctx context.Context, c *app.RequestContext) {
	var bindReq BindReq
	if err := c.BindAndValidate(&bindReq); err != nil {
		_ = c.Error(errors.WithStack(err))
		return
	}

	c.JSON(consts.StatusOK, utils.H{
		"code": consts.StatusOK,
		"data": &map[string]interface{}{
			"name":      bindReq.Name,
			"age":       bindReq.Age,
			"sex":       bindReq.Sex,
			"requestId": requestid.Get(c),
		},
		"message": "success",
	})
}
