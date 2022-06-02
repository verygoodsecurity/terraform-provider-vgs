include common/make/common.mk

RELEASE := $(shell git describe --tags $(git rev-list --branches=main --tags --max-count=1))
PRERELEASE := $(shell git describe --tags $(git rev-list --branches=main --tags --max-count=1) | awk 'BEGIN{FS=OFS="."} {$$3+=1} 1')

build:
	go build -o ./bin/terraform-provider-vgs_$(PRERELEASE)

install-local:
	mkdir -p ~/.terraform.d/plugins/local.terraform.com/user/vgs/$(subst v,,$(PRERELEASE))/darwin_amd64
	# 0.13+
	go build -o ~/.terraform.d/plugins/local.terraform.com/user/vgs/$(subst v,,$(PRERELEASE))/darwin_amd64/terraform-provider-vgs_$(PRERELEASE)
	# 0.12
	go build -o ~/.terraform.d/plugins/terraform-provider-vgs_$(PRERELEASE)

testacc:
	TF_ACC=true go test -v ./...

lint: go-lint ## Lint all go files
test: go-test ## Run all un-tagged go tests
integration: go-test-integration ## Run all integrations tests, concourse helper
install-tools: go-install-tools alderson-install-tools ## Install all required tools
init-ci: init-github-actions ## Creates push_to_pr.yml and merged.yml workflows for github. This will overwrite current workflows by the same names
update-ci: init-github-actions ## Creates push_to_pr.yml and merged.yml workflows for github. This will overwrite current workflows by the same names
docker-build: docker-build-app # Build all docker images
