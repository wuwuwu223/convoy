package convoy

import "time"

type Server struct {
	Id          string      `json:"id"`
	Uuid        string      `json:"uuid"`
	NodeId      int         `json:"node_id"`
	Hostname    string      `json:"hostname"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	Status      interface{} `json:"status"`
	Usages      struct {
		Bandwidth int `json:"bandwidth"`
	} `json:"usages"`
	Limits struct {
		Cpu       int   `json:"cpu"`
		Memory    int64 `json:"memory"`
		Disk      int64 `json:"disk"`
		Snapshots int   `json:"snapshots"`
		Backups   int   `json:"backups"`
		Bandwidth int64 `json:"bandwidth"`
		Addresses struct {
			Ipv4 []Address `json:"ipv4"`
			Ipv6 []Address `json:"ipv6"`
		} `json:"addresses"`
		MacAddress interface{} `json:"mac_address"`
	} `json:"limits"`
	UserId     int `json:"user_id"`
	Vmid       int `json:"vmid"`
	InternalId int `json:"internal_id"`
}
type User struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	RootAdmin       bool      `json:"root_admin"`
	ServersCount    int       `json:"servers_count"`
}
type Node struct {
	Id                 int    `json:"id"`
	LocationId         int    `json:"location_id"`
	Name               string `json:"name"`
	Cluster            string `json:"cluster"`
	Fqdn               string `json:"fqdn"`
	Port               int    `json:"port"`
	Memory             int64  `json:"memory"`
	MemoryOverallocate int    `json:"memory_overallocate"`
	MemoryAllocated    int64  `json:"memory_allocated"`
	Disk               int64  `json:"disk"`
	DiskOverallocate   int    `json:"disk_overallocate"`
	DiskAllocated      int64  `json:"disk_allocated"`
	VmStorage          string `json:"vm_storage"`
	BackupStorage      string `json:"backup_storage"`
	IsoStorage         string `json:"iso_storage"`
	Network            string `json:"network"`
	ServersCount       int    `json:"servers_count"`
}

type Location struct {
	Id           int    `json:"id"`
	ShortCode    string `json:"short_code"`
	Description  string `json:"description"`
	NodesCount   int    `json:"nodes_count"`
	ServersCount int    `json:"servers_count"`
}

type TemplateGroup struct {
	Id          int    `json:"id"`
	NodeId      int    `json:"node_id"`
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Hidden      int    `json:"hidden"`
	OrderColumn int    `json:"order_column"`
	Templates   struct {
		Data []Template `json:"data"`
	} `json:"templates"`
}

type Template struct {
	Id              int    `json:"id"`
	TemplateGroupId int    `json:"template_group_id"`
	Uuid            string `json:"uuid"`
	Name            string `json:"name"`
	Vmid            int    `json:"vmid"`
	Hidden          int    `json:"hidden"`
	OrderColumn     int    `json:"order_column"`
}

type Address struct {
	Id         int         `json:"id,omitempty"`
	ServerId   int         `json:"server_id,omitempty"`
	Type       string      `json:"type,omitempty"`
	Address    string      `json:"address,omitempty"`
	Cidr       int         `json:"cidr,omitempty"`
	Gateway    string      `json:"gateway,omitempty"`
	MacAddress interface{} `json:"mac_address,omitempty"`
}

type Meta struct {
	Pagination struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
		Links       struct {
		} `json:"links"`
	} `json:"pagination"`
}

type ErrMsg struct {
	Message string `json:"message"`
	Errors  any    `json:"errors"`
}

type CreateServerReq struct {
	NodeId   int    `json:"node_id"`
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Vmid     *int   `json:"vmid"`
	Limits   struct {
		Cpu        int    `json:"cpu"`
		Memory     int64  `json:"memory"`
		Disk       int64  `json:"disk"`
		Snapshots  int    `json:"snapshots"`
		Backups    *int   `json:"backups"`
		Bandwidth  *int64 `json:"bandwidth"`
		AddressIds []int  `json:"address_ids"`
	} `json:"limits"`
	AccountPassword    string `json:"account_password"`
	ShouldCreateServer bool   `json:"should_create_server"`
	TemplateUuid       string `json:"template_uuid"`
	StartOnCompletion  bool   `json:"start_on_completion"`
}

type ServerState struct {
	State       string  `json:"state"`
	CpuUsed     float64 `json:"cpu_used"`
	MemoryTotal int64   `json:"memory_total"`
	MemoryUsed  int     `json:"memory_used"`
	Uptime      int     `json:"uptime"`
}

type StateReq struct {
	State string `json:"state"`
}

type ReinstallReq struct {
	TemplateUuid      string `json:"template_uuid"`
	AccountPassword   string `json:"account_password"`
	StartOnCompletion bool   `json:"start_on_completion"`
}

type VNC struct {
	Token string `json:"token"`
	Node  string `json:"node"`
	Vmid  int    `json:"vmid"`
	Fqdn  string `json:"fqdn"`
	Port  int    `json:"port"`
}
