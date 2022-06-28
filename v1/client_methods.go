// Package proxmox was automatically generated
package proxmox

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
)

// ClusterReplicationCreate POST /cluster/replication
// Create a new replication job
//
func (c *Client) ClusterReplicationCreate(req ClusterReplicationCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterReplicationCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterReplicationCreate: %v", err)
	}

	err = c.Request("POST", "/cluster/replication", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterReplicationCreate: %v", err)
	}

	// output
	return nil
}

// ClusterReplicationDelete DELETE /cluster/replication/{id}
// Mark replication job for removal.
//
func (c *Client) ClusterReplicationDelete(req ClusterReplicationDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterReplicationDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/replication/%v", req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterReplicationDelete: %v", err)
	}

	// output
	return nil
}

// ClusterReplicationRead GET /cluster/replication/{id}
// Read replication job configuration.
//
func (c *Client) ClusterReplicationRead(req ClusterReplicationReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterReplicationRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/replication/%v", req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterReplicationRead: %v", err)
	}

	// output
	return nil
}

// ClusterReplicationUpdate PUT /cluster/replication/{id}
// Update replication job configuration.
//
func (c *Client) ClusterReplicationUpdate(req ClusterReplicationUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterReplicationUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterReplicationUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/replication/%v", req.Id), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterReplicationUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterMetricsServerIndex GET /cluster/metrics/server
// List configured metric servers.
//
func (c *Client) ClusterMetricsServerIndex() ([]*ClusterMetricsServerIndexResponse, error) {
	var err error

	var res []*ClusterMetricsServerIndexResponse
	err = c.Request("GET", "/cluster/metrics/server", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterMetricsServerIndex: %v", err)
	}

	// output
	return res, nil
}

// ClusterMetricsServerDelete DELETE /cluster/metrics/server/{id}
// Remove Metric server.
//
func (c *Client) ClusterMetricsServerDelete(req ClusterMetricsServerDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterMetricsServerDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/metrics/server/%v", req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterMetricsServerDelete: %v", err)
	}

	// output
	return nil
}

// ClusterMetricsServerRead GET /cluster/metrics/server/{id}
// Read metric server configuration.
//
func (c *Client) ClusterMetricsServerRead(req ClusterMetricsServerReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterMetricsServerRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/metrics/server/%v", req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterMetricsServerRead: %v", err)
	}

	// output
	return nil
}

// ClusterMetricsServerCreate POST /cluster/metrics/server/{id}
// Create a new external metric server config
//
func (c *Client) ClusterMetricsServerCreate(req ClusterMetricsServerCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterMetricsServerCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterMetricsServerCreate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/cluster/metrics/server/%v", req.Id), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterMetricsServerCreate: %v", err)
	}

	// output
	return nil
}

// ClusterMetricsServerUpdate PUT /cluster/metrics/server/{id}
// Update metric server configuration.
//
func (c *Client) ClusterMetricsServerUpdate(req ClusterMetricsServerUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterMetricsServerUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterMetricsServerUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/metrics/server/%v", req.Id), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterMetricsServerUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterConfigCreate POST /cluster/config
// Generate new cluster configuration. If no links given, default to local IP address as link0.
//
func (c *Client) ClusterConfigCreate(req ClusterConfigCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterConfigCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterConfigCreate: %v", err)
	}

	err = c.Request("POST", "/cluster/config", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterConfigCreate: %v", err)
	}

	// output
	return nil
}

// ClusterConfigNodes GET /cluster/config/nodes
// Corosync node list.
//
func (c *Client) ClusterConfigNodes() ([]*ClusterConfigNodesResponse, error) {
	var err error

	var res []*ClusterConfigNodesResponse
	err = c.Request("GET", "/cluster/config/nodes", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterConfigNodes: %v", err)
	}

	// output
	return res, nil
}

// ClusterConfigNodesDelnode DELETE /cluster/config/nodes/{node}
// Removes a node from the cluster configuration.
//
func (c *Client) ClusterConfigNodesDelnode(req ClusterConfigNodesDelnodeRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterConfigNodesDelnode: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/config/nodes/%v", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterConfigNodesDelnode: %v", err)
	}

	// output
	return nil
}

// ClusterConfigNodesAddnode POST /cluster/config/nodes/{node}
// Adds a node to the cluster configuration. This call is for internal use.
//
func (c *Client) ClusterConfigNodesAddnode(req ClusterConfigNodesAddnodeRequest) (*ClusterConfigNodesAddnodeResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterConfigNodesAddnode: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("ClusterConfigNodesAddnode: %v", err)
	}

	var res *ClusterConfigNodesAddnodeResponse
	err = c.Request("POST", fmt.Sprintf("/cluster/config/nodes/%v", req.Node), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterConfigNodesAddnode: %v", err)
	}

	// output
	return res, nil
}

// ClusterConfigJoinInfo GET /cluster/config/join
// Get information needed to join this cluster over the connected node.
//
func (c *Client) ClusterConfigJoinInfo(req ClusterConfigJoinInfoRequest) (*ClusterConfigJoinInfoResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterConfigJoinInfo: %v", err)
	}

	var res *ClusterConfigJoinInfoResponse
	err = c.Request("GET", "/cluster/config/join", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterConfigJoinInfo: %v", err)
	}

	// output
	return res, nil
}

// ClusterConfigJoin POST /cluster/config/join
// Joins this node into an existing cluster. If no links are given, default to IP resolved by node's hostname on single link (fallback fails for clusters with multiple links).
//
func (c *Client) ClusterConfigJoin(req ClusterConfigJoinRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterConfigJoin: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterConfigJoin: %v", err)
	}

	err = c.Request("POST", "/cluster/config/join", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterConfigJoin: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallGroupsListSecurity GET /cluster/firewall/groups
// List security groups.
//
func (c *Client) ClusterFirewallGroupsListSecurity() ([]*ClusterFirewallGroupsListSecurityResponse, error) {
	var err error

	var res []*ClusterFirewallGroupsListSecurityResponse
	err = c.Request("GET", "/cluster/firewall/groups", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallGroupsListSecurity: %v", err)
	}

	// output
	return res, nil
}

// ClusterFirewallGroupsCreateSecurityGroup POST /cluster/firewall/groups
// Create new security group.
//
func (c *Client) ClusterFirewallGroupsCreateSecurityGroup(req ClusterFirewallGroupsCreateSecurityGroupRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallGroupsCreateSecurityGroup: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallGroupsCreateSecurityGroup: %v", err)
	}

	err = c.Request("POST", "/cluster/firewall/groups", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallGroupsCreateSecurityGroup: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallGroupsDeleteSecurityGroup DELETE /cluster/firewall/groups/{group}
// Delete security group.
//
func (c *Client) ClusterFirewallGroupsDeleteSecurityGroup(req ClusterFirewallGroupsDeleteSecurityGroupRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallGroupsDeleteSecurityGroup: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/firewall/groups/%v", req.Group), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallGroupsDeleteSecurityGroup: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallGroupsGetRules GET /cluster/firewall/groups/{group}
// List rules.
//
func (c *Client) ClusterFirewallGroupsGetRules(req ClusterFirewallGroupsGetRulesRequest) ([]*ClusterFirewallGroupsGetRulesResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterFirewallGroupsGetRules: %v", err)
	}

	var res []*ClusterFirewallGroupsGetRulesResponse
	err = c.Request("GET", fmt.Sprintf("/cluster/firewall/groups/%v", req.Group), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallGroupsGetRules: %v", err)
	}

	// output
	return res, nil
}

// ClusterFirewallGroupsCreateRule POST /cluster/firewall/groups/{group}
// Create new rule.
//
func (c *Client) ClusterFirewallGroupsCreateRule(req ClusterFirewallGroupsCreateRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallGroupsCreateRule: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallGroupsCreateRule: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/cluster/firewall/groups/%v", req.Group), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallGroupsCreateRule: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallGroupsUpdateRule PUT /cluster/firewall/groups/{group}/{pos}
// Modify rule data.
//
func (c *Client) ClusterFirewallGroupsUpdateRule(req ClusterFirewallGroupsUpdateRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallGroupsUpdateRule: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallGroupsUpdateRule: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/firewall/groups/%v/%v", req.Group, req.Pos), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallGroupsUpdateRule: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallGroupsDeleteRule DELETE /cluster/firewall/groups/{group}/{pos}
// Delete rule.
//
func (c *Client) ClusterFirewallGroupsDeleteRule(req ClusterFirewallGroupsDeleteRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallGroupsDeleteRule: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/firewall/groups/%v/%v", req.Group, req.Pos), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallGroupsDeleteRule: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallGroupsGetRule GET /cluster/firewall/groups/{group}/{pos}
// Get single rule data.
//
func (c *Client) ClusterFirewallGroupsGetRule(req ClusterFirewallGroupsGetRuleRequest) (*ClusterFirewallGroupsGetRuleResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterFirewallGroupsGetRule: %v", err)
	}

	var res *ClusterFirewallGroupsGetRuleResponse
	err = c.Request("GET", fmt.Sprintf("/cluster/firewall/groups/%v/%v", req.Group, req.Pos), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallGroupsGetRule: %v", err)
	}

	// output
	return res, nil
}

// ClusterFirewallRulesGet GET /cluster/firewall/rules
// List rules.
//
func (c *Client) ClusterFirewallRulesGet() ([]*ClusterFirewallRulesGetResponse, error) {
	var err error

	var res []*ClusterFirewallRulesGetResponse
	err = c.Request("GET", "/cluster/firewall/rules", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallRulesGet: %v", err)
	}

	// output
	return res, nil
}

// ClusterFirewallRulesCreateRule POST /cluster/firewall/rules
// Create new rule.
//
func (c *Client) ClusterFirewallRulesCreateRule(req ClusterFirewallRulesCreateRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallRulesCreateRule: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallRulesCreateRule: %v", err)
	}

	err = c.Request("POST", "/cluster/firewall/rules", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallRulesCreateRule: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallRulesDeleteRule DELETE /cluster/firewall/rules/{pos}
// Delete rule.
//
func (c *Client) ClusterFirewallRulesDeleteRule(req ClusterFirewallRulesDeleteRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallRulesDeleteRule: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/firewall/rules/%v", req.Pos), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallRulesDeleteRule: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallRulesGetRule GET /cluster/firewall/rules/{pos}
// Get single rule data.
//
func (c *Client) ClusterFirewallRulesGetRule(req ClusterFirewallRulesGetRuleRequest) (*ClusterFirewallRulesGetRuleResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterFirewallRulesGetRule: %v", err)
	}

	var res *ClusterFirewallRulesGetRuleResponse
	err = c.Request("GET", fmt.Sprintf("/cluster/firewall/rules/%v", req.Pos), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallRulesGetRule: %v", err)
	}

	// output
	return res, nil
}

// ClusterFirewallRulesUpdateRule PUT /cluster/firewall/rules/{pos}
// Modify rule data.
//
func (c *Client) ClusterFirewallRulesUpdateRule(req ClusterFirewallRulesUpdateRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallRulesUpdateRule: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallRulesUpdateRule: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/firewall/rules/%v", req.Pos), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallRulesUpdateRule: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallIpsetIndex GET /cluster/firewall/ipset
// List IPSets
//
func (c *Client) ClusterFirewallIpsetIndex() ([]*ClusterFirewallIpsetIndexResponse, error) {
	var err error

	var res []*ClusterFirewallIpsetIndexResponse
	err = c.Request("GET", "/cluster/firewall/ipset", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallIpsetIndex: %v", err)
	}

	// output
	return res, nil
}

// ClusterFirewallIpsetCreate POST /cluster/firewall/ipset
// Create new IPSet
//
func (c *Client) ClusterFirewallIpsetCreate(req ClusterFirewallIpsetCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallIpsetCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallIpsetCreate: %v", err)
	}

	err = c.Request("POST", "/cluster/firewall/ipset", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallIpsetCreate: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallIpsetCreateIp POST /cluster/firewall/ipset/{name}
// Add IP or Network to IPSet.
//
func (c *Client) ClusterFirewallIpsetCreateIp(req ClusterFirewallIpsetCreateIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallIpsetCreateIp: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallIpsetCreateIp: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/cluster/firewall/ipset/%v", req.Name), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallIpsetCreateIp: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallIpsetDelete DELETE /cluster/firewall/ipset/{name}
// Delete IPSet
//
func (c *Client) ClusterFirewallIpsetDelete(req ClusterFirewallIpsetDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallIpsetDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/firewall/ipset/%v", req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallIpsetDelete: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallIpsetGet GET /cluster/firewall/ipset/{name}
// List IPSet content
//
func (c *Client) ClusterFirewallIpsetGet(req ClusterFirewallIpsetGetRequest) ([]*ClusterFirewallIpsetGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterFirewallIpsetGet: %v", err)
	}

	var res []*ClusterFirewallIpsetGetResponse
	err = c.Request("GET", fmt.Sprintf("/cluster/firewall/ipset/%v", req.Name), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallIpsetGet: %v", err)
	}

	// output
	return res, nil
}

// ClusterFirewallIpsetRemoveIp DELETE /cluster/firewall/ipset/{name}/{cidr}
// Remove IP or Network from IPSet.
//
func (c *Client) ClusterFirewallIpsetRemoveIp(req ClusterFirewallIpsetRemoveIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallIpsetRemoveIp: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/firewall/ipset/%v/%v", req.Name, req.Cidr), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallIpsetRemoveIp: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallIpsetReadIp GET /cluster/firewall/ipset/{name}/{cidr}
// Read IP or Network settings from IPSet.
//
func (c *Client) ClusterFirewallIpsetReadIp(req ClusterFirewallIpsetReadIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallIpsetReadIp: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/firewall/ipset/%v/%v", req.Name, req.Cidr), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallIpsetReadIp: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallIpsetUpdateIp PUT /cluster/firewall/ipset/{name}/{cidr}
// Update IP or Network settings
//
func (c *Client) ClusterFirewallIpsetUpdateIp(req ClusterFirewallIpsetUpdateIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallIpsetUpdateIp: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallIpsetUpdateIp: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/firewall/ipset/%v/%v", req.Name, req.Cidr), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallIpsetUpdateIp: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallAliasesGet GET /cluster/firewall/aliases
// List aliases
//
func (c *Client) ClusterFirewallAliasesGet() ([]*ClusterFirewallAliasesGetResponse, error) {
	var err error

	var res []*ClusterFirewallAliasesGetResponse
	err = c.Request("GET", "/cluster/firewall/aliases", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallAliasesGet: %v", err)
	}

	// output
	return res, nil
}

// ClusterFirewallAliasesCreateAlias POST /cluster/firewall/aliases
// Create IP or Network Alias.
//
func (c *Client) ClusterFirewallAliasesCreateAlias(req ClusterFirewallAliasesCreateAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallAliasesCreateAlias: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallAliasesCreateAlias: %v", err)
	}

	err = c.Request("POST", "/cluster/firewall/aliases", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallAliasesCreateAlias: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallAliasesReadAlias GET /cluster/firewall/aliases/{name}
// Read alias.
//
func (c *Client) ClusterFirewallAliasesReadAlias(req ClusterFirewallAliasesReadAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallAliasesReadAlias: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/firewall/aliases/%v", req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallAliasesReadAlias: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallAliasesUpdateAlias PUT /cluster/firewall/aliases/{name}
// Update IP or Network alias.
//
func (c *Client) ClusterFirewallAliasesUpdateAlias(req ClusterFirewallAliasesUpdateAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallAliasesUpdateAlias: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallAliasesUpdateAlias: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/firewall/aliases/%v", req.Name), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallAliasesUpdateAlias: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallAliasesRemoveAlias DELETE /cluster/firewall/aliases/{name}
// Remove IP or Network alias.
//
func (c *Client) ClusterFirewallAliasesRemoveAlias(req ClusterFirewallAliasesRemoveAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallAliasesRemoveAlias: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/firewall/aliases/%v", req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallAliasesRemoveAlias: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallOptionsGet GET /cluster/firewall/options
// Get Firewall options.
//
func (c *Client) ClusterFirewallOptionsGet() (*ClusterFirewallOptionsGetResponse, error) {
	var err error

	var res *ClusterFirewallOptionsGetResponse
	err = c.Request("GET", "/cluster/firewall/options", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallOptionsGet: %v", err)
	}

	// output
	return res, nil
}

// ClusterFirewallOptionsSet PUT /cluster/firewall/options
// Set Firewall options.
//
func (c *Client) ClusterFirewallOptionsSet(req ClusterFirewallOptionsSetRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterFirewallOptionsSet: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterFirewallOptionsSet: %v", err)
	}

	err = c.Request("PUT", "/cluster/firewall/options", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterFirewallOptionsSet: %v", err)
	}

	// output
	return nil
}

// ClusterFirewallMacrosGet GET /cluster/firewall/macros
// List available macros
//
func (c *Client) ClusterFirewallMacrosGet() ([]*ClusterFirewallMacrosGetResponse, error) {
	var err error

	var res []*ClusterFirewallMacrosGetResponse
	err = c.Request("GET", "/cluster/firewall/macros", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallMacrosGet: %v", err)
	}

	// output
	return res, nil
}

// ClusterFirewallRefs GET /cluster/firewall/refs
// Lists possible IPSet/Alias reference which are allowed in source/dest properties.
//
func (c *Client) ClusterFirewallRefs(req ClusterFirewallRefsRequest) ([]*ClusterFirewallRefsResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterFirewallRefs: %v", err)
	}

	var res []*ClusterFirewallRefsResponse
	err = c.Request("GET", "/cluster/firewall/refs", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterFirewallRefs: %v", err)
	}

	// output
	return res, nil
}

// ClusterBackupCreateJob POST /cluster/backup
// Create new vzdump backup job.
//
func (c *Client) ClusterBackupCreateJob(req ClusterBackupCreateJobRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterBackupCreateJob: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterBackupCreateJob: %v", err)
	}

	err = c.Request("POST", "/cluster/backup", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterBackupCreateJob: %v", err)
	}

	// output
	return nil
}

// ClusterBackupDeleteJob DELETE /cluster/backup/{id}
// Delete vzdump backup job definition.
//
func (c *Client) ClusterBackupDeleteJob(req ClusterBackupDeleteJobRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterBackupDeleteJob: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/backup/%v", req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterBackupDeleteJob: %v", err)
	}

	// output
	return nil
}

// ClusterBackupReadJob GET /cluster/backup/{id}
// Read vzdump backup job definition.
//
func (c *Client) ClusterBackupReadJob(req ClusterBackupReadJobRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterBackupReadJob: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/backup/%v", req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterBackupReadJob: %v", err)
	}

	// output
	return nil
}

// ClusterBackupUpdateJob PUT /cluster/backup/{id}
// Update vzdump backup job definition.
//
func (c *Client) ClusterBackupUpdateJob(req ClusterBackupUpdateJobRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterBackupUpdateJob: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterBackupUpdateJob: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/backup/%v", req.Id), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterBackupUpdateJob: %v", err)
	}

	// output
	return nil
}

// ClusterBackupIncludedVolumesGetVolumeBackupIncluded GET /cluster/backup/{id}/included_volumes
// Returns included guests and the backup status of their disks. Optimized to be used in ExtJS tree views.
//
func (c *Client) ClusterBackupIncludedVolumesGetVolumeBackupIncluded(req ClusterBackupIncludedVolumesGetVolumeBackupIncludedRequest) (*ClusterBackupIncludedVolumesGetVolumeBackupIncludedResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterBackupIncludedVolumesGetVolumeBackupIncluded: %v", err)
	}

	var res *ClusterBackupIncludedVolumesGetVolumeBackupIncludedResponse
	err = c.Request("GET", fmt.Sprintf("/cluster/backup/%v/included_volumes", req.Id), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterBackupIncludedVolumesGetVolumeBackupIncluded: %v", err)
	}

	// output
	return res, nil
}

// ClusterBackupInfoNotBackedUpGetGuestsNotInBackup GET /cluster/backup-info/not-backed-up
// Shows all guests which are not covered by any backup job.
//
func (c *Client) ClusterBackupInfoNotBackedUpGetGuestsNotInBackup() ([]*ClusterBackupInfoNotBackedUpGetGuestsNotInBackupResponse, error) {
	var err error

	var res []*ClusterBackupInfoNotBackedUpGetGuestsNotInBackupResponse
	err = c.Request("GET", "/cluster/backup-info/not-backed-up", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterBackupInfoNotBackedUpGetGuestsNotInBackup: %v", err)
	}

	// output
	return res, nil
}

// ClusterHaResourcesCreate POST /cluster/ha/resources
// Create a new HA resource.
//
func (c *Client) ClusterHaResourcesCreate(req ClusterHaResourcesCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterHaResourcesCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterHaResourcesCreate: %v", err)
	}

	err = c.Request("POST", "/cluster/ha/resources", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterHaResourcesCreate: %v", err)
	}

	// output
	return nil
}

// ClusterHaResourcesDelete DELETE /cluster/ha/resources/{sid}
// Delete resource configuration.
//
func (c *Client) ClusterHaResourcesDelete(req ClusterHaResourcesDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterHaResourcesDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/ha/resources/%v", req.Sid), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterHaResourcesDelete: %v", err)
	}

	// output
	return nil
}

// ClusterHaResourcesRead GET /cluster/ha/resources/{sid}
// Read resource configuration.
//
func (c *Client) ClusterHaResourcesRead(req ClusterHaResourcesReadRequest) (*ClusterHaResourcesReadResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterHaResourcesRead: %v", err)
	}

	var res *ClusterHaResourcesReadResponse
	err = c.Request("GET", fmt.Sprintf("/cluster/ha/resources/%v", req.Sid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterHaResourcesRead: %v", err)
	}

	// output
	return res, nil
}

// ClusterHaResourcesUpdate PUT /cluster/ha/resources/{sid}
// Update resource configuration.
//
func (c *Client) ClusterHaResourcesUpdate(req ClusterHaResourcesUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterHaResourcesUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterHaResourcesUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/ha/resources/%v", req.Sid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterHaResourcesUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterHaResourcesMigrate POST /cluster/ha/resources/{sid}/migrate
// Request resource migration (online) to another node.
//
func (c *Client) ClusterHaResourcesMigrate(req ClusterHaResourcesMigrateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterHaResourcesMigrate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterHaResourcesMigrate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/cluster/ha/resources/%v/migrate", req.Sid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterHaResourcesMigrate: %v", err)
	}

	// output
	return nil
}

// ClusterHaResourcesRelocate POST /cluster/ha/resources/{sid}/relocate
// Request resource relocatzion to another node. This stops the service on the old node, and restarts it on the target node.
//
func (c *Client) ClusterHaResourcesRelocate(req ClusterHaResourcesRelocateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterHaResourcesRelocate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterHaResourcesRelocate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/cluster/ha/resources/%v/relocate", req.Sid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterHaResourcesRelocate: %v", err)
	}

	// output
	return nil
}

