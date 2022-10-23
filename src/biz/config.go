package biz

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/logger"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/model"
)

var serverConf *model.ServerConfig

func readConfig(safeMode bool) {
	b, err := os.ReadFile("./conf/config.json")
	conf := &model.ServerConfig{}
	if err != nil {
		errMsg := fmt.Sprintf("Open config file failed, err = %v", err)
		if !safeMode {
			panic(errMsg)
		} else {
			logger.Errorf(errMsg)
		}
	}

	err = json.Unmarshal(b, &conf)
	if err != nil {
		errMsg := fmt.Sprintf("Load config file failed, please check your config file, err = %v", err)
		if !safeMode {
			panic(errMsg)
		} else {
			logger.Errorf(errMsg)
		}
	}

	if conf.AlicloudConfig == nil || conf.AlicloudConfig.AccessKeyId == "" || conf.AlicloudConfig.AccessKeySecret == "" {
		errMsg := fmt.Sprintf("Cannot read server access key config, please check your config file")
		if !safeMode {
			panic(errMsg)
		} else {
			logger.Errorf(errMsg)
		}
	}

	serverConf = conf
}

func InitConfig() {
	readConfig(false)
}

func GetAKSK() (ak string, sk string) {
	return serverConf.AlicloudConfig.AccessKeyId, serverConf.AlicloudConfig.AccessKeySecret
}

func GetServerByServerName(serverName string) *model.MCServerInfo {
	for _, server := range serverConf.MinecraftServerConfig {
		if server.ServerName == serverName {
			return server
		}
	}
	return nil
}

func GetServerByServerID(id int64) *model.MCServerInfo {
	if len(serverConf.MinecraftServerConfig) > int(id) {
		return serverConf.MinecraftServerConfig[id]
	}
	return nil
}

func GetServer(serverName *string, serverID *int64) *model.MCServerInfo {
	var resp *model.MCServerInfo
	if serverID != nil {
		resp = GetServerByServerID(*serverID)
		if resp != nil {
			return resp
		}
	}
	if serverName != nil {
		for _, server := range serverConf.MinecraftServerConfig {
			if server.ServerName == *serverName {
				return server
			}
		}
	}
	return nil
}

func GetServerConfigList() []*model.MCServerInfo {
	return serverConf.MinecraftServerConfig
}

func GetInstanceID(serverName *string, serverID *int64) string {
	if serverID != nil && len(serverConf.MinecraftServerConfig) > int(*serverID) {
		return serverConf.MinecraftServerConfig[*serverID].InstanceID
	}
	for _, server := range serverConf.MinecraftServerConfig {
		if serverName != nil && server.ServerName == *serverName {
			return server.InstanceID
		}
	}
	return ""
}

func GetServerByInstanceID(instanceID string) *model.MCServerInfo {
	for _, server := range serverConf.MinecraftServerConfig {
		if server.InstanceID == instanceID {
			return server
		}
	}
	return nil
}
