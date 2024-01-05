Simple network TCP or UDP client and server utility. Used for validation that the host and port is open. (i.e., not blocked by firewalls)

Note: This is a repo to test out various things. If you really want a utility to test tcp and udp connections, check out *netcat*.
Start tcp server:
  nc -l 1234
Start tcp client:
  nc 127.0.01 1234
Start udp server:
  nc -l 1235 -u
Start udp client:
  nc 127.0.01 1235 -u

Start TCP server:
  snet -port 3000
Start TCP client:
  snet -client -host:1.2.3.4 -port 3000
When done, type in "STOP" in the client and "control-c" on the server.

Start UDP server:
  snet -port 3001 -udp
Start UDP client:
  snet -client -host:1.2.3.4 -port 3001 -udp

On windows, might need to open an elevated prompt:
  netsh advfirewall firewall add rule name="TCP Port 3000" dir=in action=allow protocol=TCP localport=3000
  netsh advfirewall firewall add rule name="UDP Port 3001" dir=in action=allow protocol=UDP localport=3001

On linux, if using UFW:
  sudo ufw allow 3000
  sudo ufw allow 3001/udp
  sudo ufw allow from 1.2.3.4 to any port 3000

inspired by: https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/

# Development notes:

# trying tcp server
# start server
go run main.go -port=1234
# start client
go run main.go -client -host=127.0.0.1 -port=1234
send "Hello!"
send "STOP"

# initial
go mod init github.com/mkinney/snet
go mod tidy

Note: Generated a personal access token with perms to do releases.

workflow for a release:
- commit changes
- tag release
git tag -a v0.2 -m "v0.2"
git push origin --tags
- create a release in GitHub for that tag
git push


See these other repos:
https://github.com/mkinney/snet_win To build nupkg file that can be used for a Windows Chocolatey install 
https://github.com/mkinney/snet_rpm To build rpm package for Red Hat based linux systems
https://github.com/mkinney/snet_packer To build an Ubuntu based docker image with snet installed

