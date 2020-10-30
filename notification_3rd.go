package jpushclient

type Notification3rd struct {
	Content   string                 `json:"content,omitempty"`
	Title     string                 `json:"title,omitempty"`
	ChannelId string                 `json:"channel_id,omitempty"`
	Extras    map[string]interface{} `json:"extras,omitempty"`
}

func (c *Notification3rd) SetTitle(title string) {
	c.Title = title
}

func (c *Notification3rd) SetContent(content string) {
	c.Content = content
}

func (c *Notification3rd) SetExtras(extras map[string]interface{}) {
	c.Extras = extras
}
