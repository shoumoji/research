http2Dir := $(CURDIR)/http2
http3Dir := $(CURDIR)/http3


.PHONY: http2-start
http2-start:
	@echo "Starting HTTP/2 server..."
	@cd $(http2Dir) && go run main.go

.PHONY: http3-start
http3-start:
	@echo "Starting HTTP/3 server..."
	@cd $(http3Dir) && go run main.go

# make start -j 2
.PHONY: start
start: http2-start http3-start
