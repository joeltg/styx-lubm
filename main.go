package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	ld "github.com/underlay/json-gold/ld"
)

var test = regexp.MustCompile("^University\\d+_\\d+\\.owl\\.nt$")

func main() {
	start := time.Now()
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}
	for _, info := range files {
		name := info.Name()
		if test.MatchString(name) {
			fmt.Println("Canonizing", name)
			file, err := os.Open(name)
			if err != nil {
				log.Fatalln(err)
			}
			ns := &ld.NQuadRDFSerializer{}
			dataset, err := ns.Parse(file)
			if err != nil {
				log.Fatalln(err)
			}
			err = file.Close()
			if err != nil {
				log.Fatalln(err)
			}
			na := ld.NewNormalisationAlgorithm("URDNA2015")
			opts := ld.NewJsonLdOptions("")
			opts.Algorithm = "URDNA2015"
			opts.Format = "application/n-quads"
			res, err := na.Main(dataset, opts)
			if err != nil {
				log.Fatalln(err)
			}
			output, err := os.Create("data/" + name)
			if err != nil {
				log.Fatalln(err)
			}
			_, err = output.WriteString(res.(string))
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	fmt.Println("Canonizing completed in", time.Now().Sub(start).Milliseconds(), "milliseconds")
}
