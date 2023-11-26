FROM golang:1.21-bookworm
RUN apt-get update && apt-get install -y \
    bind9-utils \
    iproute2 \
    curl \
    netcat-openbsd \
    vim \
    golang
RUN mkdir -p /srv/app/
COPY app/app.go /srv/app/
COPY cmd.go /srv
COPY go.mod /srv
WORKDIR /srv
RUN go build cmd.go
CMD ["/srv/cmd"]