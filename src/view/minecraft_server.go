package view

import (
	"context"

	"github.com/Zarxdy34/alicloud_minecraft_server_manager/src/biz"
	"github.com/Zarxdy34/alicloud_minecraft_server_manager/src/common/utils"
	"github.com/Zarxdy34/alicloud_minecraft_server_manager/src/model"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
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
