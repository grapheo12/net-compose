INSTALL_DIR = /usr/bin
GO = go
FLAGS = GOOS=linux GOARCH=386

.PHONY: build/net-compose
build/net-compose:
	${FLAGS} ${GO} build -o build/net-compose ./cmd/net-compose

.PHONY: install
install: build/net-compose	
	@cp build/net-compose ${INSTALL_DIR}