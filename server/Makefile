clean:
	go clean
	rm -f server
	rm -f server.zip

build:
	GOOS=linux go build -o server .

install: clean build
	zip server.zip server

deploy: install
	aws lambda update-function-code \
		--function-name arn:aws:lambda:us-west-2:202214144554:function:wh3_server \
		--zip-file fileb://server.zip
	rm -f server.zip