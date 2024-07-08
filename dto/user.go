package dto

// User 用户
type User struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	Avatar           string `json:"avatar"`
	Bot              bool   `json:"bot"`
	UserOpenid       string `json:"user_openid"`        // c2c私聊时用户openid
	MemberOpenid     string `json:"member_openid"`      //群聊@机器人时用户openid
	UnionOpenID      string `json:"union_openid"`       // 特殊关联应用的 openid
	UnionUserAccount string `json:"union_user_account"` // 机器人关联的用户信息，与union_openid关联的应用是同一个
}
