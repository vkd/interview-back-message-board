# [Job interview - Back] Message board

The Backend of a message board application.

## TODO

[ ] admin users storage
[ ] email validation

## API

### Structure
#### Response

In general a response looks like this:

```Status: 200 OK```
```
{
  "status": "ok",
  "data": ...
}
```
The reason why I decided to use this way - because there is no way to see the difference in case `Status: 404 Not found` between 2 cases: `wrong url` and `message with that :id is not found`.

### Status codes
TODO: describe the status codes that can be returned and what they mean (as example if password is wrong)

## Public API methods

### Post new message
`POST /message`
Request:
```
{
  "name": "my_name",
  "email": "my_email",
  "text": "my_text"
}
```
Response:
```
{
  "status": "ok",
  "data": {
    "id": "2C7BCEC7-CD14-D6E5-3FBF-F9551375429A"
  }
}
```

## Private API methods

That method use HTTP Basic Authentication.
A request should contain the header:
`Authorization: Basic base64(username:password)`

### List of messages
`GET /message?order=-created_at`

Response:
```
{
  "status": "ok",
  "data": [
    {
      "id": "2C7BCEC7-CD14-D6E5-3FBF-F9551375429A",
      "name": "my_name",
      "email": "my_email",
      "text": "my_text",
      "creation_time": "2020-01-01T10:00:00"
    },
    ...
  ]
}
```

### Get a message by id
`GET /message/:id`

Response:
```
{
  "status": "ok",
  "data": {
    "id": "2C7BCEC7-CD14-D6E5-3FBF-F9551375429A",
    "name": "my_name",
    "email": "my_email",
    "text": "my_text",
    "creation_time": "2020-01-01T10:00:00"
  }
}
```

### Update the message text by id
`POST /message/:id`

Request:
```
{
  "name": "new_name",
  "email": "new_email",
  "text": "new_text"
}
```

Response:
```
{
  "status": "ok"
}
```
