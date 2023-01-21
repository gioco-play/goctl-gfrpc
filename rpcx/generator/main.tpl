package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/joho/godotenv"
    "github.com/spf13/viper"
	{{.imports}}

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
    "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var (
    configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")
    envFile    = flag.String("env", "etc/.env", "the env file")
)

func main() {
	flag.Parse()

    err := godotenv.Load(*envFile)
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
    // 配置額外參數
    viper.SetConfigName(".env")
    viper.SetConfigType("env")
    viper.SetConfigFile(*envFile)
    err = viper.ReadInConfig() // Find and read the config file
    if err != nil {            // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %w \n", err))
    }
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
{{range .serviceNames}}       {{.Pkg}}.Register{{.Service}}Server(grpcServer, {{.ServerPkg}}.New{{.Service}}Server(ctx))
{{end}}
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 注册服务
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
