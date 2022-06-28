package proxmox

type Ticket struct {
    CSRFPreventionToken string `json:"CSRFPreventionToken"`
    Ticket              string `json:"ticket"`
    Username            string `json:"username"`
}
