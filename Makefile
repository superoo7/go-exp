conc-mutex:
	@go run concurrency/main.go mutex

conc-chan:
	@go run concurrency/main.go channel

powlite:
	@go run pow-lite/main.go pow-lite/blockchain.go pow-lite/server.go

