#makefile

run-fe:
	cd server/frontend; npm run start

run-server:
	cd server; go run main.go types.go

build-fe:
	cd server/frontend; npm run build

build-be:
	cd server; GOOS=linux go build -o server main.go types.go

build-lambda:
	make build-be; zip wherefrom.zip server/server

watch-server:
	ulimit -n 1000
	/Users/mike/.asdf/installs/golang/1.20.1/packages/bin/reflex -s -r '\.go$$' make run-server
