package dependencies

import (
	"context"
	"github.com/henriquegmendes/go-base-api-rest/client"
	"github.com/henriquegmendes/go-base-api-rest/repository"
	"github.com/henriquegmendes/go-base-api-rest/service"
	"log"
)

type GlobalDeps struct {
	MongoDep       *client.MongoDep
	ExampleService service.ExampleService
}

func LoadGlobalDependencies(ctx context.Context) *GlobalDeps {
	mongoDep, err := client.NewMongoClient(ctx)
	if err != nil {
		log.Fatalf("error loading mongodb client. Error: %s", err.Error())
	}

	exampleRepository := repository.NewExampleRepository(mongoDep.ExampleDatabase)
	exampleService := service.NewExampleService(exampleRepository)

	return &GlobalDeps{
		MongoDep:       mongoDep,
		ExampleService: exampleService,
	}
}
