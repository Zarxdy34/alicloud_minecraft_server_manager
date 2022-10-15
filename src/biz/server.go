package biz

import (
	"context"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/base_handler"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/consts"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/utils"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/model"
)

type MinecraftServerManager struct {
	*base_handler.BaseHandler
	req  *model.MinecraftServerRequest
	resp *model.MinecraftServerResponse
	conf *model.ServerConfig
}

func NewMinecraftServerManager(ctx context.Context, req *model.MinecraftServerRequest) *MinecraftServerManager {
	return &MinecraftServerManager{
		BaseHandler: base_handler.New(ctx, "MinecraftServerManager"),
		req:         req,
		resp:        &model.MinecraftServerResponse{},
		conf:        ServerConf,
	}
}

func (h *MinecraftServerManager) Run() (*model.MinecraftServerResponse, error) {
	var err error
	switch h.req.Type {
	case consts.StartServer:
		err = h.StartServer()
	case consts.StopServer:
		err = h.StopServer()
	case consts.QueryServer:
		err = h.QueryServer()
	}
	if err != nil {
		h.LogError("Processing request failed, err = %v, request = %v", err, utils.Marshal(h.req))
		return h.resp, err
	}
	return h.resp, nil
}

func (h *MinecraftServerManager) StartServer() error {
	if h.req.InstanceId == "" {
		h.resp.Message = "missing instance id"
		return fmt.Errorf("missing instance id")
	}

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(h.conf.AccessKeyId, h.conf.AccessKeySecret)
	client, err := ecs.NewClientWithOptions("cn-beijing", config, credential)
	if err != nil {
		panic(err)
	}

	request := ecs.CreateStartInstanceRequest()
	request.Scheme = "https"
	request.InstanceId = h.req.InstanceId

	response, err := client.StartInstance(request)
	if err != nil {
		h.resp.Message = "request to alicloud failed, see more details in logs"
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %v\n", utils.Marshal(response))
	return nil
}

func (h *MinecraftServerManager) StopServer() error {
	if h.req.InstanceId == "" {
		h.resp.Message = "missing instance id"
		return fmt.Errorf("missing instance id")
	}

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(h.conf.AccessKeyId, h.conf.AccessKeySecret)
	client, err := ecs.NewClientWithOptions("cn-beijing", config, credential)
	if err != nil {
		panic(err)
	}

	request := ecs.CreateStopInstanceRequest()
	request.Scheme = "https"
	request.InstanceId = h.req.InstanceId
	request.StoppedMode = "StopCharging"

	response, err := client.StopInstance(request)
	if err != nil {
		h.resp.Message = "request to alicloud failed, see more details in logs"
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %v\n", utils.Marshal(response))
	return nil
}

func (h *MinecraftServerManager) QueryServer() error {
	return nil
}
