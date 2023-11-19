package cli

import (
	"log"
	"strings"
)

func ParseKVFlagValue(value string) map[string]string {
	kvs := map[string]string{}
	records := strings.Split(value, ",")

	for _, record := range records {
		kv := strings.Split(record, "=")
		if len(kv) != 2 {
			log.Printf("invalid flag passed: %s\n", record)
			continue
		}

		key := kv[0]
		val := kv[1]
		kvs[key] = val
	}

	return kvs
}
