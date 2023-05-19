package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	. "vehicle-search/config"
	"vehicle-search/models"
)

var config Config
var collection = new(mongo.Collection)

const VehiclesCollection = "vehicles"

type VehiclesRepository struct{}

func init() {
	config.Read()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Database.Uri))
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(config.Database.DatabaseName).Collection(VehiclesCollection)
}
func (p *VehiclesRepository) Find(placa string) (models.Vehicle, error) {
	filterAntiga := bson.M{"placa_modelo_antigo": placa}
	filterNova := bson.M{"placa_modelo_novo": placa}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var vehicle models.Vehicle
	err := collection.FindOne(ctx, filterAntiga).Decode(&vehicle)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = collection.FindOne(ctx, filterNova).Decode(&vehicle)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					vehicle = models.Vehicle{
						Placa_modelo_antigo: "",
						Placa_modelo_novo:   "",
						Ano_fabricacao:      0,
						Ano_modelo:          0,
						Chassi:              "",
						Combustivel:         "",
						Cor_veiculo:         "",
						Especie_veiculo:     "",
						Marca_modelo:        "",
						Municipio:           "",
						Uf_placa:            "",
						Nacionalidade:       "",
						Tipo_veiculo:        "",
					}
				} else {
					return models.Vehicle{}, err
				}
			}
		} else {
			return models.Vehicle{}, err
		}
	}
	return vehicle, nil
}
