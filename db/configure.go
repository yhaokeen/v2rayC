package db

import (
	"log"
)

func GetServers() []ServerRaw {
	r := make([]ServerRaw, 0)
	raw, err := GetBucketAll("server")
	if err == nil {
		for _, b := range raw {
			t, e := Bytes2ServerRaw(b)
			if e != nil {
				log.Warn("GetServers: %v", e)
				continue
			}
			r = append(r, *t)
		}
	}
	return r
}

func GetSubscriptions() map[string]SubscriptionRaw {
	r := make(map[string]SubscriptionRaw)
	raw, err := GetBucketAll("subscription")
	if err == nil {
		for k, b := range raw {
			t, e := Bytes2SubscriptionRaw(b)
			if e != nil {
				log.Warn("%v", e)
				continue
			}
			r[k] = *t
		}
	}
	return r
}
