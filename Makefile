VERSION := $(shell git describe --exact-match --tags 2> /dev/null || git rev-parse --short HEAD)
default:
	go build -o ./bin/terraform-provider-vgs_$(VERSION)

all:
	GOOS=linux GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(VERSION)_linux_amd64
	GOOS=linux GOARCH=386 go build -o ./bin/terraform-provider-vgs_$(VERSION)_linux_386
	GOOS=linux GOARCH=arm go build -o ./bin/terraform-provider-vgs_$(VERSION)_linux_arm
	GOOS=darwin GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(VERSION)_darwin_amd64
	GOOS=darwin GOARCH=386 go build -o ./bin/terraform-provider-vgs_$(VERSION)_darwin_386
	GOOS=darwin GOARCH=arm go build -o ./bin/terraform-provider-vgs_$(VERSION)_darwin_arm
	GOOS=windows GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(VERSION)_windows_amd64
	GOOS=windows GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(VERSION)_windows_386
	GOOS=windows GOARCH=arm go build -o ./bin/terraform-provider-vgs_$(VERSION)_windows_arm
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(VERSION)_freebsd_amd64
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(VERSION)_freebsd_386
	GOOS=freebsd GOARCH=arm go build -o ./bin/terraform-provider-vgs_$(VERSION)_freebsd_arm
