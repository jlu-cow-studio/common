package mysql

type Tag struct {
	ID         int64   `gorm:"column:id;primaryKey;autoIncrement"`
	Name       string  `gorm:"column:name;not null"`
	Weight     float64 `gorm:"column:weight;not null;type:decimal(8,4);default:0.0000"`
	MarkObject string  `gorm:"column:mark_object;not null"`
	CategoryID int64   `gorm:"column:category_id;not null"`
}

func (Tag) TableName() string {
	return "tag"
}

type TagCategory struct {
	ID       int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:name"`
	ParentID int    `gorm:"column:parent_id"`
	Level    int8   `gorm:"column:level"`
}

type TagWithCategory struct {
	ID         int64        `gorm:"column:id;primaryKey;autoIncrement"`
	Name       string       `gorm:"column:name"`
	Weight     float64      `gorm:"column:weight"`
	MarkObject string       `gorm:"column:mark_object"`
	CategoryID int64        `gorm:"column:category_id"`
	Category   *TagCategory `gorm:"foreignKey:CategoryID"`
}

type TagCategoryWithList struct {
	TagCategory
	TagList []*Tag `gorm:"foreignKey:CategoryID"`
}
