package mysql

import (
	"time"
)

type ID int64

type PersonalInfo struct {
	Name    string `gorm:"column:name;type:varchar(255);not null"`
	Email   string `gorm:"column:email;type:varchar(255);unique;not null"`
	Phone   string `gorm:"column:phone;type:varchar(20)"`
	Address string `gorm:"column:address;type:text"`
}

type Employee struct {
	ID        ID           `gorm:"column:id;primaryKey;autoIncrement:false"` // 使用 Snowflake ID
	Info      PersonalInfo `gorm:"embedded"`
	Position  int          `gorm:"column:position;type:tinyint;not null"`  // 職位 (Engineer, Manager, Admin)
	JobLevel  int          `gorm:"column:job_level;type:tinyint;not null"` // 工作等級 (1~5)
	Status    int          `gorm:"column:status;type:tinyint;not null"`    // 0: Resigned, 1: Active
	CreatedAt time.Time    `gorm:"column:created_at;autoCreateTime"`       // 自動設定建立時間
	UpdatedAt time.Time    `gorm:"column:updated_at;autoUpdateTime"`       // 自動更新修改時間
}
