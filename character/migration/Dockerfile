FROM alpine

COPY ./dist/bin/migration /go/migration
COPY ./postgresql /go/postgresql

ENTRYPOINT /go/migration
