FROM golang

RUN mkdir -p /go/src//Users/mikehsu/Desktop/go_project/go-kit-reddit-demo/internal

ADD . /go/src//Users/mikehsu/Desktop/go_project/go-kit-reddit-demo/internal

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run /Users/mikehsu/Desktop/go_project/go-kit-reddit-demo/internal/auth/cmd  -watch /Users/mikehsu/Desktop/go_project/go-kit-reddit-demo/internal/auth
