install:
	go get -u github.com/antonfisher/nested-logrus-formatter
	go get -u github.com/gin-gonic/gin
	go get -u golang.org/x/crypto
	go get -u github.com/sirupsen/logrus
	go get -u github.com/joho/godotenv

run:
	@go run ./cmd/app/main.go > /dev/null