// ClusterHaGroupsCreate POST /cluster/ha/groups
// Create a new HA group.
//
func (c *Client) ClusterHaGroupsCreate(req ClusterHaGroupsCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterHaGroupsCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterHaGroupsCreate: %v", err)
	}

	err = c.Request("POST", "/cluster/ha/groups", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterHaGroupsCreate: %v", err)
	}

	// output
	return nil
}

// ClusterHaGroupsRead GET /cluster/ha/groups/{group}
// Read ha group configuration.
//
func (c *Client) ClusterHaGroupsRead(req ClusterHaGroupsReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterHaGroupsRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/ha/groups/%v", req.Group), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterHaGroupsRead: %v", err)
	}

	// output
	return nil
}

// ClusterHaGroupsUpdate PUT /cluster/ha/groups/{group}
// Update ha group configuration.
//
func (c *Client) ClusterHaGroupsUpdate(req ClusterHaGroupsUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterHaGroupsUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterHaGroupsUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/ha/groups/%v", req.Group), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterHaGroupsUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterHaGroupsDelete DELETE /cluster/ha/groups/{group}
// Delete ha group configuration.
//
func (c *Client) ClusterHaGroupsDelete(req ClusterHaGroupsDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterHaGroupsDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/ha/groups/%v", req.Group), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterHaGroupsDelete: %v", err)
	}

	// output
	return nil
}

// ClusterAcmePluginsAddPlugin POST /cluster/acme/plugins
// Add ACME plugin configuration.
//
func (c *Client) ClusterAcmePluginsAddPlugin(req ClusterAcmePluginsAddPluginRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterAcmePluginsAddPlugin: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterAcmePluginsAddPlugin: %v", err)
	}

	err = c.Request("POST", "/cluster/acme/plugins", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterAcmePluginsAddPlugin: %v", err)
	}

	// output
	return nil
}

// ClusterAcmePluginsDeletePlugin DELETE /cluster/acme/plugins/{id}
// Delete ACME plugin configuration.
//
func (c *Client) ClusterAcmePluginsDeletePlugin(req ClusterAcmePluginsDeletePluginRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterAcmePluginsDeletePlugin: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/acme/plugins/%v", req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterAcmePluginsDeletePlugin: %v", err)
	}

	// output
	return nil
}

// ClusterAcmePluginsGetPluginConfig GET /cluster/acme/plugins/{id}
// Get ACME plugin configuration.
//
func (c *Client) ClusterAcmePluginsGetPluginConfig(req ClusterAcmePluginsGetPluginConfigRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterAcmePluginsGetPluginConfig: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/acme/plugins/%v", req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterAcmePluginsGetPluginConfig: %v", err)
	}

	// output
	return nil
}

// ClusterAcmePluginsUpdatePlugin PUT /cluster/acme/plugins/{id}
// Update ACME plugin configuration.
//
func (c *Client) ClusterAcmePluginsUpdatePlugin(req ClusterAcmePluginsUpdatePluginRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterAcmePluginsUpdatePlugin: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterAcmePluginsUpdatePlugin: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/acme/plugins/%v", req.Id), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterAcmePluginsUpdatePlugin: %v", err)
	}

	// output
	return nil
}

// ClusterAcmeAccountRegister POST /cluster/acme/account
// Register a new ACME account with CA.
//
func (c *Client) ClusterAcmeAccountRegister(req ClusterAcmeAccountRegisterRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterAcmeAccountRegister: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterAcmeAccountRegister: %v", err)
	}

	err = c.Request("POST", "/cluster/acme/account", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterAcmeAccountRegister: %v", err)
	}

	// output
	return nil
}

// ClusterAcmeAccountDeactivate DELETE /cluster/acme/account/{name}
// Deactivate existing ACME account at CA.
//
func (c *Client) ClusterAcmeAccountDeactivate(req ClusterAcmeAccountDeactivateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterAcmeAccountDeactivate: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/acme/account/%v", req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterAcmeAccountDeactivate: %v", err)
	}

	// output
	return nil
}

// ClusterAcmeAccountGet GET /cluster/acme/account/{name}
// Return existing ACME account information.
//
func (c *Client) ClusterAcmeAccountGet(req ClusterAcmeAccountGetRequest) (*ClusterAcmeAccountGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterAcmeAccountGet: %v", err)
	}

	var res *ClusterAcmeAccountGetResponse
	err = c.Request("GET", fmt.Sprintf("/cluster/acme/account/%v", req.Name), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterAcmeAccountGet: %v", err)
	}

	// output
	return res, nil
}

// ClusterAcmeAccountUpdate PUT /cluster/acme/account/{name}
// Update existing ACME account information with CA. Note: not specifying any new account information triggers a refresh.
//
func (c *Client) ClusterAcmeAccountUpdate(req ClusterAcmeAccountUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterAcmeAccountUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterAcmeAccountUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/acme/account/%v", req.Name), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterAcmeAccountUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterAcmeTosGet GET /cluster/acme/tos
// Retrieve ACME TermsOfService URL from CA.
//
func (c *Client) ClusterAcmeTosGet(req ClusterAcmeTosGetRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterAcmeTosGet: %v", err)
	}

	err = c.Request("GET", "/cluster/acme/tos", nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterAcmeTosGet: %v", err)
	}

	// output
	return nil
}

// ClusterAcmeDirectoriesGet GET /cluster/acme/directories
// Get named known ACME directory endpoints.
//
func (c *Client) ClusterAcmeDirectoriesGet() ([]*ClusterAcmeDirectoriesGetResponse, error) {
	var err error

	var res []*ClusterAcmeDirectoriesGetResponse
	err = c.Request("GET", "/cluster/acme/directories", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterAcmeDirectoriesGet: %v", err)
	}

	// output
	return res, nil
}

// ClusterAcmeChallengeSchemaChallengeschema GET /cluster/acme/challenge-schema
// Get schema of ACME challenge types.
//
func (c *Client) ClusterAcmeChallengeSchemaChallengeschema() ([]*ClusterAcmeChallengeSchemaChallengeschemaResponse, error) {
	var err error

	var res []*ClusterAcmeChallengeSchemaChallengeschemaResponse
	err = c.Request("GET", "/cluster/acme/challenge-schema", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterAcmeChallengeSchemaChallengeschema: %v", err)
	}

	// output
	return res, nil
}

// ClusterCephMetadata GET /cluster/ceph/metadata
// Get ceph metadata.
//
func (c *Client) ClusterCephMetadata(req ClusterCephMetadataRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterCephMetadata: %v", err)
	}

	err = c.Request("GET", "/cluster/ceph/metadata", nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterCephMetadata: %v", err)
	}

	// output
	return nil
}

// ClusterCephFlagsGetAll GET /cluster/ceph/flags
// get the status of all ceph flags
//
func (c *Client) ClusterCephFlagsGetAll() ([]*ClusterCephFlagsGetAllResponse, error) {
	var err error

	var res []*ClusterCephFlagsGetAllResponse
	err = c.Request("GET", "/cluster/ceph/flags", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterCephFlagsGetAll: %v", err)
	}

	// output
	return res, nil
}

// ClusterCephFlagsSet PUT /cluster/ceph/flags
// Set/Unset multiple ceph flags at once.
//
func (c *Client) ClusterCephFlagsSet(req ClusterCephFlagsSetRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterCephFlagsSet: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterCephFlagsSet: %v", err)
	}

	err = c.Request("PUT", "/cluster/ceph/flags", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterCephFlagsSet: %v", err)
	}

	// output
	return nil
}

// ClusterCephFlagsGetFlag GET /cluster/ceph/flags/{flag}
// Get the status of a specific ceph flag.
//
func (c *Client) ClusterCephFlagsGetFlag(req ClusterCephFlagsGetFlagRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterCephFlagsGetFlag: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/ceph/flags/%v", req.Flag), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterCephFlagsGetFlag: %v", err)
	}

	// output
	return nil
}

// ClusterCephFlagsUpdateFlag PUT /cluster/ceph/flags/{flag}
// Set or clear (unset) a specific ceph flag
//
func (c *Client) ClusterCephFlagsUpdateFlag(req ClusterCephFlagsUpdateFlagRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterCephFlagsUpdateFlag: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterCephFlagsUpdateFlag: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/ceph/flags/%v", req.Flag), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterCephFlagsUpdateFlag: %v", err)
	}

	// output
	return nil
}

// ClusterJobsScheduleAnalyze GET /cluster/jobs/schedule-analyze
// Returns a list of future schedule runtimes.
//
func (c *Client) ClusterJobsScheduleAnalyze(req ClusterJobsScheduleAnalyzeRequest) ([]*ClusterJobsScheduleAnalyzeResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterJobsScheduleAnalyze: %v", err)
	}

	var res []*ClusterJobsScheduleAnalyzeResponse
	err = c.Request("GET", "/cluster/jobs/schedule-analyze", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterJobsScheduleAnalyze: %v", err)
	}

	// output
	return res, nil
}

// ClusterSdnReload PUT /cluster/sdn
// Apply sdn controller changes && reload.
//
func (c *Client) ClusterSdnReload() error {
	var err error

	err = c.Request("PUT", "/cluster/sdn", nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnReload: %v", err)
	}

	// output
	return nil
}

// ClusterSdnVnetsCreate POST /cluster/sdn/vnets
// Create a new sdn vnet object.
//
func (c *Client) ClusterSdnVnetsCreate(req ClusterSdnVnetsCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnVnetsCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsCreate: %v", err)
	}

	err = c.Request("POST", "/cluster/sdn/vnets", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsCreate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnVnetsDelete DELETE /cluster/sdn/vnets/{vnet}
// Delete sdn vnet object configuration.
//
func (c *Client) ClusterSdnVnetsDelete(req ClusterSdnVnetsDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnVnetsDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/sdn/vnets/%v", req.Vnet), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsDelete: %v", err)
	}

	// output
	return nil
}

// ClusterSdnVnetsRead GET /cluster/sdn/vnets/{vnet}
// Read sdn vnet configuration.
//
func (c *Client) ClusterSdnVnetsRead(req ClusterSdnVnetsReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnVnetsRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/sdn/vnets/%v", req.Vnet), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsRead: %v", err)
	}

	// output
	return nil
}

// ClusterSdnVnetsUpdate PUT /cluster/sdn/vnets/{vnet}
// Update sdn vnet object configuration.
//
func (c *Client) ClusterSdnVnetsUpdate(req ClusterSdnVnetsUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnVnetsUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/sdn/vnets/%v", req.Vnet), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnVnetsSubnetsCreate POST /cluster/sdn/vnets/{vnet}/subnets
// Create a new sdn subnet object.
//
func (c *Client) ClusterSdnVnetsSubnetsCreate(req ClusterSdnVnetsSubnetsCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnVnetsSubnetsCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsSubnetsCreate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/cluster/sdn/vnets/%v/subnets", req.Vnet), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsSubnetsCreate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnVnetsSubnetsDelete DELETE /cluster/sdn/vnets/{vnet}/subnets/{subnet}
// Delete sdn subnet object configuration.
//
func (c *Client) ClusterSdnVnetsSubnetsDelete(req ClusterSdnVnetsSubnetsDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnVnetsSubnetsDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/sdn/vnets/%v/subnets/%v", req.Vnet, req.Subnet), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsSubnetsDelete: %v", err)
	}

	// output
	return nil
}

// ClusterSdnVnetsSubnetsRead GET /cluster/sdn/vnets/{vnet}/subnets/{subnet}
// Read sdn subnet configuration.
//
func (c *Client) ClusterSdnVnetsSubnetsRead(req ClusterSdnVnetsSubnetsReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnVnetsSubnetsRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/sdn/vnets/%v/subnets/%v", req.Vnet, req.Subnet), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsSubnetsRead: %v", err)
	}

	// output
	return nil
}

// ClusterSdnVnetsSubnetsUpdate PUT /cluster/sdn/vnets/{vnet}/subnets/{subnet}
// Update sdn subnet object configuration.
//
func (c *Client) ClusterSdnVnetsSubnetsUpdate(req ClusterSdnVnetsSubnetsUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnVnetsSubnetsUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsSubnetsUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/sdn/vnets/%v/subnets/%v", req.Vnet, req.Subnet), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnVnetsSubnetsUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnZonesCreate POST /cluster/sdn/zones
// Create a new sdn zone object.
//
func (c *Client) ClusterSdnZonesCreate(req ClusterSdnZonesCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnZonesCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnZonesCreate: %v", err)
	}

	err = c.Request("POST", "/cluster/sdn/zones", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnZonesCreate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnZonesDelete DELETE /cluster/sdn/zones/{zone}
// Delete sdn zone object configuration.
//
func (c *Client) ClusterSdnZonesDelete(req ClusterSdnZonesDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnZonesDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/sdn/zones/%v", req.Zone), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnZonesDelete: %v", err)
	}

	// output
	return nil
}

// ClusterSdnZonesRead GET /cluster/sdn/zones/{zone}
// Read sdn zone configuration.
//
func (c *Client) ClusterSdnZonesRead(req ClusterSdnZonesReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnZonesRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/sdn/zones/%v", req.Zone), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnZonesRead: %v", err)
	}

	// output
	return nil
}

// ClusterSdnZonesUpdate PUT /cluster/sdn/zones/{zone}
// Update sdn zone object configuration.
//
func (c *Client) ClusterSdnZonesUpdate(req ClusterSdnZonesUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnZonesUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnZonesUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/sdn/zones/%v", req.Zone), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnZonesUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnControllersCreate POST /cluster/sdn/controllers
// Create a new sdn controller object.
//
func (c *Client) ClusterSdnControllersCreate(req ClusterSdnControllersCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnControllersCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnControllersCreate: %v", err)
	}

	err = c.Request("POST", "/cluster/sdn/controllers", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnControllersCreate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnControllersUpdate PUT /cluster/sdn/controllers/{controller}
// Update sdn controller object configuration.
//
func (c *Client) ClusterSdnControllersUpdate(req ClusterSdnControllersUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnControllersUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnControllersUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/sdn/controllers/%v", req.Controller), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnControllersUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnControllersDelete DELETE /cluster/sdn/controllers/{controller}
// Delete sdn controller object configuration.
//
func (c *Client) ClusterSdnControllersDelete(req ClusterSdnControllersDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnControllersDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/sdn/controllers/%v", req.Controller), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnControllersDelete: %v", err)
	}

	// output
	return nil
}

// ClusterSdnControllersRead GET /cluster/sdn/controllers/{controller}
// Read sdn controller configuration.
//
func (c *Client) ClusterSdnControllersRead(req ClusterSdnControllersReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnControllersRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/sdn/controllers/%v", req.Controller), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnControllersRead: %v", err)
	}

	// output
	return nil
}

// ClusterSdnIpamsCreate POST /cluster/sdn/ipams
// Create a new sdn ipam object.
//
func (c *Client) ClusterSdnIpamsCreate(req ClusterSdnIpamsCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnIpamsCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnIpamsCreate: %v", err)
	}

	err = c.Request("POST", "/cluster/sdn/ipams", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnIpamsCreate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnIpamsDelete DELETE /cluster/sdn/ipams/{ipam}
// Delete sdn ipam object configuration.
//
func (c *Client) ClusterSdnIpamsDelete(req ClusterSdnIpamsDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnIpamsDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/sdn/ipams/%v", req.Ipam), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnIpamsDelete: %v", err)
	}

	// output
	return nil
}

// ClusterSdnIpamsRead GET /cluster/sdn/ipams/{ipam}
// Read sdn ipam configuration.
//
func (c *Client) ClusterSdnIpamsRead(req ClusterSdnIpamsReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnIpamsRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/sdn/ipams/%v", req.Ipam), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnIpamsRead: %v", err)
	}

	// output
	return nil
}

// ClusterSdnIpamsUpdate PUT /cluster/sdn/ipams/{ipam}
// Update sdn ipam object configuration.
//
func (c *Client) ClusterSdnIpamsUpdate(req ClusterSdnIpamsUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnIpamsUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnIpamsUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/sdn/ipams/%v", req.Ipam), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnIpamsUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnDnsCreate POST /cluster/sdn/dns
// Create a new sdn dns object.
//
func (c *Client) ClusterSdnDnsCreate(req ClusterSdnDnsCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnDnsCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnDnsCreate: %v", err)
	}

	err = c.Request("POST", "/cluster/sdn/dns", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnDnsCreate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnDnsUpdate PUT /cluster/sdn/dns/{dns}
// Update sdn dns object configuration.
//
func (c *Client) ClusterSdnDnsUpdate(req ClusterSdnDnsUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnDnsUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterSdnDnsUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/cluster/sdn/dns/%v", req.Dns), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnDnsUpdate: %v", err)
	}

	// output
	return nil
}

// ClusterSdnDnsDelete DELETE /cluster/sdn/dns/{dns}
// Delete sdn dns object configuration.
//
func (c *Client) ClusterSdnDnsDelete(req ClusterSdnDnsDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnDnsDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/cluster/sdn/dns/%v", req.Dns), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnDnsDelete: %v", err)
	}

	// output
	return nil
}

// ClusterSdnDnsRead GET /cluster/sdn/dns/{dns}
// Read sdn dns configuration.
//
func (c *Client) ClusterSdnDnsRead(req ClusterSdnDnsReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterSdnDnsRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/cluster/sdn/dns/%v", req.Dns), nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterSdnDnsRead: %v", err)
	}

	// output
	return nil
}

// ClusterLog GET /cluster/log
// Read cluster log
//
func (c *Client) ClusterLog(req ClusterLogRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterLog: %v", err)
	}

	err = c.Request("GET", "/cluster/log", nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterLog: %v", err)
	}

	// output
	return nil
}

// ClusterResources GET /cluster/resources
// Resources index (cluster wide).
//
func (c *Client) ClusterResources(req ClusterResourcesRequest) ([]*ClusterResourcesResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("ClusterResources: %v", err)
	}

	var res []*ClusterResourcesResponse
	err = c.Request("GET", "/cluster/resources", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterResources: %v", err)
	}

	// output
	return res, nil
}

// ClusterTasks GET /cluster/tasks
// List recent tasks (cluster wide).
//
func (c *Client) ClusterTasks() ([]*ClusterTasksResponse, error) {
	var err error

	var res []*ClusterTasksResponse
	err = c.Request("GET", "/cluster/tasks", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterTasks: %v", err)
	}

	// output
	return res, nil
}

// ClusterOptionsSet PUT /cluster/options
// Set datacenter options.
//
func (c *Client) ClusterOptionsSet(req ClusterOptionsSetRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterOptionsSet: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("ClusterOptionsSet: %v", err)
	}

	err = c.Request("PUT", "/cluster/options", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("ClusterOptionsSet: %v", err)
	}

	// output
	return nil
}

// ClusterStatusGet GET /cluster/status
// Get cluster status information.
//
func (c *Client) ClusterStatusGet() ([]*ClusterStatusGetResponse, error) {
	var err error

	var res []*ClusterStatusGetResponse
	err = c.Request("GET", "/cluster/status", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("ClusterStatusGet: %v", err)
	}

	// output
	return res, nil
}

// ClusterNextid GET /cluster/nextid
// Get next free VMID. Pass a VMID to assert that its free (at time of check).
//
func (c *Client) ClusterNextid(req ClusterNextidRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("ClusterNextid: %v", err)
	}

	err = c.Request("GET", "/cluster/nextid", nil, nil)

	if err != nil {
		return fmt.Errorf("ClusterNextid: %v", err)
	}

	// output
	return nil
}

// NodesQemuVmlist GET /nodes/{node}/qemu
// Virtual machine index (per node).
//
func (c *Client) NodesQemuVmlist(req NodesQemuVmlistRequest) ([]*NodesQemuVmlistResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuVmlist: %v", err)
	}

	var res []*NodesQemuVmlistResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuVmlist: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuCreateVm POST /nodes/{node}/qemu
