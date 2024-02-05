package message

type Markdown struct {
	Message
	At
	MarkdownBody *MarkdownBody `json:"markdown"`
}

func NewMarkdown() *Markdown {
	m := &Markdown{}
	m.SetType("markdown")

	return m
}

// markdown格式的消息
func (m *Markdown) SetText(text string) *Markdown {
	if m.MarkdownBody == nil {
		m.setDefaultMarkdownBody()
	}
	m.MarkdownBody.Text = text

	return m
}

// 首屏会话透出的展示内容
func (m *Markdown) SetTitle(title string) *Markdown {
	if m.MarkdownBody == nil {
		m.setDefaultMarkdownBody()
	}
	m.MarkdownBody.Title = title

	return m
}

func (m *Markdown) setDefaultMarkdownBody() *Markdown {
	m.MarkdownBody = &MarkdownBody{}

	return m
}

type MarkdownBody struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
