package proxmox

const (
    NetworkLinuxBridge = "bridge"
    NetworkBond        = "bond"
    NetworkETH         = "eth"
    NetworkAlias       = "alias"
    NetworkVLAN        = "vlan"
    NetworkOVSBridge   = "OVSBridge"
    NetworkOVSBond     = "OVSBond"
    NetworkOVSPort     = "OVSPort"
    NetworkOVSIntPort  = "OVSIntPort"
    NetworkUnknown     = "unknown"
)

type Ticket struct {
    CSRFPreventionToken string `json:"CSRFPreventionToken"`
    Ticket              string `json:"ticket"`
    Username            string `json:"username"`
}
