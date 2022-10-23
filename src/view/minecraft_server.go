package view

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/logger"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/biz"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/utils"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/model"
)

// MinecraftServerManage
func MinecraftServerManage(ctx context.Context, c *app.RequestContext) {
	logger.Infof("Received request = %v, uri = %v, path = %v, body = %v", utils.Marshal(c), c.Request.URI(), string(c.Request.Path()), string(c.Request.Body()))
	req := &model.MinecraftServerRequest{}
	err := c.Bind(&req)
	if err != nil {
		logger.Infof("Parse request failed")
		logger.Infof("Request = %v", utils.Marshal(c))
	}
	resp, _ := biz.NewMinecraftServerManager(ctx, req).Run()
	c.JSON(200, resp)
	return
}
