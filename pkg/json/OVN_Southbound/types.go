package OVN_Southbound

import "github.com/ibm/ovsdb-etcd/pkg/json"

type Address_Set struct {
	Addresses []string  `json:"addresses,omitempty"`
	Name      string    `json:"name,omitempty"`
	Version   json.Uuid `json:"_version,omitempty"`
	Uuid      json.Uuid `json:"uuid,omitempty"`
}

type Chassis struct {
	Encaps                []json.Uuid       `json:"encaps,omitempty"`
	External_ids          map[string]string `json:"external_ids,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Nb_cfg                int64             `json:"nb_cfg,omitempty"`
	Other_config          map[string]string `json:"other_config,omitempty"`
	Transport_zones       []string          `json:"transport_zones,omitempty"`
	Vtep_logical_switches []string          `json:"vtep_logical_switches,omitempty"`
	Version               json.Uuid         `json:"_version,omitempty"`
	Uuid                  json.Uuid         `json:"uuid,omitempty"`
}

type Chassis_Private struct {
	Chassis          json.Uuid         `json:"chassis,omitempty"`
	External_ids     map[string]string `json:"external_ids,omitempty"`
	Name             string            `json:"name,omitempty"`
	Nb_cfg           int64             `json:"nb_cfg,omitempty"`
	Nb_cfg_timestamp int64             `json:"nb_cfg_timestamp,omitempty"`
	Version          json.Uuid         `json:"_version,omitempty"`
	Uuid             json.Uuid         `json:"uuid,omitempty"`
}

type Connection struct {
	External_ids     map[string]string `json:"external_ids,omitempty"`
	Inactivity_probe int64             `json:"inactivity_probe,omitempty"`
	Is_connected     bool              `json:"is_connected,omitempty"`
	Max_backoff      int64             `json:"max_backoff,omitempty"`
	Other_config     map[string]string `json:"other_config,omitempty"`
	Read_only        bool              `json:"read_only,omitempty"`
	Role             string            `json:"role,omitempty"`
	Status           map[string]string `json:"status,omitempty"`
	Target           string            `json:"target,omitempty"`
	Version          json.Uuid         `json:"_version,omitempty"`
	Uuid             json.Uuid         `json:"uuid,omitempty"`
}

type Controller_Event struct {
	Chassis    json.Uuid         `json:"chassis,omitempty"`
	Event_info map[string]string `json:"event_info,omitempty"`
	Event_type string            `json:"event_type,omitempty"`
	Seq_num    int64             `json:"seq_num,omitempty"`
	Version    json.Uuid         `json:"_version,omitempty"`
	Uuid       json.Uuid         `json:"uuid,omitempty"`
}

type DHCP_Options struct {
	Code    int64     `json:"code,omitempty"`
	Name    string    `json:"name,omitempty"`
	Type    string    `json:"type,omitempty"`
	Version json.Uuid `json:"_version,omitempty"`
	Uuid    json.Uuid `json:"uuid,omitempty"`
}

type DHCPv6_Options struct {
	Code    int64     `json:"code,omitempty"`
	Name    string    `json:"name,omitempty"`
	Type    string    `json:"type,omitempty"`
	Version json.Uuid `json:"_version,omitempty"`
	Uuid    json.Uuid `json:"uuid,omitempty"`
}

type DNS struct {
	Datapaths    []json.Uuid       `json:"datapaths,omitempty"`
	External_ids map[string]string `json:"external_ids,omitempty"`
	Records      map[string]string `json:"records,omitempty"`
	Version      json.Uuid         `json:"_version,omitempty"`
	Uuid         json.Uuid         `json:"uuid,omitempty"`
}

type Datapath_Binding struct {
	External_ids   map[string]string `json:"external_ids,omitempty"`
	Load_balancers []json.Uuid       `json:"load_balancers,omitempty"`
	Tunnel_key     int64             `json:"tunnel_key,omitempty"`
	Version        json.Uuid         `json:"_version,omitempty"`
	Uuid           json.Uuid         `json:"uuid,omitempty"`
}

type Encap struct {
	Chassis_name string            `json:"chassis_name,omitempty"`
	Ip           string            `json:"ip,omitempty"`
	Options      map[string]string `json:"options,omitempty"`
	Type         string            `json:"type,omitempty"`
	Version      json.Uuid         `json:"_version,omitempty"`
	Uuid         json.Uuid         `json:"uuid,omitempty"`
}

type Gateway_Chassis struct {
	Chassis      json.Uuid         `json:"chassis,omitempty"`
	External_ids map[string]string `json:"external_ids,omitempty"`
	Name         string            `json:"name,omitempty"`
	Options      map[string]string `json:"options,omitempty"`
	Priority     int64             `json:"priority,omitempty"`
	Version      json.Uuid         `json:"_version,omitempty"`
	Uuid         json.Uuid         `json:"uuid,omitempty"`
}

type HA_Chassis struct {
	Chassis      json.Uuid         `json:"chassis,omitempty"`
	External_ids map[string]string `json:"external_ids,omitempty"`
	Priority     int64             `json:"priority,omitempty"`
	Version      json.Uuid         `json:"_version,omitempty"`
	Uuid         json.Uuid         `json:"uuid,omitempty"`
}

type HA_Chassis_Group struct {
	External_ids map[string]string `json:"external_ids,omitempty"`
	Ha_chassis   []json.Uuid       `json:"ha_chassis,omitempty"`
	Name         string            `json:"name,omitempty"`
	Ref_chassis  []json.Uuid       `json:"ref_chassis,omitempty"`
	Version      json.Uuid         `json:"_version,omitempty"`
	Uuid         json.Uuid         `json:"uuid,omitempty"`
}

type IGMP_Group struct {
	Address  string      `json:"address,omitempty"`
	Chassis  json.Uuid   `json:"chassis,omitempty"`
	Datapath json.Uuid   `json:"datapath,omitempty"`
	Ports    []json.Uuid `json:"ports,omitempty"`
	Version  json.Uuid   `json:"_version,omitempty"`
	Uuid     json.Uuid   `json:"uuid,omitempty"`
}

type IP_Multicast struct {
	Datapath       json.Uuid `json:"datapath,omitempty"`
	Enabled        bool      `json:"enabled,omitempty"`
	Eth_src        string    `json:"eth_src,omitempty"`
	Idle_timeout   int64     `json:"idle_timeout,omitempty"`
	Ip4_src        string    `json:"ip4_src,omitempty"`
	Ip6_src        string    `json:"ip6_src,omitempty"`
	Querier        bool      `json:"querier,omitempty"`
	Query_interval int64     `json:"query_interval,omitempty"`
	Query_max_resp int64     `json:"query_max_resp,omitempty"`
	Seq_no         int64     `json:"seq_no,omitempty"`
	Table_size     int64     `json:"table_size,omitempty"`
	Version        json.Uuid `json:"_version,omitempty"`
	Uuid           json.Uuid `json:"uuid,omitempty"`
}

type Load_Balancer struct {
	Datapaths    []json.Uuid       `json:"datapaths,omitempty"`
	External_ids map[string]string `json:"external_ids,omitempty"`
	Name         string            `json:"name,omitempty"`
	Protocol     string            `json:"protocol,omitempty"`
	Vips         map[string]string `json:"vips,omitempty"`
	Version      json.Uuid         `json:"_version,omitempty"`
	Uuid         json.Uuid         `json:"uuid,omitempty"`
}

type Logical_DP_Group struct {
	Datapaths []json.Uuid `json:"datapaths,omitempty"`
	Version   json.Uuid   `json:"_version,omitempty"`
	Uuid      json.Uuid   `json:"uuid,omitempty"`
}

type Logical_Flow struct {
	Actions          string            `json:"actions,omitempty"`
	External_ids     map[string]string `json:"external_ids,omitempty"`
	Logical_datapath json.Uuid         `json:"logical_datapath,omitempty"`
	Logical_dp_group json.Uuid         `json:"logical_dp_group,omitempty"`
	Match            string            `json:"match,omitempty"`
	Pipeline         string            `json:"pipeline,omitempty"`
	Priority         int64             `json:"priority,omitempty"`
	Table_id         int64             `json:"table_id,omitempty"`
	Version          json.Uuid         `json:"_version,omitempty"`
	Uuid             json.Uuid         `json:"uuid,omitempty"`
}

type MAC_Binding struct {
	Datapath     json.Uuid `json:"datapath,omitempty"`
	Ip           string    `json:"ip,omitempty"`
	Logical_port string    `json:"logical_port,omitempty"`
	Mac          string    `json:"mac,omitempty"`
	Version      json.Uuid `json:"_version,omitempty"`
	Uuid         json.Uuid `json:"uuid,omitempty"`
}

type Meter struct {
	Bands   []json.Uuid `json:"bands,omitempty"`
	Name    string      `json:"name,omitempty"`
	Unit    string      `json:"unit,omitempty"`
	Version json.Uuid   `json:"_version,omitempty"`
	Uuid    json.Uuid   `json:"uuid,omitempty"`
}

type Meter_Band struct {
	Action     string    `json:"action,omitempty"`
	Burst_size int64     `json:"burst_size,omitempty"`
	Rate       int64     `json:"rate,omitempty"`
	Version    json.Uuid `json:"_version,omitempty"`
	Uuid       json.Uuid `json:"uuid,omitempty"`
}

type Multicast_Group struct {
	Datapath   json.Uuid   `json:"datapath,omitempty"`
	Name       string      `json:"name,omitempty"`
	Ports      []json.Uuid `json:"ports,omitempty"`
	Tunnel_key int64       `json:"tunnel_key,omitempty"`
	Version    json.Uuid   `json:"_version,omitempty"`
	Uuid       json.Uuid   `json:"uuid,omitempty"`
}

type Port_Binding struct {
	Chassis          json.Uuid         `json:"chassis,omitempty"`
	Datapath         json.Uuid         `json:"datapath,omitempty"`
	Encap            json.Uuid         `json:"encap,omitempty"`
	External_ids     map[string]string `json:"external_ids,omitempty"`
	Gateway_chassis  []json.Uuid       `json:"gateway_chassis,omitempty"`
	Ha_chassis_group json.Uuid         `json:"ha_chassis_group,omitempty"`
	Logical_port     string            `json:"logical_port,omitempty"`
	Mac              []string          `json:"mac,omitempty"`
	Nat_addresses    []string          `json:"nat_addresses,omitempty"`
	Options          map[string]string `json:"options,omitempty"`
	Parent_port      string            `json:"parent_port,omitempty"`
	Tag              int64             `json:"tag,omitempty"`
	Tunnel_key       int64             `json:"tunnel_key,omitempty"`
	Type             string            `json:"type,omitempty"`
	Virtual_parent   string            `json:"virtual_parent,omitempty"`
	Version          json.Uuid         `json:"_version,omitempty"`
	Uuid             json.Uuid         `json:"uuid,omitempty"`
}

type Port_Group struct {
	Name    string    `json:"name,omitempty"`
	Ports   []string  `json:"ports,omitempty"`
	Version json.Uuid `json:"_version,omitempty"`
	Uuid    json.Uuid `json:"uuid,omitempty"`
}

type RBAC_Permission struct {
	Authorization []string  `json:"authorization,omitempty"`
	Insert_delete bool      `json:"insert_delete,omitempty"`
	Table         string    `json:"table,omitempty"`
	Update        []string  `json:"update,omitempty"`
	Version       json.Uuid `json:"_version,omitempty"`
	Uuid          json.Uuid `json:"uuid,omitempty"`
}

type RBAC_Role struct {
	Name        string               `json:"name,omitempty"`
	Permissions map[string]json.Uuid `json:"permissions,omitempty"`
	Version     json.Uuid            `json:"_version,omitempty"`
	Uuid        json.Uuid            `json:"uuid,omitempty"`
}

type SB_Global struct {
	Connections  []json.Uuid       `json:"connections,omitempty"`
	External_ids map[string]string `json:"external_ids,omitempty"`
	Ipsec        bool              `json:"ipsec,omitempty"`
	Nb_cfg       int64             `json:"nb_cfg,omitempty"`
	Options      map[string]string `json:"options,omitempty"`
	Ssl          json.Uuid         `json:"ssl,omitempty"`
	Version      json.Uuid         `json:"_version,omitempty"`
	Uuid         json.Uuid         `json:"uuid,omitempty"`
}

type SSL struct {
	Bootstrap_ca_cert bool              `json:"bootstrap_ca_cert,omitempty"`
	Ca_cert           string            `json:"ca_cert,omitempty"`
	Certificate       string            `json:"certificate,omitempty"`
	External_ids      map[string]string `json:"external_ids,omitempty"`
	Private_key       string            `json:"private_key,omitempty"`
	Ssl_ciphers       string            `json:"ssl_ciphers,omitempty"`
	Ssl_protocols     string            `json:"ssl_protocols,omitempty"`
	Version           json.Uuid         `json:"_version,omitempty"`
	Uuid              json.Uuid         `json:"uuid,omitempty"`
}

type Service_Monitor struct {
	External_ids map[string]string `json:"external_ids,omitempty"`
	Ip           string            `json:"ip,omitempty"`
	Logical_port string            `json:"logical_port,omitempty"`
	Options      map[string]string `json:"options,omitempty"`
	Port         int64             `json:"port,omitempty"`
	Protocol     string            `json:"protocol,omitempty"`
	Src_ip       string            `json:"src_ip,omitempty"`
	Src_mac      string            `json:"src_mac,omitempty"`
	Status       string            `json:"status,omitempty"`
	Version      json.Uuid         `json:"_version,omitempty"`
	Uuid         json.Uuid         `json:"uuid,omitempty"`
}