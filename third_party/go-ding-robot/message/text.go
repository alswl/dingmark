package message

type Text struct {
	Message
	At
	TextBody *TextBody `json:"text"`
}

func NewText() *Text {
	t := &Text{}
	t.SetType("text")

	return t
}

// 消息内容
func (t *Text) SetContent(content string) *Text {
	if t.TextBody == nil {
		t.setDefaultTextBody()
	}
	t.TextBody.Content = content

	return t
}

func (t *Text) setDefaultTextBody() *Text {
	t.TextBody = &TextBody{}

	return t
}

type TextBody struct {
	Content string `json:"content"`
}
