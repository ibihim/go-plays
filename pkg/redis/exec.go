package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Exec performs some redis operations
func Exec() error {
	conn, err := Setup()
	if err != nil {
		return err
	}

	cl := "ibihim was here"
	// value is an interface, we can store whatever
	// the last argument is the redis expiration
	conn.Set("key", cl, 5*time.Second)

	var result string
	if err := conn.Get("key").Scan(&result); err != nil {
		switch err {
		// this means the key
		// was not found
		case redis.Nil:
			return nil
		default:
			return err
		}
	}

	fmt.Println("result =", result)

	return nil
}
