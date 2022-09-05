package sso

type User struct {
	Id          string `json:"id" gorm:"id"`
	Username    string `json:"username" gorm:"username"`
	DisplayName string `json:"displayname" gorm:"displayname"`
	Password    string `json:"password" gorm:"password"`
	Salt        string `json:"salt" gorm:"salt"`
	Status      string `json:"status" gorm:"status"`
	Token       string `json:"token" gorm:"token"`
	DeletedAt   string `json:"deletedAt" gorm:"deleted_at"`
	CreatedAt   string `json:"createdAt" gorm:"created_at"`
	UpdatedAt   string `json:"updatedAt" gorm:"updated_at"`
	Remember    bool   `json:"remember" gorm:"remember"`
}
