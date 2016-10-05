FROM alpine

RUN apk --update add ca-certificates

COPY ./s3-club-7 /s3-club-7

EXPOSE 8000
ENTRYPOINT ["/s3-club-7"]