// Create or restore a virtual machine.
//
func (c *Client) NodesQemuCreateVm(req NodesQemuCreateVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuCreateVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuCreateVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuCreateVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuDestroyVm DELETE /nodes/{node}/qemu/{vmid}
// Destroy the VM and  all used/owned volumes. Removes any VM specific permissions and firewall rules
//
func (c *Client) NodesQemuDestroyVm(req NodesQemuDestroyVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuDestroyVm: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/qemu/%v", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuDestroyVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuVmdiridx GET /nodes/{node}/qemu/{vmid}
// Directory index
//
func (c *Client) NodesQemuVmdiridx(req NodesQemuVmdiridxRequest) ([]*NodesQemuVmdiridxResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuVmdiridx: %v", err)
	}

	var res []*NodesQemuVmdiridxResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuVmdiridx: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuFirewallRulesGet GET /nodes/{node}/qemu/{vmid}/firewall/rules
// List rules.
//
func (c *Client) NodesQemuFirewallRulesGet(req NodesQemuFirewallRulesGetRequest) ([]*NodesQemuFirewallRulesGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuFirewallRulesGet: %v", err)
	}

	var res []*NodesQemuFirewallRulesGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/rules", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuFirewallRulesGet: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuFirewallRulesCreateRule POST /nodes/{node}/qemu/{vmid}/firewall/rules
// Create new rule.
//
func (c *Client) NodesQemuFirewallRulesCreateRule(req NodesQemuFirewallRulesCreateRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallRulesCreateRule: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuFirewallRulesCreateRule: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/rules", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallRulesCreateRule: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallRulesDeleteRule DELETE /nodes/{node}/qemu/{vmid}/firewall/rules/{pos}
// Delete rule.
//
func (c *Client) NodesQemuFirewallRulesDeleteRule(req NodesQemuFirewallRulesDeleteRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallRulesDeleteRule: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/rules/%v", req.Node, req.Vmid, req.Pos), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallRulesDeleteRule: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallRulesGetRule GET /nodes/{node}/qemu/{vmid}/firewall/rules/{pos}
// Get single rule data.
//
func (c *Client) NodesQemuFirewallRulesGetRule(req NodesQemuFirewallRulesGetRuleRequest) (*NodesQemuFirewallRulesGetRuleResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuFirewallRulesGetRule: %v", err)
	}

	var res *NodesQemuFirewallRulesGetRuleResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/rules/%v", req.Node, req.Vmid, req.Pos), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuFirewallRulesGetRule: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuFirewallRulesUpdateRule PUT /nodes/{node}/qemu/{vmid}/firewall/rules/{pos}
// Modify rule data.
//
func (c *Client) NodesQemuFirewallRulesUpdateRule(req NodesQemuFirewallRulesUpdateRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallRulesUpdateRule: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuFirewallRulesUpdateRule: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/rules/%v", req.Node, req.Vmid, req.Pos), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallRulesUpdateRule: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallAliasesCreateAlias POST /nodes/{node}/qemu/{vmid}/firewall/aliases
// Create IP or Network Alias.
//
func (c *Client) NodesQemuFirewallAliasesCreateAlias(req NodesQemuFirewallAliasesCreateAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallAliasesCreateAlias: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuFirewallAliasesCreateAlias: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/aliases", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallAliasesCreateAlias: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallAliasesGet GET /nodes/{node}/qemu/{vmid}/firewall/aliases
// List aliases
//
func (c *Client) NodesQemuFirewallAliasesGet(req NodesQemuFirewallAliasesGetRequest) ([]*NodesQemuFirewallAliasesGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuFirewallAliasesGet: %v", err)
	}

	var res []*NodesQemuFirewallAliasesGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/aliases", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuFirewallAliasesGet: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuFirewallAliasesReadAlias GET /nodes/{node}/qemu/{vmid}/firewall/aliases/{name}
// Read alias.
//
func (c *Client) NodesQemuFirewallAliasesReadAlias(req NodesQemuFirewallAliasesReadAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallAliasesReadAlias: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/aliases/%v", req.Node, req.Vmid, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallAliasesReadAlias: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallAliasesUpdateAlias PUT /nodes/{node}/qemu/{vmid}/firewall/aliases/{name}
// Update IP or Network alias.
//
func (c *Client) NodesQemuFirewallAliasesUpdateAlias(req NodesQemuFirewallAliasesUpdateAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallAliasesUpdateAlias: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuFirewallAliasesUpdateAlias: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/aliases/%v", req.Node, req.Vmid, req.Name), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallAliasesUpdateAlias: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallAliasesRemoveAlias DELETE /nodes/{node}/qemu/{vmid}/firewall/aliases/{name}
// Remove IP or Network alias.
//
func (c *Client) NodesQemuFirewallAliasesRemoveAlias(req NodesQemuFirewallAliasesRemoveAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallAliasesRemoveAlias: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/aliases/%v", req.Node, req.Vmid, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallAliasesRemoveAlias: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallIpsetCreate POST /nodes/{node}/qemu/{vmid}/firewall/ipset
// Create new IPSet
//
func (c *Client) NodesQemuFirewallIpsetCreate(req NodesQemuFirewallIpsetCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallIpsetCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuFirewallIpsetCreate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/ipset", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallIpsetCreate: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallIpsetIndex GET /nodes/{node}/qemu/{vmid}/firewall/ipset
// List IPSets
//
func (c *Client) NodesQemuFirewallIpsetIndex(req NodesQemuFirewallIpsetIndexRequest) ([]*NodesQemuFirewallIpsetIndexResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuFirewallIpsetIndex: %v", err)
	}

	var res []*NodesQemuFirewallIpsetIndexResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/ipset", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuFirewallIpsetIndex: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuFirewallIpsetDelete DELETE /nodes/{node}/qemu/{vmid}/firewall/ipset/{name}
// Delete IPSet
//
func (c *Client) NodesQemuFirewallIpsetDelete(req NodesQemuFirewallIpsetDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallIpsetDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/ipset/%v", req.Node, req.Vmid, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallIpsetDelete: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallIpsetGet GET /nodes/{node}/qemu/{vmid}/firewall/ipset/{name}
// List IPSet content
//
func (c *Client) NodesQemuFirewallIpsetGet(req NodesQemuFirewallIpsetGetRequest) ([]*NodesQemuFirewallIpsetGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuFirewallIpsetGet: %v", err)
	}

	var res []*NodesQemuFirewallIpsetGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/ipset/%v", req.Node, req.Vmid, req.Name), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuFirewallIpsetGet: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuFirewallIpsetCreateIp POST /nodes/{node}/qemu/{vmid}/firewall/ipset/{name}
// Add IP or Network to IPSet.
//
func (c *Client) NodesQemuFirewallIpsetCreateIp(req NodesQemuFirewallIpsetCreateIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallIpsetCreateIp: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuFirewallIpsetCreateIp: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/ipset/%v", req.Node, req.Vmid, req.Name), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallIpsetCreateIp: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallIpsetReadIp GET /nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}
// Read IP or Network settings from IPSet.
//
func (c *Client) NodesQemuFirewallIpsetReadIp(req NodesQemuFirewallIpsetReadIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallIpsetReadIp: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/ipset/%v/%v", req.Node, req.Vmid, req.Name, req.Cidr), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallIpsetReadIp: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallIpsetUpdateIp PUT /nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}
// Update IP or Network settings
//
func (c *Client) NodesQemuFirewallIpsetUpdateIp(req NodesQemuFirewallIpsetUpdateIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallIpsetUpdateIp: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuFirewallIpsetUpdateIp: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/ipset/%v/%v", req.Node, req.Vmid, req.Name, req.Cidr), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallIpsetUpdateIp: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallIpsetRemoveIp DELETE /nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}
// Remove IP or Network from IPSet.
//
func (c *Client) NodesQemuFirewallIpsetRemoveIp(req NodesQemuFirewallIpsetRemoveIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallIpsetRemoveIp: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/ipset/%v/%v", req.Node, req.Vmid, req.Name, req.Cidr), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallIpsetRemoveIp: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallOptionsGet GET /nodes/{node}/qemu/{vmid}/firewall/options
// Get VM firewall options.
//
func (c *Client) NodesQemuFirewallOptionsGet(req NodesQemuFirewallOptionsGetRequest) (*NodesQemuFirewallOptionsGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuFirewallOptionsGet: %v", err)
	}

	var res *NodesQemuFirewallOptionsGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/options", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuFirewallOptionsGet: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuFirewallOptionsSet PUT /nodes/{node}/qemu/{vmid}/firewall/options
// Set Firewall options.
//
func (c *Client) NodesQemuFirewallOptionsSet(req NodesQemuFirewallOptionsSetRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuFirewallOptionsSet: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuFirewallOptionsSet: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/options", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuFirewallOptionsSet: %v", err)
	}

	// output
	return nil
}

// NodesQemuFirewallLog GET /nodes/{node}/qemu/{vmid}/firewall/log
// Read firewall log
//
func (c *Client) NodesQemuFirewallLog(req NodesQemuFirewallLogRequest) ([]*NodesQemuFirewallLogResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuFirewallLog: %v", err)
	}

	var res []*NodesQemuFirewallLogResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/log", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuFirewallLog: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuFirewallRefs GET /nodes/{node}/qemu/{vmid}/firewall/refs
// Lists possible IPSet/Alias reference which are allowed in source/dest properties.
//
func (c *Client) NodesQemuFirewallRefs(req NodesQemuFirewallRefsRequest) ([]*NodesQemuFirewallRefsResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuFirewallRefs: %v", err)
	}

	var res []*NodesQemuFirewallRefsResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/firewall/refs", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuFirewallRefs: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuAgent POST /nodes/{node}/qemu/{vmid}/agent
// Execute Qemu Guest Agent commands.
//
func (c *Client) NodesQemuAgent(req NodesQemuAgentRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgent: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgent: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgent: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentFsfreezeFreeze POST /nodes/{node}/qemu/{vmid}/agent/fsfreeze-freeze
// Execute fsfreeze-freeze.
//
func (c *Client) NodesQemuAgentFsfreezeFreeze(req NodesQemuAgentFsfreezeFreezeRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentFsfreezeFreeze: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentFsfreezeFreeze: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/fsfreeze-freeze", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentFsfreezeFreeze: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentFsfreezeStatus POST /nodes/{node}/qemu/{vmid}/agent/fsfreeze-status
// Execute fsfreeze-status.
//
func (c *Client) NodesQemuAgentFsfreezeStatus(req NodesQemuAgentFsfreezeStatusRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentFsfreezeStatus: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentFsfreezeStatus: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/fsfreeze-status", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentFsfreezeStatus: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentFsfreezeThaw POST /nodes/{node}/qemu/{vmid}/agent/fsfreeze-thaw
// Execute fsfreeze-thaw.
//
func (c *Client) NodesQemuAgentFsfreezeThaw(req NodesQemuAgentFsfreezeThawRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentFsfreezeThaw: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentFsfreezeThaw: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/fsfreeze-thaw", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentFsfreezeThaw: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentFstrim POST /nodes/{node}/qemu/{vmid}/agent/fstrim
// Execute fstrim.
//
func (c *Client) NodesQemuAgentFstrim(req NodesQemuAgentFstrimRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentFstrim: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentFstrim: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/fstrim", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentFstrim: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentGetFsinfo GET /nodes/{node}/qemu/{vmid}/agent/get-fsinfo
// Execute get-fsinfo.
//
func (c *Client) NodesQemuAgentGetFsinfo(req NodesQemuAgentGetFsinfoRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentGetFsinfo: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/get-fsinfo", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentGetFsinfo: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentGetHostName GET /nodes/{node}/qemu/{vmid}/agent/get-host-name
// Execute get-host-name.
//
func (c *Client) NodesQemuAgentGetHostName(req NodesQemuAgentGetHostNameRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentGetHostName: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/get-host-name", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentGetHostName: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentGetMemoryBlockInfo GET /nodes/{node}/qemu/{vmid}/agent/get-memory-block-info
// Execute get-memory-block-info.
//
func (c *Client) NodesQemuAgentGetMemoryBlockInfo(req NodesQemuAgentGetMemoryBlockInfoRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentGetMemoryBlockInfo: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/get-memory-block-info", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentGetMemoryBlockInfo: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentGetMemoryBlocks GET /nodes/{node}/qemu/{vmid}/agent/get-memory-blocks
// Execute get-memory-blocks.
//
func (c *Client) NodesQemuAgentGetMemoryBlocks(req NodesQemuAgentGetMemoryBlocksRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentGetMemoryBlocks: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/get-memory-blocks", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentGetMemoryBlocks: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentGetOsinfo GET /nodes/{node}/qemu/{vmid}/agent/get-osinfo
// Execute get-osinfo.
//
func (c *Client) NodesQemuAgentGetOsinfo(req NodesQemuAgentGetOsinfoRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentGetOsinfo: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/get-osinfo", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentGetOsinfo: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentGetTime GET /nodes/{node}/qemu/{vmid}/agent/get-time
// Execute get-time.
//
func (c *Client) NodesQemuAgentGetTime(req NodesQemuAgentGetTimeRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentGetTime: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/get-time", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentGetTime: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentGetTimezone GET /nodes/{node}/qemu/{vmid}/agent/get-timezone
// Execute get-timezone.
//
func (c *Client) NodesQemuAgentGetTimezone(req NodesQemuAgentGetTimezoneRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentGetTimezone: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/get-timezone", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentGetTimezone: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentGetUsers GET /nodes/{node}/qemu/{vmid}/agent/get-users
// Execute get-users.
//
func (c *Client) NodesQemuAgentGetUsers(req NodesQemuAgentGetUsersRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentGetUsers: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/get-users", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentGetUsers: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentGetVcpus GET /nodes/{node}/qemu/{vmid}/agent/get-vcpus
// Execute get-vcpus.
//
func (c *Client) NodesQemuAgentGetVcpus(req NodesQemuAgentGetVcpusRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentGetVcpus: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/get-vcpus", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentGetVcpus: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentInfo GET /nodes/{node}/qemu/{vmid}/agent/info
// Execute info.
//
func (c *Client) NodesQemuAgentInfo(req NodesQemuAgentInfoRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentInfo: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/info", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentInfo: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentNetworkGetInterfaces GET /nodes/{node}/qemu/{vmid}/agent/network-get-interfaces
// Execute network-get-interfaces.
//
func (c *Client) NodesQemuAgentNetworkGetInterfaces(req NodesQemuAgentNetworkGetInterfacesRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentNetworkGetInterfaces: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/network-get-interfaces", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentNetworkGetInterfaces: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentPing POST /nodes/{node}/qemu/{vmid}/agent/ping
// Execute ping.
//
func (c *Client) NodesQemuAgentPing(req NodesQemuAgentPingRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentPing: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentPing: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/ping", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentPing: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentShutdown POST /nodes/{node}/qemu/{vmid}/agent/shutdown
// Execute shutdown.
//
func (c *Client) NodesQemuAgentShutdown(req NodesQemuAgentShutdownRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentShutdown: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentShutdown: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/shutdown", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentShutdown: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentSuspendDisk POST /nodes/{node}/qemu/{vmid}/agent/suspend-disk
// Execute suspend-disk.
//
func (c *Client) NodesQemuAgentSuspendDisk(req NodesQemuAgentSuspendDiskRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentSuspendDisk: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentSuspendDisk: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/suspend-disk", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentSuspendDisk: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentSuspendHybrid POST /nodes/{node}/qemu/{vmid}/agent/suspend-hybrid
// Execute suspend-hybrid.
//
func (c *Client) NodesQemuAgentSuspendHybrid(req NodesQemuAgentSuspendHybridRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentSuspendHybrid: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentSuspendHybrid: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/suspend-hybrid", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentSuspendHybrid: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentSuspendRam POST /nodes/{node}/qemu/{vmid}/agent/suspend-ram
// Execute suspend-ram.
//
func (c *Client) NodesQemuAgentSuspendRam(req NodesQemuAgentSuspendRamRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentSuspendRam: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentSuspendRam: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/suspend-ram", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentSuspendRam: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentSetUserPassword POST /nodes/{node}/qemu/{vmid}/agent/set-user-password
// Sets the password for the given user to the given password
//
func (c *Client) NodesQemuAgentSetUserPassword(req NodesQemuAgentSetUserPasswordRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentSetUserPassword: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentSetUserPassword: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/set-user-password", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentSetUserPassword: %v", err)
	}

	// output
	return nil
}

// NodesQemuAgentExec POST /nodes/{node}/qemu/{vmid}/agent/exec
// Executes the given command in the vm via the guest-agent and returns an object with the pid.
//
func (c *Client) NodesQemuAgentExec(req NodesQemuAgentExecRequest) (*NodesQemuAgentExecResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuAgentExec: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesQemuAgentExec: %v", err)
	}

	var res *NodesQemuAgentExecResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/exec", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuAgentExec: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuAgentExecStatus GET /nodes/{node}/qemu/{vmid}/agent/exec-status
// Gets the status of the given pid started by the guest-agent
//
func (c *Client) NodesQemuAgentExecStatus(req NodesQemuAgentExecStatusRequest) (*NodesQemuAgentExecStatusResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuAgentExecStatus: %v", err)
	}

	var res *NodesQemuAgentExecStatusResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/exec-status", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuAgentExecStatus: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuAgentFileRead GET /nodes/{node}/qemu/{vmid}/agent/file-read
// Reads the given file via guest agent. Is limited to 16777216 bytes.
//
func (c *Client) NodesQemuAgentFileRead(req NodesQemuAgentFileReadRequest) (*NodesQemuAgentFileReadResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuAgentFileRead: %v", err)
	}

	var res *NodesQemuAgentFileReadResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/agent/file-read", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuAgentFileRead: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuAgentFileWrite POST /nodes/{node}/qemu/{vmid}/agent/file-write
// Writes the given file via guest agent.
//
func (c *Client) NodesQemuAgentFileWrite(req NodesQemuAgentFileWriteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuAgentFileWrite: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuAgentFileWrite: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/agent/file-write", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuAgentFileWrite: %v", err)
	}

	// output
	return nil
}

// NodesQemuRrd GET /nodes/{node}/qemu/{vmid}/rrd
// Read VM RRD statistics (returns PNG)
//
func (c *Client) NodesQemuRrd(req NodesQemuRrdRequest) (*NodesQemuRrdResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuRrd: %v", err)
	}

	var res *NodesQemuRrdResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/rrd", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuRrd: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuRrddata GET /nodes/{node}/qemu/{vmid}/rrddata
// Read VM RRD statistics
//
func (c *Client) NodesQemuRrddata(req NodesQemuRrddataRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuRrddata: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/rrddata", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuRrddata: %v", err)
	}

	// output
	return nil
}

// NodesQemuConfigVm GET /nodes/{node}/qemu/{vmid}/config
// Get the virtual machine configuration with pending configuration changes applied. Set the 'current' parameter to get the current configuration instead.
//
func (c *Client) NodesQemuConfigVm(req NodesQemuConfigVmRequest) (*NodesQemuConfigVmResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuConfigVm: %v", err)
	}

	var res *NodesQemuConfigVmResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/config", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuConfigVm: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuConfigUpdateVmAsync POST /nodes/{node}/qemu/{vmid}/config
// Set virtual machine options (asynchrounous API).
//
func (c *Client) NodesQemuConfigUpdateVmAsync(req NodesQemuConfigUpdateVmAsyncRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuConfigUpdateVmAsync: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuConfigUpdateVmAsync: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/config", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuConfigUpdateVmAsync: %v", err)
	}

	// output
	return nil
}

// NodesQemuConfigUpdateVm PUT /nodes/{node}/qemu/{vmid}/config
// Set virtual machine options (synchrounous API) - You should consider using the POST method instead for any actions involving hotplug or storage allocation.
//
func (c *Client) NodesQemuConfigUpdateVm(req NodesQemuConfigUpdateVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuConfigUpdateVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuConfigUpdateVm: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/qemu/%v/config", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuConfigUpdateVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuPendingVm GET /nodes/{node}/qemu/{vmid}/pending
// Get the virtual machine configuration with both current and pending values.
//
func (c *Client) NodesQemuPendingVm(req NodesQemuPendingVmRequest) ([]*NodesQemuPendingVmResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuPendingVm: %v", err)
	}

	var res []*NodesQemuPendingVmResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/pending", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuPendingVm: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuUnlink PUT /nodes/{node}/qemu/{vmid}/unlink
// Unlink/delete disk images.
//
func (c *Client) NodesQemuUnlink(req NodesQemuUnlinkRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuUnlink: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuUnlink: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/qemu/%v/unlink", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuUnlink: %v", err)
	}

	// output
	return nil
}

// NodesQemuVncproxy POST /nodes/{node}/qemu/{vmid}/vncproxy
// Creates a TCP VNC proxy connections.
//
func (c *Client) NodesQemuVncproxy(req NodesQemuVncproxyRequest) (*NodesQemuVncproxyResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuVncproxy: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesQemuVncproxy: %v", err)
	}

	var res *NodesQemuVncproxyResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/vncproxy", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuVncproxy: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuTermproxy POST /nodes/{node}/qemu/{vmid}/termproxy
// Creates a TCP proxy connections.
//
func (c *Client) NodesQemuTermproxy(req NodesQemuTermproxyRequest) (*NodesQemuTermproxyResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuTermproxy: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesQemuTermproxy: %v", err)
	}

	var res *NodesQemuTermproxyResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/termproxy", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuTermproxy: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuVncwebsocket GET /nodes/{node}/qemu/{vmid}/vncwebsocket
// Opens a weksocket for VNC traffic.
//
func (c *Client) NodesQemuVncwebsocket(req NodesQemuVncwebsocketRequest) (*NodesQemuVncwebsocketResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuVncwebsocket: %v", err)
	}

	var res *NodesQemuVncwebsocketResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/vncwebsocket", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuVncwebsocket: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuSpiceproxy POST /nodes/{node}/qemu/{vmid}/spiceproxy
// Returns a SPICE configuration to connect to the VM.
//
func (c *Client) NodesQemuSpiceproxy(req NodesQemuSpiceproxyRequest) (*NodesQemuSpiceproxyResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuSpiceproxy: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesQemuSpiceproxy: %v", err)
	}

	var res *NodesQemuSpiceproxyResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/spiceproxy", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuSpiceproxy: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuStatusVmcmdidx GET /nodes/{node}/qemu/{vmid}/status
// Directory index
//
func (c *Client) NodesQemuStatusVmcmdidx(req NodesQemuStatusVmcmdidxRequest) ([]*NodesQemuStatusVmcmdidxResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuStatusVmcmdidx: %v", err)
	}

	var res []*NodesQemuStatusVmcmdidxResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/status", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuStatusVmcmdidx: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuStatusCurrentVmStatus GET /nodes/{node}/qemu/{vmid}/status/current
// Get virtual machine status.
//
func (c *Client) NodesQemuStatusCurrentVmStatus(req NodesQemuStatusCurrentVmStatusRequest) (*NodesQemuStatusCurrentVmStatusResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuStatusCurrentVmStatus: %v", err)
	}

	var res *NodesQemuStatusCurrentVmStatusResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/status/current", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuStatusCurrentVmStatus: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuStatusStartVm POST /nodes/{node}/qemu/{vmid}/status/start
// Start virtual machine.
//
func (c *Client) NodesQemuStatusStartVm(req NodesQemuStatusStartVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuStatusStartVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuStatusStartVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/status/start", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuStatusStartVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuStatusStopVm POST /nodes/{node}/qemu/{vmid}/status/stop
// Stop virtual machine. The qemu process will exit immediately. Thisis akin to pulling the power plug of a running computer and may damage the VM data
//
func (c *Client) NodesQemuStatusStopVm(req NodesQemuStatusStopVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuStatusStopVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuStatusStopVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/status/stop", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuStatusStopVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuStatusResetVm POST /nodes/{node}/qemu/{vmid}/status/reset
// Reset virtual machine.
//
func (c *Client) NodesQemuStatusResetVm(req NodesQemuStatusResetVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuStatusResetVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuStatusResetVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/status/reset", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuStatusResetVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuStatusShutdownVm POST /nodes/{node}/qemu/{vmid}/status/shutdown
// Shutdown virtual machine. This is similar to pressing the power button on a physical machine.This will send an ACPI event for the guest OS, which should then proceed to a clean shutdown.
//
func (c *Client) NodesQemuStatusShutdownVm(req NodesQemuStatusShutdownVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuStatusShutdownVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuStatusShutdownVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/status/shutdown", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuStatusShutdownVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuStatusRebootVm POST /nodes/{node}/qemu/{vmid}/status/reboot
// Reboot the VM by shutting it down, and starting it again. Applies pending changes.
//
func (c *Client) NodesQemuStatusRebootVm(req NodesQemuStatusRebootVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuStatusRebootVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuStatusRebootVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/status/reboot", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuStatusRebootVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuStatusSuspendVm POST /nodes/{node}/qemu/{vmid}/status/suspend
// Suspend virtual machine.
//
func (c *Client) NodesQemuStatusSuspendVm(req NodesQemuStatusSuspendVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuStatusSuspendVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuStatusSuspendVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/status/suspend", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuStatusSuspendVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuStatusResumeVm POST /nodes/{node}/qemu/{vmid}/status/resume
// Resume virtual machine.
//
func (c *Client) NodesQemuStatusResumeVm(req NodesQemuStatusResumeVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuStatusResumeVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuStatusResumeVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/status/resume", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuStatusResumeVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuSendkeyVm PUT /nodes/{node}/qemu/{vmid}/sendkey
// Send key event to virtual machine.
//
func (c *Client) NodesQemuSendkeyVm(req NodesQemuSendkeyVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuSendkeyVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuSendkeyVm: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/qemu/%v/sendkey", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuSendkeyVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuFeatureVm GET /nodes/{node}/qemu/{vmid}/feature
// Check if feature for virtual machine is available.
//
func (c *Client) NodesQemuFeatureVm(req NodesQemuFeatureVmRequest) (*NodesQemuFeatureVmResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuFeatureVm: %v", err)
	}

	var res *NodesQemuFeatureVmResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/feature", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuFeatureVm: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuCloneVm POST /nodes/{node}/qemu/{vmid}/clone
// Create a copy of virtual machine/template.
//
func (c *Client) NodesQemuCloneVm(req NodesQemuCloneVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuCloneVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuCloneVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/clone", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuCloneVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuMoveDiskMoveVmDisk POST /nodes/{node}/qemu/{vmid}/move_disk
// Move volume to different storage or to a different VM.
//
func (c *Client) NodesQemuMoveDiskMoveVmDisk(req NodesQemuMoveDiskMoveVmDiskRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuMoveDiskMoveVmDisk: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuMoveDiskMoveVmDisk: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/move_disk", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuMoveDiskMoveVmDisk: %v", err)
	}

	// output
	return nil
}

// NodesQemuMigrateVm POST /nodes/{node}/qemu/{vmid}/migrate
// Migrate virtual machine. Creates a new migration task.
//
func (c *Client) NodesQemuMigrateVm(req NodesQemuMigrateVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuMigrateVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuMigrateVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/migrate", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuMigrateVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuMigrateVmPrecondition GET /nodes/{node}/qemu/{vmid}/migrate
// Get preconditions for migration.
//
func (c *Client) NodesQemuMigrateVmPrecondition(req NodesQemuMigrateVmPreconditionRequest) (*NodesQemuMigrateVmPreconditionResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuMigrateVmPrecondition: %v", err)
	}

	var res *NodesQemuMigrateVmPreconditionResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/migrate", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuMigrateVmPrecondition: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuMonitor POST /nodes/{node}/qemu/{vmid}/monitor
// Execute Qemu monitor commands.
//
func (c *Client) NodesQemuMonitor(req NodesQemuMonitorRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuMonitor: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuMonitor: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/monitor", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuMonitor: %v", err)
	}

	// output
	return nil
}

// NodesQemuResizeVm PUT /nodes/{node}/qemu/{vmid}/resize
// Extend volume size.
//
func (c *Client) NodesQemuResizeVm(req NodesQemuResizeVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuResizeVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuResizeVm: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/qemu/%v/resize", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuResizeVm: %v", err)
	}

	// output
	return nil
}

// NodesQemuSnapshot POST /nodes/{node}/qemu/{vmid}/snapshot
// Snapshot a VM.
//
func (c *Client) NodesQemuSnapshot(req NodesQemuSnapshotRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuSnapshot: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuSnapshot: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/snapshot", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuSnapshot: %v", err)
	}

	// output
	return nil
}

// NodesQemuSnapshotList GET /nodes/{node}/qemu/{vmid}/snapshot
// List all snapshots.
//
func (c *Client) NodesQemuSnapshotList(req NodesQemuSnapshotListRequest) ([]*NodesQemuSnapshotListResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQemuSnapshotList: %v", err)
	}

	var res []*NodesQemuSnapshotListResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/snapshot", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQemuSnapshotList: %v", err)
	}

	// output
	return res, nil
}

// NodesQemuSnapshotDelsnapshot DELETE /nodes/{node}/qemu/{vmid}/snapshot/{snapname}
// Delete a VM snapshot.
//
func (c *Client) NodesQemuSnapshotDelsnapshot(req NodesQemuSnapshotDelsnapshotRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuSnapshotDelsnapshot: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/qemu/%v/snapshot/%v", req.Node, req.Vmid, req.Snapname), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuSnapshotDelsnapshot: %v", err)
	}

	// output
	return nil
}

// NodesQemuSnapshotCmdIdx GET /nodes/{node}/qemu/{vmid}/snapshot/{snapname}
//
//
func (c *Client) NodesQemuSnapshotCmdIdx(req NodesQemuSnapshotCmdIdxRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuSnapshotCmdIdx: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/snapshot/%v", req.Node, req.Vmid, req.Snapname), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuSnapshotCmdIdx: %v", err)
	}

	// output
	return nil
}

// NodesQemuSnapshotConfigGetSnapshot GET /nodes/{node}/qemu/{vmid}/snapshot/{snapname}/config
// Get snapshot configuration
//
func (c *Client) NodesQemuSnapshotConfigGetSnapshot(req NodesQemuSnapshotConfigGetSnapshotRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuSnapshotConfigGetSnapshot: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/snapshot/%v/config", req.Node, req.Vmid, req.Snapname), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuSnapshotConfigGetSnapshot: %v", err)
	}

	// output
	return nil
}

// NodesQemuSnapshotConfigUpdateSnapshot PUT /nodes/{node}/qemu/{vmid}/snapshot/{snapname}/config
// Update snapshot metadata.
//
func (c *Client) NodesQemuSnapshotConfigUpdateSnapshot(req NodesQemuSnapshotConfigUpdateSnapshotRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuSnapshotConfigUpdateSnapshot: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuSnapshotConfigUpdateSnapshot: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/qemu/%v/snapshot/%v/config", req.Node, req.Vmid, req.Snapname), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuSnapshotConfigUpdateSnapshot: %v", err)
	}

	// output
	return nil
}

// NodesQemuSnapshotRollback POST /nodes/{node}/qemu/{vmid}/snapshot/{snapname}/rollback
// Rollback VM state to specified snapshot.
//
func (c *Client) NodesQemuSnapshotRollback(req NodesQemuSnapshotRollbackRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuSnapshotRollback: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuSnapshotRollback: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/snapshot/%v/rollback", req.Node, req.Vmid, req.Snapname), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuSnapshotRollback: %v", err)
	}

	// output
	return nil
}

// NodesQemuTemplate POST /nodes/{node}/qemu/{vmid}/template
// Create a Template.
//
func (c *Client) NodesQemuTemplate(req NodesQemuTemplateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuTemplate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesQemuTemplate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/qemu/%v/template", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesQemuTemplate: %v", err)
	}

	// output
	return nil
}

// NodesQemuCloudinitDumpCloudinitGeneratedConfig GET /nodes/{node}/qemu/{vmid}/cloudinit/dump
// Get automatically generated cloudinit config.
//
func (c *Client) NodesQemuCloudinitDumpCloudinitGeneratedConfig(req NodesQemuCloudinitDumpCloudinitGeneratedConfigRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesQemuCloudinitDumpCloudinitGeneratedConfig: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/qemu/%v/cloudinit/dump", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesQemuCloudinitDumpCloudinitGeneratedConfig: %v", err)
	}

	// output
	return nil
}

// NodesLxcVmlist GET /nodes/{node}/lxc
// LXC container index (per node).
//
func (c *Client) NodesLxcVmlist(req NodesLxcVmlistRequest) ([]*NodesLxcVmlistResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcVmlist: %v", err)
	}

	var res []*NodesLxcVmlistResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcVmlist: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcCreateVm POST /nodes/{node}/lxc
// Create or restore a container.
//
func (c *Client) NodesLxcCreateVm(req NodesLxcCreateVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcCreateVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcCreateVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcCreateVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcDestroyVm DELETE /nodes/{node}/lxc/{vmid}
// Destroy the container (also delete all uses files).
//
func (c *Client) NodesLxcDestroyVm(req NodesLxcDestroyVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcDestroyVm: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/lxc/%v", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcDestroyVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcVmdiridx GET /nodes/{node}/lxc/{vmid}
// Directory index
//
func (c *Client) NodesLxcVmdiridx(req NodesLxcVmdiridxRequest) ([]*NodesLxcVmdiridxResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcVmdiridx: %v", err)
	}

	var res []*NodesLxcVmdiridxResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcVmdiridx: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcConfigVm GET /nodes/{node}/lxc/{vmid}/config
// Get container configuration.
//
func (c *Client) NodesLxcConfigVm(req NodesLxcConfigVmRequest) (*NodesLxcConfigVmResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcConfigVm: %v", err)
	}

	var res *NodesLxcConfigVmResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/config", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcConfigVm: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcConfigUpdateVm PUT /nodes/{node}/lxc/{vmid}/config
// Set container options.
//
func (c *Client) NodesLxcConfigUpdateVm(req NodesLxcConfigUpdateVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcConfigUpdateVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcConfigUpdateVm: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/lxc/%v/config", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcConfigUpdateVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcStatusVmcmdidx GET /nodes/{node}/lxc/{vmid}/status
// Directory index
//
func (c *Client) NodesLxcStatusVmcmdidx(req NodesLxcStatusVmcmdidxRequest) ([]*NodesLxcStatusVmcmdidxResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcStatusVmcmdidx: %v", err)
	}

	var res []*NodesLxcStatusVmcmdidxResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/status", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcStatusVmcmdidx: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcStatusCurrentVmStatus GET /nodes/{node}/lxc/{vmid}/status/current
// Get virtual machine status.
//
func (c *Client) NodesLxcStatusCurrentVmStatus(req NodesLxcStatusCurrentVmStatusRequest) (*NodesLxcStatusCurrentVmStatusResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcStatusCurrentVmStatus: %v", err)
	}

	var res *NodesLxcStatusCurrentVmStatusResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/status/current", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcStatusCurrentVmStatus: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcStatusStartVm POST /nodes/{node}/lxc/{vmid}/status/start
// Start the container.
//
func (c *Client) NodesLxcStatusStartVm(req NodesLxcStatusStartVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcStatusStartVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcStatusStartVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/status/start", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcStatusStartVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcStatusStopVm POST /nodes/{node}/lxc/{vmid}/status/stop
// Stop the container. This will abruptly stop all processes running in the container.
//
func (c *Client) NodesLxcStatusStopVm(req NodesLxcStatusStopVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcStatusStopVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcStatusStopVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/status/stop", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcStatusStopVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcStatusShutdownVm POST /nodes/{node}/lxc/{vmid}/status/shutdown
// Shutdown the container. This will trigger a clean shutdown of the container, see lxc-stop(1) for details.
//
func (c *Client) NodesLxcStatusShutdownVm(req NodesLxcStatusShutdownVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcStatusShutdownVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcStatusShutdownVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/status/shutdown", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcStatusShutdownVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcStatusSuspendVm POST /nodes/{node}/lxc/{vmid}/status/suspend
// Suspend the container. This is experimental.
//
func (c *Client) NodesLxcStatusSuspendVm(req NodesLxcStatusSuspendVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcStatusSuspendVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcStatusSuspendVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/status/suspend", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcStatusSuspendVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcStatusResumeVm POST /nodes/{node}/lxc/{vmid}/status/resume
// Resume the container.
//
func (c *Client) NodesLxcStatusResumeVm(req NodesLxcStatusResumeVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcStatusResumeVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcStatusResumeVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/status/resume", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcStatusResumeVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcStatusRebootVm POST /nodes/{node}/lxc/{vmid}/status/reboot
// Reboot the container by shutting it down, and starting it again. Applies pending changes.
//
func (c *Client) NodesLxcStatusRebootVm(req NodesLxcStatusRebootVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcStatusRebootVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcStatusRebootVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/status/reboot", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcStatusRebootVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcSnapshot POST /nodes/{node}/lxc/{vmid}/snapshot
// Snapshot a container.
//
func (c *Client) NodesLxcSnapshot(req NodesLxcSnapshotRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcSnapshot: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcSnapshot: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/snapshot", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcSnapshot: %v", err)
	}

	// output
	return nil
}

// NodesLxcSnapshotList GET /nodes/{node}/lxc/{vmid}/snapshot
// List all snapshots.
//
func (c *Client) NodesLxcSnapshotList(req NodesLxcSnapshotListRequest) ([]*NodesLxcSnapshotListResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcSnapshotList: %v", err)
	}

	var res []*NodesLxcSnapshotListResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/snapshot", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcSnapshotList: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcSnapshotDelsnapshot DELETE /nodes/{node}/lxc/{vmid}/snapshot/{snapname}
// Delete a LXC snapshot.
//
func (c *Client) NodesLxcSnapshotDelsnapshot(req NodesLxcSnapshotDelsnapshotRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcSnapshotDelsnapshot: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/lxc/%v/snapshot/%v", req.Node, req.Vmid, req.Snapname), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcSnapshotDelsnapshot: %v", err)
	}

	// output
	return nil
}

// NodesLxcSnapshotCmdIdx GET /nodes/{node}/lxc/{vmid}/snapshot/{snapname}
//
//
func (c *Client) NodesLxcSnapshotCmdIdx(req NodesLxcSnapshotCmdIdxRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcSnapshotCmdIdx: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/snapshot/%v", req.Node, req.Vmid, req.Snapname), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcSnapshotCmdIdx: %v", err)
	}

	// output
	return nil
}

// NodesLxcSnapshotRollback POST /nodes/{node}/lxc/{vmid}/snapshot/{snapname}/rollback
// Rollback LXC state to specified snapshot.
//
func (c *Client) NodesLxcSnapshotRollback(req NodesLxcSnapshotRollbackRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcSnapshotRollback: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcSnapshotRollback: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/snapshot/%v/rollback", req.Node, req.Vmid, req.Snapname), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcSnapshotRollback: %v", err)
	}

	// output
	return nil
}

// NodesLxcSnapshotConfigUpdateSnapshot PUT /nodes/{node}/lxc/{vmid}/snapshot/{snapname}/config
// Update snapshot metadata.
//
func (c *Client) NodesLxcSnapshotConfigUpdateSnapshot(req NodesLxcSnapshotConfigUpdateSnapshotRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcSnapshotConfigUpdateSnapshot: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcSnapshotConfigUpdateSnapshot: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/lxc/%v/snapshot/%v/config", req.Node, req.Vmid, req.Snapname), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcSnapshotConfigUpdateSnapshot: %v", err)
	}

	// output
	return nil
}

// NodesLxcSnapshotConfigGetSnapshot GET /nodes/{node}/lxc/{vmid}/snapshot/{snapname}/config
// Get snapshot configuration
//
func (c *Client) NodesLxcSnapshotConfigGetSnapshot(req NodesLxcSnapshotConfigGetSnapshotRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcSnapshotConfigGetSnapshot: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/snapshot/%v/config", req.Node, req.Vmid, req.Snapname), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcSnapshotConfigGetSnapshot: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallRulesGet GET /nodes/{node}/lxc/{vmid}/firewall/rules
// List rules.
//
func (c *Client) NodesLxcFirewallRulesGet(req NodesLxcFirewallRulesGetRequest) ([]*NodesLxcFirewallRulesGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcFirewallRulesGet: %v", err)
	}

	var res []*NodesLxcFirewallRulesGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/rules", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcFirewallRulesGet: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcFirewallRulesCreateRule POST /nodes/{node}/lxc/{vmid}/firewall/rules
// Create new rule.
//
func (c *Client) NodesLxcFirewallRulesCreateRule(req NodesLxcFirewallRulesCreateRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallRulesCreateRule: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcFirewallRulesCreateRule: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/rules", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallRulesCreateRule: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallRulesDeleteRule DELETE /nodes/{node}/lxc/{vmid}/firewall/rules/{pos}
// Delete rule.
//
func (c *Client) NodesLxcFirewallRulesDeleteRule(req NodesLxcFirewallRulesDeleteRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallRulesDeleteRule: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/rules/%v", req.Node, req.Vmid, req.Pos), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallRulesDeleteRule: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallRulesGetRule GET /nodes/{node}/lxc/{vmid}/firewall/rules/{pos}
// Get single rule data.
//
func (c *Client) NodesLxcFirewallRulesGetRule(req NodesLxcFirewallRulesGetRuleRequest) (*NodesLxcFirewallRulesGetRuleResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcFirewallRulesGetRule: %v", err)
	}

	var res *NodesLxcFirewallRulesGetRuleResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/rules/%v", req.Node, req.Vmid, req.Pos), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcFirewallRulesGetRule: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcFirewallRulesUpdateRule PUT /nodes/{node}/lxc/{vmid}/firewall/rules/{pos}
// Modify rule data.
//
func (c *Client) NodesLxcFirewallRulesUpdateRule(req NodesLxcFirewallRulesUpdateRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallRulesUpdateRule: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcFirewallRulesUpdateRule: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/rules/%v", req.Node, req.Vmid, req.Pos), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallRulesUpdateRule: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallAliasesGet GET /nodes/{node}/lxc/{vmid}/firewall/aliases
// List aliases
//
func (c *Client) NodesLxcFirewallAliasesGet(req NodesLxcFirewallAliasesGetRequest) ([]*NodesLxcFirewallAliasesGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcFirewallAliasesGet: %v", err)
	}

	var res []*NodesLxcFirewallAliasesGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/aliases", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcFirewallAliasesGet: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcFirewallAliasesCreateAlias POST /nodes/{node}/lxc/{vmid}/firewall/aliases
// Create IP or Network Alias.
//
func (c *Client) NodesLxcFirewallAliasesCreateAlias(req NodesLxcFirewallAliasesCreateAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallAliasesCreateAlias: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcFirewallAliasesCreateAlias: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/aliases", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallAliasesCreateAlias: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallAliasesUpdateAlias PUT /nodes/{node}/lxc/{vmid}/firewall/aliases/{name}
// Update IP or Network alias.
//
func (c *Client) NodesLxcFirewallAliasesUpdateAlias(req NodesLxcFirewallAliasesUpdateAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallAliasesUpdateAlias: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcFirewallAliasesUpdateAlias: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/aliases/%v", req.Node, req.Vmid, req.Name), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallAliasesUpdateAlias: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallAliasesRemoveAlias DELETE /nodes/{node}/lxc/{vmid}/firewall/aliases/{name}
// Remove IP or Network alias.
//
func (c *Client) NodesLxcFirewallAliasesRemoveAlias(req NodesLxcFirewallAliasesRemoveAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallAliasesRemoveAlias: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/aliases/%v", req.Node, req.Vmid, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallAliasesRemoveAlias: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallAliasesReadAlias GET /nodes/{node}/lxc/{vmid}/firewall/aliases/{name}
// Read alias.
//
func (c *Client) NodesLxcFirewallAliasesReadAlias(req NodesLxcFirewallAliasesReadAliasRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallAliasesReadAlias: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/aliases/%v", req.Node, req.Vmid, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallAliasesReadAlias: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallIpsetIndex GET /nodes/{node}/lxc/{vmid}/firewall/ipset
// List IPSets
//
func (c *Client) NodesLxcFirewallIpsetIndex(req NodesLxcFirewallIpsetIndexRequest) ([]*NodesLxcFirewallIpsetIndexResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcFirewallIpsetIndex: %v", err)
	}

	var res []*NodesLxcFirewallIpsetIndexResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/ipset", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcFirewallIpsetIndex: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcFirewallIpsetCreate POST /nodes/{node}/lxc/{vmid}/firewall/ipset
// Create new IPSet
//
func (c *Client) NodesLxcFirewallIpsetCreate(req NodesLxcFirewallIpsetCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallIpsetCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcFirewallIpsetCreate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/ipset", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallIpsetCreate: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallIpsetDelete DELETE /nodes/{node}/lxc/{vmid}/firewall/ipset/{name}
// Delete IPSet
//
func (c *Client) NodesLxcFirewallIpsetDelete(req NodesLxcFirewallIpsetDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallIpsetDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/ipset/%v", req.Node, req.Vmid, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallIpsetDelete: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallIpsetGet GET /nodes/{node}/lxc/{vmid}/firewall/ipset/{name}
// List IPSet content
//
func (c *Client) NodesLxcFirewallIpsetGet(req NodesLxcFirewallIpsetGetRequest) ([]*NodesLxcFirewallIpsetGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcFirewallIpsetGet: %v", err)
	}

	var res []*NodesLxcFirewallIpsetGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/ipset/%v", req.Node, req.Vmid, req.Name), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcFirewallIpsetGet: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcFirewallIpsetCreateIp POST /nodes/{node}/lxc/{vmid}/firewall/ipset/{name}
// Add IP or Network to IPSet.
//
func (c *Client) NodesLxcFirewallIpsetCreateIp(req NodesLxcFirewallIpsetCreateIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallIpsetCreateIp: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcFirewallIpsetCreateIp: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/ipset/%v", req.Node, req.Vmid, req.Name), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallIpsetCreateIp: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallIpsetRemoveIp DELETE /nodes/{node}/lxc/{vmid}/firewall/ipset/{name}/{cidr}
// Remove IP or Network from IPSet.
//
func (c *Client) NodesLxcFirewallIpsetRemoveIp(req NodesLxcFirewallIpsetRemoveIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallIpsetRemoveIp: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/ipset/%v/%v", req.Node, req.Vmid, req.Name, req.Cidr), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallIpsetRemoveIp: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallIpsetReadIp GET /nodes/{node}/lxc/{vmid}/firewall/ipset/{name}/{cidr}
// Read IP or Network settings from IPSet.
//
func (c *Client) NodesLxcFirewallIpsetReadIp(req NodesLxcFirewallIpsetReadIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallIpsetReadIp: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/ipset/%v/%v", req.Node, req.Vmid, req.Name, req.Cidr), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallIpsetReadIp: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallIpsetUpdateIp PUT /nodes/{node}/lxc/{vmid}/firewall/ipset/{name}/{cidr}
// Update IP or Network settings
//
func (c *Client) NodesLxcFirewallIpsetUpdateIp(req NodesLxcFirewallIpsetUpdateIpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallIpsetUpdateIp: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcFirewallIpsetUpdateIp: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/ipset/%v/%v", req.Node, req.Vmid, req.Name, req.Cidr), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallIpsetUpdateIp: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallOptionsGet GET /nodes/{node}/lxc/{vmid}/firewall/options
// Get VM firewall options.
//
func (c *Client) NodesLxcFirewallOptionsGet(req NodesLxcFirewallOptionsGetRequest) (*NodesLxcFirewallOptionsGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcFirewallOptionsGet: %v", err)
	}

	var res *NodesLxcFirewallOptionsGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/options", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcFirewallOptionsGet: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcFirewallOptionsSet PUT /nodes/{node}/lxc/{vmid}/firewall/options
// Set Firewall options.
//
func (c *Client) NodesLxcFirewallOptionsSet(req NodesLxcFirewallOptionsSetRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcFirewallOptionsSet: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcFirewallOptionsSet: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/options", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcFirewallOptionsSet: %v", err)
	}

	// output
	return nil
}

// NodesLxcFirewallLog GET /nodes/{node}/lxc/{vmid}/firewall/log
// Read firewall log
//
func (c *Client) NodesLxcFirewallLog(req NodesLxcFirewallLogRequest) ([]*NodesLxcFirewallLogResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcFirewallLog: %v", err)
	}

	var res []*NodesLxcFirewallLogResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/log", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcFirewallLog: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcFirewallRefs GET /nodes/{node}/lxc/{vmid}/firewall/refs
// Lists possible IPSet/Alias reference which are allowed in source/dest properties.
//
func (c *Client) NodesLxcFirewallRefs(req NodesLxcFirewallRefsRequest) ([]*NodesLxcFirewallRefsResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcFirewallRefs: %v", err)
	}

	var res []*NodesLxcFirewallRefsResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/firewall/refs", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcFirewallRefs: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcRrd GET /nodes/{node}/lxc/{vmid}/rrd
// Read VM RRD statistics (returns PNG)
//
func (c *Client) NodesLxcRrd(req NodesLxcRrdRequest) (*NodesLxcRrdResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcRrd: %v", err)
	}

	var res *NodesLxcRrdResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/rrd", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcRrd: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcRrddata GET /nodes/{node}/lxc/{vmid}/rrddata
// Read VM RRD statistics
//
func (c *Client) NodesLxcRrddata(req NodesLxcRrddataRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcRrddata: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/rrddata", req.Node, req.Vmid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesLxcRrddata: %v", err)
	}

	// output
	return nil
}

// NodesLxcVncproxy POST /nodes/{node}/lxc/{vmid}/vncproxy
// Creates a TCP VNC proxy connections.
//
func (c *Client) NodesLxcVncproxy(req NodesLxcVncproxyRequest) (*NodesLxcVncproxyResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcVncproxy: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesLxcVncproxy: %v", err)
	}

	var res *NodesLxcVncproxyResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/vncproxy", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcVncproxy: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcTermproxy POST /nodes/{node}/lxc/{vmid}/termproxy
// Creates a TCP proxy connection.
//
func (c *Client) NodesLxcTermproxy(req NodesLxcTermproxyRequest) (*NodesLxcTermproxyResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcTermproxy: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesLxcTermproxy: %v", err)
	}

	var res *NodesLxcTermproxyResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/termproxy", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcTermproxy: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcVncwebsocket GET /nodes/{node}/lxc/{vmid}/vncwebsocket
// Opens a weksocket for VNC traffic.
//
func (c *Client) NodesLxcVncwebsocket(req NodesLxcVncwebsocketRequest) (*NodesLxcVncwebsocketResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcVncwebsocket: %v", err)
	}

	var res *NodesLxcVncwebsocketResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/vncwebsocket", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcVncwebsocket: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcSpiceproxy POST /nodes/{node}/lxc/{vmid}/spiceproxy
// Returns a SPICE configuration to connect to the CT.
//
func (c *Client) NodesLxcSpiceproxy(req NodesLxcSpiceproxyRequest) (*NodesLxcSpiceproxyResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcSpiceproxy: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesLxcSpiceproxy: %v", err)
	}

	var res *NodesLxcSpiceproxyResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/spiceproxy", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcSpiceproxy: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcMigrateVm POST /nodes/{node}/lxc/{vmid}/migrate
// Migrate the container to another node. Creates a new migration task.
//
func (c *Client) NodesLxcMigrateVm(req NodesLxcMigrateVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcMigrateVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcMigrateVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/migrate", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcMigrateVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcFeatureVm GET /nodes/{node}/lxc/{vmid}/feature
// Check if feature for virtual machine is available.
//
func (c *Client) NodesLxcFeatureVm(req NodesLxcFeatureVmRequest) (*NodesLxcFeatureVmResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcFeatureVm: %v", err)
	}

	var res *NodesLxcFeatureVmResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/feature", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcFeatureVm: %v", err)
	}

	// output
	return res, nil
}

// NodesLxcTemplate POST /nodes/{node}/lxc/{vmid}/template
// Create a Template.
//
func (c *Client) NodesLxcTemplate(req NodesLxcTemplateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcTemplate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcTemplate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/template", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcTemplate: %v", err)
	}

	// output
	return nil
}

// NodesLxcCloneVm POST /nodes/{node}/lxc/{vmid}/clone
// Create a container clone/copy
//
func (c *Client) NodesLxcCloneVm(req NodesLxcCloneVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcCloneVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcCloneVm: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/clone", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcCloneVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcResizeVm PUT /nodes/{node}/lxc/{vmid}/resize
// Resize a container mount point.
//
func (c *Client) NodesLxcResizeVm(req NodesLxcResizeVmRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcResizeVm: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcResizeVm: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/lxc/%v/resize", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcResizeVm: %v", err)
	}

	// output
	return nil
}

// NodesLxcMoveVolume POST /nodes/{node}/lxc/{vmid}/move_volume
// Move a rootfs-/mp-volume to a different storage or to a different container.
//
func (c *Client) NodesLxcMoveVolume(req NodesLxcMoveVolumeRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesLxcMoveVolume: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesLxcMoveVolume: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/lxc/%v/move_volume", req.Node, req.Vmid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesLxcMoveVolume: %v", err)
	}

	// output
	return nil
}

// NodesLxcPendingVm GET /nodes/{node}/lxc/{vmid}/pending
// Get container configuration, including pending changes.
//
func (c *Client) NodesLxcPendingVm(req NodesLxcPendingVmRequest) ([]*NodesLxcPendingVmResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesLxcPendingVm: %v", err)
	}

	var res []*NodesLxcPendingVmResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/lxc/%v/pending", req.Node, req.Vmid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesLxcPendingVm: %v", err)
	}

	// output
	return res, nil
}

// NodesCephOsdCreateosd POST /nodes/{node}/ceph/osd
// Create OSD
//
func (c *Client) NodesCephOsdCreateosd(req NodesCephOsdCreateosdRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephOsdCreateosd: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephOsdCreateosd: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/osd", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephOsdCreateosd: %v", err)
	}

	// output
	return nil
}

// NodesCephOsdDestroyosd DELETE /nodes/{node}/ceph/osd/{osdid}
// Destroy OSD
//
func (c *Client) NodesCephOsdDestroyosd(req NodesCephOsdDestroyosdRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephOsdDestroyosd: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/ceph/osd/%v", req.Node, req.Osdid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCephOsdDestroyosd: %v", err)
	}

	// output
	return nil
}

// NodesCephOsdIn POST /nodes/{node}/ceph/osd/{osdid}/in
// ceph osd in
//
func (c *Client) NodesCephOsdIn(req NodesCephOsdInRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephOsdIn: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephOsdIn: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/osd/%v/in", req.Node, req.Osdid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephOsdIn: %v", err)
	}

	// output
	return nil
}

// NodesCephOsdOut POST /nodes/{node}/ceph/osd/{osdid}/out
// ceph osd out
//
func (c *Client) NodesCephOsdOut(req NodesCephOsdOutRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephOsdOut: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephOsdOut: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/osd/%v/out", req.Node, req.Osdid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephOsdOut: %v", err)
	}

	// output
	return nil
}

// NodesCephOsdScrub POST /nodes/{node}/ceph/osd/{osdid}/scrub
// Instruct the OSD to scrub.
//
func (c *Client) NodesCephOsdScrub(req NodesCephOsdScrubRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephOsdScrub: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephOsdScrub: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/osd/%v/scrub", req.Node, req.Osdid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephOsdScrub: %v", err)
	}

	// output
	return nil
}

