package convoy

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/url"
	"os"
	"testing"
)

var c *Client

func init() {
	//从.env文件中读取环境变量
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	address := os.Getenv("CONVOY_ADDRESS")
	token := os.Getenv("CONVOY_TOKEN")
	c = New(address, token)
}

func TestClient_GetNodes(t *testing.T) {
	nodes, err := c.GetNodes()
	if err != nil {
		t.Fatal(err)
	}
	for _, node := range nodes {
		t.Log(node)
	}
}

func TestClient_GetNode(t *testing.T) {
	node, err := c.GetNode(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(node)
}

func TestClient_GetLocations(t *testing.T) {
	locations, err := c.GetLocations()
	if err != nil {
		t.Fatal(err)
	}
	for _, location := range locations {
		t.Log(location)
	}
}

func TestClient_GetUsers(t *testing.T) {
	users, err := c.GetUsers()
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range users {
		t.Log(user)
	}
}

func TestClient_GetUser(t *testing.T) {
	user, err := c.GetUser(4)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestClient_CreateUser(t *testing.T) {
	user := User{
		RootAdmin: false,
		Name:      "test2",
		Email:     "test@test2.com",
		Password:  "qCG2xHoA^%@%g",
	}
	newuser, err := c.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(newuser)
}

func TestClient_UpdateUser(t *testing.T) {
	user := User{
		RootAdmin: false,
		Name:      "test23",
		Email:     "ttest@test2.com",
	}
	newuser, err := c.UpdateUser(4, user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(newuser)
}

func TestClient_DeleteUser(t *testing.T) {
	err := c.DeleteUser(2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetAddresses(t *testing.T) {
	addresses, err := c.GetAddresses("1")
	if err != nil {
		t.Fatal(err)
	}
	for _, address := range addresses {
		t.Log(address)
	}
}

func TestClient_UpdateAddress(t *testing.T) {
	address := Address{
		Address:  "10.10.10.35",
		Type:     "ipv4",
		Cidr:     24,
		Gateway:  "10.10.10.1",
		ServerId: 0,
	}
	newaddress, err := c.UpdateAddress("1", "6", address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(newaddress)
}

func TestClient_DeleteAddress(t *testing.T) {
	err := c.DeleteAddress("1", "3")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_CreateAddress(t *testing.T) {
	address := Address{
		Address:  "10.10.10.36",
		Type:     "ipv4",
		Cidr:     24,
		Gateway:  "10.10.10.1",
		ServerId: 0,
	}
	newaddress, err := c.CreateAddress("1", address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(newaddress)
}

func TestClient_GetTemplates(t *testing.T) {
	templates, err := c.GetTemplates("1")
	if err != nil {
		t.Fatal(err)
	}
	for _, template := range templates {
		t.Log(template)
	}
}

func TestClient_GetSsoToken(t *testing.T) {
	token, err := c.GetSsoToken("1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestClient_GetServers(t *testing.T) {
	servers, err := c.GetServers()
	if err != nil {
		t.Fatal(err)
	}
	for _, server := range servers {
		t.Log(server)
	}
}

func TestClient_GetServer(t *testing.T) {
	uuid := "660d52e3-028a-455c-afa3-f6d38f305a45"
	server, err := c.GetServer(uuid)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(server)
}

func TestClient_CreateServer(t *testing.T) {
	sreq := CreateServerReq{
		NodeId:   1,
		UserId:   1,
		Name:     "teststs",
		Hostname: "advinservers.com",
		Limits: struct {
			Cpu        int    `json:"cpu"`
			Memory     int64  `json:"memory"`
			Disk       int64  `json:"disk"`
			Snapshots  int    `json:"snapshots"`
			Backups    *int   `json:"backups"`
			Bandwidth  *int64 `json:"bandwidth"`
			AddressIds []int  `json:"address_ids"`
		}{
			Cpu:        3,
			Memory:     4 * 1024 * 1024 * 1024, //4GB
			Disk:       3 * 1024 * 1024 * 1024, //3GB
			Snapshots:  1,
			Backups:    nil,
			Bandwidth:  new(int64),
			AddressIds: []int{},
		},
		AccountPassword:    "q%#tUyLPAm@2q",
		ShouldCreateServer: true,
		TemplateUuid:       "c8b1de32-4e02-4a87-b40d-e6abbdbc9a4a",
		StartOnCompletion:  false,
	}
	*sreq.Limits.Bandwidth = 1024 * 1024 * 1024 * 1024 //1TB
	newserver, err := c.CreateServer(sreq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(newserver)
}

func TestClient_SuspendServer(t *testing.T) {
	err := c.SuspendServer("5fdaec4d-ba74-4c0e-889c-95a5fda54f6f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UnsuspendServer(t *testing.T) {
	err := c.UnsuspendServer("5fdaec4d-ba74-4c0e-889c-95a5fda54f6f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_DeleteServer(t *testing.T) {
	err := c.DeleteServer("5fdaec4d-ba74-4c0e-889c-95a5fda54f6f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetServerState(t *testing.T) {
	state, err := c.GetServerState("660d52e3-028a-455c-afa3-f6d38f305a45")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(state)
}

func TestClient_UpdateServerState(t *testing.T) {
	//start, stop, restart, kill, reinstall
	err := c.UpdateServerState("660d52e3-028a-455c-afa3-f6d38f305a45", "restart")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetServerAvailableOS(t *testing.T) {
	os, err := c.GetServerAvailableOS("660d52e3-028a-455c-afa3-f6d38f305a45")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(os)
}

func TestClient_ReinstallServerOS(t *testing.T) {
	req := ReinstallReq{
		TemplateUuid:      "c8b1de32-4e02-4a87-b40d-e6abbdbc9a4a",
		AccountPassword:   "q%#tUyLPAm@2q",
		StartOnCompletion: true,
	}
	err := c.ReinstallServerOS("660d52e3-028a-455c-afa3-f6d38f305a45", req)
	if err != nil {
		t.Fatal(err)
	}
	TestClient_GetServer(t)
}

// 测试获取VNC地址
func TestClient_GetServerVNC(t *testing.T) {
	vnc, err := c.GetServerVNC("660d52e3-028a-455c-afa3-f6d38f305a45")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("https://%s:%d/novnc/novnc.html?console=qemu&virtualization=qemu&node=%s&vmid=%d&token=%s", vnc.Fqdn, vnc.Port, vnc.Node, vnc.Vmid, url.QueryEscape(vnc.Token)))
}
