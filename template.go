package convoy

import "fmt"

// GetTemplates 获取模板列表
func (c *Client) GetTemplates(nodeId string) (templates []TemplateGroup, err error) {
	req, err := c.buildReq("GET", "/api/application/nodes/"+nodeId+"/template-groups", nil)
	if err != nil {
		return
	}
	var resp struct {
		Data []TemplateGroup `json:"data"`
		Meta Meta            `json:"meta"`
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
	templates = resp.Data
	return
}
