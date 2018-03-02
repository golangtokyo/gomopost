
OUTDIR = $(CURDIR)/android/app/libs
AAR    = $(OUTDIR)/gomopost.aar

all: cmd bind

.PHONY: cmd
cmd:
	cd $(CURDIR)/cmd/gomopost && go build

.PHONY: bind
bind: $(AAR)

$(AAR): $(CURDIR)/gomopost.go
	mkdir -p $(OUTDIR)
	gomobile bind -o $(AAR)

.PHONY: lint
lint:
	gometalinter $(CURDIR)/... --exclude='cmd/server'

clean:
	rm -rf $(OUTDIR)/*
