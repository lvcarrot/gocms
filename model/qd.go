package model

import (
	"sdbackend/domain"
)

type AdminQD struct {
	ID int64  `gorm:"primary_key"`
	QD string `gorm:"primary_key"`
}

type QDInstallRuns struct {
	ID                   int64 `gorm:"primary_key;auto_increment"`
	Date                 string
	QD                   string
	InstallStart         int64
	InstallEnd           int64
	UninstallStart       int64
	UninstallEnd         int64
	MFShow               int64
	ServerRun            int64
	Coefficient          int64
	Price                int64
	MFShowRetention1     int64
	MFShowRetention3     int64
	MFShowRetention7     int64
	MFShowRetention30    int64
	ServerRunRetention1  int64
	ServerRunRetention3  int64
	ServerRunRetention7  int64
	ServerRunRetention30 int64
	Total                int64 `gorm:"-"`
}

type GroupCoefficient struct {
	ID          int64 `gorm:"primary_key;auto_increment"`
	GroupName   string
	Coefficient int
	Price       int64
	Start       string
}

func AdmindQDs(id int64) ([]string, error) {
	var qds []string
	if err := db.New().Model(new(AdminQD)).Select("qd").Where("id = ?", id).Pluck("qd", &qds).Error; err != nil {
		return nil, err
	}
	if len(qds) == 0 {
		return nil, nil
	}

	return qds, nil
}

func AllQDs() ([]string, error) {
	var qds []string
	if err := db.New().Model(new(QDInstallRuns)).Select("qd").Where("install_end >= 10").Group("qd").Pluck("qd", &qds).Error; err != nil {
		return nil, err
	}
	if len(qds) == 0 {
		return nil, nil
	}
	return qds, nil
}

func InstallRunsByQD(qds []string, limit, offset int) ([]QDInstallRuns, error) {
	var ins []QDInstallRuns
	if err := db.New().Where("qd in (?)", qds).Where("install_end >= 10").Order("date desc").Limit(limit).Offset(offset).Find(&ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func TotalInstallRunsByQD(qds []string) (int64, error) {
	var total int64
	if err := db.New().Model(new(QDInstallRuns)).Where("qd in (?)", qds).Where("install_end >= 10").Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func GetGroupCoefficient(groupName string) (*GroupCoefficient, error) {
	var c GroupCoefficient
	if err := db.New().Model(new(GroupCoefficient)).Where("group_name = ?", groupName).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func MonthSettleByQD(qds []string, limit, offset int) ([]domain.MonthSettle, error) {
	var ins []domain.MonthSettle
	if err := db.New().Where("qd in (?)", qds).Order("month desc").Limit(limit).Offset(offset).Find(&ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func TotalMonthSettleByQD(qds []string) (int64, error) {
	var total int64
	if err := db.New().Model(new(domain.MonthSettle)).Where("qd in (?)", qds).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func GetQDRetentions(limit, offset int, rounds []int, from, to, qd string) ([]*domain.QDRetention, error) {
	var ins []*domain.QDRetention
	db := db.New().Model(&domain.QDRetention{}).Where("round in (?)", rounds)
	if from != "" {
		db = db.Where("date >= ?", from)
	}
	if to != "" {
		db = db.Where("date <= ?", to)
	}
	if qd != "" {
		db = db.Where("qd  = ?", qd)
	}
	dbResult := db.Order("date DESC").Order("round ASC").Limit(limit).Offset(offset).Find(&ins)
	if dbResult.RecordNotFound() {
		return nil, nil
	}
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return ins, nil
}

func TotalQDRetention(rounds []int, from, to, qd string) (int64, error) {
	var total int64
	db := db.New().Model(&domain.QDRetention{}).Where("round in (?)", rounds)
	if from != "" {
		db = db.Where("date >= ?", from)
	}
	if to != "" {
		db = db.Where("date <= ?", to)
	}
	if qd != "" {
		db = db.Where("qd = ?", qd)
	}
	if err := db.Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
