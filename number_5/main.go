package main

import {
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
}

type ItemsDetail struct {
	ID int64 
	Name string
}

func createItem(gormDB *gorm.DB) error {
	item := ItemsDetail{
		ID: 1,
		Name: "Desk Fan",
	}

	result := gormDB.Create(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func readItem(gormDB *gorm.DB) {
	var listItems []ItemsDetail

	gormDB.find(&listItems)

	fmt.Println(listItemDetails)
}

func updateItem(ID int64, gormDB *gorm.DB) {
	gormDB.Model(&ItemsDetail{}).Where("id  = ?", ID).Update("Name", "wall fan")
}

func deleteItem(ID int64, gormDB *gorm.DB) {
	db.Delete(&ItemsDetail{}, ID)
}

func main() {
	// using gorm as orm
	dsn := "host=localhost user=psql password=psql dbname=testnumber5 port=5321 sslmode=disable"
	gormDB, err := gorm.Open(postgres.Open(dsn),&gorm.Config())
	if err != nil {
		log.Fatal(err)
	}
	defer gormDB.Close()

	// create
	err = createItem(gormDB)
	if err != nil {
		log.Fatal(err)
	}

	// read
	readItem(gormDB)

	// update
	updateItem(1, gormDB)

	// delete
	deleteItem(1, gormDB)

}
