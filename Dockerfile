FROM golang:1.14.3-alpine

RUN mkdir /app
WORKDIR /app

ADD ./tableauproxy /app/
ADD ./templates /app/templates
ADD ./ui/build /app/ui/build

EXPOSE 8080/tcp

CMD ["/app/tableauproxy"]
