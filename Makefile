build-arm:
	GOOS=linux GOARCH=arm64 go build -o packetnet-capsule-arm64

build:
	GOOS=linux GOARCH=amd64 go build -o packetnet-capsule-amd64

clean:
	rm packetnet-capsule-{amd64,arm64}

docker-arm64: build-arm
	docker build -t jmarhee/packetnet-capsule.arm64:${TAG} -f Dockerfile.arm64 .
docker-arm64: clean

docker-amd64: build
	docker build -t jmarhee/packetnet-capsule.amd64:${TAG} -f Dockerfile.amd64 .
docker-amd64: clean
