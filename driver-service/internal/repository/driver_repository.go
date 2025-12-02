package repository

import (
	"context"
	"math"
	"sort"
	"time"

	"github.com/aemreakyuz/bitaksi-taxihub/driver-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DriverRepository struct {
	collection *mongo.Collection
}

func NewDriverRepository(client *mongo.Client) *DriverRepository {
	collection := client.Database("taxihub").Collection("drivers")
	return &DriverRepository{
		collection: collection,
	}
}

func (r *DriverRepository) Create(driver *model.Driver) error {
	driver.CreatedAt = time.Now()
	driver.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(context.Background(), driver)
	if err != nil {
		return err
	}

	// Set the generated ID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		driver.ID = oid.Hex()
	}

	return nil
}

func (r *DriverRepository) Update(id string, driver *model.Driver) error {
	driver.UpdatedAt = time.Now()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": driver}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *DriverRepository) FindAll(page, pageSize int) ([]*model.Driver, error) {
	var drivers []*model.Driver

	skip := (page - 1) * pageSize

	cursor, err := r.collection.Find(
		context.Background(),
		bson.M{},
		options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize)),
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &drivers); err != nil {
		return nil, err
	}

	return drivers, nil
}

func (r *DriverRepository) FindNearby(lat, lon float64, taxiType string, radiusKm float64) ([]*model.Driver, error) {
	var drivers []*model.Driver

	filter := bson.M{}
	if taxiType != "" {
		filter["taxiType"] = taxiType
	}

	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &drivers); err != nil {
		return nil, err
	}

	// Filter by distance and calculate distances
	var nearbyDrivers []*model.Driver
	for _, driver := range drivers {
		distance := haversine(lat, lon, driver.Lat, driver.Lon)
		if distance <= radiusKm {
			nearbyDrivers = append(nearbyDrivers, driver)
		}
	}

	// Sort by distance (closest first)
	sort.Slice(nearbyDrivers, func(i, j int) bool {
		distI := haversine(lat, lon, nearbyDrivers[i].Lat, nearbyDrivers[i].Lon)
		distJ := haversine(lat, lon, nearbyDrivers[j].Lat, nearbyDrivers[j].Lon)
		return distI < distJ
	})

	return nearbyDrivers, nil
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadiusKm = 6371.0

	dLat := (lat2 - lat1) * (math.Pi / 180.0)
	dLon := (lon2 - lon1) * (math.Pi / 180.0)

	lat1Rad := lat1 * (math.Pi / 180.0)
	lat2Rad := lat2 * (math.Pi / 180.0)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}
