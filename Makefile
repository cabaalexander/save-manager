.PHONY: hi
hi:
	@echo zup

.PHONY: test-backend
test-backend:
	@go test -v ./backend

.PHONY: test-frontend-watch
test-frontend-watch:
	@cd frontend && yarn test

.PHONY: test-frontend
test-frontend:
	@cd frontend && yarn test --run

.PHONY: test
test: test-backend test-frontend
