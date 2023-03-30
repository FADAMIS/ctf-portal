build:
	- go build -o bin/
	- sudo setcap 'cap_net_bind_service=+ep' bin/ctf-portal