// NodesCephMdsDestroymds DELETE /nodes/{node}/ceph/mds/{name}
// Destroy Ceph Metadata Server
//
func (c *Client) NodesCephMdsDestroymds(req NodesCephMdsDestroymdsRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephMdsDestroymds: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/ceph/mds/%v", req.Node, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCephMdsDestroymds: %v", err)
	}

	// output
	return nil
}

// NodesCephMdsCreatemds POST /nodes/{node}/ceph/mds/{name}
// Create Ceph Metadata Server (MDS)
//
func (c *Client) NodesCephMdsCreatemds(req NodesCephMdsCreatemdsRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephMdsCreatemds: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephMdsCreatemds: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/mds/%v", req.Node, req.Name), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephMdsCreatemds: %v", err)
	}

	// output
	return nil
}

// NodesCephMgrCreatemgr POST /nodes/{node}/ceph/mgr/{id}
// Create Ceph Manager
//
func (c *Client) NodesCephMgrCreatemgr(req NodesCephMgrCreatemgrRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephMgrCreatemgr: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephMgrCreatemgr: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/mgr/%v", req.Node, req.Id), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephMgrCreatemgr: %v", err)
	}

	// output
	return nil
}

// NodesCephMgrDestroymgr DELETE /nodes/{node}/ceph/mgr/{id}
// Destroy Ceph Manager.
//
func (c *Client) NodesCephMgrDestroymgr(req NodesCephMgrDestroymgrRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephMgrDestroymgr: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/ceph/mgr/%v", req.Node, req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCephMgrDestroymgr: %v", err)
	}

	// output
	return nil
}

