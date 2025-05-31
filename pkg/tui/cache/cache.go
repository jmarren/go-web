package cache

import (
	"context"
	"fmt"

	"github.com/valkey-io/valkey-go"
)

type Cache struct{}

func (c *Cache) Connect() {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{"127.0.0.1:300"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()
	// SET key val NX
	err = client.Do(ctx, client.B().Set().Key("key").Value("val").Build()).Error()
	// if err != nil {
	// 	panic(err)
	// }
	// HGETALL hm
	val, err := client.Do(ctx, client.B().Get().Key("key").Build()).AsStrSlice()
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Printf("key: %v\n", val)
}
