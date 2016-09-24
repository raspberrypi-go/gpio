PWD := $(shell pwd)
PKG_CUT := github.com/raspberrypi-go
PKG := $(PKG_CUT)/gpio
APP := watch

REMOTE_ADDR := 192.168.5.28
REMOTE_GOPATH := /opt/go


push:
	rsync -vaz --delete --exclude=.git ./ root@$(REMOTE_ADDR):$(REMOTE_GOPATH)/src/$(PKG)
