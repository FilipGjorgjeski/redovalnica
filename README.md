# Redovalnica

CLI za upravljanje ocen.

## Namestitev

```bash
go get github.com/FilipGjorgjeski/redovalnica
```

## Uporaba CLI

```bash
# Gradnja
mkdir -p bin
go build -o bin/redovalnica ./cmd/redovalnica

# Dodaj ocene
./bin/redovalnica --minOcena 1 --maxOcena 5 dodaj 3
./bin/redovalnica dodaj 5

# Izpis vseh ocen
./bin/redovalnica izpisi

# Koncni uspeh
./bin/redovalnica --stOcen 2 uspeh
```

### Stikala
- `--stOcen`: najmanjše število ocen za pozitivno oceno
- `--minOcena`: najmanjša možna ocena
- `--maxOcena`: največja možna ocena

## Paket `redovalnica`

Izvožene funkcije:
- `DodajOceno(grade, minOcena, maxOcena int) bool`
- `IzpisVsehOcen() []int`
- `IzpisiKoncniUspeh(stOcen int) string`
