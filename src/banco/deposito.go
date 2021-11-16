package banco

import "go.mongodb.org/mongo-driver/bson/primitive"

type Deposito struct {
	Valor         float64             `bson:"valor" json:"valor"`
	Data          primitive.Timestamp `bson:"data" json:"data"`
	NumeroConta   string              `bson:"numero_conta" json:"numero_conta"`
	NumeroAgencia string              `bson:"numero_agencia" json:"numero_agencia"`
}
