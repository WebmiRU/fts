# PROTOCOL DESCRIPTION

First, you need login to chat server

#### Request
```json
{
  "type": "auth/login",
  "payload": {
    "token": "abc_user_token_123"
  }
}
```

#### Success response
```json
{
  "type": "auth/login",
  "payload": {
    "success": true
  }
}
```

#### Error response
```json
{
  "type": "auth/login",
  "payload": {
    "success": false
  }
}
```

