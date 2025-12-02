package main

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func dodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Student s to vpisno številko ne obstaja.")
		return
	}
	if ocena < 1 || ocena > 10 {
		fmt.Println("Ocena mora biti med 1 in 10.")
		return
	}
	student.ocene = append(student.ocene, ocena)
	studenti[vpisnaStevilka] = student
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}
	if len(student.ocene) < 6 {
		return 0.0
	}
	sum := 0.0
	for _, ocena := range student.ocene {
		sum += float64(ocena)
	}
	return sum / float64(len(student.ocene))
}

func izpisRedovalnice(studenti map[string]Student) {
	fmt.Printf("REDOVALNICA:\n")
	for vpisnaStevilka, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisnaStevilka, student.ime, student.priimek, student.ocene)
	}
}

func izpisiKoncniUspeh(studenti map[string]Student) {
	for vpisnaStevilka, student := range studenti {
		povp := povprecje(studenti, vpisnaStevilka)
		fmt.Printf("%s %s: povprečna ocena %f -> ", student.ime, student.priimek, povp)
		if povp < 6.0 {
			fmt.Printf("Neuspešen študent\n")
		} else if povp >= 6.0 && povp < 9.0 {
			fmt.Printf("Povprečen študent\n")
		} else {
			fmt.Printf("Odličen študent\n")
		}
	}
}

func main() {
	var studenti map[string]Student = make(map[string]Student)
	studenti["63230000"] = Student{"File", "Zemen", []int{10, 10, 10, 10, 10}}
	studenti["63230001"] = Student{"Rajko", "Drvar", []int{8, 9, 10, 7, 9, 7, 8}}
	studenti["63230002"] = Student{"Trigli", "Ceridi", []int{6, 4, 2, 6, 4}}

	dodajOceno(studenti, "63230000", 9)
	fmt.Println(studenti["63230000"])
	dodajOceno(studenti, "63230001", 11)
	dodajOceno(studenti, "63230101", 10)

	fmt.Println(povprecje(studenti, "63230000"))
	fmt.Println(povprecje(studenti, "63230001"))
	fmt.Println(povprecje(studenti, "63230002"))
	fmt.Println(povprecje(studenti, "63230101"))

	izpisRedovalnice(studenti)

	izpisiKoncniUspeh(studenti)

}
