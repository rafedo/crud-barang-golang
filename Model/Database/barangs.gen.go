// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

const TableNameBarang = "barangs"

// Barang mapped from table <barangs>
type Barang struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	Category string `gorm:"column:category;not null" json:"category"`
	Price    int64  `gorm:"column:price;not null" json:"price"`
}

// TableName Barang's table name
func (*Barang) TableName() string {
	return TableNameBarang
}
