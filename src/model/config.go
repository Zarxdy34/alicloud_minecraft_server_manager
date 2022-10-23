package model

type ServerConfig struct {
	AlicloudConfig        *CloudAuthInfo  `json:"alicloud_config"`
	MinecraftServerConfig []*MCServerInfo `json:"minecraft_server_config"`
}

type CloudAuthInfo struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

type MCServerInfo struct {
	ServerName                               string `json:"server_name"`
	RemoteServerIP                           string `json:"remote_server_ip"`
	RemoteServerPort                         string `json:"remote_server_port"`
	InstanceID                               string `json:"instance_id"`
	RconPort                                 string `json:"rcon_port"`
	RconPassword                             string `json:"rcon_password"`
	ShutdownAfterNoPlayerSecond              int64  `json:"shutdown_after_no_player_second"`
	ShutdownAfterNoPlayerCheckIntervalSecond int64  `json:"shutdown_after_no_player_check_interval_second"`
}
