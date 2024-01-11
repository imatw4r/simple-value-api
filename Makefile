install:
	go get github.com/antonfisher/nested-logrus-formatter
	go get github.com/gin-gonic/gin
	go get golang.org/x/crypto
	go get github.com/sirupsen/logrus
	go get github.com/joho/godotenv
	go get gopkg.in/yaml.v2


install-dev: install
	go get github.com/stretchr/testify


run:
	@if [ ! -f config/dev.yaml ]; then \
		cp config/example.yaml config/dev.yaml; \
	fi
	@go run ./cmd/web/


test:
	@CONFIG_PATH=../../../config/test.yaml go test ./tests/...
