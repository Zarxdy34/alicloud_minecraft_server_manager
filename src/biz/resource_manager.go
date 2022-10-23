package biz

import (
	"context"
	"time"

	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/base_handler"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/consts"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/model"
)

type resourceManager struct {
	*base_handler.BaseHandler
	NoPlayerTimeCount map[int64]int64
}

var rsm *resourceManager

func InitResourceManager() {
	rsm = &resourceManager{
		BaseHandler:       base_handler.New(context.Background(), "ResourceManager"),
		NoPlayerTimeCount: make(map[int64]int64),
	}
	rsm.Loop()
}

func (h *resourceManager) Loop() {
	for idx, server := range GetServerConfigList() {
		if server.ShutdownAfterNoPlayerSecond == 0 {
			h.LogInfo("Server %s dont need auto stopped when idle", server.ServerName)
			continue
		}
		go func(serverID int64, server *model.MCServerInfo) {
			for {
				time.Sleep(time.Second * time.Duration(server.ShutdownAfterNoPlayerCheckIntervalSecond))
				resp, err := NewMinecraftServerManager(h.GetContext(), &model.MinecraftServerRequest{
					Type:     consts.QueryServer,
					ServerID: &serverID,
				}).Run()
				if err != nil || resp.ServerStatus == nil {
					h.LogError("Server %s get status failed", server.ServerName)
					h.NoPlayerTimeCount[serverID] = 0
					continue
				}
				serverStatus := resp.ServerStatus
				if !serverStatus.InstanceOnline || !serverStatus.ServerOnline || serverStatus.OnlinePlayerNumber == 0 {
					h.NoPlayerTimeCount[serverID] = 0
					continue
				}
				h.NoPlayerTimeCount[serverID]++
				if h.NoPlayerTimeCount[serverID]*server.ShutdownAfterNoPlayerCheckIntervalSecond >= server.ShutdownAfterNoPlayerSecond {
					h.LogInfo("Longer than %d seconds no player online, stop server", h.NoPlayerTimeCount)
					_, _ = NewMinecraftServerManager(h.GetContext(), &model.MinecraftServerRequest{
						Type:     consts.StopServer,
						ServerID: &serverID,
					}).Run()
				}
			}
		}(int64(idx), server)
	}
}
