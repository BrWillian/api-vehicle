package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vehicle struct {
	id                  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Ano_fabricacao      int16              `bson:"ano_fabricacao" json:"ano_fabricacao"`
	Ano_modelo          int16              `bson:"ano_modelo" json:"ano_modelo"`
	Chassi              string             `bson:"chassi" json:"chassi"`
	Combustivel         string             `bson:"combustivel" json:"combustivel"`
	Cor_veiculo         string             `bson:"cor_veiculo" json:"cor_veiculo"`
	Especie_veiculo     string             `bson:"especie_veiculo" json:"especie_veiculo"`
	Marca_modelo        string             `bson:"marca_modelo" json:"marca_modelo"`
	Municipio           string             `bson:"municipio" json:"municipio,omitempty"`
	Uf_placa            string             `bson:"uf_placa" json:"uf_placa"`
	Nacionalidade       string             `bson:"nacionalidade" json:"nacionalidade"`
	Placa_modelo_antigo string             `bson:"placa_modelo_antigo" json:"placa_modelo_antigo"`
	Placa_modelo_novo   string             `bson:"placa_modelo_novo" json:"placa_modelo_novo"`
	Tipo_veiculo        string             `bson:"tipo_veiculo" json:"tipo_veiculo"`
}
