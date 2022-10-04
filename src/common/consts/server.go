package consts

type ActionType string

const (
	StartServer ActionType = "start_server"
	StopServer  ActionType = "stop_server"
	QueryServer ActionType = "query_server"
)
