

.PHONY: proto
proto:
	docker run --rm -v d:/GOLANG/src/taobao/cartApi:/d/GOLANG/src/taobao/cartApi -w /d/GOLANG/src/taobao/cartApi  -e ICODE=2606C833CD172F4C cap1573/cap-protoc -I ./ --micro_out=./ --go_out=./ ./proto/cartApi/cartApi.proto

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cartApi-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t cartApi-api:latest
