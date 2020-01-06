clean:
	rm -rf data src classes readme.txt log.txt Generator.java canonize

uba1.7.zip:
	curl http://swat.cse.lehigh.edu/projects/lubm/uba1.7.zip -o uba1.7.zip
	unzip uba1.7.zip

src: uba1.7.zip
	unzip uba1.7.zip

GeneratorLinuxFix.zip:
	curl http://swat.cse.lehigh.edu/projects/lubm/GeneratorLinuxFix.zip -o GeneratorLinuxFix.zip

Generator.java: GeneratorLinuxFix.zip
	unzip GeneratorLinuxFix.zip

classes: Generator.java
	cp Generator.java src/edu/lehigh/swat/bench/uba/Generator.java
	javac -d classes src/edu/lehigh/swat/bench/uba/*.java

generate: src classes
	java -cp classes edu.lehigh.swat.bench.uba.Generator -univ 1 -index 0 -seed 0 -onto http://swat.cse.lehigh.edu/onto/univ-bench.owl

data: generate canonize
	mkdir data
	find *.owl -maxdepth 1 -type f -exec sh -c "rapper -i rdfxml -o nquads -q {} > {}.nt" \;
	go run .
	rm University*.owl.nt
	rm University*.owl

root: data
	ipfs add -r -Q --raw-leaves --pin=true --cid-version=1 data | tr -d '\n' > data.cid