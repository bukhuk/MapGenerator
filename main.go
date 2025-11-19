package main

import (
	"MapGenerator/Generation"
	"flag"
	"fmt"
	"time"
)

func main() {
	n := flag.Int("n", 256, "Number of strings")
	m := flag.Int("m", 256, "Number of columns")
	seed := flag.Int("seed", 0, "Generation seed")

	flag.Parse()

	start := time.Now()

	Generation.Generate(*n, *m, *seed)

	duration := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", duration)
}