// NodesCephMonListmon GET /nodes/{node}/ceph/mon
// Get Ceph monitor list.
//
func (c *Client) NodesCephMonListmon(req NodesCephMonListmonRequest) ([]*NodesCephMonListmonResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesCephMonListmon: %v", err)
	}

	var res []*NodesCephMonListmonResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/ceph/mon", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesCephMonListmon: %v", err)
	}

	// output
	return res, nil
}

// NodesCephMonDestroymon DELETE /nodes/{node}/ceph/mon/{monid}
// Destroy Ceph Monitor and Manager.
//
func (c *Client) NodesCephMonDestroymon(req NodesCephMonDestroymonRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephMonDestroymon: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/ceph/mon/%v", req.Node, req.Monid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCephMonDestroymon: %v", err)
	}

	// output
	return nil
}

// NodesCephMonCreatemon POST /nodes/{node}/ceph/mon/{monid}
// Create Ceph Monitor and Manager
//
func (c *Client) NodesCephMonCreatemon(req NodesCephMonCreatemonRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephMonCreatemon: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephMonCreatemon: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/mon/%v", req.Node, req.Monid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephMonCreatemon: %v", err)
	}

	// output
	return nil
}

// NodesCephFsCreatefs POST /nodes/{node}/ceph/fs/{name}
// Create a Ceph filesystem
//
func (c *Client) NodesCephFsCreatefs(req NodesCephFsCreatefsRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephFsCreatefs: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephFsCreatefs: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/fs/%v", req.Node, req.Name), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephFsCreatefs: %v", err)
	}

	// output
	return nil
}

// NodesCephPoolsLspools GET /nodes/{node}/ceph/pools
// List all pools.
//
func (c *Client) NodesCephPoolsLspools(req NodesCephPoolsLspoolsRequest) ([]*NodesCephPoolsLspoolsResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesCephPoolsLspools: %v", err)
	}

	var res []*NodesCephPoolsLspoolsResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/ceph/pools", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesCephPoolsLspools: %v", err)
	}

	// output
	return res, nil
}

// NodesCephPoolsCreatepool POST /nodes/{node}/ceph/pools
// Create Ceph pool
//
func (c *Client) NodesCephPoolsCreatepool(req NodesCephPoolsCreatepoolRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephPoolsCreatepool: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephPoolsCreatepool: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/pools", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephPoolsCreatepool: %v", err)
	}

	// output
	return nil
}

// NodesCephPoolsDestroypool DELETE /nodes/{node}/ceph/pools/{name}
// Destroy pool
//
func (c *Client) NodesCephPoolsDestroypool(req NodesCephPoolsDestroypoolRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephPoolsDestroypool: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/ceph/pools/%v", req.Node, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCephPoolsDestroypool: %v", err)
	}

	// output
	return nil
}

// NodesCephPoolsGetpool GET /nodes/{node}/ceph/pools/{name}
// List pool settings.
//
func (c *Client) NodesCephPoolsGetpool(req NodesCephPoolsGetpoolRequest) (*NodesCephPoolsGetpoolResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesCephPoolsGetpool: %v", err)
	}

	var res *NodesCephPoolsGetpoolResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/ceph/pools/%v", req.Node, req.Name), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesCephPoolsGetpool: %v", err)
	}

	// output
	return res, nil
}

// NodesCephPoolsSetpool PUT /nodes/{node}/ceph/pools/{name}
// Change POOL settings
//
func (c *Client) NodesCephPoolsSetpool(req NodesCephPoolsSetpoolRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephPoolsSetpool: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephPoolsSetpool: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/ceph/pools/%v", req.Node, req.Name), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephPoolsSetpool: %v", err)
	}

	// output
	return nil
}

// NodesCephConfig GET /nodes/{node}/ceph/config
// Get Ceph configuration.
//
func (c *Client) NodesCephConfig(req NodesCephConfigRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephConfig: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/ceph/config", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCephConfig: %v", err)
	}

	// output
	return nil
}

// NodesCephConfigdb GET /nodes/{node}/ceph/configdb
// Get Ceph configuration database.
//
func (c *Client) NodesCephConfigdb(req NodesCephConfigdbRequest) ([]*NodesCephConfigdbResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesCephConfigdb: %v", err)
	}

	var res []*NodesCephConfigdbResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/ceph/configdb", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesCephConfigdb: %v", err)
	}

	// output
	return res, nil
}

// NodesCephInit POST /nodes/{node}/ceph/init
// Create initial ceph default configuration and setup symlinks.
//
func (c *Client) NodesCephInit(req NodesCephInitRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephInit: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephInit: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/init", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephInit: %v", err)
	}

	// output
	return nil
}

// NodesCephStop POST /nodes/{node}/ceph/stop
// Stop ceph services.
//
func (c *Client) NodesCephStop(req NodesCephStopRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephStop: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephStop: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/stop", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephStop: %v", err)
	}

	// output
	return nil
}

// NodesCephStart POST /nodes/{node}/ceph/start
// Start ceph services.
//
func (c *Client) NodesCephStart(req NodesCephStartRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephStart: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephStart: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/start", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephStart: %v", err)
	}

	// output
	return nil
}

// NodesCephRestart POST /nodes/{node}/ceph/restart
// Restart ceph services.
//
func (c *Client) NodesCephRestart(req NodesCephRestartRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephRestart: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCephRestart: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/ceph/restart", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCephRestart: %v", err)
	}

	// output
	return nil
}

// NodesCephStatus GET /nodes/{node}/ceph/status
// Get ceph status.
//
func (c *Client) NodesCephStatus(req NodesCephStatusRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephStatus: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/ceph/status", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCephStatus: %v", err)
	}

	// output
	return nil
}

// NodesCephCrush GET /nodes/{node}/ceph/crush
// Get OSD crush map
//
func (c *Client) NodesCephCrush(req NodesCephCrushRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephCrush: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/ceph/crush", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCephCrush: %v", err)
	}

	// output
	return nil
}

// NodesCephLog GET /nodes/{node}/ceph/log
// Read ceph log
//
func (c *Client) NodesCephLog(req NodesCephLogRequest) ([]*NodesCephLogResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesCephLog: %v", err)
	}

	var res []*NodesCephLogResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/ceph/log", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesCephLog: %v", err)
	}

	// output
	return res, nil
}

// NodesCephRules GET /nodes/{node}/ceph/rules
// List ceph rules.
//
func (c *Client) NodesCephRules(req NodesCephRulesRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCephRules: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/ceph/rules", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCephRules: %v", err)
	}

	// output
	return nil
}

// NodesVzdump POST /nodes/{node}/vzdump
// Create backup.
//
func (c *Client) NodesVzdump(req NodesVzdumpRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesVzdump: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesVzdump: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/vzdump", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesVzdump: %v", err)
	}

	// output
	return nil
}

// NodesVzdumpDefaults GET /nodes/{node}/vzdump/defaults
// Get the currently configured vzdump defaults.
//
func (c *Client) NodesVzdumpDefaults(req NodesVzdumpDefaultsRequest) (*NodesVzdumpDefaultsResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesVzdumpDefaults: %v", err)
	}

	var res *NodesVzdumpDefaultsResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/vzdump/defaults", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesVzdumpDefaults: %v", err)
	}

	// output
	return res, nil
}

// NodesVzdumpExtractconfig GET /nodes/{node}/vzdump/extractconfig
// Extract configuration from vzdump backup archive.
//
func (c *Client) NodesVzdumpExtractconfig(req NodesVzdumpExtractconfigRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesVzdumpExtractconfig: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/vzdump/extractconfig", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesVzdumpExtractconfig: %v", err)
	}

	// output
	return nil
}

// NodesServicesSrvcmdidx GET /nodes/{node}/services/{service}
// Directory index
//
func (c *Client) NodesServicesSrvcmdidx(req NodesServicesSrvcmdidxRequest) ([]*NodesServicesSrvcmdidxResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesServicesSrvcmdidx: %v", err)
	}

	var res []*NodesServicesSrvcmdidxResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/services/%v", req.Node, req.Service), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesServicesSrvcmdidx: %v", err)
	}

	// output
	return res, nil
}

// NodesServicesStateService GET /nodes/{node}/services/{service}/state
// Read service properties
//
func (c *Client) NodesServicesStateService(req NodesServicesStateServiceRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesServicesStateService: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/services/%v/state", req.Node, req.Service), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesServicesStateService: %v", err)
	}

	// output
	return nil
}

