package banco

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DATABASE         = "cambio"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func Conectar() (*mongo.Client, error) {
	//Performando a conexão com o mongo apenas uma vez
	mongoOnce.Do(func() {
		// Definindo configurações para o Client
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		// Conecta com o mongo
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		// Verifica a conexão
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
