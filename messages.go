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

//type messageChannelMessagePayload struct {
//	ChannelId uint64 `json:"channel_id"`
//	Message   string `json:"message"`
//	Text      string `json:"text"`
//}
//
//type MessageChannelMessage struct {
//	Type    string                       `json:"type"`
//	Payload messageChannelMessagePayload `json:"payload"`
//}
