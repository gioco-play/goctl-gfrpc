
 # 版號更新時務必使用 1.0.1 三個數字組成
 # 安裝
```shell
 go install github.com/gioco-play/goctl-gfrpc@latest
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

# add-dep（注入依賴到既有服務）
```shell
 goctl-gfrpc rpc add-dep --name=transaction --remote https://github.com/gioco-play/gf-template
 goctl-gfrpc rpc add-dep --name=notify,grabber --remote https://github.com/gioco-play/gf-template
```
須在服務目錄（含 `internal/`、`etc/` 的那層）下執行。`--name` 可用逗號分隔一次注入多個依賴（依序執行，非交易性）。依賴定義見 [gf-template 的 deps/ 說明](https://github.com/gioco-play/gf-template#deps)。
