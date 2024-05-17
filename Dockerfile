FROM golang:1.22
ADD . /usr/bin/post
ADD ./db /usr/bin/post/db

RUN ls
COPY . /usr/bin/post
COPY db /usr/bin/post
RUN ls
WORKDIR /usr/bin/post/internal/post
RUN echo $GOPATH
RUN go version
RUN go mod download
RUN go build
RUN go install -v ./...
ENTRYPOINT post
EXPOSE 8080

