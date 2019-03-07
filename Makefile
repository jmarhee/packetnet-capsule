build-arm:
	GOOS=linux GOARCH=arm64 go build -o packetnet-capsule-arm64

build:
	GOOS=linux GOARCH=amd64 go build -o packetnet-capsule-amd64

clean:
	rm packetnet-capsule-{amd64,arm64}

