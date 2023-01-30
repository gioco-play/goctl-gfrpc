lang:
	easyi18n generate --pkg=locales ../locales ../locales/locales.go

rpc:
	 goctl-gfrpc rpc protoc {{.serviceName}}.proto --zrpc_out=. --go-grpc_out=. --go_out=. --home ../template

run:
	go run {{.serviceName}}.go -f etc/{{.serviceName}}.yaml -env etc/.env

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/{{.serviceName}}_service {{.serviceName}}.go