default: test

# Load environment variables from .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Run acceptance tests
# Usage: make test [NAME="Definition Name"] [DEBUG=1]
# NAME must be a definition name (e.g., "cEdge AAA", "Transport WAN VPN") matching the name: field in gen/definitions/{feature_templates,profile_parcels,generic}/*.yaml
.PHONY: test
test:
	@echo "========================================="
	@echo "Running acceptance tests..."
	@echo "========================================="
	@if [ -z "$(SDWAN_URL)" ]; then \
		echo "SKIPPED: SDWAN_URL is not configured"; \
		echo "To enable tests, configure SDWAN_URL, SDWAN_USERNAME, and SDWAN_PASSWORD in your .env file"; \
	else \
		TEST_NAME=""; \
		if [ -n "$(NAME)" ]; then \
			MATCH=$$(grep -rl '^name: $(NAME)$$' gen/definitions/feature_templates/*.yaml gen/definitions/profile_parcels/*.yaml gen/definitions/generic/*.yaml 2>/dev/null | head -1); \
			if [ -n "$${MATCH}" ]; then \
				CAMEL=$$(echo "$(NAME)" | sed 's/-/ /g' | awk '{ result=""; for(i=1;i<=NF;i++) { word=$$i; result=result toupper(substr(word,1,1)) substr(word,2) } print result }'); \
				TEST_NAME="TestAcc.*Sdwan$${CAMEL}"; \
				echo "Resolved definition '$(NAME)' to test pattern: $${TEST_NAME}"; \
			else \
				echo "ERROR: No definition found matching '$(NAME)' in gen/definitions/"; \
				exit 1; \
			fi; \
		fi; \
		$(if $(DEBUG),echo "Debug mode enabled - logs will be written to test-output.log";) \
		if [ -n "$${TEST_NAME}" ]; then echo "Running tests matching: $${TEST_NAME}"; fi; \
		TF_ACC=1 \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $${TEST_NAME:+-run "$${TEST_NAME}"} $(TESTARGS) -count 1 -timeout 120m ./internal/provider $(if $(DEBUG),2>&1 | tee test-output.log); \
	fi

# Update all files from a single definition
# Usage: make gen NAME="Transport WAN VPN"
# NAME: The name of the definition, e.g. "Transport WAN VPN"
.PHONY: gen
gen:
	go run ./gen/generator.go "$(NAME)"
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go run gen/doc_category.go
