package mysql

type User struct {
	Uid      string `gorm:"column:uid"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Province string `gorm:"column:province"`
	City     string `gorm:"column:city"`
	District string `gorm:"column:district"`
	Role     string `gorm:"column:role"`
}
