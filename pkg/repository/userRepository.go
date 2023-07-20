package repository

import (
	"github.com/athunlal/bookNow-Account-Services/pkg/domain"
	interfaces "github.com/athunlal/bookNow-Account-Services/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserDataBase struct {
	DB *gorm.DB
}

func (r *UserDataBase) FindProfile(user domain.User) (domain.User, int64) {
	result := r.DB.Raw("select * from users where id = ?", user.Id).Scan(&user)
	return user, result.RowsAffected
}

func (r *UserDataBase) EditProfile(user domain.User) int {
	result := r.DB.Exec("UPDATE users SET username = ?, dateofbirth = ?, gender = ? where id = ?", user.Username, user.Dateofbirth, user.Gender, user.Id)
	return int(result.RowsAffected)
}

func (r *UserDataBase) UpdatePassword(passwordData domain.Password) int64 {
	result := r.DB.Exec("UPDATE users SET password = ? where id = ?", passwordData.Newpassword, passwordData.Id)
	return result.RowsAffected
}

func (r *UserDataBase) CreateAddress(addressData domain.Address) (domain.Address, error) {
	result := r.DB.Create(&addressData)
	return addressData, result.Error
}
func (r *UserDataBase) ViewAllAddress(addressData domain.Address) ([]domain.Address, int64) {
	var address []domain.Address
	result := r.DB.Raw("select * from addresses where uid = ?", addressData.Uid).Scan(&address)
	return address, result.RowsAffected
}

func (r *UserDataBase) EditAddress(addressData domain.Address) (domain.Address, int64) {
	result := r.DB.Exec("UPDATE addresses SET type = ?, locationaddress = ?, complete_address = ?, landmark = ?, floorno = ? where uid = ? AND addressid = ?", addressData.Type, addressData.Locationaddress, addressData.CompleteAddress, addressData.Landmark, addressData.Floorno, addressData.Uid, addressData.Addressid)
	return addressData, result.RowsAffected
}

func (r *UserDataBase) FindByUserName(user domain.User) (domain.User, int64) {
	result := r.DB.Raw("select * from users where username LIKE ?", user.Username).Scan(&domain.User{})
	return user, result.RowsAffected
}

func (r *UserDataBase) ViewAddressByID(addressData domain.Address) (domain.Address, int64) {
	result := r.DB.Raw("select * from addresses where uid = ? and addressid = ?", addressData.Uid, addressData.Addressid).Scan(&addressData)
	return addressData, result.RowsAffected
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserDataBase{
		DB: db,
	}
}
