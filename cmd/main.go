package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/garciaolais/challenge-backend/challengelib"
	"github.com/garciaolais/challenge-backend/version"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	versionPtr := flag.Bool("version", false, "get the version number")
	firstPtr := flag.Int("first", 1, "first number in the list")
	lastPtr := flag.Int("last", 100, "last number in the list")
	flag.Parse()

	if *versionPtr {
		fmt.Println(version.Version)
		return 0
	}

	if *firstPtr < 1 {
		log.Fatalf("-first [%d] can't be negative or zero", *firstPtr)
		return 1
	}

	if *lastPtr < 1 {
		log.Fatalf("-last [%d] can't be negative or zero", *lastPtr)
		return 1
	}

	if *firstPtr > *lastPtr {
		log.Fatalf("-firstPtr [%d] can't be greater than -last [%d]", *firstPtr, *lastPtr)
		return 1
	}

	challengelib.Run(*firstPtr, *lastPtr)
	return 0
}
