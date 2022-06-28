# Proxmox-API #

proxmox-api is a Go client for accessing and managing Proxmox VE instances using API.

Now, full disclosure. I only recently started my journey with proxmox while migrating from different hypervisor. 
One of my requirements when selecting new hypervisor was API coverage good enough to automate everything from provisioning VM's to designing network around them.
This project is my effort to provide such API for PVE.

There are few things you need to know before you try it:
- Files prefixed with client_* are 100% generated from json schema.
- I couldn't find everything I needed to generate this code using "official" tools, so I wrote custom generator to do it for me (I will publish it soon as well) 
- I spend only few hours on it, so I took every shortcut I could to make this work ;-)

While https://github.com/Telmate/proxmox-api-go exists it only implements small subset of API endpoints. Also, I had to try and see if generator will work ;-)   

# TODO #

- Enum fields validation support
- Format field validation support
