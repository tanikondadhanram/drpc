gen-go:
	protoc --go_out=. --go-drpc_out=. data.proto
gen-dart:
	protoc --dart_out=. proto/data.proto
clean:
	rm pb/*.go
run:
	go run main.go
