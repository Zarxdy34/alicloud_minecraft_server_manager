package model

import (
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/consts"
)

type MinecraftServerRequest struct {
	Type       consts.ActionType `json:"type"`
	ServerID   *int64            `json:"server_id"`
	ServerName *string           `json:"server_name"`
	Message    *string           `json:"message"`
}

type RequestServerListItem struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ServerStatus struct {
	InstanceOnline     bool     `json:"instance_online"`
	ServerOnline       bool     `json:"server_online"`
	OnlinePlayerNumber int64    `json:"online_player_number"`
	OnlinePlayerList   []string `json:"online_player_list"`
}

type MinecraftServerResponse struct {
	ServerList   []*RequestServerListItem `json:"server_list,omitempty"`
	ServerStatus *ServerStatus            `json:"server_status,omitempty"`
	Message      string                   `json:"message,omitempty"`
}
