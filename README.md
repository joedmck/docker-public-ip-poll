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

```bash
docker run ghcr.io/joedmck/docker-public-ip-poll:main
```
