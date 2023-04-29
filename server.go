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

type UpdateServerReq struct {
	AddressIds     []int       `json:"address_ids"`
	SnapshotLimit  int         `json:"snapshot_limit"`
	BackupLimit    interface{} `json:"backup_limit"`
	BandwidthLimit interface{} `json:"bandwidth_limit"`
	BandwidthUsage int         `json:"bandwidth_usage"`
	Cpu            int         `json:"cpu"`
	Memory         int64       `json:"memory"`
	Disk           int64       `json:"disk"`
}

func (c *Client) UpdateServer(uuid string, sreq UpdateServerReq) (err error) {
	req, err := c.buildReq("PATCH", "/api/application/servers/"+uuid+"/settings/build", sreq)
	if err != nil {
		return
	}
	var resp struct {
		Server
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

//下面的函数需要修改convoy源码才能用

// GetServerState 获取服务器状态
func (c *Client) GetServerState(uuid string) (state ServerState, err error) {
	req, err := c.buildReq("GET", "/api/application/servers/"+uuid+"/state", nil)
	if err != nil {
		return
	}
	var resp struct {
		Data ServerState `json:"data"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	state = resp.Data
	return
}

// UpdateServerState 更新服务器状态
func (c *Client) UpdateServerState(uuid string, state string) (err error) {
	statereq := StateReq{
		State: state,
	}
	req, err := c.buildReq("PATCH", "/api/application/servers/"+uuid+"/state", statereq)
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

// GetServerAvailableOS 获取服务器可用的操作系统
func (c *Client) GetServerAvailableOS(uuid string) (os []TemplateGroup, err error) {
	req, err := c.buildReq("GET", "/api/application/servers/"+uuid+"/settings/template-groups", nil)
	if err != nil {
		return
	}
	var resp struct {
		Data []TemplateGroup `json:"data"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	os = resp.Data
	return
}

// ReinstallServerOS 重装系统
func (c *Client) ReinstallServerOS(uuid string, reiq ReinstallReq) (err error) {
	req, _ := c.buildReq("POST", "/api/application/servers/"+uuid+"/settings/reinstall", reiq)
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

// GetServerVNC 获取vnc信息
func (c *Client) GetServerVNC(uuid string) (vnc VNC, err error) {
	req, err := c.buildReq("GET", "/api/application/servers/"+uuid+"/terminal", nil)
	if err != nil {
		return
	}
	var resp struct {
		Data VNC `json:"data"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	vnc = resp.Data
	return
}
