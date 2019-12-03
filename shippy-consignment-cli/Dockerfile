FROM golang:latest as builder

ARG SSH_PRIVATE_KEY

WORKDIR /app/shippy-consignment-cli

RUN mkdir -p ~/.ssh && umask 0077 && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa \
    && git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com/" \
    && ssh-keyscan gitlab.com >> ~/.ssh/known_hosts

COPY . .

RUN GO111MODULE=on go mod download

RUN echo "Building. This may take some time..."

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-consignment-cli

RUN echo "Creating container"

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir -p /app
WORKDIR /app

ADD consignment.json /app/consignment.json
COPY --from=builder /app/shippy-consignment-cli .

CMD ["./shippy-consignment-cli"]
