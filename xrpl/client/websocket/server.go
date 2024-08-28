package websocket

import "github.com/Peersyst/xrpl-go/xrpl/model/client/server"

func (c *WebsocketClient) GetServerInfo(req *server.ServerInfoRequest) (*server.ServerInfoResponse, error) {
	res, err := c.SendRequest(req)
	if err != nil {
		return nil, err
	}
	var sir server.ServerInfoResponse
	err = res.GetResult(&sir)
	if err != nil {
		return nil, err
	}
	return &sir, err
}
