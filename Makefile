appname = goapple

Default:
	go run cmd/main.go
start:
	./bin/$(appname)
build:
	go build -o bin/$(appname) cmd/main.go