// NodesServicesStartService POST /nodes/{node}/services/{service}/start
// Start service.
//
func (c *Client) NodesServicesStartService(req NodesServicesStartServiceRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesServicesStartService: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesServicesStartService: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/services/%v/start", req.Node, req.Service), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesServicesStartService: %v", err)
	}

	// output
	return nil
}

// NodesServicesStopService POST /nodes/{node}/services/{service}/stop
// Stop service.
//
func (c *Client) NodesServicesStopService(req NodesServicesStopServiceRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesServicesStopService: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesServicesStopService: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/services/%v/stop", req.Node, req.Service), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesServicesStopService: %v", err)
	}

	// output
	return nil
}

// NodesServicesRestartService POST /nodes/{node}/services/{service}/restart
// Hard restart service. Use reload if you want to reduce interruptions.
//
func (c *Client) NodesServicesRestartService(req NodesServicesRestartServiceRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesServicesRestartService: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesServicesRestartService: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/services/%v/restart", req.Node, req.Service), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesServicesRestartService: %v", err)
	}

	// output
	return nil
}

// NodesServicesReloadService POST /nodes/{node}/services/{service}/reload
// Reload service. Falls back to restart if service cannot be reloaded.
//
func (c *Client) NodesServicesReloadService(req NodesServicesReloadServiceRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesServicesReloadService: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesServicesReloadService: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/services/%v/reload", req.Node, req.Service), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesServicesReloadService: %v", err)
	}

	// output
	return nil
}

// NodesSubscriptionDelete DELETE /nodes/{node}/subscription
// Delete subscription key of this node.
//
func (c *Client) NodesSubscriptionDelete(req NodesSubscriptionDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesSubscriptionDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/subscription", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesSubscriptionDelete: %v", err)
	}

	// output
	return nil
}

// NodesSubscriptionGet GET /nodes/{node}/subscription
// Read subscription info.
//
func (c *Client) NodesSubscriptionGet(req NodesSubscriptionGetRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesSubscriptionGet: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/subscription", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesSubscriptionGet: %v", err)
	}

	// output
	return nil
}

// NodesSubscriptionUpdate POST /nodes/{node}/subscription
// Update subscription info.
//
func (c *Client) NodesSubscriptionUpdate(req NodesSubscriptionUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesSubscriptionUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesSubscriptionUpdate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/subscription", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesSubscriptionUpdate: %v", err)
	}

	// output
	return nil
}

// NodesSubscriptionSet PUT /nodes/{node}/subscription
// Set subscription key.
//
func (c *Client) NodesSubscriptionSet(req NodesSubscriptionSetRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesSubscriptionSet: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesSubscriptionSet: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/subscription", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesSubscriptionSet: %v", err)
	}

	// output
	return nil
}

// NodesNetworkCreate POST /nodes/{node}/network
// Create network device configuration
//
func (c *Client) NodesNetworkCreate(req NodesNetworkCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesNetworkCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesNetworkCreate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/network", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesNetworkCreate: %v", err)
	}

	// output
	return nil
}

// NodesNetworkReloadConfig PUT /nodes/{node}/network
// Reload network configuration
//
func (c *Client) NodesNetworkReloadConfig(req NodesNetworkReloadConfigRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesNetworkReloadConfig: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesNetworkReloadConfig: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/network", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesNetworkReloadConfig: %v", err)
	}

	// output
	return nil
}

// NodesNetworkRevertChanges DELETE /nodes/{node}/network
// Revert network configuration changes.
//
func (c *Client) NodesNetworkRevertChanges(req NodesNetworkRevertChangesRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesNetworkRevertChanges: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/network", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesNetworkRevertChanges: %v", err)
	}

	// output
	return nil
}

// NodesNetworkConfig GET /nodes/{node}/network/{iface}
// Read network device configuration
//
func (c *Client) NodesNetworkConfig(req NodesNetworkConfigRequest) (*NodesNetworkConfigResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesNetworkConfig: %v", err)
	}

	var res *NodesNetworkConfigResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/network/%v", req.Node, req.Iface), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesNetworkConfig: %v", err)
	}

	// output
	return res, nil
}

// NodesNetworkUpdate PUT /nodes/{node}/network/{iface}
// Update network device configuration
//
func (c *Client) NodesNetworkUpdate(req NodesNetworkUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesNetworkUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesNetworkUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/network/%v", req.Node, req.Iface), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesNetworkUpdate: %v", err)
	}

	// output
	return nil
}

// NodesNetworkDelete DELETE /nodes/{node}/network/{iface}
// Delete network device configuration
//
func (c *Client) NodesNetworkDelete(req NodesNetworkDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesNetworkDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/network/%v", req.Node, req.Iface), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesNetworkDelete: %v", err)
	}

	// output
	return nil
}

// NodesTasksNode GET /nodes/{node}/tasks
// Read task list for one node (finished tasks).
//
func (c *Client) NodesTasksNode(req NodesTasksNodeRequest) ([]*NodesTasksNodeResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesTasksNode: %v", err)
	}

	var res []*NodesTasksNodeResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/tasks", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesTasksNode: %v", err)
	}

	// output
	return res, nil
}

// NodesTasksStopTask DELETE /nodes/{node}/tasks/{upid}
// Stop a task.
//
func (c *Client) NodesTasksStopTask(req NodesTasksStopTaskRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesTasksStopTask: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/tasks/%v", req.Node, req.Upid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesTasksStopTask: %v", err)
	}

	// output
	return nil
}

// NodesTasksUpidIndex GET /nodes/{node}/tasks/{upid}
//
//
func (c *Client) NodesTasksUpidIndex(req NodesTasksUpidIndexRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesTasksUpidIndex: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/tasks/%v", req.Node, req.Upid), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesTasksUpidIndex: %v", err)
	}

	// output
	return nil
}

// NodesTasksLogReadTask GET /nodes/{node}/tasks/{upid}/log
// Read task log.
//
func (c *Client) NodesTasksLogReadTask(req NodesTasksLogReadTaskRequest) ([]*NodesTasksLogReadTaskResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesTasksLogReadTask: %v", err)
	}

	var res []*NodesTasksLogReadTaskResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/tasks/%v/log", req.Node, req.Upid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesTasksLogReadTask: %v", err)
	}

	// output
	return res, nil
}

// NodesTasksStatusReadTask GET /nodes/{node}/tasks/{upid}/status
// Read task status.
//
func (c *Client) NodesTasksStatusReadTask(req NodesTasksStatusReadTaskRequest) (*NodesTasksStatusReadTaskResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesTasksStatusReadTask: %v", err)
	}

	var res *NodesTasksStatusReadTaskResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/tasks/%v/status", req.Node, req.Upid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesTasksStatusReadTask: %v", err)
	}

	// output
	return res, nil
}

// NodesScanNfsScan GET /nodes/{node}/scan/nfs
// Scan remote NFS server.
//
func (c *Client) NodesScanNfsScan(req NodesScanNfsScanRequest) ([]*NodesScanNfsScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesScanNfsScan: %v", err)
	}

	var res []*NodesScanNfsScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/scan/nfs", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesScanNfsScan: %v", err)
	}

	// output
	return res, nil
}

// NodesScanCifsScan GET /nodes/{node}/scan/cifs
// Scan remote CIFS server.
//
func (c *Client) NodesScanCifsScan(req NodesScanCifsScanRequest) ([]*NodesScanCifsScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesScanCifsScan: %v", err)
	}

	var res []*NodesScanCifsScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/scan/cifs", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesScanCifsScan: %v", err)
	}

	// output
	return res, nil
}

// NodesScanPbsScan GET /nodes/{node}/scan/pbs
// Scan remote Proxmox Backup Server.
//
func (c *Client) NodesScanPbsScan(req NodesScanPbsScanRequest) ([]*NodesScanPbsScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesScanPbsScan: %v", err)
	}

	var res []*NodesScanPbsScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/scan/pbs", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesScanPbsScan: %v", err)
	}

	// output
	return res, nil
}

// NodesScanGlusterfsScan GET /nodes/{node}/scan/glusterfs
// Scan remote GlusterFS server.
//
func (c *Client) NodesScanGlusterfsScan(req NodesScanGlusterfsScanRequest) ([]*NodesScanGlusterfsScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesScanGlusterfsScan: %v", err)
	}

	var res []*NodesScanGlusterfsScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/scan/glusterfs", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesScanGlusterfsScan: %v", err)
	}

	// output
	return res, nil
}

// NodesScanIscsiScan GET /nodes/{node}/scan/iscsi
// Scan remote iSCSI server.
//
func (c *Client) NodesScanIscsiScan(req NodesScanIscsiScanRequest) ([]*NodesScanIscsiScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesScanIscsiScan: %v", err)
	}

	var res []*NodesScanIscsiScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/scan/iscsi", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesScanIscsiScan: %v", err)
	}

	// output
	return res, nil
}

// NodesScanLvmScan GET /nodes/{node}/scan/lvm
// List local LVM volume groups.
//
func (c *Client) NodesScanLvmScan(req NodesScanLvmScanRequest) ([]*NodesScanLvmScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesScanLvmScan: %v", err)
	}

	var res []*NodesScanLvmScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/scan/lvm", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesScanLvmScan: %v", err)
	}

	// output
	return res, nil
}

// NodesScanLvmthinScan GET /nodes/{node}/scan/lvmthin
// List local LVM Thin Pools.
//
func (c *Client) NodesScanLvmthinScan(req NodesScanLvmthinScanRequest) ([]*NodesScanLvmthinScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesScanLvmthinScan: %v", err)
	}

	var res []*NodesScanLvmthinScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/scan/lvmthin", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesScanLvmthinScan: %v", err)
	}

	// output
	return res, nil
}

// NodesScanZfsScan GET /nodes/{node}/scan/zfs
// Scan zfs pool list on local node.
//
func (c *Client) NodesScanZfsScan(req NodesScanZfsScanRequest) ([]*NodesScanZfsScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesScanZfsScan: %v", err)
	}

	var res []*NodesScanZfsScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/scan/zfs", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesScanZfsScan: %v", err)
	}

	// output
	return res, nil
}

// NodesHardwarePciScan GET /nodes/{node}/hardware/pci
// List local PCI devices.
//
func (c *Client) NodesHardwarePciScan(req NodesHardwarePciScanRequest) ([]*NodesHardwarePciScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesHardwarePciScan: %v", err)
	}

	var res []*NodesHardwarePciScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/hardware/pci", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesHardwarePciScan: %v", err)
	}

	// output
	return res, nil
}

// NodesHardwarePciIndex GET /nodes/{node}/hardware/pci/{pciid}
// Index of available pci methods
//
func (c *Client) NodesHardwarePciIndex(req NodesHardwarePciIndexRequest) ([]*NodesHardwarePciIndexResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesHardwarePciIndex: %v", err)
	}

	var res []*NodesHardwarePciIndexResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/hardware/pci/%v", req.Node, req.Pciid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesHardwarePciIndex: %v", err)
	}

	// output
	return res, nil
}

// NodesHardwarePciMdevScan GET /nodes/{node}/hardware/pci/{pciid}/mdev
// List mediated device types for given PCI device.
//
func (c *Client) NodesHardwarePciMdevScan(req NodesHardwarePciMdevScanRequest) ([]*NodesHardwarePciMdevScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesHardwarePciMdevScan: %v", err)
	}

	var res []*NodesHardwarePciMdevScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/hardware/pci/%v/mdev", req.Node, req.Pciid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesHardwarePciMdevScan: %v", err)
	}

	// output
	return res, nil
}

// NodesHardwareUsbScan GET /nodes/{node}/hardware/usb
// List local USB devices.
//
func (c *Client) NodesHardwareUsbScan(req NodesHardwareUsbScanRequest) ([]*NodesHardwareUsbScanResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesHardwareUsbScan: %v", err)
	}

	var res []*NodesHardwareUsbScanResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/hardware/usb", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesHardwareUsbScan: %v", err)
	}

	// output
	return res, nil
}

// NodesCapabilitiesQemuCapsIndex GET /nodes/{node}/capabilities/qemu
// QEMU capabilities index.
//
func (c *Client) NodesCapabilitiesQemuCapsIndex(req NodesCapabilitiesQemuCapsIndexRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCapabilitiesQemuCapsIndex: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/capabilities/qemu", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCapabilitiesQemuCapsIndex: %v", err)
	}

	// output
	return nil
}

// NodesCapabilitiesQemuMachinesTypes GET /nodes/{node}/capabilities/qemu/machines
// Get available QEMU/KVM machine types.
//
func (c *Client) NodesCapabilitiesQemuMachinesTypes(req NodesCapabilitiesQemuMachinesTypesRequest) ([]*NodesCapabilitiesQemuMachinesTypesResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesCapabilitiesQemuMachinesTypes: %v", err)
	}

	var res []*NodesCapabilitiesQemuMachinesTypesResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/capabilities/qemu/machines", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesCapabilitiesQemuMachinesTypes: %v", err)
	}

	// output
	return res, nil
}

// NodesStorageDiridx GET /nodes/{node}/storage/{storage}
//
//
func (c *Client) NodesStorageDiridx(req NodesStorageDiridxRequest) ([]*NodesStorageDiridxResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesStorageDiridx: %v", err)
	}

	var res []*NodesStorageDiridxResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/storage/%v", req.Node, req.Storage), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesStorageDiridx: %v", err)
	}

	// output
	return res, nil
}

// NodesStoragePrunebackupsDelete DELETE /nodes/{node}/storage/{storage}/prunebackups
// Prune backups. Only those using the standard naming scheme are considered.
//
func (c *Client) NodesStoragePrunebackupsDelete(req NodesStoragePrunebackupsDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStoragePrunebackupsDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/storage/%v/prunebackups", req.Node, req.Storage), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesStoragePrunebackupsDelete: %v", err)
	}

	// output
	return nil
}

// NodesStoragePrunebackupsDryrun GET /nodes/{node}/storage/{storage}/prunebackups
// Get prune information for backups. NOTE: this is only a preview and might not be what a subsequent prune call does if backups are removed/added in the meantime.
//
func (c *Client) NodesStoragePrunebackupsDryrun(req NodesStoragePrunebackupsDryrunRequest) ([]*NodesStoragePrunebackupsDryrunResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesStoragePrunebackupsDryrun: %v", err)
	}

	var res []*NodesStoragePrunebackupsDryrunResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/storage/%v/prunebackups", req.Node, req.Storage), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesStoragePrunebackupsDryrun: %v", err)
	}

	// output
	return res, nil
}

// NodesStorageContentCreate POST /nodes/{node}/storage/{storage}/content
// Allocate disk images.
//
func (c *Client) NodesStorageContentCreate(req NodesStorageContentCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStorageContentCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesStorageContentCreate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/storage/%v/content", req.Node, req.Storage), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesStorageContentCreate: %v", err)
	}

	// output
	return nil
}

// NodesStorageContentDelete DELETE /nodes/{node}/storage/{storage}/content/{volume}
// Delete volume
//
func (c *Client) NodesStorageContentDelete(req NodesStorageContentDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStorageContentDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/storage/%v/content/%v", req.Node, req.Storage, req.Volume), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesStorageContentDelete: %v", err)
	}

	// output
	return nil
}

// NodesStorageContentInfo GET /nodes/{node}/storage/{storage}/content/{volume}
// Get volume attributes
//
func (c *Client) NodesStorageContentInfo(req NodesStorageContentInfoRequest) (*NodesStorageContentInfoResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesStorageContentInfo: %v", err)
	}

	var res *NodesStorageContentInfoResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/storage/%v/content/%v", req.Node, req.Storage, req.Volume), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesStorageContentInfo: %v", err)
	}

	// output
	return res, nil
}

// NodesStorageContentCopy POST /nodes/{node}/storage/{storage}/content/{volume}
// Copy a volume. This is experimental code - do not use.
//
func (c *Client) NodesStorageContentCopy(req NodesStorageContentCopyRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStorageContentCopy: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesStorageContentCopy: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/storage/%v/content/%v", req.Node, req.Storage, req.Volume), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesStorageContentCopy: %v", err)
	}

	// output
	return nil
}

// NodesStorageContentUpdateattributes PUT /nodes/{node}/storage/{storage}/content/{volume}
// Update volume attributes
//
func (c *Client) NodesStorageContentUpdateattributes(req NodesStorageContentUpdateattributesRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStorageContentUpdateattributes: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesStorageContentUpdateattributes: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/storage/%v/content/%v", req.Node, req.Storage, req.Volume), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesStorageContentUpdateattributes: %v", err)
	}

	// output
	return nil
}

// NodesStorageFileRestoreList GET /nodes/{node}/storage/{storage}/file-restore/list
// List files and directories for single file restore under the given path.
//
func (c *Client) NodesStorageFileRestoreList(req NodesStorageFileRestoreListRequest) ([]*NodesStorageFileRestoreListResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesStorageFileRestoreList: %v", err)
	}

	var res []*NodesStorageFileRestoreListResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/storage/%v/file-restore/list", req.Node, req.Storage), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesStorageFileRestoreList: %v", err)
	}

	// output
	return res, nil
}

// NodesStorageFileRestoreDownload GET /nodes/{node}/storage/{storage}/file-restore/download
// Extract a file or directory (as zip archive) from a PBS backup.
//
func (c *Client) NodesStorageFileRestoreDownload(req NodesStorageFileRestoreDownloadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStorageFileRestoreDownload: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/storage/%v/file-restore/download", req.Node, req.Storage), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesStorageFileRestoreDownload: %v", err)
	}

	// output
	return nil
}

// NodesStorageStatusRead GET /nodes/{node}/storage/{storage}/status
// Read storage status.
//
func (c *Client) NodesStorageStatusRead(req NodesStorageStatusReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStorageStatusRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/storage/%v/status", req.Node, req.Storage), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesStorageStatusRead: %v", err)
	}

	// output
	return nil
}

// NodesStorageRrd GET /nodes/{node}/storage/{storage}/rrd
// Read storage RRD statistics (returns PNG).
//
func (c *Client) NodesStorageRrd(req NodesStorageRrdRequest) (*NodesStorageRrdResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesStorageRrd: %v", err)
	}

	var res *NodesStorageRrdResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/storage/%v/rrd", req.Node, req.Storage), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesStorageRrd: %v", err)
	}

	// output
	return res, nil
}

// NodesStorageRrddata GET /nodes/{node}/storage/{storage}/rrddata
// Read storage RRD statistics.
//
func (c *Client) NodesStorageRrddata(req NodesStorageRrddataRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStorageRrddata: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/storage/%v/rrddata", req.Node, req.Storage), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesStorageRrddata: %v", err)
	}

	// output
	return nil
}

// NodesStorageUpload POST /nodes/{node}/storage/{storage}/upload
// Upload templates and ISO images.
//
func (c *Client) NodesStorageUpload(req NodesStorageUploadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStorageUpload: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesStorageUpload: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/storage/%v/upload", req.Node, req.Storage), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesStorageUpload: %v", err)
	}

	// output
	return nil
}

// NodesStorageDownloadUrlDownloadUrl POST /nodes/{node}/storage/{storage}/download-url
// Download templates and ISO images by using an URL.
//
func (c *Client) NodesStorageDownloadUrlDownloadUrl(req NodesStorageDownloadUrlDownloadUrlRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStorageDownloadUrlDownloadUrl: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesStorageDownloadUrlDownloadUrl: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/storage/%v/download-url", req.Node, req.Storage), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesStorageDownloadUrlDownloadUrl: %v", err)
	}

	// output
	return nil
}

// NodesDisksLvmCreate POST /nodes/{node}/disks/lvm
// Create an LVM Volume Group
//
func (c *Client) NodesDisksLvmCreate(req NodesDisksLvmCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDisksLvmCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesDisksLvmCreate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/disks/lvm", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesDisksLvmCreate: %v", err)
	}

	// output
	return nil
}

// NodesDisksLvmDelete DELETE /nodes/{node}/disks/lvm/{name}
// Remove an LVM Volume Group.
//
func (c *Client) NodesDisksLvmDelete(req NodesDisksLvmDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDisksLvmDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/disks/lvm/%v", req.Node, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesDisksLvmDelete: %v", err)
	}

	// output
	return nil
}

// NodesDisksLvmthinCreate POST /nodes/{node}/disks/lvmthin
// Create an LVM thinpool
//
func (c *Client) NodesDisksLvmthinCreate(req NodesDisksLvmthinCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDisksLvmthinCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesDisksLvmthinCreate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/disks/lvmthin", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesDisksLvmthinCreate: %v", err)
	}

	// output
	return nil
}

// NodesDisksLvmthinDelete DELETE /nodes/{node}/disks/lvmthin/{name}
// Remove an LVM thin pool.
//
func (c *Client) NodesDisksLvmthinDelete(req NodesDisksLvmthinDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDisksLvmthinDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/disks/lvmthin/%v", req.Node, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesDisksLvmthinDelete: %v", err)
	}

	// output
	return nil
}

// NodesDisksDirectoryCreate POST /nodes/{node}/disks/directory
// Create a Filesystem on an unused disk. Will be mounted under '/mnt/pve/NAME'.
//
func (c *Client) NodesDisksDirectoryCreate(req NodesDisksDirectoryCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDisksDirectoryCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesDisksDirectoryCreate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/disks/directory", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesDisksDirectoryCreate: %v", err)
	}

	// output
	return nil
}

// NodesDisksDirectoryDelete DELETE /nodes/{node}/disks/directory/{name}
// Unmounts the storage and removes the mount unit.
//
func (c *Client) NodesDisksDirectoryDelete(req NodesDisksDirectoryDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDisksDirectoryDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/disks/directory/%v", req.Node, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesDisksDirectoryDelete: %v", err)
	}

	// output
	return nil
}

