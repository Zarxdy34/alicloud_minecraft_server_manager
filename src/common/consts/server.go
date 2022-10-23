package consts

type ActionType string

const (
	StartServer     ActionType = "start_server"
	StopServer      ActionType = "stop_server"
	QueryServer     ActionType = "query_server"
	QueryServerList ActionType = "query_server_list"
	SendMessage     ActionType = "send_message"
)
