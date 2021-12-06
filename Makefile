RELEASE := $(shell git describe --tags $(git rev-list --branches=main --tags --max-count=1))
PRERELEASE := $(shell git describe --tags $(git rev-list --branches=main --tags --max-count=1) | awk 'BEGIN{FS=OFS="."} {$$3+=1} 1')

build:
	go build -o ./bin/terraform-provider-vgs_$(PRERELEASE)

install-local:
	go build -o ~/.terraform.d/plugins/terraform-provider-vgs_$(PRERELEASE)

test:
	
