FROM golang:1.14.3-alpine

RUN mkdir /app
WORKDIR /app

ADD ./tableauproxy .
ADD ./templates ./templates
ADD ./ui/build ./ui/build

EXPOSE 8080/tcp

CMD ["/app/tableauproxy"]
