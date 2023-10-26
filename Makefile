RUNNABLE := todo-app

format:
	templ fmt .
	go fmt *.go

run:
	go build -o ${RUNNABLE} *.go
	./${RUNNABLE}

test:
	go test main_test.go
