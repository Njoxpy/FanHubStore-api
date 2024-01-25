// connect to a database

package initializers

// import packages(gorm and postGress)
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )

  
func ConnectDB() {
	// create a database using elephant sql

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
