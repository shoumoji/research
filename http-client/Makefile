COUNT := 100
FORMAT := csv

.SILENT:
.PHONY: run
run:
	go run $(CURDIR)/*.go --count $(COUNT) --format $(FORMAT)

.PHONY: run-http2
run-http2:
	@for i in {1..$(COUNT)}; do go run $(CURDIR)/*.go --count 1 --format json --http2; done

.PHONY: run-http3
run-http3:
	@for i in {1..$(COUNT)}; do go run $(CURDIR)/*.go --count 1 --format json --http3; done
