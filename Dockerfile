FROM golang:latest

EXPOSE 8080

WORKDIR /var/lib/lemmas

COPY . .

RUN go build -v -a

CMD /var/lib/lemmas/lemmas -lemmas /var/lib/lemmas/data/lemmas.db
