package usecases

import (
	"errors"

	"github.com/jinzhu/copier"
	"github.com/sgokul961/GO-PROJECT/pkg/domain"
	"github.com/sgokul961/GO-PROJECT/pkg/repository"
	"github.com/sgokul961/GO-PROJECT/pkg/utils/models"
)

type AdminUseCaseinterface interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
}

type adminUseCase struct {
	AdminRepository repository.AdminRepositoryinterfaces
}

func NewAdminUseCase(repo repository.AdminRepositoryinterfaces) AdminUseCaseinterface {
	return &adminUseCase{
		AdminRepository: repo,
	}
}

func (ad *adminUseCase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	//getting details of the admin based on the email provided

	adminFromDB, err := ad.AdminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	//compare password that provided from the databse provided from admins
	// 2. Compare password
	if adminDetails.Password != adminFromDB.Password {
		return domain.TokenAdmin{}, errors.New("invalid password")
	}

	// err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	// fmt.Println(err)
	// if err != nil {
	// 	return domain.TokenAdmin{}, err
	// }

	var adminDetailsResponse models.AdminDetailsResponse

	err = copier.Copy(&adminDetailsResponse, &adminFromDB)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin: adminDetailsResponse,
		//Token: tokenString,
	}, nil

}
