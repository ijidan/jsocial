APP     			= jsocial
PACKAGE 			=
OUTPUT_BUILD_DIR 	= /data
PROJECT_DIR			=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))
API_DIR				=$(PROJECT_DIR)api
BUILD_DIR			=$(PROJECT_DIR)build
CMD_DIR				=$(PROJECT_DIR)cmd
CONFIGS_DIR			=$(PROJECT_DIR)configs
INTERNAL_DIR		=$(PROJECT_DIR)internal

.PHONY :all help dl build run compose clean gormt test grpcurl start status stop proto_update command compose_up compose_down docker_stop docker_rm

help:
	echo "make check -  查看相关信息"
	echo "make proto -  grpc编译"
	echo "make tidy -  Go Mod tidy"
	echo "make dl -  Go Mod下载"
	echo "make build - 编译 Go 代码, 生成二进制文件"
	echo "make run - 直接运行 Go 代码"
	echo "make compose - docker-compose直接运行 Go 代码"
	echo "make clean - 清除vendor"
	echo "make gormt - 使用gormt自动生成model"
	echo "make test - 执行测试代码"
	echo "make grpcurl - 查看当前的服务"
	echo "make start - goreman start"
	echo "make status - goreman run status"
	echo "make stop - goreman run stop"
	echo "make proto_update -proto 更新"
	echo "make command -显示所有bin目录命令"
	echo "make compose_up  -启动docker-compose"
	echo "make compose_down  -关闭docker-compose"
	echo "make docker_stop -停止所有docker容器"
	echo "make docker_rm -- 删除所有docker容器"
	echo "make test_api -- 运行API"
	echo "make build_api -- 构建API"
	echo "make start_api -- 启动API"
	echo "make ping_api -- ping API"
	echo "make wire --执行wire"


check:
	echo $(API_DIR)
proto: dl
	protoc -I=internal/jim_proto/proto \
    		-I=$(GOPATH)/pkg/mod \
    		-I=$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gatewayv1.16.0/third_party/googleapis \
    		-I=$(GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validatev0.6.3 \
          --doc_out=internal/jim_proto/ --doc_opt=html,docs.html  internal/jim_proto/proto/*.proto

	protoc -I=internal/jim_proto/proto \
		-I=$(GOPATH)/pkg/mod \
		-I=$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gatewayv1.16.0/third_party/googleapis \
		-I=$(GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validatev0.6.3 \
      --go_out=internal/jim_proto/ --go-grpc_out=internal/jim_proto/ --grpc-gateway_out=internal/jim_proto/ --validate_out="lang=go:internal/jim_proto/" \
      --doc_out=internal/jim_proto/ --doc_opt=markdown,docs.md  \
      --grpc-gateway_opt logtostderr=true internal/jim_proto/proto/*.proto
tidy:
	go mod tidy
dl:
	go mod download
build:
	echo "Building  app..."
	rm -rf ${OUTPUT_BUILD_DIR}
	mkdir -p ${OUTPUT_BUILD_DIR}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o ${OUTPUT_BUILD_DIR}/${APP} -ldflags '-w -s'
run: build
	${OUTPUT_BUILD_DIR}/${APP}
compose:
	echo "Building ${APP} app in docker..."
	docker-compose up --remove-orphans -d
clean:
	-echo "Cleaning..."
	-rm -rf vendor
gormt:
	gormt
gen:
	go run cmd/main/main.go gen_gorm
token:
	go run cmd/main/main.go gen_token
test:
	go test -v  ./test
grpcurl:
	 grpcurl -plaintext 127.0.0.1:9093 list
grpcui:run
	grpcui -plaintext 127.0.0.1:9093
start:build
	goreman start
status:
	goreman run status
stop:
	goreman run stop
check:
proto_update:
	cd internal/jim_proto && git pull origin master
db_up:
	migrate -path migrate -database "mysql://root:roottcp(127.0.0.1:3306)/jim" -verbose up
db_down:
	migrate -path migrate -database "mysql://root:roottcp(127.0.0.1:3306)/jim" -verbose down
compose_up:
	docker-compose -f ./deployments/docker-compose.yaml  up -d
compose_down:
	docker-compose -f ./deployments/docker-compose.yaml  down
docker_stop:
	docker stop $(docker ps -a -q)
docker_rm: docker_stop
	docker rm $(docker ps -a -q)
test_api:
	go run $(CMD_DIR)/api/main.go -c $(CONFIGS_DIR)/api.yaml -d $(PROJECT_DIR)
build_api:
	go build -o $(CMD_DIR)/api/api $(CMD_DIR)/api/main.go
start_api:build_api
	nohup $(CMD_DIR)/api/api &
ping_api:
	curl 127.0.0.1:9096/api/v1/ping
wire:
	cd $(INTERNAL_DIR)/injector && wire
command:
	echo "goreman  gormt  grpcui  grpcurl  protoc-gen-doc  protoc-gen-go  protoc-gen-go-grpc  protoc-gen-govalidators  protoc-gen-grpc-gateway  protoc-gen-openapiv2  protoc-gen-validate"

