
default: build

.PHONY: lint
lint:
	gometalinter ./...

.PHONY: clean
clean:
	@rm -f prof pq.test pq bench.txt

.PHONY: benchmark
benchmark: clean
	go test -c
	./pq.test -test.bench=. -test.count=5 >bench.txt
	@cat bench.txt

.PHONY: mem-profile
mem-profile: clean
	go test -run=XXX -bench=. -memprofile=prof kkn.fi/pq
	go tool pprof pq.test prof

.PHONY: cpu-profile
cpu-profile: clean
	go test -run=XXX -bench=. -cpuprofile=prof kkn.fi/pq
	go tool pprof pq.test prof

.PHONY: build
build:
	go build kkn.fi/pq

.PHONY: test
test:
	go test kkn.fi/pq
