
 # 版號更新時務必使用 1.0.1 三個數字組成
 # 安裝
```shell
 go install github.com/gioco-play/goctl-gfrpc
```

 
 # 使用
#### test 為資料夾
```shell
 goctl-gfrpc rpc protoc test/test.proto --zrpc_out=test --go-grpc_out=test --go_out=test --home template
```

# 遠端範本
```shell
goctl-gfrpc rpc protoc test.proto --zrpc_out=. --go-grpc_out=. --go_out=. --remote https://github.com/gioco-play/gf-template
```
 
 # 安裝依賴 
```shell
 go mod tidy 
```

 
 # 範本
```shell
 goctl-gfrpc template init --home $(pwd)/template
```