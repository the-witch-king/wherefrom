#makefile

run-fe:
	cd server/frontend; npm run start

run-server:
	cd server; go run main.go types.go

build-fe:
	cd server/frontend; npm run build

build-be:
	cd server; GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../main main.go types.go

build-lambda:
	make build-be; zip wherefrom.zip main

watch-server:
	ulimit -n 1000
	/Users/mike/.asdf/installs/golang/1.20.1/packages/bin/reflex -s -r '\.go$$' make run-server
