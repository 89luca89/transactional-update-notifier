VERSION ?= "1.0.0"

all:
	CGO_ENABLED=0 go build \
		-ldflags  "-X main.Version=${VERSION}" \
		-o build/transactional-update-notifier
	cp transactional-update-notifier@.service build/

install:
	cp build/transactional-update-notifier /usr/bin/
	cp transactional-update-notifier@.service /etc/systemd/system
