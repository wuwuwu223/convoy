package convoy

import "fmt"

// GetServers 获取服务器列表
func (c *Client) GetServers() (servers []Server, err error) {
	req, err := c.buildReq("GET", "/api/application/servers", nil)
	if err != nil {
		return
	}
	var resp struct {
		Data []Server `json:"data"`
		Meta Meta     `json:"meta"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	servers = resp.Data
	return
}

// GetServer 获取服务器
func (c *Client) GetServer(uuid string) (server Server, err error) {
	req, err := c.buildReq("GET", "/api/application/servers/"+uuid, nil)
	if err != nil {
		return
	}
	var resp struct {
		Data Server `json:"data"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	server = resp.Data
	return
}

// CreateServer 创建服务器
func (c *Client) CreateServer(sreq CreateServerReq) (server Server, err error) {
	req, err := c.buildReq("POST", "/api/application/servers", sreq)
	if err != nil {
		return
	}
	var resp struct {
		Data Server `json:"data"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	if resp.ErrMsg.Message != "" {
		err = fmt.Errorf(resp.ErrMsg.Message)
		fmt.Println(resp.ErrMsg)
		return
	}
	server = resp.Data
	return
}

// SuspendServer 暂停服务器
func (c *Client) SuspendServer(uuid string) (err error) {
	req, err := c.buildReq("POST", "/api/application/servers/"+uuid+"/settings/suspend", nil)
	if err != nil {
		return
	}
	var resp struct {
		ErrMsg
	}
	err = c.doReq(req, &resp)
	if resp.ErrMsg.Message != "" {
		err = fmt.Errorf(resp.ErrMsg.Message)
		return
	}
	return
}

// UnsuspendServer 恢复服务器
func (c *Client) UnsuspendServer(uuid string) (err error) {
	req, err := c.buildReq("POST", "/api/application/servers/"+uuid+"/settings/unsuspend", nil)
	if err != nil {
		return
	}
	var resp struct {
		ErrMsg
	}
	err = c.doReq(req, &resp)
	if resp.ErrMsg.Message != "" {
		err = fmt.Errorf(resp.ErrMsg.Message)
		return
	}
	return
}

// DeleteServer 删除服务器
func (c *Client) DeleteServer(uuid string) (err error) {
	req, err := c.buildReq("DELETE", "/api/application/servers/"+uuid, nil)
	if err != nil {
		return
	}
	var resp struct {
		ErrMsg
	}
	err = c.doReq(req, &resp)
	if resp.ErrMsg.Message != "" {
		err = fmt.Errorf(resp.ErrMsg.Message)
		return
	}
	return
}
