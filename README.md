# http-redirect
Redirect HTTP to HTTP(S), redirect HTTP ports, when a service only listens on one port you can redirect to another

This is useful for tools like Gogs (https://gogs.io) that only listen on one port (e.g. you configure it to listen on 443) but don't provide graceful redirection, obviously you can use a proxy, but for simple use-cases its a bit overkill - hence this simple tool.

As usual to get going, something like the following..

```go get github.com/penolver/http-redirect
cd src/github.com/penolver/http-redirect
go build redirector.go```
