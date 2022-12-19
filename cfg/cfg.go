package cfg

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type Environment struct {
	Port            int           `envconfig:"PORT"`
	MongodbURI      string        `envconfig:"MONGODB_URI"`
	MongodbDatabase string        `envconfig:"MONGODB_DATABASE"`
	Example         time.Duration `envconfig:"EXAMPLE"`
}

var env *Environment

func Env() *Environment {
	if env == nil {
		err := godotenv.Load()
		if err != nil {
			log.Printf("error loading .env file: %s", err.Error())
		}

		var newEnv Environment
		err = envconfig.Process("", &newEnv)
		if err != nil {
			log.Fatalf("error mapping envs into Environment struct: %s", err.Error())
		}

		env = &newEnv
	}

	return env
}
