package redis

import (
	"context"
	"fmt"

	"acortlink/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(config config.Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Dirección del servidor Redis
		Password: "",               // Contraseña (si no hay, dejar en blanco)
		DB:       0,                // Base de datos a usar (por defecto es 0)
	})

	// Verificar la conexión con un Ping
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error al conectar a Redis:", err)
		panic(err)
	}

	fmt.Println("Connection to redis")

	return client
}
