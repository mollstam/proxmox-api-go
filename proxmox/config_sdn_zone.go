package proxmox

import (
	"encoding/json"
	"fmt"
)

// ConfigSDNZone describes the Zone configurable element
type ConfigSDNZone struct {
	Type                     string `json:"type"`
	Zone                     string `json:"zone"`
	AdvertiseSubnets         bool   `json:"advertise-subnets,omitempty"`
	Bridge                   string `json:"bridge,omitempty"`
	BridgeDisableMacLearning bool   `json:"bridge-disable-mac-learning,omitempty"`
	Controller               string `json:"controller,omitempty"`
	DisableARPNDSuppression  bool   `json:"disable-arp-nd-suppression,omitempty"`
	DNS                      string `json:"dns,omitempty"`
	DNSZone                  string `json:"dnszone,omitempty"`
	DPID                     int    `json:"dp-id,omitempty"`
	ExitNodes                string `json:"exitnodes,omitempty"`
	ExitNodesLocalRouting    bool   `json:"exitnodes-local-routing,omitempty"`
	ExitNodesPrimary         string `json:"exitnodes-primary,omitempty"`
	IPAM                     string `json:"ipam,omitempty"`
	MAC                      string `json:"mac,omitempty"`
	MTU                      int    `json:"mtu,omitempty"`
	Nodes                    string `json:"nodes,omitempty"`
	Peers                    string `json:"peers,omitempty"`
	ReverseDNS               string `json:"reversedns,omitempty"`
	RTImport                 string `json:"rt-import,omitempty"`
	Tag                      int    `json:"tag,omitempty"`
	VlanProtocol             string `json:"vlan-protocol,omitempty"`
	VrfVxlan                 int    `json:"vrf-vxlan,omitempty"`
	// Pass a string of attributes to be deleted from the remote object
	Delete string `json:"delete,omitempty"`
	// Digest allows for a form of optimistic locking
	Digest string `json:"digest,omitempty"`
}

// NewConfigNetworkFromJSON takes in a byte array from a json encoded SDN Zone
// configuration and stores it in config.
// It returns the newly created config with the passed in configuration stored
// and an error if one occurs unmarshalling the input data.
func NewConfigSDNZoneFromJson(input []byte) (config *ConfigSDNZone, err error) {
	config = &ConfigSDNZone{}
	err = json.Unmarshal([]byte(input), config)
	return
}

func NewConfigSDNZoneFromApi(id string, client *Client) (config *ConfigSDNZone, err error) {
	zoneConfig, err := client.GetSDNZone(id)
	if err != nil {
		return nil, err
	}

	var zoneConfigData map[string]interface{}
	if _, isSet := zoneConfig["data"]; isSet {
		zoneConfigData = zoneConfig["data"].(map[string]interface{})
	}

	config = &ConfigSDNZone{}

	zone := ""
	if _, isSet := zoneConfigData["zone"]; isSet {
		zone = zoneConfigData["zone"].(string)
	}
	typ := ""
	if _, isSet := zoneConfigData["type"]; isSet {
		typ = zoneConfigData["type"].(string)
	}
	advertiseSubnets := false
	if _, isSet := zoneConfigData["advertise-subnets"]; isSet {
		advertiseSubnets = Itob(int(zoneConfigData["advertise-subnets"].(float64)))
	}
	bridge := ""
	if _, isSet := zoneConfigData["bridge"]; isSet {
		bridge = zoneConfigData["bridge"].(string)
	}
	bridgeDisableMacLearning := false
	if _, isSet := zoneConfigData["bridge-disable-mac-learning"]; isSet {
		bridgeDisableMacLearning = Itob(int(zoneConfigData["bridge-disable-mac-learning"].(float64)))
	}
	controller := ""
	if _, isSet := zoneConfigData["controller"]; isSet {
		controller = zoneConfigData["controller"].(string)
	}
	disableARPNDSuppression := false
	if _, isSet := zoneConfigData["disable-arp-nd-suppression"]; isSet {
		disableARPNDSuppression = Itob(int(zoneConfigData["disable-arp-nd-suppression"].(float64)))
	}
	dns := ""
	if _, isSet := zoneConfigData["dns"]; isSet {
		dns = zoneConfigData["dns"].(string)
	}
	dnsZone := ""
	if _, isSet := zoneConfigData["dnszone"]; isSet {
		dnsZone = zoneConfigData["dnszone"].(string)
	}
	dpid := 0
	if _, isSet := zoneConfigData["dp-id"]; isSet {
		dpid = int(zoneConfigData["dp-id"].(float64))
	}
	exitNodes := ""
	if _, isSet := zoneConfigData["exitnodes"]; isSet {
		exitNodes = zoneConfigData["exitnodes"].(string)
	}
	exitNodesLocalRouting := false
	if _, isSet := zoneConfigData["exitnodes-local-routing"]; isSet {
		exitNodesLocalRouting = Itob(int(zoneConfigData["exitnodes-local-routing"].(float64)))
	}
	exitNodesPrimary := ""
	if _, isSet := zoneConfigData["exitnodes-primary"]; isSet {
		exitNodesPrimary = zoneConfigData["exitnodes-primary"].(string)
	}
	ipam := ""
	if _, isSet := zoneConfigData["ipam"]; isSet {
		ipam = zoneConfigData["ipam"].(string)
	}
	mac := ""
	if _, isSet := zoneConfigData["mac"]; isSet {
		mac = zoneConfigData["mac"].(string)
	}
	mtu := 0
	if _, isSet := zoneConfigData["mtu"]; isSet {
		mtu = int(zoneConfigData["mtu"].(float64))
	}
	nodes := ""
	if _, isSet := zoneConfigData["nodes"]; isSet {
		nodes = zoneConfigData["nodes"].(string)
	}
	peers := ""
	if _, isSet := zoneConfigData["peers"]; isSet {
		peers = zoneConfigData["peers"].(string)
	}
	reverseDNS := ""
	if _, isSet := zoneConfigData["reversedns"]; isSet {
		reverseDNS = zoneConfigData["reversedns"].(string)
	}
	rtImport := ""
	if _, isSet := zoneConfigData["rt-import"]; isSet {
		rtImport = zoneConfigData["rt-import"].(string)
	}
	tag := 0
	if _, isSet := zoneConfigData["tag"]; isSet {
		tag = int(zoneConfigData["tag"].(float64))
	}
	vlanProtocol := ""
	if _, isSet := zoneConfigData["vlan-protocol"]; isSet {
		vlanProtocol = zoneConfigData["vlan-protocol"].(string)
	}
	vrfVxlan := 0
	if _, isSet := zoneConfigData["vrf-vxlan"]; isSet {
		vrfVxlan = int(zoneConfigData["vrf-vxlan"].(float64))
	}

	config.Zone = zone
	config.Type = typ
	config.AdvertiseSubnets = advertiseSubnets
	config.Bridge = bridge
	config.BridgeDisableMacLearning = bridgeDisableMacLearning
	config.Controller = controller
	config.DisableARPNDSuppression = disableARPNDSuppression
	config.DNS = dns
	config.DNSZone = dnsZone
	config.DPID = dpid
	config.ExitNodes = exitNodes
	config.ExitNodesLocalRouting = exitNodesLocalRouting
	config.ExitNodesPrimary = exitNodesPrimary
	config.IPAM = ipam
	config.MAC = mac
	config.MTU = mtu
	config.Nodes = nodes
	config.Peers = peers
	config.ReverseDNS = reverseDNS
	config.RTImport = rtImport
	config.Tag = tag
	config.VlanProtocol = vlanProtocol
	config.VrfVxlan = vrfVxlan

	return
}

