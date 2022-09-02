package pfsenseapi

import (
	"context"
	"encoding/json"
)

const (
	systemStatusEndpoint      = "api/v1/status/system"
	interfaceStatusEndpoint   = "api/v1/status/interface"
	gatewayStatusEndpoint     = "api/v1/status/gateway"
	firewallLogStatusEndpoint = "api/v1/status/log/firewall"
	systemLogStatusEndpoint   = "api/v1/status/log/system"
	dhcpLogStatusEndpoint     = "api/v1/status/log/dhcp"
)

// StatusService provides Status API methods
type StatusService service

type SystemStatus struct {
	SystemPlatform  string    `json:"system_platform"`
	SystemSerial    string    `json:"system_serial"`
	SystemNetgateId string    `json:"system_netgate_id"`
	BiosVendor      string    `json:"bios_vendor"`
	BiosVersion     string    `json:"bios_version"`
	BiosDate        string    `json:"bios_date"`
	CpuModel        string    `json:"cpu_model"`
	KernelPti       bool      `json:"kernel_pti"`
	MdsMitigation   string    `json:"mds_mitigation"`
	TempC           int       `json:"temp_c"`
	TempF           float64   `json:"temp_f"`
	LoadAvg         []float64 `json:"load_avg"`
	MbufUsage       float64   `json:"mbuf_usage"`
	MemUsage        float64   `json:"mem_usage"`
	SwapUsage       int       `json:"swap_usage"`
	DiskUsage       float64   `json:"disk_usage"`
}

type systemStatusResponse struct {
	apiResponse
	Data *SystemStatus `json:"data"`
}

type InterfaceStatus struct {
	Name          string `json:"name"`
	Descr         string `json:"descr"`
	Hwif          string `json:"hwif"`
	Enable        bool   `json:"enable"`
	If            string `json:"if"`
	Status        string `json:"status"`
	Macaddr       string `json:"macaddr"`
	Mtu           int    `json:"mtu"`
	Ipaddr        string `json:"ipaddr"`
	Subnet        string `json:"subnet"`
	Linklocal     string `json:"linklocal"`
	Ipaddrv6      string `json:"ipaddrv6"`
	Subnetv6      int    `json:"subnetv6"`
	Inerrs        int    `json:"inerrs"`
	Outerrs       int    `json:"outerrs"`
	Collisions    int    `json:"collisions"`
	Inbytespass   int64  `json:"inbytespass"`
	Outbytespass  int64  `json:"outbytespass"`
	Inpktspass    int    `json:"inpktspass"`
	Outpktspass   int    `json:"outpktspass"`
	Inbytesblock  int    `json:"inbytesblock"`
	Outbytesblock int    `json:"outbytesblock"`
	Inpktsblock   int    `json:"inpktsblock"`
	Outpktsblock  int    `json:"outpktsblock"`
	Inbytes       int64  `json:"inbytes"`
	Outbytes      int64  `json:"outbytes"`
	Inpkts        int    `json:"inpkts"`
	Outpkts       int    `json:"outpkts"`
	Dhcplink      string `json:"dhcplink"`
	Media         string `json:"media"`
	Gateway       string `json:"gateway"`
	Gatewayv6     string `json:"gatewayv6"`
}

type interfaceStatusResponse struct {
	apiResponse
	Data []*InterfaceStatus `json:"data"`
}

type GatewayStatus struct {
	Monitorip string  `json:"monitorip"`
	Srcip     string  `json:"srcip"`
	Name      string  `json:"name"`
	Delay     float64 `json:"delay"`
	Stddev    float64 `json:"stddev"`
	Loss      int     `json:"loss"`
	Status    string  `json:"status"`
	Substatus string  `json:"substatus"`
}

type gatewayStatusResponse struct {
	apiResponse
	Data []*GatewayStatus `json:"data"`
}

type logStatusResponse struct {
	apiResponse
	Data []string `json:"data"`
}

// GetSystemStatus returns the system status
func (s StatusService) GetSystemStatus(ctx context.Context) (*SystemStatus, error) {
	response, err := s.client.get(ctx, systemStatusEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(systemStatusResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// ListInterfaceStatus returns the interface status
func (s StatusService) ListInterfaceStatus(ctx context.Context) ([]*InterfaceStatus, error) {
	response, err := s.client.get(ctx, interfaceStatusEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceStatusResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// ListGatewayStatus returns the interface status
func (s StatusService) ListGatewayStatus(ctx context.Context) ([]*GatewayStatus, error) {
	response, err := s.client.get(ctx, gatewayStatusEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(gatewayStatusResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// genericLogRequest returns the a generic Log response
func (s StatusService) genericLogRequest(ctx context.Context, endpoint string) ([]string, error) {
	response, err := s.client.get(ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(logStatusResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// DHCPLog returns the DHCP log
func (s StatusService) DHCPLog(ctx context.Context) ([]string, error) {
	return s.genericLogRequest(ctx, dhcpLogStatusEndpoint)
}

// FirewallLog returns the firewall log
func (s StatusService) FirewallLog(ctx context.Context) ([]string, error) {
	return s.genericLogRequest(ctx, firewallLogStatusEndpoint)
}

// SystemLog returns the firewall log
func (s StatusService) SystemLog(ctx context.Context) ([]string, error) {
	return s.genericLogRequest(ctx, systemLogStatusEndpoint)
}
