GOARGS=-a -installsuffix cgo -x
GOENV=CGO_ENABLED=0
GOOS=linux
BINDIR=/usr/local/sbin
EXECUTABLE=s3-club-7
CONTAINER_TAG=quay.io/financialtimes/s3-club-7:latest

all: main

main:
	$(GOENV) GOOS=$(GOOS) go build $(GOARGS)

install: all
	install -d $(BINDIR)
	install -s -m 0750 -o $(USER) $(EXECUTABLE) $(BINDIR)/$(EXECUTABLE)

build: all
	docker build -t $(CONTAINER_TAG) .

push:
	docker push $(CONTAINER_TAG)

uninstall:
	rm -rfv $(CONF)
	rm -v $(BINDIR)/$(EXECUTABLE)

clean:
	-rm -v $(EXECUTABLE)

strip:
	strip -v $(EXECUTABLE)

dist: clean main strip build push
