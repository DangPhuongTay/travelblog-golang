package models

type Blog struct {
	Id      uint   `son:"id"`
	Title   string `json:title`
	Topic   string `json:topic`
	Time    string `json:time`
	Desc1   string `json:desc1`
	Desc2   string `json:desc2`
	Desc3   string `json:desc3`
	Contact string `json:contact`
	Price   string `json:price`
	Image   string `json:image`
	Address string `json:address`
	Map     string `json:map`
	UserID  uint   `json:user_id`
	User    User   `json:"user";gorm:"foreignkey:UserID"`
}
