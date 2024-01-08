# Redirects

This package provides a way to handle redirects in a Go application using the Gorilla Mux router. It allows you to define redirects in a YAML file and apply them to your router.

## Usage

First, define your redirects in a YAML file. Each redirect should have a `from` path, a `to` path, and a `status` code. The `from` path can include wildcards, which are represented by `*`.

```yaml
redirects:
  - from: '/old-path/(.*)'
    to: '/new-path?new=$1'
    status: 301
  - from: /another-old-path
    to: /another-new-path
    status: 301
  - from: /everything/(.*)/(.*)/(.*)/(.*)
    to: /?one=$1&two=$2&three=$3&four=$4
    status: 301
```
