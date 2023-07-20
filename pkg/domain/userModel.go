package domain

type Address struct {
	Addressid       uint   `JSON:"addressid" gorm:"primarykey;unique"`
	User            User   `gorm:"ForeignKey:uid"`
	Uid             uint   `JSON:"uid"`
	Type            string `JSON:"type" gorm:"not null"`
	Locationaddress string `JSON:"locationaddress" gorm:"not null"`
	CompleteAddress string `JSON:"completeAddress" gorm:"not null"`
	Landmark        string `JSON:"landmark" gorm:"not null"`
	Floorno         string `JSON:"floorno" gorm:"not null"`
}

type User struct {
	Id          uint   `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Username    string `json:"username" gorm:"unique" validate:"required,min=2,max=50"`
	Password    string `json:"password"`
	Email       string `json:"email" gorm:"notnull;unique" validate:"email,required"`
	Phone       string `json:"phone" validate:"required,len=10"`
	Profile     string `json:"profile"`
	Dateofbirth string `json:"dateofbirth" gorm:"default:null"`
	Gender      string `json:"gender" gorm:"default:null" `
}

type Password struct {
	Id          uint   `json:"id"`
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}
