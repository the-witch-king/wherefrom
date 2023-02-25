#makefile

.DEFAULT_GOAL := serve

run:
	cd server; go run main.go types.go

watch:
	ulimit -n 1000
	/Users/mike/.asdf/installs/golang/1.20.1/packages/bin/reflex -s -r '\.go$$' make run
