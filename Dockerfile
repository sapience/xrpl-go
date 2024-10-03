FROM golang:1.22 AS install

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

FROM install AS lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN make lint

FROM lint AS test
RUN make test