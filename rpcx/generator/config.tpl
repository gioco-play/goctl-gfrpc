package config

import (
    "github.com/zeromicro/go-zero/zrpc"
    "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Consul consul.Conf
    Postgres struct {
        Host              string
        Port              int
        SlavePort         int
        UserName          string
        Password          string
        DBName            string
        DBTimezone        string
        DBPoolMin         int
        DBPoolMax         int
        DBConnMaxLifetime int
        DBDebugLevel      string
    }
    RedisCache struct {
        RedisSentinelNode string
        RedisMasterName   string
        RedisDB           int
    }
}
