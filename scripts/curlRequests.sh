#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/messages/a2181017-5c53-422b-b6bc-036b27c04fc8
curl localhost:9090/messages/conversation/a2181017-5c53-422b-b6bc-036b27c04fc8 # All messages from the conversation
curl localhost:9090/messages -XPOST -d '{"user_id":"a2181017-5c53-422b-b6bc-036b27c04fc8", "conversation_id":"a2181017-5c53-422b-b6bc-036b27c04fc8", "text":"This is a message"}'
curl localhost:9090/messages/a2181017-5c53-422b-b6bc-036b27c04fc8 -XDELETE

curl localhost:9090/conversations/72dca054-9bc3-4959-92cc-3dc1900391ab
curl localhost:9090/conversations -XPOST -d '{"user_id":["a2181017-5c53-422b-b6bc-036b27c04fc8", "e2382ea2-b5fa-4506-aa9d-d338aa52af44"], "game_id":""}'
curl localhost:9090/conversations -XPUT -d '{"user_id":["a2181017-5c53-422b-b6bc-036b27c04fc8", "e2382ea2-b5fa-4506-aa9d-d338aa52af44", "newUserID"], "game_id":"72dca054-9bc3-4959-92cc-3dc1900391ab"}'
curl localhost:9090/conversations/a2181017-5c53-422b-b6bc-036b27c04fc8 -XDELETE
