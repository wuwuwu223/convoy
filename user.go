package convoy

import (
	"fmt"
	"strconv"
)

// GetUsers 获取用户列表
func (c *Client) GetUsers() (users []User, err error) {
	req, err := c.buildReq("GET", "/api/application/users", nil)
	if err != nil {
		return
	}
	var resp struct {
		Data []User `json:"data"`
		Meta Meta   `json:"meta"`
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
	users = resp.Data
	return
}

// GetUser 获取用户信息
func (c *Client) GetUser(id int) (user User, err error) {
	req, err := c.buildReq("GET", "/api/application/users/"+strconv.Itoa(id), nil)
	if err != nil {
		return
	}
	var resp struct {
		Data User `json:"data"`
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
	user = resp.Data
	return
}

// CreateUser 创建用户
func (c *Client) CreateUser(user User) (newuser User, err error) {
	req, err := c.buildReq("POST", "/api/application/users", user)
	if err != nil {
		return
	}
	resp := struct {
		Data    User   `json:"data,omitempty"`
		Message string `json:"message"`
	}{}
	err = c.doReq(req, &resp)
	if err != nil {
		return
	}
	if resp.Message != "" {
		err = fmt.Errorf(resp.Message)
		return
	}
	newuser = resp.Data
	return
}

// UpdateUser 更新用户信息
func (c *Client) UpdateUser(id int, user User) (newuser User, err error) {
	req, err := c.buildReq("PUT", "/api/application/users/"+strconv.Itoa(id), user)
	if err != nil {
		return
	}
	resp := struct {
		Data    User   `json:"data,omitempty"`
		Message string `json:"message"`
	}{}
	err = c.doReq(req, &resp)
	if err != nil {
		return
	}
	if resp.Message != "" {
		err = fmt.Errorf(resp.Message)
		return
	}
	newuser = resp.Data
	return
}

// DeleteUser 删除用户
func (c *Client) DeleteUser(id int) (err error) {
	req, err := c.buildReq("DELETE", "/api/application/users/"+strconv.Itoa(id), nil)
	if err != nil {
		return
	}
	err = c.doReq(req, nil)
	if err != nil {
		return
	}
	return
}