// NodesDisksZfsCreate POST /nodes/{node}/disks/zfs
// Create a ZFS pool.
//
func (c *Client) NodesDisksZfsCreate(req NodesDisksZfsCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDisksZfsCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesDisksZfsCreate: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/disks/zfs", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesDisksZfsCreate: %v", err)
	}

	// output
	return nil
}

// NodesDisksZfsDelete DELETE /nodes/{node}/disks/zfs/{name}
// Destroy a ZFS pool.
//
func (c *Client) NodesDisksZfsDelete(req NodesDisksZfsDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDisksZfsDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/disks/zfs/%v", req.Node, req.Name), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesDisksZfsDelete: %v", err)
	}

	// output
	return nil
}

// NodesDisksZfsDetail GET /nodes/{node}/disks/zfs/{name}
// Get details about a zpool.
//
func (c *Client) NodesDisksZfsDetail(req NodesDisksZfsDetailRequest) (*NodesDisksZfsDetailResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesDisksZfsDetail: %v", err)
	}

	var res *NodesDisksZfsDetailResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/disks/zfs/%v", req.Node, req.Name), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesDisksZfsDetail: %v", err)
	}

	// output
	return res, nil
}

// NodesDisksList GET /nodes/{node}/disks/list
// List local disks.
//
func (c *Client) NodesDisksList(req NodesDisksListRequest) ([]*NodesDisksListResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesDisksList: %v", err)
	}

	var res []*NodesDisksListResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/disks/list", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesDisksList: %v", err)
	}

	// output
	return res, nil
}

// NodesDisksSmart GET /nodes/{node}/disks/smart
// Get SMART Health of a disk.
//
func (c *Client) NodesDisksSmart(req NodesDisksSmartRequest) (*NodesDisksSmartResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesDisksSmart: %v", err)
	}

	var res *NodesDisksSmartResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/disks/smart", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesDisksSmart: %v", err)
	}

	// output
	return res, nil
}

// NodesDisksInitgpt POST /nodes/{node}/disks/initgpt
// Initialize Disk with GPT
//
func (c *Client) NodesDisksInitgpt(req NodesDisksInitgptRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDisksInitgpt: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesDisksInitgpt: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/disks/initgpt", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesDisksInitgpt: %v", err)
	}

	// output
	return nil
}

// NodesDisksWipediskWipeDisk PUT /nodes/{node}/disks/wipedisk
// Wipe a disk or partition.
//
func (c *Client) NodesDisksWipediskWipeDisk(req NodesDisksWipediskWipeDiskRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDisksWipediskWipeDisk: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesDisksWipediskWipeDisk: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/disks/wipedisk", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesDisksWipediskWipeDisk: %v", err)
	}

	// output
	return nil
}

// NodesAptUpdateDatabase POST /nodes/{node}/apt/update
// This is used to resynchronize the package index files from their sources (apt-get update).
//
func (c *Client) NodesAptUpdateDatabase(req NodesAptUpdateDatabaseRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesAptUpdateDatabase: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesAptUpdateDatabase: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/apt/update", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesAptUpdateDatabase: %v", err)
	}

	// output
	return nil
}

// NodesAptUpdateLists GET /nodes/{node}/apt/update
// List available updates.
//
func (c *Client) NodesAptUpdateLists(req NodesAptUpdateListsRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesAptUpdateLists: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/apt/update", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesAptUpdateLists: %v", err)
	}

	// output
	return nil
}

// NodesAptChangelog GET /nodes/{node}/apt/changelog
// Get package changelogs.
//
func (c *Client) NodesAptChangelog(req NodesAptChangelogRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesAptChangelog: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/apt/changelog", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesAptChangelog: %v", err)
	}

	// output
	return nil
}

// NodesAptRepositoriesAddRepository PUT /nodes/{node}/apt/repositories
// Add a standard repository to the configuration
//
func (c *Client) NodesAptRepositoriesAddRepository(req NodesAptRepositoriesAddRepositoryRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesAptRepositoriesAddRepository: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesAptRepositoriesAddRepository: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/apt/repositories", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesAptRepositoriesAddRepository: %v", err)
	}

	// output
	return nil
}

// NodesAptRepositories GET /nodes/{node}/apt/repositories
// Get APT repository information.
//
func (c *Client) NodesAptRepositories(req NodesAptRepositoriesRequest) (*NodesAptRepositoriesResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesAptRepositories: %v", err)
	}

	var res *NodesAptRepositoriesResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/apt/repositories", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesAptRepositories: %v", err)
	}

	// output
	return res, nil
}

// NodesAptRepositoriesChangeRepository POST /nodes/{node}/apt/repositories
// Change the properties of a repository. Currently only allows enabling/disabling.
//
func (c *Client) NodesAptRepositoriesChangeRepository(req NodesAptRepositoriesChangeRepositoryRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesAptRepositoriesChangeRepository: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesAptRepositoriesChangeRepository: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/apt/repositories", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesAptRepositoriesChangeRepository: %v", err)
	}

	// output
	return nil
}

// NodesAptVersions GET /nodes/{node}/apt/versions
// Get package information for important Proxmox packages.
//
func (c *Client) NodesAptVersions(req NodesAptVersionsRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesAptVersions: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/apt/versions", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesAptVersions: %v", err)
	}

	// output
	return nil
}

// NodesFirewallRulesGet GET /nodes/{node}/firewall/rules
// List rules.
//
func (c *Client) NodesFirewallRulesGet(req NodesFirewallRulesGetRequest) ([]*NodesFirewallRulesGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesFirewallRulesGet: %v", err)
	}

	var res []*NodesFirewallRulesGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/firewall/rules", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesFirewallRulesGet: %v", err)
	}

	// output
	return res, nil
}

// NodesFirewallRulesCreateRule POST /nodes/{node}/firewall/rules
// Create new rule.
//
func (c *Client) NodesFirewallRulesCreateRule(req NodesFirewallRulesCreateRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesFirewallRulesCreateRule: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesFirewallRulesCreateRule: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/firewall/rules", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesFirewallRulesCreateRule: %v", err)
	}

	// output
	return nil
}

// NodesFirewallRulesDeleteRule DELETE /nodes/{node}/firewall/rules/{pos}
// Delete rule.
//
func (c *Client) NodesFirewallRulesDeleteRule(req NodesFirewallRulesDeleteRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesFirewallRulesDeleteRule: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/firewall/rules/%v", req.Node, req.Pos), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesFirewallRulesDeleteRule: %v", err)
	}

	// output
	return nil
}

// NodesFirewallRulesGetRule GET /nodes/{node}/firewall/rules/{pos}
// Get single rule data.
//
func (c *Client) NodesFirewallRulesGetRule(req NodesFirewallRulesGetRuleRequest) (*NodesFirewallRulesGetRuleResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesFirewallRulesGetRule: %v", err)
	}

	var res *NodesFirewallRulesGetRuleResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/firewall/rules/%v", req.Node, req.Pos), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesFirewallRulesGetRule: %v", err)
	}

	// output
	return res, nil
}

// NodesFirewallRulesUpdateRule PUT /nodes/{node}/firewall/rules/{pos}
// Modify rule data.
//
func (c *Client) NodesFirewallRulesUpdateRule(req NodesFirewallRulesUpdateRuleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesFirewallRulesUpdateRule: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesFirewallRulesUpdateRule: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/firewall/rules/%v", req.Node, req.Pos), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesFirewallRulesUpdateRule: %v", err)
	}

	// output
	return nil
}

// NodesFirewallOptionsGet GET /nodes/{node}/firewall/options
// Get host firewall options.
//
func (c *Client) NodesFirewallOptionsGet(req NodesFirewallOptionsGetRequest) (*NodesFirewallOptionsGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesFirewallOptionsGet: %v", err)
	}

	var res *NodesFirewallOptionsGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/firewall/options", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesFirewallOptionsGet: %v", err)
	}

	// output
	return res, nil
}

// NodesFirewallOptionsSet PUT /nodes/{node}/firewall/options
// Set Firewall options.
//
func (c *Client) NodesFirewallOptionsSet(req NodesFirewallOptionsSetRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesFirewallOptionsSet: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesFirewallOptionsSet: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/firewall/options", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesFirewallOptionsSet: %v", err)
	}

	// output
	return nil
}

// NodesFirewallLog GET /nodes/{node}/firewall/log
// Read firewall log
//
func (c *Client) NodesFirewallLog(req NodesFirewallLogRequest) ([]*NodesFirewallLogResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesFirewallLog: %v", err)
	}

	var res []*NodesFirewallLogResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/firewall/log", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesFirewallLog: %v", err)
	}

	// output
	return res, nil
}

// NodesReplicationStatus GET /nodes/{node}/replication
// List status of all replication jobs on this node.
//
func (c *Client) NodesReplicationStatus(req NodesReplicationStatusRequest) ([]*NodesReplicationStatusResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesReplicationStatus: %v", err)
	}

	var res []*NodesReplicationStatusResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/replication", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesReplicationStatus: %v", err)
	}

	// output
	return res, nil
}

// NodesReplicationStatusJob GET /nodes/{node}/replication/{id}/status
// Get replication job status.
//
func (c *Client) NodesReplicationStatusJob(req NodesReplicationStatusJobRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesReplicationStatusJob: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/replication/%v/status", req.Node, req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesReplicationStatusJob: %v", err)
	}

	// output
	return nil
}

// NodesReplicationLogReadJob GET /nodes/{node}/replication/{id}/log
// Read replication job log.
//
func (c *Client) NodesReplicationLogReadJob(req NodesReplicationLogReadJobRequest) ([]*NodesReplicationLogReadJobResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesReplicationLogReadJob: %v", err)
	}

	var res []*NodesReplicationLogReadJobResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/replication/%v/log", req.Node, req.Id), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesReplicationLogReadJob: %v", err)
	}

	// output
	return res, nil
}

// NodesReplicationScheduleNow POST /nodes/{node}/replication/{id}/schedule_now
// Schedule replication job to start as soon as possible.
//
func (c *Client) NodesReplicationScheduleNow(req NodesReplicationScheduleNowRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesReplicationScheduleNow: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesReplicationScheduleNow: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/replication/%v/schedule_now", req.Node, req.Id), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesReplicationScheduleNow: %v", err)
	}

	// output
	return nil
}

// NodesCertificatesAcmeCertificateNew POST /nodes/{node}/certificates/acme/certificate
// Order a new certificate from ACME-compatible CA.
//
func (c *Client) NodesCertificatesAcmeCertificateNew(req NodesCertificatesAcmeCertificateNewRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCertificatesAcmeCertificateNew: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCertificatesAcmeCertificateNew: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/certificates/acme/certificate", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCertificatesAcmeCertificateNew: %v", err)
	}

	// output
	return nil
}

// NodesCertificatesAcmeCertificateRenew PUT /nodes/{node}/certificates/acme/certificate
// Renew existing certificate from CA.
//
func (c *Client) NodesCertificatesAcmeCertificateRenew(req NodesCertificatesAcmeCertificateRenewRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCertificatesAcmeCertificateRenew: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesCertificatesAcmeCertificateRenew: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/certificates/acme/certificate", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesCertificatesAcmeCertificateRenew: %v", err)
	}

	// output
	return nil
}

// NodesCertificatesAcmeCertificateRevoke DELETE /nodes/{node}/certificates/acme/certificate
// Revoke existing certificate from CA.
//
func (c *Client) NodesCertificatesAcmeCertificateRevoke(req NodesCertificatesAcmeCertificateRevokeRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCertificatesAcmeCertificateRevoke: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/certificates/acme/certificate", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCertificatesAcmeCertificateRevoke: %v", err)
	}

	// output
	return nil
}

// NodesCertificatesInfo GET /nodes/{node}/certificates/info
// Get information about node's certificates.
//
func (c *Client) NodesCertificatesInfo(req NodesCertificatesInfoRequest) ([]*NodesCertificatesInfoResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesCertificatesInfo: %v", err)
	}

	var res []*NodesCertificatesInfoResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/certificates/info", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesCertificatesInfo: %v", err)
	}

	// output
	return res, nil
}

// NodesCertificatesCustomRemoveCert DELETE /nodes/{node}/certificates/custom
// DELETE custom certificate chain and key.
//
func (c *Client) NodesCertificatesCustomRemoveCert(req NodesCertificatesCustomRemoveCertRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesCertificatesCustomRemoveCert: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/nodes/%v/certificates/custom", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesCertificatesCustomRemoveCert: %v", err)
	}

	// output
	return nil
}

// NodesCertificatesCustomUploadCert POST /nodes/{node}/certificates/custom
// Upload or update custom certificate chain and key.
//
func (c *Client) NodesCertificatesCustomUploadCert(req NodesCertificatesCustomUploadCertRequest) (*NodesCertificatesCustomUploadCertResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesCertificatesCustomUploadCert: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesCertificatesCustomUploadCert: %v", err)
	}

	var res *NodesCertificatesCustomUploadCertResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/certificates/custom", req.Node), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesCertificatesCustomUploadCert: %v", err)
	}

	// output
	return res, nil
}

// NodesConfigGet GET /nodes/{node}/config
// Get node configuration options.
//
func (c *Client) NodesConfigGet(req NodesConfigGetRequest) (*NodesConfigGetResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesConfigGet: %v", err)
	}

	var res *NodesConfigGetResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/config", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesConfigGet: %v", err)
	}

	// output
	return res, nil
}

// NodesConfigSetOptions PUT /nodes/{node}/config
// Set node configuration options.
//
func (c *Client) NodesConfigSetOptions(req NodesConfigSetOptionsRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesConfigSetOptions: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesConfigSetOptions: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/config", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesConfigSetOptions: %v", err)
	}

	// output
	return nil
}

// NodesSdnIndex GET /nodes/{node}/sdn
// SDN index.
//
func (c *Client) NodesSdnIndex(req NodesSdnIndexRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesSdnIndex: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/sdn", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesSdnIndex: %v", err)
	}

	// output
	return nil
}

// NodesSdnZonesDiridx GET /nodes/{node}/sdn/zones/{zone}
//
//
func (c *Client) NodesSdnZonesDiridx(req NodesSdnZonesDiridxRequest) ([]*NodesSdnZonesDiridxResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesSdnZonesDiridx: %v", err)
	}

	var res []*NodesSdnZonesDiridxResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/sdn/zones/%v", req.Node, req.Zone), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesSdnZonesDiridx: %v", err)
	}

	// output
	return res, nil
}

// NodesVersion GET /nodes/{node}/version
// API version details
//
func (c *Client) NodesVersion(req NodesVersionRequest) (*NodesVersionResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesVersion: %v", err)
	}

	var res *NodesVersionResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/version", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesVersion: %v", err)
	}

	// output
	return res, nil
}

// NodesStatus GET /nodes/{node}/status
// Read node status
//
func (c *Client) NodesStatus(req NodesStatusRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStatus: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/status", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesStatus: %v", err)
	}

	// output
	return nil
}

// NodesStatusNodeCmd POST /nodes/{node}/status
// Reboot or shutdown a node.
//
func (c *Client) NodesStatusNodeCmd(req NodesStatusNodeCmdRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStatusNodeCmd: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesStatusNodeCmd: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/status", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesStatusNodeCmd: %v", err)
	}

	// output
	return nil
}

// NodesNetstat GET /nodes/{node}/netstat
// Read tap/vm network device interface counters
//
func (c *Client) NodesNetstat(req NodesNetstatRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesNetstat: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/netstat", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesNetstat: %v", err)
	}

	// output
	return nil
}

// NodesExecute POST /nodes/{node}/execute
// Execute multiple commands in order.
//
func (c *Client) NodesExecute(req NodesExecuteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesExecute: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesExecute: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/execute", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesExecute: %v", err)
	}

	// output
	return nil
}

// NodesWakeonlan POST /nodes/{node}/wakeonlan
// Try to wake a node via 'wake on LAN' network packet.
//
func (c *Client) NodesWakeonlan(req NodesWakeonlanRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesWakeonlan: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesWakeonlan: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/wakeonlan", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesWakeonlan: %v", err)
	}

	// output
	return nil
}

// NodesRrd GET /nodes/{node}/rrd
// Read node RRD statistics (returns PNG)
//
func (c *Client) NodesRrd(req NodesRrdRequest) (*NodesRrdResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesRrd: %v", err)
	}

	var res *NodesRrdResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/rrd", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesRrd: %v", err)
	}

	// output
	return res, nil
}

// NodesRrddata GET /nodes/{node}/rrddata
// Read node RRD statistics
//
func (c *Client) NodesRrddata(req NodesRrddataRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesRrddata: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/rrddata", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesRrddata: %v", err)
	}

	// output
	return nil
}

// NodesSyslog GET /nodes/{node}/syslog
// Read system log
//
func (c *Client) NodesSyslog(req NodesSyslogRequest) ([]*NodesSyslogResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesSyslog: %v", err)
	}

	var res []*NodesSyslogResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/syslog", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesSyslog: %v", err)
	}

	// output
	return res, nil
}

// NodesJournal GET /nodes/{node}/journal
// Read Journal
//
func (c *Client) NodesJournal(req NodesJournalRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesJournal: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/journal", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesJournal: %v", err)
	}

	// output
	return nil
}

// NodesVncshell POST /nodes/{node}/vncshell
// Creates a VNC Shell proxy.
//
func (c *Client) NodesVncshell(req NodesVncshellRequest) (*NodesVncshellResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesVncshell: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesVncshell: %v", err)
	}

	var res *NodesVncshellResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/vncshell", req.Node), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesVncshell: %v", err)
	}

	// output
	return res, nil
}

// NodesTermproxy POST /nodes/{node}/termproxy
// Creates a VNC Shell proxy.
//
func (c *Client) NodesTermproxy(req NodesTermproxyRequest) (*NodesTermproxyResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesTermproxy: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesTermproxy: %v", err)
	}

	var res *NodesTermproxyResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/termproxy", req.Node), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesTermproxy: %v", err)
	}

	// output
	return res, nil
}

// NodesVncwebsocket GET /nodes/{node}/vncwebsocket
// Opens a websocket for VNC traffic.
//
func (c *Client) NodesVncwebsocket(req NodesVncwebsocketRequest) (*NodesVncwebsocketResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesVncwebsocket: %v", err)
	}

	var res *NodesVncwebsocketResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/vncwebsocket", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesVncwebsocket: %v", err)
	}

	// output
	return res, nil
}

// NodesSpiceshell POST /nodes/{node}/spiceshell
// Creates a SPICE shell.
//
func (c *Client) NodesSpiceshell(req NodesSpiceshellRequest) (*NodesSpiceshellResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesSpiceshell: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("NodesSpiceshell: %v", err)
	}

	var res *NodesSpiceshellResponse
	err = c.Request("POST", fmt.Sprintf("/nodes/%v/spiceshell", req.Node), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("NodesSpiceshell: %v", err)
	}

	// output
	return res, nil
}

// NodesDns GET /nodes/{node}/dns
// Read DNS settings.
//
func (c *Client) NodesDns(req NodesDnsRequest) (*NodesDnsResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesDns: %v", err)
	}

	var res *NodesDnsResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/dns", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesDns: %v", err)
	}

	// output
	return res, nil
}

// NodesDnsUpdate PUT /nodes/{node}/dns
// Write DNS settings.
//
func (c *Client) NodesDnsUpdate(req NodesDnsUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesDnsUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesDnsUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/dns", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesDnsUpdate: %v", err)
	}

	// output
	return nil
}

// NodesTime GET /nodes/{node}/time
// Read server time and time zone settings.
//
func (c *Client) NodesTime(req NodesTimeRequest) (*NodesTimeResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesTime: %v", err)
	}

	var res *NodesTimeResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/time", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesTime: %v", err)
	}

	// output
	return res, nil
}

// NodesTimeSetzone PUT /nodes/{node}/time
// Set time zone.
//
func (c *Client) NodesTimeSetzone(req NodesTimeSetzoneRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesTimeSetzone: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesTimeSetzone: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/nodes/%v/time", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesTimeSetzone: %v", err)
	}

	// output
	return nil
}

// NodesAplinfo GET /nodes/{node}/aplinfo
// Get list of appliances.
//
func (c *Client) NodesAplinfo(req NodesAplinfoRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesAplinfo: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/aplinfo", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesAplinfo: %v", err)
	}

	// output
	return nil
}

// NodesAplinfoAplDownload POST /nodes/{node}/aplinfo
// Download appliance templates.
//
func (c *Client) NodesAplinfoAplDownload(req NodesAplinfoAplDownloadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesAplinfoAplDownload: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesAplinfoAplDownload: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/aplinfo", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesAplinfoAplDownload: %v", err)
	}

	// output
	return nil
}

// NodesQueryUrlMetadataQueryUrlMetadata GET /nodes/{node}/query-url-metadata
// Query metadata of an URL: file size, file name and mime type.
//
func (c *Client) NodesQueryUrlMetadataQueryUrlMetadata(req NodesQueryUrlMetadataQueryUrlMetadataRequest) (*NodesQueryUrlMetadataQueryUrlMetadataResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesQueryUrlMetadataQueryUrlMetadata: %v", err)
	}

	var res *NodesQueryUrlMetadataQueryUrlMetadataResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/query-url-metadata", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesQueryUrlMetadataQueryUrlMetadata: %v", err)
	}

	// output
	return res, nil
}

// NodesReport GET /nodes/{node}/report
// Gather various systems information about a node
//
func (c *Client) NodesReport(req NodesReportRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesReport: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/nodes/%v/report", req.Node), nil, nil)

	if err != nil {
		return fmt.Errorf("NodesReport: %v", err)
	}

	// output
	return nil
}

