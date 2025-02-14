package migration

import (
	"time"

	"gorm.io/gorm"

	"github.com/andys920605/hr-system/pkg/logging"
	"github.com/andys920605/hr-system/pkg/snowflake"
)

func SeedData(db *gorm.DB, logger *logging.Logging) {
	var count int64
	db.Table("employee").Model(&Employee{}).Count(&count)

	if count == 0 {
		employees := []Employee{
			{
				ID:       ID(snowflake.New()),
				Info:     PersonalInfo{Name: "Alice", Email: "alice@example.com", Phone: "123456789", Address: "Taipei"},
				Position: 1, JobLevel: 3, Status: 1,
			},
			{
				ID:       ID(snowflake.New()),
				Info:     PersonalInfo{Name: "Bob", Email: "bob@example.com", Phone: "987654321", Address: "New York"},
				Position: 2, JobLevel: 4, Status: 1,
			},
			{
				ID:       ID(snowflake.New()),
				Info:     PersonalInfo{Name: "Charlie", Email: "charlie@example.com", Phone: "555555555", Address: "Tokyo"},
				Position: 3, JobLevel: 5, Status: 0,
			},
		}

		if err := db.Table("employee").Create(&employees).Error; err != nil {
			logger.Errorf("Failed to seed data:", err)
		} else {
			logger.Info("Seed data inserted successfully")
		}
	} else {
		logger.Info("Table already has data, skipping seeding")
	}
}

type ID int64

type PersonalInfo struct {
	Name    string `gorm:"column:name;type:varchar(255);not null"`
	Email   string `gorm:"column:email;type:varchar(255);unique;not null"`
	Phone   string `gorm:"column:phone;type:varchar(20)"`
	Address string `gorm:"column:address;type:text"`
}

type Employee struct {
	ID        ID           `gorm:"column:id;primaryKey;autoIncrement:false"`
	Info      PersonalInfo `gorm:"embedded"`
	Position  int          `gorm:"column:position;type:tinyint;not null"`
	JobLevel  int          `gorm:"column:job_level;type:tinyint;not null"`
	Status    int          `gorm:"column:status;type:tinyint;not null"`
	CreatedAt time.Time    `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time    `gorm:"column:updated_at;autoUpdateTime"`
}
