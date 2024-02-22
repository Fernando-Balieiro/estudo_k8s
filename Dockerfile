FROM golang:1.22.0-alpine

COPY . .
RUN go build -o server .
CMD [ "./server" ]
