
inspired by: https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/

# trying tcp server
# start server
go run main.go -port=1234
# start client
go run main.go -client -host=127.0.0.1 -port=1234
send "Hello!"
send "STOP"

# Note: Same works with "us" and "uc".

# initial
go mod init github.com/mkinney/snet
go mod tidy

Note: Generated a personal access token with perms to do releases.

workflow:
- commit changes
- tag release
git tag -a v0.2 -m "v0.2"
git push origin --tags
- create a release in GitHub for that tag
git push

