FROM golang:1.14.3-alpine

ADD ./tableauproxy .
ADD ./ui/build ./ui/build

EXPOSE 8080/tcp

CMD ["./tableauproxy"]
