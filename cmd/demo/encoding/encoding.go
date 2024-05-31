package main

import (
	"fmt"
	"github.com/forestyc/playground/pkg/encoding/base64"
)

func main() {
	urlEncoded := []string{
		`eyJhbGciOiJIUzI1NiJ9`,
		`eyJ1c2VybmFtZSI6ImZyMDAxIiwibG9naW5UeXBlIjoiMyIsIm5ldFRhZyI6IjEiLCJ1c2VyS2V5IjoiY2xpZW50IiwiaXNzIjoiRENFIiwiaWF0IjoxNzE3MTE5MTc5LCJleHAiOjE3MTcyMDU1Nzl9`,
		`OCAT_XP2wpGgGoiCiBW77BOOmVKwhIQ--N-aid-WYJU`,
	}
	b64 := base64.Base64{}
	for _, b := range urlEncoded {
		buf, err := b64.Decode(b)
		if err != nil {
			panic(err)
		}
		fmt.Println(b64.Encode(buf))
	}
}
