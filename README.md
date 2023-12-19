# Chat server for 

It's simple chat serve based on websocket connection

After connection to websocket connection we can start read from socket and write to the socket. 
All users have nickname. To set nickname you shoud send string after connect to the socket

All messages it's simple struct of Message that have 2 fields: nickname and message payload

Example object of message:
```
{
  "text": "some text",
  "sender": "nickname"
}
```

After client connect to the ws all ws users get notification about new user in chat room.

## Routes 

${{host}}/ws - work with websocket
