package messenger

type FacebookQuickReply struct {
	ContentType string `json:"content_type"`
	Title       string `json:"title"`
	Payload     string `json:"payload"`
	ImageUrl    string `json:"image_url"`
}

type FacebookMessage struct {
	Text         string                `json:"text"`
	QuickReplies []*FacebookQuickReply `json:"quick_replies"`
}

type FacebookResponse struct {
	MessagingType string `json:"messaging_type""`
	Recipient     struct {
		Id string `json:"id"`
	} `json:"recipient"`
	Message *FacebookMessage `json:"message"`
}
