INSTALL_DIR = /usr/bin
GO = go
FLAGS = GOOS=linux GOARCH=386

net-compose: main.go
	${FLAGS} ${GO} build

.PHONY: install
install: net-compose	
	@cp ./net-compose ${INSTALL_DIR}