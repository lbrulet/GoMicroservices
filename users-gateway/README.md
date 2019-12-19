# Users Service

This is the Users service

Generated with

```
micro new github.com/lbrulet/GoMicroservices/users-gateway --namespace=go.micro --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.users
- Type: api
- Alias: users

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./users-api
```

Build a docker image
```
make docker
```