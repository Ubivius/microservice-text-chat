# Microservice-Text-Chat
Text chat microservice for our online game framework.

## Text chat endpoints

`GET` `/messages/{id}` Returns json data about a specific message. `id=[string]`

`GET` `/conversations/{id}` Returns json data about a specific conversation. `id=[string]`

`GET` `/messages/conversation/{id}` Returns the list of messages of a specific conversation. `id=[string]`

`GET` `/health/live` Returns a Status OK when live.

`GET` `/health/ready` Returns a Status OK when ready or an error when dependencies are not available.

`POST` `/messages` Add new message with specific data. </br>
__Data Params__
```json
{
  "user_id":         "string, required",
  "conversation_id": "string, required",
  "text":            "string, required",
}
```

`POST` `/conversations` Add new message with specific data. </br>
__Data Params__
```json
{
  "user_id": ["string, required"],
  "game_id": "string, required",
}
```

`PUT` `/conversations` Add new or remove users from a conversation. </br>
__Data Params__
```json
{
  "user_id": ["string, required"],
  "game_id": "string, required",
}
```

`DELETE` `/messages/{id}` Delete a message.  `id=[string]`

`DELETE` `/conversations/{id}` Delete a conversation and all associated messages.  `id=[string]`
