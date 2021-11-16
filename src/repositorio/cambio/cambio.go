package cambio

import (
	"api/src/banco"
	"api/src/repositorio/deposito"
	"context"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	COLLECTIONNAME = "moedas"
)

func CreateOne(r *http.Request) (moeda banco.Moeda, erro error) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro = json.Unmarshal(requestBody, &moeda); erro != nil {
		return moeda, erro
	}
	client, erro := banco.Conectar()

	if erro != nil {

		return moeda, erro
	}

	collection := client.Database(banco.DATABASE).Collection(COLLECTIONNAME)
	_, erro = collection.InsertOne(context.TODO(), moeda)

	if erro != nil {
		return moeda, erro
	}

	return moeda, erro

}

func GetAll() (moedas []banco.Moeda, erro error) {
	filter := bson.D{{}}
	moedas = []banco.Moeda{}

	client, erro := banco.Conectar()
	if erro != nil {
		return moedas, erro
	}

	collection := client.Database(banco.DATABASE).Collection(COLLECTIONNAME)
	cursor, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return moedas, erro
	}

	for cursor.Next(context.TODO()) {
		item := banco.Moeda{}
		err := cursor.Decode(&item)
		if err != nil {
			return moedas, erro
		}
		moedas = append(moedas, item)
	}
	cursor.Close(context.TODO())
	return moedas, erro
}

func ConverterSaldo(r *http.Request) (total float64, totalConvertido float64, erro error) {
	// Pega os parâmetros da URL (no nosso caso só o "moeda")
	parametros := mux.Vars(r)

	client, erro := banco.Conectar()

	if erro != nil {
		return total, totalConvertido, erro
	}

	//Pegando todos os Depositos
	filterDepositos := bson.D{{}}

	collectionDeposito := client.Database(banco.DATABASE).Collection(deposito.COLLECTIONNAME)
	cursorDeposito, findError := collectionDeposito.Find(context.TODO(), filterDepositos)

	if findError != nil {
		return total, totalConvertido, findError
	}

	// Percorre os depósitos encontrados (no caso todos)
	// e faz a soma do valor do depósito percorrido com a
	// variável "total"
	for cursorDeposito.Next(context.TODO()) {
		item := banco.Deposito{}
		err := cursorDeposito.Decode(&item)
		if err != nil {
			return total, totalConvertido, erro
		}
		total += item.Valor
	}
	cursorDeposito.Close(context.TODO())

	//Pegando Moeda
	var moeda bson.M // ?
	collection := client.Database(banco.DATABASE).Collection(COLLECTIONNAME)

	// Tenta encontrar um registro que bata com o filtro (bson.M{"code": "USD"}))
	// Encontrou?! Faz o decode baseado na var moeda declarada acima
	if err := collection.FindOne(context.TODO(), bson.M{"code": parametros["moeda"]}).Decode(&moeda); err != nil {
		return total, totalConvertido, err
	}

	// Pega a posição "valor" do nosso "bson.M" (variável moeda)
	// e converte para float64 (porque ele vem como "interface" ???)
	valor := moeda["valor"].(float64)
	valor = total * valor

	// Depois de multiplicar o total com o valor da moeda encontrada
	// "transforma" o total convertido em float com 2 casas decimais
	totalConvertido = math.Floor(valor*100) / 100

	return total, totalConvertido, erro
}
