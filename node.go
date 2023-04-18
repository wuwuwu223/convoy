package convoy

import (
	"fmt"
	"strconv"
)

// GetNodes 获取节点列表
func (c *Client) GetNodes() ([]Node, error) {
	req, err := c.buildReq("GET", "/api/application/nodes", nil)
	if err != nil {
		return nil, err
	}
	//把body转换成json
	var resp struct {
		Data []Node `json:"data"`
		Meta Meta   `json:"meta"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.ErrMsg.Message != "" {
		err = fmt.Errorf(resp.ErrMsg.Message)
		return nil, err
	}
	return resp.Data, nil
}

// GetNode 获取节点信息
func (c *Client) GetNode(id int) (Node, error) {
	req, err := c.buildReq("GET", "/api/application/nodes/"+strconv.Itoa(id), nil)
	if err != nil {
		return Node{}, err
	}
	//把body转换成json
	var resp struct {
		Data Node `json:"data"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	if err != nil {
		return Node{}, err
	}
	if resp.ErrMsg.Message != "" {
		err = fmt.Errorf(resp.ErrMsg.Message)
		return Node{}, err
	}
	return resp.Data, nil
}
