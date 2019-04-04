package model

import (
	"sdbackend/domain"

	"github.com/jinzhu/gorm"
)

func QDPricesTotal(start, end string) (int64, error) {
	var total int64
	db := db.New().Model(&domain.GroupCoefficient{})
	if start != "" {
		db = db.Where("start >= ?", start)
	}
	if end != "" {
		db = db.Where("start < ?", end)
	}
	if err := db.Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func QDPrices(limit, offset int, start, end string) ([]*domain.GroupCoefficient, error) {
	var results []*domain.GroupCoefficient
	db := db.New()
	if start != "" {
		db = db.Where("start >= ?", start)
	}
	if end != "" {
		db = db.Where("start < ?", end)
	}
	dbResult := db.Limit(limit).Offset(offset).Find(&results)
	if dbResult.RecordNotFound() {
		return nil, nil
	}
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return results, nil
}

func GetQDPriceByQD(qd string) (*domain.GroupCoefficient, error) {
	var ins domain.GroupCoefficient
	dbResult := db.New().Where("group_name = ?", qd).Take(&ins)
	if dbResult.RecordNotFound() {
		return nil, nil
	}
	if err := dbResult.Error; err != nil {
		return nil, err
	}
	return &ins, nil
}

func SaveQDPrice(p *domain.GroupCoefficient) error {
	var (
		updates = map[string]interface{}{
			"group_name":  p.QD,
			"coefficient": p.Coefficient,
			"price":       p.Price,
		}
		filter = func(db *gorm.DB) *gorm.DB {
			return db.Where("group_name = ?", p.QD)
		}
	)
	return db.New().Scopes(filter).Assign(updates).FirstOrCreate(p).Error
}
