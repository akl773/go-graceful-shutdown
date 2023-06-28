# go-graceful-shutdown
The Graceful library provides a convenient way to create and manage graceful shutdowns for HTTP servers in Go. It allows you to handle interrupt signals and gracefully shut down the server, ensuring that all active connections are properly closed before exiting.

## Installation
Run the following command to install the package:

```
go get github.com/akl773/go-graceful-shutdown
```

Make sure to include this import statement in your Go source files.


## Usage
The Graceful library provides a GracefulServer struct, which represents an HTTP server with graceful shutdown capabilities. Here's how you can use it:

```
srv := graceful.NewGracefulServer(":8080", handler)
err := srv.Start()
if err != nil {
    // Handle server startup error
}

```
