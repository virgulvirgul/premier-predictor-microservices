PROJECT_NAME = premier-predictor-microservices

_gendart:
	@mkdir -p model/gen/ship/dart
	@protoc -I=model/protodefs --plugin=/Users/chrisshepherd/.pub-cache/bin/protoc-gen-dart --dart_out=grpc:model/gen/ship/dart model/protodefs/*.proto

_gengo:
	@mkdir -p model/gen
	@protoc -I=model/protodefs --go_out=plugins=grpc:model/gen model/protodefs/*.proto

gen: _gengo _gendart

build: get gen
	@env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags '-w -extldflags "-static"' -o build/${PROJECT_NAME}_linux_amd64 .
	@env GOARCH=amd64 go build -ldflags '-w -extldflags "-static"' -o build/${PROJECT_NAME}_macosx_amd64 .

get:
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure

install: get gen
	@cp config_template.json config.json
