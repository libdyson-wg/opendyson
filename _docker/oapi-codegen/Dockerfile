FROM docker.io/golang:1.20-alpine3.18

RUN apk upgrade && apk add npm bash && npm i -g @redocly/cli@latest

RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

COPY root/usr/local/bin/docker-run.sh /usr/local/bin/docker-run.sh
RUN chmod +x /usr/local/bin/docker-run.sh

CMD docker-run.sh
