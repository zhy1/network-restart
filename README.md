# for network restart condition check.

SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build src/main.go


SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build src/main.go -o ./build/network-linux-x64
scp network-linux-x64 SERVERHOST