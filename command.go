package main

func cmdLogin(token string) User {
	var model User
	db.Preload("Channels").Where("token = ?", token).First(&model)

	return model
}
