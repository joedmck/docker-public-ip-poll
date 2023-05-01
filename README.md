# Docker Public IP Poll

![CI/CD Status](https://img.shields.io/github/actions/workflow/status/joedmck/docker-public-ip-poll/build-on-push.yml)

Docker image that retrieves your public IP regularly and prints to stdout.

## Configuration

Conigured via environment variables.

Variable Name | Default Value
---|---
INTERVAL | 1m
ENDPOINT | https://checkip.amazonaws.com

## Getting Started

With standard configuration:
```bash
docker run ghcr.io/joedmck/docker-public-ip-poll:main
```

Changing interval:
```bash
docker run -e INTERVAL="5s" ghcr.io/joedmck/docker-public-ip-poll:main
```

Changing endpoint:
```bash
docker run -e ENDPOINT="https://api64.ipify.org/" ghcr.io/joedmck/docker-public-ip-poll:main
```
