package message

type Link struct {
	Message
	LinkBody *LinkBody `json:"link"`
}

func NewLink() *Link {
	l := &Link{}
	l.SetType("link")

	return l
}

// 消息内容。如果太长只会部分展示
func (l *Link) SetText(text string) *Link {
	if l.LinkBody == nil {
		l.setDefaultLinkBody()
	}
	l.LinkBody.Text = text

	return l
}

// 消息标题
func (l *Link) SetTitle(title string) *Link {
	if l.LinkBody == nil {
		l.setDefaultLinkBody()
	}
	l.LinkBody.Title = title

	return l
}

// 图片URL
func (l *Link) SetPic(pic string) *Link {
	if l.LinkBody == nil {
		l.setDefaultLinkBody()
	}
	l.LinkBody.Pic = pic

	return l
}

// 点击消息跳转的URL
func (l *Link) SetUrl(url string) *Link {
	if l.LinkBody == nil {
		l.setDefaultLinkBody()
	}
	l.LinkBody.Url = url

	return l
}

func (l *Link) setDefaultLinkBody() *Link {
	l.LinkBody = &LinkBody{}

	return l
}

type LinkBody struct {
	Text  string `json:"text"`
	Title string `json:"title"`
	Pic   string `json:"picUrl"`
	Url   string `json:"messageUrl"`
}
