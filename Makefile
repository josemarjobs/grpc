clean_server:
	rm server
clean_client:
	rm client

server: clean_server
	go build grpcdemo/server

client: clean_client
	go build grpcdemo/client