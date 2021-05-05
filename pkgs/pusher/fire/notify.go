package fire

import (
	"context"

	firebase "firebase.google.com/go/v4"

	// "firebase.google.com/go/v4/auth"

	"google.golang.org/api/option"
)

var app *firebase.App

var ctx = context.Background()

func init() {
	var err error
	opt := option.WithCredentialsFile("./secrets/fire.json")
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
}

func PushNotify(data map[string]string, tokens []string) {
	// message := &messaging.MulticastMessage{
	// 	Data:   data,
	// 	Tokens: tokens,
	// 	Android: &messaging.AndroidConfig{
	// 		Priority: "high",
	// 	},
	// }
	// client, _ := app.Messaging(ctx)
	// client.SendMulticast(ctx, message)
}
