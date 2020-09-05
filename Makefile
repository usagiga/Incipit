FRONT_DIR=./front
BACK_DIR=./back

xxx:
	@echo "Please select optimal option."

front_build:
	@cd $(FRONT_DIR) && npm run build

front_clean:
	@cd $(FRONT_DIR) && rm -rf ./dist

front_run:
	@cd $(FRONT_DIR) && npm run dev

back_build:
	@cd $(BACK_DIR) && go build -o incipit .

back_clean:
	@cd $(BACK_DIR) && rm -f ./incipit

back_run:
	@cd $(BACK_DIR) && go run .

back_run_air:
	@cd $(BACK_DIR) && air

back_test:
	@cd $(BACK_DIR) && go test -v "./..."

back_test_air:
	@cd $(BACK_DIR) && air -c .air-test.toml
