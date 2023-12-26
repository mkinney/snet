
inspired by: https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/

# trying tcp server
# start server
go run ts.go 1234
# start client
go run tc.go 127.0.0.1:1234
send "Hello!"
send "STOP"

# Note: Same works with "us" and "uc".

# initial
go mod init github.com/mkinney/snet
go mod tidy

Note: Generated a personal access token with perms to do releases.

