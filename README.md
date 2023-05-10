# PROTOCOL DESCRIPTION

First, you need login to chat server

#### Request: Login/Auth

```json
{
  "type": "auth/login",
  "payload": {
    "token": "abc_user_token_123"
  }
}
```

#### Response: Success

```json
{
  "type": "auth/login",
  "payload": {
    "success": true
  }
}
```

#### Response: Error

```json
{
  "type": "auth/login",
  "payload": {
    "success": false
  }
}
```

### Get channel list

Get user channel list

#### Request: Get channel list

```json
{
  "type": "channel/list"
}
```

#### Response: Success

```json
{
  "type": "channel/list",
  "success": true,
  "payload": [
    {
      "id": 1,
      "title": "Channel 1"
    },
    {
      "id": 2,
      "title": "Channel 2"
    },
    {
      "id": 3,
      "title": "Channel 3"
    }
  ]
}
```

#### Response: Error

```json
{
  "type": "auth/login",
  "success": false,
  "payload": {
  }
}
```