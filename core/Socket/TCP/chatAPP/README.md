# Basic chat application

> https://www.thepolyglotdeveloper.com/2017/05/network-sockets-with-the-go-programming-language/

This chat application will have a server for listening and routing client communications and a client for sending and 
receiving messages from the server.

When we start the server a management process will be started. This management service will keep track of connected 
clients and queue messages to each of the clients. For the server, we will listen for connections and when established, 
the management process will keep track of them in addition to a new send and receive process being started. The server 
will have one management process and then one send and receive process for every connection.

When starting the application as a client, things are handled a bit differently. There will be no management process 
because the client should not be managing connections. Instead a process for receiving data will be started and the 
application will allow for sending of data.
