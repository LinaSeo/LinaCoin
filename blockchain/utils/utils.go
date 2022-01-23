package utils

import "log"

func HanderErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
