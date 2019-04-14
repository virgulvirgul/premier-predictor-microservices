FROM golang:1.11-alpine as builder
RUN apk add --no-cache ca-certificates git && \
      wget -qO/go/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && \
      chmod +x /go/bin/dep

ENV PROJECT github.com/cshep4/premier-predictor-microservices/src/chatservice
WORKDIR /go/src/$PROJECT

# restore dependencies
#COPY Gopkg.* ./
#RUN dep ensure --vendor-only -v

# <- COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

RUN go get -d -v
RUN go build -gcflags='-N -l' -o /chatservice .

#FROM alpine as release
#RUN apk add --no-cache ca-certificates
#RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
#    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
#    chmod +x /bin/grpc_health_probe
COPY --from=builder /chatservice /chatservice

ADD ./certs ./certs
ADD ./certs ./certs

EXPOSE 3000
ENTRYPOINT ["/chatservice"]
