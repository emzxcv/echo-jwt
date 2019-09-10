build:
	$(info [+] Dockerising container with tag name go-jwt-server ...")
	docker build -t go-jwt-server .

run:
	$(info [+] Starting app on port :1323 ...")
	docker run -d -p 1323:1323 go-jwt-server

