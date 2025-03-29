FROM golang:1.24

WORKDIR /app
RUN apt-get update && apt-get install -y gcc libc-dev
COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY env.json ./

RUN CGO_ENABLED=1 GOOS=linux go build -v -o /docker-gs-ping

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]