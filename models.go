package main

type Model struct {
	ID uint64 `gorm:"primarykey"`
}

type User struct {
	Model
	Name     string
	Token    string
	Channels []Channel `gorm:"many2many:user_m2m_channel;"`
}

type Channel struct {
	Model
	Title string
}

type ChannelMessage struct {
	Model
	UserId    uint64
	ChannelId uint64
	Message   string
	Text      string
}