// NodesStartall POST /nodes/{node}/startall
// Start all VMs and containers located on this node (by default only those with onboot=1).
//
func (c *Client) NodesStartall(req NodesStartallRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStartall: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesStartall: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/startall", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesStartall: %v", err)
	}

	// output
	return nil
}

// NodesStopall POST /nodes/{node}/stopall
// Stop all VMs and Containers.
//
func (c *Client) NodesStopall(req NodesStopallRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesStopall: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesStopall: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/stopall", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesStopall: %v", err)
	}

	// output
	return nil
}

// NodesMigrateall POST /nodes/{node}/migrateall
// Migrate all VMs and Containers.
//
func (c *Client) NodesMigrateall(req NodesMigrateallRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesMigrateall: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesMigrateall: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/migrateall", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesMigrateall: %v", err)
	}

	// output
	return nil
}

// NodesHostsGetEtc GET /nodes/{node}/hosts
// Get the content of /etc/hosts.
//
func (c *Client) NodesHostsGetEtc(req NodesHostsGetEtcRequest) (*NodesHostsGetEtcResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("NodesHostsGetEtc: %v", err)
	}

	var res *NodesHostsGetEtcResponse
	err = c.Request("GET", fmt.Sprintf("/nodes/%v/hosts", req.Node), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("NodesHostsGetEtc: %v", err)
	}

	// output
	return res, nil
}

// NodesHostsWriteEtc POST /nodes/{node}/hosts
// Write /etc/hosts.
//
func (c *Client) NodesHostsWriteEtc(req NodesHostsWriteEtcRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("NodesHostsWriteEtc: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("NodesHostsWriteEtc: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/nodes/%v/hosts", req.Node), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("NodesHostsWriteEtc: %v", err)
	}

	// output
	return nil
}

// StorageCreate POST /storage
// Create a new storage.
//
func (c *Client) StorageCreate(req StorageCreateRequest) (*StorageCreateResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("StorageCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("StorageCreate: %v", err)
	}

	var res *StorageCreateResponse
	err = c.Request("POST", "/storage", strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("StorageCreate: %v", err)
	}

	// output
	return res, nil
}

// StorageDelete DELETE /storage/{storage}
// Delete storage configuration.
//
func (c *Client) StorageDelete(req StorageDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("StorageDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/storage/%v", req.Storage), nil, nil)

	if err != nil {
		return fmt.Errorf("StorageDelete: %v", err)
	}

	// output
	return nil
}

// StorageRead GET /storage/{storage}
// Read storage configuration.
//
func (c *Client) StorageRead(req StorageReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("StorageRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/storage/%v", req.Storage), nil, nil)

	if err != nil {
		return fmt.Errorf("StorageRead: %v", err)
	}

	// output
	return nil
}

// StorageUpdate PUT /storage/{storage}
// Update storage configuration.
//
func (c *Client) StorageUpdate(req StorageUpdateRequest) (*StorageUpdateResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("StorageUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("StorageUpdate: %v", err)
	}

	var res *StorageUpdateResponse
	err = c.Request("PUT", fmt.Sprintf("/storage/%v", req.Storage), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("StorageUpdate: %v", err)
	}

	// output
	return res, nil
}

// AccessUsersCreateUser POST /access/users
// Create new user.
//
func (c *Client) AccessUsersCreateUser(req AccessUsersCreateUserRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessUsersCreateUser: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessUsersCreateUser: %v", err)
	}

	err = c.Request("POST", "/access/users", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessUsersCreateUser: %v", err)
	}

	// output
	return nil
}

// AccessUsersDeleteUser DELETE /access/users/{userid}
// Delete user.
//
func (c *Client) AccessUsersDeleteUser(req AccessUsersDeleteUserRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessUsersDeleteUser: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/access/users/%v", req.Userid), nil, nil)

	if err != nil {
		return fmt.Errorf("AccessUsersDeleteUser: %v", err)
	}

	// output
	return nil
}

// AccessUsersReadUser GET /access/users/{userid}
// Get user configuration.
//
func (c *Client) AccessUsersReadUser(req AccessUsersReadUserRequest) (*AccessUsersReadUserResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessUsersReadUser: %v", err)
	}

	var res *AccessUsersReadUserResponse
	err = c.Request("GET", fmt.Sprintf("/access/users/%v", req.Userid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("AccessUsersReadUser: %v", err)
	}

	// output
	return res, nil
}

// AccessUsersUpdateUser PUT /access/users/{userid}
// Update user configuration.
//
func (c *Client) AccessUsersUpdateUser(req AccessUsersUpdateUserRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessUsersUpdateUser: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessUsersUpdateUser: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/access/users/%v", req.Userid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessUsersUpdateUser: %v", err)
	}

	// output
	return nil
}

// AccessUsersTfaReadUserType GET /access/users/{userid}/tfa
// Get user TFA types (Personal and Realm).
//
func (c *Client) AccessUsersTfaReadUserType(req AccessUsersTfaReadUserTypeRequest) (*AccessUsersTfaReadUserTypeResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessUsersTfaReadUserType: %v", err)
	}

	var res *AccessUsersTfaReadUserTypeResponse
	err = c.Request("GET", fmt.Sprintf("/access/users/%v/tfa", req.Userid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("AccessUsersTfaReadUserType: %v", err)
	}

	// output
	return res, nil
}

// AccessUsersTokenIndex GET /access/users/{userid}/token
// Get user API tokens.
//
func (c *Client) AccessUsersTokenIndex(req AccessUsersTokenIndexRequest) ([]*AccessUsersTokenIndexResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessUsersTokenIndex: %v", err)
	}

	var res []*AccessUsersTokenIndexResponse
	err = c.Request("GET", fmt.Sprintf("/access/users/%v/token", req.Userid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("AccessUsersTokenIndex: %v", err)
	}

	// output
	return res, nil
}

// AccessUsersTokenRemove DELETE /access/users/{userid}/token/{tokenid}
// Remove API token for a specific user.
//
func (c *Client) AccessUsersTokenRemove(req AccessUsersTokenRemoveRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessUsersTokenRemove: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/access/users/%v/token/%v", req.Userid, req.Tokenid), nil, nil)

	if err != nil {
		return fmt.Errorf("AccessUsersTokenRemove: %v", err)
	}

	// output
	return nil
}

// AccessUsersTokenRead GET /access/users/{userid}/token/{tokenid}
// Get specific API token information.
//
func (c *Client) AccessUsersTokenRead(req AccessUsersTokenReadRequest) (*AccessUsersTokenReadResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessUsersTokenRead: %v", err)
	}

	var res *AccessUsersTokenReadResponse
	err = c.Request("GET", fmt.Sprintf("/access/users/%v/token/%v", req.Userid, req.Tokenid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("AccessUsersTokenRead: %v", err)
	}

	// output
	return res, nil
}

// AccessUsersTokenGenerate POST /access/users/{userid}/token/{tokenid}
// Generate a new API token for a specific user. NOTE: returns API token value, which needs to be stored as it cannot be retrieved afterwards!
//
func (c *Client) AccessUsersTokenGenerate(req AccessUsersTokenGenerateRequest) (*AccessUsersTokenGenerateResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessUsersTokenGenerate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("AccessUsersTokenGenerate: %v", err)
	}

	var res *AccessUsersTokenGenerateResponse
	err = c.Request("POST", fmt.Sprintf("/access/users/%v/token/%v", req.Userid, req.Tokenid), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("AccessUsersTokenGenerate: %v", err)
	}

	// output
	return res, nil
}

// AccessUsersTokenUpdateInfo PUT /access/users/{userid}/token/{tokenid}
// Update API token for a specific user.
//
func (c *Client) AccessUsersTokenUpdateInfo(req AccessUsersTokenUpdateInfoRequest) (*AccessUsersTokenUpdateInfoResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessUsersTokenUpdateInfo: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("AccessUsersTokenUpdateInfo: %v", err)
	}

	var res *AccessUsersTokenUpdateInfoResponse
	err = c.Request("PUT", fmt.Sprintf("/access/users/%v/token/%v", req.Userid, req.Tokenid), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("AccessUsersTokenUpdateInfo: %v", err)
	}

	// output
	return res, nil
}

// AccessGroupsCreateGroup POST /access/groups
// Create new group.
//
func (c *Client) AccessGroupsCreateGroup(req AccessGroupsCreateGroupRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessGroupsCreateGroup: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessGroupsCreateGroup: %v", err)
	}

	err = c.Request("POST", "/access/groups", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessGroupsCreateGroup: %v", err)
	}

	// output
	return nil
}

// AccessGroupsDeleteGroup DELETE /access/groups/{groupid}
// Delete group.
//
func (c *Client) AccessGroupsDeleteGroup(req AccessGroupsDeleteGroupRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessGroupsDeleteGroup: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/access/groups/%v", req.Groupid), nil, nil)

	if err != nil {
		return fmt.Errorf("AccessGroupsDeleteGroup: %v", err)
	}

	// output
	return nil
}

// AccessGroupsReadGroup GET /access/groups/{groupid}
// Get group configuration.
//
func (c *Client) AccessGroupsReadGroup(req AccessGroupsReadGroupRequest) (*AccessGroupsReadGroupResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessGroupsReadGroup: %v", err)
	}

	var res *AccessGroupsReadGroupResponse
	err = c.Request("GET", fmt.Sprintf("/access/groups/%v", req.Groupid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("AccessGroupsReadGroup: %v", err)
	}

	// output
	return res, nil
}

// AccessGroupsUpdateGroup PUT /access/groups/{groupid}
// Update group data.
//
func (c *Client) AccessGroupsUpdateGroup(req AccessGroupsUpdateGroupRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessGroupsUpdateGroup: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessGroupsUpdateGroup: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/access/groups/%v", req.Groupid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessGroupsUpdateGroup: %v", err)
	}

	// output
	return nil
}

// AccessRolesCreateRole POST /access/roles
// Create new role.
//
func (c *Client) AccessRolesCreateRole(req AccessRolesCreateRoleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessRolesCreateRole: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessRolesCreateRole: %v", err)
	}

	err = c.Request("POST", "/access/roles", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessRolesCreateRole: %v", err)
	}

	// output
	return nil
}

// AccessRolesReadRole GET /access/roles/{roleid}
// Get role configuration.
//
func (c *Client) AccessRolesReadRole(req AccessRolesReadRoleRequest) (*AccessRolesReadRoleResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessRolesReadRole: %v", err)
	}

	var res *AccessRolesReadRoleResponse
	err = c.Request("GET", fmt.Sprintf("/access/roles/%v", req.Roleid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("AccessRolesReadRole: %v", err)
	}

	// output
	return res, nil
}

// AccessRolesUpdateRole PUT /access/roles/{roleid}
// Update an existing role.
//
func (c *Client) AccessRolesUpdateRole(req AccessRolesUpdateRoleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessRolesUpdateRole: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessRolesUpdateRole: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/access/roles/%v", req.Roleid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessRolesUpdateRole: %v", err)
	}

	// output
	return nil
}

// AccessRolesDeleteRole DELETE /access/roles/{roleid}
// Delete role.
//
func (c *Client) AccessRolesDeleteRole(req AccessRolesDeleteRoleRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessRolesDeleteRole: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/access/roles/%v", req.Roleid), nil, nil)

	if err != nil {
		return fmt.Errorf("AccessRolesDeleteRole: %v", err)
	}

	// output
	return nil
}

// AccessAclRead GET /access/acl
// Get Access Control List (ACLs).
//
func (c *Client) AccessAclRead() ([]*AccessAclReadResponse, error) {
	var err error

	var res []*AccessAclReadResponse
	err = c.Request("GET", "/access/acl", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("AccessAclRead: %v", err)
	}

	// output
	return res, nil
}

// AccessAclUpdate PUT /access/acl
// Update Access Control List (add or remove permissions).
//
func (c *Client) AccessAclUpdate(req AccessAclUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessAclUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessAclUpdate: %v", err)
	}

	err = c.Request("PUT", "/access/acl", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessAclUpdate: %v", err)
	}

	// output
	return nil
}

// AccessDomainsCreate POST /access/domains
// Add an authentication server.
//
func (c *Client) AccessDomainsCreate(req AccessDomainsCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessDomainsCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessDomainsCreate: %v", err)
	}

	err = c.Request("POST", "/access/domains", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessDomainsCreate: %v", err)
	}

	// output
	return nil
}

// AccessDomainsDelete DELETE /access/domains/{realm}
// Delete an authentication server.
//
func (c *Client) AccessDomainsDelete(req AccessDomainsDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessDomainsDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/access/domains/%v", req.Realm), nil, nil)

	if err != nil {
		return fmt.Errorf("AccessDomainsDelete: %v", err)
	}

	// output
	return nil
}

// AccessDomainsRead GET /access/domains/{realm}
// Get auth server configuration.
//
func (c *Client) AccessDomainsRead(req AccessDomainsReadRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessDomainsRead: %v", err)
	}

	err = c.Request("GET", fmt.Sprintf("/access/domains/%v", req.Realm), nil, nil)

	if err != nil {
		return fmt.Errorf("AccessDomainsRead: %v", err)
	}

	// output
	return nil
}

// AccessDomainsUpdate PUT /access/domains/{realm}
// Update authentication server settings.
//
func (c *Client) AccessDomainsUpdate(req AccessDomainsUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessDomainsUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessDomainsUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/access/domains/%v", req.Realm), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessDomainsUpdate: %v", err)
	}

	// output
	return nil
}

// AccessDomainsSync POST /access/domains/{realm}/sync
// Syncs users and/or groups from the configured LDAP to user.cfg. NOTE: Synced groups will have the name 'name-$realm', so make sure those groups do not exist to prevent overwriting.
//
func (c *Client) AccessDomainsSync(req AccessDomainsSyncRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessDomainsSync: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessDomainsSync: %v", err)
	}

	err = c.Request("POST", fmt.Sprintf("/access/domains/%v/sync", req.Realm), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessDomainsSync: %v", err)
	}

	// output
	return nil
}

// AccessOpenidAuthUrlAuthUrl POST /access/openid/auth-url
// Get the OpenId Authorization Url for the specified realm.
//
func (c *Client) AccessOpenidAuthUrlAuthUrl(req AccessOpenidAuthUrlAuthUrlRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessOpenidAuthUrlAuthUrl: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessOpenidAuthUrlAuthUrl: %v", err)
	}

	err = c.Request("POST", "/access/openid/auth-url", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessOpenidAuthUrlAuthUrl: %v", err)
	}

	// output
	return nil
}

// AccessOpenidLogin POST /access/openid/login
//  Verify OpenID authorization code and create a ticket.
//
func (c *Client) AccessOpenidLogin(req AccessOpenidLoginRequest) (*AccessOpenidLoginResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessOpenidLogin: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("AccessOpenidLogin: %v", err)
	}

	var res *AccessOpenidLoginResponse
	err = c.Request("POST", "/access/openid/login", strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("AccessOpenidLogin: %v", err)
	}

	// output
	return res, nil
}

// AccessTfaList GET /access/tfa
// List TFA configurations of users.
//
func (c *Client) AccessTfaList() ([]*AccessTfaListResponse, error) {
	var err error

	var res []*AccessTfaListResponse
	err = c.Request("GET", "/access/tfa", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("AccessTfaList: %v", err)
	}

	// output
	return res, nil
}

// AccessTfaVerify POST /access/tfa
// Finish a u2f challenge.
//
func (c *Client) AccessTfaVerify(req AccessTfaVerifyRequest) (*AccessTfaVerifyResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessTfaVerify: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("AccessTfaVerify: %v", err)
	}

	var res *AccessTfaVerifyResponse
	err = c.Request("POST", "/access/tfa", strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("AccessTfaVerify: %v", err)
	}

	// output
	return res, nil
}

// AccessTfaAddEntry POST /access/tfa/{userid}
// Add a TFA entry for a user.
//
func (c *Client) AccessTfaAddEntry(req AccessTfaAddEntryRequest) (*AccessTfaAddEntryResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessTfaAddEntry: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("AccessTfaAddEntry: %v", err)
	}

	var res *AccessTfaAddEntryResponse
	err = c.Request("POST", fmt.Sprintf("/access/tfa/%v", req.Userid), strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("AccessTfaAddEntry: %v", err)
	}

	// output
	return res, nil
}

// AccessTfaListUser GET /access/tfa/{userid}
// List TFA configurations of users.
//
func (c *Client) AccessTfaListUser(req AccessTfaListUserRequest) ([]*AccessTfaListUserResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessTfaListUser: %v", err)
	}

	var res []*AccessTfaListUserResponse
	err = c.Request("GET", fmt.Sprintf("/access/tfa/%v", req.Userid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("AccessTfaListUser: %v", err)
	}

	// output
	return res, nil
}

// AccessTfaDelete DELETE /access/tfa/{userid}/{id}
// Delete a TFA entry by ID.
//
func (c *Client) AccessTfaDelete(req AccessTfaDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessTfaDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/access/tfa/%v/%v", req.Userid, req.Id), nil, nil)

	if err != nil {
		return fmt.Errorf("AccessTfaDelete: %v", err)
	}

	// output
	return nil
}

// AccessTfaGetEntry GET /access/tfa/{userid}/{id}
// Fetch a requested TFA entry if present.
//
func (c *Client) AccessTfaGetEntry(req AccessTfaGetEntryRequest) (*AccessTfaGetEntryResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessTfaGetEntry: %v", err)
	}

	var res *AccessTfaGetEntryResponse
	err = c.Request("GET", fmt.Sprintf("/access/tfa/%v/%v", req.Userid, req.Id), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("AccessTfaGetEntry: %v", err)
	}

	// output
	return res, nil
}

// AccessTfaUpdateEntry PUT /access/tfa/{userid}/{id}
// Add a TFA entry for a user.
//
func (c *Client) AccessTfaUpdateEntry(req AccessTfaUpdateEntryRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessTfaUpdateEntry: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessTfaUpdateEntry: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/access/tfa/%v/%v", req.Userid, req.Id), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessTfaUpdateEntry: %v", err)
	}

	// output
	return nil
}

// AccessTicketCreate POST /access/ticket
// Create or verify authentication ticket.
//
func (c *Client) AccessTicketCreate(req AccessTicketCreateRequest) (*AccessTicketCreateResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("AccessTicketCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return nil, fmt.Errorf("AccessTicketCreate: %v", err)
	}

	var res *AccessTicketCreateResponse
	err = c.Request("POST", "/access/ticket", strings.NewReader(c.replaceBool(input.Encode())), &res)

	if err != nil {
		return nil, fmt.Errorf("AccessTicketCreate: %v", err)
	}

	// output
	return res, nil
}

// AccessPasswordChange PUT /access/password
// Change user password.
//
func (c *Client) AccessPasswordChange(req AccessPasswordChangeRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessPasswordChange: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("AccessPasswordChange: %v", err)
	}

	err = c.Request("PUT", "/access/password", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("AccessPasswordChange: %v", err)
	}

	// output
	return nil
}

// AccessPermissions GET /access/permissions
// Retrieve effective permissions of given user/token.
//
func (c *Client) AccessPermissions(req AccessPermissionsRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("AccessPermissions: %v", err)
	}

	err = c.Request("GET", "/access/permissions", nil, nil)

	if err != nil {
		return fmt.Errorf("AccessPermissions: %v", err)
	}

	// output
	return nil
}

// PoolsCreate POST /pools
// Create new pool.
//
func (c *Client) PoolsCreate(req PoolsCreateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("PoolsCreate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("PoolsCreate: %v", err)
	}

	err = c.Request("POST", "/pools", strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("PoolsCreate: %v", err)
	}

	// output
	return nil
}

// PoolsDelete DELETE /pools/{poolid}
// Delete pool.
//
func (c *Client) PoolsDelete(req PoolsDeleteRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("PoolsDelete: %v", err)
	}

	err = c.Request("DELETE", fmt.Sprintf("/pools/%v", req.Poolid), nil, nil)

	if err != nil {
		return fmt.Errorf("PoolsDelete: %v", err)
	}

	// output
	return nil
}

// PoolsRead GET /pools/{poolid}
// Get pool configuration.
//
func (c *Client) PoolsRead(req PoolsReadRequest) (*PoolsReadResponse, error) {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return nil, fmt.Errorf("PoolsRead: %v", err)
	}

	var res *PoolsReadResponse
	err = c.Request("GET", fmt.Sprintf("/pools/%v", req.Poolid), nil, &res)

	if err != nil {
		return nil, fmt.Errorf("PoolsRead: %v", err)
	}

	// output
	return res, nil
}

// PoolsUpdate PUT /pools/{poolid}
// Update pool data.
//
func (c *Client) PoolsUpdate(req PoolsUpdateRequest) error {
	var err error

	if err = c.validate(req); err != nil {
		// values not valid, deal with errors here
		return fmt.Errorf("PoolsUpdate: %v", err)
	}

	input, err := query.Values(req)
	if err != nil {
		return fmt.Errorf("PoolsUpdate: %v", err)
	}

	err = c.Request("PUT", fmt.Sprintf("/pools/%v", req.Poolid), strings.NewReader(c.replaceBool(input.Encode())), nil)

	if err != nil {
		return fmt.Errorf("PoolsUpdate: %v", err)
	}

	// output
	return nil
}

// Version GET /version
// API version details, including some parts of the global datacenter config.
//
func (c *Client) Version() (*VersionResponse, error) {
	var err error

	var res *VersionResponse
	err = c.Request("GET", "/version", nil, &res)

	if err != nil {
		return nil, fmt.Errorf("Version: %v", err)
	}

	// output
	return res, nil
}
