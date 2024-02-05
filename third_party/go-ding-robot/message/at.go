package message

type At struct {
	AtBody *AtBody `json:"at,omitempty"`
}

// 被@人的手机号（在content里添加@人的手机号）
func (at *At) SetAtMobiles(mobiles []string) *At {
	if at.AtBody == nil {
		at.setDefaultAtBody()
	}
	at.AtBody.Mobiles = mobiles

	return at
}

// 是否@所有人
func (at *At) SetIsAll(isAll bool) *At {
	if at.AtBody == nil {
		at.setDefaultAtBody()
	}
	if isAll {
		at.SetAtMobiles(nil)
	}
	at.AtBody.IsAll = isAll

	return at
}

func (at *At) setDefaultAtBody() *At {
	at.AtBody = &AtBody{}

	return at
}

type AtBody struct {
	Mobiles []string `json:"atMobiles,omitempty"`
	IsAll   bool     `json:"isAtAll,omitempty"`
}
