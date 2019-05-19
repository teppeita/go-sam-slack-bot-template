.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./remind-okan/remind-okan

build:
	GOOS=linux GOARCH=amd64 go build -o remind-okan/remind-okan ./remind-okan

package:
	sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket remind-okan --profile teppei

deploy:
	sam deploy --template-file output-template.yaml --stack-name remind-okan --capabilities CAPABILITY_IAM --profile teppei
