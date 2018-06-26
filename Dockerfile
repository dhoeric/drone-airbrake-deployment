FROM alpine:3.4

RUN apk update && \
  apk add -U --no-cache \
  ca-certificates && \
  rm -rf /var/cache/apk/*

LABEL org.label-schema.version=latest
LABEL org.label-schema.vcs-url="https://github.com/dhoeric/drone-airbrake-deployment.git"
LABEL org.label-schema.name="drone-airbrake-deployment"
LABEL org.label-schema.vendor="Eric Ho"
LABEL org.label-schema.schema-version="0.0.1"

ADD release/linux/amd64/drone-airbrake-deployment /bin/
ENTRYPOINT ["/bin/drone-airbrake-deployment"]
