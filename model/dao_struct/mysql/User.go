package mysql

import "github.com/jlu-cow-studio/common/model/dao_struct/redis"

const (
	RoleNormal          = "normal"
	RoleProducer        = "producer"
	RoleMerchant        = "merchant"
	RoleVeterinary      = "veterinary"
	RoleBreeder         = "breeder"
	RoleServiceProvider = "service_provider"
)

type User struct {
	Uid      string `gorm:"column:uid"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Province string `gorm:"column:province"`
	City     string `gorm:"column:city"`
	District string `gorm:"column:district"`
	Role     string `gorm:"column:role"`
}

func UserFromRedis(userRedis *redis.UserInfo) *User {
	return &User{
		Uid:      userRedis.Uid,
		Username: userRedis.Username,
		Province: userRedis.Province,
		City:     userRedis.City,
		District: userRedis.District,
		Role:     userRedis.Role,
	}
}

func (user *User) ToRedis() *redis.UserInfo {
	return &redis.UserInfo{
		Uid:      user.Uid,
		Username: user.Username,
		Province: user.Province,
		City:     user.City,
		District: user.District,
		Role:     user.Role,
	}
}
