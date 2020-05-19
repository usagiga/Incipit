xxx:
	@echo "Please select optimal option."

back_build:
	@go build -o ./back/incipit ./back

back_clean:
	@rm -f ./back/incipit

back_run:
	@go run ./back

back_test:
	@go test -v "./back/..."
