#makefile

.DEFAULT_GOAL := serve

serve:
	cd server/frontend; npm run build
	cd server; go run main.go types.go
