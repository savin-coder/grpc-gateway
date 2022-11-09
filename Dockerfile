FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go .
COPY generate.sh .
COPY proto proto
RUN chmod u+x generate.sh
RUN mkdir generated
RUN go install "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
RUN go install "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
RUN go install "google.golang.org/protobuf/cmd/protoc-gen-go"
RUN PATH="$PATH:$(go env GOPATH)/bin"
RUN export PATH
RUN apk add protoc
RUN ./generate.sh
RUN go build

FROM alpine
WORKDIR /grpc
COPY --from=builder /app/ecards-grpc-gateway .
CMD ["./ecards-grpc-gateway"]