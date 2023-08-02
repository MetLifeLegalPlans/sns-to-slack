ZIP := bundle.zip
EXE := notifier
IMG := sns-to-slack

GOFLAGS := -tags lambda.norpc

bundle: notifier
	zip $(ZIP) $(EXE)

notifier:
	go build -tags lambda.norpc -o notifier

start: rie docker
	docker run -v ~/.aws-lambda-rie:/aws-lambda -p 8080:8080 --entrypoint /aws-lambda/aws-lambda-rie $(IMG):latest /main

docker:
	docker build -t $(IMG):latest .

rie:
	mkdir -p ~/.aws-lambda-rie
	curl -Lo ~/.aws-lambda-rie/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie
	chmod +x ~/.aws-lambda-rie/aws-lambda-rie
