PKGNAME    = sup
DESTDIR   ?=
PREFIX    ?= /usr
BINDIR     = $(PREFIX)/bin
DATADIR    = $(PREFIX)/share/$(PKGNAME)
THEMEDIR   = $(DATADIR)/themes

GOPROJROOT  = $(GOSRC)/$(PROJREPO)

GOLDFLAGS   = -ldflags "-s -w -X github.com/DataDrake/sup/themes.ThemeDir=$(THEMEDIR)"
GOCC        = go
GOFMT       = $(GOCC) fmt -x
GOGET       = $(GOCC) get $(GOLDFLAGS)
GOBUILD     = $(GOCC) build -v $(GOLDFLAGS) $(GOTAGS)
GOTEST      = $(GOCC) test
GOVET       = $(GOCC) vet
GOINSTALL   = $(GOCC) install $(GOLDFLAGS)

include Makefile.waterlog

GOLINT = golint -set_exit_status

all: build

build:
	@$(call stage,BUILD)
	@$(GOBUILD)
	@$(call pass,BUILD)

test: build
	@$(call stage,TEST)
	@$(GOTEST) ./...
	@$(call pass,TEST)

validate:
	@$(call stage,FORMAT)
	@$(GOFMT) ./...
	@$(call pass,FORMAT)
	@$(call stage,VET)
	@$(call task,Running 'go vet'...)
	@$(GOVET) ./...
	@$(call pass,VET)
	@$(call stage,LINT)
	@$(call task,Running 'golint'...)
	@$(GOLINT) ./...
	@$(call pass,LINT)

install:
	@$(call stage,INSTALL)
	install -Dm 00755 $(PKGNAME) $(DESTDIR)$(BINDIR)/$(PKGNAME)
	install -Dm 00644 data/bash.sh $(DESTDIR)$(DATADIR)/bash.sh
	install -Dm 00644 data/sup.zsh $(DESTDIR)$(DATADIR)/sup.zsh
	install -Dm 00644 data/themes/default.json $(DESTDIR)$(THEMEDIR)/default.json
	install -Dm 00644 data/themes/warm.json $(DESTDIR)$(THEMEDIR)/warm.json
	@$(call pass,INSTALL)

uninstall:
	@$(call stage,UNINSTALL)
	rm -f $(DESTDIR)$(BINDIR)/$(PKGNAME)
	rm -rf $(DESTDIR)$(DATADIR)
	@$(call pass,UNINSTALL)

clean:
	@$(call stage,CLEAN)
	@$(call task,Removing executable...)
	@rm $(PKGNAME)
	@$(call pass,CLEAN)
