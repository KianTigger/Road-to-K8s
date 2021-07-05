FROM golang:1.10
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["/app/main"]