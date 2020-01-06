package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"

	cid "github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	ipfs "github.com/ipfs/go-ipfs-http-client"
	path "github.com/ipfs/interface-go-ipfs-core/path"
	ld "github.com/underlay/json-gold/ld"

	styx "github.com/underlay/styx/db"
)

var httpAPI *ipfs.HttpApi

func init() {
	var err error
	httpAPI, err = ipfs.NewURLApiWithClient("http://localhost:5001", http.DefaultClient)
	if err != nil {
		log.Fatalln(err)
	}
}

var university = regexp.MustCompile("^University\\d+_\\d+\\.owl\\.c\\.nt$")

func TestIngestLubm(t *testing.T) {
	// Get root CID
	rootCid, err := ioutil.ReadFile("data.cid")
	if err != nil {
		t.Error(err)
		return
	}

	root, err := cid.Decode(string(rootCid))
	if err != nil {
		t.Error(err)
		return
	}

	// Remove old db
	fmt.Println("removing path", styx.DefaultPath)
	err = os.RemoveAll(styx.DefaultPath)
	if err != nil {
		t.Fatal(err)
	}

	db, err := styx.OpenDB(styx.DefaultPath, nil)
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	ctx := context.Background()

	p := path.IpfsPath(root)
	node, err := httpAPI.Unixfs().Get(ctx, p)
	if err != nil {
		t.Error(err)
		return
	}

	object := httpAPI.Object()
	dir := files.ToDir(node)
	start := time.Now()
	for iter := dir.Entries(); iter.Next(); {
		file, name := files.ToFile(iter.Node()), iter.Name()
		stat, err := object.Stat(ctx, path.Join(p, name))
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println("Inserting", name, stat.Cid.String())
		dataset, err := ld.ParseQuads(file)
		if err != nil {
			t.Error(err)
			return
		}
		err = db.Insert(stat.Cid, dataset)
		if err != nil {
			t.Error(err)
			return
		}
	}
	duration := time.Now().Sub(start).Milliseconds()
	fmt.Println("Inserted data in", duration, "milliseconds")
}

func TestQuery1(t *testing.T) {
	err := testQuery("query1.json")
	if err != nil {
		t.Error(err)
	}
}

func TestQuery2(t *testing.T) {
	err := testQuery("query2.json")
	if err != nil {
		t.Error(err)
	}
}

func TestQuery3(t *testing.T) {
	err := testQuery("query3.json")
	if err != nil {
		t.Error(err)
	}
}

func testQuery(name string) error {
	query, err := getQuery(name)
	if err != nil {
		return err
	}

	db, err := styx.OpenDB(styx.DefaultPath, nil)
	if err != nil {
		return err
	}

	defer db.Close()

	start := time.Now()
	cursor, err := db.Query(query, nil, nil)
	if err != nil {
		return err
	}

	if cursor != nil {
		defer cursor.Close()
		for d := cursor.Domain(); d != nil; d, err = cursor.Next(nil) {
			for _, b := range d {
				fmt.Printf("%s: %s\n", b.Attribute, cursor.Get(b).GetValue())
			}
			fmt.Println("-----")
		}
	} else {
		fmt.Println("No results")
	}
	elapsed := time.Now().Sub(start)
	log.Println("Iteration completed in", elapsed.Microseconds(), "microseconds")
	return nil
}

func getQuery(name string) ([]*ld.Quad, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	var query map[string]interface{}
	err = json.NewDecoder(file).Decode(&query)
	if err != nil {
		return nil, err
	}

	proc := ld.NewJsonLdProcessor()
	opts := ld.NewJsonLdOptions("")
	opts.ProduceGeneralizedRdf = true
	rdf, err := proc.ToRDF(query, opts)
	if err != nil {
		return nil, err
	}
	return rdf.(*ld.RDFDataset).GetQuads("@default"), nil
}