func (config *ConfigSDNZone) CreateWithValidate(id string, client *Client) (err error) {
	err = config.Validate(id, true, client)
	if err != nil {
		return
	}
	return config.Create(id, client)
}

func (config *ConfigSDNZone) Create(id string, client *Client) (err error) {
	config.Zone = id
	params := config.mapToApiValues(true)
	return client.CreateSDNZone(params)
}

func (config *ConfigSDNZone) UpdateWithValidate(id string, client *Client) (err error) {
	err = config.Validate(id, false, client)
	if err != nil {
		return
	}
	return config.Update(id, client)
}

func (config *ConfigSDNZone) Update(id string, client *Client) (err error) {
	config.Zone = id
	params := config.mapToApiValues(false)
	err = client.UpdateSDNZone(id, params)
	if err != nil {
		params, _ := json.Marshal(&params)
		return fmt.Errorf("error updating SDN Zone: %v, (params: %v)", err, string(params))
	}
	return
}

func (c *ConfigSDNZone) Validate(id string, create bool, client *Client) (err error) {
	exists, err := client.CheckSDNZoneExistance(id)
	if err != nil {
		return
	}
	if exists && create {
		return ErrorItemExists(id, "zone")
	}
	if !exists && !create {
		return ErrorItemNotExists(id, "zone")
	}

	if create {
		err = ValidateStringInArray([]string{"evpn", "qinq", "simple", "vlan", "vxlan"}, c.Type, "type")
		if err != nil {
			return
		}
	}
	switch c.Type {
	case "simple":
	case "vlan":
		if create {
			if c.Bridge == "" {
				return ErrorKeyEmpty("bridge")
			}
		}
	case "qinq":
		if create {
			if c.Bridge == "" {
				return ErrorKeyEmpty("bridge")
			}
			if c.Tag <= 0 {
				return ErrorKeyEmpty("tag")
			}
			if c.VlanProtocol == "" {
				return ErrorKeyEmpty("vlan-protocol")
			}
		}
	case "vxlan":
		if create {
			if c.Peers == "" {
				return ErrorKeyEmpty("peers")
			}
		}
	case "evpn":
		if create {
			if c.VrfVxlan < 0 {
				return ErrorKeyEmpty("vrf-vxlan")
			}
			if c.Controller == "" {
				return ErrorKeyEmpty("controller")
			}
		}
	}
	if c.VlanProtocol != "" {
		err = ValidateStringInArray([]string{"802.1q", "802.1ad"}, c.VlanProtocol, "vlan-protocol")
		if err != nil {
			return
		}
	}
	return
}

func (config *ConfigSDNZone) mapToApiValues(create bool) (params map[string]interface{}) {

	d, _ := json.Marshal(config)
	json.Unmarshal(d, &params)

	boolsToFix := []string{
		"advertise-subnets",
		"bridge-disable-mac-learning",
		"disable-arp-nd-suppression",
		"exitnodes-local-routing",
	}
	for _, key := range boolsToFix {
		if v, has := params[key]; has {
			params[key] = Btoi(v.(bool))
		}
	}
	if !create {
		delete(params, "type")
	}
	return
}
