package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderItem struct {
	ProductID string  `json:"product_id" bson:"product_id"`
	Quantity  int     `json:"quantity" bson:"quantity"`
	Price     float64 `json:"price" bson:"price"`
}

type Order struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    string            `json:"user_id" bson:"user_id"`
	Items     []OrderItem       `json:"items" bson:"items"`
	Total     float64           `json:"total" bson:"total"`
	Status    string            `json:"status" bson:"status"`
	CreatedAt time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" bson:"updated_at"`
}

type OrderService struct {
	collection     *mongo.Collection
	productService string
}

func NewOrderService() (*OrderService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return nil, err
	}

	db := client.Database(os.Getenv("DB_NAME"))
	collection := db.Collection("orders")

	return &OrderService{
		collection:     collection,
		productService: os.Getenv("PRODUCT_SERVICE_URL"),
	}, nil
}

func (s *OrderService) validateProducts(items []OrderItem) error {
	for _, item := range items {
		resp, err := http.Get(fmt.Sprintf("%s/api/products/%s", s.productService, item.ProductID))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("product %s not found", item.ProductID)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("error validating product %s", item.ProductID)
		}

		var product struct {
			Stock int     `json:"stock"`
			Price float64 `json:"price"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
			return err
		}

		if product.Stock < item.Quantity {
			return fmt.Errorf("insufficient stock for product %s", item.ProductID)
		}
		item.Price = product.Price
	}
	return nil
}

func (s *OrderService) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

