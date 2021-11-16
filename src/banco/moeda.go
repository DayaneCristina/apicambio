package banco

type Moeda struct {
	Code  string  `bson:"code" json:"code"`
	Valor float64 `bson:"valor" json:"valor"`
}
