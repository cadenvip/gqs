all:install

install:
	go install github.com/sunisdown/gqs
	go install github.com/sunisdown/gqs/services/greeter_server

start:
	./misc/server.sh start

stop:
	./misc/server.sh stop
