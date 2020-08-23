.PHONY: help
help:
	@echo "Make Targets:"
	@echo "- init:     initialize terraform"
	@echo "- plan:     plan terraform"
	@echo "- apply:    apply terraform"

.PHONY: init
init:
	terraform init terraform/

.PHONY: plan
plan:
	terraform plan -out terraform/plan.tfplan terraform/

.PHONY: apply
tf-apply: plan
	terraform apply terraform/plan.tfplan

.PHONY: clean
clean:
	rm -rf terraform/plan.tfplan

.PHONY: docker
docker:
	cd .. && docker build -t url-shortener -f .\url-shortener\Dockerfile .

