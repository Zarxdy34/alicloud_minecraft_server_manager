package rcon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/logger"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/model"
)

func baseCommand(server *model.MCServerInfo) []string {
	cmd := fmt.Sprintf("-H %s -P %s -p %s", server.RemoteServerIP, server.RconPort, server.RconPassword)
	return strings.Split(cmd, " ")
}

func SaySomething(server *model.MCServerInfo, msg string) error {
	cmd := baseCommand(server)
	cmd = append(cmd, fmt.Sprintf("\"say %s\"", msg))
	resp, err := exec.Command("mcrcon", cmd...).Output()
	logger.Infof("Exec command %s, resp = %v", strings.Join(cmd, " "), resp)
	if err != nil {
		return fmt.Errorf("Run mcron failed, err = %v", err)
	}
	return nil
}
