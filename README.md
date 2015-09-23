# delphi-course 2015 bmstu

Dependencies for generator:
```
go get gopkg.in/alecthomas/kingpin.v2
go get gopkg.in/yaml.v2
```

Run generator:
```go build gen.go && ./gen --base=base1.yml  > out && pdflatex out && evince out.pdf```

Generate answers page:
```go build gen.go && ./gen --base=base1.yml --answers > out && pdflatex out && evince out.pdf```

Common way with check-pages:
```
go build gen.go
./gen --base=base1.yml > out && pdflatex out
while [[ $(pdfinfo out.pdf | grep Pages | perl -p -E 's/Pages:\s+//') != '2' ]]; do
	./gen --base=base1.yml > out && pdflatex out
done
evince out.pdf
```

Common way with check-pages:
```
go build gen.go
./gen --base=base1.yml > out && pdflatex out
while [[ $(pdfinfo out.pdf | grep Pages | perl -p -E 's/Pages:\s+//') != '2' ]]; do
	./gen --base=base1.yml > out && pdflatex out
done
evince out.pdf
```

Print 22 different versions:
```
go build gen.go
for i in {1..22}; do
	./gen --base=base1.yml > out && pdflatex out
	while [[ $(pdfinfo out.pdf | grep Pages | perl -p -E 's/Pages:\s+//') != '2' ]]; do
		./gen --base=base1.yml > out && pdflatex out
	done
	lp out.pdf
done
```