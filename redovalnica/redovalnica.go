package redovalnica

// Package redovalnica provides gradebook functionality.
//
// Exported functions:
// - DodajOceno: adds a grade to the gradebook.
// - IzpisVsehOcen: returns a slice of all grades
// - IzpisiKoncniUspeh: computes and returns the final result message
//   based on average and minimum required number of grades
//
// The function povprecje is unexported.
import "fmt"

var ocene []int

func DodajOceno(grade, minOcena, maxOcena int) bool {
	if grade < minOcena || grade > maxOcena {
		return false
	}
	ocene = append(ocene, grade)
	return true
}

func IzpisVsehOcen() []int {
	fmt.Println(ocene)
	return ocene
}

func IzpisiKoncniUspeh(stOcen int) string {
	if len(ocene) < stOcen {
		return "Premalo ocen za koncno oceno."
	}
	avg := povprecje(ocene)
	if avg >= 2.0 {
		return "Pozitiven uspeh (povprecje: " + fmt.Sprintf("%.2f", avg) + ")"
	}
	return "Negativen uspeh (povprecje: " + fmt.Sprintf("%.2f", avg) + ")"
}

func povprecje(vals []int) float64 {
	if len(vals) == 0 {
		return 0
	}
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return float64(sum) / float64(len(vals))
}
