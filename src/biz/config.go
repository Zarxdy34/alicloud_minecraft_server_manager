package biz

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Zarxdy34/alicloud_minecraft_server_manager/src/model"
)

var ServerConf *model.ServerConfig

func InitConfig() {
	b, err := os.ReadFile("./conf/config.json")
	if err != nil {
		panic(fmt.Sprintf("Open config file failed, err = %v", err))
	}

	err = json.Unmarshal(b, &ServerConf)
	if err != nil {
		panic(fmt.Sprintf("Load config failed, please check your config format"))
	}

	if ServerConf.AccessKeyId == "" || ServerConf.AccessKeySecret == "" {
		panic(fmt.Sprintf("Cannot read server access key config, please check your config file format"))
	}
}
