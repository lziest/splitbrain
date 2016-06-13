package main

import (
	"flag"
	"fmt"
	"github.com/lziest/funkybrain"
	"github.com/lziest/mood"
)

func main() {
	flag.StringVar(&mood.Status, "mood", "happy", "put in your mood")
	flag.Parse()
	fmt.Printf("It appears you can change your mood to what you want: ")
	mood.Show()
	fmt.Println("Deep inside the funkybrain, it never changes: I'm", funkybrain.Mood())
}
