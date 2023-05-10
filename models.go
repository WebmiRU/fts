package main

type Model struct {
	ID uint64 `gorm:"primarykey" json:"id"`
}

type User struct {
	Model
	Name     string
	Token    string
	Channels []Channel `gorm:"many2many:user_m2m_channel;"`
}

type Channel struct {
	Model
	Title string `json:"title"`
}

type ChannelMessage struct {
	Model
	UserId    uint64 `json:"user_id"`
	ChannelId uint64 `json:"channel_id"`
	Message   string `json:"message"`
	Text      string `json:"text"`
}
