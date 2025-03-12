package redis

import (
	"fmt"
	"log"
)

func NewRedisCacheClient(attr InitAttribute) Cache {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	client := &RedisClient{
		cache: attr.Client,
	}

	return client
}

func (init InitAttribute) validate() error {
	if !init.Client.validate() {
		return fmt.Errorf("missing mq client : %+v", init.Client)
	}

	return nil
}

func (client Client) validate() bool {
	return client.Client != nil
}
