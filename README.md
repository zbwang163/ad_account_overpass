# ad.info.account的rpc代码
生成客户端代码和服务端代码命令：
```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ad_info_account_rpc.proto
```