APP_NAME={{.serviceName}}.rpc
APP_HOST={{.host}}
APP_TAG=["{{.serviceName}}"]

DB_HOST=192.168.10.101
DB_PORT=5432
DB_SLAVE_PORT=5432
DB_DATABASE=gpsbo
DB_USERNAME=postgres
DB_PASSWORD=1qaz2wsx
DB_POOL_MIN=1
DB_POOL_MAX=10
DB_TIMEZONE=UTC
DB_CONN_MAX_LIFETIME=180
DB_DEBUG_LEVEL=error

CONSUL_HOST=192.168.10.93:8500
CONSUL_TTL=20

TRACE_ENDPOINT=http://172.16.204.124:14268/api/traces

LOG_MODE=console

REDIS_SENTINEL_NODE=192.168.30.38:26379;192.168.30.39:26379;192.168.30.40:26379
REDIS_MASTER_NAME=mymaster
REDIS_DB=1
REDIS_PREFIX=gps::