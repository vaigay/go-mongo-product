package impl

import (
	"context"
	"go-mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepoImpl struct {
	Db *mongo.Database
}

func NewProductRepo(db *mongo.Database) *ProductRepoImpl {
	return &ProductRepoImpl{
		Db: db,
	}
}

func (mongoImpl *ProductRepoImpl) CreateProduct(product *models.Product) error {
	bbytes, err := bson.Marshal(product)
	if err != nil {
		return err
	}

	insertResult, err := mongoImpl.Db.Collection("products").InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}
	product.Id = insertResult.InsertedID.(primitive.ObjectID)

	return nil
}

func (mongoImpl *ProductRepoImpl) FindAll() ([]models.Product, error) {
	listProduct := make([]models.Product, 0)
	ctx := context.Background()
	cursor, err := mongoImpl.Db.Collection("products").Find(ctx, bson.M{})
	if err != nil {
		return listProduct, err
	}
	defer cursor.Close(ctx)
	var p models.Product
	for cursor.Next(ctx) {
		cursor.Decode(&p)
		listProduct = append(listProduct, p)
	}
	return listProduct, nil
}

func (mongoImpl *ProductRepoImpl) EditProductById(product *models.Product, id string) error {
	ctx := context.Background()
	pUpdate := bson.M{
		"$set": product,
	}
	idP, err := primitive.ObjectIDFromHex(id)
	defer mongoImpl.Db.Client().Disconnect(ctx)
	if err != nil {
		return err
	}
	rs := mongoImpl.Db.Collection("products").FindOneAndUpdate(ctx, bson.M{"_id": bson.M{"$eq": idP}}, pUpdate)
	return rs.Err()
}

func (mongoImpl *ProductRepoImpl) GetProductByName(name string) ([]models.Product, error) {
	listProduct := make([]models.Product, 0)
	ctx := context.Background()
	cursor, err := mongoImpl.Db.Collection("products").Find(ctx, bson.M{"name": name})
	if err != nil {
		return listProduct, err
	}
	defer cursor.Close(ctx)
	var p models.Product
	for cursor.Next(ctx) {
		cursor.Decode(&p)
		listProduct = append(listProduct, p)
	}
	return listProduct, nil
}
