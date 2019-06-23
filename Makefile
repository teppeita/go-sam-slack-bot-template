.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./template/template

build:
	GOOS=linux GOARCH=amd64 go build -o template/template ./template

package:
	sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket template --profile teppei

deploy:
	sam deploy --template-file output-template.yaml --stack-name template --capabilities CAPABILITY_IAM --profile teppei
