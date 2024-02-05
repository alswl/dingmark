package message

const (
	BtnOrientationVertical   = "0"
	BtnOrientationHorizontal = "1"
)

type ActionCard struct {
	Message
	ActionCardBody *ActionCardBody `json:"actionCard"`
}

func NewActionCard() *ActionCard {
	ac := &ActionCard{}
	ac.SetType("actionCard")

	return ac
}

// 首屏会话透出的展示内容
func (ac *ActionCard) SetTitle(title string) *ActionCard {
	if ac.ActionCardBody == nil {
		ac.setDefaultActionCardBody()
	}
	ac.ActionCardBody.Title = title

	return ac
}

// markdown格式的消息
func (ac *ActionCard) SetText(text string) *ActionCard {
	if ac.ActionCardBody == nil {
		ac.setDefaultActionCardBody()
	}
	ac.ActionCardBody.Text = text

	return ac
}

// 0-按钮竖直排列，1-按钮横向排列
func (ac *ActionCard) SetBtnOrientation(orientation string) *ActionCard {
	if ac.ActionCardBody == nil {
		ac.setDefaultActionCardBody()
	}
	ac.ActionCardBody.BtnOrientation = orientation

	return ac
}

// 单个按钮的标题。(设置此项和singleURL后btns无效)
func (ac *ActionCard) SetBtnTitle(title string) *ActionCard {
	if ac.ActionCardBody == nil {
		ac.setDefaultActionCardBody()
	}
	ac.ActionCardBody.BtnTitle = title

	return ac
}

// 点击singleTitle按钮触发的URL
func (ac *ActionCard) SetBtnUrl(url string) *ActionCard {
	if ac.ActionCardBody == nil {
		ac.setDefaultActionCardBody()
	}
	ac.ActionCardBody.BtnURL = url

	return ac
}

// 按钮组
func (ac *ActionCard) SetBtns(btns []*ActionCardBtnBody) *ActionCard {
	if ac.ActionCardBody == nil {
		ac.setDefaultActionCardBody()
	}
	ac.ActionCardBody.BtnsBody = btns

	return ac
}

func (ac *ActionCard) AddBtn(btn *ActionCardBtnBody) *ActionCard {
	if ac.ActionCardBody == nil {
		ac.setDefaultActionCardBody()
	}
	ac.ActionCardBody.BtnsBody = append(ac.ActionCardBody.BtnsBody, btn)

	return ac
}

func (ac *ActionCard) setDefaultActionCardBody() *ActionCard {
	ac.ActionCardBody = &ActionCardBody{}

	return ac
}

type ActionCardBody struct {
	Title          string               `json:"title"`
	Text           string               `json:"text"`
	BtnOrientation string               `json:"btnOrientation,omitempty"`
	BtnTitle       string               `json:"singleTitle,omitempty"`
	BtnURL         string               `json:"singleURL,omitempty"`
	BtnsBody       []*ActionCardBtnBody `json:"btns,omitempty"`
}

type ActionCardBtnBody struct {
	Title string `json:"title"`
	Url   string `json:"actionURL"`
}

func NewActionCardBtn() *ActionCardBtnBody {
	return &ActionCardBtnBody{}
}

// 按钮标题
func (acbb *ActionCardBtnBody) SetTitle(text string) *ActionCardBtnBody {
	acbb.Title = text

	return acbb
}

// 点击按钮触发的URL
func (acbb *ActionCardBtnBody) SetUrl(url string) *ActionCardBtnBody {
	acbb.Url = url

	return acbb
}
