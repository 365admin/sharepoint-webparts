FROM golang:1.20-buster
WORKDIR /app.
COPY . .
RUN go install

CMD [ "sleep","infinity"]