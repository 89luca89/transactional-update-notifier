PROGRAM = transactional-update-notifier
VERSION = 1.1.0

DESTDIR ?=
PREFIX = /usr
BINDIR := $(PREFIX)/bin
DBUSCONFDIR := $(PREFIX)/share/dbus-1/system.d/
USERUNITDIR := $(PREFIX)/lib/systemd/user
USERPRESETDIR := $(PREFIX)/lib/systemd/user-preset
BUILDDIR = build

WITH_SYSTEMD_PRESET = 1

GO = go
GOFLAGS ?=
LDFLAGS ?= -ldflags "-X main.Version=$(VERSION)"
CGOFLAGS ?= CGO_ENABLED=0

all: $(PROGRAM) systemd.service dbus.policy

$(PROGRAM):
	mkdir -p $(BUILDDIR)
	$(CGOFLAGS) $(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BUILDDIR)/$(PROGRAM)

systemd.service:
	sed -e 's|@bindir@|$(BINDIR)|' $(PROGRAM).service.in > $(BUILDDIR)/$(PROGRAM).service

dbus.policy:
	cp tukit-notifier.conf $(BUILDDIR)/$(PROGRAM).conf

install:
	install -m 0755 -D $(BUILDDIR)/$(PROGRAM) $(DESTDIR)$(BINDIR)/$(PROGRAM)
	install -m 0644 -D $(BUILDDIR)/$(PROGRAM).service $(DESTDIR)$(USERUNITDIR)/$(PROGRAM).service
	install -m 0644 -D $(BUILDDIR)/$(PROGRAM).conf $(DESTDIR)$(DBUSCONFDIR)/$(PROGRAM).conf

ifeq ($(WITH_SYSTEMD_PRESET), 1)
	install -m 0644 -D 96-$(PROGRAM).preset $(DESTDIR)$(USERPRESETDIR)/96-$(PROGRAM).preset
endif

clean:
	rm -r $(BUILDDIR)
