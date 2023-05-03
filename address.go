package convoy

import (
	"fmt"
)

// GetAddresses 获取地址列表
func (c *Client) GetAddresses(nodeId string) (addresses []Address, err error) {

	req, err := c.buildReq("GET", "/api/application/nodes/"+nodeId+"/addresses?filter[server_id]=", nil)
	if err != nil {
		return
	}
	var resp struct {
		Data []Address `json:"data"`
		Meta Meta      `json:"meta"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	if err != nil {
		return
	}
	if resp.ErrMsg.Message != "" {
		err = fmt.Errorf(resp.ErrMsg.Message)
		return
	}
	addresses = resp.Data
	return
}

// UpdateAddress 更新地址
func (c *Client) UpdateAddress(nodeId string, ip_id string, address Address) (newaddress Address, err error) {
	req, err := c.buildReq("PUT", "/api/application/nodes/"+nodeId+"/addresses/"+ip_id, address)
	if err != nil {
		return
	}
	var resp struct {
		Data Address `json:"data,omitempty"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	if err != nil {
		return
	}
	if resp.ErrMsg.Message != "" {
		err = fmt.Errorf(resp.ErrMsg.Message)
		return
	}
	newaddress = resp.Data
	return
}

// DeleteAddress 删除地址
func (c *Client) DeleteAddress(nodeId string, ip_id string) (err error) {
	req, err := c.buildReq("DELETE", "/api/application/nodes/"+nodeId+"/addresses/"+ip_id, nil)
	if err != nil {
		return
	}
	var resp struct {
		ErrMsg
	}
	err = c.doReq(req, &resp)
	if err != nil {
		return
	}
	if resp.ErrMsg.Message != "" {
		err = fmt.Errorf(resp.ErrMsg.Message)
		return
	}
	return
}

// CreateAddress 创建地址
func (c *Client) CreateAddress(nodeId string, address Address) (newaddress Address, err error) {
	req, err := c.buildReq("POST", "/api/application/nodes/"+nodeId+"/addresses", address)
	if err != nil {
		return
	}
	var resp struct {
		Data Address `json:"data,omitempty"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	if err != nil {
		return
	}
	if resp.ErrMsg.Message != "" {
		err = fmt.Errorf(resp.ErrMsg.Message)
		return
	}
	newaddress = resp.Data
	return
}
