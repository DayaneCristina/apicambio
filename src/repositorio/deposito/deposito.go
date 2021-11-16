package deposito

import (
	"api/src/banco"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	COLLECTIONNAME = "deposito"
)

func CreateOne(request *http.Request) (deposito banco.Deposito, erro error) {

	//Associa as informações do Body da Request na variavel deposito
	requestBody, erro := ioutil.ReadAll(request.Body)
	if erro = json.Unmarshal(requestBody, &deposito); erro != nil {
		return deposito, erro
	}

	//Pega conexão com o BD
	client, erro := banco.Conectar()

	if erro != nil {
		return deposito, erro
	}

	collection := client.Database(banco.DATABASE).Collection(COLLECTIONNAME)
	_, erro = collection.InsertOne(context.TODO(), deposito)

	if erro != nil {
		return deposito, erro
	}

	return deposito, erro

}

func GetAll() (depositos []banco.Deposito, erro error) {
	filter := bson.D{{}} //bson.D {{}} significa todos os documentos
	depositos = []banco.Deposito{}

	//Pega conexão com BD
	client, erro := banco.Conectar()
	if erro != nil {
		return depositos, erro
	}

	collection := client.Database(banco.DATABASE).Collection(COLLECTIONNAME)
	cursor, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return depositos, erro
	}

	//Mapeia o resultado da Query para "cortar"
	for cursor.Next(context.TODO()) {
		item := banco.Deposito{}
		err := cursor.Decode(&item)
		if err != nil {
			return depositos, erro
		}
		depositos = append(depositos, item)
	}
	cursor.Close(context.TODO())
	return depositos, erro
}
