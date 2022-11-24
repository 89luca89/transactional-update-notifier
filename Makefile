PROGRAM = transactional-update-notifier
VERSION = 1.0.0

DESTDIR ?=
PREFIX = /usr
BINDIR := $(PREFIX)/bin
USERUNITDIR := $(PREFIX)/lib/systemd/user
USERPRESETDIR := $(PREFIX)/lib/systemd/user-preset
BUILDDIR = build

WITH_SYSTEMD_PRESET = 1

GO = go
GOFLAGS ?= 
LDFLAGS ?= -ldflags "-X main.Version=$(VERSION)"
CGOFLAGS ?= CGO_ENABLED=0

all: $(PROGRAM) systemd.service

$(PROGRAM):
	mkdir $(BUILDDIR)
	$(CGOFLAGS) $(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BUILDDIR)/$(PROGRAM)

systemd.service:
	sed -e 's|@bindir@|$(BINDIR)|' $(PROGRAM).service.in > $(BUILDDIR)/$(PROGRAM).service

install: 
	install -m 0755 -D $(BUILDDIR)/$(PROGRAM) $(DESTDIR)$(BINDIR)/$(PROGRAM)
	install -m 0644 -D $(BUILDDIR)/$(PROGRAM).service $(DESTDIR)$(USERUNITDIR)/$(PROGRAM).service

ifeq ($(WITH_SYSTEMD_PRESET), 1)
	install -m 0644 -D 96-$(PROGRAM).preset $(DESTDIR)$(USERPRESETDIR)/96-$(PROGRAM).preset
endif

clean:
	rm -r $(BUILDDIR)
