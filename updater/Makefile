clean:
	go clean
	rm -f wh3_updater
	rm -f wh3_updater.zip

build:
	GOOS=linux go build -o wh3_updater .

install: clean build
	zip server.zip wh3_updater *.json

deploy: install
	aws lambda update-function-code \
		--function-name arn:aws:lambda:us-west-2:XXXXXXX \
		--zip-file fileb://wh3_updater.zip
	rm -f wh3_updater.zip