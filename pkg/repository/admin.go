package repository

import (
	"github.com/sgokul961/GO-PROJECT/pkg/domain"
	"github.com/sgokul961/GO-PROJECT/pkg/utils/models"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}
type AdminRepositoryinterfaces interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
}

func NewAdminRepository(DB *gorm.DB) AdminRepositoryinterfaces {
	return &AdminRepository{
		DB: DB,
	}

}
func (ad *AdminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := ad.DB.Raw("select *from admins where email =? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}
	return adminCompareDetails, nil

}
