FROM alpine:3.12.0

COPY service .
COPY build/*.dart.js dist/js/
COPY views views

RUN mkdir -p /views/_shared

EXPOSE 8096

ENTRYPOINT [ "./service" ]
