package main

type messageAuthLoginOKPayload struct {
	Success bool `json:"success"`
}

type messageAuthLoginOK struct {
	Type    string                    `json:"type"`
	Payload messageAuthLoginOKPayload `json:"payload"`
}

var MessageAuthLoginOK = messageAuthLoginOK{
	Type: "auth/login",
	Payload: messageAuthLoginOKPayload{
		Success: true,
	},
}

type messageAuthLoginErrorPayload struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
}

type messageAuthLoginError struct {
	Type    string                       `json:"type"`
	Payload messageAuthLoginErrorPayload `json:"payload"`
}

var MessageAuthLoginError = messageAuthLoginError{
	Type: "auth/login",
	Payload: messageAuthLoginErrorPayload{
		Success: false,
		Code:    999,
		Text:    "Auth error",
	},
}

type messageUserChannelList struct {
	Type    string    `json:"type"`
	Payload []Channel `json:"payload"`
}
