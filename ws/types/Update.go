package types

type Update struct {
	Type string
	Payload Payload
	Sender string
	Reciver string
}

type Payload struct {
	Data interface{}
}

