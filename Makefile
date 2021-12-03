RELEASE := $(shell git describe --tags $(git rev-list --branches=main --tags --max-count=1))
PRERELEASE := $(shell git describe --tags $(git rev-list --branches=main --tags --max-count=1) | awk 'BEGIN{FS=OFS="."} {$$3+=1} 1')
default:
	go build -o ./bin/terraform-provider-vgs_$(RELEASE)

install-local:	
	mkdir -p ~/.terraform.d/plugins/local.terraform.com/user/vgs/$(subst v,,$(PRERELEASE))/darwin_amd64
	go build -o ~/.terraform.d/plugins/local.terraform.com/user/vgs/$(subst v,,$(PRERELEASE))/darwin_amd64/terraform-provider-vgs_$(PRERELEASE)

test:
	echo "$(shell sha256sum release/terraform-provider-vgs_v0.1.1_linux_386.zip) 111"

all:
	rm -rf ./release && mkdir -p ./release
	GOOS=linux GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(RELEASE)_linux_amd64/terraform-provider-vgs_$(RELEASE) && zip -r ./release/terraform-provider-vgs_$(RELEASE)_linux_amd64.zip ./bin/terraform-provider-vgs_$(RELEASE)_linux_amd64
	GOOS=linux GOARCH=386 go build -o ./bin/terraform-provider-vgs_$(RELEASE)_linux_386/terraform-provider-vgs_$(RELEASE) && zip -r ./release/terraform-provider-vgs_$(RELEASE)_linux_386.zip ./bin/terraform-provider-vgs_$(RELEASE)_linux_386
	GOOS=linux GOARCH=arm go build -o ./bin/terraform-provider-vgs_$(RELEASE)_linux_arm/terraform-provider-vgs_$(RELEASE) && zip -r ./release/terraform-provider-vgs_$(RELEASE)_linux_arm.zip ./bin/terraform-provider-vgs_$(RELEASE)_linux_arm

	GOOS=darwin GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(RELEASE)_darwin_amd64/terraform-provider-vgs_$(RELEASE) && zip -r ./release/terraform-provider-vgs_$(RELEASE)_darwin_amd64.zip ./bin/terraform-provider-vgs_$(RELEASE)_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o ./bin/terraform-provider-vgs_$(RELEASE)_darwin_arm64/terraform-provider-vgs_$(RELEASE) && zip -r ./release/terraform-provider-vgs_$(RELEASE)_darwin_arm.zip ./bin/terraform-provider-vgs_$(RELEASE)_darwin_arm64

	GOOS=windows GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(RELEASE)_windows_amd64/terraform-provider-vgs_$(RELEASE).exe && zip -r ./release/terraform-provider-vgs_$(RELEASE)_windows_amd64.zip ./bin/terraform-provider-vgs_$(RELEASE)_windows_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/terraform-provider-vgs_$(RELEASE)_windows_386/terraform-provider-vgs_$(RELEASE).exe && zip -r ./release/terraform-provider-vgs_$(RELEASE)_windows_386.zip ./bin/terraform-provider-vgs_$(RELEASE)_windows_386
	GOOS=windows GOARCH=arm go build -o ./bin/terraform-provider-vgs_$(RELEASE)_windows_arm/terraform-provider-vgs_$(RELEASE).exe && zip -r ./release/terraform-provider-vgs_$(RELEASE)_windows_arm.zip ./bin/terraform-provider-vgs_$(RELEASE)_windows_arm

	GOOS=freebsd GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(RELEASE)_freebsd_amd64/terraform-provider-vgs_$(RELEASE) && zip -r ./release/terraform-provider-vgs_$(RELEASE)_freebsd_amd64.zip ./bin/terraform-provider-vgs_$(RELEASE)_freebsd_amd64
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/terraform-provider-vgs_$(RELEASE)_freebsd_386/terraform-provider-vgs_$(RELEASE) && zip -r ./release/terraform-provider-vgs_$(RELEASE)_freebsd_386.zip ./bin/terraform-provider-vgs_$(RELEASE)_freebsd_386
	GOOS=freebsd GOARCH=arm go build -o ./bin/terraform-provider-vgs_$(RELEASE)_freebsd_arm/terraform-provider-vgs_$(RELEASE) && zip -r ./release/terraform-provider-vgs_$(RELEASE)_freebsd_arm.zip ./bin/terraform-provider-vgs_$(RELEASE)_freebsd_arm

checksums:
	rm -f ./release/terraform-provider-vgs_$(RELEASE)_sha256sums && touch ./release/terraform-provider-vgs_$(RELEASE)_sha256sums
	find ./release -type f ! -name '*sha256sums' -exec basename {} \; | xargs -I '{}' bash -c 'echo "$$(sha256sum ./release/{}) {}" >> ./release/terraform-provider-vgs_$(RELEASE)_sha256sums'

