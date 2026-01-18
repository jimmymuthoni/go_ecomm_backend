package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/anthdm/weavebox"
	"github.com/jimmymuthoni/go_ecomm_backend/api"
	"github.com/jimmymuthoni/go_ecomm_backend/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func handleAPIError(ctx *weavebox.Context, err error){
	fmt.Println("API Error !!!:", err)
	ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
}

func main(){
	app := weavebox.New()
	app.ErrorHandler = handleAPIError

	adminMW := &api.AdminAuthMiddleware{}
	adminRoute := app.Box("/admin")
	adminRoute.Use(adminMW.Authenticate)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongo://localhost:8080"))
	if err != nil {
		panic(err)
	}
	productStore := store.NewMongoProductStore(client.Database("go_ecomm_backend"))
	productHandler := api.NewProductHandler(productStore)

	//adminproduct
	adminProductRoute := adminRoute.Box("/product")
	adminProductRoute.Get("/:id", productHandler.HandleGetProductByID)
	adminProductRoute.Get("/", productHandler.HandleGetProducts)
	adminProductRoute.Post("/", productHandler.HandlePostProduct)
	
	app.Serve(3001)
}