RUNNABLE := todo-app

format:
	templ fmt .
	go fmt *.go

generate:
	templ generate

run: generate
	go build -o ${RUNNABLE} *.go
	./${RUNNABLE}

test:
	go test main_test.go

clean:
	find . -name "*_templ.go" -type f -print0 | xargs -t -0 -I _ rm _
