BINARY				:= $(shell basename -s .git `git config --get remote.origin.url`)

.PHONY: run-1
run-1: 
	SERVICE_NAME="my-local-1" \
	SERVICE_ADDR=localhost:8001 \
	CLUSTER_HOSTS="localhost:8002,localhost:8003,localhost:8004,localhost:8005" \
	ELECTION_TIMEOUT=150 go run ./cmd/$(BINARY)/main.go
.PHONY: run-2
run-2: 
	SERVICE_NAME="my-local-2" \
	SERVICE_ADDR=localhost:8002 \
	CLUSTER_HOSTS="localhost:8001,localhost:8003,localhost:8004,localhost:8005" \
	ELECTION_TIMEOUT=250 go run ./cmd/$(BINARY)/main.go

.PHONY: run-3
run-3: 
	SERVICE_NAME="my-local-3" \
	SERVICE_ADDR=localhost:8003 \
	CLUSTER_HOSTS="localhost:8001,localhost:8002,localhost:8004,localhost:8005" \
	ELECTION_TIMEOUT=300 go run ./cmd/$(BINARY)/main.go

.PHONY: run-4
run-4: 
	SERVICE_NAME="my-local-4" \
	SERVICE_ADDR=localhost:8004 \
	CLUSTER_HOSTS="localhost:8001,localhost:8002,localhost:8003,localhost:8005" \
	ELECTION_TIMEOUT=400 go run ./cmd/$(BINARY)/main.go

.PHONY: run-5
run-5: 
	SERVICE_NAME="my-local-5" \
	SERVICE_ADDR=localhost:8005 \
	CLUSTER_HOSTS="localhost:8001,localhost:8002,localhost:8003,localhost:8004" \
	ELECTION_TIMEOUT=500 go run ./cmd/$(BINARY)/main.go

.PHONY: proto
proto:
	./gen-go.sh