FROM golang:1.22.4-alpine
COPY go.mod go.sum ./

RUN go mod download
COPY *.go ./
COPY pkg/*.go ./pkg

#RUN go build -o /go-docker-app
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-docker-app

CMD [ "/go-docker-app" ]