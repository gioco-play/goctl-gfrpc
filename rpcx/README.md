# Rpc Generation

Goctl Rpc 是 `goctl` 腳手架下的一個 rpc 服務程式碼產生模組，支援 proto 範本產生和 rpc 服務程式碼產生，透過此工具產生程式碼你只需要專注在商業邏輯的撰寫，不用再寫一些重複性的程式碼。這讓我們能把心力放在商業邏輯上，進而加快開發效率並降低出錯率。

## 特性

* 簡單易用
* 快速提升開發效率
* 出錯率低
* 貼近 protoc


## 快速開始

### 方式一：快速產生 greet 服務

  透過指令 `goctl rpc new ${servieName}` 產生

  如產生 greet rpc 服務：

  ```Bash
  goctl rpc new greet
  ```

  執行後程式碼結構如下:

```text
.
└── greet
    ├── etc
    │   └── greet.yaml
    ├── greet
    │   ├── greet.go
    │   ├── greet.pb.go
    │   └── greet_grpc.pb.go
    ├── greet.go
    ├── greet.proto
    └── internal
        ├── config
        │   └── config.go
        ├── logic
        │   └── pinglogic.go
        ├── server
        │   └── greetserver.go
        └── svc
            └── servicecontext.go
```

### 方式二：透過指定 proto 產生 rpc 服務

* 產生 proto 範本

```Bash
$ goctl rpc template -o=user.proto
```
  
```proto
syntax = "proto3";

package user;
option go_package="./user";

message Request {
  string ping = 1;
}

message Response {
  string pong = 1;
}

service User {
  rpc Ping(Request) returns(Response);
}
```
  

* 產生 rpc 服務程式碼

```bash
$ goctl rpc protoc  user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```


## 用法

### rpc 服務產生用法

```Bash
$ goctl rpc protoc -h
Generate grpc code

Usage:
  goctl rpc protoc [flags]

Examples:
goctl rpc protoc xx.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.

Flags:
      --branch string     The branch of the remote repo, it does work with --remote
  -h, --help              help for protoc
      --home string       The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
  -m, --multiple          Generated in multiple rpc service mode
      --remote string     The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                          	The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
      --style string      The file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md] (default "gozero")
  -v, --verbose           Enable log output
      --zrpc_out string   The zrpc output directory
```

### 參數說明

* --branch 指定遠端倉庫範本分支
* --home 指定 goctl 範本根目錄
* -m, --multiple 指定產生多個 rpc 服務模式，預設為 false，若為 false，則只支援產生一個 rpc service；若為 true，則支援產生多個 rpc service，且多個 rpc service 會分組。
* --style 指定檔案輸出格式
* -v, --verbose 顯示日誌
* --zrpc_out 指定 zrpc 輸出目錄

> ## --multiple
> 是否開啟多個 rpc service 產生，若開啟，則具備以下新特性
> 1. 支援 1 到多個 rpc service
> 2. 產生的 rpc 服務會依照服務名稱分組（即便只有一個 rpc service）
> 3. rpc client 的檔案目錄改為固定名稱 `client`
>
> 若不開啟，則和舊版本 rpc 產生邏輯一樣（相容）
> 1. 有且只能有一個 rpc service

### add-dep 用法

```Bash
$ goctl-gfrpc rpc add-dep -h
Inject a dependency into an existing service

Usage:
  goctl-gfrpc rpc add-dep [flags]

Examples:
goctl-gfrpc rpc add-dep --name=transaction --remote https://github.com/gioco-play/gf-template

Flags:
      --branch string   The branch of the remote repo, it does work with --remote
  -h, --help            help for add-dep
      --home string     The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --name string     The dependency name, matching <template repo>/deps/<name>.tpl
      --remote string   The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
```

#### 參數說明

* --name 指定要注入的依賴名稱，對應 `<template repo>/deps/<name>.tpl`
* --home 指定 goctl 範本根目錄
* --remote 指定遠端倉庫範本網址
* --branch 指定遠端倉庫範本分支

#### 用途

把一個依賴（rpc client、consul 設定讀取等）注入到**已經產生好**的服務裡，補齊 `internal/config/config.go`、`internal/svc/servicecontext.go`、`etc/*.yaml`、`etc/.env` 中該依賴需要的部分。必須在服務目錄（含 `internal/`、`etc/` 的那層，跟 `make rpc` 同一層）下執行。重複執行同一個 `--name` 不會重複插入。依賴的定義方式見 [gf-template 的 deps/ 說明](https://github.com/gioco-play/gf-template#deps)。

#### 範例

```Bash
cd obfish-vendor-go/rpc
goctl-gfrpc rpc add-dep --name=transaction --remote https://github.com/gioco-play/gf-template
```

服務目錄下的 `makefile` 已內建對應 target，也可以直接用：

```Bash
make add-dep NAME=transaction
```

`REMOTE` 預設為 `https://github.com/gioco-play/gf-template`，可用 `make add-dep NAME=xxx REMOTE=...` 覆寫。
> 註：此 target 只會出現在新產生的服務（`rpc new` / 第一次 `rpc protoc`）；已存在的舊 `makefile` 需自行手動加上。

## rpc 服務產生 example
詳情見 [example/rpc](https://github.com/zeromicro/go-zero/tree/master/tools/goctl/example)

## --multiple 為 true 和 false 的目錄差異
來源 proto 檔案

```protobuf
syntax = "proto3";

package hello;

option go_package = "./hello";

message HelloReq {
  string in = 1;
}

message HelloResp {
  string msg = 1;
}

service Greet {
  rpc SayHello(HelloReq) returns (HelloResp);
}
```

### --multiple=true

```text
hello
├── client // 差異1：rpc client 目錄固定為 client 名稱
│   └── greet // 差異2：會依照 rpc service 名稱分組
│       └── greet.go
├── etc
│   └── hello.yaml
├── hello.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── logic
│   │   └── greet // 差異2：會依照 rpc service 名稱分組
│   │       └── sayhellologic.go
│   ├── server
│   │   └── greet // 差異2：會依照 rpc service 名稱分組
│   │       └── greetserver.go
│   └── svc
│       └── servicecontext.go
└── pb
    └── hello
        ├── hello.pb.go
        └── hello_grpc.pb.go
```

### --multiple=false (舊版本目錄，向後相容)
```text
hello
├── etc
│   └── hello.yaml
├── greet
│   └── greet.go
├── hello.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── logic
│   │   └── sayhellologic.go
│   ├── server
│   │   └── greetserver.go
│   └── svc
│       └── servicecontext.go
└── pb
    └── hello
        ├── hello.pb.go
        └── hello_grpc.pb.go
```
