package repository

import (
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/rishad004/bw01/microservice-01/pkg/domain"
	pb "github.com/rishad004/bw01_proto-files/microservice-01"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
	RD *redis.Client
}

func NewRepository(Db *gorm.DB, Rd *redis.Client) Repo {
	return &repo{DB: Db, RD: Rd}
}

func (r *repo) UserCreate(name, email string) (int32, error) {

	user := domain.User{
		Name:  name,
		Email: email,
	}

	if err := r.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	r.RD.Set(strconv.Itoa(int(user.ID)), name+","+email, 24*time.Hour)

	return int32(user.ID), nil
}

func (r *repo) UserFetch(Id int32) (domain.User, error) {
	var user domain.User

	details, err := r.RD.Get(strconv.Itoa(int(Id))).Result()
	if err == nil {
		data := strings.Split(details, ",")
		user.Name = data[0]
		user.Email = data[1]
		user.ID = uint(Id)

		return user, nil
	}

	if err := r.DB.First(&user, Id).Error; err != nil {
		return domain.User{}, err
	}

	if err := r.RD.Set(strconv.Itoa(int(user.ID)), user.Name+","+user.Email, 24*time.Hour).Err(); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *repo) UserUpdate(id int32, name, email string) error {
	var user domain.User

	if err := r.DB.First(&user, id).Error; err != nil {
		return err
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}

	if err := r.DB.Save(&user).Error; err != nil {
		return err
	}

	if err := r.RD.Set(strconv.Itoa(int(id)), user.Name+","+user.Email, 24*time.Hour).Err(); err != nil {
		return err
	}

	return nil
}

func (r *repo) UserDelete(Id int32) error {
	if err := r.DB.Delete(&domain.User{}, Id).Error; err != nil {
		return err
	}

	r.RD.Del(strconv.Itoa(int(Id)))

	return nil
}

func (r *repo) FromMethod() (*pb.Data, error) {
	var users []domain.User

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	f := &pb.Data{}

	for _, v := range users {
		k := &pb.Details{
			Id:    int32(v.ID),
			Name:  v.Name,
			Email: v.Email,
		}
		f.Users = append(f.Users, k)
	}

	return f, nil
}
