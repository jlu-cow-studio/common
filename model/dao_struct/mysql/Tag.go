package mysql

const (
	MarkObject_WholeCattle    = "whole_cattle"
	MarkObject_CattleProduct  = "cattle_product"
	MarkObject_Breeding       = "breeding"
	MarkObject_Service        = "service"
	MarkObject_ServiceProduct = "service_product"
	MarkObject_Live           = "live"
	MarkObject_Tweet          = "tweet"
	MarkObject_User           = "user"
	MarkObject_Interest       = "interest"
)

var MarkObjectMap = map[string]interface{}{
	MarkObject_WholeCattle:    nil,
	MarkObject_CattleProduct:  nil,
	MarkObject_Breeding:       nil,
	MarkObject_Service:        nil,
	MarkObject_ServiceProduct: nil,
	MarkObject_Live:           nil,
	MarkObject_Tweet:          nil,
	MarkObject_User:           nil,
	MarkObject_Interest:       nil,
}

type Tag struct {
	ID         int64   `gorm:"column:tag_id;primaryKey;autoIncrement"`
	Name       string  `gorm:"column:tag_name;not null"`
	Weight     float64 `gorm:"column:weight;not null;type:decimal(8,4);default:0.0000"`
	MarkObject string  `gorm:"column:mark_object;not null"`
	CategoryID int64   `gorm:"column:category_id;not null"`
}

func (Tag) TableName() string {
	return "tag"
}

type TagCategory struct {
	ID       int    `gorm:"column:tag_category_id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:tag_category_name"`
	ParentID int    `gorm:"column:parent_id"`
	Level    int8   `gorm:"column:level"`
}

type TagWithCategory struct {
	ID         int64        `gorm:"column:tag_id;primaryKey;autoIncrement"`
	Name       string       `gorm:"column:tag_name"`
	Weight     float64      `gorm:"column:weight"`
	MarkObject string       `gorm:"column:mark_object"`
	CategoryID int64        `gorm:"column:category_id"`
	Category   *TagCategory `gorm:"foreignKey:CategoryID"`
}

type TagCategoryWithList struct {
	TagCategory
	TagList []*Tag `gorm:"foreignKey:category_id"`
}
