package main

import (
	"context"
	"solid/db"
	"solid/helper"
	"solid/models"
	"solid/repositories"
)

func main() {

	ctx := context.Background()

	db, err := db.NewGormDatabase()
	if err != nil {
		panic("error connecting to db!")
	}

	db.AutoMigrate(&models.User{})

	//setup db
	database := repositories.NewDbOperations(db)

	//create
	create := database.Create(ctx, models.User{Name: "Wow"})
	if !create {
		panic("error creating to db")
	}

	//read
	get, err := database.Read(ctx)
	helper.DataOutput(get)

	//update
	update := database.Update(ctx, models.User{Name: "Berotak Senku 👉🏿🤓👈🏿"}, get.ID)
	if !update {
		panic("update to db error!")
	}

	//read
	getAfterUpdate, err := database.Read(ctx)
	helper.DataOutput(getAfterUpdate)

	//delete
	deleteData := database.Delete(ctx, get.ID)

	if !deleteData {
		panic("cannot delete")
	}
}
