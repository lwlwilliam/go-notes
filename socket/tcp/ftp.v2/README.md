### Application-Leval Protocol

> https://github.com/tumregels/Network-Programming-with-Go/blob/master/applevelprotocols/simple_example.md

This example deals with a directory browsing protocol - basically a stripped down version of FTP, but without even 
the file transfer part.

server.go

```
$ go run server.go
```

client.go

```
$ go run client.go
>>> 
```

Available commands:

*   cd
*   clear
*   help
*   ls
*   pwd
*   quit
