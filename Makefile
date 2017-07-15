.PHONY: db
ipapi:
	go run main.go ipapi google.com
keycdn:
	go run main.go keycdn google.com
db:
	docker run --name tidb-server -d -p 4000:4000 pingcap/tidb:latest
