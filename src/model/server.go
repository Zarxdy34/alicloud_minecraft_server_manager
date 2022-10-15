package model

import "github.com/zarxdy34/alicloud_minecraft_server_manager/src/common/consts"

type MinecraftServerRequest struct {
	Type       consts.ActionType `json:"type"`
	InstanceId string            `json:"instance_id"`
}

type MinecraftServerResponse struct {
	Message string `json:"message,omitempty"`
}
