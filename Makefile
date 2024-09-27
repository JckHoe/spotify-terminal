GO = $(shell which go 2>/dev/null)
COPY = $(shell which cp 2>/dev/null)


ifeq ($(GO),)
$(warning "go is not in your system PATH")
else
$(info "go found")
endif


.PHONY: all clean

all: clean app-build

app-build:
	$(GO) build -o bin/app cmd/*.go

clean:
	$(RM) -rf bin/*
