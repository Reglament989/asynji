package types

type Notification struct {
	Notifications []NotificationPayload `json:"notifications"`
}

type NotificationPayload struct {
	Tokens   []string `json:"tokens"`
	Platform int      `json:"platform"`
	Message  string   `json:"message"`
	Title    string   `json:"title"`
}

type Response struct {
	Counts  int      `json:"counts"`
	Logs    []string `json:"logs"`
	Success string   `json:"success"`
}
