package convoy

// GetLocations 获取地域列表
func (c *Client) GetLocations() (locations []Location, err error) {
	req, err := c.buildReq("GET", "/api/application/locations", nil)
	if err != nil {
		return
	}
	var resp struct {
		Data []Location `json:"data"`
		Meta Meta       `json:"meta"`
		ErrMsg
	}
	err = c.doReq(req, &resp)
	locations = resp.Data
	return
}
