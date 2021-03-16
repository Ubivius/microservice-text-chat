#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/messages
curl localhost:9090/messages/1
curl localhost:9090/messages/conversation/1 # All messages from the conversation
curl localhost:9090/messages -XPOST -d '{"userid":1, "conversationid":1, "text":"This is a message"}'
curl localhost:9090/messages/1 -XDELETE

curl localhost:9090/conversations
curl localhost:9090/conversations/1
curl localhost:9090/conversations -XPOST -d '{"userid":[1, 2], "gameid":-1}'
curl localhost:9090/conversations/1 -XDELETE