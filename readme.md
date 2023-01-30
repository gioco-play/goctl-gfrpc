```
 # 版號更新時務必使用 1.0.1 三個數字組成
 # 安裝
 go install github.com/gioco-play/goctl-gfrpc@latest 
 
 # 使用
 ### test 為資料夾
 goctl-gfrpc rpc protoc test/test.proto --zrpc_out=test --go-grpc_out=test --go_out=test --home template
 
 # 安裝依賴 
 go mod tidy 
 
 # 範本
 goctl-gfrpc template init --home $(pwd)/template
```