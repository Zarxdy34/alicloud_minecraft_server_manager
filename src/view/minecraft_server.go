package view

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/biz"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/utils"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/model"
)

// MinecraftServerManage
func MinecraftServerManage(ctx context.Context, c *app.RequestContext) {
	req := &model.MinecraftServerRequest{}
	err := c.Bind(&req)
	if err != nil {
		logger.CtxInfof(ctx, "Parse request failed")
		logger.CtxDebugf(ctx, "Request = %v", utils.Marshal(c))
	}
	resp, _ := biz.NewMinecraftServerManager(ctx, req).Run()
	c.JSON(200, resp)
	return
}
