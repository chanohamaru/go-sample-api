format:
	@find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/chanohamaru/go-sample-api"
verify:
	@staticcheck ./... && go vet ./...
serve:
	@docker-compose -f build/docker-compose.yml up