FROM golang
WORKDIR /go/src/hello-world-api
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...
RUN go build
CMD ["./hello-world-api"]