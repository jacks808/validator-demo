package test

import (
	"log"
	"testing"
)

import (
	"github.com/faceair/jio"
)

func Test(t *testing.T) {
	data := []byte(`{
        "debug": "on",
        "window": {
            "title": "Sample Widget",
            "size": [500, 500]
        }
    }`)
	_, err := jio.ValidateJSON(&data, jio.Object().Keys(jio.K{
		"debug": jio.Bool().Truthy("on").Required(),
		"window": jio.Object().Keys(jio.K{
			"title": jio.String().Min(3).Max(18),
			"size":  jio.Array().Items(jio.Number().Integer()).Length(2).Required(),
		}).Without("name", "title").Required(),
	}))
	if err != nil {
		panic(err)
	}
	log.Printf("%s", data) // {"debug":true,"window":{"size":[500,500],"title":"Sample Widget"}}
}
