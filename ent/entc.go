//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/masseelch/elk"
)

func main() {
	exREST, err := elk.NewExtension(
		elk.GenerateSpec("../openapi.json", elk.SpecTitle("Inventory API"), elk.SpecDescription("An API for agriculture inventory management.")),
		elk.GenerateHandlers(),
	)
	if err != nil {
		log.Fatalf("creating elk extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(exREST))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
