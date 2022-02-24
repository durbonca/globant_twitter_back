package db

import (
	"context"
	"fmt"
	"time"

	"github.com/durbonca/globant_twitter_back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProfile(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnection.Database("globant-twitter")
	col := db.Collection("usuarios")
	registro := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Bio) > 0 {
		registro["bio"] = u.Bio
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.Web) > 0 {
		registro["web"] = u.Web
	}
	registro["fechaNac"] = u.FechaNac

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	fmt.Println(objID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}
	fmt.Println(filtro)
	_, err := col.UpdateOne(ctx, filtro, updateString)

	if err != nil {
		return false, err
	}
	return true, nil
}
