# Proxmox-API #

proxmox-api is a Go client for accessing and managing Proxmox VE instances using API.

Now, full disclosure. I only recently started my journey with proxmox while migrating from different hypervisor. 
One of my requirements when selecting new hypervisor was API coverage good enough to automate everything from provisioning VM's to designing network around them.
Since there was none, this project is my effort to provide such API for PVE.

There are few things you need to know before you try it:
- Files prefixed with client_* are 100% generated from json schema.
- This repository contains 450 generated client methods and 618 generated structs
- I couldn't find everything I needed to generate this code using "official" tools, so I wrote custom generator to do it for me (I will publish it soon as well) 
- I spend only few hours on it, so I took every shortcut I could to make this work ;-)

While https://github.com/Telmate/proxmox-api-go exists it only implements small subset of API endpoints. Also, I had to try and see if generator will work ;-)   

# Examples

Each example should have this at the top of your go file:
```go
import (
    "crypto/tls"
    "fmt"
    "log"
    
    "github.com/r3dsh/proxmox-api/v1"
)

c, err := proxmox.NewClient(&proxmox.Config{
    ApiHost:  "https://pve.local:8006",
    Username: "proxmox-client@pve",
    Password: "Passw0rd",
    TLS:      &tls.Config{InsecureSkipVerify: true}, // you know, remove it for production ;-)
})
if err != nil {
    log.Fatalln(err)
}
// c.Simulate = true
```

Getting proxmox version:
```go
ver, err := c.Version()
if err != nil {
    log.Println(err)
}
fmt.Println("PROXMOX VERSION:", ver.Version, ver.Release, ver.Repoid)
```
output:
```shell
PROXMOX VERSION: 7.2-5 7.2 12f1e639
```

Listing containers:
```go
containers, err := c.NodesLxcVmlist(proxmox.NodesLxcVmlistRequest{Node: "dc1"})
if err != nil {
    log.Fatal(err)
}

for _, c := range containers {
    fmt.Println("   ", vm.Vmid, vm.Name, vm.Cpus, vm.Maxmem, vm.Status)
}
```
output:
```shell
103 test2 1 536870912 running
105 test10 1 536870912 running
101 bastion2 1 536870912 running
100 bastion1 1 536870912 running
```

Creating new bridge:
```go
err = c.NodesNetworkCreate(proxmox.NodesNetworkCreateRequest{
    Node:      "dc1",
    Iface:     "vmbr8",
    Cidr:      "10.122.22.0/24",
    Type:      proxmox.NetworkOVSBridge,
    Comments:  "NEW SDK Created Device - this is crazy!",
    Autostart: true,
})
if err != nil {
    log.Fatal(err)
}
// we need to reload proxmox configuration
err = c.NodesNetworkReloadConfig(proxmox.NodesNetworkReloadConfigRequest{Node: "dc1"})
if err != nil {
    log.Fatal(err)
}
```
No output, but bridge should be there! Let's remove it:
```go
err = c.NodesNetworkDelete(proxmox.NodesNetworkDeleteRequest{
    Node:  "dc1",
    Iface: "vmbr8",
})
if err != nil {
    log.Println(err)
}
```

Every method, type and property are documented based on proxmox docs attached to the schema.

# TODO #

- Enum fields validation support
- Format field validation support
