# alicloud_minecraft_server_manager


## 1. Put your access key config at ./conf/config.json

    {
        "alicloud_config": {
            "access_key_id": "",
            "access_key_secret": ""
        },
        "minecraft_server_config": [
            {
                "server_name": "测试用服务器",
                "remote_server_ip": "",
                "remote_server_port": "",
                "instance_id": "",
                "rcon_port": "",
                "rcon_password": "",
                "shutdown_after_no_player_second": 60,
                "shutdown_after_no_player_check_interval_second": 20
            }
        ]
    }

## 2. Run build.sh, then run.sh