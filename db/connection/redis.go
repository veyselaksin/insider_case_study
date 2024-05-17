package connection

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
)

func initializeRedis() *redis.Client {

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalln("REDIS_DB is not a number")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       db,
	})

	return rdb
}
