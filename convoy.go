package convoy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Token   string
	Address string
	Client  *http.Client
}

func New(address, token string) (client *Client) {
	client = &Client{
		Token:   token,
		Address: address,
		Client:  http.DefaultClient,
	}
	return
}
func (c *Client) buildReq(mtd, uri string, param any) (req *http.Request, err error) {
	//如果有参数，就把参数转换成json
	if param != nil {
		var b []byte
		b, err = json.Marshal(param)
		if err != nil {
			return
		}
		req, err = http.NewRequest(mtd, c.Address+uri, bytes.NewReader(b))
		if err != nil {
			return
		}
	} else {
		req, err = http.NewRequest(mtd, c.Address+uri, nil)
		if err != nil {
			return
		}
	}
	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return
}

func (c *Client) doReq(req *http.Request, v interface{}) (err error) {
	do, err := c.Client.Do(req)
	if err != nil {
		return
	}
	defer do.Body.Close()
	//如果返回的状态码不是200，就返回错误
	if do.StatusCode != 200 && do.StatusCode != 204 && do.StatusCode != 500 {
		err = fmt.Errorf("status code: %d", do.StatusCode)
		return
	}
	body, err := io.ReadAll(do.Body)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return
	}
	return
}
