- Session is a network connection between a client application and the server
- User represents a human being who connects to the server with a session.
- Topic is a named communication channel which routes content between sessions.
- Clients such as mobile or web applications create sessions by connecting to the server over a websocket

``` bash
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
oapi-codegen --output-config --old-config-style  --package=ports --generate="types,gorilla" openapi.yaml
oapi-codegen --config openapi/oapi_server.yaml openapi/openapi.yaml > ports/openapi_server.gen.go
oapi-codegen --config openapi/oapi_types.yaml openapi/openapi.yaml > dto/openapi_types.gen.go
```

Use of Pub/Sub

A channel is a representation of a line of communication for users. eg a group chat, one to one chat
A membership is a way of linking users to a channel i.e being a member of a channel is denoted by having a membership. A channel's memberships indicates it's members/users. Each membership should have a unique ID
A session is a connection between a client (user application) and the server.

Pubsub is used in order to route messages where: 
- Each user's membership is a topic
- Any active user session is a subscriber to the topic(s). An inactive/closed session should have it's subscription removed from the topics.

A message sent from one user is first saved in the db broadcasted to all the members of the channel i.e memberships (sent to all the topics). An active session (subscription) receives the message and sends it to it's connected client