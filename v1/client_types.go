// Package proxmox was automatically generated
package proxmox

// ClusterReplicationCreateRequest replication - Create a new replication job
// Create a new replication job
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/replication
//
type ClusterReplicationCreateRequest struct {
	Comment   string `json:"comment,omitempty" url:"comment,omitempty,optional"`         // Description.
	Id        string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`         // Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	Rate      int    `json:"rate,omitempty" url:"rate,omitempty,optional"`               // Rate limit in mbps (megabytes per second) as floating point number.
	Schedule  string `json:"schedule,omitempty" url:"schedule,omitempty,optional"`       // Storage replication schedule. The format is a subset of `systemd` calendar events.
	Source    string `json:"source,omitempty" url:"source,omitempty,optional"`           // For internal use, to detect if the guest was stolen.
	Target    string `json:"target,omitempty" url:"target,omitempty" validate:"nonzero"` // Target node.
	Type      string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     // Section type.
	Disable   bool   `json:"disable,omitempty" url:"disable,omitempty,optional"`         // Flag to disable/deactivate the entry.
	RemoveJob string `json:"remove_job,omitempty" url:"remove_job,omitempty,optional"`   // Mark the replication job for removal. The job will remove all local replication snapshots. When set to 'full', it also tries to remove replicated volumes on the target. The job then removes itself from the configuration file.
}

// ClusterReplicationDeleteRequest {id} - Mark replication job for removal.
// Mark replication job for removal.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/replication/{id}
//
type ClusterReplicationDeleteRequest struct {
	Force bool   `json:"force,omitempty" url:"force,omitempty,optional"`     // Will remove the jobconfig entry, but will not cleanup.
	Id    string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"` // Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	Keep  bool   `json:"keep,omitempty" url:"keep,omitempty,optional"`       // Keep replicated data at target (do not remove).
}

// ClusterReplicationReadRequest {id} - Read replication job configuration.
// Read replication job configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/replication/{id}
//
type ClusterReplicationReadRequest struct {
	Id string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"` // Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
}

// ClusterReplicationUpdateRequest {id} - Update replication job configuration.
// Update replication job configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/replication/{id}
//
type ClusterReplicationUpdateRequest struct {
	Delete    string `json:"delete,omitempty" url:"delete,omitempty,optional"`         // A list of settings you want to delete.
	Disable   bool   `json:"disable,omitempty" url:"disable,omitempty,optional"`       // Flag to disable/deactivate the entry.
	Comment   string `json:"comment,omitempty" url:"comment,omitempty,optional"`       // Description.
	Digest    string `json:"digest,omitempty" url:"digest,omitempty,optional"`         // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Id        string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`       // Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	Rate      int    `json:"rate,omitempty" url:"rate,omitempty,optional"`             // Rate limit in mbps (megabytes per second) as floating point number.
	RemoveJob string `json:"remove_job,omitempty" url:"remove_job,omitempty,optional"` // Mark the replication job for removal. The job will remove all local replication snapshots. When set to 'full', it also tries to remove replicated volumes on the target. The job then removes itself from the configuration file.
	Schedule  string `json:"schedule,omitempty" url:"schedule,omitempty,optional"`     // Storage replication schedule. The format is a subset of `systemd` calendar events.
	Source    string `json:"source,omitempty" url:"source,omitempty,optional"`         // For internal use, to detect if the guest was stolen.
}

// ClusterMetricsServerIndexResponse server
// List configured metric servers.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/metrics/server
//
type ClusterMetricsServerIndexResponse struct {
	Disable bool   `json:"disable,omitempty" url:"disable,omitempty" validate:"nonzero"` // Flag to disable the plugin.
	Id      string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`           // The ID of the entry.
	Port    int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`       // Server network port
	Server  string `json:"server,omitempty" url:"server,omitempty" validate:"nonzero"`   // Server dns name or IP address
	Type    string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`       // Plugin type.
}

// ClusterMetricsServerUpdateRequest {id} - Update metric server configuration.
// Update metric server configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/metrics/server/{id}
//
type ClusterMetricsServerUpdateRequest struct {
	ApiPathPrefix     string `json:"api-path-prefix,omitempty" url:"api-path-prefix,omitempty,optional"`       // An API path prefix inserted between '<host>:<port>/' and '/api2/'. Can be useful if the InfluxDB service runs behind a reverse proxy.
	Delete            string `json:"delete,omitempty" url:"delete,omitempty,optional"`                         // A list of settings you want to delete.
	Id                string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`                       // The ID of the entry.
	MaxBodySize       int    `json:"max-body-size,omitempty" url:"max-body-size,omitempty,optional"`           // InfluxDB max-body-size in bytes. Requests are batched up to this size.
	Bucket            string `json:"bucket,omitempty" url:"bucket,omitempty,optional"`                         // The InfluxDB bucket/db. Only necessary when using the http v2 api.
	Server            string `json:"server,omitempty" url:"server,omitempty" validate:"nonzero"`               // server dns name or IP address
	Token             string `json:"token,omitempty" url:"token,omitempty,optional"`                           // The InfluxDB access token. Only necessary when using the http v2 api. If the v2 compatibility api is used, use 'user:password' instead.
	VerifyCertificate bool   `json:"verify-certificate,omitempty" url:"verify-certificate,omitempty,optional"` // Set to 0 to disable certificate verification for https endpoints.
	Digest            string `json:"digest,omitempty" url:"digest,omitempty,optional"`                         // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Organization      string `json:"organization,omitempty" url:"organization,omitempty,optional"`             // The InfluxDB organization. Only necessary when using the http v2 api. Has no meaning when using v2 compatibility api.
	Port              int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`                   // server network port
	Proto             string `json:"proto,omitempty" url:"proto,omitempty,optional"`                           // Protocol to send graphite data. TCP or UDP (default)
	Timeout           int    `json:"timeout,omitempty" url:"timeout,omitempty,optional"`                       // graphite TCP socket timeout (default=1)
	Disable           bool   `json:"disable,omitempty" url:"disable,omitempty,optional"`                       // Flag to disable the plugin.
	Influxdbproto     string `json:"influxdbproto,omitempty" url:"influxdbproto,omitempty,optional"`           //
	Mtu               int    `json:"mtu,omitempty" url:"mtu,omitempty,optional"`                               // MTU for metrics transmission over UDP
	Path              string `json:"path,omitempty" url:"path,omitempty,optional"`                             // root graphite path (ex: proxmox.mycluster.mykey)
}

// ClusterMetricsServerDeleteRequest {id} - Remove Metric server.
// Remove Metric server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/metrics/server/{id}
//
type ClusterMetricsServerDeleteRequest struct {
	Id string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"` //
}

// ClusterMetricsServerReadRequest {id} - Read metric server configuration.
// Read metric server configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/metrics/server/{id}
//
type ClusterMetricsServerReadRequest struct {
	Id string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"` //
}

// ClusterMetricsServerCreateRequest {id} - Create a new external metric server config
// Create a new external metric server config
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/metrics/server/{id}
//
type ClusterMetricsServerCreateRequest struct {
	ApiPathPrefix     string `json:"api-path-prefix,omitempty" url:"api-path-prefix,omitempty,optional"`       // An API path prefix inserted between '<host>:<port>/' and '/api2/'. Can be useful if the InfluxDB service runs behind a reverse proxy.
	Disable           bool   `json:"disable,omitempty" url:"disable,omitempty,optional"`                       // Flag to disable the plugin.
	Influxdbproto     string `json:"influxdbproto,omitempty" url:"influxdbproto,omitempty,optional"`           //
	Server            string `json:"server,omitempty" url:"server,omitempty" validate:"nonzero"`               // server dns name or IP address
	Timeout           int    `json:"timeout,omitempty" url:"timeout,omitempty,optional"`                       // graphite TCP socket timeout (default=1)
	VerifyCertificate bool   `json:"verify-certificate,omitempty" url:"verify-certificate,omitempty,optional"` // Set to 0 to disable certificate verification for https endpoints.
	Mtu               int    `json:"mtu,omitempty" url:"mtu,omitempty,optional"`                               // MTU for metrics transmission over UDP
	Organization      string `json:"organization,omitempty" url:"organization,omitempty,optional"`             // The InfluxDB organization. Only necessary when using the http v2 api. Has no meaning when using v2 compatibility api.
	Path              string `json:"path,omitempty" url:"path,omitempty,optional"`                             // root graphite path (ex: proxmox.mycluster.mykey)
	Port              int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`                   // server network port
	Proto             string `json:"proto,omitempty" url:"proto,omitempty,optional"`                           // Protocol to send graphite data. TCP or UDP (default)
	Type              string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`                   // Plugin type.
	Bucket            string `json:"bucket,omitempty" url:"bucket,omitempty,optional"`                         // The InfluxDB bucket/db. Only necessary when using the http v2 api.
	Id                string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`                       // The ID of the entry.
	MaxBodySize       int    `json:"max-body-size,omitempty" url:"max-body-size,omitempty,optional"`           // InfluxDB max-body-size in bytes. Requests are batched up to this size.
	Token             string `json:"token,omitempty" url:"token,omitempty,optional"`                           // The InfluxDB access token. Only necessary when using the http v2 api. If the v2 compatibility api is used, use 'user:password' instead.
}

// ClusterConfigCreateRequest config - Generate new cluster configuration. If no links given, default to local IP address as link0.
// Generate new cluster configuration. If no links given, default to local IP address as link0.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/config
//
type ClusterConfigCreateRequest struct {
	Clustername string   `json:"clustername,omitempty" url:"clustername,omitempty" validate:"nonzero"` // The name of the cluster.
	Link        []string `json:"link[n],omitempty" url:"link[n],omitempty,optional"`                   // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	Nodeid      int      `json:"nodeid,omitempty" url:"nodeid,omitempty,optional"`                     // Node id for this node.
	Votes       int      `json:"votes,omitempty" url:"votes,omitempty,optional"`                       // Number of votes for this node.
}

// ClusterConfigNodesResponse nodes
// Corosync node list.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/config/nodes
//
type ClusterConfigNodesResponse struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` //
}

// ClusterConfigNodesDelnodeRequest {node} - Removes a node from the cluster configuration.
// Removes a node from the cluster configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/config/nodes/{node}
//
type ClusterConfigNodesDelnodeRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// ClusterConfigNodesAddnodeRequest {node} - Adds a node to the cluster configuration. This call is for internal use.
// Adds a node to the cluster configuration. This call is for internal use.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/config/nodes/{node}
//
type ClusterConfigNodesAddnodeRequest struct {
	Link       []string `json:"link[n],omitempty" url:"link[n],omitempty,optional"`         // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	NewNodeIp  string   `json:"new_node_ip,omitempty" url:"new_node_ip,omitempty,optional"` // IP Address of node to add. Used as fallback if no links are given.
	Node       string   `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Nodeid     int      `json:"nodeid,omitempty" url:"nodeid,omitempty,optional"`           // Node id for this node.
	Votes      int      `json:"votes,omitempty" url:"votes,omitempty,optional"`             // Number of votes for this node
	Apiversion int      `json:"apiversion,omitempty" url:"apiversion,omitempty,optional"`   // The JOIN_API_VERSION of the new node.
	Force      bool     `json:"force,omitempty" url:"force,omitempty,optional"`             // Do not throw error if node already exists.
}

// ClusterConfigNodesAddnodeResponse {node}
// Adds a node to the cluster configuration. This call is for internal use.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/config/nodes/{node}
//
type ClusterConfigNodesAddnodeResponse struct {
	CorosyncAuthkey string `json:"corosync_authkey,omitempty" url:"corosync_authkey,omitempty" validate:"nonzero"` //
	CorosyncConf    string `json:"corosync_conf,omitempty" url:"corosync_conf,omitempty" validate:"nonzero"`       //
	Warnings        []struct {
	} `json:"warnings,omitempty" url:"warnings,omitempty" validate:"nonzero"` //
}

// ClusterConfigJoinInfoRequest join - Get information needed to join this cluster over the connected node.
// Get information needed to join this cluster over the connected node.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/config/join
//
type ClusterConfigJoinInfoRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty,optional"` // The node for which the joinee gets the nodeinfo.
}

// ClusterConfigJoinInfoResponse join
// Get information needed to join this cluster over the connected node.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/config/join
//
type ClusterConfigJoinInfoResponse struct {
	ConfigDigest string `json:"config_digest,omitempty" url:"config_digest,omitempty" validate:"nonzero"` //
	Nodelist     []struct {
	} `json:"nodelist,omitempty" url:"nodelist,omitempty" validate:"nonzero"` //
	PreferredNode string   `json:"preferred_node,omitempty" url:"preferred_node,omitempty" validate:"nonzero"` // The cluster node name.
	Totem         struct { //
	} `json:"totem,omitempty" url:"totem,omitempty" validate:"nonzero"` //
}

// ClusterConfigJoinRequest join - Joins this node into an existing cluster. If no links are given, default to IP resolved by node's hostname on single link (fallback fails for clusters with multiple links).
// Joins this node into an existing cluster. If no links are given, default to IP resolved by node's hostname on single link (fallback fails for clusters with multiple links).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/config/join
//
type ClusterConfigJoinRequest struct {
	Link        []string `json:"link[n],omitempty" url:"link[n],omitempty,optional"`                   // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	Nodeid      int      `json:"nodeid,omitempty" url:"nodeid,omitempty,optional"`                     // Node id for this node.
	Password    string   `json:"password,omitempty" url:"password,omitempty" validate:"nonzero"`       // Superuser (root) password of peer node.
	Votes       int      `json:"votes,omitempty" url:"votes,omitempty,optional"`                       // Number of votes for this node
	Fingerprint string   `json:"fingerprint,omitempty" url:"fingerprint,omitempty" validate:"nonzero"` // Certificate SHA 256 fingerprint.
	Force       bool     `json:"force,omitempty" url:"force,omitempty,optional"`                       // Do not throw error if node already exists.
	Hostname    string   `json:"hostname,omitempty" url:"hostname,omitempty" validate:"nonzero"`       // Hostname (or IP) of an existing cluster member.
}

// ClusterFirewallGroupsCreateSecurityGroupRequest groups - Create new security group.
// Create new security group.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/groups
//
type ClusterFirewallGroupsCreateSecurityGroupRequest struct {
	Rename  string `json:"rename,omitempty" url:"rename,omitempty,optional"`         // Rename/update an existing security group. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing group.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`       //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`         // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Group   string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"` // Security Group name.
}

// ClusterFirewallGroupsListSecurityResponse groups
// List security groups.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/groups
//
type ClusterFirewallGroupsListSecurityResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Group   string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"`   // Security Group name.
}

// ClusterFirewallGroupsDeleteSecurityGroupRequest {group} - Delete security group.
// Delete security group.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/groups/{group}
//
type ClusterFirewallGroupsDeleteSecurityGroupRequest struct {
	Group string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"` // Security Group name.
}

// ClusterFirewallGroupsGetRulesRequest {group} - List rules.
// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/groups/{group}
//
type ClusterFirewallGroupsGetRulesRequest struct {
	Group string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"` // Security Group name.
}

// ClusterFirewallGroupsGetRulesResponse {group}
// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/groups/{group}
//
type ClusterFirewallGroupsGetRulesResponse struct {
	Pos int `json:"pos,omitempty" url:"pos,omitempty" validate:"nonzero"` //
}

// ClusterFirewallGroupsCreateRuleRequest {group} - Create new rule.
// Create new rule.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/groups/{group}
//
type ClusterFirewallGroupsCreateRuleRequest struct {
	Comment  string `json:"comment,omitempty" url:"comment,omitempty,optional"`         // Descriptive comment.
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`           // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Iface    string `json:"iface,omitempty" url:"iface,omitempty,optional"`             // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Sport    string `json:"sport,omitempty" url:"sport,omitempty,optional"`             // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Pos      int    `json:"pos,omitempty" url:"pos,omitempty,optional"`                 // Update rule at position <pos>.
	Action   string `json:"action,omitempty" url:"action,omitempty" validate:"nonzero"` // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Dest     string `json:"dest,omitempty" url:"dest,omitempty,optional"`               // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dport    string `json:"dport,omitempty" url:"dport,omitempty,optional"`             // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Group    string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"`   // Security Group name.
	Macro    string `json:"macro,omitempty" url:"macro,omitempty,optional"`             // Use predefined standard macro.
	Enable   int    `json:"enable,omitempty" url:"enable,omitempty,optional"`           // Flag to enable/disable a rule.
	Proto    string `json:"proto,omitempty" url:"proto,omitempty,optional"`             // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Source   string `json:"source,omitempty" url:"source,omitempty,optional"`           // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Type     string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     // Rule type.
	IcmpType string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`     // Specify icmp-type. Only valid if proto equals 'icmp'.
	Log      string `json:"log,omitempty" url:"log,omitempty,optional"`                 // Log level for firewall rule.
}

// ClusterFirewallGroupsDeleteRuleRequest {pos} - Delete rule.
// Delete rule.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/groups/{group}/{pos}
//
type ClusterFirewallGroupsDeleteRuleRequest struct {
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`         // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Group  string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"` // Security Group name.
	Pos    int    `json:"pos,omitempty" url:"pos,omitempty,optional"`               // Update rule at position <pos>.
}

// ClusterFirewallGroupsGetRuleRequest {pos} - Get single rule data.
// Get single rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/groups/{group}/{pos}
//
type ClusterFirewallGroupsGetRuleRequest struct {
	Pos   int    `json:"pos,omitempty" url:"pos,omitempty,optional"`               // Update rule at position <pos>.
	Group string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"` // Security Group name.
}

// ClusterFirewallGroupsGetRuleResponse {pos}
// Get single rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/groups/{group}/{pos}
//
type ClusterFirewallGroupsGetRuleResponse struct {
	Ipversion int    `json:"ipversion,omitempty" url:"ipversion,omitempty,optional"`     //
	Dport     string `json:"dport,omitempty" url:"dport,omitempty,optional"`             //
	Enable    int    `json:"enable,omitempty" url:"enable,omitempty,optional"`           //
	Log       string `json:"log,omitempty" url:"log,omitempty,optional"`                 // Log level for firewall rule
	Macro     string `json:"macro,omitempty" url:"macro,omitempty,optional"`             //
	Pos       int    `json:"pos,omitempty" url:"pos,omitempty" validate:"nonzero"`       //
	Sport     string `json:"sport,omitempty" url:"sport,omitempty,optional"`             //
	Dest      string `json:"dest,omitempty" url:"dest,omitempty,optional"`               //
	Source    string `json:"source,omitempty" url:"source,omitempty,optional"`           //
	Type      string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     //
	Action    string `json:"action,omitempty" url:"action,omitempty" validate:"nonzero"` //
	Comment   string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	IcmpType  string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`     //
	Iface     string `json:"iface,omitempty" url:"iface,omitempty,optional"`             //
	Proto     string `json:"proto,omitempty" url:"proto,omitempty,optional"`             //
}

// ClusterFirewallGroupsUpdateRuleRequest {pos} - Modify rule data.
// Modify rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/groups/{group}/{pos}
//
type ClusterFirewallGroupsUpdateRuleRequest struct {
	Comment  string `json:"comment,omitempty" url:"comment,omitempty,optional"`       // Descriptive comment.
	Group    string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"` // Security Group name.
	Macro    string `json:"macro,omitempty" url:"macro,omitempty,optional"`           // Use predefined standard macro.
	Proto    string `json:"proto,omitempty" url:"proto,omitempty,optional"`           // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Source   string `json:"source,omitempty" url:"source,omitempty,optional"`         // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Log      string `json:"log,omitempty" url:"log,omitempty,optional"`               // Log level for firewall rule.
	Moveto   int    `json:"moveto,omitempty" url:"moveto,omitempty,optional"`         // Move rule to new position <moveto>. Other arguments are ignored.
	Pos      int    `json:"pos,omitempty" url:"pos,omitempty,optional"`               // Update rule at position <pos>.
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`         // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Dport    string `json:"dport,omitempty" url:"dport,omitempty,optional"`           // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	IcmpType string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`   // Specify icmp-type. Only valid if proto equals 'icmp'.
	Sport    string `json:"sport,omitempty" url:"sport,omitempty,optional"`           // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Type     string `json:"type,omitempty" url:"type,omitempty,optional"`             // Rule type.
	Action   string `json:"action,omitempty" url:"action,omitempty,optional"`         // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Delete   string `json:"delete,omitempty" url:"delete,omitempty,optional"`         // A list of settings you want to delete.
	Dest     string `json:"dest,omitempty" url:"dest,omitempty,optional"`             // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Enable   int    `json:"enable,omitempty" url:"enable,omitempty,optional"`         // Flag to enable/disable a rule.
	Iface    string `json:"iface,omitempty" url:"iface,omitempty,optional"`           // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
}

// ClusterFirewallRulesGetResponse rules
// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/rules
//
type ClusterFirewallRulesGetResponse struct {
	Pos int `json:"pos,omitempty" url:"pos,omitempty" validate:"nonzero"` //
}

// ClusterFirewallRulesCreateRuleRequest rules - Create new rule.
// Create new rule.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/rules
//
type ClusterFirewallRulesCreateRuleRequest struct {
	Iface    string `json:"iface,omitempty" url:"iface,omitempty,optional"`             // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Macro    string `json:"macro,omitempty" url:"macro,omitempty,optional"`             // Use predefined standard macro.
	Dest     string `json:"dest,omitempty" url:"dest,omitempty,optional"`               // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Enable   int    `json:"enable,omitempty" url:"enable,omitempty,optional"`           // Flag to enable/disable a rule.
	Log      string `json:"log,omitempty" url:"log,omitempty,optional"`                 // Log level for firewall rule.
	Source   string `json:"source,omitempty" url:"source,omitempty,optional"`           // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Type     string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     // Rule type.
	Action   string `json:"action,omitempty" url:"action,omitempty" validate:"nonzero"` // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	IcmpType string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`     // Specify icmp-type. Only valid if proto equals 'icmp'.
	Pos      int    `json:"pos,omitempty" url:"pos,omitempty,optional"`                 // Update rule at position <pos>.
	Sport    string `json:"sport,omitempty" url:"sport,omitempty,optional"`             // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Comment  string `json:"comment,omitempty" url:"comment,omitempty,optional"`         // Descriptive comment.
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`           // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Dport    string `json:"dport,omitempty" url:"dport,omitempty,optional"`             // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Proto    string `json:"proto,omitempty" url:"proto,omitempty,optional"`             // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
}

// ClusterFirewallRulesDeleteRuleRequest {pos} - Delete rule.
// Delete rule.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/rules/{pos}
//
type ClusterFirewallRulesDeleteRuleRequest struct {
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Pos    int    `json:"pos,omitempty" url:"pos,omitempty,optional"`       // Update rule at position <pos>.
}

// ClusterFirewallRulesGetRuleRequest {pos} - Get single rule data.
// Get single rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/rules/{pos}
//
type ClusterFirewallRulesGetRuleRequest struct {
	Pos int `json:"pos,omitempty" url:"pos,omitempty,optional"` // Update rule at position <pos>.
}

// ClusterFirewallRulesGetRuleResponse {pos}
// Get single rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/rules/{pos}
//
type ClusterFirewallRulesGetRuleResponse struct {
	Comment   string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Dport     string `json:"dport,omitempty" url:"dport,omitempty,optional"`             //
	IcmpType  string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`     //
	Source    string `json:"source,omitempty" url:"source,omitempty,optional"`           //
	Action    string `json:"action,omitempty" url:"action,omitempty" validate:"nonzero"` //
	Iface     string `json:"iface,omitempty" url:"iface,omitempty,optional"`             //
	Pos       int    `json:"pos,omitempty" url:"pos,omitempty" validate:"nonzero"`       //
	Sport     string `json:"sport,omitempty" url:"sport,omitempty,optional"`             //
	Dest      string `json:"dest,omitempty" url:"dest,omitempty,optional"`               //
	Enable    int    `json:"enable,omitempty" url:"enable,omitempty,optional"`           //
	Macro     string `json:"macro,omitempty" url:"macro,omitempty,optional"`             //
	Ipversion int    `json:"ipversion,omitempty" url:"ipversion,omitempty,optional"`     //
	Log       string `json:"log,omitempty" url:"log,omitempty,optional"`                 // Log level for firewall rule
	Proto     string `json:"proto,omitempty" url:"proto,omitempty,optional"`             //
	Type      string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     //
}

// ClusterFirewallRulesUpdateRuleRequest {pos} - Modify rule data.
// Modify rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/rules/{pos}
//
type ClusterFirewallRulesUpdateRuleRequest struct {
	Dest     string `json:"dest,omitempty" url:"dest,omitempty,optional"`           // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dport    string `json:"dport,omitempty" url:"dport,omitempty,optional"`         // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Type     string `json:"type,omitempty" url:"type,omitempty,optional"`           // Rule type.
	Enable   int    `json:"enable,omitempty" url:"enable,omitempty,optional"`       // Flag to enable/disable a rule.
	Proto    string `json:"proto,omitempty" url:"proto,omitempty,optional"`         // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	IcmpType string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"` // Specify icmp-type. Only valid if proto equals 'icmp'.
	Iface    string `json:"iface,omitempty" url:"iface,omitempty,optional"`         // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Macro    string `json:"macro,omitempty" url:"macro,omitempty,optional"`         // Use predefined standard macro.
	Moveto   int    `json:"moveto,omitempty" url:"moveto,omitempty,optional"`       // Move rule to new position <moveto>. Other arguments are ignored.
	Pos      int    `json:"pos,omitempty" url:"pos,omitempty,optional"`             // Update rule at position <pos>.
	Action   string `json:"action,omitempty" url:"action,omitempty,optional"`       // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Comment  string `json:"comment,omitempty" url:"comment,omitempty,optional"`     // Descriptive comment.
	Delete   string `json:"delete,omitempty" url:"delete,omitempty,optional"`       // A list of settings you want to delete.
	Log      string `json:"log,omitempty" url:"log,omitempty,optional"`             // Log level for firewall rule.
	Source   string `json:"source,omitempty" url:"source,omitempty,optional"`       // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Sport    string `json:"sport,omitempty" url:"sport,omitempty,optional"`         // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
}

// ClusterFirewallIpsetIndexResponse ipset
// List IPSets
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/ipset
//
type ClusterFirewallIpsetIndexResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`     // IP set name.
}

// ClusterFirewallIpsetCreateRequest ipset - Create new IPSet
// Create new IPSet
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/ipset
//
type ClusterFirewallIpsetCreateRequest struct {
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Rename  string `json:"rename,omitempty" url:"rename,omitempty,optional"`       // Rename an existing IPSet. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing IPSet.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}

// ClusterFirewallIpsetDeleteRequest {name} - Delete IPSet
// Delete IPSet
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/ipset/{name}
//
type ClusterFirewallIpsetDeleteRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
}

// ClusterFirewallIpsetGetRequest {name} - List IPSet content
// List IPSet content
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/ipset/{name}
//
type ClusterFirewallIpsetGetRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
}

// ClusterFirewallIpsetGetResponse {name}
// List IPSet content
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/ipset/{name}
//
type ClusterFirewallIpsetGetResponse struct {
	Nomatch bool   `json:"nomatch,omitempty" url:"nomatch,omitempty,optional"`         //
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"`     //
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}

// ClusterFirewallIpsetCreateIpRequest {name} - Add IP or Network to IPSet.
// Add IP or Network to IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/ipset/{name}
//
type ClusterFirewallIpsetCreateIpRequest struct {
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Nomatch bool   `json:"nomatch,omitempty" url:"nomatch,omitempty,optional"`     //
}

// ClusterFirewallIpsetRemoveIpRequest {cidr} - Remove IP or Network from IPSet.
// Remove IP or Network from IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/ipset/{name}/{cidr}
//
type ClusterFirewallIpsetRemoveIpRequest struct {
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Cidr   string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
}

// ClusterFirewallIpsetReadIpRequest {cidr} - Read IP or Network settings from IPSet.
// Read IP or Network settings from IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/ipset/{name}/{cidr}
//
type ClusterFirewallIpsetReadIpRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Cidr string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
}

// ClusterFirewallIpsetUpdateIpRequest {cidr} - Update IP or Network settings
// Update IP or Network settings
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/ipset/{name}/{cidr}
//
type ClusterFirewallIpsetUpdateIpRequest struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Nomatch bool   `json:"nomatch,omitempty" url:"nomatch,omitempty,optional"`     //
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
}

// ClusterFirewallAliasesGetResponse aliases
// List aliases
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/aliases
//
type ClusterFirewallAliasesGetResponse struct {
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`     //
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"`     //
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}

// ClusterFirewallAliasesCreateAliasRequest aliases - Create IP or Network Alias.
// Create IP or Network Alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/aliases
//
type ClusterFirewallAliasesCreateAliasRequest struct {
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
}

// ClusterFirewallAliasesUpdateAliasRequest {name} - Update IP or Network alias.
// Update IP or Network alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/aliases/{name}
//
type ClusterFirewallAliasesUpdateAliasRequest struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
	Rename  string `json:"rename,omitempty" url:"rename,omitempty,optional"`       // Rename an existing alias.
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
}

// ClusterFirewallAliasesRemoveAliasRequest {name} - Remove IP or Network alias.
// Remove IP or Network alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/aliases/{name}
//
type ClusterFirewallAliasesRemoveAliasRequest struct {
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
}

// ClusterFirewallAliasesReadAliasRequest {name} - Read alias.
// Read alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/aliases/{name}
//
type ClusterFirewallAliasesReadAliasRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
}

// ClusterFirewallOptionsGetResponse options
// Get Firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/options
//
type ClusterFirewallOptionsGetResponse struct {
	Enable       int    `json:"enable,omitempty" url:"enable,omitempty,optional"`               // Enable or disable the firewall cluster wide.
	LogRatelimit string `json:"log_ratelimit,omitempty" url:"log_ratelimit,omitempty,optional"` // Log ratelimiting settings
	PolicyIn     string `json:"policy_in,omitempty" url:"policy_in,omitempty,optional"`         // Input policy.
	PolicyOut    string `json:"policy_out,omitempty" url:"policy_out,omitempty,optional"`       // Output policy.
	Ebtables     bool   `json:"ebtables,omitempty" url:"ebtables,omitempty,optional"`           // Enable ebtables rules cluster wide.
}

// ClusterFirewallOptionsSetRequest options - Set Firewall options.
// Set Firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/options
//
type ClusterFirewallOptionsSetRequest struct {
	Ebtables     bool   `json:"ebtables,omitempty" url:"ebtables,omitempty,optional"`           // Enable ebtables rules cluster wide.
	Enable       int    `json:"enable,omitempty" url:"enable,omitempty,optional"`               // Enable or disable the firewall cluster wide.
	LogRatelimit string `json:"log_ratelimit,omitempty" url:"log_ratelimit,omitempty,optional"` // Log ratelimiting settings
	PolicyIn     string `json:"policy_in,omitempty" url:"policy_in,omitempty,optional"`         // Input policy.
	PolicyOut    string `json:"policy_out,omitempty" url:"policy_out,omitempty,optional"`       // Output policy.
	Delete       string `json:"delete,omitempty" url:"delete,omitempty,optional"`               // A list of settings you want to delete.
	Digest       string `json:"digest,omitempty" url:"digest,omitempty,optional"`               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}

// ClusterFirewallMacrosGetResponse macros
// List available macros
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/macros
//
type ClusterFirewallMacrosGetResponse struct {
	Descr string `json:"descr,omitempty" url:"descr,omitempty" validate:"nonzero"` // More verbose description (if available).
	Macro string `json:"macro,omitempty" url:"macro,omitempty" validate:"nonzero"` // Macro name.
}

// ClusterFirewallRefsRequest refs - Lists possible IPSet/Alias reference which are allowed in source/dest properties.
// Lists possible IPSet/Alias reference which are allowed in source/dest properties.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/refs
//
type ClusterFirewallRefsRequest struct {
	Type string `json:"type,omitempty" url:"type,omitempty,optional"` // Only list references of specified type.
}

// ClusterFirewallRefsResponse refs
// Lists possible IPSet/Alias reference which are allowed in source/dest properties.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/firewall/refs
//
type ClusterFirewallRefsResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` //
	Ref     string `json:"ref,omitempty" url:"ref,omitempty" validate:"nonzero"`   //
	Type    string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"` //
}

// ClusterBackupCreateJobRequest backup - Create new vzdump backup job.
// Create new vzdump backup job.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/backup
//
type ClusterBackupCreateJobRequest struct {
	Protected        bool   `json:"protected,omitempty" url:"protected,omitempty,optional"`               // If true, mark backup(s) as protected.
	Stop             bool   `json:"stop,omitempty" url:"stop,omitempty,optional"`                         // Stop running backup jobs on this host.
	Stopwait         int    `json:"stopwait,omitempty" url:"stopwait,omitempty,optional"`                 // Maximal time to wait until a guest system is stopped (minutes).
	Zstd             int    `json:"zstd,omitempty" url:"zstd,omitempty,optional"`                         // Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
	Ionice           int    `json:"ionice,omitempty" url:"ionice,omitempty,optional"`                     // Set CFQ ionice priority.
	Node             string `json:"node,omitempty" url:"node,omitempty,optional"`                         // Only run if executed on this node.
	Exclude          string `json:"exclude,omitempty" url:"exclude,omitempty,optional"`                   // Exclude specified guest systems (assumes --all)
	Maxfiles         int    `json:"maxfiles,omitempty" url:"maxfiles,omitempty,optional"`                 // Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	Mode             string `json:"mode,omitempty" url:"mode,omitempty,optional"`                         // Backup mode.
	Quiet            bool   `json:"quiet,omitempty" url:"quiet,omitempty,optional"`                       // Be quiet.
	Bwlimit          int    `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`                   // Limit I/O bandwidth (KBytes per second).
	Enabled          bool   `json:"enabled,omitempty" url:"enabled,omitempty,optional"`                   // Enable or disable the job.
	Pigz             int    `json:"pigz,omitempty" url:"pigz,omitempty,optional"`                         // Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Remove           bool   `json:"remove,omitempty" url:"remove,omitempty,optional"`                     // Prune older backups according to 'prune-backups'.
	Script           string `json:"script,omitempty" url:"script,omitempty,optional"`                     // Use specified hook script.
	Starttime        string `json:"starttime,omitempty" url:"starttime,omitempty,optional"`               // Job Start time.
	Vmid             string `json:"vmid,omitempty" url:"vmid,omitempty,optional"`                         // The ID of the guest system you want to backup.
	All              bool   `json:"all,omitempty" url:"all,omitempty,optional"`                           // Backup all known guest systems on this host.
	Id               string `json:"id,omitempty" url:"id,omitempty,optional"`                             // Job ID (will be autogenerated).
	Mailnotification string `json:"mailnotification,omitempty" url:"mailnotification,omitempty,optional"` // Specify when to send an email
	Dow              string `json:"dow,omitempty" url:"dow,omitempty,optional"`                           // Day of week selection.
	NotesTemplate    string `json:"notes-template,omitempty" url:"notes-template,omitempty,optional"`     // Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future.
	Stdexcludes      bool   `json:"stdexcludes,omitempty" url:"stdexcludes,omitempty,optional"`           // Exclude temporary files and logs.
	Mailto           string `json:"mailto,omitempty" url:"mailto,omitempty,optional"`                     // Comma-separated list of email addresses or users that should receive email notifications.
	Dumpdir          string `json:"dumpdir,omitempty" url:"dumpdir,omitempty,optional"`                   // Store resulting files to specified directory.
	ExcludePath      string `json:"exclude-path,omitempty" url:"exclude-path,omitempty,optional"`         // Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root,  other paths match relative to each subdirectory.
	PruneBackups     string `json:"prune-backups,omitempty" url:"prune-backups,omitempty,optional"`       // Use these retention options instead of those from the storage configuration.
	Schedule         string `json:"schedule,omitempty" url:"schedule,omitempty,optional"`                 // Backup schedule. The format is a subset of `systemd` calendar events.
	Comment          string `json:"comment,omitempty" url:"comment,omitempty,optional"`                   // Description for the Job.
	Lockwait         int    `json:"lockwait,omitempty" url:"lockwait,omitempty,optional"`                 // Maximal time to wait for the global lock (minutes).
	Pool             string `json:"pool,omitempty" url:"pool,omitempty,optional"`                         // Backup all known guest systems included in the specified pool.
	Storage          string `json:"storage,omitempty" url:"storage,omitempty,optional"`                   // Store resulting file to this storage.
	Tmpdir           string `json:"tmpdir,omitempty" url:"tmpdir,omitempty,optional"`                     // Store temporary files to specified directory.
	Compress         string `json:"compress,omitempty" url:"compress,omitempty,optional"`                 // Compress dump file.
}

// ClusterBackupReadJobRequest {id} - Read vzdump backup job definition.
// Read vzdump backup job definition.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/backup/{id}
//
type ClusterBackupReadJobRequest struct {
	Id string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"` // The job ID.
}

// ClusterBackupUpdateJobRequest {id} - Update vzdump backup job definition.
// Update vzdump backup job definition.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/backup/{id}
//
type ClusterBackupUpdateJobRequest struct {
	Zstd             int    `json:"zstd,omitempty" url:"zstd,omitempty,optional"`                         // Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
	Compress         string `json:"compress,omitempty" url:"compress,omitempty,optional"`                 // Compress dump file.
	Dumpdir          string `json:"dumpdir,omitempty" url:"dumpdir,omitempty,optional"`                   // Store resulting files to specified directory.
	ExcludePath      string `json:"exclude-path,omitempty" url:"exclude-path,omitempty,optional"`         // Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root,  other paths match relative to each subdirectory.
	Mode             string `json:"mode,omitempty" url:"mode,omitempty,optional"`                         // Backup mode.
	Node             string `json:"node,omitempty" url:"node,omitempty,optional"`                         // Only run if executed on this node.
	Protected        bool   `json:"protected,omitempty" url:"protected,omitempty,optional"`               // If true, mark backup(s) as protected.
	Stopwait         int    `json:"stopwait,omitempty" url:"stopwait,omitempty,optional"`                 // Maximal time to wait until a guest system is stopped (minutes).
	Dow              string `json:"dow,omitempty" url:"dow,omitempty,optional"`                           // Day of week selection.
	Maxfiles         int    `json:"maxfiles,omitempty" url:"maxfiles,omitempty,optional"`                 // Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	Script           string `json:"script,omitempty" url:"script,omitempty,optional"`                     // Use specified hook script.
	Ionice           int    `json:"ionice,omitempty" url:"ionice,omitempty,optional"`                     // Set CFQ ionice priority.
	Lockwait         int    `json:"lockwait,omitempty" url:"lockwait,omitempty,optional"`                 // Maximal time to wait for the global lock (minutes).
	PruneBackups     string `json:"prune-backups,omitempty" url:"prune-backups,omitempty,optional"`       // Use these retention options instead of those from the storage configuration.
	Quiet            bool   `json:"quiet,omitempty" url:"quiet,omitempty,optional"`                       // Be quiet.
	Delete           string `json:"delete,omitempty" url:"delete,omitempty,optional"`                     // A list of settings you want to delete.
	Enabled          bool   `json:"enabled,omitempty" url:"enabled,omitempty,optional"`                   // Enable or disable the job.
	Id               string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`                   // The job ID.
	Mailto           string `json:"mailto,omitempty" url:"mailto,omitempty,optional"`                     // Comma-separated list of email addresses or users that should receive email notifications.
	Storage          string `json:"storage,omitempty" url:"storage,omitempty,optional"`                   // Store resulting file to this storage.
	Remove           bool   `json:"remove,omitempty" url:"remove,omitempty,optional"`                     // Prune older backups according to 'prune-backups'.
	Bwlimit          int    `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`                   // Limit I/O bandwidth (KBytes per second).
	Pigz             int    `json:"pigz,omitempty" url:"pigz,omitempty,optional"`                         // Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Starttime        string `json:"starttime,omitempty" url:"starttime,omitempty,optional"`               // Job Start time.
	Stop             bool   `json:"stop,omitempty" url:"stop,omitempty,optional"`                         // Stop running backup jobs on this host.
	Tmpdir           string `json:"tmpdir,omitempty" url:"tmpdir,omitempty,optional"`                     // Store temporary files to specified directory.
	Comment          string `json:"comment,omitempty" url:"comment,omitempty,optional"`                   // Description for the Job.
	Exclude          string `json:"exclude,omitempty" url:"exclude,omitempty,optional"`                   // Exclude specified guest systems (assumes --all)
	NotesTemplate    string `json:"notes-template,omitempty" url:"notes-template,omitempty,optional"`     // Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future.
	Pool             string `json:"pool,omitempty" url:"pool,omitempty,optional"`                         // Backup all known guest systems included in the specified pool.
	Vmid             string `json:"vmid,omitempty" url:"vmid,omitempty,optional"`                         // The ID of the guest system you want to backup.
	All              bool   `json:"all,omitempty" url:"all,omitempty,optional"`                           // Backup all known guest systems on this host.
	Mailnotification string `json:"mailnotification,omitempty" url:"mailnotification,omitempty,optional"` // Specify when to send an email
	Schedule         string `json:"schedule,omitempty" url:"schedule,omitempty,optional"`                 // Backup schedule. The format is a subset of `systemd` calendar events.
	Stdexcludes      bool   `json:"stdexcludes,omitempty" url:"stdexcludes,omitempty,optional"`           // Exclude temporary files and logs.
}

// ClusterBackupDeleteJobRequest {id} - Delete vzdump backup job definition.
// Delete vzdump backup job definition.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/backup/{id}
//
type ClusterBackupDeleteJobRequest struct {
	Id string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"` // The job ID.
}

// ClusterBackupIncludedVolumesGetVolumeBackupIncludedRequest included_volumes - Returns included guests and the backup status of their disks. Optimized to be used in ExtJS tree views.
// Returns included guests and the backup status of their disks. Optimized to be used in ExtJS tree views.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/backup/{id}/included_volumes
//
type ClusterBackupIncludedVolumesGetVolumeBackupIncludedRequest struct {
	Id string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"` // The job ID.
}

// ClusterBackupIncludedVolumesGetVolumeBackupIncludedResponse included_volumes
// Returns included guests and the backup status of their disks. Optimized to be used in ExtJS tree views.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/backup/{id}/included_volumes
//
type ClusterBackupIncludedVolumesGetVolumeBackupIncludedResponse struct {
	Children []struct {
	} `json:"children,omitempty" url:"children,omitempty" validate:"nonzero"` //
}

// ClusterBackupInfoNotBackedUpGetGuestsNotInBackupResponse not-backed-up
// Shows all guests which are not covered by any backup job.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/backup-info/not-backed-up
//
type ClusterBackupInfoNotBackedUpGetGuestsNotInBackupResponse struct {
	Name string `json:"name,omitempty" url:"name,omitempty,optional"`           // Name of the guest
	Type string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"` // Type of the guest.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // VMID of the guest.
}

// ClusterHaResourcesCreateRequest resources - Create a new HA resource.
// Create a new HA resource.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/resources
//
type ClusterHaResourcesCreateRequest struct {
	State       string `json:"state,omitempty" url:"state,omitempty,optional"`               // Requested resource state.
	Type        string `json:"type,omitempty" url:"type,omitempty,optional"`                 // Resource type.
	Comment     string `json:"comment,omitempty" url:"comment,omitempty,optional"`           // Description.
	Group       string `json:"group,omitempty" url:"group,omitempty,optional"`               // The HA group identifier.
	MaxRelocate int    `json:"max_relocate,omitempty" url:"max_relocate,omitempty,optional"` // Maximal number of service relocate tries when a service failes to start.
	MaxRestart  int    `json:"max_restart,omitempty" url:"max_restart,omitempty,optional"`   // Maximal number of tries to restart the service on a node after its start failed.
	Sid         string `json:"sid,omitempty" url:"sid,omitempty" validate:"nonzero"`         // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
}

// ClusterHaResourcesDeleteRequest {sid} - Delete resource configuration.
// Delete resource configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/resources/{sid}
//
type ClusterHaResourcesDeleteRequest struct {
	Sid string `json:"sid,omitempty" url:"sid,omitempty" validate:"nonzero"` // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
}

// ClusterHaResourcesReadRequest {sid} - Read resource configuration.
// Read resource configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/resources/{sid}
//
type ClusterHaResourcesReadRequest struct {
	Sid string `json:"sid,omitempty" url:"sid,omitempty" validate:"nonzero"` // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
}

// ClusterHaResourcesReadResponse {sid}
// Read resource configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/resources/{sid}
//
type ClusterHaResourcesReadResponse struct {
	MaxRestart  int    `json:"max_restart,omitempty" url:"max_restart,omitempty,optional"`   // Maximal number of tries to restart the service on a node after its start failed.
	Sid         string `json:"sid,omitempty" url:"sid,omitempty" validate:"nonzero"`         // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
	State       string `json:"state,omitempty" url:"state,omitempty,optional"`               // Requested resource state.
	Type        string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`       // The type of the resources.
	Comment     string `json:"comment,omitempty" url:"comment,omitempty,optional"`           // Description.
	Digest      string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"`   // Can be used to prevent concurrent modifications.
	Group       string `json:"group,omitempty" url:"group,omitempty,optional"`               // The HA group identifier.
	MaxRelocate int    `json:"max_relocate,omitempty" url:"max_relocate,omitempty,optional"` // Maximal number of service relocate tries when a service failes to start.
}

// ClusterHaResourcesUpdateRequest {sid} - Update resource configuration.
// Update resource configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/resources/{sid}
//
type ClusterHaResourcesUpdateRequest struct {
	State       string `json:"state,omitempty" url:"state,omitempty,optional"`               // Requested resource state.
	Comment     string `json:"comment,omitempty" url:"comment,omitempty,optional"`           // Description.
	Delete      string `json:"delete,omitempty" url:"delete,omitempty,optional"`             // A list of settings you want to delete.
	Digest      string `json:"digest,omitempty" url:"digest,omitempty,optional"`             // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Group       string `json:"group,omitempty" url:"group,omitempty,optional"`               // The HA group identifier.
	MaxRelocate int    `json:"max_relocate,omitempty" url:"max_relocate,omitempty,optional"` // Maximal number of service relocate tries when a service failes to start.
	MaxRestart  int    `json:"max_restart,omitempty" url:"max_restart,omitempty,optional"`   // Maximal number of tries to restart the service on a node after its start failed.
	Sid         string `json:"sid,omitempty" url:"sid,omitempty" validate:"nonzero"`         // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
}

// ClusterHaResourcesMigrateRequest migrate - Request resource migration (online) to another node.
// Request resource migration (online) to another node.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/resources/{sid}/migrate
//
type ClusterHaResourcesMigrateRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // Target node.
	Sid  string `json:"sid,omitempty" url:"sid,omitempty" validate:"nonzero"`   // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
}

// ClusterHaResourcesRelocateRequest relocate - Request resource relocatzion to another node. This stops the service on the old node, and restarts it on the target node.
// Request resource relocatzion to another node. This stops the service on the old node, and restarts it on the target node.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/resources/{sid}/relocate
//
type ClusterHaResourcesRelocateRequest struct {
	Sid  string `json:"sid,omitempty" url:"sid,omitempty" validate:"nonzero"`   // HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // Target node.
}

// ClusterHaGroupsCreateRequest groups - Create a new HA group.
// Create a new HA group.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/groups
//
type ClusterHaGroupsCreateRequest struct {
	Comment    string `json:"comment,omitempty" url:"comment,omitempty,optional"`       // Description.
	Group      string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"` // The HA group identifier.
	Nodes      string `json:"nodes,omitempty" url:"nodes,omitempty" validate:"nonzero"` // List of cluster node names with optional priority.
	Nofailback bool   `json:"nofailback,omitempty" url:"nofailback,omitempty,optional"` // The CRM tries to run services on the node with the highest priority. If a node with higher priority comes online, the CRM migrates the service to that node. Enabling nofailback prevents that behavior.
	Restricted bool   `json:"restricted,omitempty" url:"restricted,omitempty,optional"` // Resources bound to restricted groups may only run on nodes defined by the group.
	Type       string `json:"type,omitempty" url:"type,omitempty,optional"`             // Group type.
}

// ClusterHaGroupsDeleteRequest {group} - Delete ha group configuration.
// Delete ha group configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/groups/{group}
//
type ClusterHaGroupsDeleteRequest struct {
	Group string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"` // The HA group identifier.
}

// ClusterHaGroupsReadRequest {group} - Read ha group configuration.
// Read ha group configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/groups/{group}
//
type ClusterHaGroupsReadRequest struct {
	Group string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"` // The HA group identifier.
}

// ClusterHaGroupsUpdateRequest {group} - Update ha group configuration.
// Update ha group configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ha/groups/{group}
//
type ClusterHaGroupsUpdateRequest struct {
	Delete     string `json:"delete,omitempty" url:"delete,omitempty,optional"`         // A list of settings you want to delete.
	Digest     string `json:"digest,omitempty" url:"digest,omitempty,optional"`         // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Group      string `json:"group,omitempty" url:"group,omitempty" validate:"nonzero"` // The HA group identifier.
	Nodes      string `json:"nodes,omitempty" url:"nodes,omitempty,optional"`           // List of cluster node names with optional priority.
	Nofailback bool   `json:"nofailback,omitempty" url:"nofailback,omitempty,optional"` // The CRM tries to run services on the node with the highest priority. If a node with higher priority comes online, the CRM migrates the service to that node. Enabling nofailback prevents that behavior.
	Restricted bool   `json:"restricted,omitempty" url:"restricted,omitempty,optional"` // Resources bound to restricted groups may only run on nodes defined by the group.
	Comment    string `json:"comment,omitempty" url:"comment,omitempty,optional"`       // Description.
}

// ClusterAcmePluginsAddPluginRequest plugins - Add ACME plugin configuration.
// Add ACME plugin configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/plugins
//
type ClusterAcmePluginsAddPluginRequest struct {
	Type            string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`               // ACME challenge type.
	ValidationDelay int    `json:"validation-delay,omitempty" url:"validation-delay,omitempty,optional"` // Extra delay in seconds to wait before requesting validation. Allows to cope with a long TTL of DNS records.
	Api             string `json:"api,omitempty" url:"api,omitempty,optional"`                           // API plugin name
	Data            string `json:"data,omitempty" url:"data,omitempty,optional"`                         // DNS plugin data. (base64 encoded)
	Disable         bool   `json:"disable,omitempty" url:"disable,omitempty,optional"`                   // Flag to disable the config.
	Id              string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`                   // ACME Plugin ID name
	Nodes           string `json:"nodes,omitempty" url:"nodes,omitempty,optional"`                       // List of cluster node names.
}

// ClusterAcmePluginsGetPluginConfigRequest {id} - Get ACME plugin configuration.
// Get ACME plugin configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/plugins/{id}
//
type ClusterAcmePluginsGetPluginConfigRequest struct {
	Id string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"` // Unique identifier for ACME plugin instance.
}

// ClusterAcmePluginsUpdatePluginRequest {id} - Update ACME plugin configuration.
// Update ACME plugin configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/plugins/{id}
//
type ClusterAcmePluginsUpdatePluginRequest struct {
	ValidationDelay int    `json:"validation-delay,omitempty" url:"validation-delay,omitempty,optional"` // Extra delay in seconds to wait before requesting validation. Allows to cope with a long TTL of DNS records.
	Api             string `json:"api,omitempty" url:"api,omitempty,optional"`                           // API plugin name
	Data            string `json:"data,omitempty" url:"data,omitempty,optional"`                         // DNS plugin data. (base64 encoded)
	Delete          string `json:"delete,omitempty" url:"delete,omitempty,optional"`                     // A list of settings you want to delete.
	Digest          string `json:"digest,omitempty" url:"digest,omitempty,optional"`                     // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Disable         bool   `json:"disable,omitempty" url:"disable,omitempty,optional"`                   // Flag to disable the config.
	Id              string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`                   // ACME Plugin ID name
	Nodes           string `json:"nodes,omitempty" url:"nodes,omitempty,optional"`                       // List of cluster node names.
}

// ClusterAcmePluginsDeletePluginRequest {id} - Delete ACME plugin configuration.
// Delete ACME plugin configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/plugins/{id}
//
type ClusterAcmePluginsDeletePluginRequest struct {
	Id string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"` // Unique identifier for ACME plugin instance.
}

// ClusterAcmeAccountRegisterRequest account - Register a new ACME account with CA.
// Register a new ACME account with CA.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/account
//
type ClusterAcmeAccountRegisterRequest struct {
	Name      string `json:"name,omitempty" url:"name,omitempty,optional"`                 // ACME account config file name.
	TosUrl    string `json:"tos_url,omitempty" url:"tos_url,omitempty,optional"`           // URL of CA TermsOfService - setting this indicates agreement.
	Contact   string `json:"contact,omitempty" url:"contact,omitempty" validate:"nonzero"` // Contact email addresses.
	Directory string `json:"directory,omitempty" url:"directory,omitempty,optional"`       // URL of ACME CA directory endpoint.
}

// ClusterAcmeAccountDeactivateRequest {name} - Deactivate existing ACME account at CA.
// Deactivate existing ACME account at CA.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/account/{name}
//
type ClusterAcmeAccountDeactivateRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty,optional"` // ACME account config file name.
}

// ClusterAcmeAccountGetRequest {name} - Return existing ACME account information.
// Return existing ACME account information.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/account/{name}
//
type ClusterAcmeAccountGetRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty,optional"` // ACME account config file name.
}

// ClusterAcmeAccountGetResponse {name}
// Return existing ACME account information.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/account/{name}
//
type ClusterAcmeAccountGetResponse struct {
	Location string   `json:"location,omitempty" url:"location,omitempty,optional"` //
	Tos      string   `json:"tos,omitempty" url:"tos,omitempty,optional"`           //
	Account  struct { //
	} `json:"account,omitempty" url:"account,omitempty,optional"` //
	Directory string `json:"directory,omitempty" url:"directory,omitempty,optional"` // URL of ACME CA directory endpoint.
}

// ClusterAcmeAccountUpdateRequest {name} - Update existing ACME account information with CA. Note: not specifying any new account information triggers a refresh.
// Update existing ACME account information with CA. Note: not specifying any new account information triggers a refresh.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/account/{name}
//
type ClusterAcmeAccountUpdateRequest struct {
	Contact string `json:"contact,omitempty" url:"contact,omitempty,optional"` // Contact email addresses.
	Name    string `json:"name,omitempty" url:"name,omitempty,optional"`       // ACME account config file name.
}

// ClusterAcmeTosGetRequest tos - Retrieve ACME TermsOfService URL from CA.
// Retrieve ACME TermsOfService URL from CA.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/tos
//
type ClusterAcmeTosGetRequest struct {
	Directory string `json:"directory,omitempty" url:"directory,omitempty,optional"` // URL of ACME CA directory endpoint.
}

// ClusterAcmeDirectoriesGetResponse directories
// Get named known ACME directory endpoints.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/directories
//
type ClusterAcmeDirectoriesGetResponse struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` //
	Url  string `json:"url,omitempty" url:"url,omitempty" validate:"nonzero"`   // URL of ACME CA directory endpoint.
}

// ClusterAcmeChallengeSchemaChallengeschemaResponse challenge-schema
// Get schema of ACME challenge types.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/acme/challenge-schema
//
type ClusterAcmeChallengeSchemaChallengeschemaResponse struct {
	Name   string   `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Human readable name, falls back to id
	Schema struct { //
	} `json:"schema,omitempty" url:"schema,omitempty" validate:"nonzero"` //
	Type string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"` //
	Id   string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`     //
}

// ClusterCephMetadataRequest metadata - Get ceph metadata.
// Get ceph metadata.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ceph/metadata
//
type ClusterCephMetadataRequest struct {
	Scope string `json:"scope,omitempty" url:"scope,omitempty,optional"` //
}

// ClusterCephFlagsGetAllResponse flags
// get the status of all ceph flags
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ceph/flags
//
type ClusterCephFlagsGetAllResponse struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Flag name.
}

// ClusterCephFlagsSetRequest flags - Set/Unset multiple ceph flags at once.
// Set/Unset multiple ceph flags at once.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ceph/flags
//
type ClusterCephFlagsSetRequest struct {
	Notieragent bool `json:"notieragent,omitempty" url:"notieragent,omitempty,optional"`   // Cache tiering activity is suspended.
	Nobackfill  bool `json:"nobackfill,omitempty" url:"nobackfill,omitempty,optional"`     // Backfilling of PGs is suspended.
	NodeepScrub bool `json:"nodeep-scrub,omitempty" url:"nodeep-scrub,omitempty,optional"` // Deep Scrubbing is disabled.
	Noout       bool `json:"noout,omitempty" url:"noout,omitempty,optional"`               // OSDs will not automatically be marked out after the configured interval.
	Norebalance bool `json:"norebalance,omitempty" url:"norebalance,omitempty,optional"`   // Rebalancing of PGs is suspended.
	Norecover   bool `json:"norecover,omitempty" url:"norecover,omitempty,optional"`       // Recovery of PGs is suspended.
	Noscrub     bool `json:"noscrub,omitempty" url:"noscrub,omitempty,optional"`           // Scrubbing is disabled.
	Nodown      bool `json:"nodown,omitempty" url:"nodown,omitempty,optional"`             // OSD failure reports are being ignored, such that the monitors will not mark OSDs down.
	Noin        bool `json:"noin,omitempty" url:"noin,omitempty,optional"`                 // OSDs that were previously marked out will not be marked back in when they start.
	Noup        bool `json:"noup,omitempty" url:"noup,omitempty,optional"`                 // OSDs are not allowed to start.
	Pause       bool `json:"pause,omitempty" url:"pause,omitempty,optional"`               // Pauses read and writes.
}

// ClusterCephFlagsGetFlagRequest {flag} - Get the status of a specific ceph flag.
// Get the status of a specific ceph flag.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ceph/flags/{flag}
//
type ClusterCephFlagsGetFlagRequest struct {
	Flag string `json:"flag,omitempty" url:"flag,omitempty" validate:"nonzero"` // The name of the flag name to get.
}

// ClusterCephFlagsUpdateFlagRequest {flag} - Set or clear (unset) a specific ceph flag
// Set or clear (unset) a specific ceph flag
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/ceph/flags/{flag}
//
type ClusterCephFlagsUpdateFlagRequest struct {
	Flag  string `json:"flag,omitempty" url:"flag,omitempty" validate:"nonzero"`   // The ceph flag to update
	Value bool   `json:"value,omitempty" url:"value,omitempty" validate:"nonzero"` // The new value of the flag
}

// ClusterJobsScheduleAnalyzeRequest schedule-analyze - Returns a list of future schedule runtimes.
// Returns a list of future schedule runtimes.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/jobs/schedule-analyze
//
type ClusterJobsScheduleAnalyzeRequest struct {
	Iterations int    `json:"iterations,omitempty" url:"iterations,omitempty,optional"`       // Number of event-iteration to simulate and return.
	Schedule   string `json:"schedule,omitempty" url:"schedule,omitempty" validate:"nonzero"` // Job schedule. The format is a subset of `systemd` calendar events.
	Starttime  int    `json:"starttime,omitempty" url:"starttime,omitempty,optional"`         // UNIX timestamp to start the calculation from. Defaults to the current time.
}

// ClusterJobsScheduleAnalyzeResponse schedule-analyze
// Returns a list of future schedule runtimes.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/jobs/schedule-analyze
//
type ClusterJobsScheduleAnalyzeResponse struct {
	Timestamp int    `json:"timestamp,omitempty" url:"timestamp,omitempty" validate:"nonzero"` // UNIX timestamp for the run.
	Utc       string `json:"utc,omitempty" url:"utc,omitempty" validate:"nonzero"`             // UTC timestamp for the run.
}

// ClusterSdnVnetsCreateRequest vnets - Create a new sdn vnet object.
// Create a new sdn vnet object.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/vnets
//
type ClusterSdnVnetsCreateRequest struct {
	Zone      string `json:"zone,omitempty" url:"zone,omitempty" validate:"nonzero"` // zone id
	Alias     string `json:"alias,omitempty" url:"alias,omitempty,optional"`         // alias name of the vnet
	Tag       int    `json:"tag,omitempty" url:"tag,omitempty,optional"`             // vlan or vxlan id
	Type      string `json:"type,omitempty" url:"type,omitempty,optional"`           // Type
	Vlanaware bool   `json:"vlanaware,omitempty" url:"vlanaware,omitempty,optional"` // Allow vm VLANs to pass through this vnet.
	Vnet      string `json:"vnet,omitempty" url:"vnet,omitempty" validate:"nonzero"` // The SDN vnet object identifier.
}

// ClusterSdnVnetsDeleteRequest {vnet} - Delete sdn vnet object configuration.
// Delete sdn vnet object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/vnets/{vnet}
//
type ClusterSdnVnetsDeleteRequest struct {
	Vnet string `json:"vnet,omitempty" url:"vnet,omitempty" validate:"nonzero"` // The SDN vnet object identifier.
}

// ClusterSdnVnetsReadRequest {vnet} - Read sdn vnet configuration.
// Read sdn vnet configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/vnets/{vnet}
//
type ClusterSdnVnetsReadRequest struct {
	Pending bool   `json:"pending,omitempty" url:"pending,omitempty,optional"`     // Display pending config.
	Running bool   `json:"running,omitempty" url:"running,omitempty,optional"`     // Display running config.
	Vnet    string `json:"vnet,omitempty" url:"vnet,omitempty" validate:"nonzero"` // The SDN vnet object identifier.
}

// ClusterSdnVnetsUpdateRequest {vnet} - Update sdn vnet object configuration.
// Update sdn vnet object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/vnets/{vnet}
//
type ClusterSdnVnetsUpdateRequest struct {
	Digest    string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Tag       int    `json:"tag,omitempty" url:"tag,omitempty,optional"`             // vlan or vxlan id
	Vlanaware bool   `json:"vlanaware,omitempty" url:"vlanaware,omitempty,optional"` // Allow vm VLANs to pass through this vnet.
	Vnet      string `json:"vnet,omitempty" url:"vnet,omitempty" validate:"nonzero"` // The SDN vnet object identifier.
	Zone      string `json:"zone,omitempty" url:"zone,omitempty,optional"`           // zone id
	Alias     string `json:"alias,omitempty" url:"alias,omitempty,optional"`         // alias name of the vnet
	Delete    string `json:"delete,omitempty" url:"delete,omitempty,optional"`       // A list of settings you want to delete.
}

// ClusterSdnVnetsSubnetsCreateRequest subnets - Create a new sdn subnet object.
// Create a new sdn subnet object.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/vnets/{vnet}/subnets
//
type ClusterSdnVnetsSubnetsCreateRequest struct {
	Snat          bool   `json:"snat,omitempty" url:"snat,omitempty,optional"`                   // enable masquerade for this subnet if pve-firewall
	Subnet        string `json:"subnet,omitempty" url:"subnet,omitempty" validate:"nonzero"`     // The SDN subnet object identifier.
	Type          string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`         //
	Vnet          string `json:"vnet,omitempty" url:"vnet,omitempty" validate:"nonzero"`         // associated vnet
	Dnszoneprefix string `json:"dnszoneprefix,omitempty" url:"dnszoneprefix,omitempty,optional"` // dns domain zone prefix  ex: 'adm' -> <hostname>.adm.mydomain.com
	Gateway       string `json:"gateway,omitempty" url:"gateway,omitempty,optional"`             // Subnet Gateway: Will be assign on vnet for layer3 zones
}

// ClusterSdnVnetsSubnetsReadRequest {subnet} - Read sdn subnet configuration.
// Read sdn subnet configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/vnets/{vnet}/subnets/{subnet}
//
type ClusterSdnVnetsSubnetsReadRequest struct {
	Pending bool   `json:"pending,omitempty" url:"pending,omitempty,optional"`         // Display pending config.
	Running bool   `json:"running,omitempty" url:"running,omitempty,optional"`         // Display running config.
	Subnet  string `json:"subnet,omitempty" url:"subnet,omitempty" validate:"nonzero"` // The SDN subnet object identifier.
	Vnet    string `json:"vnet,omitempty" url:"vnet,omitempty" validate:"nonzero"`     // The SDN vnet object identifier.
}

// ClusterSdnVnetsSubnetsUpdateRequest {subnet} - Update sdn subnet object configuration.
// Update sdn subnet object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/vnets/{vnet}/subnets/{subnet}
//
type ClusterSdnVnetsSubnetsUpdateRequest struct {
	Delete        string `json:"delete,omitempty" url:"delete,omitempty,optional"`               // A list of settings you want to delete.
	Digest        string `json:"digest,omitempty" url:"digest,omitempty,optional"`               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Dnszoneprefix string `json:"dnszoneprefix,omitempty" url:"dnszoneprefix,omitempty,optional"` // dns domain zone prefix  ex: 'adm' -> <hostname>.adm.mydomain.com
	Gateway       string `json:"gateway,omitempty" url:"gateway,omitempty,optional"`             // Subnet Gateway: Will be assign on vnet for layer3 zones
	Snat          bool   `json:"snat,omitempty" url:"snat,omitempty,optional"`                   // enable masquerade for this subnet if pve-firewall
	Subnet        string `json:"subnet,omitempty" url:"subnet,omitempty" validate:"nonzero"`     // The SDN subnet object identifier.
	Vnet          string `json:"vnet,omitempty" url:"vnet,omitempty,optional"`                   // associated vnet
}

// ClusterSdnVnetsSubnetsDeleteRequest {subnet} - Delete sdn subnet object configuration.
// Delete sdn subnet object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/vnets/{vnet}/subnets/{subnet}
//
type ClusterSdnVnetsSubnetsDeleteRequest struct {
	Subnet string `json:"subnet,omitempty" url:"subnet,omitempty" validate:"nonzero"` // The SDN subnet object identifier.
	Vnet   string `json:"vnet,omitempty" url:"vnet,omitempty" validate:"nonzero"`     // The SDN vnet object identifier.
}

// ClusterSdnZonesCreateRequest zones - Create a new sdn zone object.
// Create a new sdn zone object.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/zones
//
type ClusterSdnZonesCreateRequest struct {
	Reversedns               string `json:"reversedns,omitempty" url:"reversedns,omitempty,optional"`                                   // reverse dns api server
	Type                     string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`                                     // Plugin type.
	VlanProtocol             string `json:"vlan-protocol,omitempty" url:"vlan-protocol,omitempty,optional"`                             //
	Zone                     string `json:"zone,omitempty" url:"zone,omitempty" validate:"nonzero"`                                     // The SDN zone object identifier.
	Exitnodes                string `json:"exitnodes,omitempty" url:"exitnodes,omitempty,optional"`                                     // List of cluster node names.
	ExitnodesLocalRouting    bool   `json:"exitnodes-local-routing,omitempty" url:"exitnodes-local-routing,omitempty,optional"`         // Allow exitnodes to connect to evpn guests
	Ipam                     string `json:"ipam,omitempty" url:"ipam,omitempty,optional"`                                               // use a specific ipam
	Peers                    string `json:"peers,omitempty" url:"peers,omitempty,optional"`                                             // peers address list.
	RtImport                 string `json:"rt-import,omitempty" url:"rt-import,omitempty,optional"`                                     // Route-Target import
	Tag                      int    `json:"tag,omitempty" url:"tag,omitempty,optional"`                                                 // Service-VLAN Tag
	Bridge                   string `json:"bridge,omitempty" url:"bridge,omitempty,optional"`                                           //
	DpId                     int    `json:"dp-id,omitempty" url:"dp-id,omitempty,optional"`                                             // Faucet dataplane id
	Mac                      string `json:"mac,omitempty" url:"mac,omitempty,optional"`                                                 // Anycast logical router mac address
	VrfVxlan                 int    `json:"vrf-vxlan,omitempty" url:"vrf-vxlan,omitempty,optional"`                                     // l3vni.
	BridgeDisableMacLearning bool   `json:"bridge-disable-mac-learning,omitempty" url:"bridge-disable-mac-learning,omitempty,optional"` // Disable auto mac learning.
	Controller               string `json:"controller,omitempty" url:"controller,omitempty,optional"`                                   // Frr router name
	Dnszone                  string `json:"dnszone,omitempty" url:"dnszone,omitempty,optional"`                                         // dns domain zone  ex: mydomain.com
	ExitnodesPrimary         string `json:"exitnodes-primary,omitempty" url:"exitnodes-primary,omitempty,optional"`                     // Force traffic to this exitnode first.
	Mtu                      int    `json:"mtu,omitempty" url:"mtu,omitempty,optional"`                                                 // MTU
	Nodes                    string `json:"nodes,omitempty" url:"nodes,omitempty,optional"`                                             // List of cluster node names.
	AdvertiseSubnets         bool   `json:"advertise-subnets,omitempty" url:"advertise-subnets,omitempty,optional"`                     // Advertise evpn subnets if you have silent hosts
	DisableArpNdSuppression  bool   `json:"disable-arp-nd-suppression,omitempty" url:"disable-arp-nd-suppression,omitempty,optional"`   // Disable ipv4 arp && ipv6 neighbour discovery suppression
	Dns                      string `json:"dns,omitempty" url:"dns,omitempty,optional"`                                                 // dns api server
}

// ClusterSdnZonesDeleteRequest {zone} - Delete sdn zone object configuration.
// Delete sdn zone object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/zones/{zone}
//
type ClusterSdnZonesDeleteRequest struct {
	Zone string `json:"zone,omitempty" url:"zone,omitempty" validate:"nonzero"` // The SDN zone object identifier.
}

// ClusterSdnZonesReadRequest {zone} - Read sdn zone configuration.
// Read sdn zone configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/zones/{zone}
//
type ClusterSdnZonesReadRequest struct {
	Pending bool   `json:"pending,omitempty" url:"pending,omitempty,optional"`     // Display pending config.
	Running bool   `json:"running,omitempty" url:"running,omitempty,optional"`     // Display running config.
	Zone    string `json:"zone,omitempty" url:"zone,omitempty" validate:"nonzero"` // The SDN zone object identifier.
}

// ClusterSdnZonesUpdateRequest {zone} - Update sdn zone object configuration.
// Update sdn zone object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/zones/{zone}
//
type ClusterSdnZonesUpdateRequest struct {
	ExitnodesLocalRouting    bool   `json:"exitnodes-local-routing,omitempty" url:"exitnodes-local-routing,omitempty,optional"`         // Allow exitnodes to connect to evpn guests
	Peers                    string `json:"peers,omitempty" url:"peers,omitempty,optional"`                                             // peers address list.
	RtImport                 string `json:"rt-import,omitempty" url:"rt-import,omitempty,optional"`                                     // Route-Target import
	Dns                      string `json:"dns,omitempty" url:"dns,omitempty,optional"`                                                 // dns api server
	ExitnodesPrimary         string `json:"exitnodes-primary,omitempty" url:"exitnodes-primary,omitempty,optional"`                     // Force traffic to this exitnode first.
	Ipam                     string `json:"ipam,omitempty" url:"ipam,omitempty,optional"`                                               // use a specific ipam
	Nodes                    string `json:"nodes,omitempty" url:"nodes,omitempty,optional"`                                             // List of cluster node names.
	VrfVxlan                 int    `json:"vrf-vxlan,omitempty" url:"vrf-vxlan,omitempty,optional"`                                     // l3vni.
	Delete                   string `json:"delete,omitempty" url:"delete,omitempty,optional"`                                           // A list of settings you want to delete.
	Controller               string `json:"controller,omitempty" url:"controller,omitempty,optional"`                                   // Frr router name
	Digest                   string `json:"digest,omitempty" url:"digest,omitempty,optional"`                                           // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	DisableArpNdSuppression  bool   `json:"disable-arp-nd-suppression,omitempty" url:"disable-arp-nd-suppression,omitempty,optional"`   // Disable ipv4 arp && ipv6 neighbour discovery suppression
	Dnszone                  string `json:"dnszone,omitempty" url:"dnszone,omitempty,optional"`                                         // dns domain zone  ex: mydomain.com
	Mtu                      int    `json:"mtu,omitempty" url:"mtu,omitempty,optional"`                                                 // MTU
	Reversedns               string `json:"reversedns,omitempty" url:"reversedns,omitempty,optional"`                                   // reverse dns api server
	Tag                      int    `json:"tag,omitempty" url:"tag,omitempty,optional"`                                                 // Service-VLAN Tag
	BridgeDisableMacLearning bool   `json:"bridge-disable-mac-learning,omitempty" url:"bridge-disable-mac-learning,omitempty,optional"` // Disable auto mac learning.
	Zone                     string `json:"zone,omitempty" url:"zone,omitempty" validate:"nonzero"`                                     // The SDN zone object identifier.
	VlanProtocol             string `json:"vlan-protocol,omitempty" url:"vlan-protocol,omitempty,optional"`                             //
	Bridge                   string `json:"bridge,omitempty" url:"bridge,omitempty,optional"`                                           //
	DpId                     int    `json:"dp-id,omitempty" url:"dp-id,omitempty,optional"`                                             // Faucet dataplane id
	Exitnodes                string `json:"exitnodes,omitempty" url:"exitnodes,omitempty,optional"`                                     // List of cluster node names.
	Mac                      string `json:"mac,omitempty" url:"mac,omitempty,optional"`                                                 // Anycast logical router mac address
	AdvertiseSubnets         bool   `json:"advertise-subnets,omitempty" url:"advertise-subnets,omitempty,optional"`                     // Advertise evpn subnets if you have silent hosts
}

// ClusterSdnControllersCreateRequest controllers - Create a new sdn controller object.
// Create a new sdn controller object.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/controllers
//
type ClusterSdnControllersCreateRequest struct {
	EbgpMultihop            int    `json:"ebgp-multihop,omitempty" url:"ebgp-multihop,omitempty,optional"`                             //
	Peers                   string `json:"peers,omitempty" url:"peers,omitempty,optional"`                                             // peers address list.
	Ebgp                    bool   `json:"ebgp,omitempty" url:"ebgp,omitempty,optional"`                                               // Enable ebgp. (remote-as external)
	BgpMultipathAsPathRelax bool   `json:"bgp-multipath-as-path-relax,omitempty" url:"bgp-multipath-as-path-relax,omitempty,optional"` //
	Controller              string `json:"controller,omitempty" url:"controller,omitempty" validate:"nonzero"`                         // The SDN controller object identifier.
	Loopback                string `json:"loopback,omitempty" url:"loopback,omitempty,optional"`                                       // source loopback interface.
	Node                    string `json:"node,omitempty" url:"node,omitempty,optional"`                                               // The cluster node name.
	Type                    string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`                                     // Plugin type.
	Asn                     int    `json:"asn,omitempty" url:"asn,omitempty,optional"`                                                 // autonomous system number
}

// ClusterSdnControllersDeleteRequest {controller} - Delete sdn controller object configuration.
// Delete sdn controller object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/controllers/{controller}
//
type ClusterSdnControllersDeleteRequest struct {
	Controller string `json:"controller,omitempty" url:"controller,omitempty" validate:"nonzero"` // The SDN controller object identifier.
}

// ClusterSdnControllersReadRequest {controller} - Read sdn controller configuration.
// Read sdn controller configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/controllers/{controller}
//
type ClusterSdnControllersReadRequest struct {
	Pending    bool   `json:"pending,omitempty" url:"pending,omitempty,optional"`                 // Display pending config.
	Running    bool   `json:"running,omitempty" url:"running,omitempty,optional"`                 // Display running config.
	Controller string `json:"controller,omitempty" url:"controller,omitempty" validate:"nonzero"` // The SDN controller object identifier.
}

// ClusterSdnControllersUpdateRequest {controller} - Update sdn controller object configuration.
// Update sdn controller object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/controllers/{controller}
//
type ClusterSdnControllersUpdateRequest struct {
	Digest                  string `json:"digest,omitempty" url:"digest,omitempty,optional"`                                           // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Peers                   string `json:"peers,omitempty" url:"peers,omitempty,optional"`                                             // peers address list.
	Controller              string `json:"controller,omitempty" url:"controller,omitempty" validate:"nonzero"`                         // The SDN controller object identifier.
	Delete                  string `json:"delete,omitempty" url:"delete,omitempty,optional"`                                           // A list of settings you want to delete.
	Ebgp                    bool   `json:"ebgp,omitempty" url:"ebgp,omitempty,optional"`                                               // Enable ebgp. (remote-as external)
	EbgpMultihop            int    `json:"ebgp-multihop,omitempty" url:"ebgp-multihop,omitempty,optional"`                             //
	Loopback                string `json:"loopback,omitempty" url:"loopback,omitempty,optional"`                                       // source loopback interface.
	Node                    string `json:"node,omitempty" url:"node,omitempty,optional"`                                               // The cluster node name.
	Asn                     int    `json:"asn,omitempty" url:"asn,omitempty,optional"`                                                 // autonomous system number
	BgpMultipathAsPathRelax bool   `json:"bgp-multipath-as-path-relax,omitempty" url:"bgp-multipath-as-path-relax,omitempty,optional"` //
}

// ClusterSdnIpamsCreateRequest ipams - Create a new sdn ipam object.
// Create a new sdn ipam object.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/ipams
//
type ClusterSdnIpamsCreateRequest struct {
	Ipam    string `json:"ipam,omitempty" url:"ipam,omitempty" validate:"nonzero"` // The SDN ipam object identifier.
	Section int    `json:"section,omitempty" url:"section,omitempty,optional"`     //
	Token   string `json:"token,omitempty" url:"token,omitempty,optional"`         //
	Type    string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"` // Plugin type.
	Url     string `json:"url,omitempty" url:"url,omitempty,optional"`             //
}

// ClusterSdnIpamsUpdateRequest {ipam} - Update sdn ipam object configuration.
// Update sdn ipam object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/ipams/{ipam}
//
type ClusterSdnIpamsUpdateRequest struct {
	Delete  string `json:"delete,omitempty" url:"delete,omitempty,optional"`       // A list of settings you want to delete.
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Ipam    string `json:"ipam,omitempty" url:"ipam,omitempty" validate:"nonzero"` // The SDN ipam object identifier.
	Section int    `json:"section,omitempty" url:"section,omitempty,optional"`     //
	Token   string `json:"token,omitempty" url:"token,omitempty,optional"`         //
	Url     string `json:"url,omitempty" url:"url,omitempty,optional"`             //
}

// ClusterSdnIpamsDeleteRequest {ipam} - Delete sdn ipam object configuration.
// Delete sdn ipam object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/ipams/{ipam}
//
type ClusterSdnIpamsDeleteRequest struct {
	Ipam string `json:"ipam,omitempty" url:"ipam,omitempty" validate:"nonzero"` // The SDN ipam object identifier.
}

// ClusterSdnIpamsReadRequest {ipam} - Read sdn ipam configuration.
// Read sdn ipam configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/ipams/{ipam}
//
type ClusterSdnIpamsReadRequest struct {
	Ipam string `json:"ipam,omitempty" url:"ipam,omitempty" validate:"nonzero"` // The SDN ipam object identifier.
}

// ClusterSdnDnsCreateRequest dns - Create a new sdn dns object.
// Create a new sdn dns object.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/dns
//
type ClusterSdnDnsCreateRequest struct {
	Dns           string `json:"dns,omitempty" url:"dns,omitempty" validate:"nonzero"`           // The SDN dns object identifier.
	Key           string `json:"key,omitempty" url:"key,omitempty" validate:"nonzero"`           //
	Reversemaskv6 int    `json:"reversemaskv6,omitempty" url:"reversemaskv6,omitempty,optional"` //
	Reversev6mask int    `json:"reversev6mask,omitempty" url:"reversev6mask,omitempty,optional"` //
	Ttl           int    `json:"ttl,omitempty" url:"ttl,omitempty,optional"`                     //
	Type          string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`         // Plugin type.
	Url           string `json:"url,omitempty" url:"url,omitempty" validate:"nonzero"`           //
}

// ClusterSdnDnsDeleteRequest {dns} - Delete sdn dns object configuration.
// Delete sdn dns object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/dns/{dns}
//
type ClusterSdnDnsDeleteRequest struct {
	Dns string `json:"dns,omitempty" url:"dns,omitempty" validate:"nonzero"` // The SDN dns object identifier.
}

// ClusterSdnDnsReadRequest {dns} - Read sdn dns configuration.
// Read sdn dns configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/dns/{dns}
//
type ClusterSdnDnsReadRequest struct {
	Dns string `json:"dns,omitempty" url:"dns,omitempty" validate:"nonzero"` // The SDN dns object identifier.
}

// ClusterSdnDnsUpdateRequest {dns} - Update sdn dns object configuration.
// Update sdn dns object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/sdn/dns/{dns}
//
type ClusterSdnDnsUpdateRequest struct {
	Dns           string `json:"dns,omitempty" url:"dns,omitempty" validate:"nonzero"`           // The SDN dns object identifier.
	Key           string `json:"key,omitempty" url:"key,omitempty,optional"`                     //
	Reversemaskv6 int    `json:"reversemaskv6,omitempty" url:"reversemaskv6,omitempty,optional"` //
	Ttl           int    `json:"ttl,omitempty" url:"ttl,omitempty,optional"`                     //
	Url           string `json:"url,omitempty" url:"url,omitempty,optional"`                     //
	Delete        string `json:"delete,omitempty" url:"delete,omitempty,optional"`               // A list of settings you want to delete.
	Digest        string `json:"digest,omitempty" url:"digest,omitempty,optional"`               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}

// ClusterLogRequest log - Read cluster log
// Read cluster log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/log
//
type ClusterLogRequest struct {
	Max int `json:"max,omitempty" url:"max,omitempty,optional"` // Maximum number of entries.
}

// ClusterResourcesRequest resources - Resources index (cluster wide).
// Resources index (cluster wide).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/resources
//
type ClusterResourcesRequest struct {
	Type string `json:"type,omitempty" url:"type,omitempty,optional"` //
}

// ClusterResourcesResponse resources
// Resources index (cluster wide).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/resources
//
type ClusterResourcesResponse struct {
	Content    string `json:"content,omitempty" url:"content,omitempty,optional"`       // Allowed storage content types (when type == storage).
	Level      string `json:"level,omitempty" url:"level,omitempty,optional"`           // Support level (when type == node).
	Node       string `json:"node,omitempty" url:"node,omitempty,optional"`             // The cluster node name (when type in node,storage,qemu,lxc).
	Uptime     int    `json:"uptime,omitempty" url:"uptime,omitempty,optional"`         // Node uptime in seconds (when type in node,qemu,lxc).
	Cpu        int    `json:"cpu,omitempty" url:"cpu,omitempty,optional"`               // CPU utilization (when type in node,qemu,lxc).
	Maxcpu     int    `json:"maxcpu,omitempty" url:"maxcpu,omitempty,optional"`         // Number of available CPUs (when type in node,qemu,lxc).
	Plugintype string `json:"plugintype,omitempty" url:"plugintype,omitempty,optional"` // More specific type, if available.
	Pool       string `json:"pool,omitempty" url:"pool,omitempty,optional"`             // The pool name (when type in pool,qemu,lxc).
	Type       string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`   // Resource type.
	Disk       string `json:"disk,omitempty" url:"disk,omitempty,optional"`             // Used disk space in bytes (when type in storage), used root image spave for VMs (type in qemu,lxc).
	Mem        string `json:"mem,omitempty" url:"mem,omitempty,optional"`               // Used memory in bytes (when type in node,qemu,lxc).
	Storage    string `json:"storage,omitempty" url:"storage,omitempty,optional"`       // The storage identifier (when type == storage).
	Status     string `json:"status,omitempty" url:"status,omitempty,optional"`         // Resource type dependent status.
	Hastate    string `json:"hastate,omitempty" url:"hastate,omitempty,optional"`       // HA service status (for HA managed VMs).
	Id         string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`       //
	Maxdisk    int    `json:"maxdisk,omitempty" url:"maxdisk,omitempty,optional"`       // Storage size in bytes (when type in storage), root image size for VMs (type in qemu,lxc).
	Maxmem     int    `json:"maxmem,omitempty" url:"maxmem,omitempty,optional"`         // Number of available memory in bytes (when type in node,qemu,lxc).
	Name       string `json:"name,omitempty" url:"name,omitempty,optional"`             // Name of the resource.
}

// ClusterTasksResponse tasks
// List recent tasks (cluster wide).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/tasks
//
type ClusterTasksResponse struct {
	Upid string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"` //
}

// ClusterOptionsSetRequest options - Set datacenter options.
// Set datacenter options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/options
//
type ClusterOptionsSetRequest struct {
	Console           string `json:"console,omitempty" url:"console,omitempty,optional"`                       // Select the default Console viewer. You can either use the builtin java applet (VNC; deprecated and maps to html5), an external virt-viewer comtatible application (SPICE), an HTML5 based vnc viewer (noVNC), or an HTML5 based console client (xtermjs). If the selected viewer is not available (e.g. SPICE not activated for the VM), the fallback is noVNC.
	EmailFrom         string `json:"email_from,omitempty" url:"email_from,omitempty,optional"`                 // Specify email address to send notification from (default is root@$hostname)
	Migration         string `json:"migration,omitempty" url:"migration,omitempty,optional"`                   // For cluster wide migration settings.
	Webauthn          string `json:"webauthn,omitempty" url:"webauthn,omitempty,optional"`                     // webauthn configuration
	Fencing           string `json:"fencing,omitempty" url:"fencing,omitempty,optional"`                       // Set the fencing mode of the HA cluster. Hardware mode needs a valid configuration of fence devices in /etc/pve/ha/fence.cfg. With both all two modes are used.  WARNING: 'hardware' and 'both' are EXPERIMENTAL & WIP
	Language          string `json:"language,omitempty" url:"language,omitempty,optional"`                     // Default GUI language.
	NextId            string `json:"next-id,omitempty" url:"next-id,omitempty,optional"`                       // Control the range for the free VMID auto-selection pool.
	Delete            string `json:"delete,omitempty" url:"delete,omitempty,optional"`                         // A list of settings you want to delete.
	Description       string `json:"description,omitempty" url:"description,omitempty,optional"`               // Datacenter description. Shown in the web-interface datacenter notes panel. This is saved as comment inside the configuration file.
	Ha                string `json:"ha,omitempty" url:"ha,omitempty,optional"`                                 // Cluster wide HA settings.
	HttpProxy         string `json:"http_proxy,omitempty" url:"http_proxy,omitempty,optional"`                 // Specify external http proxy which is used for downloads (example: 'http://username:password@host:port/')
	Keyboard          string `json:"keyboard,omitempty" url:"keyboard,omitempty,optional"`                     // Default keybord layout for vnc server.
	MigrationUnsecure bool   `json:"migration_unsecure,omitempty" url:"migration_unsecure,omitempty,optional"` // Migration is secure using SSH tunnel by default. For secure private networks you can disable it to speed up migration. Deprecated, use the 'migration' property instead!
	U2f               string `json:"u2f,omitempty" url:"u2f,omitempty,optional"`                               // u2f
	Bwlimit           string `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`                       // Set bandwidth/io limits various operations.
	MacPrefix         string `json:"mac_prefix,omitempty" url:"mac_prefix,omitempty,optional"`                 // Prefix for autogenerated MAC addresses.
	MaxWorkers        int    `json:"max_workers,omitempty" url:"max_workers,omitempty,optional"`               // Defines how many workers (per node) are maximal started  on actions like 'stopall VMs' or task from the ha-manager.
}

// ClusterStatusGetResponse status
// Get cluster status information.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/status
//
type ClusterStatusGetResponse struct {
	Ip      string `json:"ip,omitempty" url:"ip,omitempty,optional"`               // [node] IP of the resolved nodename.
	Level   string `json:"level,omitempty" url:"level,omitempty,optional"`         // [node] Proxmox VE Subscription level, indicates if eligible for enterprise support as well as access to the stable Proxmox VE Enterprise Repository.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` //
	Nodeid  int    `json:"nodeid,omitempty" url:"nodeid,omitempty,optional"`       // [node] ID of the node from the corosync configuration.
	Online  bool   `json:"online,omitempty" url:"online,omitempty,optional"`       // [node] Indicates if the node is online or offline.
	Quorate bool   `json:"quorate,omitempty" url:"quorate,omitempty,optional"`     // [cluster] Indicates if there is a majority of nodes online to make decisions
	Type    string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"` // Indicates the type, either cluster or node. The type defines the object properties e.g. quorate available for type cluster.
	Id      string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`     //
	Local   bool   `json:"local,omitempty" url:"local,omitempty,optional"`         // [node] Indicates if this is the responding node.
	Nodes   int    `json:"nodes,omitempty" url:"nodes,omitempty,optional"`         // [cluster] Nodes count, including offline nodes.
	Version int    `json:"version,omitempty" url:"version,omitempty,optional"`     // [cluster] Current version of the corosync configuration file.
}

// ClusterNextidRequest nextid - Get next free VMID. Pass a VMID to assert that its free (at time of check).
// Get next free VMID. Pass a VMID to assert that its free (at time of check).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/nextid
//
type ClusterNextidRequest struct {
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty,optional"` // The (unique) ID of the VM.
}

// NodesQemuVmlistRequest qemu - Virtual machine index (per node).
// Virtual machine index (per node).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu
//
type NodesQemuVmlistRequest struct {
	Full bool   `json:"full,omitempty" url:"full,omitempty,optional"`           // Determine the full status of active VMs.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesQemuVmlistResponse qemu
// Virtual machine index (per node).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu
//
type NodesQemuVmlistResponse struct {
	Tags           string `json:"tags,omitempty" url:"tags,omitempty,optional"`                       // The current configured tags, if any
	Uptime         int    `json:"uptime,omitempty" url:"uptime,omitempty,optional"`                   // Uptime.
	Cpus           int    `json:"cpus,omitempty" url:"cpus,omitempty,optional"`                       // Maximum usable CPUs.
	Name           string `json:"name,omitempty" url:"name,omitempty,optional"`                       // VM name.
	Pid            int    `json:"pid,omitempty" url:"pid,omitempty,optional"`                         // PID of running qemu process.
	RunningMachine string `json:"running-machine,omitempty" url:"running-machine,omitempty,optional"` // The currently running machine type (if running).
	RunningQemu    string `json:"running-qemu,omitempty" url:"running-qemu,omitempty,optional"`       // The currently running QEMU version (if running).
	Vmid           string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`             // The (unique) ID of the VM.
	Lock           string `json:"lock,omitempty" url:"lock,omitempty,optional"`                       // The current config lock, if any.
	Maxdisk        int    `json:"maxdisk,omitempty" url:"maxdisk,omitempty,optional"`                 // Root disk size in bytes.
	Maxmem         int    `json:"maxmem,omitempty" url:"maxmem,omitempty,optional"`                   // Maximum memory in bytes.
	Qmpstatus      string `json:"qmpstatus,omitempty" url:"qmpstatus,omitempty,optional"`             // Qemu QMP agent status.
	Status         string `json:"status,omitempty" url:"status,omitempty" validate:"nonzero"`         // Qemu process status.
}

// NodesQemuCreateVmRequest qemu - Create or restore a virtual machine.
// Create or restore a virtual machine.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu
//
type NodesQemuCreateVmRequest struct {
	Cpu               string   `json:"cpu,omitempty" url:"cpu,omitempty,optional"`                               // Emulated CPU type.
	Kvm               bool     `json:"kvm,omitempty" url:"kvm,omitempty,optional"`                               // Enable/disable KVM hardware virtualization.
	Name              string   `json:"name,omitempty" url:"name,omitempty,optional"`                             // Set a name for the VM. Only used on the configuration web interface.
	Ostype            string   `json:"ostype,omitempty" url:"ostype,omitempty,optional"`                         // Specify guest operating system.
	Searchdomain      string   `json:"searchdomain,omitempty" url:"searchdomain,omitempty,optional"`             // cloud-init: Sets DNS search domains for a container. Create will' 	    .' automatically use the setting from the host if neither searchdomain nor nameserver' 	    .' are set.
	Startup           string   `json:"startup,omitempty" url:"startup,omitempty,optional"`                       // Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Agent             string   `json:"agent,omitempty" url:"agent,omitempty,optional"`                           // Enable/disable communication with the Qemu Guest Agent and its properties.
	Autostart         bool     `json:"autostart,omitempty" url:"autostart,omitempty,optional"`                   // Automatic restart after crash (currently ignored).
	Ciuser            string   `json:"ciuser,omitempty" url:"ciuser,omitempty,optional"`                         // cloud-init: User name to change ssh keys and password for instead of the image's configured default user.
	Ivshmem           string   `json:"ivshmem,omitempty" url:"ivshmem,omitempty,optional"`                       // Inter-VM shared memory. Useful for direct communication between VMs, or to the host.
	Localtime         bool     `json:"localtime,omitempty" url:"localtime,omitempty,optional"`                   // Set the real time clock (RTC) to local time. This is enabled by default if the `ostype` indicates a Microsoft Windows OS.
	Start             bool     `json:"start,omitempty" url:"start,omitempty,optional"`                           // Start VM after it was created successfully.
	Tags              string   `json:"tags,omitempty" url:"tags,omitempty,optional"`                             // Tags of the VM. This is only meta information.
	Bwlimit           int      `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`                       // Override I/O bandwidth limit (in KiB/s).
	Cicustom          string   `json:"cicustom,omitempty" url:"cicustom,omitempty,optional"`                     // cloud-init: Specify custom files to replace the automatically generated ones at start.
	SpiceEnhancements string   `json:"spice_enhancements,omitempty" url:"spice_enhancements,omitempty,optional"` // Configure additional enhancements for SPICE.
	Acpi              bool     `json:"acpi,omitempty" url:"acpi,omitempty,optional"`                             // Enable/disable ACPI.
	Pool              string   `json:"pool,omitempty" url:"pool,omitempty,optional"`                             // Add the VM to the specified pool.
	LiveRestore       bool     `json:"live-restore,omitempty" url:"live-restore,omitempty,optional"`             // Start the VM immediately from the backup and restore in background. PBS only.
	Protection        bool     `json:"protection,omitempty" url:"protection,omitempty,optional"`                 // Sets the protection flag of the VM. This will disable the remove VM and remove disk operations.
	Sata              []string `json:"sata[n],omitempty" url:"sata[n],omitempty,optional"`                       // Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Scsi              []string `json:"scsi[n],omitempty" url:"scsi[n],omitempty,optional"`                       // Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Cpulimit          int      `json:"cpulimit,omitempty" url:"cpulimit,omitempty,optional"`                     // Limit of CPU usage.
	Force             bool     `json:"force,omitempty" url:"force,omitempty,optional"`                           // Allow to overwrite existing VM.
	Hotplug           string   `json:"hotplug,omitempty" url:"hotplug,omitempty,optional"`                       // Selectively enable hotplug features. This is a comma separated list of hotplug features: 'network', 'disk', 'cpu', 'memory' and 'usb'. Use '0' to disable hotplug completely. Using '1' as value is an alias for the default `network,disk,usb`.
	Tdf               bool     `json:"tdf,omitempty" url:"tdf,omitempty,optional"`                               // Enable/disable time drift fix.
	Usb               []string `json:"usb[n],omitempty" url:"usb[n],omitempty,optional"`                         // Configure an USB device (n is 0 to 4).
	Archive           string   `json:"archive,omitempty" url:"archive,omitempty,optional"`                       // The backup archive. Either the file system path to a .tar or .vma file (use '-' to pipe data from stdin) or a proxmox storage backup volume identifier.
	Bios              string   `json:"bios,omitempty" url:"bios,omitempty,optional"`                             // Select BIOS implementation.
	Virtio            []string `json:"virtio[n],omitempty" url:"virtio[n],omitempty,optional"`                   // Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Rng0              string   `json:"rng0,omitempty" url:"rng0,omitempty,optional"`                             // Configure a VirtIO-based Random Number Generator.
	Smbios1           string   `json:"smbios1,omitempty" url:"smbios1,omitempty,optional"`                       // Specify SMBIOS type 1 fields.
	Hostpci           []string `json:"hostpci[n],omitempty" url:"hostpci[n],omitempty,optional"`                 // Map host PCI devices into guest.
	Memory            int      `json:"memory,omitempty" url:"memory,omitempty,optional"`                         // Amount of RAM for the VM in MB. This is the maximum available memory when you use the balloon device.
	Node              string   `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                   // The cluster node name.
	Smp               int      `json:"smp,omitempty" url:"smp,omitempty,optional"`                               // The number of CPUs. Please use option -sockets instead.
	Sshkeys           string   `json:"sshkeys,omitempty" url:"sshkeys,omitempty,optional"`                       // cloud-init: Setup public SSH keys (one key per line, OpenSSH format).
	Watchdog          string   `json:"watchdog,omitempty" url:"watchdog,omitempty,optional"`                     // Create a virtual hardware watchdog device.
	Balloon           int      `json:"balloon,omitempty" url:"balloon,omitempty,optional"`                       // Amount of target RAM for the VM in MB. Using zero disables the ballon driver.
	Citype            string   `json:"citype,omitempty" url:"citype,omitempty,optional"`                         // Specifies the cloud-init configuration format. The default depends on the configured operating system type (`ostype`. We use the `nocloud` format for Linux, and `configdrive2` for windows.
	Cipassword        string   `json:"cipassword,omitempty" url:"cipassword,omitempty,optional"`                 // cloud-init: Password to assign the user. Using this is generally not recommended. Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords.
	Cpuunits          int      `json:"cpuunits,omitempty" url:"cpuunits,omitempty,optional"`                     // CPU weight for a VM, will be clamped to [1, 10000] in cgroup v2.
	Description       string   `json:"description,omitempty" url:"description,omitempty,optional"`               // Description for the VM. Shown in the web-interface VM's summary. This is saved as comment inside the configuration file.
	Vcpus             int      `json:"vcpus,omitempty" url:"vcpus,omitempty,optional"`                           // Number of hotplugged vcpus.
	Vga               string   `json:"vga,omitempty" url:"vga,omitempty,optional"`                               // Configure the VGA hardware.
	Vmstatestorage    string   `json:"vmstatestorage,omitempty" url:"vmstatestorage,omitempty,optional"`         // Default storage for VM state volumes/files.
	Args              string   `json:"args,omitempty" url:"args,omitempty,optional"`                             // Arbitrary arguments passed to kvm.
	Cdrom             string   `json:"cdrom,omitempty" url:"cdrom,omitempty,optional"`                           // This is an alias for option -ide2
	Template          bool     `json:"template,omitempty" url:"template,omitempty,optional"`                     // Enable/disable Template.
	Vmid              string   `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`                   // The (unique) ID of the VM.
	Arch              string   `json:"arch,omitempty" url:"arch,omitempty,optional"`                             // Virtual processor architecture. Defaults to the host.
	Freeze            bool     `json:"freeze,omitempty" url:"freeze,omitempty,optional"`                         // Freeze CPU at startup (use 'c' monitor command to start execution).
	Nameserver        string   `json:"nameserver,omitempty" url:"nameserver,omitempty,optional"`                 // cloud-init: Sets DNS server IP address for a container. Create will' 	    .' automatically use the setting from the host if neither searchdomain nor nameserver' 	    .' are set.
	Reboot            bool     `json:"reboot,omitempty" url:"reboot,omitempty,optional"`                         // Allow reboot. If set to '0' the VM exit on reboot.
	Sockets           int      `json:"sockets,omitempty" url:"sockets,omitempty,optional"`                       // The number of CPU sockets.
	Startdate         string   `json:"startdate,omitempty" url:"startdate,omitempty,optional"`                   // Set the initial date of the real time clock. Valid format for date are:'now' or '2006-06-17T16:01:21' or '2006-06-17'.
	Boot              string   `json:"boot,omitempty" url:"boot,omitempty,optional"`                             // Specify guest boot order. Use the 'order=' sub-property as usage with no key or 'legacy=' is deprecated.
	Cores             int      `json:"cores,omitempty" url:"cores,omitempty,optional"`                           // The number of cores per socket.
	Unused            []string `json:"unused[n],omitempty" url:"unused[n],omitempty,optional"`                   // Reference to unused volumes. This is used internally, and should not be modified manually.
	Keephugepages     bool     `json:"keephugepages,omitempty" url:"keephugepages,omitempty,optional"`           // Use together with hugepages. If enabled, hugepages will not not be deleted after VM shutdown and can be used for subsequent starts.
	Serial            []string `json:"serial[n],omitempty" url:"serial[n],omitempty,optional"`                   // Create a serial device inside the VM (n is 0 to 3)
	MigrateDowntime   int      `json:"migrate_downtime,omitempty" url:"migrate_downtime,omitempty,optional"`     // Set maximum tolerated downtime (in seconds) for migrations.
	Scsihw            string   `json:"scsihw,omitempty" url:"scsihw,omitempty,optional"`                         // SCSI controller model
	Unique            bool     `json:"unique,omitempty" url:"unique,omitempty,optional"`                         // Assign a unique random ethernet address.
	Ide               []string `json:"ide[n],omitempty" url:"ide[n],omitempty,optional"`                         // Use volume as IDE hard disk or CD-ROM (n is 0 to 3). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Keyboard          string   `json:"keyboard,omitempty" url:"keyboard,omitempty,optional"`                     // Keyboard layout for VNC server. This option is generally not required and is often better handled from within the guest OS.
	Tablet            bool     `json:"tablet,omitempty" url:"tablet,omitempty,optional"`                         // Enable/disable the USB tablet device.
	Tpmstate0         string   `json:"tpmstate0,omitempty" url:"tpmstate0,omitempty,optional"`                   // Configure a Disk for storing TPM state. The format is fixed to 'raw'. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and 4 MiB will be used instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	NumaEnabled       bool     `json:"numa,omitempty" url:"numa,omitempty,optional"`                             // Enable/disable NUMA.
	Shares            int      `json:"shares,omitempty" url:"shares,omitempty,optional"`                         // Amount of memory shares for auto-ballooning. The larger the number is, the more memory this VM gets. Number is relative to weights of all other running VMs. Using zero disables auto-ballooning. Auto-ballooning is done by pvestatd.
	Storage           string   `json:"storage,omitempty" url:"storage,omitempty,optional"`                       // Default storage.
	Hookscript        string   `json:"hookscript,omitempty" url:"hookscript,omitempty,optional"`                 // Script that will be executed during various steps in the vms lifetime.
	Ipconfig          []string `json:"ipconfig[n],omitempty" url:"ipconfig[n],omitempty,optional"`               // cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string 'dhcp' can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string 'auto' can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.
	Numa              []string `json:"numa[n],omitempty" url:"numa[n],omitempty,optional"`                       // NUMA topology.
	Vmgenid           string   `json:"vmgenid,omitempty" url:"vmgenid,omitempty,optional"`                       // Set VM Generation ID. Use '1' to autogenerate on create or update, pass '0' to disable explicitly.
	Bootdisk          string   `json:"bootdisk,omitempty" url:"bootdisk,omitempty,optional"`                     // Enable booting from specified disk. Deprecated: Use 'boot: order=foo;bar' instead.
	Efidisk0          string   `json:"efidisk0,omitempty" url:"efidisk0,omitempty,optional"`                     // Configure a Disk for storing EFI vars. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and that the default EFI vars are copied to the volume instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Lock              string   `json:"lock,omitempty" url:"lock,omitempty,optional"`                             // Lock/unlock the VM.
	Machine           string   `json:"machine,omitempty" url:"machine,omitempty,optional"`                       // Specifies the Qemu machine type.
	MigrateSpeed      int      `json:"migrate_speed,omitempty" url:"migrate_speed,omitempty,optional"`           // Set maximum speed (in MB/s) for migrations. Value 0 is no limit.
	Net               []string `json:"net[n],omitempty" url:"net[n],omitempty,optional"`                         // Specify network devices.
	Onboot            bool     `json:"onboot,omitempty" url:"onboot,omitempty,optional"`                         // Specifies whether a VM will be started during system bootup.
	Parallel          []string `json:"parallel[n],omitempty" url:"parallel[n],omitempty,optional"`               // Map host parallel devices (n is 0 to 2).
	Audio0            string   `json:"audio0,omitempty" url:"audio0,omitempty,optional"`                         // Configure a audio device, useful in combination with QXL/Spice.
	Hugepages         string   `json:"hugepages,omitempty" url:"hugepages,omitempty,optional"`                   // Enable/disable hugepages memory.
}

// NodesQemuVmdiridxRequest {vmid} - Directory index
// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}
//
type NodesQemuVmdiridxRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuVmdiridxResponse {vmid}
// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}
//
type NodesQemuVmdiridxResponse struct {
	Subdir string `json:"subdir,omitempty" url:"subdir,omitempty" validate:"nonzero"` //
}

// NodesQemuDestroyVmRequest {vmid} - Destroy the VM and  all used/owned volumes. Removes any VM specific permissions and firewall rules
// Destroy the VM and  all used/owned volumes. Removes any VM specific permissions and firewall rules
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}
//
type NodesQemuDestroyVmRequest struct {
	DestroyUnreferencedDisks bool   `json:"destroy-unreferenced-disks,omitempty" url:"destroy-unreferenced-disks,omitempty,optional"` // If set, destroy additionally all disks not referenced in the config but with a matching VMID from all enabled storages.
	Node                     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                                   // The cluster node name.
	Purge                    bool   `json:"purge,omitempty" url:"purge,omitempty,optional"`                                           // Remove VMID from configurations, like backup & replication jobs and HA.
	Skiplock                 bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`                                     // Ignore locks - only root is allowed to use this option.
	Vmid                     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`                                   // The (unique) ID of the VM.
}

// NodesQemuFirewallRulesGetRequest rules - List rules.
// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/rules
//
type NodesQemuFirewallRulesGetRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallRulesGetResponse rules
// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/rules
//
type NodesQemuFirewallRulesGetResponse struct {
	Pos int `json:"pos,omitempty" url:"pos,omitempty" validate:"nonzero"` //
}

// NodesQemuFirewallRulesCreateRuleRequest rules - Create new rule.
// Create new rule.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/rules
//
type NodesQemuFirewallRulesCreateRuleRequest struct {
	Dest     string `json:"dest,omitempty" url:"dest,omitempty,optional"`               // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dport    string `json:"dport,omitempty" url:"dport,omitempty,optional"`             // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Pos      int    `json:"pos,omitempty" url:"pos,omitempty,optional"`                 // Update rule at position <pos>.
	Proto    string `json:"proto,omitempty" url:"proto,omitempty,optional"`             // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Sport    string `json:"sport,omitempty" url:"sport,omitempty,optional"`             // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`           // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Iface    string `json:"iface,omitempty" url:"iface,omitempty,optional"`             // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Log      string `json:"log,omitempty" url:"log,omitempty,optional"`                 // Log level for firewall rule.
	Type     string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     // Rule type.
	Action   string `json:"action,omitempty" url:"action,omitempty" validate:"nonzero"` // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Comment  string `json:"comment,omitempty" url:"comment,omitempty,optional"`         // Descriptive comment.
	IcmpType string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`     // Specify icmp-type. Only valid if proto equals 'icmp'.
	Macro    string `json:"macro,omitempty" url:"macro,omitempty,optional"`             // Use predefined standard macro.
	Source   string `json:"source,omitempty" url:"source,omitempty,optional"`           // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`     // The (unique) ID of the VM.
	Enable   int    `json:"enable,omitempty" url:"enable,omitempty,optional"`           // Flag to enable/disable a rule.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
}

// NodesQemuFirewallRulesDeleteRuleRequest {pos} - Delete rule.
// Delete rule.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/rules/{pos}
//
type NodesQemuFirewallRulesDeleteRuleRequest struct {
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Pos    int    `json:"pos,omitempty" url:"pos,omitempty,optional"`             // Update rule at position <pos>.
	Vmid   string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}

// NodesQemuFirewallRulesGetRuleRequest {pos} - Get single rule data.
// Get single rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/rules/{pos}
//
type NodesQemuFirewallRulesGetRuleRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Pos  int    `json:"pos,omitempty" url:"pos,omitempty,optional"`             // Update rule at position <pos>.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallRulesGetRuleResponse {pos}
// Get single rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/rules/{pos}
//
type NodesQemuFirewallRulesGetRuleResponse struct {
	Iface     string `json:"iface,omitempty" url:"iface,omitempty,optional"`             //
	Pos       int    `json:"pos,omitempty" url:"pos,omitempty" validate:"nonzero"`       //
	Comment   string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Dport     string `json:"dport,omitempty" url:"dport,omitempty,optional"`             //
	IcmpType  string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`     //
	Log       string `json:"log,omitempty" url:"log,omitempty,optional"`                 // Log level for firewall rule
	Proto     string `json:"proto,omitempty" url:"proto,omitempty,optional"`             //
	Ipversion int    `json:"ipversion,omitempty" url:"ipversion,omitempty,optional"`     //
	Source    string `json:"source,omitempty" url:"source,omitempty,optional"`           //
	Type      string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     //
	Action    string `json:"action,omitempty" url:"action,omitempty" validate:"nonzero"` //
	Dest      string `json:"dest,omitempty" url:"dest,omitempty,optional"`               //
	Enable    int    `json:"enable,omitempty" url:"enable,omitempty,optional"`           //
	Macro     string `json:"macro,omitempty" url:"macro,omitempty,optional"`             //
	Sport     string `json:"sport,omitempty" url:"sport,omitempty,optional"`             //
}

// NodesQemuFirewallRulesUpdateRuleRequest {pos} - Modify rule data.
// Modify rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/rules/{pos}
//
type NodesQemuFirewallRulesUpdateRuleRequest struct {
	Delete   string `json:"delete,omitempty" url:"delete,omitempty,optional"`       // A list of settings you want to delete.
	IcmpType string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"` // Specify icmp-type. Only valid if proto equals 'icmp'.
	Source   string `json:"source,omitempty" url:"source,omitempty,optional"`       // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Sport    string `json:"sport,omitempty" url:"sport,omitempty,optional"`         // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Type     string `json:"type,omitempty" url:"type,omitempty,optional"`           // Rule type.
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Proto    string `json:"proto,omitempty" url:"proto,omitempty,optional"`         // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Dest     string `json:"dest,omitempty" url:"dest,omitempty,optional"`           // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dport    string `json:"dport,omitempty" url:"dport,omitempty,optional"`         // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Moveto   int    `json:"moveto,omitempty" url:"moveto,omitempty,optional"`       // Move rule to new position <moveto>. Other arguments are ignored.
	Action   string `json:"action,omitempty" url:"action,omitempty,optional"`       // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Comment  string `json:"comment,omitempty" url:"comment,omitempty,optional"`     // Descriptive comment.
	Enable   int    `json:"enable,omitempty" url:"enable,omitempty,optional"`       // Flag to enable/disable a rule.
	Iface    string `json:"iface,omitempty" url:"iface,omitempty,optional"`         // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Log      string `json:"log,omitempty" url:"log,omitempty,optional"`             // Log level for firewall rule.
	Macro    string `json:"macro,omitempty" url:"macro,omitempty,optional"`         // Use predefined standard macro.
	Pos      int    `json:"pos,omitempty" url:"pos,omitempty,optional"`             // Update rule at position <pos>.
}

// NodesQemuFirewallAliasesGetRequest aliases - List aliases
// List aliases
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/aliases
//
type NodesQemuFirewallAliasesGetRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallAliasesGetResponse aliases
// List aliases
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/aliases
//
type NodesQemuFirewallAliasesGetResponse struct {
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"`     //
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`     //
}

// NodesQemuFirewallAliasesCreateAliasRequest aliases - Create IP or Network Alias.
// Create IP or Network Alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/aliases
//
type NodesQemuFirewallAliasesCreateAliasRequest struct {
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallAliasesUpdateAliasRequest {name} - Update IP or Network alias.
// Update IP or Network alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/aliases/{name}
//
type NodesQemuFirewallAliasesUpdateAliasRequest struct {
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Rename  string `json:"rename,omitempty" url:"rename,omitempty,optional"`       // Rename an existing alias.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}

// NodesQemuFirewallAliasesRemoveAliasRequest {name} - Remove IP or Network alias.
// Remove IP or Network alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/aliases/{name}
//
type NodesQemuFirewallAliasesRemoveAliasRequest struct {
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid   string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallAliasesReadAliasRequest {name} - Read alias.
// Read alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/aliases/{name}
//
type NodesQemuFirewallAliasesReadAliasRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallIpsetIndexRequest ipset - List IPSets
// List IPSets
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/ipset
//
type NodesQemuFirewallIpsetIndexRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallIpsetIndexResponse ipset
// List IPSets
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/ipset
//
type NodesQemuFirewallIpsetIndexResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`     // IP set name.
}

// NodesQemuFirewallIpsetCreateRequest ipset - Create new IPSet
// Create new IPSet
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/ipset
//
type NodesQemuFirewallIpsetCreateRequest struct {
	Rename  string `json:"rename,omitempty" url:"rename,omitempty,optional"`       // Rename an existing IPSet. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing IPSet.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesQemuFirewallIpsetDeleteRequest {name} - Delete IPSet
// Delete IPSet
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}
//
type NodesQemuFirewallIpsetDeleteRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallIpsetGetRequest {name} - List IPSet content
// List IPSet content
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}
//
type NodesQemuFirewallIpsetGetRequest struct {
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesQemuFirewallIpsetGetResponse {name}
// List IPSet content
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}
//
type NodesQemuFirewallIpsetGetResponse struct {
	Digest  string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Nomatch bool   `json:"nomatch,omitempty" url:"nomatch,omitempty,optional"`         //
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"`     //
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
}

// NodesQemuFirewallIpsetCreateIpRequest {name} - Add IP or Network to IPSet.
// Add IP or Network to IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}
//
type NodesQemuFirewallIpsetCreateIpRequest struct {
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Nomatch bool   `json:"nomatch,omitempty" url:"nomatch,omitempty,optional"`     //
}

// NodesQemuFirewallIpsetUpdateIpRequest {cidr} - Update IP or Network settings
// Update IP or Network settings
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}
//
type NodesQemuFirewallIpsetUpdateIpRequest struct {
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Nomatch bool   `json:"nomatch,omitempty" url:"nomatch,omitempty,optional"`     //
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
}

// NodesQemuFirewallIpsetRemoveIpRequest {cidr} - Remove IP or Network from IPSet.
// Remove IP or Network from IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}
//
type NodesQemuFirewallIpsetRemoveIpRequest struct {
	Cidr   string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid   string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallIpsetReadIpRequest {cidr} - Read IP or Network settings from IPSet.
// Read IP or Network settings from IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}
//
type NodesQemuFirewallIpsetReadIpRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Cidr string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
}

// NodesQemuFirewallOptionsSetRequest options - Set Firewall options.
// Set Firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/options
//
type NodesQemuFirewallOptionsSetRequest struct {
	Dhcp        bool   `json:"dhcp,omitempty" url:"dhcp,omitempty,optional"`                   // Enable DHCP.
	Radv        bool   `json:"radv,omitempty" url:"radv,omitempty,optional"`                   // Allow sending Router Advertisement.
	Vmid        string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Delete      string `json:"delete,omitempty" url:"delete,omitempty,optional"`               // A list of settings you want to delete.
	LogLevelIn  string `json:"log_level_in,omitempty" url:"log_level_in,omitempty,optional"`   // Log level for incoming traffic.
	PolicyIn    string `json:"policy_in,omitempty" url:"policy_in,omitempty,optional"`         // Input policy.
	PolicyOut   string `json:"policy_out,omitempty" url:"policy_out,omitempty,optional"`       // Output policy.
	Macfilter   bool   `json:"macfilter,omitempty" url:"macfilter,omitempty,optional"`         // Enable/disable MAC address filter.
	Digest      string `json:"digest,omitempty" url:"digest,omitempty,optional"`               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Enable      bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`               // Enable/disable firewall rules.
	Ipfilter    bool   `json:"ipfilter,omitempty" url:"ipfilter,omitempty,optional"`           // Enable default IP filters. This is equivalent to adding an empty ipfilter-net<id> ipset for every interface. Such ipsets implicitly contain sane default restrictions such as restricting IPv6 link local addresses to the one derived from the interface's MAC address. For containers the configured IP addresses will be implicitly added.
	LogLevelOut string `json:"log_level_out,omitempty" url:"log_level_out,omitempty,optional"` // Log level for outgoing traffic.
	Ndp         bool   `json:"ndp,omitempty" url:"ndp,omitempty,optional"`                     // Enable NDP (Neighbor Discovery Protocol).
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
}

// NodesQemuFirewallOptionsGetRequest options - Get VM firewall options.
// Get VM firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/options
//
type NodesQemuFirewallOptionsGetRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallOptionsGetResponse options
// Get VM firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/options
//
type NodesQemuFirewallOptionsGetResponse struct {
	Ipfilter    bool   `json:"ipfilter,omitempty" url:"ipfilter,omitempty,optional"`           // Enable default IP filters. This is equivalent to adding an empty ipfilter-net<id> ipset for every interface. Such ipsets implicitly contain sane default restrictions such as restricting IPv6 link local addresses to the one derived from the interface's MAC address. For containers the configured IP addresses will be implicitly added.
	LogLevelOut string `json:"log_level_out,omitempty" url:"log_level_out,omitempty,optional"` // Log level for outgoing traffic.
	Macfilter   bool   `json:"macfilter,omitempty" url:"macfilter,omitempty,optional"`         // Enable/disable MAC address filter.
	PolicyOut   string `json:"policy_out,omitempty" url:"policy_out,omitempty,optional"`       // Output policy.
	Radv        bool   `json:"radv,omitempty" url:"radv,omitempty,optional"`                   // Allow sending Router Advertisement.
	Dhcp        bool   `json:"dhcp,omitempty" url:"dhcp,omitempty,optional"`                   // Enable DHCP.
	Enable      bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`               // Enable/disable firewall rules.
	LogLevelIn  string `json:"log_level_in,omitempty" url:"log_level_in,omitempty,optional"`   // Log level for incoming traffic.
	Ndp         bool   `json:"ndp,omitempty" url:"ndp,omitempty,optional"`                     // Enable NDP (Neighbor Discovery Protocol).
	PolicyIn    string `json:"policy_in,omitempty" url:"policy_in,omitempty,optional"`         // Input policy.
}

// NodesQemuFirewallLogRequest log - Read firewall log
// Read firewall log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/log
//
type NodesQemuFirewallLogRequest struct {
	Limit int    `json:"limit,omitempty" url:"limit,omitempty,optional"`         //
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Start int    `json:"start,omitempty" url:"start,omitempty,optional"`         //
	Vmid  string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuFirewallLogResponse log
// Read firewall log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/log
//
type NodesQemuFirewallLogResponse struct {
	T string `json:"t,omitempty" url:"t,omitempty" validate:"nonzero"` // Line text
	N int    `json:"n,omitempty" url:"n,omitempty" validate:"nonzero"` // Line number
}

// NodesQemuFirewallRefsRequest refs - Lists possible IPSet/Alias reference which are allowed in source/dest properties.
// Lists possible IPSet/Alias reference which are allowed in source/dest properties.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/refs
//
type NodesQemuFirewallRefsRequest struct {
	Type string `json:"type,omitempty" url:"type,omitempty,optional"`           // Only list references of specified type.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesQemuFirewallRefsResponse refs
// Lists possible IPSet/Alias reference which are allowed in source/dest properties.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/firewall/refs
//
type NodesQemuFirewallRefsResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` //
	Type    string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"` //
}

// NodesQemuAgentRequest agent - Execute Qemu Guest Agent commands.
// Execute Qemu Guest Agent commands.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent
//
type NodesQemuAgentRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`       // The (unique) ID of the VM.
	Command string `json:"command,omitempty" url:"command,omitempty" validate:"nonzero"` // The QGA command.
}

// NodesQemuAgentFsfreezeFreezeRequest fsfreeze-freeze - Execute fsfreeze-freeze.
// Execute fsfreeze-freeze.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/fsfreeze-freeze
//
type NodesQemuAgentFsfreezeFreezeRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentFsfreezeStatusRequest fsfreeze-status - Execute fsfreeze-status.
// Execute fsfreeze-status.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/fsfreeze-status
//
type NodesQemuAgentFsfreezeStatusRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentFsfreezeThawRequest fsfreeze-thaw - Execute fsfreeze-thaw.
// Execute fsfreeze-thaw.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/fsfreeze-thaw
//
type NodesQemuAgentFsfreezeThawRequest struct {
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesQemuAgentFstrimRequest fstrim - Execute fstrim.
// Execute fstrim.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/fstrim
//
type NodesQemuAgentFstrimRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentGetFsinfoRequest get-fsinfo - Execute get-fsinfo.
// Execute get-fsinfo.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/get-fsinfo
//
type NodesQemuAgentGetFsinfoRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentGetHostNameRequest get-host-name - Execute get-host-name.
// Execute get-host-name.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/get-host-name
//
type NodesQemuAgentGetHostNameRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentGetMemoryBlockInfoRequest get-memory-block-info - Execute get-memory-block-info.
// Execute get-memory-block-info.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/get-memory-block-info
//
type NodesQemuAgentGetMemoryBlockInfoRequest struct {
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesQemuAgentGetMemoryBlocksRequest get-memory-blocks - Execute get-memory-blocks.
// Execute get-memory-blocks.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/get-memory-blocks
//
type NodesQemuAgentGetMemoryBlocksRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentGetOsinfoRequest get-osinfo - Execute get-osinfo.
// Execute get-osinfo.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/get-osinfo
//
type NodesQemuAgentGetOsinfoRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentGetTimeRequest get-time - Execute get-time.
// Execute get-time.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/get-time
//
type NodesQemuAgentGetTimeRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentGetTimezoneRequest get-timezone - Execute get-timezone.
// Execute get-timezone.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/get-timezone
//
type NodesQemuAgentGetTimezoneRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentGetUsersRequest get-users - Execute get-users.
// Execute get-users.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/get-users
//
type NodesQemuAgentGetUsersRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentGetVcpusRequest get-vcpus - Execute get-vcpus.
// Execute get-vcpus.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/get-vcpus
//
type NodesQemuAgentGetVcpusRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentInfoRequest info - Execute info.
// Execute info.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/info
//
type NodesQemuAgentInfoRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentNetworkGetInterfacesRequest network-get-interfaces - Execute network-get-interfaces.
// Execute network-get-interfaces.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/network-get-interfaces
//
type NodesQemuAgentNetworkGetInterfacesRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentPingRequest ping - Execute ping.
// Execute ping.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/ping
//
type NodesQemuAgentPingRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentShutdownRequest shutdown - Execute shutdown.
// Execute shutdown.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/shutdown
//
type NodesQemuAgentShutdownRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentSuspendDiskRequest suspend-disk - Execute suspend-disk.
// Execute suspend-disk.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/suspend-disk
//
type NodesQemuAgentSuspendDiskRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentSuspendHybridRequest suspend-hybrid - Execute suspend-hybrid.
// Execute suspend-hybrid.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/suspend-hybrid
//
type NodesQemuAgentSuspendHybridRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentSuspendRamRequest suspend-ram - Execute suspend-ram.
// Execute suspend-ram.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/suspend-ram
//
type NodesQemuAgentSuspendRamRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentSetUserPasswordRequest set-user-password - Sets the password for the given user to the given password
// Sets the password for the given user to the given password
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/set-user-password
//
type NodesQemuAgentSetUserPasswordRequest struct {
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Password string `json:"password,omitempty" url:"password,omitempty" validate:"nonzero"` // The new password.
	Username string `json:"username,omitempty" url:"username,omitempty" validate:"nonzero"` // The user to set the password for.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Crypted  bool   `json:"crypted,omitempty" url:"crypted,omitempty,optional"`             // set to 1 if the password has already been passed through crypt()
}

// NodesQemuAgentExecRequest exec - Executes the given command in the vm via the guest-agent and returns an object with the pid.
// Executes the given command in the vm via the guest-agent and returns an object with the pid.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/exec
//
type NodesQemuAgentExecRequest struct {
	Command   string `json:"command,omitempty" url:"command,omitempty,optional"`       // The command as a list of program + arguments
	InputData string `json:"input-data,omitempty" url:"input-data,omitempty,optional"` // Data to pass as 'input-data' to the guest. Usually treated as STDIN to 'command'.
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
	Vmid      string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`   // The (unique) ID of the VM.
}

// NodesQemuAgentExecResponse exec
// Executes the given command in the vm via the guest-agent and returns an object with the pid.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/exec
//
type NodesQemuAgentExecResponse struct {
	Pid int `json:"pid,omitempty" url:"pid,omitempty" validate:"nonzero"` // The PID of the process started by the guest-agent.
}

// NodesQemuAgentExecStatusRequest exec-status - Gets the status of the given pid started by the guest-agent
// Gets the status of the given pid started by the guest-agent
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/exec-status
//
type NodesQemuAgentExecStatusRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Pid  int    `json:"pid,omitempty" url:"pid,omitempty" validate:"nonzero"`   // The PID to query
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentExecStatusResponse exec-status
// Gets the status of the given pid started by the guest-agent
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/exec-status
//
type NodesQemuAgentExecStatusResponse struct {
	Signal       int    `json:"signal,omitempty" url:"signal,omitempty,optional"`               // signal number or exception code if the process was abnormally terminated.
	ErrData      string `json:"err-data,omitempty" url:"err-data,omitempty,optional"`           // stderr of the process
	ErrTruncated bool   `json:"err-truncated,omitempty" url:"err-truncated,omitempty,optional"` // true if stderr was not fully captured
	Exitcode     int    `json:"exitcode,omitempty" url:"exitcode,omitempty,optional"`           // process exit code if it was normally terminated.
	Exited       bool   `json:"exited,omitempty" url:"exited,omitempty" validate:"nonzero"`     // Tells if the given command has exited yet.
	OutData      string `json:"out-data,omitempty" url:"out-data,omitempty,optional"`           // stdout of the process
	OutTruncated bool   `json:"out-truncated,omitempty" url:"out-truncated,omitempty,optional"` // true if stdout was not fully captured
}

// NodesQemuAgentFileReadRequest file-read - Reads the given file via guest agent. Is limited to 16777216 bytes.
// Reads the given file via guest agent. Is limited to 16777216 bytes.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/file-read
//
type NodesQemuAgentFileReadRequest struct {
	File string `json:"file,omitempty" url:"file,omitempty" validate:"nonzero"` // The path to the file
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuAgentFileReadResponse file-read
// Reads the given file via guest agent. Is limited to 16777216 bytes.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/file-read
//
type NodesQemuAgentFileReadResponse struct {
	Content   string `json:"content,omitempty" url:"content,omitempty" validate:"nonzero"` // The content of the file, maximum 16777216
	Truncated bool   `json:"truncated,omitempty" url:"truncated,omitempty,optional"`       // If set to 1, the output is truncated and not complete
}

// NodesQemuAgentFileWriteRequest file-write - Writes the given file via guest agent.
// Writes the given file via guest agent.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/agent/file-write
//
type NodesQemuAgentFileWriteRequest struct {
	Content string `json:"content,omitempty" url:"content,omitempty" validate:"nonzero"` // The content to write into the file.
	Encode  bool   `json:"encode,omitempty" url:"encode,omitempty,optional"`             // If set, the content will be encoded as base64 (required by QEMU).Otherwise the content needs to be encoded beforehand - defaults to true.
	File    string `json:"file,omitempty" url:"file,omitempty" validate:"nonzero"`       // The path to the file.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`       // The (unique) ID of the VM.
}

// NodesQemuRrdRequest rrd - Read VM RRD statistics (returns PNG)
// Read VM RRD statistics (returns PNG)
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/rrd
//
type NodesQemuRrdRequest struct {
	Cf        string `json:"cf,omitempty" url:"cf,omitempty,optional"`                         // The RRD consolidation function
	Ds        string `json:"ds,omitempty" url:"ds,omitempty" validate:"nonzero"`               // The list of datasources you want to display.
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Timeframe string `json:"timeframe,omitempty" url:"timeframe,omitempty" validate:"nonzero"` // Specify the time frame you are interested in.
	Vmid      string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`           // The (unique) ID of the VM.
}

// NodesQemuRrdResponse rrd
// Read VM RRD statistics (returns PNG)
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/rrd
//
type NodesQemuRrdResponse struct {
	Filename string `json:"filename,omitempty" url:"filename,omitempty" validate:"nonzero"` //
}

// NodesQemuRrddataRequest rrddata - Read VM RRD statistics
// Read VM RRD statistics
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/rrddata
//
type NodesQemuRrddataRequest struct {
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Timeframe string `json:"timeframe,omitempty" url:"timeframe,omitempty" validate:"nonzero"` // Specify the time frame you are interested in.
	Vmid      string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`           // The (unique) ID of the VM.
	Cf        string `json:"cf,omitempty" url:"cf,omitempty,optional"`                         // The RRD consolidation function
}

// NodesQemuConfigVmRequest config - Get the virtual machine configuration with pending configuration changes applied. Set the 'current' parameter to get the current configuration instead.
// Get the virtual machine configuration with pending configuration changes applied. Set the 'current' parameter to get the current configuration instead.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/config
//
type NodesQemuConfigVmRequest struct {
	Snapshot string `json:"snapshot,omitempty" url:"snapshot,omitempty,optional"`   // Fetch config values from given snapshot.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Current  bool   `json:"current,omitempty" url:"current,omitempty,optional"`     // Get current values (instead of pending values).
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesQemuConfigVmResponse config
// Get the virtual machine configuration with pending configuration changes applied. Set the 'current' parameter to get the current configuration instead.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/config
//
type NodesQemuConfigVmResponse struct {
	Ipconfig          []string `json:"ipconfig[n],omitempty" url:"ipconfig[n],omitempty,optional"`               // cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string 'dhcp' can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string 'auto' can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.
	Tags              string   `json:"tags,omitempty" url:"tags,omitempty,optional"`                             // Tags of the VM. This is only meta information.
	Vmstatestorage    string   `json:"vmstatestorage,omitempty" url:"vmstatestorage,omitempty,optional"`         // Default storage for VM state volumes/files.
	Audio0            string   `json:"audio0,omitempty" url:"audio0,omitempty,optional"`                         // Configure a audio device, useful in combination with QXL/Spice.
	MigrateDowntime   int      `json:"migrate_downtime,omitempty" url:"migrate_downtime,omitempty,optional"`     // Set maximum tolerated downtime (in seconds) for migrations.
	Name              string   `json:"name,omitempty" url:"name,omitempty,optional"`                             // Set a name for the VM. Only used on the configuration web interface.
	Reboot            bool     `json:"reboot,omitempty" url:"reboot,omitempty,optional"`                         // Allow reboot. If set to '0' the VM exit on reboot.
	Virtio            []string `json:"virtio[n],omitempty" url:"virtio[n],omitempty,optional"`                   // Use volume as VIRTIO hard disk (n is 0 to 15).
	Agent             string   `json:"agent,omitempty" url:"agent,omitempty,optional"`                           // Enable/disable communication with the Qemu Guest Agent and its properties.
	Efidisk0          string   `json:"efidisk0,omitempty" url:"efidisk0,omitempty,optional"`                     // Configure a Disk for storing EFI vars.
	Scsihw            string   `json:"scsihw,omitempty" url:"scsihw,omitempty,optional"`                         // SCSI controller model
	SpiceEnhancements string   `json:"spice_enhancements,omitempty" url:"spice_enhancements,omitempty,optional"` // Configure additional enhancements for SPICE.
	Tablet            bool     `json:"tablet,omitempty" url:"tablet,omitempty,optional"`                         // Enable/disable the USB tablet device.
	Tdf               bool     `json:"tdf,omitempty" url:"tdf,omitempty,optional"`                               // Enable/disable time drift fix.
	Usb               []string `json:"usb[n],omitempty" url:"usb[n],omitempty,optional"`                         // Configure an USB device (n is 0 to 4).
	Cpuunits          int      `json:"cpuunits,omitempty" url:"cpuunits,omitempty,optional"`                     // CPU weight for a VM, will be clamped to [1, 10000] in cgroup v2.
	Numa              []string `json:"numa[n],omitempty" url:"numa[n],omitempty,optional"`                       // NUMA topology.
	Onboot            bool     `json:"onboot,omitempty" url:"onboot,omitempty,optional"`                         // Specifies whether a VM will be started during system bootup.
	Protection        bool     `json:"protection,omitempty" url:"protection,omitempty,optional"`                 // Sets the protection flag of the VM. This will disable the remove VM and remove disk operations.
	Sockets           int      `json:"sockets,omitempty" url:"sockets,omitempty,optional"`                       // The number of CPU sockets.
	Unused            []string `json:"unused[n],omitempty" url:"unused[n],omitempty,optional"`                   // Reference to unused volumes. This is used internally, and should not be modified manually.
	Acpi              bool     `json:"acpi,omitempty" url:"acpi,omitempty,optional"`                             // Enable/disable ACPI.
	Arch              string   `json:"arch,omitempty" url:"arch,omitempty,optional"`                             // Virtual processor architecture. Defaults to the host.
	Cpu               string   `json:"cpu,omitempty" url:"cpu,omitempty,optional"`                               // Emulated CPU type.
	Digest            string   `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"`               // SHA1 digest of configuration file. This can be used to prevent concurrent modifications.
	Lock              string   `json:"lock,omitempty" url:"lock,omitempty,optional"`                             // Lock/unlock the VM.
	Bootdisk          string   `json:"bootdisk,omitempty" url:"bootdisk,omitempty,optional"`                     // Enable booting from specified disk. Deprecated: Use 'boot: order=foo;bar' instead.
	Net               []string `json:"net[n],omitempty" url:"net[n],omitempty,optional"`                         // Specify network devices.
	Smp               int      `json:"smp,omitempty" url:"smp,omitempty,optional"`                               // The number of CPUs. Please use option -sockets instead.
	Ciuser            string   `json:"ciuser,omitempty" url:"ciuser,omitempty,optional"`                         // cloud-init: User name to change ssh keys and password for instead of the image's configured default user.
	Cores             int      `json:"cores,omitempty" url:"cores,omitempty,optional"`                           // The number of cores per socket.
	Sata              []string `json:"sata[n],omitempty" url:"sata[n],omitempty,optional"`                       // Use volume as SATA hard disk or CD-ROM (n is 0 to 5).
	Tpmstate0         string   `json:"tpmstate0,omitempty" url:"tpmstate0,omitempty,optional"`                   // Configure a Disk for storing TPM state. The format is fixed to 'raw'.
	Watchdog          string   `json:"watchdog,omitempty" url:"watchdog,omitempty,optional"`                     // Create a virtual hardware watchdog device.
	Cdrom             string   `json:"cdrom,omitempty" url:"cdrom,omitempty,optional"`                           // This is an alias for option -ide2
	Description       string   `json:"description,omitempty" url:"description,omitempty,optional"`               // Description for the VM. Shown in the web-interface VM's summary. This is saved as comment inside the configuration file.
	Hotplug           string   `json:"hotplug,omitempty" url:"hotplug,omitempty,optional"`                       // Selectively enable hotplug features. This is a comma separated list of hotplug features: 'network', 'disk', 'cpu', 'memory' and 'usb'. Use '0' to disable hotplug completely. Using '1' as value is an alias for the default `network,disk,usb`.
	NumaEnabled       bool     `json:"numa,omitempty" url:"numa,omitempty,optional"`                             // Enable/disable NUMA.
	Ostype            string   `json:"ostype,omitempty" url:"ostype,omitempty,optional"`                         // Specify guest operating system.
	Shares            int      `json:"shares,omitempty" url:"shares,omitempty,optional"`                         // Amount of memory shares for auto-ballooning. The larger the number is, the more memory this VM gets. Number is relative to weights of all other running VMs. Using zero disables auto-ballooning. Auto-ballooning is done by pvestatd.
	Kvm               bool     `json:"kvm,omitempty" url:"kvm,omitempty,optional"`                               // Enable/disable KVM hardware virtualization.
	Bios              string   `json:"bios,omitempty" url:"bios,omitempty,optional"`                             // Select BIOS implementation.
	Cipassword        string   `json:"cipassword,omitempty" url:"cipassword,omitempty,optional"`                 // cloud-init: Password to assign the user. Using this is generally not recommended. Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords.
	Citype            string   `json:"citype,omitempty" url:"citype,omitempty,optional"`                         // Specifies the cloud-init configuration format. The default depends on the configured operating system type (`ostype`. We use the `nocloud` format for Linux, and `configdrive2` for windows.
	Machine           string   `json:"machine,omitempty" url:"machine,omitempty,optional"`                       // Specifies the Qemu machine type.
	Scsi              []string `json:"scsi[n],omitempty" url:"scsi[n],omitempty,optional"`                       // Use volume as SCSI hard disk or CD-ROM (n is 0 to 30).
	Parallel          []string `json:"parallel[n],omitempty" url:"parallel[n],omitempty,optional"`               // Map host parallel devices (n is 0 to 2).
	Startdate         string   `json:"startdate,omitempty" url:"startdate,omitempty,optional"`                   // Set the initial date of the real time clock. Valid format for date are:'now' or '2006-06-17T16:01:21' or '2006-06-17'.
	Autostart         bool     `json:"autostart,omitempty" url:"autostart,omitempty,optional"`                   // Automatic restart after crash (currently ignored).
	Balloon           int      `json:"balloon,omitempty" url:"balloon,omitempty,optional"`                       // Amount of target RAM for the VM in MB. Using zero disables the ballon driver.
	Boot              string   `json:"boot,omitempty" url:"boot,omitempty,optional"`                             // Specify guest boot order. Use the 'order=' sub-property as usage with no key or 'legacy=' is deprecated.
	Hugepages         string   `json:"hugepages,omitempty" url:"hugepages,omitempty,optional"`                   // Enable/disable hugepages memory.
	Ide               []string `json:"ide[n],omitempty" url:"ide[n],omitempty,optional"`                         // Use volume as IDE hard disk or CD-ROM (n is 0 to 3).
	Keyboard          string   `json:"keyboard,omitempty" url:"keyboard,omitempty,optional"`                     // Keyboard layout for VNC server. This option is generally not required and is often better handled from within the guest OS.
	Template          bool     `json:"template,omitempty" url:"template,omitempty,optional"`                     // Enable/disable Template.
	Cpulimit          int      `json:"cpulimit,omitempty" url:"cpulimit,omitempty,optional"`                     // Limit of CPU usage.
	Memory            int      `json:"memory,omitempty" url:"memory,omitempty,optional"`                         // Amount of RAM for the VM in MB. This is the maximum available memory when you use the balloon device.
	Vcpus             int      `json:"vcpus,omitempty" url:"vcpus,omitempty,optional"`                           // Number of hotplugged vcpus.
	Vga               string   `json:"vga,omitempty" url:"vga,omitempty,optional"`                               // Configure the VGA hardware.
	Vmgenid           string   `json:"vmgenid,omitempty" url:"vmgenid,omitempty,optional"`                       // Set VM Generation ID. Use '1' to autogenerate on create or update, pass '0' to disable explicitly.
	Hookscript        string   `json:"hookscript,omitempty" url:"hookscript,omitempty,optional"`                 // Script that will be executed during various steps in the vms lifetime.
	Startup           string   `json:"startup,omitempty" url:"startup,omitempty,optional"`                       // Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Nameserver        string   `json:"nameserver,omitempty" url:"nameserver,omitempty,optional"`                 // cloud-init: Sets DNS server IP address for a container. Create will' 	    .' automatically use the setting from the host if neither searchdomain nor nameserver' 	    .' are set.
	Searchdomain      string   `json:"searchdomain,omitempty" url:"searchdomain,omitempty,optional"`             // cloud-init: Sets DNS search domains for a container. Create will' 	    .' automatically use the setting from the host if neither searchdomain nor nameserver' 	    .' are set.
	Args              string   `json:"args,omitempty" url:"args,omitempty,optional"`                             // Arbitrary arguments passed to kvm.
	Cicustom          string   `json:"cicustom,omitempty" url:"cicustom,omitempty,optional"`                     // cloud-init: Specify custom files to replace the automatically generated ones at start.
	Freeze            bool     `json:"freeze,omitempty" url:"freeze,omitempty,optional"`                         // Freeze CPU at startup (use 'c' monitor command to start execution).
	Hostpci           []string `json:"hostpci[n],omitempty" url:"hostpci[n],omitempty,optional"`                 // Map host PCI devices into guest.
	Ivshmem           string   `json:"ivshmem,omitempty" url:"ivshmem,omitempty,optional"`                       // Inter-VM shared memory. Useful for direct communication between VMs, or to the host.
	Localtime         bool     `json:"localtime,omitempty" url:"localtime,omitempty,optional"`                   // Set the real time clock (RTC) to local time. This is enabled by default if the `ostype` indicates a Microsoft Windows OS.
	Keephugepages     bool     `json:"keephugepages,omitempty" url:"keephugepages,omitempty,optional"`           // Use together with hugepages. If enabled, hugepages will not not be deleted after VM shutdown and can be used for subsequent starts.
	Rng0              string   `json:"rng0,omitempty" url:"rng0,omitempty,optional"`                             // Configure a VirtIO-based Random Number Generator.
	Serial            []string `json:"serial[n],omitempty" url:"serial[n],omitempty,optional"`                   // Create a serial device inside the VM (n is 0 to 3)
	Smbios1           string   `json:"smbios1,omitempty" url:"smbios1,omitempty,optional"`                       // Specify SMBIOS type 1 fields.
	MigrateSpeed      int      `json:"migrate_speed,omitempty" url:"migrate_speed,omitempty,optional"`           // Set maximum speed (in MB/s) for migrations. Value 0 is no limit.
	Sshkeys           string   `json:"sshkeys,omitempty" url:"sshkeys,omitempty,optional"`                       // cloud-init: Setup public SSH keys (one key per line, OpenSSH format).
}

// NodesQemuConfigUpdateVmAsyncRequest config - Set virtual machine options (asynchrounous API).
// Set virtual machine options (asynchrounous API).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/config
//
type NodesQemuConfigUpdateVmAsyncRequest struct {
	Ostype            string   `json:"ostype,omitempty" url:"ostype,omitempty,optional"`                         // Specify guest operating system.
	Reboot            bool     `json:"reboot,omitempty" url:"reboot,omitempty,optional"`                         // Allow reboot. If set to '0' the VM exit on reboot.
	Sshkeys           string   `json:"sshkeys,omitempty" url:"sshkeys,omitempty,optional"`                       // cloud-init: Setup public SSH keys (one key per line, OpenSSH format).
	Freeze            bool     `json:"freeze,omitempty" url:"freeze,omitempty,optional"`                         // Freeze CPU at startup (use 'c' monitor command to start execution).
	Net               []string `json:"net[n],omitempty" url:"net[n],omitempty,optional"`                         // Specify network devices.
	Unused            []string `json:"unused[n],omitempty" url:"unused[n],omitempty,optional"`                   // Reference to unused volumes. This is used internally, and should not be modified manually.
	Acpi              bool     `json:"acpi,omitempty" url:"acpi,omitempty,optional"`                             // Enable/disable ACPI.
	Bios              string   `json:"bios,omitempty" url:"bios,omitempty,optional"`                             // Select BIOS implementation.
	Tags              string   `json:"tags,omitempty" url:"tags,omitempty,optional"`                             // Tags of the VM. This is only meta information.
	Arch              string   `json:"arch,omitempty" url:"arch,omitempty,optional"`                             // Virtual processor architecture. Defaults to the host.
	Hookscript        string   `json:"hookscript,omitempty" url:"hookscript,omitempty,optional"`                 // Script that will be executed during various steps in the vms lifetime.
	Cpuunits          int      `json:"cpuunits,omitempty" url:"cpuunits,omitempty,optional"`                     // CPU weight for a VM, will be clamped to [1, 10000] in cgroup v2.
	Hostpci           []string `json:"hostpci[n],omitempty" url:"hostpci[n],omitempty,optional"`                 // Map host PCI devices into guest.
	Hugepages         string   `json:"hugepages,omitempty" url:"hugepages,omitempty,optional"`                   // Enable/disable hugepages memory.
	Keephugepages     bool     `json:"keephugepages,omitempty" url:"keephugepages,omitempty,optional"`           // Use together with hugepages. If enabled, hugepages will not not be deleted after VM shutdown and can be used for subsequent starts.
	MigrateDowntime   int      `json:"migrate_downtime,omitempty" url:"migrate_downtime,omitempty,optional"`     // Set maximum tolerated downtime (in seconds) for migrations.
	Revert            string   `json:"revert,omitempty" url:"revert,omitempty,optional"`                         // Revert a pending change.
	Bootdisk          string   `json:"bootdisk,omitempty" url:"bootdisk,omitempty,optional"`                     // Enable booting from specified disk. Deprecated: Use 'boot: order=foo;bar' instead.
	Ciuser            string   `json:"ciuser,omitempty" url:"ciuser,omitempty,optional"`                         // cloud-init: User name to change ssh keys and password for instead of the image's configured default user.
	Searchdomain      string   `json:"searchdomain,omitempty" url:"searchdomain,omitempty,optional"`             // cloud-init: Sets DNS search domains for a container. Create will' 	    .' automatically use the setting from the host if neither searchdomain nor nameserver' 	    .' are set.
	Virtio            []string `json:"virtio[n],omitempty" url:"virtio[n],omitempty,optional"`                   // Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Vmstatestorage    string   `json:"vmstatestorage,omitempty" url:"vmstatestorage,omitempty,optional"`         // Default storage for VM state volumes/files.
	Watchdog          string   `json:"watchdog,omitempty" url:"watchdog,omitempty,optional"`                     // Create a virtual hardware watchdog device.
	Cpu               string   `json:"cpu,omitempty" url:"cpu,omitempty,optional"`                               // Emulated CPU type.
	Lock              string   `json:"lock,omitempty" url:"lock,omitempty,optional"`                             // Lock/unlock the VM.
	Cpulimit          int      `json:"cpulimit,omitempty" url:"cpulimit,omitempty,optional"`                     // Limit of CPU usage.
	Ivshmem           string   `json:"ivshmem,omitempty" url:"ivshmem,omitempty,optional"`                       // Inter-VM shared memory. Useful for direct communication between VMs, or to the host.
	Onboot            bool     `json:"onboot,omitempty" url:"onboot,omitempty,optional"`                         // Specifies whether a VM will be started during system bootup.
	Shares            int      `json:"shares,omitempty" url:"shares,omitempty,optional"`                         // Amount of memory shares for auto-ballooning. The larger the number is, the more memory this VM gets. Number is relative to weights of all other running VMs. Using zero disables auto-ballooning. Auto-ballooning is done by pvestatd.
	Startdate         string   `json:"startdate,omitempty" url:"startdate,omitempty,optional"`                   // Set the initial date of the real time clock. Valid format for date are:'now' or '2006-06-17T16:01:21' or '2006-06-17'.
	Startup           string   `json:"startup,omitempty" url:"startup,omitempty,optional"`                       // Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Balloon           int      `json:"balloon,omitempty" url:"balloon,omitempty,optional"`                       // Amount of target RAM for the VM in MB. Using zero disables the ballon driver.
	Cores             int      `json:"cores,omitempty" url:"cores,omitempty,optional"`                           // The number of cores per socket.
	Vmgenid           string   `json:"vmgenid,omitempty" url:"vmgenid,omitempty,optional"`                       // Set VM Generation ID. Use '1' to autogenerate on create or update, pass '0' to disable explicitly.
	Name              string   `json:"name,omitempty" url:"name,omitempty,optional"`                             // Set a name for the VM. Only used on the configuration web interface.
	Tablet            bool     `json:"tablet,omitempty" url:"tablet,omitempty,optional"`                         // Enable/disable the USB tablet device.
	Usb               []string `json:"usb[n],omitempty" url:"usb[n],omitempty,optional"`                         // Configure an USB device (n is 0 to 4).
	Vcpus             int      `json:"vcpus,omitempty" url:"vcpus,omitempty,optional"`                           // Number of hotplugged vcpus.
	Vga               string   `json:"vga,omitempty" url:"vga,omitempty,optional"`                               // Configure the VGA hardware.
	Description       string   `json:"description,omitempty" url:"description,omitempty,optional"`               // Description for the VM. Shown in the web-interface VM's summary. This is saved as comment inside the configuration file.
	Machine           string   `json:"machine,omitempty" url:"machine,omitempty,optional"`                       // Specifies the Qemu machine type.
	Digest            string   `json:"digest,omitempty" url:"digest,omitempty,optional"`                         // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Kvm               bool     `json:"kvm,omitempty" url:"kvm,omitempty,optional"`                               // Enable/disable KVM hardware virtualization.
	Parallel          []string `json:"parallel[n],omitempty" url:"parallel[n],omitempty,optional"`               // Map host parallel devices (n is 0 to 2).
	Scsihw            string   `json:"scsihw,omitempty" url:"scsihw,omitempty,optional"`                         // SCSI controller model
	Skiplock          bool     `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`                     // Ignore locks - only root is allowed to use this option.
	Tdf               bool     `json:"tdf,omitempty" url:"tdf,omitempty,optional"`                               // Enable/disable time drift fix.
	Autostart         bool     `json:"autostart,omitempty" url:"autostart,omitempty,optional"`                   // Automatic restart after crash (currently ignored).
	Boot              string   `json:"boot,omitempty" url:"boot,omitempty,optional"`                             // Specify guest boot order. Use the 'order=' sub-property as usage with no key or 'legacy=' is deprecated.
	Keyboard          string   `json:"keyboard,omitempty" url:"keyboard,omitempty,optional"`                     // Keyboard layout for VNC server. This option is generally not required and is often better handled from within the guest OS.
	NumaEnabled       bool     `json:"numa,omitempty" url:"numa,omitempty,optional"`                             // Enable/disable NUMA.
	Numa              []string `json:"numa[n],omitempty" url:"numa[n],omitempty,optional"`                       // NUMA topology.
	Args              string   `json:"args,omitempty" url:"args,omitempty,optional"`                             // Arbitrary arguments passed to kvm.
	BackgroundDelay   int      `json:"background_delay,omitempty" url:"background_delay,omitempty,optional"`     // Time to wait for the task to finish. We return 'null' if the task finish within that time.
	Node              string   `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                   // The cluster node name.
	Serial            []string `json:"serial[n],omitempty" url:"serial[n],omitempty,optional"`                   // Create a serial device inside the VM (n is 0 to 3)
	Protection        bool     `json:"protection,omitempty" url:"protection,omitempty,optional"`                 // Sets the protection flag of the VM. This will disable the remove VM and remove disk operations.
	SpiceEnhancements string   `json:"spice_enhancements,omitempty" url:"spice_enhancements,omitempty,optional"` // Configure additional enhancements for SPICE.
	Audio0            string   `json:"audio0,omitempty" url:"audio0,omitempty,optional"`                         // Configure a audio device, useful in combination with QXL/Spice.
	Efidisk0          string   `json:"efidisk0,omitempty" url:"efidisk0,omitempty,optional"`                     // Configure a Disk for storing EFI vars. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and that the default EFI vars are copied to the volume instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Ide               []string `json:"ide[n],omitempty" url:"ide[n],omitempty,optional"`                         // Use volume as IDE hard disk or CD-ROM (n is 0 to 3). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Memory            int      `json:"memory,omitempty" url:"memory,omitempty,optional"`                         // Amount of RAM for the VM in MB. This is the maximum available memory when you use the balloon device.
	Smp               int      `json:"smp,omitempty" url:"smp,omitempty,optional"`                               // The number of CPUs. Please use option -sockets instead.
	Delete            string   `json:"delete,omitempty" url:"delete,omitempty,optional"`                         // A list of settings you want to delete.
	Force             bool     `json:"force,omitempty" url:"force,omitempty,optional"`                           // Force physical removal. Without this, we simple remove the disk from the config file and create an additional configuration entry called 'unused[n]', which contains the volume ID. Unlink of unused[n] always cause physical removal.
	MigrateSpeed      int      `json:"migrate_speed,omitempty" url:"migrate_speed,omitempty,optional"`           // Set maximum speed (in MB/s) for migrations. Value 0 is no limit.
	Template          bool     `json:"template,omitempty" url:"template,omitempty,optional"`                     // Enable/disable Template.
	Tpmstate0         string   `json:"tpmstate0,omitempty" url:"tpmstate0,omitempty,optional"`                   // Configure a Disk for storing TPM state. The format is fixed to 'raw'. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and 4 MiB will be used instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Agent             string   `json:"agent,omitempty" url:"agent,omitempty,optional"`                           // Enable/disable communication with the Qemu Guest Agent and its properties.
	Cicustom          string   `json:"cicustom,omitempty" url:"cicustom,omitempty,optional"`                     // cloud-init: Specify custom files to replace the automatically generated ones at start.
	Ipconfig          []string `json:"ipconfig[n],omitempty" url:"ipconfig[n],omitempty,optional"`               // cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string 'dhcp' can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string 'auto' can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.
	Scsi              []string `json:"scsi[n],omitempty" url:"scsi[n],omitempty,optional"`                       // Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Citype            string   `json:"citype,omitempty" url:"citype,omitempty,optional"`                         // Specifies the cloud-init configuration format. The default depends on the configured operating system type (`ostype`. We use the `nocloud` format for Linux, and `configdrive2` for windows.
	Localtime         bool     `json:"localtime,omitempty" url:"localtime,omitempty,optional"`                   // Set the real time clock (RTC) to local time. This is enabled by default if the `ostype` indicates a Microsoft Windows OS.
	Nameserver        string   `json:"nameserver,omitempty" url:"nameserver,omitempty,optional"`                 // cloud-init: Sets DNS server IP address for a container. Create will' 	    .' automatically use the setting from the host if neither searchdomain nor nameserver' 	    .' are set.
	Rng0              string   `json:"rng0,omitempty" url:"rng0,omitempty,optional"`                             // Configure a VirtIO-based Random Number Generator.
	Sockets           int      `json:"sockets,omitempty" url:"sockets,omitempty,optional"`                       // The number of CPU sockets.
	Vmid              string   `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`                   // The (unique) ID of the VM.
	Cdrom             string   `json:"cdrom,omitempty" url:"cdrom,omitempty,optional"`                           // This is an alias for option -ide2
	Cipassword        string   `json:"cipassword,omitempty" url:"cipassword,omitempty,optional"`                 // cloud-init: Password to assign the user. Using this is generally not recommended. Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords.
	Smbios1           string   `json:"smbios1,omitempty" url:"smbios1,omitempty,optional"`                       // Specify SMBIOS type 1 fields.
	Hotplug           string   `json:"hotplug,omitempty" url:"hotplug,omitempty,optional"`                       // Selectively enable hotplug features. This is a comma separated list of hotplug features: 'network', 'disk', 'cpu', 'memory' and 'usb'. Use '0' to disable hotplug completely. Using '1' as value is an alias for the default `network,disk,usb`.
	Sata              []string `json:"sata[n],omitempty" url:"sata[n],omitempty,optional"`                       // Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
}

// NodesQemuConfigUpdateVmRequest config - Set virtual machine options (synchrounous API) - You should consider using the POST method instead for any actions involving hotplug or storage allocation.
// Set virtual machine options (synchrounous API) - You should consider using the POST method instead for any actions involving hotplug or storage allocation.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/config
//
type NodesQemuConfigUpdateVmRequest struct {
	Boot              string   `json:"boot,omitempty" url:"boot,omitempty,optional"`                             // Specify guest boot order. Use the 'order=' sub-property as usage with no key or 'legacy=' is deprecated.
	Cores             int      `json:"cores,omitempty" url:"cores,omitempty,optional"`                           // The number of cores per socket.
	Ide               []string `json:"ide[n],omitempty" url:"ide[n],omitempty,optional"`                         // Use volume as IDE hard disk or CD-ROM (n is 0 to 3). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Numa              []string `json:"numa[n],omitempty" url:"numa[n],omitempty,optional"`                       // NUMA topology.
	Parallel          []string `json:"parallel[n],omitempty" url:"parallel[n],omitempty,optional"`               // Map host parallel devices (n is 0 to 2).
	Rng0              string   `json:"rng0,omitempty" url:"rng0,omitempty,optional"`                             // Configure a VirtIO-based Random Number Generator.
	Startup           string   `json:"startup,omitempty" url:"startup,omitempty,optional"`                       // Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Vcpus             int      `json:"vcpus,omitempty" url:"vcpus,omitempty,optional"`                           // Number of hotplugged vcpus.
	Acpi              bool     `json:"acpi,omitempty" url:"acpi,omitempty,optional"`                             // Enable/disable ACPI.
	Arch              string   `json:"arch,omitempty" url:"arch,omitempty,optional"`                             // Virtual processor architecture. Defaults to the host.
	Bios              string   `json:"bios,omitempty" url:"bios,omitempty,optional"`                             // Select BIOS implementation.
	Bootdisk          string   `json:"bootdisk,omitempty" url:"bootdisk,omitempty,optional"`                     // Enable booting from specified disk. Deprecated: Use 'boot: order=foo;bar' instead.
	Nameserver        string   `json:"nameserver,omitempty" url:"nameserver,omitempty,optional"`                 // cloud-init: Sets DNS server IP address for a container. Create will' 	    .' automatically use the setting from the host if neither searchdomain nor nameserver' 	    .' are set.
	Usb               []string `json:"usb[n],omitempty" url:"usb[n],omitempty,optional"`                         // Configure an USB device (n is 0 to 4).
	Autostart         bool     `json:"autostart,omitempty" url:"autostart,omitempty,optional"`                   // Automatic restart after crash (currently ignored).
	Cpulimit          int      `json:"cpulimit,omitempty" url:"cpulimit,omitempty,optional"`                     // Limit of CPU usage.
	Onboot            bool     `json:"onboot,omitempty" url:"onboot,omitempty,optional"`                         // Specifies whether a VM will be started during system bootup.
	Scsi              []string `json:"scsi[n],omitempty" url:"scsi[n],omitempty,optional"`                       // Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Tags              string   `json:"tags,omitempty" url:"tags,omitempty,optional"`                             // Tags of the VM. This is only meta information.
	Ivshmem           string   `json:"ivshmem,omitempty" url:"ivshmem,omitempty,optional"`                       // Inter-VM shared memory. Useful for direct communication between VMs, or to the host.
	Name              string   `json:"name,omitempty" url:"name,omitempty,optional"`                             // Set a name for the VM. Only used on the configuration web interface.
	Watchdog          string   `json:"watchdog,omitempty" url:"watchdog,omitempty,optional"`                     // Create a virtual hardware watchdog device.
	Virtio            []string `json:"virtio[n],omitempty" url:"virtio[n],omitempty,optional"`                   // Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Vmgenid           string   `json:"vmgenid,omitempty" url:"vmgenid,omitempty,optional"`                       // Set VM Generation ID. Use '1' to autogenerate on create or update, pass '0' to disable explicitly.
	Cpuunits          int      `json:"cpuunits,omitempty" url:"cpuunits,omitempty,optional"`                     // CPU weight for a VM, will be clamped to [1, 10000] in cgroup v2.
	Hostpci           []string `json:"hostpci[n],omitempty" url:"hostpci[n],omitempty,optional"`                 // Map host PCI devices into guest.
	Lock              string   `json:"lock,omitempty" url:"lock,omitempty,optional"`                             // Lock/unlock the VM.
	Net               []string `json:"net[n],omitempty" url:"net[n],omitempty,optional"`                         // Specify network devices.
	Tpmstate0         string   `json:"tpmstate0,omitempty" url:"tpmstate0,omitempty,optional"`                   // Configure a Disk for storing TPM state. The format is fixed to 'raw'. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and 4 MiB will be used instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Agent             string   `json:"agent,omitempty" url:"agent,omitempty,optional"`                           // Enable/disable communication with the Qemu Guest Agent and its properties.
	Digest            string   `json:"digest,omitempty" url:"digest,omitempty,optional"`                         // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Ostype            string   `json:"ostype,omitempty" url:"ostype,omitempty,optional"`                         // Specify guest operating system.
	Tdf               bool     `json:"tdf,omitempty" url:"tdf,omitempty,optional"`                               // Enable/disable time drift fix.
	Template          bool     `json:"template,omitempty" url:"template,omitempty,optional"`                     // Enable/disable Template.
	Vmstatestorage    string   `json:"vmstatestorage,omitempty" url:"vmstatestorage,omitempty,optional"`         // Default storage for VM state volumes/files.
	Cpu               string   `json:"cpu,omitempty" url:"cpu,omitempty,optional"`                               // Emulated CPU type.
	Freeze            bool     `json:"freeze,omitempty" url:"freeze,omitempty,optional"`                         // Freeze CPU at startup (use 'c' monitor command to start execution).
	MigrateDowntime   int      `json:"migrate_downtime,omitempty" url:"migrate_downtime,omitempty,optional"`     // Set maximum tolerated downtime (in seconds) for migrations.
	Sata              []string `json:"sata[n],omitempty" url:"sata[n],omitempty,optional"`                       // Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Scsihw            string   `json:"scsihw,omitempty" url:"scsihw,omitempty,optional"`                         // SCSI controller model
	Hugepages         string   `json:"hugepages,omitempty" url:"hugepages,omitempty,optional"`                   // Enable/disable hugepages memory.
	Kvm               bool     `json:"kvm,omitempty" url:"kvm,omitempty,optional"`                               // Enable/disable KVM hardware virtualization.
	Cipassword        string   `json:"cipassword,omitempty" url:"cipassword,omitempty,optional"`                 // cloud-init: Password to assign the user. Using this is generally not recommended. Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords.
	Delete            string   `json:"delete,omitempty" url:"delete,omitempty,optional"`                         // A list of settings you want to delete.
	Node              string   `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                   // The cluster node name.
	NumaEnabled       bool     `json:"numa,omitempty" url:"numa,omitempty,optional"`                             // Enable/disable NUMA.
	Shares            int      `json:"shares,omitempty" url:"shares,omitempty,optional"`                         // Amount of memory shares for auto-ballooning. The larger the number is, the more memory this VM gets. Number is relative to weights of all other running VMs. Using zero disables auto-ballooning. Auto-ballooning is done by pvestatd.
	Keyboard          string   `json:"keyboard,omitempty" url:"keyboard,omitempty,optional"`                     // Keyboard layout for VNC server. This option is generally not required and is often better handled from within the guest OS.
	Startdate         string   `json:"startdate,omitempty" url:"startdate,omitempty,optional"`                   // Set the initial date of the real time clock. Valid format for date are:'now' or '2006-06-17T16:01:21' or '2006-06-17'.
	Tablet            bool     `json:"tablet,omitempty" url:"tablet,omitempty,optional"`                         // Enable/disable the USB tablet device.
	Unused            []string `json:"unused[n],omitempty" url:"unused[n],omitempty,optional"`                   // Reference to unused volumes. This is used internally, and should not be modified manually.
	Vga               string   `json:"vga,omitempty" url:"vga,omitempty,optional"`                               // Configure the VGA hardware.
	Searchdomain      string   `json:"searchdomain,omitempty" url:"searchdomain,omitempty,optional"`             // cloud-init: Sets DNS search domains for a container. Create will' 	    .' automatically use the setting from the host if neither searchdomain nor nameserver' 	    .' are set.
	Sshkeys           string   `json:"sshkeys,omitempty" url:"sshkeys,omitempty,optional"`                       // cloud-init: Setup public SSH keys (one key per line, OpenSSH format).
	Args              string   `json:"args,omitempty" url:"args,omitempty,optional"`                             // Arbitrary arguments passed to kvm.
	Balloon           int      `json:"balloon,omitempty" url:"balloon,omitempty,optional"`                       // Amount of target RAM for the VM in MB. Using zero disables the ballon driver.
	Cdrom             string   `json:"cdrom,omitempty" url:"cdrom,omitempty,optional"`                           // This is an alias for option -ide2
	Ciuser            string   `json:"ciuser,omitempty" url:"ciuser,omitempty,optional"`                         // cloud-init: User name to change ssh keys and password for instead of the image's configured default user.
	Description       string   `json:"description,omitempty" url:"description,omitempty,optional"`               // Description for the VM. Shown in the web-interface VM's summary. This is saved as comment inside the configuration file.
	Hotplug           string   `json:"hotplug,omitempty" url:"hotplug,omitempty,optional"`                       // Selectively enable hotplug features. This is a comma separated list of hotplug features: 'network', 'disk', 'cpu', 'memory' and 'usb'. Use '0' to disable hotplug completely. Using '1' as value is an alias for the default `network,disk,usb`.
	SpiceEnhancements string   `json:"spice_enhancements,omitempty" url:"spice_enhancements,omitempty,optional"` // Configure additional enhancements for SPICE.
	Audio0            string   `json:"audio0,omitempty" url:"audio0,omitempty,optional"`                         // Configure a audio device, useful in combination with QXL/Spice.
	Force             bool     `json:"force,omitempty" url:"force,omitempty,optional"`                           // Force physical removal. Without this, we simple remove the disk from the config file and create an additional configuration entry called 'unused[n]', which contains the volume ID. Unlink of unused[n] always cause physical removal.
	Keephugepages     bool     `json:"keephugepages,omitempty" url:"keephugepages,omitempty,optional"`           // Use together with hugepages. If enabled, hugepages will not not be deleted after VM shutdown and can be used for subsequent starts.
	Serial            []string `json:"serial[n],omitempty" url:"serial[n],omitempty,optional"`                   // Create a serial device inside the VM (n is 0 to 3)
	Cicustom          string   `json:"cicustom,omitempty" url:"cicustom,omitempty,optional"`                     // cloud-init: Specify custom files to replace the automatically generated ones at start.
	Localtime         bool     `json:"localtime,omitempty" url:"localtime,omitempty,optional"`                   // Set the real time clock (RTC) to local time. This is enabled by default if the `ostype` indicates a Microsoft Windows OS.
	Machine           string   `json:"machine,omitempty" url:"machine,omitempty,optional"`                       // Specifies the Qemu machine type.
	Sockets           int      `json:"sockets,omitempty" url:"sockets,omitempty,optional"`                       // The number of CPU sockets.
	Vmid              string   `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`                   // The (unique) ID of the VM.
	Hookscript        string   `json:"hookscript,omitempty" url:"hookscript,omitempty,optional"`                 // Script that will be executed during various steps in the vms lifetime.
	Memory            int      `json:"memory,omitempty" url:"memory,omitempty,optional"`                         // Amount of RAM for the VM in MB. This is the maximum available memory when you use the balloon device.
	Skiplock          bool     `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`                     // Ignore locks - only root is allowed to use this option.
	Smbios1           string   `json:"smbios1,omitempty" url:"smbios1,omitempty,optional"`                       // Specify SMBIOS type 1 fields.
	Reboot            bool     `json:"reboot,omitempty" url:"reboot,omitempty,optional"`                         // Allow reboot. If set to '0' the VM exit on reboot.
	Revert            string   `json:"revert,omitempty" url:"revert,omitempty,optional"`                         // Revert a pending change.
	Smp               int      `json:"smp,omitempty" url:"smp,omitempty,optional"`                               // The number of CPUs. Please use option -sockets instead.
	Citype            string   `json:"citype,omitempty" url:"citype,omitempty,optional"`                         // Specifies the cloud-init configuration format. The default depends on the configured operating system type (`ostype`. We use the `nocloud` format for Linux, and `configdrive2` for windows.
	Efidisk0          string   `json:"efidisk0,omitempty" url:"efidisk0,omitempty,optional"`                     // Configure a Disk for storing EFI vars. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and that the default EFI vars are copied to the volume instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Ipconfig          []string `json:"ipconfig[n],omitempty" url:"ipconfig[n],omitempty,optional"`               // cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string 'dhcp' can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string 'auto' can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.
	MigrateSpeed      int      `json:"migrate_speed,omitempty" url:"migrate_speed,omitempty,optional"`           // Set maximum speed (in MB/s) for migrations. Value 0 is no limit.
	Protection        bool     `json:"protection,omitempty" url:"protection,omitempty,optional"`                 // Sets the protection flag of the VM. This will disable the remove VM and remove disk operations.
}

// NodesQemuPendingVmRequest pending - Get the virtual machine configuration with both current and pending values.
// Get the virtual machine configuration with both current and pending values.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/pending
//
type NodesQemuPendingVmRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuPendingVmResponse pending
// Get the virtual machine configuration with both current and pending values.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/pending
//
type NodesQemuPendingVmResponse struct {
	Delete  int    `json:"delete,omitempty" url:"delete,omitempty,optional"`     // Indicates a pending delete request if present and not 0. The value 2 indicates a force-delete request.
	Key     string `json:"key,omitempty" url:"key,omitempty" validate:"nonzero"` // Configuration option name.
	Pending string `json:"pending,omitempty" url:"pending,omitempty,optional"`   // Pending value.
	Value   string `json:"value,omitempty" url:"value,omitempty,optional"`       // Current value.
}

// NodesQemuUnlinkRequest unlink - Unlink/delete disk images.
// Unlink/delete disk images.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/unlink
//
type NodesQemuUnlinkRequest struct {
	Force  bool   `json:"force,omitempty" url:"force,omitempty,optional"`             // Force physical removal. Without this, we simple remove the disk from the config file and create an additional configuration entry called 'unused[n]', which contains the volume ID. Unlink of unused[n] always cause physical removal.
	Idlist string `json:"idlist,omitempty" url:"idlist,omitempty" validate:"nonzero"` // A list of disk IDs you want to delete.
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Vmid   string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`     // The (unique) ID of the VM.
}

// NodesQemuVncproxyRequest vncproxy - Creates a TCP VNC proxy connections.
// Creates a TCP VNC proxy connections.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/vncproxy
//
type NodesQemuVncproxyRequest struct {
	Node             string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                 // The cluster node name.
	Vmid             string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`                 // The (unique) ID of the VM.
	Websocket        bool   `json:"websocket,omitempty" url:"websocket,omitempty,optional"`                 // starts websockify instead of vncproxy
	GeneratePassword bool   `json:"generate-password,omitempty" url:"generate-password,omitempty,optional"` // Generates a random password to be used as ticket instead of the API ticket.
}

// NodesQemuVncproxyResponse vncproxy
// Creates a TCP VNC proxy connections.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/vncproxy
//
type NodesQemuVncproxyResponse struct {
	User     string `json:"user,omitempty" url:"user,omitempty" validate:"nonzero"`     //
	Cert     string `json:"cert,omitempty" url:"cert,omitempty" validate:"nonzero"`     //
	Password string `json:"password,omitempty" url:"password,omitempty,optional"`       // Returned if requested with 'generate-password' param. Consists of printable ASCII characters ('!' .. '~').
	Port     int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`     //
	Ticket   string `json:"ticket,omitempty" url:"ticket,omitempty" validate:"nonzero"` //
	Upid     string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"`     //
}

// NodesQemuTermproxyRequest termproxy - Creates a TCP proxy connections.
// Creates a TCP proxy connections.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/termproxy
//
type NodesQemuTermproxyRequest struct {
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Serial string `json:"serial,omitempty" url:"serial,omitempty,optional"`       // opens a serial terminal (defaults to display)
	Vmid   string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuTermproxyResponse termproxy
// Creates a TCP proxy connections.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/termproxy
//
type NodesQemuTermproxyResponse struct {
	Port   int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`     //
	Ticket string `json:"ticket,omitempty" url:"ticket,omitempty" validate:"nonzero"` //
	Upid   string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"`     //
	User   string `json:"user,omitempty" url:"user,omitempty" validate:"nonzero"`     //
}

// NodesQemuVncwebsocketRequest vncwebsocket - Opens a weksocket for VNC traffic.
// Opens a weksocket for VNC traffic.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/vncwebsocket
//
type NodesQemuVncwebsocketRequest struct {
	Vncticket string `json:"vncticket,omitempty" url:"vncticket,omitempty" validate:"nonzero"` // Ticket from previous call to vncproxy.
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Port      int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`           // Port number returned by previous vncproxy call.
	Vmid      string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`           // The (unique) ID of the VM.
}

// NodesQemuVncwebsocketResponse vncwebsocket
// Opens a weksocket for VNC traffic.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/vncwebsocket
//
type NodesQemuVncwebsocketResponse struct {
	Port string `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"` //
}

// NodesQemuSpiceproxyRequest spiceproxy - Returns a SPICE configuration to connect to the VM.
// Returns a SPICE configuration to connect to the VM.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/spiceproxy
//
type NodesQemuSpiceproxyRequest struct {
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Proxy string `json:"proxy,omitempty" url:"proxy,omitempty,optional"`         // SPICE proxy server. This can be used by the client to specify the proxy server. All nodes in a cluster runs 'spiceproxy', so it is up to the client to choose one. By default, we return the node where the VM is currently running. As reasonable setting is to use same node you use to connect to the API (This is window.location.hostname for the JS GUI).
	Vmid  string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuSpiceproxyResponse spiceproxy
// Returns a SPICE configuration to connect to the VM.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/spiceproxy
//
type NodesQemuSpiceproxyResponse struct {
	TlsPort  int    `json:"tls-port,omitempty" url:"tls-port,omitempty" validate:"nonzero"` //
	Type     string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`         //
	Host     string `json:"host,omitempty" url:"host,omitempty" validate:"nonzero"`         //
	Password string `json:"password,omitempty" url:"password,omitempty" validate:"nonzero"` //
	Proxy    string `json:"proxy,omitempty" url:"proxy,omitempty" validate:"nonzero"`       //
}

// NodesQemuStatusVmcmdidxRequest status - Directory index
// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status
//
type NodesQemuStatusVmcmdidxRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuStatusVmcmdidxResponse status
// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status
//
type NodesQemuStatusVmcmdidxResponse struct {
	Subdir string `json:"subdir,omitempty" url:"subdir,omitempty" validate:"nonzero"` //
}

// NodesQemuStatusCurrentVmStatusRequest current - Get virtual machine status.
// Get virtual machine status.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status/current
//
type NodesQemuStatusCurrentVmStatusRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuStatusCurrentVmStatusResponse current
// Get virtual machine status.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status/current
//
type NodesQemuStatusCurrentVmStatusResponse struct {
	Pid            int      `json:"pid,omitempty" url:"pid,omitempty,optional"`                         // PID of running qemu process.
	Qmpstatus      string   `json:"qmpstatus,omitempty" url:"qmpstatus,omitempty,optional"`             // Qemu QMP agent status.
	Spice          bool     `json:"spice,omitempty" url:"spice,omitempty,optional"`                     // Qemu VGA configuration supports spice.
	Vmid           string   `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`             // The (unique) ID of the VM.
	Name           string   `json:"name,omitempty" url:"name,omitempty,optional"`                       // VM name.
	RunningMachine string   `json:"running-machine,omitempty" url:"running-machine,omitempty,optional"` // The currently running machine type (if running).
	Cpus           int      `json:"cpus,omitempty" url:"cpus,omitempty,optional"`                       // Maximum usable CPUs.
	Ha             struct { //
	} `json:"ha,omitempty" url:"ha,omitempty" validate:"nonzero"` // HA manager service status.
	Lock        string `json:"lock,omitempty" url:"lock,omitempty,optional"`                 // The current config lock, if any.
	Maxdisk     int    `json:"maxdisk,omitempty" url:"maxdisk,omitempty,optional"`           // Root disk size in bytes.
	Maxmem      int    `json:"maxmem,omitempty" url:"maxmem,omitempty,optional"`             // Maximum memory in bytes.
	Uptime      int    `json:"uptime,omitempty" url:"uptime,omitempty,optional"`             // Uptime.
	Agent       bool   `json:"agent,omitempty" url:"agent,omitempty,optional"`               // Qemu GuestAgent enabled in config.
	RunningQemu string `json:"running-qemu,omitempty" url:"running-qemu,omitempty,optional"` // The currently running QEMU version (if running).
	Status      string `json:"status,omitempty" url:"status,omitempty" validate:"nonzero"`   // Qemu process status.
	Tags        string `json:"tags,omitempty" url:"tags,omitempty,optional"`                 // The current configured tags, if any
}

// NodesQemuStatusStartVmRequest start - Start virtual machine.
// Start virtual machine.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status/start
//
type NodesQemuStatusStartVmRequest struct {
	Timeout          int    `json:"timeout,omitempty" url:"timeout,omitempty,optional"`                     // Wait maximal timeout seconds.
	Vmid             string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`                 // The (unique) ID of the VM.
	ForceCpu         string `json:"force-cpu,omitempty" url:"force-cpu,omitempty,optional"`                 // Override QEMU's -cpu argument with the given string.
	MigrationType    string `json:"migration_type,omitempty" url:"migration_type,omitempty,optional"`       // Migration traffic is encrypted using an SSH tunnel by default. On secure, completely private networks this can be disabled to increase performance.
	Skiplock         bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`                   // Ignore locks - only root is allowed to use this option.
	Stateuri         string `json:"stateuri,omitempty" url:"stateuri,omitempty,optional"`                   // Some command save/restore state from this location.
	Targetstorage    string `json:"targetstorage,omitempty" url:"targetstorage,omitempty,optional"`         // Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	Machine          string `json:"machine,omitempty" url:"machine,omitempty,optional"`                     // Specifies the Qemu machine type.
	Migratedfrom     string `json:"migratedfrom,omitempty" url:"migratedfrom,omitempty,optional"`           // The cluster node name.
	MigrationNetwork string `json:"migration_network,omitempty" url:"migration_network,omitempty,optional"` // CIDR of the (sub) network that is used for migration.
	Node             string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                 // The cluster node name.
}

// NodesQemuStatusStopVmRequest stop - Stop virtual machine. The qemu process will exit immediately. Thisis akin to pulling the power plug of a running computer and may damage the VM data
// Stop virtual machine. The qemu process will exit immediately. Thisis akin to pulling the power plug of a running computer and may damage the VM data
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status/stop
//
type NodesQemuStatusStopVmRequest struct {
	Timeout      int    `json:"timeout,omitempty" url:"timeout,omitempty,optional"`           // Wait maximal timeout seconds.
	Vmid         string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`       // The (unique) ID of the VM.
	KeepActive   bool   `json:"keepActive,omitempty" url:"keepActive,omitempty,optional"`     // Do not deactivate storage volumes.
	Migratedfrom string `json:"migratedfrom,omitempty" url:"migratedfrom,omitempty,optional"` // The cluster node name.
	Node         string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Skiplock     bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`         // Ignore locks - only root is allowed to use this option.
}

// NodesQemuStatusResetVmRequest reset - Reset virtual machine.
// Reset virtual machine.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status/reset
//
type NodesQemuStatusResetVmRequest struct {
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Skiplock bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`   // Ignore locks - only root is allowed to use this option.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuStatusShutdownVmRequest shutdown - Shutdown virtual machine. This is similar to pressing the power button on a physical machine.This will send an ACPI event for the guest OS, which should then proceed to a clean shutdown.
// Shutdown virtual machine. This is similar to pressing the power button on a physical machine.This will send an ACPI event for the guest OS, which should then proceed to a clean shutdown.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status/shutdown
//
type NodesQemuStatusShutdownVmRequest struct {
	KeepActive bool   `json:"keepActive,omitempty" url:"keepActive,omitempty,optional"` // Do not deactivate storage volumes.
	Node       string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
	Skiplock   bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`     // Ignore locks - only root is allowed to use this option.
	Timeout    int    `json:"timeout,omitempty" url:"timeout,omitempty,optional"`       // Wait maximal timeout seconds.
	Vmid       string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`   // The (unique) ID of the VM.
	ForceStop  bool   `json:"forceStop,omitempty" url:"forceStop,omitempty,optional"`   // Make sure the VM stops.
}

// NodesQemuStatusRebootVmRequest reboot - Reboot the VM by shutting it down, and starting it again. Applies pending changes.
// Reboot the VM by shutting it down, and starting it again. Applies pending changes.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status/reboot
//
type NodesQemuStatusRebootVmRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Timeout int    `json:"timeout,omitempty" url:"timeout,omitempty,optional"`     // Wait maximal timeout seconds for the shutdown.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuStatusSuspendVmRequest suspend - Suspend virtual machine.
// Suspend virtual machine.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status/suspend
//
type NodesQemuStatusSuspendVmRequest struct {
	Skiplock     bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`         // Ignore locks - only root is allowed to use this option.
	Statestorage string `json:"statestorage,omitempty" url:"statestorage,omitempty,optional"` // The storage for the VM state
	Todisk       bool   `json:"todisk,omitempty" url:"todisk,omitempty,optional"`             // If set, suspends the VM to disk. Will be resumed on next VM start.
	Vmid         string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`       // The (unique) ID of the VM.
	Node         string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
}

// NodesQemuStatusResumeVmRequest resume - Resume virtual machine.
// Resume virtual machine.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/status/resume
//
type NodesQemuStatusResumeVmRequest struct {
	Nocheck  bool   `json:"nocheck,omitempty" url:"nocheck,omitempty,optional"`     //
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Skiplock bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`   // Ignore locks - only root is allowed to use this option.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuSendkeyVmRequest sendkey - Send key event to virtual machine.
// Send key event to virtual machine.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/sendkey
//
type NodesQemuSendkeyVmRequest struct {
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Key      string `json:"key,omitempty" url:"key,omitempty" validate:"nonzero"`   // The key (qemu monitor encoding).
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Skiplock bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`   // Ignore locks - only root is allowed to use this option.
}

// NodesQemuFeatureVmRequest feature - Check if feature for virtual machine is available.
// Check if feature for virtual machine is available.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/feature
//
type NodesQemuFeatureVmRequest struct {
	Feature  string `json:"feature,omitempty" url:"feature,omitempty" validate:"nonzero"` // Feature to check.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Snapname string `json:"snapname,omitempty" url:"snapname,omitempty,optional"`         // The name of the snapshot.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`       // The (unique) ID of the VM.
}

// NodesQemuFeatureVmResponse feature
// Check if feature for virtual machine is available.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/feature
//
type NodesQemuFeatureVmResponse struct {
	HasFeature bool `json:"hasFeature,omitempty" url:"hasFeature,omitempty" validate:"nonzero"` //
	Nodes      []struct {
	} `json:"nodes,omitempty" url:"nodes,omitempty" validate:"nonzero"` //
}

// NodesQemuCloneVmRequest clone - Create a copy of virtual machine/template.
// Create a copy of virtual machine/template.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/clone
//
type NodesQemuCloneVmRequest struct {
	Newid       int    `json:"newid,omitempty" url:"newid,omitempty" validate:"nonzero"`   // VMID for the clone.
	Pool        string `json:"pool,omitempty" url:"pool,omitempty,optional"`               // Add the new VM to the specified pool.
	Snapname    string `json:"snapname,omitempty" url:"snapname,omitempty,optional"`       // The name of the snapshot.
	Storage     string `json:"storage,omitempty" url:"storage,omitempty,optional"`         // Target storage for full clone.
	Target      string `json:"target,omitempty" url:"target,omitempty,optional"`           // Target node. Only allowed if the original VM is on shared storage.
	Description string `json:"description,omitempty" url:"description,omitempty,optional"` // Description for the new VM.
	Format      string `json:"format,omitempty" url:"format,omitempty,optional"`           // Target format for file storage. Only valid for full clone.
	Name        string `json:"name,omitempty" url:"name,omitempty,optional"`               // Set a name for the new VM.
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Vmid        string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`     // The (unique) ID of the VM.
	Bwlimit     int    `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`         // Override I/O bandwidth limit (in KiB/s).
	Full        bool   `json:"full,omitempty" url:"full,omitempty,optional"`               // Create a full copy of all disks. This is always done when you clone a normal VM. For VM templates, we try to create a linked clone by default.
}

// NodesQemuMoveDiskMoveVmDiskRequest move_disk - Move volume to different storage or to a different VM.
// Move volume to different storage or to a different VM.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/move_disk
//
type NodesQemuMoveDiskMoveVmDiskRequest struct {
	Vmid         string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Bwlimit      int    `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`             // Override I/O bandwidth limit (in KiB/s).
	Disk         string `json:"disk,omitempty" url:"disk,omitempty" validate:"nonzero"`         // The disk you want to move.
	Format       string `json:"format,omitempty" url:"format,omitempty,optional"`               // Target Format.
	Storage      string `json:"storage,omitempty" url:"storage,omitempty,optional"`             // Target storage.
	TargetDigest string `json:"target-digest,omitempty" url:"target-digest,omitempty,optional"` // Prevent changes if the current config file of the target VM has a" 		    ." different SHA1 digest. This can be used to detect concurrent modifications.
	TargetVmid   int    `json:"target-vmid,omitempty" url:"target-vmid,omitempty,optional"`     // The (unique) ID of the VM.
	Delete       bool   `json:"delete,omitempty" url:"delete,omitempty,optional"`               // Delete the original disk after successful copy. By default the original disk is kept as unused disk.
	Digest       string `json:"digest,omitempty" url:"digest,omitempty,optional"`               // Prevent changes if current configuration file has different SHA1" 		    ." digest. This can be used to prevent concurrent modifications.
	Node         string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	TargetDisk   string `json:"target-disk,omitempty" url:"target-disk,omitempty,optional"`     // The config key the disk will be moved to on the target VM (for example, ide0 or scsi1). Default is the source disk key.
}

// NodesQemuMigrateVmPreconditionRequest migrate - Get preconditions for migration.
// Get preconditions for migration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/migrate
//
type NodesQemuMigrateVmPreconditionRequest struct {
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Target string `json:"target,omitempty" url:"target,omitempty,optional"`       // Target node.
	Vmid   string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuMigrateVmPreconditionResponse migrate
// Get preconditions for migration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/migrate
//
type NodesQemuMigrateVmPreconditionResponse struct {
	LocalDisks []struct {
	} `json:"local_disks,omitempty" url:"local_disks,omitempty" validate:"nonzero"` // List local disks including CD-Rom, unsused and not referenced disks
	LocalResources []struct {
	} `json:"local_resources,omitempty" url:"local_resources,omitempty" validate:"nonzero"` // List local resources e.g. pci, usb
	NotAllowedNodes struct { //
	} `json:"not_allowed_nodes,omitempty" url:"not_allowed_nodes,omitempty,optional"` // List not allowed nodes with additional informations, only passed if VM is offline
	Running      bool `json:"running,omitempty" url:"running,omitempty" validate:"nonzero"` //
	AllowedNodes []struct {
	} `json:"allowed_nodes,omitempty" url:"allowed_nodes,omitempty,optional"` // List nodes allowed for offline migration, only passed if VM is offline
}

// NodesQemuMigrateVmRequest migrate - Migrate virtual machine. Creates a new migration task.
// Migrate virtual machine. Creates a new migration task.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/migrate
//
type NodesQemuMigrateVmRequest struct {
	Force            bool   `json:"force,omitempty" url:"force,omitempty,optional"`                         // Allow to migrate VMs which use local devices. Only root may use this option.
	Node             string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                 // The cluster node name.
	Vmid             string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`                 // The (unique) ID of the VM.
	WithLocalDisks   bool   `json:"with-local-disks,omitempty" url:"with-local-disks,omitempty,optional"`   // Enable live storage migration for local disk
	Bwlimit          int    `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`                     // Override I/O bandwidth limit (in KiB/s).
	MigrationNetwork string `json:"migration_network,omitempty" url:"migration_network,omitempty,optional"` // CIDR of the (sub) network that is used for migration.
	MigrationType    string `json:"migration_type,omitempty" url:"migration_type,omitempty,optional"`       // Migration traffic is encrypted using an SSH tunnel by default. On secure, completely private networks this can be disabled to increase performance.
	Online           bool   `json:"online,omitempty" url:"online,omitempty,optional"`                       // Use online/live migration if VM is running. Ignored if VM is stopped.
	Target           string `json:"target,omitempty" url:"target,omitempty" validate:"nonzero"`             // Target node.
	Targetstorage    string `json:"targetstorage,omitempty" url:"targetstorage,omitempty,optional"`         // Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
}

// NodesQemuMonitorRequest monitor - Execute Qemu monitor commands.
// Execute Qemu monitor commands.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/monitor
//
type NodesQemuMonitorRequest struct {
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`       // The (unique) ID of the VM.
	Command string `json:"command,omitempty" url:"command,omitempty" validate:"nonzero"` // The monitor command.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
}

// NodesQemuResizeVmRequest resize - Extend volume size.
// Extend volume size.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/resize
//
type NodesQemuResizeVmRequest struct {
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Disk     string `json:"disk,omitempty" url:"disk,omitempty" validate:"nonzero"` // The disk you want to resize.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Size     string `json:"size,omitempty" url:"size,omitempty" validate:"nonzero"` // The new size. With the `+` sign the value is added to the actual size of the volume and without it, the value is taken as an absolute one. Shrinking disk size is not supported.
	Skiplock bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`   // Ignore locks - only root is allowed to use this option.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuSnapshotListRequest snapshot - List all snapshots.
// List all snapshots.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/snapshot
//
type NodesQemuSnapshotListRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuSnapshotListResponse snapshot
// List all snapshots.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/snapshot
//
type NodesQemuSnapshotListResponse struct {
	Description string `json:"description,omitempty" url:"description,omitempty" validate:"nonzero"` // Snapshot description.
	Name        string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`               // Snapshot identifier. Value 'current' identifies the current VM.
	Parent      string `json:"parent,omitempty" url:"parent,omitempty,optional"`                     // Parent snapshot identifier.
	Snaptime    int    `json:"snaptime,omitempty" url:"snaptime,omitempty,optional"`                 // Snapshot creation time
	Vmstate     bool   `json:"vmstate,omitempty" url:"vmstate,omitempty,optional"`                   // Snapshot includes RAM.
}

// NodesQemuSnapshotRequest snapshot - Snapshot a VM.
// Snapshot a VM.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/snapshot
//
type NodesQemuSnapshotRequest struct {
	Snapname    string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
	Vmid        string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Vmstate     bool   `json:"vmstate,omitempty" url:"vmstate,omitempty,optional"`             // Save the vmstate
	Description string `json:"description,omitempty" url:"description,omitempty,optional"`     // A textual description or comment.
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
}

// NodesQemuSnapshotDelsnapshotRequest {snapname} - Delete a VM snapshot.
// Delete a VM snapshot.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/snapshot/{snapname}
//
type NodesQemuSnapshotDelsnapshotRequest struct {
	Snapname string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Force    bool   `json:"force,omitempty" url:"force,omitempty,optional"`                 // For removal from config file, even if removing disk snapshots fails.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
}

// NodesQemuSnapshotCmdIdxRequest {snapname} -
//
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/snapshot/{snapname}
//
type NodesQemuSnapshotCmdIdxRequest struct {
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Snapname string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
}

// NodesQemuSnapshotConfigGetSnapshotRequest config - Get snapshot configuration
// Get snapshot configuration
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/snapshot/{snapname}/config
//
type NodesQemuSnapshotConfigGetSnapshotRequest struct {
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Snapname string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
}

// NodesQemuSnapshotConfigUpdateSnapshotRequest config - Update snapshot metadata.
// Update snapshot metadata.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/snapshot/{snapname}/config
//
type NodesQemuSnapshotConfigUpdateSnapshotRequest struct {
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Snapname    string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
	Vmid        string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Description string `json:"description,omitempty" url:"description,omitempty,optional"`     // A textual description or comment.
}

// NodesQemuSnapshotRollbackRequest rollback - Rollback VM state to specified snapshot.
// Rollback VM state to specified snapshot.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/snapshot/{snapname}/rollback
//
type NodesQemuSnapshotRollbackRequest struct {
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Snapname string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
}

// NodesQemuTemplateRequest template - Create a Template.
// Create a Template.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/template
//
type NodesQemuTemplateRequest struct {
	Disk string `json:"disk,omitempty" url:"disk,omitempty,optional"`           // If you want to convert only 1 disk to base image.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesQemuCloudinitDumpCloudinitGeneratedConfigRequest dump - Get automatically generated cloudinit config.
// Get automatically generated cloudinit config.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/cloudinit/dump
//
type NodesQemuCloudinitDumpCloudinitGeneratedConfigRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Type string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"` // Config type.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcVmlistRequest lxc - LXC container index (per node).
// LXC container index (per node).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc
//
type NodesLxcVmlistRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesLxcVmlistResponse lxc
// LXC container index (per node).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc
//
type NodesLxcVmlistResponse struct {
	Maxdisk int    `json:"maxdisk,omitempty" url:"maxdisk,omitempty,optional"`         // Root disk size in bytes.
	Name    string `json:"name,omitempty" url:"name,omitempty,optional"`               // Container name.
	Status  string `json:"status,omitempty" url:"status,omitempty" validate:"nonzero"` // LXC Container status.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`     // The (unique) ID of the VM.
	Cpus    int    `json:"cpus,omitempty" url:"cpus,omitempty,optional"`               // Maximum usable CPUs.
	Lock    string `json:"lock,omitempty" url:"lock,omitempty,optional"`               // The current config lock, if any.
	Maxmem  int    `json:"maxmem,omitempty" url:"maxmem,omitempty,optional"`           // Maximum memory in bytes.
	Maxswap int    `json:"maxswap,omitempty" url:"maxswap,omitempty,optional"`         // Maximum SWAP memory in bytes.
	Tags    string `json:"tags,omitempty" url:"tags,omitempty,optional"`               // The current configured tags, if any.
	Uptime  int    `json:"uptime,omitempty" url:"uptime,omitempty,optional"`           // Uptime.
}

// NodesLxcCreateVmRequest lxc - Create or restore a container.
// Create or restore a container.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc
//
type NodesLxcCreateVmRequest struct {
	Cpulimit           int      `json:"cpulimit,omitempty" url:"cpulimit,omitempty,optional"`                         // Limit of CPU usage.  NOTE: If the computer has 2 CPUs, it has a total of '2' CPU time. Value '0' indicates no CPU limit.
	Mp                 []string `json:"mp[n],omitempty" url:"mp[n],omitempty,optional"`                               // Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	Nameserver         string   `json:"nameserver,omitempty" url:"nameserver,omitempty,optional"`                     // Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Node               string   `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                       // The cluster node name.
	Rootfs             string   `json:"rootfs,omitempty" url:"rootfs,omitempty,optional"`                             // Use volume as container root.
	Storage            string   `json:"storage,omitempty" url:"storage,omitempty,optional"`                           // Default Storage.
	Tags               string   `json:"tags,omitempty" url:"tags,omitempty,optional"`                                 // Tags of the Container. This is only meta information.
	Features           string   `json:"features,omitempty" url:"features,omitempty,optional"`                         // Allow containers access to advanced features.
	Hookscript         string   `json:"hookscript,omitempty" url:"hookscript,omitempty,optional"`                     // Script that will be exectued during various steps in the containers lifetime.
	SshPublicKeys      string   `json:"ssh-public-keys,omitempty" url:"ssh-public-keys,omitempty,optional"`           // Setup public SSH keys (one key per line, OpenSSH format).
	Bwlimit            int      `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`                           // Override I/O bandwidth limit (in KiB/s).
	Lock               string   `json:"lock,omitempty" url:"lock,omitempty,optional"`                                 // Lock/unlock the VM.
	Hostname           string   `json:"hostname,omitempty" url:"hostname,omitempty,optional"`                         // Set a host name for the container.
	Memory             int      `json:"memory,omitempty" url:"memory,omitempty,optional"`                             // Amount of RAM for the VM in MB.
	Ostemplate         string   `json:"ostemplate,omitempty" url:"ostemplate,omitempty" validate:"nonzero"`           // The OS template or backup file.
	Timezone           string   `json:"timezone,omitempty" url:"timezone,omitempty,optional"`                         // Time zone to use in the container. If option isn't set, then nothing will be done. Can be set to 'host' to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab
	Cmode              string   `json:"cmode,omitempty" url:"cmode,omitempty,optional"`                               // Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to 'console' it tries to attach to /dev/console instead. If you set cmode to 'shell', it simply invokes a shell inside the container (no login).
	Force              bool     `json:"force,omitempty" url:"force,omitempty,optional"`                               // Allow to overwrite existing container.
	Debug              bool     `json:"debug,omitempty" url:"debug,omitempty,optional"`                               // Try to be more verbose. For now this only enables debug log-level on start.
	Description        string   `json:"description,omitempty" url:"description,omitempty,optional"`                   // Description for the Container. Shown in the web-interface CT's summary. This is saved as comment inside the configuration file.
	IgnoreUnpackErrors bool     `json:"ignore-unpack-errors,omitempty" url:"ignore-unpack-errors,omitempty,optional"` // Ignore errors when extracting the template.
	Ostype             string   `json:"ostype,omitempty" url:"ostype,omitempty,optional"`                             // OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/<ostype>.common.conf. Value 'unmanaged' can be used to skip and OS specific setup.
	Password           string   `json:"password,omitempty" url:"password,omitempty,optional"`                         // Sets root password inside container.
	Pool               string   `json:"pool,omitempty" url:"pool,omitempty,optional"`                                 // Add the VM to the specified pool.
	Cores              int      `json:"cores,omitempty" url:"cores,omitempty,optional"`                               // The number of cores assigned to the container. A container can use all available cores by default.
	Cpuunits           int      `json:"cpuunits,omitempty" url:"cpuunits,omitempty,optional"`                         // CPU weight for a VM. Argument is used in the kernel fair scheduler. The larger the number is, the more CPU time this VM gets. Number is relative to the weights of all the other running VMs.  NOTE: You can disable fair-scheduler configuration by setting this to 0.
	Swap               int      `json:"swap,omitempty" url:"swap,omitempty,optional"`                                 // Amount of SWAP for the VM in MB.
	Template           bool     `json:"template,omitempty" url:"template,omitempty,optional"`                         // Enable/disable Template.
	Protection         bool     `json:"protection,omitempty" url:"protection,omitempty,optional"`                     // Sets the protection flag of the container. This will prevent the CT or CT's disk remove/update operation.
	Start              bool     `json:"start,omitempty" url:"start,omitempty,optional"`                               // Start the CT after its creation finished successfully.
	Startup            string   `json:"startup,omitempty" url:"startup,omitempty,optional"`                           // Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Onboot             bool     `json:"onboot,omitempty" url:"onboot,omitempty,optional"`                             // Specifies whether a VM will be started during system bootup.
	Restore            bool     `json:"restore,omitempty" url:"restore,omitempty,optional"`                           // Mark this as restore task.
	Unique             bool     `json:"unique,omitempty" url:"unique,omitempty,optional"`                             // Assign a unique random ethernet address.
	Unprivileged       bool     `json:"unprivileged,omitempty" url:"unprivileged,omitempty,optional"`                 // Makes the container run as unprivileged user. (Should not be modified manually.)
	Vmid               string   `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`                       // The (unique) ID of the VM.
	Net                []string `json:"net[n],omitempty" url:"net[n],omitempty,optional"`                             // Specifies network interfaces for the container.
	Tty                int      `json:"tty,omitempty" url:"tty,omitempty,optional"`                                   // Specify the number of tty available to the container
	Searchdomain       string   `json:"searchdomain,omitempty" url:"searchdomain,omitempty,optional"`                 // Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Unused             []string `json:"unused[n],omitempty" url:"unused[n],omitempty,optional"`                       // Reference to unused volumes. This is used internally, and should not be modified manually.
	Arch               string   `json:"arch,omitempty" url:"arch,omitempty,optional"`                                 // OS architecture type.
	Console            bool     `json:"console,omitempty" url:"console,omitempty,optional"`                           // Attach a console device (/dev/console) to the container.
}

// NodesLxcVmdiridxRequest {vmid} - Directory index
// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}
//
type NodesLxcVmdiridxRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcVmdiridxResponse {vmid}
// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}
//
type NodesLxcVmdiridxResponse struct {
	Subdir string `json:"subdir,omitempty" url:"subdir,omitempty" validate:"nonzero"` //
}

// NodesLxcDestroyVmRequest {vmid} - Destroy the container (also delete all uses files).
// Destroy the container (also delete all uses files).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}
//
type NodesLxcDestroyVmRequest struct {
	DestroyUnreferencedDisks bool   `json:"destroy-unreferenced-disks,omitempty" url:"destroy-unreferenced-disks,omitempty,optional"` // If set, destroy additionally all disks with the VMID from all enabled storages which are not referenced in the config.
	Force                    bool   `json:"force,omitempty" url:"force,omitempty,optional"`                                           // Force destroy, even if running.
	Node                     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                                   // The cluster node name.
	Purge                    bool   `json:"purge,omitempty" url:"purge,omitempty,optional"`                                           // Remove container from all related configurations. For example, backup jobs, replication jobs or HA. Related ACLs and Firewall entries will *always* be removed.
	Vmid                     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`                                   // The (unique) ID of the VM.
}

// NodesLxcConfigVmRequest config - Get container configuration.
// Get container configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/config
//
type NodesLxcConfigVmRequest struct {
	Snapshot string `json:"snapshot,omitempty" url:"snapshot,omitempty,optional"`   // Fetch config values from given snapshot.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Current  bool   `json:"current,omitempty" url:"current,omitempty,optional"`     // Get current values (instead of pending values).
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesLxcConfigVmResponse config
// Get container configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/config
//
type NodesLxcConfigVmResponse struct {
	Digest       string   `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"`   // SHA1 digest of configuration file. This can be used to prevent concurrent modifications.
	Tags         string   `json:"tags,omitempty" url:"tags,omitempty,optional"`                 // Tags of the Container. This is only meta information.
	Timezone     string   `json:"timezone,omitempty" url:"timezone,omitempty,optional"`         // Time zone to use in the container. If option isn't set, then nothing will be done. Can be set to 'host' to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab
	Unprivileged bool     `json:"unprivileged,omitempty" url:"unprivileged,omitempty,optional"` // Makes the container run as unprivileged user. (Should not be modified manually.)
	Description  string   `json:"description,omitempty" url:"description,omitempty,optional"`   // Description for the Container. Shown in the web-interface CT's summary. This is saved as comment inside the configuration file.
	Net          []string `json:"net[n],omitempty" url:"net[n],omitempty,optional"`             // Specifies network interfaces for the container.
	Swap         int      `json:"swap,omitempty" url:"swap,omitempty,optional"`                 // Amount of SWAP for the VM in MB.
	Searchdomain string   `json:"searchdomain,omitempty" url:"searchdomain,omitempty,optional"` // Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Hostname     string   `json:"hostname,omitempty" url:"hostname,omitempty,optional"`         // Set a host name for the container.
	Lock         string   `json:"lock,omitempty" url:"lock,omitempty,optional"`                 // Lock/unlock the VM.
	Nameserver   string   `json:"nameserver,omitempty" url:"nameserver,omitempty,optional"`     // Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Arch         string   `json:"arch,omitempty" url:"arch,omitempty,optional"`                 // OS architecture type.
	Cpuunits     int      `json:"cpuunits,omitempty" url:"cpuunits,omitempty,optional"`         // CPU weight for a VM. Argument is used in the kernel fair scheduler. The larger the number is, the more CPU time this VM gets. Number is relative to the weights of all the other running VMs.  NOTE: You can disable fair-scheduler configuration by setting this to 0.
	Protection   bool     `json:"protection,omitempty" url:"protection,omitempty,optional"`     // Sets the protection flag of the container. This will prevent the CT or CT's disk remove/update operation.
	Features     string   `json:"features,omitempty" url:"features,omitempty,optional"`         // Allow containers access to advanced features.
	Onboot       bool     `json:"onboot,omitempty" url:"onboot,omitempty,optional"`             // Specifies whether a VM will be started during system bootup.
	Ostype       string   `json:"ostype,omitempty" url:"ostype,omitempty,optional"`             // OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/<ostype>.common.conf. Value 'unmanaged' can be used to skip and OS specific setup.
	Startup      string   `json:"startup,omitempty" url:"startup,omitempty,optional"`           // Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Tty          int      `json:"tty,omitempty" url:"tty,omitempty,optional"`                   // Specify the number of tty available to the container
	Unused       []string `json:"unused[n],omitempty" url:"unused[n],omitempty,optional"`       // Reference to unused volumes. This is used internally, and should not be modified manually.
	Cpulimit     int      `json:"cpulimit,omitempty" url:"cpulimit,omitempty,optional"`         // Limit of CPU usage.  NOTE: If the computer has 2 CPUs, it has a total of '2' CPU time. Value '0' indicates no CPU limit.
	Hookscript   string   `json:"hookscript,omitempty" url:"hookscript,omitempty,optional"`     // Script that will be exectued during various steps in the containers lifetime.
	Memory       int      `json:"memory,omitempty" url:"memory,omitempty,optional"`             // Amount of RAM for the VM in MB.
	Mp           []string `json:"mp[n],omitempty" url:"mp[n],omitempty,optional"`               // Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	Rootfs       string   `json:"rootfs,omitempty" url:"rootfs,omitempty,optional"`             // Use volume as container root.
	Cmode        string   `json:"cmode,omitempty" url:"cmode,omitempty,optional"`               // Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to 'console' it tries to attach to /dev/console instead. If you set cmode to 'shell', it simply invokes a shell inside the container (no login).
	Console      bool     `json:"console,omitempty" url:"console,omitempty,optional"`           // Attach a console device (/dev/console) to the container.
	Cores        int      `json:"cores,omitempty" url:"cores,omitempty,optional"`               // The number of cores assigned to the container. A container can use all available cores by default.
	Debug        bool     `json:"debug,omitempty" url:"debug,omitempty,optional"`               // Try to be more verbose. For now this only enables debug log-level on start.
	Lxc          []struct {
	} `json:"lxc,omitempty" url:"lxc,omitempty,optional"` // Array of lxc low-level configurations ([[key1, value1], [key2, value2] ...]).
	Template bool `json:"template,omitempty" url:"template,omitempty,optional"` // Enable/disable Template.
}

// NodesLxcConfigUpdateVmRequest config - Set container options.
// Set container options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/config
//
type NodesLxcConfigUpdateVmRequest struct {
	Cpuunits     int      `json:"cpuunits,omitempty" url:"cpuunits,omitempty,optional"`         // CPU weight for a VM. Argument is used in the kernel fair scheduler. The larger the number is, the more CPU time this VM gets. Number is relative to the weights of all the other running VMs.  NOTE: You can disable fair-scheduler configuration by setting this to 0.
	Hookscript   string   `json:"hookscript,omitempty" url:"hookscript,omitempty,optional"`     // Script that will be exectued during various steps in the containers lifetime.
	Memory       int      `json:"memory,omitempty" url:"memory,omitempty,optional"`             // Amount of RAM for the VM in MB.
	Unused       []string `json:"unused[n],omitempty" url:"unused[n],omitempty,optional"`       // Reference to unused volumes. This is used internally, and should not be modified manually.
	Timezone     string   `json:"timezone,omitempty" url:"timezone,omitempty,optional"`         // Time zone to use in the container. If option isn't set, then nothing will be done. Can be set to 'host' to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab
	Digest       string   `json:"digest,omitempty" url:"digest,omitempty,optional"`             // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Features     string   `json:"features,omitempty" url:"features,omitempty,optional"`         // Allow containers access to advanced features.
	Lock         string   `json:"lock,omitempty" url:"lock,omitempty,optional"`                 // Lock/unlock the VM.
	Mp           []string `json:"mp[n],omitempty" url:"mp[n],omitempty,optional"`               // Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	Onboot       bool     `json:"onboot,omitempty" url:"onboot,omitempty,optional"`             // Specifies whether a VM will be started during system bootup.
	Searchdomain string   `json:"searchdomain,omitempty" url:"searchdomain,omitempty,optional"` // Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Cmode        string   `json:"cmode,omitempty" url:"cmode,omitempty,optional"`               // Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to 'console' it tries to attach to /dev/console instead. If you set cmode to 'shell', it simply invokes a shell inside the container (no login).
	Cores        int      `json:"cores,omitempty" url:"cores,omitempty,optional"`               // The number of cores assigned to the container. A container can use all available cores by default.
	Cpulimit     int      `json:"cpulimit,omitempty" url:"cpulimit,omitempty,optional"`         // Limit of CPU usage.  NOTE: If the computer has 2 CPUs, it has a total of '2' CPU time. Value '0' indicates no CPU limit.
	Hostname     string   `json:"hostname,omitempty" url:"hostname,omitempty,optional"`         // Set a host name for the container.
	Ostype       string   `json:"ostype,omitempty" url:"ostype,omitempty,optional"`             // OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/<ostype>.common.conf. Value 'unmanaged' can be used to skip and OS specific setup.
	Protection   bool     `json:"protection,omitempty" url:"protection,omitempty,optional"`     // Sets the protection flag of the container. This will prevent the CT or CT's disk remove/update operation.
	Rootfs       string   `json:"rootfs,omitempty" url:"rootfs,omitempty,optional"`             // Use volume as container root.
	Unprivileged bool     `json:"unprivileged,omitempty" url:"unprivileged,omitempty,optional"` // Makes the container run as unprivileged user. (Should not be modified manually.)
	Delete       string   `json:"delete,omitempty" url:"delete,omitempty,optional"`             // A list of settings you want to delete.
	Revert       string   `json:"revert,omitempty" url:"revert,omitempty,optional"`             // Revert a pending change.
	Swap         int      `json:"swap,omitempty" url:"swap,omitempty,optional"`                 // Amount of SWAP for the VM in MB.
	Description  string   `json:"description,omitempty" url:"description,omitempty,optional"`   // Description for the Container. Shown in the web-interface CT's summary. This is saved as comment inside the configuration file.
	Arch         string   `json:"arch,omitempty" url:"arch,omitempty,optional"`                 // OS architecture type.
	Console      bool     `json:"console,omitempty" url:"console,omitempty,optional"`           // Attach a console device (/dev/console) to the container.
	Nameserver   string   `json:"nameserver,omitempty" url:"nameserver,omitempty,optional"`     // Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Net          []string `json:"net[n],omitempty" url:"net[n],omitempty,optional"`             // Specifies network interfaces for the container.
	Node         string   `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Vmid         string   `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`       // The (unique) ID of the VM.
	Debug        bool     `json:"debug,omitempty" url:"debug,omitempty,optional"`               // Try to be more verbose. For now this only enables debug log-level on start.
	Startup      string   `json:"startup,omitempty" url:"startup,omitempty,optional"`           // Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Tags         string   `json:"tags,omitempty" url:"tags,omitempty,optional"`                 // Tags of the Container. This is only meta information.
	Template     bool     `json:"template,omitempty" url:"template,omitempty,optional"`         // Enable/disable Template.
	Tty          int      `json:"tty,omitempty" url:"tty,omitempty,optional"`                   // Specify the number of tty available to the container
}

// NodesLxcStatusVmcmdidxRequest status - Directory index
// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/status
//
type NodesLxcStatusVmcmdidxRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcStatusVmcmdidxResponse status
// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/status
//
type NodesLxcStatusVmcmdidxResponse struct {
	Subdir string `json:"subdir,omitempty" url:"subdir,omitempty" validate:"nonzero"` //
}

// NodesLxcStatusCurrentVmStatusRequest current - Get virtual machine status.
// Get virtual machine status.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/status/current
//
type NodesLxcStatusCurrentVmStatusRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcStatusCurrentVmStatusResponse current
// Get virtual machine status.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/status/current
//
type NodesLxcStatusCurrentVmStatusResponse struct {
	Maxswap int      `json:"maxswap,omitempty" url:"maxswap,omitempty,optional"` // Maximum SWAP memory in bytes.
	Cpus    int      `json:"cpus,omitempty" url:"cpus,omitempty,optional"`       // Maximum usable CPUs.
	Ha      struct { //
	} `json:"ha,omitempty" url:"ha,omitempty" validate:"nonzero"` // HA manager service status.
	Lock    string `json:"lock,omitempty" url:"lock,omitempty,optional"`               // The current config lock, if any.
	Maxdisk int    `json:"maxdisk,omitempty" url:"maxdisk,omitempty,optional"`         // Root disk size in bytes.
	Uptime  int    `json:"uptime,omitempty" url:"uptime,omitempty,optional"`           // Uptime.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`     // The (unique) ID of the VM.
	Maxmem  int    `json:"maxmem,omitempty" url:"maxmem,omitempty,optional"`           // Maximum memory in bytes.
	Name    string `json:"name,omitempty" url:"name,omitempty,optional"`               // Container name.
	Status  string `json:"status,omitempty" url:"status,omitempty" validate:"nonzero"` // LXC Container status.
	Tags    string `json:"tags,omitempty" url:"tags,omitempty,optional"`               // The current configured tags, if any.
}

// NodesLxcStatusStartVmRequest start - Start the container.
// Start the container.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/status/start
//
type NodesLxcStatusStartVmRequest struct {
	Debug    bool   `json:"debug,omitempty" url:"debug,omitempty,optional"`         // If set, enables very verbose debug log-level on start.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Skiplock bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`   // Ignore locks - only root is allowed to use this option.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcStatusStopVmRequest stop - Stop the container. This will abruptly stop all processes running in the container.
// Stop the container. This will abruptly stop all processes running in the container.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/status/stop
//
type NodesLxcStatusStopVmRequest struct {
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Skiplock bool   `json:"skiplock,omitempty" url:"skiplock,omitempty,optional"`   // Ignore locks - only root is allowed to use this option.
}

// NodesLxcStatusShutdownVmRequest shutdown - Shutdown the container. This will trigger a clean shutdown of the container, see lxc-stop(1) for details.
// Shutdown the container. This will trigger a clean shutdown of the container, see lxc-stop(1) for details.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/status/shutdown
//
type NodesLxcStatusShutdownVmRequest struct {
	ForceStop bool   `json:"forceStop,omitempty" url:"forceStop,omitempty,optional"` // Make sure the Container stops.
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Timeout   int    `json:"timeout,omitempty" url:"timeout,omitempty,optional"`     // Wait maximal timeout seconds.
	Vmid      string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcStatusSuspendVmRequest suspend - Suspend the container. This is experimental.
// Suspend the container. This is experimental.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/status/suspend
//
type NodesLxcStatusSuspendVmRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcStatusResumeVmRequest resume - Resume the container.
// Resume the container.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/status/resume
//
type NodesLxcStatusResumeVmRequest struct {
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesLxcStatusRebootVmRequest reboot - Reboot the container by shutting it down, and starting it again. Applies pending changes.
// Reboot the container by shutting it down, and starting it again. Applies pending changes.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/status/reboot
//
type NodesLxcStatusRebootVmRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Timeout int    `json:"timeout,omitempty" url:"timeout,omitempty,optional"`     // Wait maximal timeout seconds for the shutdown.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcSnapshotListRequest snapshot - List all snapshots.
// List all snapshots.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/snapshot
//
type NodesLxcSnapshotListRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcSnapshotListResponse snapshot
// List all snapshots.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/snapshot
//
type NodesLxcSnapshotListResponse struct {
	Parent      string `json:"parent,omitempty" url:"parent,omitempty,optional"`                     // Parent snapshot identifier.
	Snaptime    int    `json:"snaptime,omitempty" url:"snaptime,omitempty,optional"`                 // Snapshot creation time
	Description string `json:"description,omitempty" url:"description,omitempty" validate:"nonzero"` // Snapshot description.
	Name        string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`               // Snapshot identifier. Value 'current' identifies the current VM.
}

// NodesLxcSnapshotRequest snapshot - Snapshot a container.
// Snapshot a container.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/snapshot
//
type NodesLxcSnapshotRequest struct {
	Description string `json:"description,omitempty" url:"description,omitempty,optional"`     // A textual description or comment.
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Snapname    string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
	Vmid        string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
}

// NodesLxcSnapshotCmdIdxRequest {snapname} -
//
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/snapshot/{snapname}
//
type NodesLxcSnapshotCmdIdxRequest struct {
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Snapname string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
}

// NodesLxcSnapshotDelsnapshotRequest {snapname} - Delete a LXC snapshot.
// Delete a LXC snapshot.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/snapshot/{snapname}
//
type NodesLxcSnapshotDelsnapshotRequest struct {
	Snapname string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Force    bool   `json:"force,omitempty" url:"force,omitempty,optional"`                 // For removal from config file, even if removing disk snapshots fails.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
}

// NodesLxcSnapshotRollbackRequest rollback - Rollback LXC state to specified snapshot.
// Rollback LXC state to specified snapshot.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/snapshot/{snapname}/rollback
//
type NodesLxcSnapshotRollbackRequest struct {
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Snapname string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
}

// NodesLxcSnapshotConfigGetSnapshotRequest config - Get snapshot configuration
// Get snapshot configuration
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/snapshot/{snapname}/config
//
type NodesLxcSnapshotConfigGetSnapshotRequest struct {
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Snapname string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
}

// NodesLxcSnapshotConfigUpdateSnapshotRequest config - Update snapshot metadata.
// Update snapshot metadata.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/snapshot/{snapname}/config
//
type NodesLxcSnapshotConfigUpdateSnapshotRequest struct {
	Description string `json:"description,omitempty" url:"description,omitempty,optional"`     // A textual description or comment.
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Snapname    string `json:"snapname,omitempty" url:"snapname,omitempty" validate:"nonzero"` // The name of the snapshot.
	Vmid        string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
}

// NodesLxcFirewallRulesGetRequest rules - List rules.
// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/rules
//
type NodesLxcFirewallRulesGetRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallRulesGetResponse rules
// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/rules
//
type NodesLxcFirewallRulesGetResponse struct {
	Pos int `json:"pos,omitempty" url:"pos,omitempty" validate:"nonzero"` //
}

// NodesLxcFirewallRulesCreateRuleRequest rules - Create new rule.
// Create new rule.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/rules
//
type NodesLxcFirewallRulesCreateRuleRequest struct {
	Dest     string `json:"dest,omitempty" url:"dest,omitempty,optional"`               // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Enable   int    `json:"enable,omitempty" url:"enable,omitempty,optional"`           // Flag to enable/disable a rule.
	Log      string `json:"log,omitempty" url:"log,omitempty,optional"`                 // Log level for firewall rule.
	Macro    string `json:"macro,omitempty" url:"macro,omitempty,optional"`             // Use predefined standard macro.
	Proto    string `json:"proto,omitempty" url:"proto,omitempty,optional"`             // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	IcmpType string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`     // Specify icmp-type. Only valid if proto equals 'icmp'.
	Iface    string `json:"iface,omitempty" url:"iface,omitempty,optional"`             // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Pos      int    `json:"pos,omitempty" url:"pos,omitempty,optional"`                 // Update rule at position <pos>.
	Type     string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     // Rule type.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`     // The (unique) ID of the VM.
	Comment  string `json:"comment,omitempty" url:"comment,omitempty,optional"`         // Descriptive comment.
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`           // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Dport    string `json:"dport,omitempty" url:"dport,omitempty,optional"`             // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Action   string `json:"action,omitempty" url:"action,omitempty" validate:"nonzero"` // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Source   string `json:"source,omitempty" url:"source,omitempty,optional"`           // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Sport    string `json:"sport,omitempty" url:"sport,omitempty,optional"`             // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
}

// NodesLxcFirewallRulesDeleteRuleRequest {pos} - Delete rule.
// Delete rule.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/rules/{pos}
//
type NodesLxcFirewallRulesDeleteRuleRequest struct {
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Pos    int    `json:"pos,omitempty" url:"pos,omitempty,optional"`             // Update rule at position <pos>.
	Vmid   string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallRulesGetRuleRequest {pos} - Get single rule data.
// Get single rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/rules/{pos}
//
type NodesLxcFirewallRulesGetRuleRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Pos  int    `json:"pos,omitempty" url:"pos,omitempty,optional"`             // Update rule at position <pos>.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallRulesGetRuleResponse {pos}
// Get single rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/rules/{pos}
//
type NodesLxcFirewallRulesGetRuleResponse struct {
	Enable    int    `json:"enable,omitempty" url:"enable,omitempty,optional"`           //
	Macro     string `json:"macro,omitempty" url:"macro,omitempty,optional"`             //
	Sport     string `json:"sport,omitempty" url:"sport,omitempty,optional"`             //
	Comment   string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Pos       int    `json:"pos,omitempty" url:"pos,omitempty" validate:"nonzero"`       //
	Ipversion int    `json:"ipversion,omitempty" url:"ipversion,omitempty,optional"`     //
	Iface     string `json:"iface,omitempty" url:"iface,omitempty,optional"`             //
	Log       string `json:"log,omitempty" url:"log,omitempty,optional"`                 // Log level for firewall rule
	Source    string `json:"source,omitempty" url:"source,omitempty,optional"`           //
	Dest      string `json:"dest,omitempty" url:"dest,omitempty,optional"`               //
	Dport     string `json:"dport,omitempty" url:"dport,omitempty,optional"`             //
	IcmpType  string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`     //
	Proto     string `json:"proto,omitempty" url:"proto,omitempty,optional"`             //
	Type      string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     //
	Action    string `json:"action,omitempty" url:"action,omitempty" validate:"nonzero"` //
}

// NodesLxcFirewallRulesUpdateRuleRequest {pos} - Modify rule data.
// Modify rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/rules/{pos}
//
type NodesLxcFirewallRulesUpdateRuleRequest struct {
	Delete   string `json:"delete,omitempty" url:"delete,omitempty,optional"`       // A list of settings you want to delete.
	Dest     string `json:"dest,omitempty" url:"dest,omitempty,optional"`           // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dport    string `json:"dport,omitempty" url:"dport,omitempty,optional"`         // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	IcmpType string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"` // Specify icmp-type. Only valid if proto equals 'icmp'.
	Proto    string `json:"proto,omitempty" url:"proto,omitempty,optional"`         // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Sport    string `json:"sport,omitempty" url:"sport,omitempty,optional"`         // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Action   string `json:"action,omitempty" url:"action,omitempty,optional"`       // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Type     string `json:"type,omitempty" url:"type,omitempty,optional"`           // Rule type.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Comment  string `json:"comment,omitempty" url:"comment,omitempty,optional"`     // Descriptive comment.
	Iface    string `json:"iface,omitempty" url:"iface,omitempty,optional"`         // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Macro    string `json:"macro,omitempty" url:"macro,omitempty,optional"`         // Use predefined standard macro.
	Moveto   int    `json:"moveto,omitempty" url:"moveto,omitempty,optional"`       // Move rule to new position <moveto>. Other arguments are ignored.
	Enable   int    `json:"enable,omitempty" url:"enable,omitempty,optional"`       // Flag to enable/disable a rule.
	Log      string `json:"log,omitempty" url:"log,omitempty,optional"`             // Log level for firewall rule.
	Pos      int    `json:"pos,omitempty" url:"pos,omitempty,optional"`             // Update rule at position <pos>.
	Source   string `json:"source,omitempty" url:"source,omitempty,optional"`       // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
}

// NodesLxcFirewallAliasesGetRequest aliases - List aliases
// List aliases
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/aliases
//
type NodesLxcFirewallAliasesGetRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallAliasesGetResponse aliases
// List aliases
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/aliases
//
type NodesLxcFirewallAliasesGetResponse struct {
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"`     //
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`     //
}

// NodesLxcFirewallAliasesCreateAliasRequest aliases - Create IP or Network Alias.
// Create IP or Network Alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/aliases
//
type NodesLxcFirewallAliasesCreateAliasRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
}

// NodesLxcFirewallAliasesUpdateAliasRequest {name} - Update IP or Network alias.
// Update IP or Network alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/aliases/{name}
//
type NodesLxcFirewallAliasesUpdateAliasRequest struct {
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Rename  string `json:"rename,omitempty" url:"rename,omitempty,optional"`       // Rename an existing alias.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallAliasesRemoveAliasRequest {name} - Remove IP or Network alias.
// Remove IP or Network alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/aliases/{name}
//
type NodesLxcFirewallAliasesRemoveAliasRequest struct {
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid   string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
}

// NodesLxcFirewallAliasesReadAliasRequest {name} - Read alias.
// Read alias.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/aliases/{name}
//
type NodesLxcFirewallAliasesReadAliasRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Alias name.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallIpsetIndexRequest ipset - List IPSets
// List IPSets
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/ipset
//
type NodesLxcFirewallIpsetIndexRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallIpsetIndexResponse ipset
// List IPSets
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/ipset
//
type NodesLxcFirewallIpsetIndexResponse struct {
	Digest  string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`     // IP set name.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
}

// NodesLxcFirewallIpsetCreateRequest ipset - Create new IPSet
// Create new IPSet
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/ipset
//
type NodesLxcFirewallIpsetCreateRequest struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Rename  string `json:"rename,omitempty" url:"rename,omitempty,optional"`       // Rename an existing IPSet. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing IPSet.
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallIpsetDeleteRequest {name} - Delete IPSet
// Delete IPSet
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}
//
type NodesLxcFirewallIpsetDeleteRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallIpsetGetRequest {name} - List IPSet content
// List IPSet content
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}
//
type NodesLxcFirewallIpsetGetRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallIpsetGetResponse {name}
// List IPSet content
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}
//
type NodesLxcFirewallIpsetGetResponse struct {
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"`     //
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Nomatch bool   `json:"nomatch,omitempty" url:"nomatch,omitempty,optional"`         //
}

// NodesLxcFirewallIpsetCreateIpRequest {name} - Add IP or Network to IPSet.
// Add IP or Network to IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}
//
type NodesLxcFirewallIpsetCreateIpRequest struct {
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Nomatch bool   `json:"nomatch,omitempty" url:"nomatch,omitempty,optional"`     //
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallIpsetUpdateIpRequest {cidr} - Update IP or Network settings
// Update IP or Network settings
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}/{cidr}
//
type NodesLxcFirewallIpsetUpdateIpRequest struct {
	Cidr    string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Nomatch bool   `json:"nomatch,omitempty" url:"nomatch,omitempty,optional"`     //
	Vmid    string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallIpsetRemoveIpRequest {cidr} - Remove IP or Network from IPSet.
// Remove IP or Network from IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}/{cidr}
//
type NodesLxcFirewallIpsetRemoveIpRequest struct {
	Cidr   string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid   string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallIpsetReadIpRequest {cidr} - Read IP or Network settings from IPSet.
// Read IP or Network settings from IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/ipset/{name}/{cidr}
//
type NodesLxcFirewallIpsetReadIpRequest struct {
	Cidr string `json:"cidr,omitempty" url:"cidr,omitempty" validate:"nonzero"` // Network/IP specification in CIDR format.
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // IP set name.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallOptionsGetRequest options - Get VM firewall options.
// Get VM firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/options
//
type NodesLxcFirewallOptionsGetRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcFirewallOptionsGetResponse options
// Get VM firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/options
//
type NodesLxcFirewallOptionsGetResponse struct {
	Dhcp        bool   `json:"dhcp,omitempty" url:"dhcp,omitempty,optional"`                   // Enable DHCP.
	Ipfilter    bool   `json:"ipfilter,omitempty" url:"ipfilter,omitempty,optional"`           // Enable default IP filters. This is equivalent to adding an empty ipfilter-net<id> ipset for every interface. Such ipsets implicitly contain sane default restrictions such as restricting IPv6 link local addresses to the one derived from the interface's MAC address. For containers the configured IP addresses will be implicitly added.
	Ndp         bool   `json:"ndp,omitempty" url:"ndp,omitempty,optional"`                     // Enable NDP (Neighbor Discovery Protocol).
	PolicyIn    string `json:"policy_in,omitempty" url:"policy_in,omitempty,optional"`         // Input policy.
	Radv        bool   `json:"radv,omitempty" url:"radv,omitempty,optional"`                   // Allow sending Router Advertisement.
	Enable      bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`               // Enable/disable firewall rules.
	LogLevelIn  string `json:"log_level_in,omitempty" url:"log_level_in,omitempty,optional"`   // Log level for incoming traffic.
	LogLevelOut string `json:"log_level_out,omitempty" url:"log_level_out,omitempty,optional"` // Log level for outgoing traffic.
	Macfilter   bool   `json:"macfilter,omitempty" url:"macfilter,omitempty,optional"`         // Enable/disable MAC address filter.
	PolicyOut   string `json:"policy_out,omitempty" url:"policy_out,omitempty,optional"`       // Output policy.
}

// NodesLxcFirewallOptionsSetRequest options - Set Firewall options.
// Set Firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/options
//
type NodesLxcFirewallOptionsSetRequest struct {
	Vmid        string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Digest      string `json:"digest,omitempty" url:"digest,omitempty,optional"`               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Enable      bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`               // Enable/disable firewall rules.
	LogLevelOut string `json:"log_level_out,omitempty" url:"log_level_out,omitempty,optional"` // Log level for outgoing traffic.
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	PolicyIn    string `json:"policy_in,omitempty" url:"policy_in,omitempty,optional"`         // Input policy.
	Ipfilter    bool   `json:"ipfilter,omitempty" url:"ipfilter,omitempty,optional"`           // Enable default IP filters. This is equivalent to adding an empty ipfilter-net<id> ipset for every interface. Such ipsets implicitly contain sane default restrictions such as restricting IPv6 link local addresses to the one derived from the interface's MAC address. For containers the configured IP addresses will be implicitly added.
	LogLevelIn  string `json:"log_level_in,omitempty" url:"log_level_in,omitempty,optional"`   // Log level for incoming traffic.
	Ndp         bool   `json:"ndp,omitempty" url:"ndp,omitempty,optional"`                     // Enable NDP (Neighbor Discovery Protocol).
	PolicyOut   string `json:"policy_out,omitempty" url:"policy_out,omitempty,optional"`       // Output policy.
	Delete      string `json:"delete,omitempty" url:"delete,omitempty,optional"`               // A list of settings you want to delete.
	Dhcp        bool   `json:"dhcp,omitempty" url:"dhcp,omitempty,optional"`                   // Enable DHCP.
	Macfilter   bool   `json:"macfilter,omitempty" url:"macfilter,omitempty,optional"`         // Enable/disable MAC address filter.
	Radv        bool   `json:"radv,omitempty" url:"radv,omitempty,optional"`                   // Allow sending Router Advertisement.
}

// NodesLxcFirewallLogRequest log - Read firewall log
// Read firewall log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/log
//
type NodesLxcFirewallLogRequest struct {
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Start int    `json:"start,omitempty" url:"start,omitempty,optional"`         //
	Vmid  string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Limit int    `json:"limit,omitempty" url:"limit,omitempty,optional"`         //
}

// NodesLxcFirewallLogResponse log
// Read firewall log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/log
//
type NodesLxcFirewallLogResponse struct {
	N int    `json:"n,omitempty" url:"n,omitempty" validate:"nonzero"` // Line number
	T string `json:"t,omitempty" url:"t,omitempty" validate:"nonzero"` // Line text
}

// NodesLxcFirewallRefsRequest refs - Lists possible IPSet/Alias reference which are allowed in source/dest properties.
// Lists possible IPSet/Alias reference which are allowed in source/dest properties.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/refs
//
type NodesLxcFirewallRefsRequest struct {
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Type string `json:"type,omitempty" url:"type,omitempty,optional"`           // Only list references of specified type.
}

// NodesLxcFirewallRefsResponse refs
// Lists possible IPSet/Alias reference which are allowed in source/dest properties.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/firewall/refs
//
type NodesLxcFirewallRefsResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`     //
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` //
	Type    string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"` //
}

// NodesLxcRrdRequest rrd - Read VM RRD statistics (returns PNG)
// Read VM RRD statistics (returns PNG)
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/rrd
//
type NodesLxcRrdRequest struct {
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Timeframe string `json:"timeframe,omitempty" url:"timeframe,omitempty" validate:"nonzero"` // Specify the time frame you are interested in.
	Vmid      string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`           // The (unique) ID of the VM.
	Cf        string `json:"cf,omitempty" url:"cf,omitempty,optional"`                         // The RRD consolidation function
	Ds        string `json:"ds,omitempty" url:"ds,omitempty" validate:"nonzero"`               // The list of datasources you want to display.
}

// NodesLxcRrdResponse rrd
// Read VM RRD statistics (returns PNG)
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/rrd
//
type NodesLxcRrdResponse struct {
	Filename string `json:"filename,omitempty" url:"filename,omitempty" validate:"nonzero"` //
}

// NodesLxcRrddataRequest rrddata - Read VM RRD statistics
// Read VM RRD statistics
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/rrddata
//
type NodesLxcRrddataRequest struct {
	Timeframe string `json:"timeframe,omitempty" url:"timeframe,omitempty" validate:"nonzero"` // Specify the time frame you are interested in.
	Vmid      string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`           // The (unique) ID of the VM.
	Cf        string `json:"cf,omitempty" url:"cf,omitempty,optional"`                         // The RRD consolidation function
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
}

// NodesLxcVncproxyRequest vncproxy - Creates a TCP VNC proxy connections.
// Creates a TCP VNC proxy connections.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/vncproxy
//
type NodesLxcVncproxyRequest struct {
	Vmid      string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Websocket bool   `json:"websocket,omitempty" url:"websocket,omitempty,optional"` // use websocket instead of standard VNC.
	Width     int    `json:"width,omitempty" url:"width,omitempty,optional"`         // sets the width of the console in pixels.
	Height    int    `json:"height,omitempty" url:"height,omitempty,optional"`       // sets the height of the console in pixels.
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesLxcVncproxyResponse vncproxy
// Creates a TCP VNC proxy connections.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/vncproxy
//
type NodesLxcVncproxyResponse struct {
	Port   int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`     //
	Ticket string `json:"ticket,omitempty" url:"ticket,omitempty" validate:"nonzero"` //
	Upid   string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"`     //
	User   string `json:"user,omitempty" url:"user,omitempty" validate:"nonzero"`     //
	Cert   string `json:"cert,omitempty" url:"cert,omitempty" validate:"nonzero"`     //
}

// NodesLxcTermproxyRequest termproxy - Creates a TCP proxy connection.
// Creates a TCP proxy connection.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/termproxy
//
type NodesLxcTermproxyRequest struct {
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesLxcTermproxyResponse termproxy
// Creates a TCP proxy connection.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/termproxy
//
type NodesLxcTermproxyResponse struct {
	Port   int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`     //
	Ticket string `json:"ticket,omitempty" url:"ticket,omitempty" validate:"nonzero"` //
	Upid   string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"`     //
	User   string `json:"user,omitempty" url:"user,omitempty" validate:"nonzero"`     //
}

// NodesLxcVncwebsocketRequest vncwebsocket - Opens a weksocket for VNC traffic.
// Opens a weksocket for VNC traffic.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/vncwebsocket
//
type NodesLxcVncwebsocketRequest struct {
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Port      int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`           // Port number returned by previous vncproxy call.
	Vmid      string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`           // The (unique) ID of the VM.
	Vncticket string `json:"vncticket,omitempty" url:"vncticket,omitempty" validate:"nonzero"` // Ticket from previous call to vncproxy.
}

// NodesLxcVncwebsocketResponse vncwebsocket
// Opens a weksocket for VNC traffic.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/vncwebsocket
//
type NodesLxcVncwebsocketResponse struct {
	Port string `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"` //
}

// NodesLxcSpiceproxyRequest spiceproxy - Returns a SPICE configuration to connect to the CT.
// Returns a SPICE configuration to connect to the CT.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/spiceproxy
//
type NodesLxcSpiceproxyRequest struct {
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Proxy string `json:"proxy,omitempty" url:"proxy,omitempty,optional"`         // SPICE proxy server. This can be used by the client to specify the proxy server. All nodes in a cluster runs 'spiceproxy', so it is up to the client to choose one. By default, we return the node where the VM is currently running. As reasonable setting is to use same node you use to connect to the API (This is window.location.hostname for the JS GUI).
	Vmid  string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcSpiceproxyResponse spiceproxy
// Returns a SPICE configuration to connect to the CT.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/spiceproxy
//
type NodesLxcSpiceproxyResponse struct {
	Host     string `json:"host,omitempty" url:"host,omitempty" validate:"nonzero"`         //
	Password string `json:"password,omitempty" url:"password,omitempty" validate:"nonzero"` //
	Proxy    string `json:"proxy,omitempty" url:"proxy,omitempty" validate:"nonzero"`       //
	TlsPort  int    `json:"tls-port,omitempty" url:"tls-port,omitempty" validate:"nonzero"` //
	Type     string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`         //
}

// NodesLxcMigrateVmRequest migrate - Migrate the container to another node. Creates a new migration task.
// Migrate the container to another node. Creates a new migration task.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/migrate
//
type NodesLxcMigrateVmRequest struct {
	Bwlimit       int    `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`               // Override I/O bandwidth limit (in KiB/s).
	Node          string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Online        bool   `json:"online,omitempty" url:"online,omitempty,optional"`                 // Use online/live migration.
	Restart       bool   `json:"restart,omitempty" url:"restart,omitempty,optional"`               // Use restart migration
	Target        string `json:"target,omitempty" url:"target,omitempty" validate:"nonzero"`       // Target node.
	TargetStorage string `json:"target-storage,omitempty" url:"target-storage,omitempty,optional"` // Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	Timeout       int    `json:"timeout,omitempty" url:"timeout,omitempty,optional"`               // Timeout in seconds for shutdown for restart migration
	Vmid          string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`           // The (unique) ID of the VM.
}

// NodesLxcFeatureVmRequest feature - Check if feature for virtual machine is available.
// Check if feature for virtual machine is available.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/feature
//
type NodesLxcFeatureVmRequest struct {
	Feature  string `json:"feature,omitempty" url:"feature,omitempty" validate:"nonzero"` // Feature to check.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Snapname string `json:"snapname,omitempty" url:"snapname,omitempty,optional"`         // The name of the snapshot.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`       // The (unique) ID of the VM.
}

// NodesLxcFeatureVmResponse feature
// Check if feature for virtual machine is available.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/feature
//
type NodesLxcFeatureVmResponse struct {
	HasFeature bool `json:"hasFeature,omitempty" url:"hasFeature,omitempty" validate:"nonzero"` //
}

// NodesLxcTemplateRequest template - Create a Template.
// Create a Template.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/template
//
type NodesLxcTemplateRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcCloneVmRequest clone - Create a container clone/copy
// Create a container clone/copy
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/clone
//
type NodesLxcCloneVmRequest struct {
	Storage     string `json:"storage,omitempty" url:"storage,omitempty,optional"`         // Target storage for full clone.
	Target      string `json:"target,omitempty" url:"target,omitempty,optional"`           // Target node. Only allowed if the original VM is on shared storage.
	Bwlimit     int    `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`         // Override I/O bandwidth limit (in KiB/s).
	Description string `json:"description,omitempty" url:"description,omitempty,optional"` // Description for the new CT.
	Newid       int    `json:"newid,omitempty" url:"newid,omitempty" validate:"nonzero"`   // VMID for the clone.
	Pool        string `json:"pool,omitempty" url:"pool,omitempty,optional"`               // Add the new CT to the specified pool.
	Vmid        string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`     // The (unique) ID of the VM.
	Full        bool   `json:"full,omitempty" url:"full,omitempty,optional"`               // Create a full copy of all disks. This is always done when you clone a normal CT. For CT templates, we try to create a linked clone by default.
	Hostname    string `json:"hostname,omitempty" url:"hostname,omitempty,optional"`       // Set a hostname for the new CT.
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Snapname    string `json:"snapname,omitempty" url:"snapname,omitempty,optional"`       // The name of the snapshot.
}

// NodesLxcResizeVmRequest resize - Resize a container mount point.
// Resize a container mount point.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/resize
//
type NodesLxcResizeVmRequest struct {
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Disk   string `json:"disk,omitempty" url:"disk,omitempty" validate:"nonzero"` // The disk you want to resize.
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Size   string `json:"size,omitempty" url:"size,omitempty" validate:"nonzero"` // The new size. With the '+' sign the value is added to the actual size of the volume and without it, the value is taken as an absolute one. Shrinking disk size is not supported.
	Vmid   string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcMoveVolumeRequest move_volume - Move a rootfs-/mp-volume to a different storage or to a different container.
// Move a rootfs-/mp-volume to a different storage or to a different container.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/move_volume
//
type NodesLxcMoveVolumeRequest struct {
	Vmid         string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // The (unique) ID of the VM.
	Volume       string `json:"volume,omitempty" url:"volume,omitempty" validate:"nonzero"`     // Volume which will be moved.
	Delete       bool   `json:"delete,omitempty" url:"delete,omitempty,optional"`               // Delete the original volume after successful copy. By default the original is kept as an unused volume entry.
	Storage      string `json:"storage,omitempty" url:"storage,omitempty,optional"`             // Target Storage.
	TargetDigest string `json:"target-digest,omitempty" url:"target-digest,omitempty,optional"` // Prevent changes if current configuration file of the target " . 		    "container has a different SHA1 digest. This can be used to prevent " . 		    "concurrent modifications.
	TargetVmid   int    `json:"target-vmid,omitempty" url:"target-vmid,omitempty,optional"`     // The (unique) ID of the VM.
	TargetVolume string `json:"target-volume,omitempty" url:"target-volume,omitempty,optional"` // The config key the volume will be moved to. Default is the source volume key.
	Bwlimit      int    `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`             // Override I/O bandwidth limit (in KiB/s).
	Digest       string `json:"digest,omitempty" url:"digest,omitempty,optional"`               // Prevent changes if current configuration file has different SHA1 " . 		    "digest. This can be used to prevent concurrent modifications.
	Node         string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
}

// NodesLxcPendingVmRequest pending - Get container configuration, including pending changes.
// Get container configuration, including pending changes.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/pending
//
type NodesLxcPendingVmRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vmid string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"` // The (unique) ID of the VM.
}

// NodesLxcPendingVmResponse pending
// Get container configuration, including pending changes.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc/{vmid}/pending
//
type NodesLxcPendingVmResponse struct {
	Value   string `json:"value,omitempty" url:"value,omitempty,optional"`       // Current value.
	Delete  int    `json:"delete,omitempty" url:"delete,omitempty,optional"`     // Indicates a pending delete request if present and not 0.
	Key     string `json:"key,omitempty" url:"key,omitempty" validate:"nonzero"` // Configuration option name.
	Pending string `json:"pending,omitempty" url:"pending,omitempty,optional"`   // Pending value.
}

// NodesCephOsdCreateosdRequest osd - Create OSD
// Create OSD
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/osd
//
type NodesCephOsdCreateosdRequest struct {
	WalDev           string `json:"wal_dev,omitempty" url:"wal_dev,omitempty,optional"`                       // Block device name for block.wal.
	WalDevSize       int    `json:"wal_dev_size,omitempty" url:"wal_dev_size,omitempty,optional"`             // Size in GiB for block.wal.
	CrushDeviceClass string `json:"crush-device-class,omitempty" url:"crush-device-class,omitempty,optional"` // Set the device class of the OSD in crush.
	DbDev            string `json:"db_dev,omitempty" url:"db_dev,omitempty,optional"`                         // Block device name for block.db.
	DbDevSize        int    `json:"db_dev_size,omitempty" url:"db_dev_size,omitempty,optional"`               // Size in GiB for block.db.
	Dev              string `json:"dev,omitempty" url:"dev,omitempty" validate:"nonzero"`                     // Block device name.
	Encrypted        bool   `json:"encrypted,omitempty" url:"encrypted,omitempty,optional"`                   // Enables encryption of the OSD.
	Node             string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                   // The cluster node name.
}

// NodesCephOsdDestroyosdRequest {osdid} - Destroy OSD
// Destroy OSD
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/osd/{osdid}
//
type NodesCephOsdDestroyosdRequest struct {
	Osdid   int    `json:"osdid,omitempty" url:"osdid,omitempty" validate:"nonzero"` // OSD ID
	Cleanup bool   `json:"cleanup,omitempty" url:"cleanup,omitempty,optional"`       // If set, we remove partition table entries.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
}

// NodesCephOsdInRequest in - ceph osd in
// ceph osd in
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/osd/{osdid}/in
//
type NodesCephOsdInRequest struct {
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
	Osdid int    `json:"osdid,omitempty" url:"osdid,omitempty" validate:"nonzero"` // OSD ID
}

// NodesCephOsdOutRequest out - ceph osd out
// ceph osd out
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/osd/{osdid}/out
//
type NodesCephOsdOutRequest struct {
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
	Osdid int    `json:"osdid,omitempty" url:"osdid,omitempty" validate:"nonzero"` // OSD ID
}

// NodesCephOsdScrubRequest scrub - Instruct the OSD to scrub.
// Instruct the OSD to scrub.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/osd/{osdid}/scrub
//
type NodesCephOsdScrubRequest struct {
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
	Osdid int    `json:"osdid,omitempty" url:"osdid,omitempty" validate:"nonzero"` // OSD ID
	Deep  bool   `json:"deep,omitempty" url:"deep,omitempty,optional"`             // If set, instructs a deep scrub instead of a normal one.
}

// NodesCephMdsDestroymdsRequest {name} - Destroy Ceph Metadata Server
// Destroy Ceph Metadata Server
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/mds/{name}
//
type NodesCephMdsDestroymdsRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // The name (ID) of the mds
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCephMdsCreatemdsRequest {name} - Create Ceph Metadata Server (MDS)
// Create Ceph Metadata Server (MDS)
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/mds/{name}
//
type NodesCephMdsCreatemdsRequest struct {
	Hotstandby bool   `json:"hotstandby,omitempty" url:"hotstandby,omitempty,optional"` // Determines whether a ceph-mds daemon should poll and replay the log of an active MDS. Faster switch on MDS failure, but needs more idle resources.
	Name       string `json:"name,omitempty" url:"name,omitempty,optional"`             // The ID for the mds, when omitted the same as the nodename
	Node       string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
}

// NodesCephMgrDestroymgrRequest {id} - Destroy Ceph Manager.
// Destroy Ceph Manager.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/mgr/{id}
//
type NodesCephMgrDestroymgrRequest struct {
	Id   string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`     // The ID of the manager
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCephMgrCreatemgrRequest {id} - Create Ceph Manager
// Create Ceph Manager
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/mgr/{id}
//
type NodesCephMgrCreatemgrRequest struct {
	Id   string `json:"id,omitempty" url:"id,omitempty,optional"`               // The ID for the manager, when omitted the same as the nodename
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCephMonListmonRequest mon - Get Ceph monitor list.
// Get Ceph monitor list.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/mon
//
type NodesCephMonListmonRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCephMonListmonResponse mon
// Get Ceph monitor list.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/mon
//
type NodesCephMonListmonResponse struct {
	Addr string `json:"addr,omitempty" url:"addr,omitempty,optional"`           //
	Host string `json:"host,omitempty" url:"host,omitempty,optional"`           //
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` //
}

// NodesCephMonDestroymonRequest {monid} - Destroy Ceph Monitor and Manager.
// Destroy Ceph Monitor and Manager.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/mon/{monid}
//
type NodesCephMonDestroymonRequest struct {
	Monid string `json:"monid,omitempty" url:"monid,omitempty" validate:"nonzero"` // Monitor ID
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
}

// NodesCephMonCreatemonRequest {monid} - Create Ceph Monitor and Manager
// Create Ceph Monitor and Manager
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/mon/{monid}
//
type NodesCephMonCreatemonRequest struct {
	Monid      string `json:"monid,omitempty" url:"monid,omitempty,optional"`             // The ID for the monitor, when omitted the same as the nodename
	Node       string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	MonAddress string `json:"mon-address,omitempty" url:"mon-address,omitempty,optional"` // Overwrites autodetected monitor IP address(es). Must be in the public network(s) of Ceph.
}

// NodesCephFsCreatefsRequest {name} - Create a Ceph filesystem
// Create a Ceph filesystem
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/fs/{name}
//
type NodesCephFsCreatefsRequest struct {
	Name       string `json:"name,omitempty" url:"name,omitempty,optional"`               // The ceph filesystem name.
	Node       string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	PgNum      int    `json:"pg_num,omitempty" url:"pg_num,omitempty,optional"`           // Number of placement groups for the backing data pool. The metadata pool will use a quarter of this.
	AddStorage bool   `json:"add-storage,omitempty" url:"add-storage,omitempty,optional"` // Configure the created CephFS as storage for this cluster.
}

// NodesCephPoolsLspoolsRequest pools - List all pools.
// List all pools.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/pools
//
type NodesCephPoolsLspoolsRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCephPoolsLspoolsResponse pools
// List all pools.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/pools
//
type NodesCephPoolsLspoolsResponse struct {
	TargetSizeRatio int      `json:"target_size_ratio,omitempty" url:"target_size_ratio,omitempty,optional"`       //
	PercentUsed     int      `json:"percent_used,omitempty" url:"percent_used,omitempty" validate:"nonzero"`       //
	PgAutoscaleMode string   `json:"pg_autoscale_mode,omitempty" url:"pg_autoscale_mode,omitempty,optional"`       //
	Pool            int      `json:"pool,omitempty" url:"pool,omitempty" validate:"nonzero"`                       //
	PoolName        string   `json:"pool_name,omitempty" url:"pool_name,omitempty" validate:"nonzero"`             //
	Size            int      `json:"size,omitempty" url:"size,omitempty" validate:"nonzero"`                       //
	Type            string   `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`                       //
	MinSize         int      `json:"min_size,omitempty" url:"min_size,omitempty" validate:"nonzero"`               //
	PgNumFinal      int      `json:"pg_num_final,omitempty" url:"pg_num_final,omitempty,optional"`                 //
	PgNumMin        int      `json:"pg_num_min,omitempty" url:"pg_num_min,omitempty,optional"`                     //
	TargetSize      int      `json:"target_size,omitempty" url:"target_size,omitempty,optional"`                   //
	CrushRule       int      `json:"crush_rule,omitempty" url:"crush_rule,omitempty" validate:"nonzero"`           //
	CrushRuleName   string   `json:"crush_rule_name,omitempty" url:"crush_rule_name,omitempty" validate:"nonzero"` //
	PgNum           int      `json:"pg_num,omitempty" url:"pg_num,omitempty" validate:"nonzero"`                   //
	AutoscaleStatus struct { //
	} `json:"autoscale_status,omitempty" url:"autoscale_status,omitempty,optional"` //
	BytesUsed int `json:"bytes_used,omitempty" url:"bytes_used,omitempty" validate:"nonzero"` //
}

// NodesCephPoolsCreatepoolRequest pools - Create Ceph pool
// Create Ceph pool
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/pools
//
type NodesCephPoolsCreatepoolRequest struct {
	Application     string `json:"application,omitempty" url:"application,omitempty,optional"`             // The application of the pool.
	CrushRule       string `json:"crush_rule,omitempty" url:"crush_rule,omitempty,optional"`               // The rule to use for mapping object placement in the cluster.
	MinSize         int    `json:"min_size,omitempty" url:"min_size,omitempty,optional"`                   // Minimum number of replicas per object
	Name            string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`                 // The name of the pool. It must be unique.
	PgNumMin        int    `json:"pg_num_min,omitempty" url:"pg_num_min,omitempty,optional"`               // Minimal number of placement groups.
	Size            int    `json:"size,omitempty" url:"size,omitempty,optional"`                           // Number of replicas per object
	TargetSize      string `json:"target_size,omitempty" url:"target_size,omitempty,optional"`             // The estimated target size of the pool for the PG autoscaler.
	AddStorages     bool   `json:"add_storages,omitempty" url:"add_storages,omitempty,optional"`           // Configure VM and CT storage using the new pool.
	Node            string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                 // The cluster node name.
	PgAutoscaleMode string `json:"pg_autoscale_mode,omitempty" url:"pg_autoscale_mode,omitempty,optional"` // The automatic PG scaling mode of the pool.
	PgNum           int    `json:"pg_num,omitempty" url:"pg_num,omitempty,optional"`                       // Number of placement groups.
	TargetSizeRatio int    `json:"target_size_ratio,omitempty" url:"target_size_ratio,omitempty,optional"` // The estimated target ratio of the pool for the PG autoscaler.
	ErasureCoding   string `json:"erasure-coding,omitempty" url:"erasure-coding,omitempty,optional"`       // Create an erasure coded pool for RBD with an accompaning replicated pool for metadata storage. With EC, the common ceph options 'size', 'min_size' and 'crush_rule' parameters will be applied to the metadata pool.
}

// NodesCephPoolsDestroypoolRequest {name} - Destroy pool
// Destroy pool
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/pools/{name}
//
type NodesCephPoolsDestroypoolRequest struct {
	Force           bool   `json:"force,omitempty" url:"force,omitempty,optional"`                       // If true, destroys pool even if in use
	Name            string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`               // The name of the pool. It must be unique.
	Node            string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`               // The cluster node name.
	RemoveEcprofile bool   `json:"remove_ecprofile,omitempty" url:"remove_ecprofile,omitempty,optional"` // Remove the erasure code profile. Defaults to true, if applicable.
	RemoveStorages  bool   `json:"remove_storages,omitempty" url:"remove_storages,omitempty,optional"`   // Remove all pveceph-managed storages configured for this pool
}

// NodesCephPoolsGetpoolRequest {name} - List pool settings.
// List pool settings.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/pools/{name}
//
type NodesCephPoolsGetpoolRequest struct {
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // The name of the pool. It must be unique.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Verbose bool   `json:"verbose,omitempty" url:"verbose,omitempty,optional"`     // If enabled, will display additional data(eg. statistics).
}

// NodesCephPoolsGetpoolResponse {name}
// List pool settings.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/pools/{name}
//
type NodesCephPoolsGetpoolResponse struct {
	PgNumMin        int      `json:"pg_num_min,omitempty" url:"pg_num_min,omitempty,optional"`               // Minimal number of placement groups.
	Size            int      `json:"size,omitempty" url:"size,omitempty,optional"`                           // Number of replicas per object
	Id              int      `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`                     //
	MinSize         int      `json:"min_size,omitempty" url:"min_size,omitempty,optional"`                   // Minimum number of replicas per object
	Nosizechange    bool     `json:"nosizechange,omitempty" url:"nosizechange,omitempty" validate:"nonzero"` //
	PgAutoscaleMode string   `json:"pg_autoscale_mode,omitempty" url:"pg_autoscale_mode,omitempty,optional"` // The automatic PG scaling mode of the pool.
	PgNum           int      `json:"pg_num,omitempty" url:"pg_num,omitempty,optional"`                       // Number of placement groups.
	AutoscaleStatus struct { //
	} `json:"autoscale_status,omitempty" url:"autoscale_status,omitempty,optional"` //
	CrushRule            string `json:"crush_rule,omitempty" url:"crush_rule,omitempty,optional"`                                   // The rule to use for mapping object placement in the cluster.
	Hashpspool           bool   `json:"hashpspool,omitempty" url:"hashpspool,omitempty" validate:"nonzero"`                         //
	Nodelete             bool   `json:"nodelete,omitempty" url:"nodelete,omitempty" validate:"nonzero"`                             //
	TargetSizeRatio      int    `json:"target_size_ratio,omitempty" url:"target_size_ratio,omitempty,optional"`                     // The estimated target ratio of the pool for the PG autoscaler.
	Application          string `json:"application,omitempty" url:"application,omitempty,optional"`                                 // The application of the pool.
	Name                 string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`                                     // The name of the pool. It must be unique.
	NodeepScrub          bool   `json:"nodeep-scrub,omitempty" url:"nodeep-scrub,omitempty" validate:"nonzero"`                     //
	Noscrub              bool   `json:"noscrub,omitempty" url:"noscrub,omitempty" validate:"nonzero"`                               //
	PgpNum               int    `json:"pgp_num,omitempty" url:"pgp_num,omitempty" validate:"nonzero"`                               //
	UseGmtHitset         bool   `json:"use_gmt_hitset,omitempty" url:"use_gmt_hitset,omitempty" validate:"nonzero"`                 //
	WriteFadviseDontneed bool   `json:"write_fadvise_dontneed,omitempty" url:"write_fadvise_dontneed,omitempty" validate:"nonzero"` //
	ApplicationList      []struct {
	} `json:"application_list,omitempty" url:"application_list,omitempty,optional"` //
	FastRead   bool     `json:"fast_read,omitempty" url:"fast_read,omitempty" validate:"nonzero"`   //
	Nopgchange bool     `json:"nopgchange,omitempty" url:"nopgchange,omitempty" validate:"nonzero"` //
	Statistics struct { //
	} `json:"statistics,omitempty" url:"statistics,omitempty,optional"` //
	TargetSize string `json:"target_size,omitempty" url:"target_size,omitempty,optional"` // The estimated target size of the pool for the PG autoscaler.
}

// NodesCephPoolsSetpoolRequest {name} - Change POOL settings
// Change POOL settings
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/pools/{name}
//
type NodesCephPoolsSetpoolRequest struct {
	Application     string `json:"application,omitempty" url:"application,omitempty,optional"`             // The application of the pool.
	MinSize         int    `json:"min_size,omitempty" url:"min_size,omitempty,optional"`                   // Minimum number of replicas per object
	Node            string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                 // The cluster node name.
	PgAutoscaleMode string `json:"pg_autoscale_mode,omitempty" url:"pg_autoscale_mode,omitempty,optional"` // The automatic PG scaling mode of the pool.
	Size            int    `json:"size,omitempty" url:"size,omitempty,optional"`                           // Number of replicas per object
	TargetSizeRatio int    `json:"target_size_ratio,omitempty" url:"target_size_ratio,omitempty,optional"` // The estimated target ratio of the pool for the PG autoscaler.
	CrushRule       string `json:"crush_rule,omitempty" url:"crush_rule,omitempty,optional"`               // The rule to use for mapping object placement in the cluster.
	Name            string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`                 // The name of the pool. It must be unique.
	PgNum           int    `json:"pg_num,omitempty" url:"pg_num,omitempty,optional"`                       // Number of placement groups.
	PgNumMin        int    `json:"pg_num_min,omitempty" url:"pg_num_min,omitempty,optional"`               // Minimal number of placement groups.
	TargetSize      string `json:"target_size,omitempty" url:"target_size,omitempty,optional"`             // The estimated target size of the pool for the PG autoscaler.
}

// NodesCephConfigRequest config - Get Ceph configuration.
// Get Ceph configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/config
//
type NodesCephConfigRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCephConfigdbRequest configdb - Get Ceph configuration database.
// Get Ceph configuration database.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/configdb
//
type NodesCephConfigdbRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCephConfigdbResponse configdb
// Get Ceph configuration database.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/configdb
//
type NodesCephConfigdbResponse struct {
	Level              string `json:"level,omitempty" url:"level,omitempty" validate:"nonzero"`                                 //
	Mask               string `json:"mask,omitempty" url:"mask,omitempty" validate:"nonzero"`                                   //
	Name               string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`                                   //
	Section            string `json:"section,omitempty" url:"section,omitempty" validate:"nonzero"`                             //
	Value              string `json:"value,omitempty" url:"value,omitempty" validate:"nonzero"`                                 //
	CanUpdateAtRuntime bool   `json:"can_update_at_runtime,omitempty" url:"can_update_at_runtime,omitempty" validate:"nonzero"` //
}

// NodesCephInitRequest init - Create initial ceph default configuration and setup symlinks.
// Create initial ceph default configuration and setup symlinks.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/init
//
type NodesCephInitRequest struct {
	MinSize        int    `json:"min_size,omitempty" url:"min_size,omitempty,optional"`               // Minimum number of available replicas per object to allow I/O
	Network        string `json:"network,omitempty" url:"network,omitempty,optional"`                 // Use specific network for all ceph related traffic
	Node           string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`             // The cluster node name.
	PgBits         int    `json:"pg_bits,omitempty" url:"pg_bits,omitempty,optional"`                 // Placement group bits, used to specify the default number of placement groups.  NOTE: 'osd pool default pg num' does not work for default pools.
	Size           int    `json:"size,omitempty" url:"size,omitempty,optional"`                       // Targeted number of replicas per object
	ClusterNetwork string `json:"cluster-network,omitempty" url:"cluster-network,omitempty,optional"` // Declare a separate cluster network, OSDs will routeheartbeat, object replication and recovery traffic over it
	DisableCephx   bool   `json:"disable_cephx,omitempty" url:"disable_cephx,omitempty,optional"`     // Disable cephx authentication.  WARNING: cephx is a security feature protecting against man-in-the-middle attacks. Only consider disabling cephx if your network is private!
}

// NodesCephStopRequest stop - Stop ceph services.
// Stop ceph services.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/stop
//
type NodesCephStopRequest struct {
	Service string `json:"service,omitempty" url:"service,omitempty,optional"`     // Ceph service name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCephStartRequest start - Start ceph services.
// Start ceph services.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/start
//
type NodesCephStartRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Service string `json:"service,omitempty" url:"service,omitempty,optional"`     // Ceph service name.
}

// NodesCephRestartRequest restart - Restart ceph services.
// Restart ceph services.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/restart
//
type NodesCephRestartRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Service string `json:"service,omitempty" url:"service,omitempty,optional"`     // Ceph service name.
}

// NodesCephStatusRequest status - Get ceph status.
// Get ceph status.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/status
//
type NodesCephStatusRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCephCrushRequest crush - Get OSD crush map
// Get OSD crush map
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/crush
//
type NodesCephCrushRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCephLogRequest log - Read ceph log
// Read ceph log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/log
//
type NodesCephLogRequest struct {
	Limit int    `json:"limit,omitempty" url:"limit,omitempty,optional"`         //
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Start int    `json:"start,omitempty" url:"start,omitempty,optional"`         //
}

// NodesCephLogResponse log
// Read ceph log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/log
//
type NodesCephLogResponse struct {
	T string `json:"t,omitempty" url:"t,omitempty" validate:"nonzero"` // Line text
	N int    `json:"n,omitempty" url:"n,omitempty" validate:"nonzero"` // Line number
}

// NodesCephRulesRequest rules - List ceph rules.
// List ceph rules.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/ceph/rules
//
type NodesCephRulesRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesVzdumpRequest vzdump - Create backup.
// Create backup.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/vzdump
//
type NodesVzdumpRequest struct {
	All              bool   `json:"all,omitempty" url:"all,omitempty,optional"`                           // Backup all known guest systems on this host.
	Bwlimit          int    `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`                   // Limit I/O bandwidth (KBytes per second).
	Stopwait         int    `json:"stopwait,omitempty" url:"stopwait,omitempty,optional"`                 // Maximal time to wait until a guest system is stopped (minutes).
	Exclude          string `json:"exclude,omitempty" url:"exclude,omitempty,optional"`                   // Exclude specified guest systems (assumes --all)
	ExcludePath      string `json:"exclude-path,omitempty" url:"exclude-path,omitempty,optional"`         // Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root,  other paths match relative to each subdirectory.
	Maxfiles         int    `json:"maxfiles,omitempty" url:"maxfiles,omitempty,optional"`                 // Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	Mode             string `json:"mode,omitempty" url:"mode,omitempty,optional"`                         // Backup mode.
	Stdexcludes      bool   `json:"stdexcludes,omitempty" url:"stdexcludes,omitempty,optional"`           // Exclude temporary files and logs.
	Mailnotification string `json:"mailnotification,omitempty" url:"mailnotification,omitempty,optional"` // Specify when to send an email
	NotesTemplate    string `json:"notes-template,omitempty" url:"notes-template,omitempty,optional"`     // Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future.
	Quiet            bool   `json:"quiet,omitempty" url:"quiet,omitempty,optional"`                       // Be quiet.
	Remove           bool   `json:"remove,omitempty" url:"remove,omitempty,optional"`                     // Prune older backups according to 'prune-backups'.
	Script           string `json:"script,omitempty" url:"script,omitempty,optional"`                     // Use specified hook script.
	Stop             bool   `json:"stop,omitempty" url:"stop,omitempty,optional"`                         // Stop running backup jobs on this host.
	Zstd             int    `json:"zstd,omitempty" url:"zstd,omitempty,optional"`                         // Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
	Dumpdir          string `json:"dumpdir,omitempty" url:"dumpdir,omitempty,optional"`                   // Store resulting files to specified directory.
	Ionice           int    `json:"ionice,omitempty" url:"ionice,omitempty,optional"`                     // Set CFQ ionice priority.
	Lockwait         int    `json:"lockwait,omitempty" url:"lockwait,omitempty,optional"`                 // Maximal time to wait for the global lock (minutes).
	Node             string `json:"node,omitempty" url:"node,omitempty,optional"`                         // Only run if executed on this node.
	Pool             string `json:"pool,omitempty" url:"pool,omitempty,optional"`                         // Backup all known guest systems included in the specified pool.
	PruneBackups     string `json:"prune-backups,omitempty" url:"prune-backups,omitempty,optional"`       // Use these retention options instead of those from the storage configuration.
	Tmpdir           string `json:"tmpdir,omitempty" url:"tmpdir,omitempty,optional"`                     // Store temporary files to specified directory.
	Compress         string `json:"compress,omitempty" url:"compress,omitempty,optional"`                 // Compress dump file.
	Mailto           string `json:"mailto,omitempty" url:"mailto,omitempty,optional"`                     // Comma-separated list of email addresses or users that should receive email notifications.
	Pigz             int    `json:"pigz,omitempty" url:"pigz,omitempty,optional"`                         // Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Storage          string `json:"storage,omitempty" url:"storage,omitempty,optional"`                   // Store resulting file to this storage.
	Vmid             string `json:"vmid,omitempty" url:"vmid,omitempty,optional"`                         // The ID of the guest system you want to backup.
	Protected        bool   `json:"protected,omitempty" url:"protected,omitempty,optional"`               // If true, mark backup(s) as protected.
	Stdout           bool   `json:"stdout,omitempty" url:"stdout,omitempty,optional"`                     // Write tar to stdout, not to a file.
}

// NodesVzdumpDefaultsRequest defaults - Get the currently configured vzdump defaults.
// Get the currently configured vzdump defaults.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/vzdump/defaults
//
type NodesVzdumpDefaultsRequest struct {
	Storage string `json:"storage,omitempty" url:"storage,omitempty,optional"`     // The storage identifier.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesVzdumpDefaultsResponse defaults
// Get the currently configured vzdump defaults.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/vzdump/defaults
//
type NodesVzdumpDefaultsResponse struct {
	Lockwait         int    `json:"lockwait,omitempty" url:"lockwait,omitempty,optional"`                 // Maximal time to wait for the global lock (minutes).
	Pool             string `json:"pool,omitempty" url:"pool,omitempty,optional"`                         // Backup all known guest systems included in the specified pool.
	All              bool   `json:"all,omitempty" url:"all,omitempty,optional"`                           // Backup all known guest systems on this host.
	Bwlimit          int    `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`                   // Limit I/O bandwidth (KBytes per second).
	Tmpdir           string `json:"tmpdir,omitempty" url:"tmpdir,omitempty,optional"`                     // Store temporary files to specified directory.
	Mailnotification string `json:"mailnotification,omitempty" url:"mailnotification,omitempty,optional"` // Specify when to send an email
	Quiet            bool   `json:"quiet,omitempty" url:"quiet,omitempty,optional"`                       // Be quiet.
	Zstd             int    `json:"zstd,omitempty" url:"zstd,omitempty,optional"`                         // Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
	Mode             string `json:"mode,omitempty" url:"mode,omitempty,optional"`                         // Backup mode.
	Node             string `json:"node,omitempty" url:"node,omitempty,optional"`                         // Only run if executed on this node.
	Ionice           int    `json:"ionice,omitempty" url:"ionice,omitempty,optional"`                     // Set CFQ ionice priority.
	Mailto           string `json:"mailto,omitempty" url:"mailto,omitempty,optional"`                     // Comma-separated list of email addresses or users that should receive email notifications.
	Exclude          string `json:"exclude,omitempty" url:"exclude,omitempty,optional"`                   // Exclude specified guest systems (assumes --all)
	ExcludePath      string `json:"exclude-path,omitempty" url:"exclude-path,omitempty,optional"`         // Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root,  other paths match relative to each subdirectory.
	Compress         string `json:"compress,omitempty" url:"compress,omitempty,optional"`                 // Compress dump file.
	Protected        bool   `json:"protected,omitempty" url:"protected,omitempty,optional"`               // If true, mark backup(s) as protected.
	Remove           bool   `json:"remove,omitempty" url:"remove,omitempty,optional"`                     // Prune older backups according to 'prune-backups'.
	Stdexcludes      bool   `json:"stdexcludes,omitempty" url:"stdexcludes,omitempty,optional"`           // Exclude temporary files and logs.
	Stop             bool   `json:"stop,omitempty" url:"stop,omitempty,optional"`                         // Stop running backup jobs on this host.
	Storage          string `json:"storage,omitempty" url:"storage,omitempty,optional"`                   // Store resulting file to this storage.
	Pigz             int    `json:"pigz,omitempty" url:"pigz,omitempty,optional"`                         // Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Script           string `json:"script,omitempty" url:"script,omitempty,optional"`                     // Use specified hook script.
	Vmid             string `json:"vmid,omitempty" url:"vmid,omitempty,optional"`                         // The ID of the guest system you want to backup.
	Maxfiles         int    `json:"maxfiles,omitempty" url:"maxfiles,omitempty,optional"`                 // Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	NotesTemplate    string `json:"notes-template,omitempty" url:"notes-template,omitempty,optional"`     // Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future.
	Stopwait         int    `json:"stopwait,omitempty" url:"stopwait,omitempty,optional"`                 // Maximal time to wait until a guest system is stopped (minutes).
	Dumpdir          string `json:"dumpdir,omitempty" url:"dumpdir,omitempty,optional"`                   // Store resulting files to specified directory.
	PruneBackups     string `json:"prune-backups,omitempty" url:"prune-backups,omitempty,optional"`       // Use these retention options instead of those from the storage configuration.
}

// NodesVzdumpExtractconfigRequest extractconfig - Extract configuration from vzdump backup archive.
// Extract configuration from vzdump backup archive.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/vzdump/extractconfig
//
type NodesVzdumpExtractconfigRequest struct {
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Volume string `json:"volume,omitempty" url:"volume,omitempty" validate:"nonzero"` // Volume identifier
}

// NodesServicesSrvcmdidxRequest {service} - Directory index
// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/services/{service}
//
type NodesServicesSrvcmdidxRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Service string `json:"service,omitempty" url:"service,omitempty" validate:"nonzero"` // Service ID
}

// NodesServicesSrvcmdidxResponse {service}
// Directory index
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/services/{service}
//
type NodesServicesSrvcmdidxResponse struct {
	Subdir string `json:"subdir,omitempty" url:"subdir,omitempty" validate:"nonzero"` //
}

// NodesServicesStateServiceRequest state - Read service properties
// Read service properties
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/services/{service}/state
//
type NodesServicesStateServiceRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Service string `json:"service,omitempty" url:"service,omitempty" validate:"nonzero"` // Service ID
}

// NodesServicesStartServiceRequest start - Start service.
// Start service.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/services/{service}/start
//
type NodesServicesStartServiceRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Service string `json:"service,omitempty" url:"service,omitempty" validate:"nonzero"` // Service ID
}

// NodesServicesStopServiceRequest stop - Stop service.
// Stop service.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/services/{service}/stop
//
type NodesServicesStopServiceRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Service string `json:"service,omitempty" url:"service,omitempty" validate:"nonzero"` // Service ID
}

// NodesServicesRestartServiceRequest restart - Hard restart service. Use reload if you want to reduce interruptions.
// Hard restart service. Use reload if you want to reduce interruptions.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/services/{service}/restart
//
type NodesServicesRestartServiceRequest struct {
	Service string `json:"service,omitempty" url:"service,omitempty" validate:"nonzero"` // Service ID
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
}

// NodesServicesReloadServiceRequest reload - Reload service. Falls back to restart if service cannot be reloaded.
// Reload service. Falls back to restart if service cannot be reloaded.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/services/{service}/reload
//
type NodesServicesReloadServiceRequest struct {
	Service string `json:"service,omitempty" url:"service,omitempty" validate:"nonzero"` // Service ID
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
}

// NodesSubscriptionUpdateRequest subscription - Update subscription info.
// Update subscription info.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/subscription
//
type NodesSubscriptionUpdateRequest struct {
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Force bool   `json:"force,omitempty" url:"force,omitempty,optional"`         // Always connect to server, even if we have up to date info inside local cache.
}

// NodesSubscriptionSetRequest subscription - Set subscription key.
// Set subscription key.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/subscription
//
type NodesSubscriptionSetRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Key  string `json:"key,omitempty" url:"key,omitempty" validate:"nonzero"`   // Proxmox VE subscription key
}

// NodesSubscriptionDeleteRequest subscription - Delete subscription key of this node.
// Delete subscription key of this node.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/subscription
//
type NodesSubscriptionDeleteRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesSubscriptionGetRequest subscription - Read subscription info.
// Read subscription info.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/subscription
//
type NodesSubscriptionGetRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesNetworkRevertChangesRequest network - Revert network configuration changes.
// Revert network configuration changes.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/network
//
type NodesNetworkRevertChangesRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesNetworkCreateRequest network - Create network device configuration
// Create network device configuration
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/network
//
type NodesNetworkCreateRequest struct {
	Comments6          string `json:"comments6,omitempty" url:"comments6,omitempty,optional"`                         // Comments
	Iface              string `json:"iface,omitempty" url:"iface,omitempty" validate:"nonzero"`                       // Network interface name.
	VlanRawDevice      string `json:"vlan-raw-device,omitempty" url:"vlan-raw-device,omitempty,optional"`             // Specify the raw interface for the vlan interface.
	Gateway            string `json:"gateway,omitempty" url:"gateway,omitempty,optional"`                             // Default gateway address.
	Address6           string `json:"address6,omitempty" url:"address6,omitempty,optional"`                           // IP address.
	OvsTag             int    `json:"ovs_tag,omitempty" url:"ovs_tag,omitempty,optional"`                             // Specify a VLan tag (used by OVSPort, OVSIntPort, OVSBond)
	Netmask6           int    `json:"netmask6,omitempty" url:"netmask6,omitempty,optional"`                           // Network mask.
	BridgePorts        string `json:"bridge_ports,omitempty" url:"bridge_ports,omitempty,optional"`                   // Specify the interfaces you want to add to your bridge.
	BridgeVlanAware    bool   `json:"bridge_vlan_aware,omitempty" url:"bridge_vlan_aware,omitempty,optional"`         // Enable bridge vlan support.
	Node               string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                         // The cluster node name.
	OvsBonds           string `json:"ovs_bonds,omitempty" url:"ovs_bonds,omitempty,optional"`                         // Specify the interfaces used by the bonding device.
	OvsBridge          string `json:"ovs_bridge,omitempty" url:"ovs_bridge,omitempty,optional"`                       // The OVS bridge associated with a OVS port. This is required when you create an OVS port.
	Address            string `json:"address,omitempty" url:"address,omitempty,optional"`                             // IP address.
	Type               string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`                         // Network interface type
	VlanId             int    `json:"vlan-id,omitempty" url:"vlan-id,omitempty,optional"`                             // vlan-id for a custom named vlan interface (ifupdown2 only).
	Autostart          bool   `json:"autostart,omitempty" url:"autostart,omitempty,optional"`                         // Automatically start interface on boot.
	Cidr               string `json:"cidr,omitempty" url:"cidr,omitempty,optional"`                                   // IPv4 CIDR.
	Mtu                int    `json:"mtu,omitempty" url:"mtu,omitempty,optional"`                                     // MTU.
	BondPrimary        string `json:"bond-primary,omitempty" url:"bond-primary,omitempty,optional"`                   // Specify the primary interface for active-backup bond.
	Netmask            string `json:"netmask,omitempty" url:"netmask,omitempty,optional"`                             // Network mask.
	OvsOptions         string `json:"ovs_options,omitempty" url:"ovs_options,omitempty,optional"`                     // OVS interface options.
	Comments           string `json:"comments,omitempty" url:"comments,omitempty,optional"`                           // Comments
	Gateway6           string `json:"gateway6,omitempty" url:"gateway6,omitempty,optional"`                           // Default ipv6 gateway address.
	BondMode           string `json:"bond_mode,omitempty" url:"bond_mode,omitempty,optional"`                         // Bonding mode.
	BondXmitHashPolicy string `json:"bond_xmit_hash_policy,omitempty" url:"bond_xmit_hash_policy,omitempty,optional"` // Selects the transmit hash policy to use for slave selection in balance-xor and 802.3ad modes.
	Cidr6              string `json:"cidr6,omitempty" url:"cidr6,omitempty,optional"`                                 // IPv6 CIDR.
	OvsPorts           string `json:"ovs_ports,omitempty" url:"ovs_ports,omitempty,optional"`                         // Specify the interfaces you want to add to your bridge.
	Slaves             string `json:"slaves,omitempty" url:"slaves,omitempty,optional"`                               // Specify the interfaces used by the bonding device.
}

// NodesNetworkReloadConfigRequest network - Reload network configuration
// Reload network configuration
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/network
//
type NodesNetworkReloadConfigRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesNetworkDeleteRequest {iface} - Delete network device configuration
// Delete network device configuration
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/network/{iface}
//
type NodesNetworkDeleteRequest struct {
	Iface string `json:"iface,omitempty" url:"iface,omitempty" validate:"nonzero"` // Network interface name.
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
}

// NodesNetworkConfigRequest {iface} - Read network device configuration
// Read network device configuration
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/network/{iface}
//
type NodesNetworkConfigRequest struct {
	Iface string `json:"iface,omitempty" url:"iface,omitempty" validate:"nonzero"` // Network interface name.
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
}

// NodesNetworkConfigResponse {iface}
// Read network device configuration
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/network/{iface}
//
type NodesNetworkConfigResponse struct {
	Method string `json:"method,omitempty" url:"method,omitempty" validate:"nonzero"` //
	Type   string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     //
}

// NodesNetworkUpdateRequest {iface} - Update network device configuration
// Update network device configuration
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/network/{iface}
//
type NodesNetworkUpdateRequest struct {
	BridgeVlanAware    bool   `json:"bridge_vlan_aware,omitempty" url:"bridge_vlan_aware,omitempty,optional"`         // Enable bridge vlan support.
	Type               string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`                         // Network interface type
	Address            string `json:"address,omitempty" url:"address,omitempty,optional"`                             // IP address.
	BondPrimary        string `json:"bond-primary,omitempty" url:"bond-primary,omitempty,optional"`                   // Specify the primary interface for active-backup bond.
	Delete             string `json:"delete,omitempty" url:"delete,omitempty,optional"`                               // A list of settings you want to delete.
	OvsOptions         string `json:"ovs_options,omitempty" url:"ovs_options,omitempty,optional"`                     // OVS interface options.
	Address6           string `json:"address6,omitempty" url:"address6,omitempty,optional"`                           // IP address.
	Cidr6              string `json:"cidr6,omitempty" url:"cidr6,omitempty,optional"`                                 // IPv6 CIDR.
	BridgePorts        string `json:"bridge_ports,omitempty" url:"bridge_ports,omitempty,optional"`                   // Specify the interfaces you want to add to your bridge.
	OvsPorts           string `json:"ovs_ports,omitempty" url:"ovs_ports,omitempty,optional"`                         // Specify the interfaces you want to add to your bridge.
	VlanId             int    `json:"vlan-id,omitempty" url:"vlan-id,omitempty,optional"`                             // vlan-id for a custom named vlan interface (ifupdown2 only).
	Autostart          bool   `json:"autostart,omitempty" url:"autostart,omitempty,optional"`                         // Automatically start interface on boot.
	BondMode           string `json:"bond_mode,omitempty" url:"bond_mode,omitempty,optional"`                         // Bonding mode.
	Netmask6           int    `json:"netmask6,omitempty" url:"netmask6,omitempty,optional"`                           // Network mask.
	Cidr               string `json:"cidr,omitempty" url:"cidr,omitempty,optional"`                                   // IPv4 CIDR.
	Comments           string `json:"comments,omitempty" url:"comments,omitempty,optional"`                           // Comments
	OvsBonds           string `json:"ovs_bonds,omitempty" url:"ovs_bonds,omitempty,optional"`                         // Specify the interfaces used by the bonding device.
	Iface              string `json:"iface,omitempty" url:"iface,omitempty" validate:"nonzero"`                       // Network interface name.
	Node               string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                         // The cluster node name.
	Netmask            string `json:"netmask,omitempty" url:"netmask,omitempty,optional"`                             // Network mask.
	BondXmitHashPolicy string `json:"bond_xmit_hash_policy,omitempty" url:"bond_xmit_hash_policy,omitempty,optional"` // Selects the transmit hash policy to use for slave selection in balance-xor and 802.3ad modes.
	Comments6          string `json:"comments6,omitempty" url:"comments6,omitempty,optional"`                         // Comments
	OvsTag             int    `json:"ovs_tag,omitempty" url:"ovs_tag,omitempty,optional"`                             // Specify a VLan tag (used by OVSPort, OVSIntPort, OVSBond)
	Slaves             string `json:"slaves,omitempty" url:"slaves,omitempty,optional"`                               // Specify the interfaces used by the bonding device.
	Gateway            string `json:"gateway,omitempty" url:"gateway,omitempty,optional"`                             // Default gateway address.
	OvsBridge          string `json:"ovs_bridge,omitempty" url:"ovs_bridge,omitempty,optional"`                       // The OVS bridge associated with a OVS port. This is required when you create an OVS port.
	VlanRawDevice      string `json:"vlan-raw-device,omitempty" url:"vlan-raw-device,omitempty,optional"`             // Specify the raw interface for the vlan interface.
	Gateway6           string `json:"gateway6,omitempty" url:"gateway6,omitempty,optional"`                           // Default ipv6 gateway address.
	Mtu                int    `json:"mtu,omitempty" url:"mtu,omitempty,optional"`                                     // MTU.
}

// NodesTasksNodeRequest tasks - Read task list for one node (finished tasks).
// Read task list for one node (finished tasks).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/tasks
//
type NodesTasksNodeRequest struct {
	Errors       bool   `json:"errors,omitempty" url:"errors,omitempty,optional"`             // Only list tasks with a status of ERROR.
	Since        int    `json:"since,omitempty" url:"since,omitempty,optional"`               // Only list tasks since this UNIX epoch.
	Start        int    `json:"start,omitempty" url:"start,omitempty,optional"`               // List tasks beginning from this offset.
	Typefilter   string `json:"typefilter,omitempty" url:"typefilter,omitempty,optional"`     // Only list tasks of this type (e.g., vzstart, vzdump).
	Userfilter   string `json:"userfilter,omitempty" url:"userfilter,omitempty,optional"`     // Only list tasks from this user.
	Vmid         string `json:"vmid,omitempty" url:"vmid,omitempty,optional"`                 // Only list tasks for this VM.
	Limit        int    `json:"limit,omitempty" url:"limit,omitempty,optional"`               // Only list this amount of tasks.
	Node         string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Source       string `json:"source,omitempty" url:"source,omitempty,optional"`             // List archived, active or all tasks.
	Statusfilter string `json:"statusfilter,omitempty" url:"statusfilter,omitempty,optional"` // List of Task States that should be returned.
	Until        int    `json:"until,omitempty" url:"until,omitempty,optional"`               // Only list tasks until this UNIX epoch.
}

// NodesTasksNodeResponse tasks
// Read task list for one node (finished tasks).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/tasks
//
type NodesTasksNodeResponse struct {
	Upid      string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"`           //
	Id        string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`               //
	Pstart    int    `json:"pstart,omitempty" url:"pstart,omitempty" validate:"nonzero"`       //
	Status    string `json:"status,omitempty" url:"status,omitempty,optional"`                 //
	Type      string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`           //
	User      string `json:"user,omitempty" url:"user,omitempty" validate:"nonzero"`           //
	Endtime   int    `json:"endtime,omitempty" url:"endtime,omitempty,optional"`               //
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           //
	Pid       int    `json:"pid,omitempty" url:"pid,omitempty" validate:"nonzero"`             //
	Starttime int    `json:"starttime,omitempty" url:"starttime,omitempty" validate:"nonzero"` //
}

// NodesTasksStopTaskRequest {upid} - Stop a task.
// Stop a task.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/tasks/{upid}
//
type NodesTasksStopTaskRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Upid string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"` //
}

// NodesTasksUpidIndexRequest {upid} -
//
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/tasks/{upid}
//
type NodesTasksUpidIndexRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Upid string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"` //
}

// NodesTasksLogReadTaskRequest log - Read task log.
// Read task log.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/tasks/{upid}/log
//
type NodesTasksLogReadTaskRequest struct {
	Limit int    `json:"limit,omitempty" url:"limit,omitempty,optional"`         // The maximum amount of lines that should be printed.
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Start int    `json:"start,omitempty" url:"start,omitempty,optional"`         // The line number to start printing at.
	Upid  string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"` // The task's unique ID.
}

// NodesTasksLogReadTaskResponse log
// Read task log.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/tasks/{upid}/log
//
type NodesTasksLogReadTaskResponse struct {
	N int    `json:"n,omitempty" url:"n,omitempty" validate:"nonzero"` // Line number
	T string `json:"t,omitempty" url:"t,omitempty" validate:"nonzero"` // Line text
}

// NodesTasksStatusReadTaskRequest status - Read task status.
// Read task status.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/tasks/{upid}/status
//
type NodesTasksStatusReadTaskRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Upid string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"` // The task's unique ID.
}

// NodesTasksStatusReadTaskResponse status
// Read task status.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/tasks/{upid}/status
//
type NodesTasksStatusReadTaskResponse struct {
	User       string `json:"user,omitempty" url:"user,omitempty" validate:"nonzero"`           //
	Exitstatus string `json:"exitstatus,omitempty" url:"exitstatus,omitempty,optional"`         //
	Pid        int    `json:"pid,omitempty" url:"pid,omitempty" validate:"nonzero"`             //
	Starttime  int    `json:"starttime,omitempty" url:"starttime,omitempty" validate:"nonzero"` //
	Type       string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`           //
	Id         string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`               //
	Node       string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           //
	Status     string `json:"status,omitempty" url:"status,omitempty" validate:"nonzero"`       //
	Upid       string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"`           //
}

// NodesScanNfsScanRequest nfs - Scan remote NFS server.
// Scan remote NFS server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/nfs
//
type NodesScanNfsScanRequest struct {
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Server string `json:"server,omitempty" url:"server,omitempty" validate:"nonzero"` // The server address (name or IP).
}

// NodesScanNfsScanResponse nfs
// Scan remote NFS server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/nfs
//
type NodesScanNfsScanResponse struct {
	Options string `json:"options,omitempty" url:"options,omitempty" validate:"nonzero"` // NFS export options.
	Path    string `json:"path,omitempty" url:"path,omitempty" validate:"nonzero"`       // The exported path.
}

// NodesScanCifsScanRequest cifs - Scan remote CIFS server.
// Scan remote CIFS server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/cifs
//
type NodesScanCifsScanRequest struct {
	Server   string `json:"server,omitempty" url:"server,omitempty" validate:"nonzero"` // The server address (name or IP).
	Username string `json:"username,omitempty" url:"username,omitempty,optional"`       // User name.
	Domain   string `json:"domain,omitempty" url:"domain,omitempty,optional"`           // SMB domain (Workgroup).
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Password string `json:"password,omitempty" url:"password,omitempty,optional"`       // User password.
}

// NodesScanCifsScanResponse cifs
// Scan remote CIFS server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/cifs
//
type NodesScanCifsScanResponse struct {
	Description string `json:"description,omitempty" url:"description,omitempty" validate:"nonzero"` // Descriptive text from server.
	Share       string `json:"share,omitempty" url:"share,omitempty" validate:"nonzero"`             // The cifs share name.
}

// NodesScanPbsScanRequest pbs - Scan remote Proxmox Backup Server.
// Scan remote Proxmox Backup Server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/pbs
//
type NodesScanPbsScanRequest struct {
	Fingerprint string `json:"fingerprint,omitempty" url:"fingerprint,omitempty,optional"`     // Certificate SHA 256 fingerprint.
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Password    string `json:"password,omitempty" url:"password,omitempty" validate:"nonzero"` // User password or API token secret.
	Port        int    `json:"port,omitempty" url:"port,omitempty,optional"`                   // Optional port.
	Server      string `json:"server,omitempty" url:"server,omitempty" validate:"nonzero"`     // The server address (name or IP).
	Username    string `json:"username,omitempty" url:"username,omitempty" validate:"nonzero"` // User-name or API token-ID.
}

// NodesScanPbsScanResponse pbs
// Scan remote Proxmox Backup Server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/pbs
//
type NodesScanPbsScanResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`       // Comment from server.
	Store   string `json:"store,omitempty" url:"store,omitempty" validate:"nonzero"` // The datastore name.
}

// NodesScanGlusterfsScanRequest glusterfs - Scan remote GlusterFS server.
// Scan remote GlusterFS server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/glusterfs
//
type NodesScanGlusterfsScanRequest struct {
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Server string `json:"server,omitempty" url:"server,omitempty" validate:"nonzero"` // The server address (name or IP).
}

// NodesScanGlusterfsScanResponse glusterfs
// Scan remote GlusterFS server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/glusterfs
//
type NodesScanGlusterfsScanResponse struct {
	Volname string `json:"volname,omitempty" url:"volname,omitempty" validate:"nonzero"` // The volume name.
}

// NodesScanIscsiScanRequest iscsi - Scan remote iSCSI server.
// Scan remote iSCSI server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/iscsi
//
type NodesScanIscsiScanRequest struct {
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Portal string `json:"portal,omitempty" url:"portal,omitempty" validate:"nonzero"` // The iSCSI portal (IP or DNS name with optional port).
}

// NodesScanIscsiScanResponse iscsi
// Scan remote iSCSI server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/iscsi
//
type NodesScanIscsiScanResponse struct {
	Portal string `json:"portal,omitempty" url:"portal,omitempty" validate:"nonzero"` // The iSCSI portal name.
	Target string `json:"target,omitempty" url:"target,omitempty" validate:"nonzero"` // The iSCSI target name.
}

// NodesScanLvmScanRequest lvm - List local LVM volume groups.
// List local LVM volume groups.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/lvm
//
type NodesScanLvmScanRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesScanLvmScanResponse lvm
// List local LVM volume groups.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/lvm
//
type NodesScanLvmScanResponse struct {
	Vg string `json:"vg,omitempty" url:"vg,omitempty" validate:"nonzero"` // The LVM logical volume group name.
}

// NodesScanLvmthinScanRequest lvmthin - List local LVM Thin Pools.
// List local LVM Thin Pools.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/lvmthin
//
type NodesScanLvmthinScanRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vg   string `json:"vg,omitempty" url:"vg,omitempty" validate:"nonzero"`     //
}

// NodesScanLvmthinScanResponse lvmthin
// List local LVM Thin Pools.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/lvmthin
//
type NodesScanLvmthinScanResponse struct {
	Lv string `json:"lv,omitempty" url:"lv,omitempty" validate:"nonzero"` // The LVM Thin Pool name (LVM logical volume).
}

// NodesScanZfsScanRequest zfs - Scan zfs pool list on local node.
// Scan zfs pool list on local node.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/zfs
//
type NodesScanZfsScanRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesScanZfsScanResponse zfs
// Scan zfs pool list on local node.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/scan/zfs
//
type NodesScanZfsScanResponse struct {
	Pool string `json:"pool,omitempty" url:"pool,omitempty" validate:"nonzero"` // ZFS pool name.
}

// NodesHardwarePciScanRequest pci - List local PCI devices.
// List local PCI devices.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hardware/pci
//
type NodesHardwarePciScanRequest struct {
	Node              string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                     // The cluster node name.
	PciClassBlacklist string `json:"pci-class-blacklist,omitempty" url:"pci-class-blacklist,omitempty,optional"` // A list of blacklisted PCI classes, which will not be returned. Following are filtered by default: Memory Controller (05), Bridge (06) and Processor (0b).
	Verbose           bool   `json:"verbose,omitempty" url:"verbose,omitempty,optional"`                         // If disabled, does only print the PCI IDs. Otherwise, additional information like vendor and device will be returned.
}

// NodesHardwarePciScanResponse pci
// List local PCI devices.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hardware/pci
//
type NodesHardwarePciScanResponse struct {
	Vendor              string `json:"vendor,omitempty" url:"vendor,omitempty" validate:"nonzero"`                     // The Vendor ID.
	VendorName          string `json:"vendor_name,omitempty" url:"vendor_name,omitempty,optional"`                     //
	DeviceName          string `json:"device_name,omitempty" url:"device_name,omitempty,optional"`                     //
	Mdev                bool   `json:"mdev,omitempty" url:"mdev,omitempty,optional"`                                   // If set, marks that the device is capable of creating mediated devices.
	SubsystemDevice     string `json:"subsystem_device,omitempty" url:"subsystem_device,omitempty,optional"`           // The Subsystem Device ID.
	SubsystemVendor     string `json:"subsystem_vendor,omitempty" url:"subsystem_vendor,omitempty,optional"`           // The Subsystem Vendor ID.
	SubsystemDeviceName string `json:"subsystem_device_name,omitempty" url:"subsystem_device_name,omitempty,optional"` //
	SubsystemVendorName string `json:"subsystem_vendor_name,omitempty" url:"subsystem_vendor_name,omitempty,optional"` //
	Class               string `json:"class,omitempty" url:"class,omitempty" validate:"nonzero"`                       // The PCI Class of the device.
	Device              string `json:"device,omitempty" url:"device,omitempty" validate:"nonzero"`                     // The Device ID.
	Id                  string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`                             // The PCI ID.
	Iommugroup          int    `json:"iommugroup,omitempty" url:"iommugroup,omitempty" validate:"nonzero"`             // The IOMMU group in which the device is in. If no IOMMU group is detected, it is set to -1.
}

// NodesHardwarePciIndexRequest {pciid} - Index of available pci methods
// Index of available pci methods
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hardware/pci/{pciid}
//
type NodesHardwarePciIndexRequest struct {
	Pciid string `json:"pciid,omitempty" url:"pciid,omitempty" validate:"nonzero"` //
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
}

// NodesHardwarePciIndexResponse {pciid}
// Index of available pci methods
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hardware/pci/{pciid}
//
type NodesHardwarePciIndexResponse struct {
	Method string `json:"method,omitempty" url:"method,omitempty" validate:"nonzero"` //
}

// NodesHardwarePciMdevScanRequest mdev - List mediated device types for given PCI device.
// List mediated device types for given PCI device.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hardware/pci/{pciid}/mdev
//
type NodesHardwarePciMdevScanRequest struct {
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
	Pciid string `json:"pciid,omitempty" url:"pciid,omitempty" validate:"nonzero"` // The PCI ID to list the mdev types for.
}

// NodesHardwarePciMdevScanResponse mdev
// List mediated device types for given PCI device.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hardware/pci/{pciid}/mdev
//
type NodesHardwarePciMdevScanResponse struct {
	Available   int    `json:"available,omitempty" url:"available,omitempty" validate:"nonzero"`     // The number of still available instances of this type.
	Description string `json:"description,omitempty" url:"description,omitempty" validate:"nonzero"` //
	Type        string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`               // The name of the mdev type.
}

// NodesHardwareUsbScanRequest usb - List local USB devices.
// List local USB devices.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hardware/usb
//
type NodesHardwareUsbScanRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesHardwareUsbScanResponse usb
// List local USB devices.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hardware/usb
//
type NodesHardwareUsbScanResponse struct {
	Port         int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`       //
	Product      string `json:"product,omitempty" url:"product,omitempty,optional"`           //
	Speed        string `json:"speed,omitempty" url:"speed,omitempty" validate:"nonzero"`     //
	Usbpath      string `json:"usbpath,omitempty" url:"usbpath,omitempty,optional"`           //
	Busnum       int    `json:"busnum,omitempty" url:"busnum,omitempty" validate:"nonzero"`   //
	Class        int    `json:"class,omitempty" url:"class,omitempty" validate:"nonzero"`     //
	Devnum       int    `json:"devnum,omitempty" url:"devnum,omitempty" validate:"nonzero"`   //
	Serial       string `json:"serial,omitempty" url:"serial,omitempty,optional"`             //
	Vendid       string `json:"vendid,omitempty" url:"vendid,omitempty" validate:"nonzero"`   //
	Level        int    `json:"level,omitempty" url:"level,omitempty" validate:"nonzero"`     //
	Manufacturer string `json:"manufacturer,omitempty" url:"manufacturer,omitempty,optional"` //
	Prodid       string `json:"prodid,omitempty" url:"prodid,omitempty" validate:"nonzero"`   //
}

// NodesCapabilitiesQemuCapsIndexRequest qemu - QEMU capabilities index.
// QEMU capabilities index.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/capabilities/qemu
//
type NodesCapabilitiesQemuCapsIndexRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCapabilitiesQemuMachinesTypesRequest machines - Get available QEMU/KVM machine types.
// Get available QEMU/KVM machine types.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/capabilities/qemu/machines
//
type NodesCapabilitiesQemuMachinesTypesRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCapabilitiesQemuMachinesTypesResponse machines
// Get available QEMU/KVM machine types.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/capabilities/qemu/machines
//
type NodesCapabilitiesQemuMachinesTypesResponse struct {
	Id      string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`           // Full name of machine type and version.
	Type    string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`       // The machine type.
	Version string `json:"version,omitempty" url:"version,omitempty" validate:"nonzero"` // The machine version.
}

// NodesStorageDiridxRequest {storage} -
//
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}
//
type NodesStorageDiridxRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Storage string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"` // The storage identifier.
}

// NodesStorageDiridxResponse {storage}
//
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}
//
type NodesStorageDiridxResponse struct {
	Subdir string `json:"subdir,omitempty" url:"subdir,omitempty" validate:"nonzero"` //
}

// NodesStoragePrunebackupsDeleteRequest prunebackups - Prune backups. Only those using the standard naming scheme are considered.
// Prune backups. Only those using the standard naming scheme are considered.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/prunebackups
//
type NodesStoragePrunebackupsDeleteRequest struct {
	Node         string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	PruneBackups string `json:"prune-backups,omitempty" url:"prune-backups,omitempty,optional"` // Use these retention options instead of those from the storage configuration.
	Storage      string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`   // The storage identifier.
	Type         string `json:"type,omitempty" url:"type,omitempty,optional"`                   // Either 'qemu' or 'lxc'. Only consider backups for guests of this type.
	Vmid         string `json:"vmid,omitempty" url:"vmid,omitempty,optional"`                   // Only prune backups for this VM.
}

// NodesStoragePrunebackupsDryrunRequest prunebackups - Get prune information for backups. NOTE: this is only a preview and might not be what a subsequent prune call does if backups are removed/added in the meantime.
// Get prune information for backups. NOTE: this is only a preview and might not be what a subsequent prune call does if backups are removed/added in the meantime.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/prunebackups
//
type NodesStoragePrunebackupsDryrunRequest struct {
	PruneBackups string `json:"prune-backups,omitempty" url:"prune-backups,omitempty,optional"` // Use these retention options instead of those from the storage configuration.
	Storage      string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`   // The storage identifier.
	Type         string `json:"type,omitempty" url:"type,omitempty,optional"`                   // Either 'qemu' or 'lxc'. Only consider backups for guests of this type.
	Vmid         string `json:"vmid,omitempty" url:"vmid,omitempty,optional"`                   // Only consider backups for this guest.
	Node         string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
}

// NodesStoragePrunebackupsDryrunResponse prunebackups
// Get prune information for backups. NOTE: this is only a preview and might not be what a subsequent prune call does if backups are removed/added in the meantime.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/prunebackups
//
type NodesStoragePrunebackupsDryrunResponse struct {
	Ctime int    `json:"ctime,omitempty" url:"ctime,omitempty" validate:"nonzero"` // Creation time of the backup (seconds since the UNIX epoch).
	Mark  string `json:"mark,omitempty" url:"mark,omitempty" validate:"nonzero"`   // Whether the backup would be kept or removed. Backups that are protected or don't use the standard naming scheme are not removed.
	Type  string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`   // One of 'qemu', 'lxc', 'openvz' or 'unknown'.
	Vmid  string `json:"vmid,omitempty" url:"vmid,omitempty,optional"`             // The VM the backup belongs to.
	Volid string `json:"volid,omitempty" url:"volid,omitempty" validate:"nonzero"` // Backup volume ID.
}

// NodesStorageContentCreateRequest content - Allocate disk images.
// Allocate disk images.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/content
//
type NodesStorageContentCreateRequest struct {
	Filename string `json:"filename,omitempty" url:"filename,omitempty" validate:"nonzero"` // The name of the file to create.
	Format   string `json:"format,omitempty" url:"format,omitempty,optional"`               //
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Size     string `json:"size,omitempty" url:"size,omitempty" validate:"nonzero"`         // Size in kilobyte (1024 bytes). Optional suffixes 'M' (megabyte, 1024K) and 'G' (gigabyte, 1024M)
	Storage  string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`   // The storage identifier.
	Vmid     string `json:"vmid,omitempty" url:"vmid,omitempty" validate:"nonzero"`         // Specify owner VM
}

// NodesStorageContentDeleteRequest {volume} - Delete volume
// Delete volume
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/content/{volume}
//
type NodesStorageContentDeleteRequest struct {
	Volume  string `json:"volume,omitempty" url:"volume,omitempty" validate:"nonzero"` // Volume identifier
	Delay   int    `json:"delay,omitempty" url:"delay,omitempty,optional"`             // Time to wait for the task to finish. We return 'null' if the task finish within that time.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Storage string `json:"storage,omitempty" url:"storage,omitempty,optional"`         // The storage identifier.
}

// NodesStorageContentInfoRequest {volume} - Get volume attributes
// Get volume attributes
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/content/{volume}
//
type NodesStorageContentInfoRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Storage string `json:"storage,omitempty" url:"storage,omitempty,optional"`         // The storage identifier.
	Volume  string `json:"volume,omitempty" url:"volume,omitempty" validate:"nonzero"` // Volume identifier
}

// NodesStorageContentInfoResponse {volume}
// Get volume attributes
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/content/{volume}
//
type NodesStorageContentInfoResponse struct {
	Protected bool   `json:"protected,omitempty" url:"protected,omitempty,optional"`     // Protection status. Currently only supported for backups.
	Size      int    `json:"size,omitempty" url:"size,omitempty" validate:"nonzero"`     // Volume size in bytes.
	Used      int    `json:"used,omitempty" url:"used,omitempty" validate:"nonzero"`     // Used space. Please note that most storage plugins do not report anything useful here.
	Format    string `json:"format,omitempty" url:"format,omitempty" validate:"nonzero"` // Format identifier ('raw', 'qcow2', 'subvol', 'iso', 'tgz' ...)
	Notes     string `json:"notes,omitempty" url:"notes,omitempty,optional"`             // Optional notes.
	Path      string `json:"path,omitempty" url:"path,omitempty" validate:"nonzero"`     // The Path
}

// NodesStorageContentCopyRequest {volume} - Copy a volume. This is experimental code - do not use.
// Copy a volume. This is experimental code - do not use.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/content/{volume}
//
type NodesStorageContentCopyRequest struct {
	Node       string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Storage    string `json:"storage,omitempty" url:"storage,omitempty,optional"`         // The storage identifier.
	Target     string `json:"target,omitempty" url:"target,omitempty" validate:"nonzero"` // Target volume identifier
	TargetNode string `json:"target_node,omitempty" url:"target_node,omitempty,optional"` // Target node. Default is local node.
	Volume     string `json:"volume,omitempty" url:"volume,omitempty" validate:"nonzero"` // Source volume identifier
}

// NodesStorageContentUpdateattributesRequest {volume} - Update volume attributes
// Update volume attributes
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/content/{volume}
//
type NodesStorageContentUpdateattributesRequest struct {
	Volume    string `json:"volume,omitempty" url:"volume,omitempty" validate:"nonzero"` // Volume identifier
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Notes     string `json:"notes,omitempty" url:"notes,omitempty,optional"`             // The new notes.
	Protected bool   `json:"protected,omitempty" url:"protected,omitempty,optional"`     // Protection status. Currently only supported for backups.
	Storage   string `json:"storage,omitempty" url:"storage,omitempty,optional"`         // The storage identifier.
}

// NodesStorageFileRestoreListRequest list - List files and directories for single file restore under the given path.
// List files and directories for single file restore under the given path.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/file-restore/list
//
type NodesStorageFileRestoreListRequest struct {
	Filepath string `json:"filepath,omitempty" url:"filepath,omitempty" validate:"nonzero"` // base64-path to the directory or file being listed, or "/".
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Storage  string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`   // The storage identifier.
	Volume   string `json:"volume,omitempty" url:"volume,omitempty" validate:"nonzero"`     // Backup volume ID or name. Currently only PBS snapshots are supported.
}

// NodesStorageFileRestoreListResponse list
// List files and directories for single file restore under the given path.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/file-restore/list
//
type NodesStorageFileRestoreListResponse struct {
	Mtime    int    `json:"mtime,omitempty" url:"mtime,omitempty,optional"`                 // Entry last-modified time (unix timestamp).
	Size     int    `json:"size,omitempty" url:"size,omitempty,optional"`                   // Entry file size.
	Text     string `json:"text,omitempty" url:"text,omitempty" validate:"nonzero"`         // Entry display text.
	Type     string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`         // Entry type.
	Filepath string `json:"filepath,omitempty" url:"filepath,omitempty" validate:"nonzero"` // base64 path of the current entry
	Leaf     bool   `json:"leaf,omitempty" url:"leaf,omitempty" validate:"nonzero"`         // If this entry is a leaf in the directory graph.
}

// NodesStorageFileRestoreDownloadRequest download - Extract a file or directory (as zip archive) from a PBS backup.
// Extract a file or directory (as zip archive) from a PBS backup.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/file-restore/download
//
type NodesStorageFileRestoreDownloadRequest struct {
	Filepath string `json:"filepath,omitempty" url:"filepath,omitempty" validate:"nonzero"` // base64-path to the directory or file to download.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Storage  string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`   // The storage identifier.
	Volume   string `json:"volume,omitempty" url:"volume,omitempty" validate:"nonzero"`     // Backup volume ID or name. Currently only PBS snapshots are supported.
}

// NodesStorageStatusReadRequest status - Read storage status.
// Read storage status.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/status
//
type NodesStorageStatusReadRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
	Storage string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"` // The storage identifier.
}

// NodesStorageRrdRequest rrd - Read storage RRD statistics (returns PNG).
// Read storage RRD statistics (returns PNG).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/rrd
//
type NodesStorageRrdRequest struct {
	Ds        string `json:"ds,omitempty" url:"ds,omitempty" validate:"nonzero"`               // The list of datasources you want to display.
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Storage   string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`     // The storage identifier.
	Timeframe string `json:"timeframe,omitempty" url:"timeframe,omitempty" validate:"nonzero"` // Specify the time frame you are interested in.
	Cf        string `json:"cf,omitempty" url:"cf,omitempty,optional"`                         // The RRD consolidation function
}

// NodesStorageRrdResponse rrd
// Read storage RRD statistics (returns PNG).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/rrd
//
type NodesStorageRrdResponse struct {
	Filename string `json:"filename,omitempty" url:"filename,omitempty" validate:"nonzero"` //
}

// NodesStorageRrddataRequest rrddata - Read storage RRD statistics.
// Read storage RRD statistics.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/rrddata
//
type NodesStorageRrddataRequest struct {
	Cf        string `json:"cf,omitempty" url:"cf,omitempty,optional"`                         // The RRD consolidation function
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Storage   string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`     // The storage identifier.
	Timeframe string `json:"timeframe,omitempty" url:"timeframe,omitempty" validate:"nonzero"` // Specify the time frame you are interested in.
}

// NodesStorageUploadRequest upload - Upload templates and ISO images.
// Upload templates and ISO images.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/upload
//
type NodesStorageUploadRequest struct {
	Checksum          string `json:"checksum,omitempty" url:"checksum,omitempty,optional"`                     // The expected checksum of the file.
	ChecksumAlgorithm string `json:"checksum-algorithm,omitempty" url:"checksum-algorithm,omitempty,optional"` // The algorithm to calculate the checksum of the file.
	Content           string `json:"content,omitempty" url:"content,omitempty" validate:"nonzero"`             // Content type.
	Filename          string `json:"filename,omitempty" url:"filename,omitempty" validate:"nonzero"`           // The name of the file to create. Caution: This will be normalized!
	Node              string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                   // The cluster node name.
	Storage           string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`             // The storage identifier.
	Tmpfilename       string `json:"tmpfilename,omitempty" url:"tmpfilename,omitempty,optional"`               // The source file name. This parameter is usually set by the REST handler. You can only overwrite it when connecting to the trusted port on localhost.
}

// NodesStorageDownloadUrlDownloadUrlRequest download-url - Download templates and ISO images by using an URL.
// Download templates and ISO images by using an URL.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage/{storage}/download-url
//
type NodesStorageDownloadUrlDownloadUrlRequest struct {
	Storage            string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`               // The storage identifier.
	Url                string `json:"url,omitempty" url:"url,omitempty" validate:"nonzero"`                       // The URL to download the file from.
	VerifyCertificates bool   `json:"verify-certificates,omitempty" url:"verify-certificates,omitempty,optional"` // If false, no SSL/TLS certificates will be verified.
	Checksum           string `json:"checksum,omitempty" url:"checksum,omitempty,optional"`                       // The expected checksum of the file.
	ChecksumAlgorithm  string `json:"checksum-algorithm,omitempty" url:"checksum-algorithm,omitempty,optional"`   // The algorithm to calculate the checksum of the file.
	Content            string `json:"content,omitempty" url:"content,omitempty" validate:"nonzero"`               // Content type.
	Filename           string `json:"filename,omitempty" url:"filename,omitempty" validate:"nonzero"`             // The name of the file to create. Caution: This will be normalized!
	Node               string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                     // The cluster node name.
}

// NodesDisksLvmCreateRequest lvm - Create an LVM Volume Group
// Create an LVM Volume Group
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/lvm
//
type NodesDisksLvmCreateRequest struct {
	AddStorage bool   `json:"add_storage,omitempty" url:"add_storage,omitempty,optional"` // Configure storage using the Volume Group
	Device     string `json:"device,omitempty" url:"device,omitempty" validate:"nonzero"` // The block device you want to create the volume group on
	Name       string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`     // The storage identifier.
	Node       string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
}

// NodesDisksLvmDeleteRequest {name} - Remove an LVM Volume Group.
// Remove an LVM Volume Group.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/lvm/{name}
//
type NodesDisksLvmDeleteRequest struct {
	CleanupConfig bool   `json:"cleanup-config,omitempty" url:"cleanup-config,omitempty,optional"` // Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupDisks  bool   `json:"cleanup-disks,omitempty" url:"cleanup-disks,omitempty,optional"`   // Also wipe disks so they can be repurposed afterwards.
	Name          string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`           // The storage identifier.
	Node          string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
}

// NodesDisksLvmthinCreateRequest lvmthin - Create an LVM thinpool
// Create an LVM thinpool
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/lvmthin
//
type NodesDisksLvmthinCreateRequest struct {
	Device     string `json:"device,omitempty" url:"device,omitempty" validate:"nonzero"` // The block device you want to create the thinpool on.
	Name       string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`     // The storage identifier.
	Node       string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	AddStorage bool   `json:"add_storage,omitempty" url:"add_storage,omitempty,optional"` // Configure storage using the thinpool.
}

// NodesDisksLvmthinDeleteRequest {name} - Remove an LVM thin pool.
// Remove an LVM thin pool.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/lvmthin/{name}
//
type NodesDisksLvmthinDeleteRequest struct {
	CleanupConfig bool   `json:"cleanup-config,omitempty" url:"cleanup-config,omitempty,optional"`       // Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupDisks  bool   `json:"cleanup-disks,omitempty" url:"cleanup-disks,omitempty,optional"`         // Also wipe disks so they can be repurposed afterwards.
	Name          string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`                 // The storage identifier.
	Node          string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                 // The cluster node name.
	VolumeGroup   string `json:"volume-group,omitempty" url:"volume-group,omitempty" validate:"nonzero"` // The storage identifier.
}

// NodesDisksDirectoryCreateRequest directory - Create a Filesystem on an unused disk. Will be mounted under '/mnt/pve/NAME'.
// Create a Filesystem on an unused disk. Will be mounted under '/mnt/pve/NAME'.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/directory
//
type NodesDisksDirectoryCreateRequest struct {
	Filesystem string `json:"filesystem,omitempty" url:"filesystem,omitempty,optional"`   // The desired filesystem.
	Name       string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`     // The storage identifier.
	Node       string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	AddStorage bool   `json:"add_storage,omitempty" url:"add_storage,omitempty,optional"` // Configure storage using the directory.
	Device     string `json:"device,omitempty" url:"device,omitempty" validate:"nonzero"` // The block device you want to create the filesystem on.
}

// NodesDisksDirectoryDeleteRequest {name} - Unmounts the storage and removes the mount unit.
// Unmounts the storage and removes the mount unit.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/directory/{name}
//
type NodesDisksDirectoryDeleteRequest struct {
	CleanupDisks  bool   `json:"cleanup-disks,omitempty" url:"cleanup-disks,omitempty,optional"`   // Also wipe disk so it can be repurposed afterwards.
	Name          string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`           // The storage identifier.
	Node          string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	CleanupConfig bool   `json:"cleanup-config,omitempty" url:"cleanup-config,omitempty,optional"` // Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
}

// NodesDisksZfsCreateRequest zfs - Create a ZFS pool.
// Create a ZFS pool.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/zfs
//
type NodesDisksZfsCreateRequest struct {
	Raidlevel   string `json:"raidlevel,omitempty" url:"raidlevel,omitempty" validate:"nonzero"` // The RAID level to use.
	AddStorage  bool   `json:"add_storage,omitempty" url:"add_storage,omitempty,optional"`       // Configure storage using the zpool.
	Ashift      int    `json:"ashift,omitempty" url:"ashift,omitempty,optional"`                 // Pool sector size exponent.
	Compression string `json:"compression,omitempty" url:"compression,omitempty,optional"`       // The compression algorithm to use.
	Devices     string `json:"devices,omitempty" url:"devices,omitempty" validate:"nonzero"`     // The block devices you want to create the zpool on.
	Name        string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`           // The storage identifier.
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
}

// NodesDisksZfsDetailRequest {name} - Get details about a zpool.
// Get details about a zpool.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/zfs/{name}
//
type NodesDisksZfsDetailRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // The storage identifier.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesDisksZfsDetailResponse {name}
// Get details about a zpool.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/zfs/{name}
//
type NodesDisksZfsDetailResponse struct {
	State    string `json:"state,omitempty" url:"state,omitempty" validate:"nonzero"` // The state of the zpool.
	Status   string `json:"status,omitempty" url:"status,omitempty,optional"`         // Information about the state of the zpool.
	Action   string `json:"action,omitempty" url:"action,omitempty,optional"`         // Information about the recommended action to fix the state.
	Children []struct {
	} `json:"children,omitempty" url:"children,omitempty" validate:"nonzero"` // The pool configuration information, including the vdevs for each section (e.g. spares, cache), may be nested.
	Errors string `json:"errors,omitempty" url:"errors,omitempty" validate:"nonzero"` // Information about the errors on the zpool.
	Name   string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`     // The name of the zpool.
	Scan   string `json:"scan,omitempty" url:"scan,omitempty,optional"`               // Information about the last/current scrub.
}

// NodesDisksZfsDeleteRequest {name} - Destroy a ZFS pool.
// Destroy a ZFS pool.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/zfs/{name}
//
type NodesDisksZfsDeleteRequest struct {
	CleanupConfig bool   `json:"cleanup-config,omitempty" url:"cleanup-config,omitempty,optional"` // Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupDisks  bool   `json:"cleanup-disks,omitempty" url:"cleanup-disks,omitempty,optional"`   // Also wipe disks so they can be repurposed afterwards.
	Name          string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"`           // The storage identifier.
	Node          string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
}

// NodesDisksListRequest list - List local disks.
// List local disks.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/list
//
type NodesDisksListRequest struct {
	IncludePartitions bool   `json:"include-partitions,omitempty" url:"include-partitions,omitempty,optional"` // Also include partitions.
	Node              string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                   // The cluster node name.
	Skipsmart         bool   `json:"skipsmart,omitempty" url:"skipsmart,omitempty,optional"`                   // Skip smart checks.
	Type              string `json:"type,omitempty" url:"type,omitempty,optional"`                             // Only list specific types of disks.
}

// NodesDisksListResponse list
// List local disks.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/list
//
type NodesDisksListResponse struct {
	Gpt     bool   `json:"gpt,omitempty" url:"gpt,omitempty" validate:"nonzero"`         //
	Health  string `json:"health,omitempty" url:"health,omitempty,optional"`             //
	Model   string `json:"model,omitempty" url:"model,omitempty,optional"`               //
	Osdid   int    `json:"osdid,omitempty" url:"osdid,omitempty" validate:"nonzero"`     //
	Serial  string `json:"serial,omitempty" url:"serial,omitempty,optional"`             //
	Used    string `json:"used,omitempty" url:"used,omitempty,optional"`                 //
	Vendor  string `json:"vendor,omitempty" url:"vendor,omitempty,optional"`             //
	Devpath string `json:"devpath,omitempty" url:"devpath,omitempty" validate:"nonzero"` // The device path
	Parent  string `json:"parent,omitempty" url:"parent,omitempty,optional"`             // For partitions only. The device path of the disk the partition resides on.
	Size    int    `json:"size,omitempty" url:"size,omitempty" validate:"nonzero"`       //
	Wwn     string `json:"wwn,omitempty" url:"wwn,omitempty,optional"`                   //
}

// NodesDisksSmartRequest smart - Get SMART Health of a disk.
// Get SMART Health of a disk.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/smart
//
type NodesDisksSmartRequest struct {
	Node       string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
	Disk       string `json:"disk,omitempty" url:"disk,omitempty" validate:"nonzero"`   // Block device name
	Healthonly bool   `json:"healthonly,omitempty" url:"healthonly,omitempty,optional"` // If true returns only the health status
}

// NodesDisksSmartResponse smart
// Get SMART Health of a disk.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/smart
//
type NodesDisksSmartResponse struct {
	Attributes []struct {
	} `json:"attributes,omitempty" url:"attributes,omitempty,optional"` //
	Health string `json:"health,omitempty" url:"health,omitempty" validate:"nonzero"` //
	Text   string `json:"text,omitempty" url:"text,omitempty,optional"`               //
	Type   string `json:"type,omitempty" url:"type,omitempty,optional"`               //
}

// NodesDisksInitgptRequest initgpt - Initialize Disk with GPT
// Initialize Disk with GPT
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/initgpt
//
type NodesDisksInitgptRequest struct {
	Disk string `json:"disk,omitempty" url:"disk,omitempty" validate:"nonzero"` // Block device name
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Uuid string `json:"uuid,omitempty" url:"uuid,omitempty,optional"`           // UUID for the GPT table
}

// NodesDisksWipediskWipeDiskRequest wipedisk - Wipe a disk or partition.
// Wipe a disk or partition.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/wipedisk
//
type NodesDisksWipediskWipeDiskRequest struct {
	Disk string `json:"disk,omitempty" url:"disk,omitempty" validate:"nonzero"` // Block device name
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesAptUpdateListsRequest update - List available updates.
// List available updates.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/update
//
type NodesAptUpdateListsRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesAptUpdateDatabaseRequest update - This is used to resynchronize the package index files from their sources (apt-get update).
// This is used to resynchronize the package index files from their sources (apt-get update).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/update
//
type NodesAptUpdateDatabaseRequest struct {
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Notify bool   `json:"notify,omitempty" url:"notify,omitempty,optional"`       // Send notification mail about new packages (to email address specified for user 'root@pam').
	Quiet  bool   `json:"quiet,omitempty" url:"quiet,omitempty,optional"`         // Only produces output suitable for logging, omitting progress indicators.
}

// NodesAptChangelogRequest changelog - Get package changelogs.
// Get package changelogs.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/changelog
//
type NodesAptChangelogRequest struct {
	Version string `json:"version,omitempty" url:"version,omitempty,optional"`     // Package version.
	Name    string `json:"name,omitempty" url:"name,omitempty" validate:"nonzero"` // Package name.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesAptRepositoriesRequest repositories - Get APT repository information.
// Get APT repository information.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/repositories
//
type NodesAptRepositoriesRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesAptRepositoriesResponse repositories
// Get APT repository information.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/repositories
//
type NodesAptRepositoriesResponse struct {
	Digest string `json:"digest,omitempty" url:"digest,omitempty" validate:"nonzero"` // Common digest of all files.
	Errors []struct {
	} `json:"errors,omitempty" url:"errors,omitempty" validate:"nonzero"` // List of problematic repository files.
	Files []struct {
	} `json:"files,omitempty" url:"files,omitempty" validate:"nonzero"` // List of parsed repository files.
	Infos []struct {
	} `json:"infos,omitempty" url:"infos,omitempty" validate:"nonzero"` // Additional information/warnings for APT repositories.
	StandardRepos []struct {
	} `json:"standard-repos,omitempty" url:"standard-repos,omitempty" validate:"nonzero"` // List of standard repositories and their configuration status
}

// NodesAptRepositoriesChangeRepositoryRequest repositories - Change the properties of a repository. Currently only allows enabling/disabling.
// Change the properties of a repository. Currently only allows enabling/disabling.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/repositories
//
type NodesAptRepositoriesChangeRepositoryRequest struct {
	Index   int    `json:"index,omitempty" url:"index,omitempty" validate:"nonzero"` // Index within the file (starting from 0).
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`   // The cluster node name.
	Path    string `json:"path,omitempty" url:"path,omitempty" validate:"nonzero"`   // Path to the containing file.
	Digest  string `json:"digest,omitempty" url:"digest,omitempty,optional"`         // Digest to detect modifications.
	Enabled bool   `json:"enabled,omitempty" url:"enabled,omitempty,optional"`       // Whether the repository should be enabled or not.
}

// NodesAptRepositoriesAddRepositoryRequest repositories - Add a standard repository to the configuration
// Add a standard repository to the configuration
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/repositories
//
type NodesAptRepositoriesAddRepositoryRequest struct {
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`           // Digest to detect modifications.
	Handle string `json:"handle,omitempty" url:"handle,omitempty" validate:"nonzero"` // Handle that identifies a repository.
}

// NodesAptVersionsRequest versions - Get package information for important Proxmox packages.
// Get package information for important Proxmox packages.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/versions
//
type NodesAptVersionsRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesFirewallRulesGetRequest rules - List rules.
// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/rules
//
type NodesFirewallRulesGetRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesFirewallRulesGetResponse rules
// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/rules
//
type NodesFirewallRulesGetResponse struct {
	Pos int `json:"pos,omitempty" url:"pos,omitempty" validate:"nonzero"` //
}

// NodesFirewallRulesCreateRuleRequest rules - Create new rule.
// Create new rule.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/rules
//
type NodesFirewallRulesCreateRuleRequest struct {
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`           // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Dport    string `json:"dport,omitempty" url:"dport,omitempty,optional"`             // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Log      string `json:"log,omitempty" url:"log,omitempty,optional"`                 // Log level for firewall rule.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Pos      int    `json:"pos,omitempty" url:"pos,omitempty,optional"`                 // Update rule at position <pos>.
	Sport    string `json:"sport,omitempty" url:"sport,omitempty,optional"`             // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dest     string `json:"dest,omitempty" url:"dest,omitempty,optional"`               // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Enable   int    `json:"enable,omitempty" url:"enable,omitempty,optional"`           // Flag to enable/disable a rule.
	Iface    string `json:"iface,omitempty" url:"iface,omitempty,optional"`             // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Source   string `json:"source,omitempty" url:"source,omitempty,optional"`           // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Comment  string `json:"comment,omitempty" url:"comment,omitempty,optional"`         // Descriptive comment.
	Macro    string `json:"macro,omitempty" url:"macro,omitempty,optional"`             // Use predefined standard macro.
	Proto    string `json:"proto,omitempty" url:"proto,omitempty,optional"`             // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Action   string `json:"action,omitempty" url:"action,omitempty" validate:"nonzero"` // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	IcmpType string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`     // Specify icmp-type. Only valid if proto equals 'icmp'.
	Type     string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     // Rule type.
}

// NodesFirewallRulesDeleteRuleRequest {pos} - Delete rule.
// Delete rule.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/rules/{pos}
//
type NodesFirewallRulesDeleteRuleRequest struct {
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Pos    int    `json:"pos,omitempty" url:"pos,omitempty,optional"`             // Update rule at position <pos>.
}

// NodesFirewallRulesGetRuleRequest {pos} - Get single rule data.
// Get single rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/rules/{pos}
//
type NodesFirewallRulesGetRuleRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Pos  int    `json:"pos,omitempty" url:"pos,omitempty,optional"`             // Update rule at position <pos>.
}

// NodesFirewallRulesGetRuleResponse {pos}
// Get single rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/rules/{pos}
//
type NodesFirewallRulesGetRuleResponse struct {
	Dest      string `json:"dest,omitempty" url:"dest,omitempty,optional"`               //
	Enable    int    `json:"enable,omitempty" url:"enable,omitempty,optional"`           //
	IcmpType  string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"`     //
	Macro     string `json:"macro,omitempty" url:"macro,omitempty,optional"`             //
	Pos       int    `json:"pos,omitempty" url:"pos,omitempty" validate:"nonzero"`       //
	Sport     string `json:"sport,omitempty" url:"sport,omitempty,optional"`             //
	Action    string `json:"action,omitempty" url:"action,omitempty" validate:"nonzero"` //
	Comment   string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Iface     string `json:"iface,omitempty" url:"iface,omitempty,optional"`             //
	Log       string `json:"log,omitempty" url:"log,omitempty,optional"`                 // Log level for firewall rule
	Dport     string `json:"dport,omitempty" url:"dport,omitempty,optional"`             //
	Ipversion int    `json:"ipversion,omitempty" url:"ipversion,omitempty,optional"`     //
	Proto     string `json:"proto,omitempty" url:"proto,omitempty,optional"`             //
	Source    string `json:"source,omitempty" url:"source,omitempty,optional"`           //
	Type      string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     //
}

// NodesFirewallRulesUpdateRuleRequest {pos} - Modify rule data.
// Modify rule data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/rules/{pos}
//
type NodesFirewallRulesUpdateRuleRequest struct {
	Delete   string `json:"delete,omitempty" url:"delete,omitempty,optional"`       // A list of settings you want to delete.
	Iface    string `json:"iface,omitempty" url:"iface,omitempty,optional"`         // Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Macro    string `json:"macro,omitempty" url:"macro,omitempty,optional"`         // Use predefined standard macro.
	Sport    string `json:"sport,omitempty" url:"sport,omitempty,optional"`         // Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dest     string `json:"dest,omitempty" url:"dest,omitempty,optional"`           // Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Moveto   int    `json:"moveto,omitempty" url:"moveto,omitempty,optional"`       // Move rule to new position <moveto>. Other arguments are ignored.
	Pos      int    `json:"pos,omitempty" url:"pos,omitempty,optional"`             // Update rule at position <pos>.
	Proto    string `json:"proto,omitempty" url:"proto,omitempty,optional"`         // IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Source   string `json:"source,omitempty" url:"source,omitempty,optional"`       // Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Comment  string `json:"comment,omitempty" url:"comment,omitempty,optional"`     // Descriptive comment.
	Digest   string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Log      string `json:"log,omitempty" url:"log,omitempty,optional"`             // Log level for firewall rule.
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Action   string `json:"action,omitempty" url:"action,omitempty,optional"`       // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Dport    string `json:"dport,omitempty" url:"dport,omitempty,optional"`         // Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Enable   int    `json:"enable,omitempty" url:"enable,omitempty,optional"`       // Flag to enable/disable a rule.
	IcmpType string `json:"icmp-type,omitempty" url:"icmp-type,omitempty,optional"` // Specify icmp-type. Only valid if proto equals 'icmp'.
	Type     string `json:"type,omitempty" url:"type,omitempty,optional"`           // Rule type.
}

// NodesFirewallOptionsGetRequest options - Get host firewall options.
// Get host firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/options
//
type NodesFirewallOptionsGetRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesFirewallOptionsGetResponse options
// Get host firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/options
//
type NodesFirewallOptionsGetResponse struct {
	LogLevelOut                      string `json:"log_level_out,omitempty" url:"log_level_out,omitempty,optional"`                                               // Log level for outgoing traffic.
	TcpFlagsLogLevel                 string `json:"tcp_flags_log_level,omitempty" url:"tcp_flags_log_level,omitempty,optional"`                                   // Log level for illegal tcp flags filter.
	Tcpflags                         bool   `json:"tcpflags,omitempty" url:"tcpflags,omitempty,optional"`                                                         // Filter illegal combinations of TCP flags.
	Enable                           bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`                                                             // Enable host firewall rules.
	LogNfConntrack                   bool   `json:"log_nf_conntrack,omitempty" url:"log_nf_conntrack,omitempty,optional"`                                         // Enable logging of conntrack information.
	NfConntrackMax                   int    `json:"nf_conntrack_max,omitempty" url:"nf_conntrack_max,omitempty,optional"`                                         // Maximum number of tracked connections.
	Nosmurfs                         bool   `json:"nosmurfs,omitempty" url:"nosmurfs,omitempty,optional"`                                                         // Enable SMURFS filter.
	ProtectionSynfloodBurst          int    `json:"protection_synflood_burst,omitempty" url:"protection_synflood_burst,omitempty,optional"`                       // Synflood protection rate burst by ip src.
	ProtectionSynfloodRate           int    `json:"protection_synflood_rate,omitempty" url:"protection_synflood_rate,omitempty,optional"`                         // Synflood protection rate syn/sec by ip src.
	LogLevelIn                       string `json:"log_level_in,omitempty" url:"log_level_in,omitempty,optional"`                                                 // Log level for incoming traffic.
	NfConntrackAllowInvalid          bool   `json:"nf_conntrack_allow_invalid,omitempty" url:"nf_conntrack_allow_invalid,omitempty,optional"`                     // Allow invalid packets on connection tracking.
	NfConntrackTcpTimeoutEstablished int    `json:"nf_conntrack_tcp_timeout_established,omitempty" url:"nf_conntrack_tcp_timeout_established,omitempty,optional"` // Conntrack established timeout.
	NfConntrackTcpTimeoutSynRecv     int    `json:"nf_conntrack_tcp_timeout_syn_recv,omitempty" url:"nf_conntrack_tcp_timeout_syn_recv,omitempty,optional"`       // Conntrack syn recv timeout.
	SmurfLogLevel                    string `json:"smurf_log_level,omitempty" url:"smurf_log_level,omitempty,optional"`                                           // Log level for SMURFS filter.
	Ndp                              bool   `json:"ndp,omitempty" url:"ndp,omitempty,optional"`                                                                   // Enable NDP (Neighbor Discovery Protocol).
	ProtectionSynflood               bool   `json:"protection_synflood,omitempty" url:"protection_synflood,omitempty,optional"`                                   // Enable synflood protection
}

// NodesFirewallOptionsSetRequest options - Set Firewall options.
// Set Firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/options
//
type NodesFirewallOptionsSetRequest struct {
	NfConntrackAllowInvalid          bool   `json:"nf_conntrack_allow_invalid,omitempty" url:"nf_conntrack_allow_invalid,omitempty,optional"`                     // Allow invalid packets on connection tracking.
	NfConntrackMax                   int    `json:"nf_conntrack_max,omitempty" url:"nf_conntrack_max,omitempty,optional"`                                         // Maximum number of tracked connections.
	Node                             string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                                                       // The cluster node name.
	Tcpflags                         bool   `json:"tcpflags,omitempty" url:"tcpflags,omitempty,optional"`                                                         // Filter illegal combinations of TCP flags.
	Digest                           string `json:"digest,omitempty" url:"digest,omitempty,optional"`                                                             // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	LogLevelOut                      string `json:"log_level_out,omitempty" url:"log_level_out,omitempty,optional"`                                               // Log level for outgoing traffic.
	NfConntrackTcpTimeoutEstablished int    `json:"nf_conntrack_tcp_timeout_established,omitempty" url:"nf_conntrack_tcp_timeout_established,omitempty,optional"` // Conntrack established timeout.
	ProtectionSynfloodBurst          int    `json:"protection_synflood_burst,omitempty" url:"protection_synflood_burst,omitempty,optional"`                       // Synflood protection rate burst by ip src.
	ProtectionSynfloodRate           int    `json:"protection_synflood_rate,omitempty" url:"protection_synflood_rate,omitempty,optional"`                         // Synflood protection rate syn/sec by ip src.
	SmurfLogLevel                    string `json:"smurf_log_level,omitempty" url:"smurf_log_level,omitempty,optional"`                                           // Log level for SMURFS filter.
	Enable                           bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`                                                             // Enable host firewall rules.
	LogLevelIn                       string `json:"log_level_in,omitempty" url:"log_level_in,omitempty,optional"`                                                 // Log level for incoming traffic.
	LogNfConntrack                   bool   `json:"log_nf_conntrack,omitempty" url:"log_nf_conntrack,omitempty,optional"`                                         // Enable logging of conntrack information.
	NfConntrackTcpTimeoutSynRecv     int    `json:"nf_conntrack_tcp_timeout_syn_recv,omitempty" url:"nf_conntrack_tcp_timeout_syn_recv,omitempty,optional"`       // Conntrack syn recv timeout.
	Delete                           string `json:"delete,omitempty" url:"delete,omitempty,optional"`                                                             // A list of settings you want to delete.
	Ndp                              bool   `json:"ndp,omitempty" url:"ndp,omitempty,optional"`                                                                   // Enable NDP (Neighbor Discovery Protocol).
	Nosmurfs                         bool   `json:"nosmurfs,omitempty" url:"nosmurfs,omitempty,optional"`                                                         // Enable SMURFS filter.
	ProtectionSynflood               bool   `json:"protection_synflood,omitempty" url:"protection_synflood,omitempty,optional"`                                   // Enable synflood protection
	TcpFlagsLogLevel                 string `json:"tcp_flags_log_level,omitempty" url:"tcp_flags_log_level,omitempty,optional"`                                   // Log level for illegal tcp flags filter.
}

// NodesFirewallLogRequest log - Read firewall log
// Read firewall log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/log
//
type NodesFirewallLogRequest struct {
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Start int    `json:"start,omitempty" url:"start,omitempty,optional"`         //
	Limit int    `json:"limit,omitempty" url:"limit,omitempty,optional"`         //
}

// NodesFirewallLogResponse log
// Read firewall log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/firewall/log
//
type NodesFirewallLogResponse struct {
	N int    `json:"n,omitempty" url:"n,omitempty" validate:"nonzero"` // Line number
	T string `json:"t,omitempty" url:"t,omitempty" validate:"nonzero"` // Line text
}

// NodesReplicationStatusRequest replication - List status of all replication jobs on this node.
// List status of all replication jobs on this node.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/replication
//
type NodesReplicationStatusRequest struct {
	Guest int    `json:"guest,omitempty" url:"guest,omitempty,optional"`         // Only list replication jobs for this guest.
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesReplicationStatusResponse replication
// List status of all replication jobs on this node.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/replication
//
type NodesReplicationStatusResponse struct {
	Id string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"` //
}

// NodesReplicationStatusJobRequest status - Get replication job status.
// Get replication job status.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/replication/{id}/status
//
type NodesReplicationStatusJobRequest struct {
	Id   string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`     // Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesReplicationLogReadJobRequest log - Read replication job log.
// Read replication job log.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/replication/{id}/log
//
type NodesReplicationLogReadJobRequest struct {
	Id    string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`     // Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	Limit int    `json:"limit,omitempty" url:"limit,omitempty,optional"`         //
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Start int    `json:"start,omitempty" url:"start,omitempty,optional"`         //
}

// NodesReplicationLogReadJobResponse log
// Read replication job log.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/replication/{id}/log
//
type NodesReplicationLogReadJobResponse struct {
	T string `json:"t,omitempty" url:"t,omitempty" validate:"nonzero"` // Line text
	N int    `json:"n,omitempty" url:"n,omitempty" validate:"nonzero"` // Line number
}

// NodesReplicationScheduleNowRequest schedule_now - Schedule replication job to start as soon as possible.
// Schedule replication job to start as soon as possible.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/replication/{id}/schedule_now
//
type NodesReplicationScheduleNowRequest struct {
	Id   string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`     // Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCertificatesAcmeCertificateRevokeRequest certificate - Revoke existing certificate from CA.
// Revoke existing certificate from CA.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/certificates/acme/certificate
//
type NodesCertificatesAcmeCertificateRevokeRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCertificatesAcmeCertificateNewRequest certificate - Order a new certificate from ACME-compatible CA.
// Order a new certificate from ACME-compatible CA.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/certificates/acme/certificate
//
type NodesCertificatesAcmeCertificateNewRequest struct {
	Force bool   `json:"force,omitempty" url:"force,omitempty,optional"`         // Overwrite existing custom certificate.
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCertificatesAcmeCertificateRenewRequest certificate - Renew existing certificate from CA.
// Renew existing certificate from CA.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/certificates/acme/certificate
//
type NodesCertificatesAcmeCertificateRenewRequest struct {
	Force bool   `json:"force,omitempty" url:"force,omitempty,optional"`         // Force renewal even if expiry is more than 30 days away.
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCertificatesInfoRequest info - Get information about node's certificates.
// Get information about node's certificates.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/certificates/info
//
type NodesCertificatesInfoRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesCertificatesInfoResponse info
// Get information about node's certificates.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/certificates/info
//
type NodesCertificatesInfoResponse struct {
	Notbefore     int    `json:"notbefore,omitempty" url:"notbefore,omitempty,optional"`             // Certificate's notBefore timestamp (UNIX epoch).
	Pem           string `json:"pem,omitempty" url:"pem,omitempty,optional"`                         // Certificate in PEM format
	PublicKeyBits int    `json:"public-key-bits,omitempty" url:"public-key-bits,omitempty,optional"` // Certificate's public key size
	PublicKeyType string `json:"public-key-type,omitempty" url:"public-key-type,omitempty,optional"` // Certificate's public key algorithm
	Fingerprint   string `json:"fingerprint,omitempty" url:"fingerprint,omitempty,optional"`         // Certificate SHA 256 fingerprint.
	Issuer        string `json:"issuer,omitempty" url:"issuer,omitempty,optional"`                   // Certificate issuer name.
	Notafter      int    `json:"notafter,omitempty" url:"notafter,omitempty,optional"`               // Certificate's notAfter timestamp (UNIX epoch).
	San           []struct {
	} `json:"san,omitempty" url:"san,omitempty,optional"` // List of Certificate's SubjectAlternativeName entries.
	Subject  string `json:"subject,omitempty" url:"subject,omitempty,optional"`   // Certificate subject name.
	Filename string `json:"filename,omitempty" url:"filename,omitempty,optional"` //
}

// NodesCertificatesCustomRemoveCertRequest custom - DELETE custom certificate chain and key.
// DELETE custom certificate chain and key.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/certificates/custom
//
type NodesCertificatesCustomRemoveCertRequest struct {
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Restart bool   `json:"restart,omitempty" url:"restart,omitempty,optional"`     // Restart pveproxy.
}

// NodesCertificatesCustomUploadCertRequest custom - Upload or update custom certificate chain and key.
// Upload or update custom certificate chain and key.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/certificates/custom
//
type NodesCertificatesCustomUploadCertRequest struct {
	Key          string `json:"key,omitempty" url:"key,omitempty,optional"`                             // PEM encoded private key.
	Node         string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                 // The cluster node name.
	Restart      bool   `json:"restart,omitempty" url:"restart,omitempty,optional"`                     // Restart pveproxy.
	Certificates string `json:"certificates,omitempty" url:"certificates,omitempty" validate:"nonzero"` // PEM encoded certificate (chain).
	Force        bool   `json:"force,omitempty" url:"force,omitempty,optional"`                         // Overwrite existing custom or ACME certificate files.
}

// NodesCertificatesCustomUploadCertResponse custom
// Upload or update custom certificate chain and key.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/certificates/custom
//
type NodesCertificatesCustomUploadCertResponse struct {
	Notbefore     int `json:"notbefore,omitempty" url:"notbefore,omitempty,optional"`             // Certificate's notBefore timestamp (UNIX epoch).
	PublicKeyBits int `json:"public-key-bits,omitempty" url:"public-key-bits,omitempty,optional"` // Certificate's public key size
	San           []struct {
	} `json:"san,omitempty" url:"san,omitempty,optional"` // List of Certificate's SubjectAlternativeName entries.
	Subject       string `json:"subject,omitempty" url:"subject,omitempty,optional"`                 // Certificate subject name.
	Issuer        string `json:"issuer,omitempty" url:"issuer,omitempty,optional"`                   // Certificate issuer name.
	Notafter      int    `json:"notafter,omitempty" url:"notafter,omitempty,optional"`               // Certificate's notAfter timestamp (UNIX epoch).
	Pem           string `json:"pem,omitempty" url:"pem,omitempty,optional"`                         // Certificate in PEM format
	PublicKeyType string `json:"public-key-type,omitempty" url:"public-key-type,omitempty,optional"` // Certificate's public key algorithm
	Filename      string `json:"filename,omitempty" url:"filename,omitempty,optional"`               //
	Fingerprint   string `json:"fingerprint,omitempty" url:"fingerprint,omitempty,optional"`         // Certificate SHA 256 fingerprint.
}

// NodesConfigGetRequest config - Get node configuration options.
// Get node configuration options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/config
//
type NodesConfigGetRequest struct {
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Property string `json:"property,omitempty" url:"property,omitempty,optional"`   // Return only a specific property from the node configuration.
}

// NodesConfigGetResponse config
// Get node configuration options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/config
//
type NodesConfigGetResponse struct {
	Acme                string   `json:"acme,omitempty" url:"acme,omitempty,optional"`                                   // Node specific ACME settings.
	Acmedomain          []string `json:"acmedomain[n],omitempty" url:"acmedomain[n],omitempty,optional"`                 // ACME domain and validation plugin
	Description         string   `json:"description,omitempty" url:"description,omitempty,optional"`                     // Description for the Node. Shown in the web-interface node notes panel. This is saved as comment inside the configuration file.
	Digest              string   `json:"digest,omitempty" url:"digest,omitempty,optional"`                               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	StartallOnbootDelay int      `json:"startall-onboot-delay,omitempty" url:"startall-onboot-delay,omitempty,optional"` // Initial delay in seconds, before starting all the Virtual Guests with on-boot enabled.
	Wakeonlan           string   `json:"wakeonlan,omitempty" url:"wakeonlan,omitempty,optional"`                         // MAC address for wake on LAN
}

// NodesConfigSetOptionsRequest config - Set node configuration options.
// Set node configuration options.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/config
//
type NodesConfigSetOptionsRequest struct {
	Acme                string   `json:"acme,omitempty" url:"acme,omitempty,optional"`                                   // Node specific ACME settings.
	Acmedomain          []string `json:"acmedomain[n],omitempty" url:"acmedomain[n],omitempty,optional"`                 // ACME domain and validation plugin
	Delete              string   `json:"delete,omitempty" url:"delete,omitempty,optional"`                               // A list of settings you want to delete.
	Description         string   `json:"description,omitempty" url:"description,omitempty,optional"`                     // Description for the Node. Shown in the web-interface node notes panel. This is saved as comment inside the configuration file.
	Digest              string   `json:"digest,omitempty" url:"digest,omitempty,optional"`                               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Node                string   `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                         // The cluster node name.
	StartallOnbootDelay int      `json:"startall-onboot-delay,omitempty" url:"startall-onboot-delay,omitempty,optional"` // Initial delay in seconds, before starting all the Virtual Guests with on-boot enabled.
	Wakeonlan           string   `json:"wakeonlan,omitempty" url:"wakeonlan,omitempty,optional"`                         // MAC address for wake on LAN
}

// NodesSdnIndexRequest sdn - SDN index.
// SDN index.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/sdn
//
type NodesSdnIndexRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesSdnZonesDiridxRequest {zone} -
//
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/sdn/zones/{zone}
//
type NodesSdnZonesDiridxRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Zone string `json:"zone,omitempty" url:"zone,omitempty" validate:"nonzero"` // The SDN zone object identifier.
}

// NodesSdnZonesDiridxResponse {zone}
//
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/sdn/zones/{zone}
//
type NodesSdnZonesDiridxResponse struct {
	Subdir string `json:"subdir,omitempty" url:"subdir,omitempty" validate:"nonzero"` //
}

// NodesVersionRequest version - API version details
// API version details
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/version
//
type NodesVersionRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesVersionResponse version
// API version details
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/version
//
type NodesVersionResponse struct {
	Release string `json:"release,omitempty" url:"release,omitempty" validate:"nonzero"` // The current installed Proxmox VE Release
	Repoid  string `json:"repoid,omitempty" url:"repoid,omitempty" validate:"nonzero"`   // The short git commit hash ID from which this version was build
	Version string `json:"version,omitempty" url:"version,omitempty" validate:"nonzero"` // The current installed pve-manager package version
}

// NodesStatusRequest status - Read node status
// Read node status
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/status
//
type NodesStatusRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesStatusNodeCmdRequest status - Reboot or shutdown a node.
// Reboot or shutdown a node.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/status
//
type NodesStatusNodeCmdRequest struct {
	Command string `json:"command,omitempty" url:"command,omitempty" validate:"nonzero"` // Specify the command.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`       // The cluster node name.
}

// NodesNetstatRequest netstat - Read tap/vm network device interface counters
// Read tap/vm network device interface counters
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/netstat
//
type NodesNetstatRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesExecuteRequest execute - Execute multiple commands in order.
// Execute multiple commands in order.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/execute
//
type NodesExecuteRequest struct {
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Commands string `json:"commands,omitempty" url:"commands,omitempty" validate:"nonzero"` // JSON encoded array of commands.
}

// NodesWakeonlanRequest wakeonlan - Try to wake a node via 'wake on LAN' network packet.
// Try to wake a node via 'wake on LAN' network packet.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/wakeonlan
//
type NodesWakeonlanRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // target node for wake on LAN packet
}

// NodesRrdRequest rrd - Read node RRD statistics (returns PNG)
// Read node RRD statistics (returns PNG)
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/rrd
//
type NodesRrdRequest struct {
	Ds        string `json:"ds,omitempty" url:"ds,omitempty" validate:"nonzero"`               // The list of datasources you want to display.
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Timeframe string `json:"timeframe,omitempty" url:"timeframe,omitempty" validate:"nonzero"` // Specify the time frame you are interested in.
	Cf        string `json:"cf,omitempty" url:"cf,omitempty,optional"`                         // The RRD consolidation function
}

// NodesRrdResponse rrd
// Read node RRD statistics (returns PNG)
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/rrd
//
type NodesRrdResponse struct {
	Filename string `json:"filename,omitempty" url:"filename,omitempty" validate:"nonzero"` //
}

// NodesRrddataRequest rrddata - Read node RRD statistics
// Read node RRD statistics
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/rrddata
//
type NodesRrddataRequest struct {
	Cf        string `json:"cf,omitempty" url:"cf,omitempty,optional"`                         // The RRD consolidation function
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Timeframe string `json:"timeframe,omitempty" url:"timeframe,omitempty" validate:"nonzero"` // Specify the time frame you are interested in.
}

// NodesSyslogRequest syslog - Read system log
// Read system log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/syslog
//
type NodesSyslogRequest struct {
	Limit   int    `json:"limit,omitempty" url:"limit,omitempty,optional"`         //
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Service string `json:"service,omitempty" url:"service,omitempty,optional"`     // Service ID
	Since   string `json:"since,omitempty" url:"since,omitempty,optional"`         // Display all log since this date-time string.
	Start   int    `json:"start,omitempty" url:"start,omitempty,optional"`         //
	Until   string `json:"until,omitempty" url:"until,omitempty,optional"`         // Display all log until this date-time string.
}

// NodesSyslogResponse syslog
// Read system log
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/syslog
//
type NodesSyslogResponse struct {
	N int    `json:"n,omitempty" url:"n,omitempty" validate:"nonzero"` // Line number
	T string `json:"t,omitempty" url:"t,omitempty" validate:"nonzero"` // Line text
}

// NodesJournalRequest journal - Read Journal
// Read Journal
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/journal
//
type NodesJournalRequest struct {
	Node        string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Since       int    `json:"since,omitempty" url:"since,omitempty,optional"`             // Display all log since this UNIX epoch. Conflicts with 'startcursor'.
	Startcursor string `json:"startcursor,omitempty" url:"startcursor,omitempty,optional"` // Start after the given Cursor. Conflicts with 'since'
	Until       int    `json:"until,omitempty" url:"until,omitempty,optional"`             // Display all log until this UNIX epoch. Conflicts with 'endcursor'.
	Endcursor   string `json:"endcursor,omitempty" url:"endcursor,omitempty,optional"`     // End before the given Cursor. Conflicts with 'until'
	Lastentries int    `json:"lastentries,omitempty" url:"lastentries,omitempty,optional"` // Limit to the last X lines. Conflicts with a range.
}

// NodesVncshellRequest vncshell - Creates a VNC Shell proxy.
// Creates a VNC Shell proxy.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/vncshell
//
type NodesVncshellRequest struct {
	Width     int    `json:"width,omitempty" url:"width,omitempty,optional"`         // sets the width of the console in pixels.
	Cmd       string `json:"cmd,omitempty" url:"cmd,omitempty,optional"`             // Run specific command or default to login.
	CmdOpts   string `json:"cmd-opts,omitempty" url:"cmd-opts,omitempty,optional"`   // Add parameters to a command. Encoded as null terminated strings.
	Height    int    `json:"height,omitempty" url:"height,omitempty,optional"`       // sets the height of the console in pixels.
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Websocket bool   `json:"websocket,omitempty" url:"websocket,omitempty,optional"` // use websocket instead of standard vnc.
}

// NodesVncshellResponse vncshell
// Creates a VNC Shell proxy.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/vncshell
//
type NodesVncshellResponse struct {
	Cert   string `json:"cert,omitempty" url:"cert,omitempty" validate:"nonzero"`     //
	Port   int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`     //
	Ticket string `json:"ticket,omitempty" url:"ticket,omitempty" validate:"nonzero"` //
	Upid   string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"`     //
	User   string `json:"user,omitempty" url:"user,omitempty" validate:"nonzero"`     //
}

// NodesTermproxyRequest termproxy - Creates a VNC Shell proxy.
// Creates a VNC Shell proxy.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/termproxy
//
type NodesTermproxyRequest struct {
	Cmd     string `json:"cmd,omitempty" url:"cmd,omitempty,optional"`             // Run specific command or default to login.
	CmdOpts string `json:"cmd-opts,omitempty" url:"cmd-opts,omitempty,optional"`   // Add parameters to a command. Encoded as null terminated strings.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesTermproxyResponse termproxy
// Creates a VNC Shell proxy.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/termproxy
//
type NodesTermproxyResponse struct {
	Port   int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`     //
	Ticket string `json:"ticket,omitempty" url:"ticket,omitempty" validate:"nonzero"` //
	Upid   string `json:"upid,omitempty" url:"upid,omitempty" validate:"nonzero"`     //
	User   string `json:"user,omitempty" url:"user,omitempty" validate:"nonzero"`     //
}

// NodesVncwebsocketRequest vncwebsocket - Opens a websocket for VNC traffic.
// Opens a websocket for VNC traffic.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/vncwebsocket
//
type NodesVncwebsocketRequest struct {
	Node      string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`           // The cluster node name.
	Port      int    `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"`           // Port number returned by previous vncproxy call.
	Vncticket string `json:"vncticket,omitempty" url:"vncticket,omitempty" validate:"nonzero"` // Ticket from previous call to vncproxy.
}

// NodesVncwebsocketResponse vncwebsocket
// Opens a websocket for VNC traffic.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/vncwebsocket
//
type NodesVncwebsocketResponse struct {
	Port string `json:"port,omitempty" url:"port,omitempty" validate:"nonzero"` //
}

// NodesSpiceshellRequest spiceshell - Creates a SPICE shell.
// Creates a SPICE shell.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/spiceshell
//
type NodesSpiceshellRequest struct {
	Cmd     string `json:"cmd,omitempty" url:"cmd,omitempty,optional"`             // Run specific command or default to login.
	CmdOpts string `json:"cmd-opts,omitempty" url:"cmd-opts,omitempty,optional"`   // Add parameters to a command. Encoded as null terminated strings.
	Node    string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Proxy   string `json:"proxy,omitempty" url:"proxy,omitempty,optional"`         // SPICE proxy server. This can be used by the client to specify the proxy server. All nodes in a cluster runs 'spiceproxy', so it is up to the client to choose one. By default, we return the node where the VM is currently running. As reasonable setting is to use same node you use to connect to the API (This is window.location.hostname for the JS GUI).
}

// NodesSpiceshellResponse spiceshell
// Creates a SPICE shell.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/spiceshell
//
type NodesSpiceshellResponse struct {
	Host     string `json:"host,omitempty" url:"host,omitempty" validate:"nonzero"`         //
	Password string `json:"password,omitempty" url:"password,omitempty" validate:"nonzero"` //
	Proxy    string `json:"proxy,omitempty" url:"proxy,omitempty" validate:"nonzero"`       //
	TlsPort  int    `json:"tls-port,omitempty" url:"tls-port,omitempty" validate:"nonzero"` //
	Type     string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`         //
}

// NodesDnsRequest dns - Read DNS settings.
// Read DNS settings.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/dns
//
type NodesDnsRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesDnsResponse dns
// Read DNS settings.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/dns
//
type NodesDnsResponse struct {
	Search string `json:"search,omitempty" url:"search,omitempty,optional"` // Search domain for host-name lookup.
	Dns1   string `json:"dns1,omitempty" url:"dns1,omitempty,optional"`     // First name server IP address.
	Dns2   string `json:"dns2,omitempty" url:"dns2,omitempty,optional"`     // Second name server IP address.
	Dns3   string `json:"dns3,omitempty" url:"dns3,omitempty,optional"`     // Third name server IP address.
}

// NodesDnsUpdateRequest dns - Write DNS settings.
// Write DNS settings.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/dns
//
type NodesDnsUpdateRequest struct {
	Dns1   string `json:"dns1,omitempty" url:"dns1,omitempty,optional"`               // First name server IP address.
	Dns2   string `json:"dns2,omitempty" url:"dns2,omitempty,optional"`               // Second name server IP address.
	Dns3   string `json:"dns3,omitempty" url:"dns3,omitempty,optional"`               // Third name server IP address.
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`     // The cluster node name.
	Search string `json:"search,omitempty" url:"search,omitempty" validate:"nonzero"` // Search domain for host-name lookup.
}

// NodesTimeRequest time - Read server time and time zone settings.
// Read server time and time zone settings.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/time
//
type NodesTimeRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesTimeResponse time
// Read server time and time zone settings.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/time
//
type NodesTimeResponse struct {
	Localtime int    `json:"localtime,omitempty" url:"localtime,omitempty" validate:"nonzero"` // Seconds since 1970-01-01 00:00:00 (local time)
	Time      int    `json:"time,omitempty" url:"time,omitempty" validate:"nonzero"`           // Seconds since 1970-01-01 00:00:00 UTC.
	Timezone  string `json:"timezone,omitempty" url:"timezone,omitempty" validate:"nonzero"`   // Time zone
}

// NodesTimeSetzoneRequest time - Set time zone.
// Set time zone.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/time
//
type NodesTimeSetzoneRequest struct {
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Timezone string `json:"timezone,omitempty" url:"timezone,omitempty" validate:"nonzero"` // Time zone. The file '/usr/share/zoneinfo/zone.tab' contains the list of valid names.
}

// NodesAplinfoRequest aplinfo - Get list of appliances.
// Get list of appliances.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/aplinfo
//
type NodesAplinfoRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesAplinfoAplDownloadRequest aplinfo - Download appliance templates.
// Download appliance templates.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/aplinfo
//
type NodesAplinfoAplDownloadRequest struct {
	Node     string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`         // The cluster node name.
	Storage  string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`   // The storage where the template will be stored
	Template string `json:"template,omitempty" url:"template,omitempty" validate:"nonzero"` // The template which will downloaded
}

// NodesQueryUrlMetadataQueryUrlMetadataRequest query-url-metadata - Query metadata of an URL: file size, file name and mime type.
// Query metadata of an URL: file size, file name and mime type.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/query-url-metadata
//
type NodesQueryUrlMetadataQueryUrlMetadataRequest struct {
	Node               string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`                     // The cluster node name.
	Url                string `json:"url,omitempty" url:"url,omitempty" validate:"nonzero"`                       // The URL to query the metadata from.
	VerifyCertificates bool   `json:"verify-certificates,omitempty" url:"verify-certificates,omitempty,optional"` // If false, no SSL/TLS certificates will be verified.
}

// NodesQueryUrlMetadataQueryUrlMetadataResponse query-url-metadata
// Query metadata of an URL: file size, file name and mime type.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/query-url-metadata
//
type NodesQueryUrlMetadataQueryUrlMetadataResponse struct {
	Filename string `json:"filename,omitempty" url:"filename,omitempty,optional"` //
	Mimetype string `json:"mimetype,omitempty" url:"mimetype,omitempty,optional"` //
	Size     int    `json:"size,omitempty" url:"size,omitempty,optional"`         //
}

// NodesReportRequest report - Gather various systems information about a node
// Gather various systems information about a node
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/report
//
type NodesReportRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesStartallRequest startall - Start all VMs and containers located on this node (by default only those with onboot=1).
// Start all VMs and containers located on this node (by default only those with onboot=1).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/startall
//
type NodesStartallRequest struct {
	Force bool   `json:"force,omitempty" url:"force,omitempty,optional"`         // Issue start command even if virtual guest have 'onboot' not set or set to off.
	Node  string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
	Vms   string `json:"vms,omitempty" url:"vms,omitempty,optional"`             // Only consider guests from this comma separated list of VMIDs.
}

// NodesStopallRequest stopall - Stop all VMs and Containers.
// Stop all VMs and Containers.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/stopall
//
type NodesStopallRequest struct {
	Vms  string `json:"vms,omitempty" url:"vms,omitempty,optional"`             // Only consider Guests with these IDs.
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesMigrateallRequest migrateall - Migrate all VMs and Containers.
// Migrate all VMs and Containers.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/migrateall
//
type NodesMigrateallRequest struct {
	Target         string `json:"target,omitempty" url:"target,omitempty" validate:"nonzero"`           // Target node.
	Vms            string `json:"vms,omitempty" url:"vms,omitempty,optional"`                           // Only consider Guests with these IDs.
	WithLocalDisks bool   `json:"with-local-disks,omitempty" url:"with-local-disks,omitempty,optional"` // Enable live storage migration for local disk
	Maxworkers     int    `json:"maxworkers,omitempty" url:"maxworkers,omitempty,optional"`             // Maximal number of parallel migration job. If not set use 'max_workers' from datacenter.cfg, one of both must be set!
	Node           string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"`               // The cluster node name.
}

// NodesHostsGetEtcRequest hosts - Get the content of /etc/hosts.
// Get the content of /etc/hosts.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hosts
//
type NodesHostsGetEtcRequest struct {
	Node string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// NodesHostsGetEtcResponse hosts
// Get the content of /etc/hosts.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hosts
//
type NodesHostsGetEtcResponse struct {
	Data   string `json:"data,omitempty" url:"data,omitempty" validate:"nonzero"` // The content of /etc/hosts.
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
}

// NodesHostsWriteEtcRequest hosts - Write /etc/hosts.
// Write /etc/hosts.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/hosts
//
type NodesHostsWriteEtcRequest struct {
	Data   string `json:"data,omitempty" url:"data,omitempty" validate:"nonzero"` // The target content of /etc/hosts.
	Digest string `json:"digest,omitempty" url:"digest,omitempty,optional"`       // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Node   string `json:"node,omitempty" url:"node,omitempty" validate:"nonzero"` // The cluster node name.
}

// StorageCreateRequest storage - Create a new storage.
// Create a new storage.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/storage
//
type StorageCreateRequest struct {
	FsName               string `json:"fs-name,omitempty" url:"fs-name,omitempty,optional"`                             // The Ceph filesystem name.
	MasterPubkey         string `json:"master-pubkey,omitempty" url:"master-pubkey,omitempty,optional"`                 // Base64-encoded, PEM-formatted public RSA key. Used to encrypt a copy of the encryption-key which will be added to each encrypted backup.
	Monhost              string `json:"monhost,omitempty" url:"monhost,omitempty,optional"`                             // IP addresses of monitors (for external clusters).
	Path                 string `json:"path,omitempty" url:"path,omitempty,optional"`                                   // File system path.
	Server2              string `json:"server2,omitempty" url:"server2,omitempty,optional"`                             // Backup volfile server IP or DNS name.
	Options              string `json:"options,omitempty" url:"options,omitempty,optional"`                             // NFS mount options (see 'man nfs')
	Portal               string `json:"portal,omitempty" url:"portal,omitempty,optional"`                               // iSCSI portal (IP or DNS name with optional port).
	PruneBackups         string `json:"prune-backups,omitempty" url:"prune-backups,omitempty,optional"`                 // The retention options with shorter intervals are processed first with --keep-last being the very first one. Each option covers a specific period of time. We say that backups within this period are covered by this option. The next option does not take care of already covered backups and only considers older backups.
	Authsupported        string `json:"authsupported,omitempty" url:"authsupported,omitempty,optional"`                 // Authsupported.
	Blocksize            string `json:"blocksize,omitempty" url:"blocksize,omitempty,optional"`                         // block size
	EncryptionKey        string `json:"encryption-key,omitempty" url:"encryption-key,omitempty,optional"`               // Encryption key. Use 'autogen' to generate one automatically without passphrase.
	Fingerprint          string `json:"fingerprint,omitempty" url:"fingerprint,omitempty,optional"`                     // Certificate SHA 256 fingerprint.
	Format               string `json:"format,omitempty" url:"format,omitempty,optional"`                               // Default image format.
	ComstarTg            string `json:"comstar_tg,omitempty" url:"comstar_tg,omitempty,optional"`                       // target group for comstar views
	Nocow                bool   `json:"nocow,omitempty" url:"nocow,omitempty,optional"`                                 // Set the NOCOW flag on files. Disables data checksumming and causes data errors to be unrecoverable from while allowing direct I/O. Only use this if data does not need to be any more safe than on a single ext4 formatted disk with no underlying raid system.
	SaferemoveThroughput string `json:"saferemove_throughput,omitempty" url:"saferemove_throughput,omitempty,optional"` // Wipe throughput (cstream -t parameter value).
	Type                 string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`                         // Storage type.
	Smbversion           string `json:"smbversion,omitempty" url:"smbversion,omitempty,optional"`                       // SMB protocol version. 'default' if not set, negotiates the highest SMB2+ version supported by both the client and server.
	Keyring              string `json:"keyring,omitempty" url:"keyring,omitempty,optional"`                             // Client keyring contents (for external clusters).
	MaxProtectedBackups  int    `json:"max-protected-backups,omitempty" url:"max-protected-backups,omitempty,optional"` // Maximal number of protected backups per guest. Use '-1' for unlimited.
	Mkdir                bool   `json:"mkdir,omitempty" url:"mkdir,omitempty,optional"`                                 // Create the directory if it doesn't exist.
	TaggedOnly           bool   `json:"tagged_only,omitempty" url:"tagged_only,omitempty,optional"`                     // Only use logical volumes tagged with 'pve-vm-ID'.
	Sparse               bool   `json:"sparse,omitempty" url:"sparse,omitempty,optional"`                               // use sparse volumes
	Domain               string `json:"domain,omitempty" url:"domain,omitempty,optional"`                               // CIFS domain.
	IsMountpoint         string `json:"is_mountpoint,omitempty" url:"is_mountpoint,omitempty,optional"`                 // Assume the given path is an externally managed mountpoint and consider the storage offline if it is not mounted. Using a boolean (yes/no) value serves as a shortcut to using the target path in this field.
	LioTpg               string `json:"lio_tpg,omitempty" url:"lio_tpg,omitempty,optional"`                             // target portal group for Linux LIO targets
	Namespace            string `json:"namespace,omitempty" url:"namespace,omitempty,optional"`                         // RBD Namespace.
	Share                string `json:"share,omitempty" url:"share,omitempty,optional"`                                 // CIFS share.
	DataPool             string `json:"data-pool,omitempty" url:"data-pool,omitempty,optional"`                         // Data Pool (for erasure coding only)
	Password             string `json:"password,omitempty" url:"password,omitempty,optional"`                           // Password for accessing the share/datastore.
	Subdir               string `json:"subdir,omitempty" url:"subdir,omitempty,optional"`                               // Subdir to mount.
	Pool                 string `json:"pool,omitempty" url:"pool,omitempty,optional"`                                   // Pool.
	Vgname               string `json:"vgname,omitempty" url:"vgname,omitempty,optional"`                               // Volume group name.
	Bwlimit              string `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`                             // Set bandwidth/io limits various operations.
	Maxfiles             int    `json:"maxfiles,omitempty" url:"maxfiles,omitempty,optional"`                           // Deprecated: use 'prune-backups' instead. Maximal number of backup files per VM. Use '0' for unlimited.
	Saferemove           bool   `json:"saferemove,omitempty" url:"saferemove,omitempty,optional"`                       // Zero-out data when removing LVs.
	Shared               bool   `json:"shared,omitempty" url:"shared,omitempty,optional"`                               // Mark storage as shared.
	Storage              string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`                   // The storage identifier.
	Base                 string `json:"base,omitempty" url:"base,omitempty,optional"`                                   // Base volume. This volume is automatically activated.
	Export               string `json:"export,omitempty" url:"export,omitempty,optional"`                               // NFS export path.
	Mountpoint           string `json:"mountpoint,omitempty" url:"mountpoint,omitempty,optional"`                       // mount point
	Port                 int    `json:"port,omitempty" url:"port,omitempty,optional"`                                   // For non default port.
	ComstarHg            string `json:"comstar_hg,omitempty" url:"comstar_hg,omitempty,optional"`                       // host group for comstar views
	Krbd                 bool   `json:"krbd,omitempty" url:"krbd,omitempty,optional"`                                   // Always access rbd through krbd kernel module.
	Nowritecache         bool   `json:"nowritecache,omitempty" url:"nowritecache,omitempty,optional"`                   // disable write caching on the target
	Preallocation        string `json:"preallocation,omitempty" url:"preallocation,omitempty,optional"`                 // Preallocation mode for raw and qcow2 images. Using 'metadata' on raw images results in preallocation=off.
	Server               string `json:"server,omitempty" url:"server,omitempty,optional"`                               // Server IP or DNS name.
	Target               string `json:"target,omitempty" url:"target,omitempty,optional"`                               // iSCSI target.
	Thinpool             string `json:"thinpool,omitempty" url:"thinpool,omitempty,optional"`                           // LVM thin pool LV name.
	Iscsiprovider        string `json:"iscsiprovider,omitempty" url:"iscsiprovider,omitempty,optional"`                 // iscsi provider
	Username             string `json:"username,omitempty" url:"username,omitempty,optional"`                           // RBD Id.
	Volume               string `json:"volume,omitempty" url:"volume,omitempty,optional"`                               // Glusterfs Volume.
	Nodes                string `json:"nodes,omitempty" url:"nodes,omitempty,optional"`                                 // List of cluster node names.
	Content              string `json:"content,omitempty" url:"content,omitempty,optional"`                             // Allowed content types.  NOTE: the value 'rootdir' is used for Containers, and value 'images' for VMs.
	Datastore            string `json:"datastore,omitempty" url:"datastore,omitempty,optional"`                         // Proxmox Backup Server datastore name.
	Disable              bool   `json:"disable,omitempty" url:"disable,omitempty,optional"`                             // Flag to disable the storage.
	Fuse                 bool   `json:"fuse,omitempty" url:"fuse,omitempty,optional"`                                   // Mount CephFS through FUSE.
	Transport            string `json:"transport,omitempty" url:"transport,omitempty,optional"`                         // Gluster transport: tcp or rdma
}

// StorageCreateResponse storage
// Create a new storage.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/storage
//
type StorageCreateResponse struct {
	Config struct { //
	} `json:"config,omitempty" url:"config,omitempty,optional"` // Partial, possible server generated, configuration properties.
	Storage string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"` // The ID of the created storage.
	Type    string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`       // The type of the created storage.
}

// StorageDeleteRequest {storage} - Delete storage configuration.
// Delete storage configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/storage/{storage}
//
type StorageDeleteRequest struct {
	Storage string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"` // The storage identifier.
}

// StorageReadRequest {storage} - Read storage configuration.
// Read storage configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/storage/{storage}
//
type StorageReadRequest struct {
	Storage string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"` // The storage identifier.
}

// StorageUpdateRequest {storage} - Update storage configuration.
// Update storage configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/storage/{storage}
//
type StorageUpdateRequest struct {
	Bwlimit              string `json:"bwlimit,omitempty" url:"bwlimit,omitempty,optional"`                             // Set bandwidth/io limits various operations.
	Disable              bool   `json:"disable,omitempty" url:"disable,omitempty,optional"`                             // Flag to disable the storage.
	Domain               string `json:"domain,omitempty" url:"domain,omitempty,optional"`                               // CIFS domain.
	IsMountpoint         string `json:"is_mountpoint,omitempty" url:"is_mountpoint,omitempty,optional"`                 // Assume the given path is an externally managed mountpoint and consider the storage offline if it is not mounted. Using a boolean (yes/no) value serves as a shortcut to using the target path in this field.
	Krbd                 bool   `json:"krbd,omitempty" url:"krbd,omitempty,optional"`                                   // Always access rbd through krbd kernel module.
	MasterPubkey         string `json:"master-pubkey,omitempty" url:"master-pubkey,omitempty,optional"`                 // Base64-encoded, PEM-formatted public RSA key. Used to encrypt a copy of the encryption-key which will be added to each encrypted backup.
	Mkdir                bool   `json:"mkdir,omitempty" url:"mkdir,omitempty,optional"`                                 // Create the directory if it doesn't exist.
	Password             string `json:"password,omitempty" url:"password,omitempty,optional"`                           // Password for accessing the share/datastore.
	Pool                 string `json:"pool,omitempty" url:"pool,omitempty,optional"`                                   // Pool.
	Saferemove           bool   `json:"saferemove,omitempty" url:"saferemove,omitempty,optional"`                       // Zero-out data when removing LVs.
	Storage              string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"`                   // The storage identifier.
	ComstarHg            string `json:"comstar_hg,omitempty" url:"comstar_hg,omitempty,optional"`                       // host group for comstar views
	ComstarTg            string `json:"comstar_tg,omitempty" url:"comstar_tg,omitempty,optional"`                       // target group for comstar views
	Delete               string `json:"delete,omitempty" url:"delete,omitempty,optional"`                               // A list of settings you want to delete.
	Mountpoint           string `json:"mountpoint,omitempty" url:"mountpoint,omitempty,optional"`                       // mount point
	Preallocation        string `json:"preallocation,omitempty" url:"preallocation,omitempty,optional"`                 // Preallocation mode for raw and qcow2 images. Using 'metadata' on raw images results in preallocation=off.
	Sparse               bool   `json:"sparse,omitempty" url:"sparse,omitempty,optional"`                               // use sparse volumes
	TaggedOnly           bool   `json:"tagged_only,omitempty" url:"tagged_only,omitempty,optional"`                     // Only use logical volumes tagged with 'pve-vm-ID'.
	Transport            string `json:"transport,omitempty" url:"transport,omitempty,optional"`                         // Gluster transport: tcp or rdma
	Content              string `json:"content,omitempty" url:"content,omitempty,optional"`                             // Allowed content types.  NOTE: the value 'rootdir' is used for Containers, and value 'images' for VMs.
	LioTpg               string `json:"lio_tpg,omitempty" url:"lio_tpg,omitempty,optional"`                             // target portal group for Linux LIO targets
	MaxProtectedBackups  int    `json:"max-protected-backups,omitempty" url:"max-protected-backups,omitempty,optional"` // Maximal number of protected backups per guest. Use '-1' for unlimited.
	PruneBackups         string `json:"prune-backups,omitempty" url:"prune-backups,omitempty,optional"`                 // The retention options with shorter intervals are processed first with --keep-last being the very first one. Each option covers a specific period of time. We say that backups within this period are covered by this option. The next option does not take care of already covered backups and only considers older backups.
	EncryptionKey        string `json:"encryption-key,omitempty" url:"encryption-key,omitempty,optional"`               // Encryption key. Use 'autogen' to generate one automatically without passphrase.
	Fuse                 bool   `json:"fuse,omitempty" url:"fuse,omitempty,optional"`                                   // Mount CephFS through FUSE.
	SaferemoveThroughput string `json:"saferemove_throughput,omitempty" url:"saferemove_throughput,omitempty,optional"` // Wipe throughput (cstream -t parameter value).
	FsName               string `json:"fs-name,omitempty" url:"fs-name,omitempty,optional"`                             // The Ceph filesystem name.
	Keyring              string `json:"keyring,omitempty" url:"keyring,omitempty,optional"`                             // Client keyring contents (for external clusters).
	Monhost              string `json:"monhost,omitempty" url:"monhost,omitempty,optional"`                             // IP addresses of monitors (for external clusters).
	Options              string `json:"options,omitempty" url:"options,omitempty,optional"`                             // NFS mount options (see 'man nfs')
	Server2              string `json:"server2,omitempty" url:"server2,omitempty,optional"`                             // Backup volfile server IP or DNS name.
	Shared               bool   `json:"shared,omitempty" url:"shared,omitempty,optional"`                               // Mark storage as shared.
	Smbversion           string `json:"smbversion,omitempty" url:"smbversion,omitempty,optional"`                       // SMB protocol version. 'default' if not set, negotiates the highest SMB2+ version supported by both the client and server.
	Subdir               string `json:"subdir,omitempty" url:"subdir,omitempty,optional"`                               // Subdir to mount.
	Username             string `json:"username,omitempty" url:"username,omitempty,optional"`                           // RBD Id.
	Blocksize            string `json:"blocksize,omitempty" url:"blocksize,omitempty,optional"`                         // block size
	Maxfiles             int    `json:"maxfiles,omitempty" url:"maxfiles,omitempty,optional"`                           // Deprecated: use 'prune-backups' instead. Maximal number of backup files per VM. Use '0' for unlimited.
	Nocow                bool   `json:"nocow,omitempty" url:"nocow,omitempty,optional"`                                 // Set the NOCOW flag on files. Disables data checksumming and causes data errors to be unrecoverable from while allowing direct I/O. Only use this if data does not need to be any more safe than on a single ext4 formatted disk with no underlying raid system.
	Fingerprint          string `json:"fingerprint,omitempty" url:"fingerprint,omitempty,optional"`                     // Certificate SHA 256 fingerprint.
	Server               string `json:"server,omitempty" url:"server,omitempty,optional"`                               // Server IP or DNS name.
	DataPool             string `json:"data-pool,omitempty" url:"data-pool,omitempty,optional"`                         // Data Pool (for erasure coding only)
	Digest               string `json:"digest,omitempty" url:"digest,omitempty,optional"`                               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Format               string `json:"format,omitempty" url:"format,omitempty,optional"`                               // Default image format.
	Namespace            string `json:"namespace,omitempty" url:"namespace,omitempty,optional"`                         // RBD Namespace.
	Nodes                string `json:"nodes,omitempty" url:"nodes,omitempty,optional"`                                 // List of cluster node names.
	Nowritecache         bool   `json:"nowritecache,omitempty" url:"nowritecache,omitempty,optional"`                   // disable write caching on the target
	Port                 int    `json:"port,omitempty" url:"port,omitempty,optional"`                                   // For non default port.
}

// StorageUpdateResponse {storage}
// Update storage configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/storage/{storage}
//
type StorageUpdateResponse struct {
	Config struct { //
	} `json:"config,omitempty" url:"config,omitempty,optional"` // Partial, possible server generated, configuration properties.
	Storage string `json:"storage,omitempty" url:"storage,omitempty" validate:"nonzero"` // The ID of the created storage.
	Type    string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`       // The type of the created storage.
}

// AccessUsersCreateUserRequest users - Create new user.
// Create new user.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users
//
type AccessUsersCreateUserRequest struct {
	Email     string `json:"email,omitempty" url:"email,omitempty,optional"`             //
	Enable    bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`           // Enable the account (default). You can set this to '0' to disable the account
	Expire    int    `json:"expire,omitempty" url:"expire,omitempty,optional"`           // Account expiration date (seconds since epoch). '0' means no expiration date.
	Groups    string `json:"groups,omitempty" url:"groups,omitempty,optional"`           //
	Lastname  string `json:"lastname,omitempty" url:"lastname,omitempty,optional"`       //
	Comment   string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Firstname string `json:"firstname,omitempty" url:"firstname,omitempty,optional"`     //
	Keys      string `json:"keys,omitempty" url:"keys,omitempty,optional"`               // Keys for two factor auth (yubico).
	Password  string `json:"password,omitempty" url:"password,omitempty,optional"`       // Initial password.
	Userid    string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
}

// AccessUsersReadUserRequest {userid} - Get user configuration.
// Get user configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}
//
type AccessUsersReadUserRequest struct {
	Userid string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
}

// AccessUsersReadUserResponse {userid}
// Get user configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}
//
type AccessUsersReadUserResponse struct {
	Groups []struct {
	} `json:"groups,omitempty" url:"groups,omitempty,optional"` //
	Keys      string   `json:"keys,omitempty" url:"keys,omitempty,optional"`           // Keys for two factor auth (yubico).
	Lastname  string   `json:"lastname,omitempty" url:"lastname,omitempty,optional"`   //
	Email     string   `json:"email,omitempty" url:"email,omitempty,optional"`         //
	Enable    bool     `json:"enable,omitempty" url:"enable,omitempty,optional"`       // Enable the account (default). You can set this to '0' to disable the account
	Expire    int      `json:"expire,omitempty" url:"expire,omitempty,optional"`       // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname string   `json:"firstname,omitempty" url:"firstname,omitempty,optional"` //
	Tokens    struct { //
	} `json:"tokens,omitempty" url:"tokens,omitempty,optional"` //
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"` //
}

// AccessUsersUpdateUserRequest {userid} - Update user configuration.
// Update user configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}
//
type AccessUsersUpdateUserRequest struct {
	Append    bool   `json:"append,omitempty" url:"append,omitempty,optional"`           //
	Enable    bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`           // Enable the account (default). You can set this to '0' to disable the account
	Expire    int    `json:"expire,omitempty" url:"expire,omitempty,optional"`           // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname string `json:"firstname,omitempty" url:"firstname,omitempty,optional"`     //
	Keys      string `json:"keys,omitempty" url:"keys,omitempty,optional"`               // Keys for two factor auth (yubico).
	Userid    string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
	Comment   string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Email     string `json:"email,omitempty" url:"email,omitempty,optional"`             //
	Groups    string `json:"groups,omitempty" url:"groups,omitempty,optional"`           //
	Lastname  string `json:"lastname,omitempty" url:"lastname,omitempty,optional"`       //
}

// AccessUsersDeleteUserRequest {userid} - Delete user.
// Delete user.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}
//
type AccessUsersDeleteUserRequest struct {
	Userid string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
}

// AccessUsersTfaReadUserTypeRequest tfa - Get user TFA types (Personal and Realm).
// Get user TFA types (Personal and Realm).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/tfa
//
type AccessUsersTfaReadUserTypeRequest struct {
	Multiple bool   `json:"multiple,omitempty" url:"multiple,omitempty,optional"`       // Request all entries as an array.
	Userid   string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
}

// AccessUsersTfaReadUserTypeResponse tfa
// Get user TFA types (Personal and Realm).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/tfa
//
type AccessUsersTfaReadUserTypeResponse struct {
	Realm string `json:"realm,omitempty" url:"realm,omitempty,optional"` // The type of TFA the users realm has set, if any.
	Types []struct {
	} `json:"types,omitempty" url:"types,omitempty,optional"` // Array of the user configured TFA types, if any. Only available if 'multiple' was not passed.
	User string `json:"user,omitempty" url:"user,omitempty,optional"` // The type of TFA the user has set, if any. Only set if 'multiple' was not passed.
}

// AccessUsersTokenIndexRequest token - Get user API tokens.
// Get user API tokens.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/token
//
type AccessUsersTokenIndexRequest struct {
	Userid string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
}

// AccessUsersTokenIndexResponse token
// Get user API tokens.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/token
//
type AccessUsersTokenIndexResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`           //
	Expire  int    `json:"expire,omitempty" url:"expire,omitempty,optional"`             // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep bool   `json:"privsep,omitempty" url:"privsep,omitempty,optional"`           // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
	Tokenid string `json:"tokenid,omitempty" url:"tokenid,omitempty" validate:"nonzero"` // User-specific token identifier.
}

// AccessUsersTokenRemoveRequest {tokenid} - Remove API token for a specific user.
// Remove API token for a specific user.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/token/{tokenid}
//
type AccessUsersTokenRemoveRequest struct {
	Tokenid string `json:"tokenid,omitempty" url:"tokenid,omitempty" validate:"nonzero"` // User-specific token identifier.
	Userid  string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"`   // User ID
}

// AccessUsersTokenReadRequest {tokenid} - Get specific API token information.
// Get specific API token information.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/token/{tokenid}
//
type AccessUsersTokenReadRequest struct {
	Tokenid string `json:"tokenid,omitempty" url:"tokenid,omitempty" validate:"nonzero"` // User-specific token identifier.
	Userid  string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"`   // User ID
}

// AccessUsersTokenReadResponse {tokenid}
// Get specific API token information.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/token/{tokenid}
//
type AccessUsersTokenReadResponse struct {
	Expire  int    `json:"expire,omitempty" url:"expire,omitempty,optional"`   // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep bool   `json:"privsep,omitempty" url:"privsep,omitempty,optional"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"` //
}

// AccessUsersTokenGenerateRequest {tokenid} - Generate a new API token for a specific user. NOTE: returns API token value, which needs to be stored as it cannot be retrieved afterwards!
// Generate a new API token for a specific user. NOTE: returns API token value, which needs to be stored as it cannot be retrieved afterwards!
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/token/{tokenid}
//
type AccessUsersTokenGenerateRequest struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`           //
	Expire  int    `json:"expire,omitempty" url:"expire,omitempty,optional"`             // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep bool   `json:"privsep,omitempty" url:"privsep,omitempty,optional"`           // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
	Tokenid string `json:"tokenid,omitempty" url:"tokenid,omitempty" validate:"nonzero"` // User-specific token identifier.
	Userid  string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"`   // User ID
}

// AccessUsersTokenGenerateResponse {tokenid}
// Generate a new API token for a specific user. NOTE: returns API token value, which needs to be stored as it cannot be retrieved afterwards!
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/token/{tokenid}
//
type AccessUsersTokenGenerateResponse struct {
	FullTokenid string   `json:"full-tokenid,omitempty" url:"full-tokenid,omitempty" validate:"nonzero"` // The full token id.
	Info        struct { //
	} `json:"info,omitempty" url:"info,omitempty" validate:"nonzero"` //
	Value string `json:"value,omitempty" url:"value,omitempty" validate:"nonzero"` // API token value used for authentication.
}

// AccessUsersTokenUpdateInfoRequest {tokenid} - Update API token for a specific user.
// Update API token for a specific user.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/token/{tokenid}
//
type AccessUsersTokenUpdateInfoRequest struct {
	Privsep bool   `json:"privsep,omitempty" url:"privsep,omitempty,optional"`           // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
	Tokenid string `json:"tokenid,omitempty" url:"tokenid,omitempty" validate:"nonzero"` // User-specific token identifier.
	Userid  string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"`   // User ID
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`           //
	Expire  int    `json:"expire,omitempty" url:"expire,omitempty,optional"`             // API token expiration date (seconds since epoch). '0' means no expiration date.
}

// AccessUsersTokenUpdateInfoResponse {tokenid}
// Update API token for a specific user.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/users/{userid}/token/{tokenid}
//
type AccessUsersTokenUpdateInfoResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"` //
	Expire  int    `json:"expire,omitempty" url:"expire,omitempty,optional"`   // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep bool   `json:"privsep,omitempty" url:"privsep,omitempty,optional"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
}

// AccessGroupsCreateGroupRequest groups - Create new group.
// Create new group.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/groups
//
type AccessGroupsCreateGroupRequest struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`           //
	Groupid string `json:"groupid,omitempty" url:"groupid,omitempty" validate:"nonzero"` //
}

// AccessGroupsDeleteGroupRequest {groupid} - Delete group.
// Delete group.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/groups/{groupid}
//
type AccessGroupsDeleteGroupRequest struct {
	Groupid string `json:"groupid,omitempty" url:"groupid,omitempty" validate:"nonzero"` //
}

// AccessGroupsReadGroupRequest {groupid} - Get group configuration.
// Get group configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/groups/{groupid}
//
type AccessGroupsReadGroupRequest struct {
	Groupid string `json:"groupid,omitempty" url:"groupid,omitempty" validate:"nonzero"` //
}

// AccessGroupsReadGroupResponse {groupid}
// Get group configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/groups/{groupid}
//
type AccessGroupsReadGroupResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"` //
	Members []struct {
	} `json:"members,omitempty" url:"members,omitempty" validate:"nonzero"` //
}

// AccessGroupsUpdateGroupRequest {groupid} - Update group data.
// Update group data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/groups/{groupid}
//
type AccessGroupsUpdateGroupRequest struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`           //
	Groupid string `json:"groupid,omitempty" url:"groupid,omitempty" validate:"nonzero"` //
}

// AccessRolesCreateRoleRequest roles - Create new role.
// Create new role.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/roles
//
type AccessRolesCreateRoleRequest struct {
	Privs  string `json:"privs,omitempty" url:"privs,omitempty,optional"`             //
	Roleid string `json:"roleid,omitempty" url:"roleid,omitempty" validate:"nonzero"` //
}

// AccessRolesUpdateRoleRequest {roleid} - Update an existing role.
// Update an existing role.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/roles/{roleid}
//
type AccessRolesUpdateRoleRequest struct {
	Append bool   `json:"append,omitempty" url:"append,omitempty,optional"`           //
	Privs  string `json:"privs,omitempty" url:"privs,omitempty,optional"`             //
	Roleid string `json:"roleid,omitempty" url:"roleid,omitempty" validate:"nonzero"` //
}

// AccessRolesDeleteRoleRequest {roleid} - Delete role.
// Delete role.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/roles/{roleid}
//
type AccessRolesDeleteRoleRequest struct {
	Roleid string `json:"roleid,omitempty" url:"roleid,omitempty" validate:"nonzero"` //
}

// AccessRolesReadRoleRequest {roleid} - Get role configuration.
// Get role configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/roles/{roleid}
//
type AccessRolesReadRoleRequest struct {
	Roleid string `json:"roleid,omitempty" url:"roleid,omitempty" validate:"nonzero"` //
}

// AccessRolesReadRoleResponse {roleid}
// Get role configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/roles/{roleid}
//
type AccessRolesReadRoleResponse struct {
	VMSnapshot                bool `json:"VM.Snapshot,omitempty" url:"VM.Snapshot,omitempty,optional"`                               //
	VMSnapshotRollback        bool `json:"VM.Snapshot.Rollback,omitempty" url:"VM.Snapshot.Rollback,omitempty,optional"`             //
	DatastoreAllocateSpace    bool `json:"Datastore.AllocateSpace,omitempty" url:"Datastore.AllocateSpace,omitempty,optional"`       //
	SDNAudit                  bool `json:"SDN.Audit,omitempty" url:"SDN.Audit,omitempty,optional"`                                   //
	SysPowerMgmt              bool `json:"Sys.PowerMgmt,omitempty" url:"Sys.PowerMgmt,omitempty,optional"`                           //
	UserModify                bool `json:"User.Modify,omitempty" url:"User.Modify,omitempty,optional"`                               //
	VMMigrate                 bool `json:"VM.Migrate,omitempty" url:"VM.Migrate,omitempty,optional"`                                 //
	VMConfigNetwork           bool `json:"VM.Config.Network,omitempty" url:"VM.Config.Network,omitempty,optional"`                   //
	VMConfigOptions           bool `json:"VM.Config.Options,omitempty" url:"VM.Config.Options,omitempty,optional"`                   //
	DatastoreAllocateTemplate bool `json:"Datastore.AllocateTemplate,omitempty" url:"Datastore.AllocateTemplate,omitempty,optional"` //
	SysSyslog                 bool `json:"Sys.Syslog,omitempty" url:"Sys.Syslog,omitempty,optional"`                                 //
	VMAudit                   bool `json:"VM.Audit,omitempty" url:"VM.Audit,omitempty,optional"`                                     //
	VMConfigCloudinit         bool `json:"VM.Config.Cloudinit,omitempty" url:"VM.Config.Cloudinit,omitempty,optional"`               //
	VMConfigHWType            bool `json:"VM.Config.HWType,omitempty" url:"VM.Config.HWType,omitempty,optional"`                     //
	PoolAudit                 bool `json:"Pool.Audit,omitempty" url:"Pool.Audit,omitempty,optional"`                                 //
	SysAudit                  bool `json:"Sys.Audit,omitempty" url:"Sys.Audit,omitempty,optional"`                                   //
	SysConsole                bool `json:"Sys.Console,omitempty" url:"Sys.Console,omitempty,optional"`                               //
	VMConsole                 bool `json:"VM.Console,omitempty" url:"VM.Console,omitempty,optional"`                                 //
	PermissionsModify         bool `json:"Permissions.Modify,omitempty" url:"Permissions.Modify,omitempty,optional"`                 //
	SDNAllocate               bool `json:"SDN.Allocate,omitempty" url:"SDN.Allocate,omitempty,optional"`                             //
	SysModify                 bool `json:"Sys.Modify,omitempty" url:"Sys.Modify,omitempty,optional"`                                 //
	VMAllocate                bool `json:"VM.Allocate,omitempty" url:"VM.Allocate,omitempty,optional"`                               //
	VMPowerMgmt               bool `json:"VM.PowerMgmt,omitempty" url:"VM.PowerMgmt,omitempty,optional"`                             //
	VMConfigDisk              bool `json:"VM.Config.Disk,omitempty" url:"VM.Config.Disk,omitempty,optional"`                         //
	DatastoreAudit            bool `json:"Datastore.Audit,omitempty" url:"Datastore.Audit,omitempty,optional"`                       //
	RealmAllocate             bool `json:"Realm.Allocate,omitempty" url:"Realm.Allocate,omitempty,optional"`                         //
	RealmAllocateUser         bool `json:"Realm.AllocateUser,omitempty" url:"Realm.AllocateUser,omitempty,optional"`                 //
	VMBackup                  bool `json:"VM.Backup,omitempty" url:"VM.Backup,omitempty,optional"`                                   //
	VMConfigCDROM             bool `json:"VM.Config.CDROM,omitempty" url:"VM.Config.CDROM,omitempty,optional"`                       //
	VMMonitor                 bool `json:"VM.Monitor,omitempty" url:"VM.Monitor,omitempty,optional"`                                 //
	DatastoreAllocate         bool `json:"Datastore.Allocate,omitempty" url:"Datastore.Allocate,omitempty,optional"`                 //
	GroupAllocate             bool `json:"Group.Allocate,omitempty" url:"Group.Allocate,omitempty,optional"`                         //
	PoolAllocate              bool `json:"Pool.Allocate,omitempty" url:"Pool.Allocate,omitempty,optional"`                           //
	VMClone                   bool `json:"VM.Clone,omitempty" url:"VM.Clone,omitempty,optional"`                                     //
	VMConfigCPU               bool `json:"VM.Config.CPU,omitempty" url:"VM.Config.CPU,omitempty,optional"`                           //
	VMConfigMemory            bool `json:"VM.Config.Memory,omitempty" url:"VM.Config.Memory,omitempty,optional"`                     //
}

// AccessAclUpdateRequest acl - Update Access Control List (add or remove permissions).
// Update Access Control List (add or remove permissions).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/acl
//
type AccessAclUpdateRequest struct {
	Users     string `json:"users,omitempty" url:"users,omitempty,optional"`           // List of users.
	Delete    bool   `json:"delete,omitempty" url:"delete,omitempty,optional"`         // Remove permissions (instead of adding it).
	Groups    string `json:"groups,omitempty" url:"groups,omitempty,optional"`         // List of groups.
	Path      string `json:"path,omitempty" url:"path,omitempty" validate:"nonzero"`   // Access control path
	Propagate bool   `json:"propagate,omitempty" url:"propagate,omitempty,optional"`   // Allow to propagate (inherit) permissions.
	Roles     string `json:"roles,omitempty" url:"roles,omitempty" validate:"nonzero"` // List of roles.
	Tokens    string `json:"tokens,omitempty" url:"tokens,omitempty,optional"`         // List of API tokens.
}

// AccessAclReadResponse acl
// Get Access Control List (ACLs).
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/acl
//
type AccessAclReadResponse struct {
	Roleid    string `json:"roleid,omitempty" url:"roleid,omitempty" validate:"nonzero"` //
	Type      string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     //
	Ugid      string `json:"ugid,omitempty" url:"ugid,omitempty" validate:"nonzero"`     //
	Path      string `json:"path,omitempty" url:"path,omitempty" validate:"nonzero"`     // Access control path
	Propagate bool   `json:"propagate,omitempty" url:"propagate,omitempty,optional"`     // Allow to propagate (inherit) permissions.
}

// AccessDomainsCreateRequest domains - Add an authentication server.
// Add an authentication server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/domains
//
type AccessDomainsCreateRequest struct {
	BaseDn              string `json:"base_dn,omitempty" url:"base_dn,omitempty,optional"`                             // LDAP base domain name
	Certkey             string `json:"certkey,omitempty" url:"certkey,omitempty,optional"`                             // Path to the client certificate key
	Realm               string `json:"realm,omitempty" url:"realm,omitempty" validate:"nonzero"`                       // Authentication domain ID
	SyncDefaultsOptions string `json:"sync-defaults-options,omitempty" url:"sync-defaults-options,omitempty,optional"` // The default options for behavior of synchronizations.
	Default             bool   `json:"default,omitempty" url:"default,omitempty,optional"`                             // Use this as default realm
	GroupDn             string `json:"group_dn,omitempty" url:"group_dn,omitempty,optional"`                           // LDAP base domain name for group sync. If not set, the base_dn will be used.
	Sslversion          string `json:"sslversion,omitempty" url:"sslversion,omitempty,optional"`                       // LDAPS TLS/SSL version. It's not recommended to use version older than 1.2!
	UserClasses         string `json:"user_classes,omitempty" url:"user_classes,omitempty,optional"`                   // The objectclasses for users.
	ClientId            string `json:"client-id,omitempty" url:"client-id,omitempty,optional"`                         // OpenID Client ID
	Domain              string `json:"domain,omitempty" url:"domain,omitempty,optional"`                               // AD domain name
	GroupFilter         string `json:"group_filter,omitempty" url:"group_filter,omitempty,optional"`                   // LDAP filter for group sync.
	GroupNameAttr       string `json:"group_name_attr,omitempty" url:"group_name_attr,omitempty,optional"`             // LDAP attribute representing a groups name. If not set or found, the first value of the DN will be used as name.
	UserAttr            string `json:"user_attr,omitempty" url:"user_attr,omitempty,optional"`                         // LDAP user attribute name
	ClientKey           string `json:"client-key,omitempty" url:"client-key,omitempty,optional"`                       // OpenID Client Key
	Comment             string `json:"comment,omitempty" url:"comment,omitempty,optional"`                             // Description.
	Filter              string `json:"filter,omitempty" url:"filter,omitempty,optional"`                               // LDAP filter for user sync.
	UsernameClaim       string `json:"username-claim,omitempty" url:"username-claim,omitempty,optional"`               // OpenID claim used to generate the unique username.
	Verify              bool   `json:"verify,omitempty" url:"verify,omitempty,optional"`                               // Verify the server's SSL certificate
	BindDn              string `json:"bind_dn,omitempty" url:"bind_dn,omitempty,optional"`                             // LDAP bind domain name
	Capath              string `json:"capath,omitempty" url:"capath,omitempty,optional"`                               // Path to the CA certificate store
	IssuerUrl           string `json:"issuer-url,omitempty" url:"issuer-url,omitempty,optional"`                       // OpenID Issuer Url
	Server1             string `json:"server1,omitempty" url:"server1,omitempty,optional"`                             // Server IP address (or DNS name)
	Server2             string `json:"server2,omitempty" url:"server2,omitempty,optional"`                             // Fallback Server IP address (or DNS name)
	Autocreate          bool   `json:"autocreate,omitempty" url:"autocreate,omitempty,optional"`                       // Automatically create users if they do not exist.
	CaseSensitive       bool   `json:"case-sensitive,omitempty" url:"case-sensitive,omitempty,optional"`               // username is case-sensitive
	Prompt              string `json:"prompt,omitempty" url:"prompt,omitempty,optional"`                               // Specifies whether the Authorization Server prompts the End-User for reauthentication and consent.
	Scopes              string `json:"scopes,omitempty" url:"scopes,omitempty,optional"`                               // Specifies the scopes (user details) that should be authorized and returned, for example 'email' or 'profile'.
	Secure              bool   `json:"secure,omitempty" url:"secure,omitempty,optional"`                               // Use secure LDAPS protocol. DEPRECATED: use 'mode' instead.
	AcrValues           string `json:"acr-values,omitempty" url:"acr-values,omitempty,optional"`                       // Specifies the Authentication Context Class Reference values that theAuthorization Server is being requested to use for the Auth Request.
	Mode                string `json:"mode,omitempty" url:"mode,omitempty,optional"`                                   // LDAP protocol mode.
	Port                int    `json:"port,omitempty" url:"port,omitempty,optional"`                                   // Server port.
	Tfa                 string `json:"tfa,omitempty" url:"tfa,omitempty,optional"`                                     // Use Two-factor authentication.
	Cert                string `json:"cert,omitempty" url:"cert,omitempty,optional"`                                   // Path to the client certificate
	GroupClasses        string `json:"group_classes,omitempty" url:"group_classes,omitempty,optional"`                 // The objectclasses for groups.
	Password            string `json:"password,omitempty" url:"password,omitempty,optional"`                           // LDAP bind password. Will be stored in '/etc/pve/priv/realm/<REALM>.pw'.
	SyncAttributes      string `json:"sync_attributes,omitempty" url:"sync_attributes,omitempty,optional"`             // Comma separated list of key=value pairs for specifying which LDAP attributes map to which PVE user field. For example, to map the LDAP attribute 'mail' to PVEs 'email', write  'email=mail'. By default, each PVE user field is represented  by an LDAP attribute of the same name.
	Type                string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`                         // Realm type.
}

// AccessDomainsUpdateRequest {realm} - Update authentication server settings.
// Update authentication server settings.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/domains/{realm}
//
type AccessDomainsUpdateRequest struct {
	Tfa                 string `json:"tfa,omitempty" url:"tfa,omitempty,optional"`                                     // Use Two-factor authentication.
	Autocreate          bool   `json:"autocreate,omitempty" url:"autocreate,omitempty,optional"`                       // Automatically create users if they do not exist.
	Certkey             string `json:"certkey,omitempty" url:"certkey,omitempty,optional"`                             // Path to the client certificate key
	ClientId            string `json:"client-id,omitempty" url:"client-id,omitempty,optional"`                         // OpenID Client ID
	Delete              string `json:"delete,omitempty" url:"delete,omitempty,optional"`                               // A list of settings you want to delete.
	GroupFilter         string `json:"group_filter,omitempty" url:"group_filter,omitempty,optional"`                   // LDAP filter for group sync.
	Sslversion          string `json:"sslversion,omitempty" url:"sslversion,omitempty,optional"`                       // LDAPS TLS/SSL version. It's not recommended to use version older than 1.2!
	Cert                string `json:"cert,omitempty" url:"cert,omitempty,optional"`                                   // Path to the client certificate
	SyncDefaultsOptions string `json:"sync-defaults-options,omitempty" url:"sync-defaults-options,omitempty,optional"` // The default options for behavior of synchronizations.
	UserClasses         string `json:"user_classes,omitempty" url:"user_classes,omitempty,optional"`                   // The objectclasses for users.
	Default             bool   `json:"default,omitempty" url:"default,omitempty,optional"`                             // Use this as default realm
	Digest              string `json:"digest,omitempty" url:"digest,omitempty,optional"`                               // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Domain              string `json:"domain,omitempty" url:"domain,omitempty,optional"`                               // AD domain name
	GroupClasses        string `json:"group_classes,omitempty" url:"group_classes,omitempty,optional"`                 // The objectclasses for groups.
	Secure              bool   `json:"secure,omitempty" url:"secure,omitempty,optional"`                               // Use secure LDAPS protocol. DEPRECATED: use 'mode' instead.
	BindDn              string `json:"bind_dn,omitempty" url:"bind_dn,omitempty,optional"`                             // LDAP bind domain name
	GroupNameAttr       string `json:"group_name_attr,omitempty" url:"group_name_attr,omitempty,optional"`             // LDAP attribute representing a groups name. If not set or found, the first value of the DN will be used as name.
	Mode                string `json:"mode,omitempty" url:"mode,omitempty,optional"`                                   // LDAP protocol mode.
	UserAttr            string `json:"user_attr,omitempty" url:"user_attr,omitempty,optional"`                         // LDAP user attribute name
	AcrValues           string `json:"acr-values,omitempty" url:"acr-values,omitempty,optional"`                       // Specifies the Authentication Context Class Reference values that theAuthorization Server is being requested to use for the Auth Request.
	BaseDn              string `json:"base_dn,omitempty" url:"base_dn,omitempty,optional"`                             // LDAP base domain name
	Capath              string `json:"capath,omitempty" url:"capath,omitempty,optional"`                               // Path to the CA certificate store
	ClientKey           string `json:"client-key,omitempty" url:"client-key,omitempty,optional"`                       // OpenID Client Key
	Password            string `json:"password,omitempty" url:"password,omitempty,optional"`                           // LDAP bind password. Will be stored in '/etc/pve/priv/realm/<REALM>.pw'.
	SyncAttributes      string `json:"sync_attributes,omitempty" url:"sync_attributes,omitempty,optional"`             // Comma separated list of key=value pairs for specifying which LDAP attributes map to which PVE user field. For example, to map the LDAP attribute 'mail' to PVEs 'email', write  'email=mail'. By default, each PVE user field is represented  by an LDAP attribute of the same name.
	Verify              bool   `json:"verify,omitempty" url:"verify,omitempty,optional"`                               // Verify the server's SSL certificate
	CaseSensitive       bool   `json:"case-sensitive,omitempty" url:"case-sensitive,omitempty,optional"`               // username is case-sensitive
	Filter              string `json:"filter,omitempty" url:"filter,omitempty,optional"`                               // LDAP filter for user sync.
	GroupDn             string `json:"group_dn,omitempty" url:"group_dn,omitempty,optional"`                           // LDAP base domain name for group sync. If not set, the base_dn will be used.
	Port                int    `json:"port,omitempty" url:"port,omitempty,optional"`                                   // Server port.
	Realm               string `json:"realm,omitempty" url:"realm,omitempty" validate:"nonzero"`                       // Authentication domain ID
	Server2             string `json:"server2,omitempty" url:"server2,omitempty,optional"`                             // Fallback Server IP address (or DNS name)
	Comment             string `json:"comment,omitempty" url:"comment,omitempty,optional"`                             // Description.
	IssuerUrl           string `json:"issuer-url,omitempty" url:"issuer-url,omitempty,optional"`                       // OpenID Issuer Url
	Prompt              string `json:"prompt,omitempty" url:"prompt,omitempty,optional"`                               // Specifies whether the Authorization Server prompts the End-User for reauthentication and consent.
	Scopes              string `json:"scopes,omitempty" url:"scopes,omitempty,optional"`                               // Specifies the scopes (user details) that should be authorized and returned, for example 'email' or 'profile'.
	Server1             string `json:"server1,omitempty" url:"server1,omitempty,optional"`                             // Server IP address (or DNS name)
}

// AccessDomainsDeleteRequest {realm} - Delete an authentication server.
// Delete an authentication server.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/domains/{realm}
//
type AccessDomainsDeleteRequest struct {
	Realm string `json:"realm,omitempty" url:"realm,omitempty" validate:"nonzero"` // Authentication domain ID
}

// AccessDomainsReadRequest {realm} - Get auth server configuration.
// Get auth server configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/domains/{realm}
//
type AccessDomainsReadRequest struct {
	Realm string `json:"realm,omitempty" url:"realm,omitempty" validate:"nonzero"` // Authentication domain ID
}

// AccessDomainsSyncRequest sync - Syncs users and/or groups from the configured LDAP to user.cfg. NOTE: Synced groups will have the name 'name-$realm', so make sure those groups do not exist to prevent overwriting.
// Syncs users and/or groups from the configured LDAP to user.cfg. NOTE: Synced groups will have the name 'name-$realm', so make sure those groups do not exist to prevent overwriting.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/domains/{realm}/sync
//
type AccessDomainsSyncRequest struct {
	DryRun         bool   `json:"dry-run,omitempty" url:"dry-run,omitempty,optional"`                 // If set, does not write anything.
	EnableNew      bool   `json:"enable-new,omitempty" url:"enable-new,omitempty,optional"`           // Enable newly synced users immediately.
	Full           bool   `json:"full,omitempty" url:"full,omitempty,optional"`                       // DEPRECATED: use 'remove-vanished' instead. If set, uses the LDAP Directory as source of truth, deleting users or groups not returned from the sync and removing all locally modified properties of synced users. If not set, only syncs information which is present in the synced data, and does not delete or modify anything else.
	Purge          bool   `json:"purge,omitempty" url:"purge,omitempty,optional"`                     // DEPRECATED: use 'remove-vanished' instead. Remove ACLs for users or groups which were removed from the config during a sync.
	Realm          string `json:"realm,omitempty" url:"realm,omitempty" validate:"nonzero"`           // Authentication domain ID
	RemoveVanished string `json:"remove-vanished,omitempty" url:"remove-vanished,omitempty,optional"` // A semicolon-seperated list of things to remove when they or the user vanishes during a sync. The following values are possible: 'entry' removes the user/group when not returned from the sync. 'properties' removes the set properties on existing user/group that do not appear in the source (even custom ones). 'acl' removes acls when the user/group is not returned from the sync.
	Scope          string `json:"scope,omitempty" url:"scope,omitempty,optional"`                     // Select what to sync.
}

// AccessOpenidAuthUrlAuthUrlRequest auth-url - Get the OpenId Authorization Url for the specified realm.
// Get the OpenId Authorization Url for the specified realm.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/openid/auth-url
//
type AccessOpenidAuthUrlAuthUrlRequest struct {
	Realm       string `json:"realm,omitempty" url:"realm,omitempty" validate:"nonzero"`               // Authentication domain ID
	RedirectUrl string `json:"redirect-url,omitempty" url:"redirect-url,omitempty" validate:"nonzero"` // Redirection Url. The client should set this to the used server url (location.origin).
}

// AccessOpenidLoginRequest login -  Verify OpenID authorization code and create a ticket.
//  Verify OpenID authorization code and create a ticket.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/openid/login
//
type AccessOpenidLoginRequest struct {
	Code        string `json:"code,omitempty" url:"code,omitempty" validate:"nonzero"`                 // OpenId authorization code.
	RedirectUrl string `json:"redirect-url,omitempty" url:"redirect-url,omitempty" validate:"nonzero"` // Redirection Url. The client should set this to the used server url (location.origin).
	State       string `json:"state,omitempty" url:"state,omitempty" validate:"nonzero"`               // OpenId state.
}

// AccessOpenidLoginResponse login
//  Verify OpenID authorization code and create a ticket.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/openid/login
//
type AccessOpenidLoginResponse struct {
	CSRFPreventionToken string   `json:"CSRFPreventionToken,omitempty" url:"CSRFPreventionToken,omitempty" validate:"nonzero"` //
	Cap                 struct { //
	} `json:"cap,omitempty" url:"cap,omitempty" validate:"nonzero"` //
	Clustername string `json:"clustername,omitempty" url:"clustername,omitempty,optional"`     //
	Ticket      string `json:"ticket,omitempty" url:"ticket,omitempty" validate:"nonzero"`     //
	Username    string `json:"username,omitempty" url:"username,omitempty" validate:"nonzero"` //
}

// AccessTfaListResponse tfa
// List TFA configurations of users.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa
//
type AccessTfaListResponse struct {
	Entries []struct {
	} `json:"entries,omitempty" url:"entries,omitempty" validate:"nonzero"` //
	Userid string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User this entry belongs to.
}

// AccessTfaVerifyRequest tfa - Finish a u2f challenge.
// Finish a u2f challenge.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa
//
type AccessTfaVerifyRequest struct {
	Response string `json:"response,omitempty" url:"response,omitempty" validate:"nonzero"` // The response to the current authentication challenge.
}

// AccessTfaVerifyResponse tfa
// Finish a u2f challenge.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa
//
type AccessTfaVerifyResponse struct {
	Ticket string `json:"ticket,omitempty" url:"ticket,omitempty" validate:"nonzero"` //
}

// AccessTfaListUserRequest {userid} - List TFA configurations of users.
// List TFA configurations of users.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa/{userid}
//
type AccessTfaListUserRequest struct {
	Userid string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
}

// AccessTfaListUserResponse {userid}
// List TFA configurations of users.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa/{userid}
//
type AccessTfaListUserResponse struct {
	Type        string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`               // TFA Entry Type.
	Created     int    `json:"created,omitempty" url:"created,omitempty" validate:"nonzero"`         // Creation time of this entry as unix epoch.
	Description string `json:"description,omitempty" url:"description,omitempty" validate:"nonzero"` // User chosen description for this entry.
	Enable      bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`                     // Whether this TFA entry is currently enabled.
	Id          string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`                   // The id used to reference this entry.
}

// AccessTfaAddEntryRequest {userid} - Add a TFA entry for a user.
// Add a TFA entry for a user.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa/{userid}
//
type AccessTfaAddEntryRequest struct {
	Type        string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`     // TFA Entry Type.
	Userid      string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
	Value       string `json:"value,omitempty" url:"value,omitempty,optional"`             // The current value for the provided totp URI, or a Webauthn/U2F challenge response
	Challenge   string `json:"challenge,omitempty" url:"challenge,omitempty,optional"`     // When responding to a u2f challenge: the original challenge string
	Description string `json:"description,omitempty" url:"description,omitempty,optional"` // A description to distinguish multiple entries from one another
	Password    string `json:"password,omitempty" url:"password,omitempty,optional"`       // The current password.
	Totp        string `json:"totp,omitempty" url:"totp,omitempty,optional"`               // A totp URI.
}

// AccessTfaAddEntryResponse {userid}
// Add a TFA entry for a user.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa/{userid}
//
type AccessTfaAddEntryResponse struct {
	Challenge string `json:"challenge,omitempty" url:"challenge,omitempty,optional"` // When adding u2f entries, this contains a challenge the user must respond to in order to finish the registration.
	Id        string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`     // The id of a newly added TFA entry.
	Recovery  []struct {
	} `json:"recovery,omitempty" url:"recovery,omitempty,optional"` // When adding recovery codes, this contains the list of codes to be displayed to the user
}

// AccessTfaDeleteRequest {id} - Delete a TFA entry by ID.
// Delete a TFA entry by ID.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa/{userid}/{id}
//
type AccessTfaDeleteRequest struct {
	Id       string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`         // A TFA entry id.
	Password string `json:"password,omitempty" url:"password,omitempty,optional"`       // The current password.
	Userid   string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
}

// AccessTfaGetEntryRequest {id} - Fetch a requested TFA entry if present.
// Fetch a requested TFA entry if present.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa/{userid}/{id}
//
type AccessTfaGetEntryRequest struct {
	Userid string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
	Id     string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`         // A TFA entry id.
}

// AccessTfaGetEntryResponse {id}
// Fetch a requested TFA entry if present.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa/{userid}/{id}
//
type AccessTfaGetEntryResponse struct {
	Created     int    `json:"created,omitempty" url:"created,omitempty" validate:"nonzero"`         // Creation time of this entry as unix epoch.
	Description string `json:"description,omitempty" url:"description,omitempty" validate:"nonzero"` // User chosen description for this entry.
	Enable      bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`                     // Whether this TFA entry is currently enabled.
	Id          string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`                   // The id used to reference this entry.
	Type        string `json:"type,omitempty" url:"type,omitempty" validate:"nonzero"`               // TFA Entry Type.
}

// AccessTfaUpdateEntryRequest {id} - Add a TFA entry for a user.
// Add a TFA entry for a user.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/tfa/{userid}/{id}
//
type AccessTfaUpdateEntryRequest struct {
	Id          string `json:"id,omitempty" url:"id,omitempty" validate:"nonzero"`         // A TFA entry id.
	Password    string `json:"password,omitempty" url:"password,omitempty,optional"`       // The current password.
	Userid      string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"` // User ID
	Description string `json:"description,omitempty" url:"description,omitempty,optional"` // A description to distinguish multiple entries from one another
	Enable      bool   `json:"enable,omitempty" url:"enable,omitempty,optional"`           // Whether the entry should be enabled for login.
}

// AccessTicketCreateRequest ticket - Create or verify authentication ticket.
// Create or verify authentication ticket.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/ticket
//
type AccessTicketCreateRequest struct {
	TfaChallenge string `json:"tfa-challenge,omitempty" url:"tfa-challenge,omitempty,optional"` // The signed TFA challenge string the user wants to respond to.
	Username     string `json:"username,omitempty" url:"username,omitempty" validate:"nonzero"` // User name
	NewFormat    bool   `json:"new-format,omitempty" url:"new-format,omitempty,optional"`       // With webauthn the format of half-authenticated tickts changed. New clients should pass 1 here and not worry about the old format. The old format is deprecated and will be retired with PVE-8.0
	Otp          string `json:"otp,omitempty" url:"otp,omitempty,optional"`                     // One-time password for Two-factor authentication.
	Password     string `json:"password,omitempty" url:"password,omitempty" validate:"nonzero"` // The secret password. This can also be a valid ticket.
	Path         string `json:"path,omitempty" url:"path,omitempty,optional"`                   // Verify ticket, and check if user have access 'privs' on 'path'
	Privs        string `json:"privs,omitempty" url:"privs,omitempty,optional"`                 // Verify ticket, and check if user have access 'privs' on 'path'
	Realm        string `json:"realm,omitempty" url:"realm,omitempty,optional"`                 // You can optionally pass the realm using this parameter. Normally the realm is simply added to the username <username>@<relam>.
}

// AccessTicketCreateResponse ticket
// Create or verify authentication ticket.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/ticket
//
type AccessTicketCreateResponse struct {
	Username            string `json:"username,omitempty" url:"username,omitempty" validate:"nonzero"`             //
	CSRFPreventionToken string `json:"CSRFPreventionToken,omitempty" url:"CSRFPreventionToken,omitempty,optional"` //
	Clustername         string `json:"clustername,omitempty" url:"clustername,omitempty,optional"`                 //
	Ticket              string `json:"ticket,omitempty" url:"ticket,omitempty,optional"`                           //
}

// AccessPasswordChangeRequest password - Change user password.
// Change user password.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/password
//
type AccessPasswordChangeRequest struct {
	Password string `json:"password,omitempty" url:"password,omitempty" validate:"nonzero"` // The new password.
	Userid   string `json:"userid,omitempty" url:"userid,omitempty" validate:"nonzero"`     // User ID
}

// AccessPermissionsRequest permissions - Retrieve effective permissions of given user/token.
// Retrieve effective permissions of given user/token.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/access/permissions
//
type AccessPermissionsRequest struct {
	Userid string `json:"userid,omitempty" url:"userid,omitempty,optional"` // User ID or full API token ID
	Path   string `json:"path,omitempty" url:"path,omitempty,optional"`     // Only dump this specific path, not the whole tree.
}

// PoolsCreateRequest pools - Create new pool.
// Create new pool.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/pools
//
type PoolsCreateRequest struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
	Poolid  string `json:"poolid,omitempty" url:"poolid,omitempty" validate:"nonzero"` //
}

// PoolsDeleteRequest {poolid} - Delete pool.
// Delete pool.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/pools/{poolid}
//
type PoolsDeleteRequest struct {
	Poolid string `json:"poolid,omitempty" url:"poolid,omitempty" validate:"nonzero"` //
}

// PoolsReadRequest {poolid} - Get pool configuration.
// Get pool configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/pools/{poolid}
//
type PoolsReadRequest struct {
	Poolid string `json:"poolid,omitempty" url:"poolid,omitempty" validate:"nonzero"` //
	Type   string `json:"type,omitempty" url:"type,omitempty,optional"`               //
}

// PoolsReadResponse {poolid}
// Get pool configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/pools/{poolid}
//
type PoolsReadResponse struct {
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"` //
	Members []struct {
	} `json:"members,omitempty" url:"members,omitempty" validate:"nonzero"` //
}

// PoolsUpdateRequest {poolid} - Update pool data.
// Update pool data.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/pools/{poolid}
//
type PoolsUpdateRequest struct {
	Delete  bool   `json:"delete,omitempty" url:"delete,omitempty,optional"`           // Remove vms/storage (instead of adding it).
	Poolid  string `json:"poolid,omitempty" url:"poolid,omitempty" validate:"nonzero"` //
	Storage string `json:"storage,omitempty" url:"storage,omitempty,optional"`         // List of storage IDs.
	Vms     string `json:"vms,omitempty" url:"vms,omitempty,optional"`                 // List of virtual machines.
	Comment string `json:"comment,omitempty" url:"comment,omitempty,optional"`         //
}

// VersionResponse version
// API version details, including some parts of the global datacenter config.
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/version
//
type VersionResponse struct {
	Console string `json:"console,omitempty" url:"console,omitempty,optional"`           // The default console viewer to use.
	Release string `json:"release,omitempty" url:"release,omitempty" validate:"nonzero"` // The current Proxmox VE point release in `x.y` format.
	Repoid  string `json:"repoid,omitempty" url:"repoid,omitempty" validate:"nonzero"`   // The short git revision from which this version was build.
	Version string `json:"version,omitempty" url:"version,omitempty" validate:"nonzero"` // The full pve-manager package version of this node.
}
