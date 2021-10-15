package second

import (
	"io"
	"log"
	"os"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/3 17:29
 * @description: err and err deal
 */

func close(f io.Closer) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func CsErr1Main() {
	r, err := os.Open("a")
	if err != nil {
		log.Fatalf("Open file a failed\n")
	}
	defer close(r)

	r, err = os.Open("b")
	if err != nil {
		log.Fatalf("Open file b failed\n")
	}
	defer close(r)
}
