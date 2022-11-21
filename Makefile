PROGRAM = transactional-update-notifier
VERSION = 1.0.0

PREFIX = /usr
BINDIR := $(PREFIX)/bin
USERUNITDIR := $(PREFIX)/lib/systemd/user
USERPRESETDIR := $(PREFIX)/lib/systemd/user-preset
BUILDDIR = build

WITH_SYSTEMD_PRESET = yes

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
	install -m 0755 -D $(BUILDDIR)/$(PROGRAM) $(BINDIR)/$(PROGRAM)
	install -m 0644 -D $(BUILDDIR)/$(PROGRAM).service $(USERUNITDIR)/$(PROGRAM).service

ifeq ($(WITH_SYSTEMD_PRESET), yes)
	install -m 0644 -D 96-$(PROGRAM).preset $(USERPRESETDIR)/96-$(PROGRAM).preset
endif

clean:
	rm -r $(BUILDDIR)
