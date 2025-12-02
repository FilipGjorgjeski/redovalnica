# Redovalnica

CLI aplikacija in Go paket za upravljanje ocen.

## Namestitev

```bash
go get github.com/FilipGjorgjeski/redovalnica
```

## Uporaba CLI

Globalna stikala morajo biti podana pred podukazom.

```bash
# Gradnja
mkdir -p bin
go build -o bin/redovalnica ./cmd/redovalnica

# Dodaj ocene
./bin/redovalnica --minOcena 1 --maxOcena 5 dodaj 3
./bin/redovalnica dodaj 5

# Izpis vseh ocen
./bin/redovalnica izpisi

# Koncni uspeh (globalni flagi pred podukazom)
./bin/redovalnica --stOcen 2 uspeh
```

### Stikala
- `--stOcen` (globalno): najmanjše število ocen za pozitivno oceno (privzeto 3)
- `--minOcena` (globalno): najmanjša možna ocena (privzeto 1)
- `--maxOcena` (globalno): največja možna ocena (privzeto 5)

## Paket `redovalnica`

Izvožene funkcije:
- `DodajOceno(grade, minOcena, maxOcena int) bool`
- `IzpisVsehOcen() []int`
- `IzpisiKoncniUspeh(stOcen int) string`

Skrita funkcija:
- `povprecje([]int) float64`

## Licenca
MIT
