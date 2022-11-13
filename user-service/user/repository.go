package user

import (
	"github.com/hashicorp/go-uuid"
	"gorm.io/gorm"
	pb "user-service/proto/user"
)

type IRepository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(*pb.User) error
	GetByEmailAndPassword(*pb.User) (*pb.User, error)
	GetByEmail(string) (*pb.User, error)
}

type Repository struct {
	Db *gorm.DB
}

func (repo *Repository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *Repository) Get(id string) (*pb.User, error) {
	user := new(pb.User)
	user.Id = id
	if err := repo.Db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *Repository) Create(user *pb.User) error {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return err
	}
	user.Id = id
	if err := repo.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.Db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *Repository) GetByEmail(email string) (*pb.User, error) {
	user := new(pb.User)
	if err := repo.Db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}