VERSION = "$(shell git rev-parse --short HEAD)"

ifeq ($(PREFIX),)
PREFIX = /usr/local
endif

NAME = slow

build:	
	go build -ldflags "-X main.VERSION=$(VERSION)" -o bin/$(NAME)
install:
	@mkdir -p $(PREFIX)/bin
	mv bin/$(NAME) $(PREFIX)/bin/$(NAME)
