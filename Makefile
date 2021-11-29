VERSION := $(shell git describe --exact-match --tags 2> /dev/null || git rev-parse --short HEAD)
default:
	go build -o ./bin/terraform-provider-vgs_$(VERSION)

install-local:
	go build -o ~/.terraform.d/plugins/terraform-provider-vgs_$(VERSION)

all:
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux_amd64/terraform-provider-vgs_$(VERSION)
	GOOS=linux GOARCH=386 go build -o ./bin/linux_386/terraform-provider-vgs_$(VERSION)
	GOOS=linux GOARCH=arm go build -o ./bin/linux_arm/terraform-provider-vgs_$(VERSION)
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin_amd64/terraform-provider-vgs_$(VERSION)
	GOOS=darwin GOARCH=386 go build -o ./bin/darwin_386/terraform-provider-vgs_$(VERSION)
	GOOS=darwin GOARCH=arm go build -o ./bin/darwin_arm/terraform-provider-vgs_$(VERSION)
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows_amd64/terraform-provider-vgs_$(VERSION).exe
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows_386/terraform-provider-vgs_$(VERSION).exe
	GOOS=windows GOARCH=arm go build -o ./bin/windows_arm/terraform-provider-vgs_$(VERSION).exe
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/freebsd_amd64/terraform-provider-vgs_$(VERSION)
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/freebsd_386/terraform-provider-vgs_$(VERSION)
	GOOS=freebsd GOARCH=arm go build -o ./bin/freebsd_arm/terraform-provider-vgs_$(VERSION)
