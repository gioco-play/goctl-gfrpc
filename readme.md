```
 # 安裝
 go install github.com/gioco-play/goctl-gfrpc@latest 
 
 # 使用
 ### test 為資料夾
goctl-gfrpc rpc protoc test/test.proto --zrpc_out=test --go-grpc_out=test --go_out=test 
 
 # 安裝依賴 
 go mod tidy 
```