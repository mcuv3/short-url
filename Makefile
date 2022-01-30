


short-create:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o create ./lambdas/create 
	zip create.zip create

short:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o short ./lambdas/short 
	zip short.zip short