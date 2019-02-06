conc-mutex:
	@go run concurrency/main.go mutex

conc-chan:
	@go run concurrency/main.go channel

powlite:
	@go run pow-lite/main.go pow-lite/blockchain.go pow-lite/server.go

rpc-dev:
	@go run rpc/server/main.go &
	@go run rpc/client/main.go

rpc-down:
	@kill $(lsof -t -i :4040)

rpc-server:
	@go run rpc/server/main.go

rpc-client:
	@go run rpc/client/main.go