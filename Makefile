RUNNABLE := todo-app

format:
	go fmt main.go

run:
	go build -o ${RUNNABLE} main.go
	./${RUNNABLE}

test:
	go test main_test.go
