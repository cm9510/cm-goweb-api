package models

type Members struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int8   `json:"age"`
	Job       string `json:"job"`
	CreatedAt int32  `json:"create_at"`
}

func (m *Members) TableName() string {
	return "t_users"
}
