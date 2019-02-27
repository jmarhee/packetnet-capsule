build-arm:
	GOOS=linux GOARCH=arm64 go build -o packetnet-capsule

build:
	GOOS=linux GOARCH=amd64 go build -o packetnet-capsule


