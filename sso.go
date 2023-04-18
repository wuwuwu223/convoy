package convoy

import "fmt"

// GetSsoToken 获取SSO Token
func (c *Client) GetSsoToken(userId string) (token string, err error) {
	req, err := c.buildReq("POST", "/api/application/users/"+userId+"/generate-sso-token", nil)
	if err != nil {
		return
	}
	var resp struct {
		Data struct {
			Token  string `json:"token"`
			UserId int    `json:"user_id"`
		}
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
	token = resp.Data.Token
	return
}
