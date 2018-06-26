# drone-airbrake-deployment
[![Go Report Card](https://goreportcard.com/badge/github.com/dhoeric/drone-airbrake-deployment)](https://goreportcard.com/report/github.com/dhoeric/drone-airbrake-deployment)

Drone plugin to notify airbrake when deployment is finished. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
make install
make build
```

## Docker

Build the docker image with the following commands:

```
make linux_amd64 docker_image docker_deploy tag=X.X.X
```

## Usage

Execute from the working directory:

```sh
docker run --rm \
    -e AIRBRAKE_PROJECT_ID=xxx \
    -e AIRBRAKE_PROJECT_KEY=xxx \
    -e AIRBRAKE_ENVIRONMENT=staging \
    -e DRONE_COMMIT_AUTHOR=dhoeric \
    -e DRONE_COMMIT_SHA=xxx \
    -e DRONE_REPO_LINK=https://github.com/dhoeric/drone-airbrake-deployment \
  dhoeric/drone-airbrake-deployment
```
