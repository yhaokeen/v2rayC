package db

import (
	"log"
)

func GetServers() []ServerRaw {
	r := make([]ServerRaw, 0)
	raw, err := ListGetAll("touch", "servers")
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

func GetSubscriptions() []SubscriptionRaw {
	r := make([]SubscriptionRaw, 0)
	raw, err := ListGetAll("touch", "subscriptions")
	if err == nil {
		for _, b := range raw {
			t, e := Bytes2SubscriptionRaw(b)
			if e != nil {
				log.Warn("%v", e)
				continue
			}
			r = append(r, *t)
		}
	}
	return r
}
