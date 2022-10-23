package biz

import (
	"context"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/iverly/go-mcping/mcping"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/base_handler"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/consts"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/conv"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/rcon"
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
		conf:        serverConf,
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
	case consts.QueryServerList:
		err = h.QueryServerList()
	case consts.SendMessage:
		err = h.SendMessage()
	}
	if err != nil {
		errMsg := fmt.Sprintf("Processing request failed, err = %v, request = %v", err, utils.Marshal(h.req))
		h.resp.Message = errMsg
		h.LogError(errMsg)
		return h.resp, err
	}
	return h.resp, nil
}

func (h *MinecraftServerManager) StartServer() error {
	instanceID := GetInstanceID(h.req.ServerName, h.req.ServerID)
	if instanceID == "" {
		err := fmt.Errorf("get instance_id failed, check your server_name and server_id both in request and config file")
		return err
	}

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(GetAKSK())
	client, err := ecs.NewClientWithOptions("cn-beijing", config, credential)
	if err != nil {
		h.resp.Message = "cannot build alicloud client"
		return err
	}

	request := ecs.CreateStartInstanceRequest()
	request.Scheme = "https"
	request.InstanceId = instanceID

	response, err := client.StartInstance(request)
	if err != nil {
		h.resp.Message = "request to alicloud failed"
		h.LogError("Start instance failed, err = %v", err)
		return err
	}
	h.LogInfo("response is %v", utils.Marshal(response))
	return nil
}

func (h *MinecraftServerManager) StopServer() error {
	instanceID := GetInstanceID(h.req.ServerName, h.req.ServerID)
	if instanceID == "" {
		err := fmt.Errorf("get instance_id failed, check your server_name and server_id both in request and config file")
		return err
	}

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(GetAKSK())
	client, err := ecs.NewClientWithOptions("cn-beijing", config, credential)
	if err != nil {
		h.resp.Message = "cannot build alicloud client"
		return err
	}

	request := ecs.CreateStopInstanceRequest()
	request.Scheme = "https"
	request.InstanceId = instanceID
	request.StoppedMode = "StopCharging"

	response, err := client.StopInstance(request)
	if err != nil {
		h.resp.Message = "request to alicloud failed"
		h.LogError("Stop instance failed, err = %v", err)
		return err
	}
	h.LogInfo("response is %v", utils.Marshal(response))
	return nil
}

func (h *MinecraftServerManager) QueryServer() error {
	// 查询实例状态
	server := GetServer(h.req.ServerName, h.req.ServerID)
	if server == nil {
		err := fmt.Errorf("cannot match server in config file, check your server_name and server_id both in request and config file")
		return err
	}

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(GetAKSK())
	client, err := ecs.NewClientWithOptions("cn-beijing", config, credential)
	if err != nil {
		h.resp.Message = "cannot build alicloud client"
		return err
	}

	request := ecs.CreateDescribeInstanceStatusRequest()
	request.Scheme = "https"
	request.InstanceId = &[]string{server.InstanceID}

	response, err := client.DescribeInstanceStatus(request)
	if err != nil || len(response.InstanceStatuses.InstanceStatus) == 0 ||
		response.InstanceStatuses.InstanceStatus[0].InstanceId != server.InstanceID {
		h.resp.Message = "request to alicloud failed"
		h.LogError("Stop instance failed, err = %v", err)
		return err
	}
	h.LogInfo("response is %v", utils.Marshal(response))

	// 如果服务器实例关闭，直接返回
	h.resp.ServerStatus = &model.ServerStatus{}
	if response.InstanceStatuses.InstanceStatus[0].Status != "Running" {
		h.resp.ServerStatus.InstanceOnline = false
		return nil
	}
	h.resp.ServerStatus.InstanceOnline = true

	// 查询服务器在线人数等信息
	pinger := mcping.NewPinger()
	pingResp, err := pinger.Ping(server.RemoteServerIP, uint16(conv.StrToInt64(server.RemoteServerPort)))
	h.LogInfo("Ping server resp = %s, err = %v", utils.Marshal(pingResp), nil)
	if err != nil {
		// 视作服务端未启动
		h.resp.ServerStatus.ServerOnline = false
		return nil
	}

	h.resp.ServerStatus.OnlinePlayerNumber = int64(pingResp.PlayerCount.Online)
	for _, player := range pingResp.Sample {
		h.resp.ServerStatus.OnlinePlayerList = append(h.resp.ServerStatus.OnlinePlayerList, player.Name)
	}

	return nil
}

func (h *MinecraftServerManager) QueryServerList() error {
	servers := GetServerConfigList()
	resp := []*model.RequestServerListItem{}
	for idx, server := range servers {
		resp = append(resp, &model.RequestServerListItem{
			ID:   int64(idx),
			Name: server.ServerName,
		})
	}
	h.resp.ServerList = resp
	h.LogInfo("Get server list resp = %s", utils.Marshal(resp))
	return nil
}

func (h *MinecraftServerManager) SendMessage() error {
	if h.req.Message == nil {
		return fmt.Errorf("No message set")
	}
	server := GetServer(h.req.ServerName, h.req.ServerID)
	if server == nil {
		err := fmt.Errorf("cannot match server in config file, check your server_name and server_id both in request and config file")
		return err
	}
	err := rcon.SaySomething(server, *h.req.Message)
	if err != nil {
		h.LogError("Send message failed, err = %v")
		return err
	}
	return nil
}
