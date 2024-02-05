package message

const MaxFeedCardLinks = 10

type FeedCard struct {
	Message
	FeedCardBody *FeedCardBody `json:"feedCard"`
}

func NewFeedCard() *FeedCard {
	fc := &FeedCard{}
	fc.SetType("feedCard")

	return fc
}

func (fc *FeedCard) SetLinks(links []*FeedCardLinkBody) *FeedCard {
	if fc.FeedCardBody == nil {
		fc.setDefaultFeedCardBody()
	}

	if len(links) > MaxFeedCardLinks {
		fc.FeedCardBody.Links = links[:MaxFeedCardLinks]
	} else {
		fc.FeedCardBody.Links = links
	}

	return fc
}

func (fc *FeedCard) AddLink(link *FeedCardLinkBody) *FeedCard {
	if fc.FeedCardBody == nil {
		fc.setDefaultFeedCardBody()
	}

	if fc.isFull() {
		return fc
	}

	fc.FeedCardBody.Links = append(fc.FeedCardBody.Links, link)

	return fc
}

// FeedCard目前最大支持10条
func (fc *FeedCard) isFull() bool {
	return len(fc.FeedCardBody.Links) >= MaxFeedCardLinks
}

func (fc *FeedCard) setDefaultFeedCardBody() *FeedCard {
	fc.FeedCardBody = &FeedCardBody{}

	return fc
}

type FeedCardBody struct {
	Links []*FeedCardLinkBody `json:"links"`
}

type FeedCardLinkBody struct {
	Title string `json:"title"`
	Pic   string `json:"picURL"`
	Url   string `json:"messageURL"`
}

func NewFeedCardLink() *FeedCardLinkBody {
	return &FeedCardLinkBody{}
}

// 单条信息文本
func (fclb *FeedCardLinkBody) SetTitle(text string) *FeedCardLinkBody {
	fclb.Title = text

	return fclb
}

// 单条信息后面图片的URL
func (fclb *FeedCardLinkBody) SetPic(pic string) *FeedCardLinkBody {
	fclb.Pic = pic

	return fclb
}

// 点击单条信息到跳转链接
func (fclb *FeedCardLinkBody) SetUrl(url string) *FeedCardLinkBody {
	fclb.Url = url

	return fclb
}
