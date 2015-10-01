# delphi-course 2015 bmstu

Dependencies for generator:
```
go get gopkg.in/alecthomas/kingpin.v2
go get gopkg.in/yaml.v2
```

Run generator:
```go build gen.go && ./gen --base=base1.yml --variant 10  > out && pdflatex out && evince out.pdf```

Generate answers:
```> go build gen.go && ./gen --base=base1.yml --variant 10 --answers
1: 2
2: 1
3: 3 4
4: 1 2 3
5: 1
6: 2
7: 1
8: 1
9: 3
10: 3
```

Print 22 different versions that fits in two pages:
```
go build gen.go
shift=0
for i in {1..22}; do
	variant=$((i + shift))
	./gen --base=base1.yml --variant $variant > out && pdflatex out
	while [[ $(pdfinfo out.pdf | grep Pages | perl -p -E 's/Pages:\s+//') != '2' ]]; do
		shift=$((shift + 1))
		variant=$((i + shift))
		./gen --base=base1.yml --variant $variant > out && pdflatex out
	done
	lp out.pdf -o sides=two-sided-long-edge
done
```