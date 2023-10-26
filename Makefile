RUNNABLE := todo-app

format:
	go fmt *.go

run:
	go build -o ${RUNNABLE} *.go
	./${RUNNABLE}

test:
	go test main_test.go
