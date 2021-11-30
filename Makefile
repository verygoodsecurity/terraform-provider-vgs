RELEASE := $(shell git describe --tags $(git rev-list --branches=main --tags --max-count=1))
PRERELEASE := $(shell git describe --tags $(git rev-list --branches=main --tags --max-count=1) | awk 'BEGIN{FS=OFS="."} {$$3+=1} 1')
default:
	go build -o ./bin/terraform-provider-vgs_$(RELEASE)

install-local:
	go build -o ~/.terraform.d/plugins/terraform-provider-vgs_$(PRERELEASE)

all:
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux_amd64/terraform-provider-vgs_$(RELEASE)
	GOOS=linux GOARCH=386 go build -o ./bin/linux_386/terraform-provider-vgs_$(RELEASE)
	GOOS=linux GOARCH=arm go build -o ./bin/linux_arm/terraform-provider-vgs_$(RELEASE)
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin_amd64/terraform-provider-vgs_$(RELEASE)
	GOOS=darwin GOARCH=386 go build -o ./bin/darwin_386/terraform-provider-vgs_$(RELEASE)
	GOOS=darwin GOARCH=arm go build -o ./bin/darwin_arm/terraform-provider-vgs_$(RELEASE)
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows_amd64/terraform-provider-vgs_$(RELEASE).exe
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows_386/terraform-provider-vgs_$(RELEASE).exe
	GOOS=windows GOARCH=arm go build -o ./bin/windows_arm/terraform-provider-vgs_$(RELEASE).exe
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/freebsd_amd64/terraform-provider-vgs_$(RELEASE)
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/freebsd_386/terraform-provider-vgs_$(RELEASE)
	GOOS=freebsd GOARCH=arm go build -o ./bin/freebsd_arm/terraform-provider-vgs_$(RELEASE)
