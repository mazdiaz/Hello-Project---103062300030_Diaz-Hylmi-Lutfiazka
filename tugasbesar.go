// ini adalah baris komentar baru

// ini juga baris baru 
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Nilai struct {
	uts, uas, quiz, total float64
	grade                 string
}

type Matakuliah struct {
	subject string
	grades  Nilai
}

type Mahasiswa struct {
	nama       string
	nim        string
	sks        int
	nilai      float64
	matakuliah []Matakuliah
}

func printAsciiArt() {
	asciiArt := `
  _____ _____ _     _  _____  __  __           _   _ _   _ _____     _______ ____  ____ ___ _______   __
 |_   _| ____| |   | |/ / _ \|  \/  |         | | | | \ | |_ _\ \   / / ____|  _ \/ ___|_ _|_   _\ \ / /
   | | |  _| | |   | ' / | | | |\/| |  _____  | | | |  \| || | \ \ / /|  _| | |_) \___ \| |  | |  \ V / 
   | | | |___| |___| . \ |_| | |  | | |_____| | |_| | |\  || |  \ V / | |___|  _ < ___) | |  | |   | |  
   |_| |_____|_____|_|\_\___/|_|  |_|          \___/|_| \_|___|  \_/  |_____|_| \_\____/___| |_|   |_|                                                                         
	`
	fmt.Println(asciiArt)
}

//sdsd d

func menu() {
	printAsciiArt()
	fmt.Println("-------------------------------------------------")
	fmt.Println("|1. Data Seluruh Mahasiswa                      |")
	fmt.Println("|2. Input Data Mahasiswa Baru                   |")
	fmt.Println("|3. Edit Data Mahasiswa                         |")
	fmt.Println("|4. Delete Data Mahasiswa                       |")
	fmt.Println("|5. Display Transkrip Mahasiswa                 |")
	fmt.Println("|6. Tampilkan Mahasiswa Berdasarkan Mata Kuliah |")
	fmt.Println("|7. Tampilkan Mata Kuliah Berdasarkan Mahasiswa |")
	fmt.Println("|8. Tampilkan Mahasiswa Secara Terurut          |")
	fmt.Println("|0. Exit program                                |")
	fmt.Println("-------------------------------------------------")
	printAsciiArt()
	fmt.Printf("Enter Your Choice: ")
}

func main() {
	var students []Mahasiswa
	running := true
	for running {
		var op int
		menu()
		fmt.Scan(&op)
		bufio.NewScanner(os.Stdin).Scan()
		switch op {
		case 1:
			displayAll(students)
		case 2:
			inputData(&students)
		case 3:
			editData(students)
		case 4:
			deleteSubject(&students)
		case 5:
			displayTranscript(students)
		case 6:
			displayStudentsBySubject(students)
		case 7:
			displaySubjectsByStudent(students)
		case 8:
			if len(students) == 0 {
				fmt.Printf("Data Kosong!\n")
			} else {
				rocket_sort(students)
			}
		case 0:
			saveToFile(students, "students.json")
			running = false
		default:
			fmt.Println("INVALID INPUT!")
		}
	}
}

func rocket_sort(arr []Mahasiswa) {
	students := make([]Mahasiswa, len(arr))
	copy(students, arr)
	n := len(students)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if students[j].sks < students[j+1].sks {
				students[j], students[j+1] = students[j+1], students[j]
			} else if students[j].sks == students[j+1].sks {
				if students[j].nilai < students[j+1].nilai {
					students[j], students[j+1] = students[j+1], students[j]
				}
			}
		}
	}
	i := 1
	for _, student := range students {
		fmt.Printf("%d. Nama: %s, NIM: %s, SKS: %d, Nilai: %.2f\n", i, student.nama, student.nim, student.sks, student.nilai)
		i++
	}
}

func displayTranscript(students []Mahasiswa) {
	var NIM string
	fmt.Printf("Masukkan NIM mahasiswa: ")
	fmt.Scan(&NIM)

	for _, student := range students {
		if student.nim == NIM {
			fmt.Printf("Transkrip Nilai untuk %s (NIM: %s):\n", student.nama, student.nim)
			fmt.Printf("Total SKS: %d\n", student.sks)
			for _, matkul := range student.matakuliah {
				fmt.Printf("Mata Kuliah: %s\n", matkul.subject)
				fmt.Printf("Nilai UTS: %.2f\n", matkul.grades.uts)
				fmt.Printf("Nilai UAS: %.2f\n", matkul.grades.uas)
				fmt.Printf("Nilai Quiz: %.2f\n", matkul.grades.quiz)
				fmt.Printf("Total nilai: %.2f\n", matkul.grades.total)
				fmt.Printf("Grade: %s\n", matkul.grades.grade)
				fmt.Println()
			}
			return
		}
	}
	fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan!")
}

func deleteSubject(students *[]Mahasiswa) {
	var NIM string
	fmt.Printf("Masukkan NIM mahasiswa yang ingin diubah: ")
	fmt.Scan(&NIM)
	for i := range *students {
		if (*students)[i].nim == NIM {
			fmt.Printf("Menghapus mata kuliah untuk %s\n", (*students)[i].nama)

			fmt.Println("Daftar mata kuliah:")
			for j, matkul := range (*students)[i].matakuliah {
				fmt.Printf("%d. %s\n", j+1, matkul.subject)
			}

			var index int
			fmt.Printf("Pilih nomor mata kuliah yang ingin dihapus: ")
			fmt.Scan(&index)

			if index > 0 && index <= len((*students)[i].matakuliah) {
				index--
				(*students)[i].matakuliah = append((*students)[i].matakuliah[:index], (*students)[i].matakuliah[index+1:]...)
				(*students)[i].sks -= 2
				fmt.Println("Mata kuliah telah dihapus.")
			} else {
				fmt.Println("Nomor mata kuliah tidak valid.")
			}
			totalNilai := 0.0
			for _, matkul := range (*students)[i].matakuliah {
				totalNilai += matkul.grades.total
			}
			(*students)[i].nilai = totalNilai / float64(len((*students)[i].matakuliah))
			return
		}
	}
	fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan!")
}

func displayStudentsBySubject(students []Mahasiswa) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Masukkan nama mata kuliah: ")
	scanner.Scan()
	subject := scanner.Text()
	found := false
	i := 1
	for _, student := range students {
		for _, matkul := range student.matakuliah {
			if matkul.subject == subject {
				fmt.Printf("%d. Nama: %s, NIM: %s\n", i, student.nama, student.nim)
				found = true
				i++
			}
		}
	}
	if !found {
		fmt.Println("Tidak ada mahasiswa yang mengambil mata kuliah tersebut.")
	}
}

func displaySubjectsByStudent(students []Mahasiswa) {
	var NIM string
	fmt.Printf("Masukkan NIM mahasiswa: ")
	fmt.Scan(&NIM)
	found := false
	for _, student := range students {
		if student.nim == NIM {
			fmt.Printf("Nama: %s, NIM: %s\n", student.nama, student.nim)
			fmt.Println("Daftar Mata Kuliah:")
			for i, matkul := range student.matakuliah {
				fmt.Printf("%d. %s\n", i+1, matkul.subject)
			}
			found = true
		}
	}
	if !found {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
	}
}

func editData(students []Mahasiswa) {
	var NIM string
	fmt.Printf("Masukkan NIM mahasiswa yang ingin diubah: ")
	fmt.Scan(&NIM)
	for i := range students {
		if students[i].nim == NIM {
			fmt.Printf("Mengubah data untuk %s\n", students[i].nama)

			fmt.Println("Daftar mata kuliah:")
			for j, matkul := range students[i].matakuliah {
				fmt.Printf("%d. %s\n", j+1, matkul.subject)
			}

			var subjectIndex int
			fmt.Printf("Pilih nomor mata kuliah yang ingin diubah: ")
			fmt.Scan(&subjectIndex)

			if subjectIndex > 0 && subjectIndex <= len(students[i].matakuliah) {
				subjectIndex--

				fmt.Printf("Input nilai UTS untuk %s: ", students[i].matakuliah[subjectIndex].subject)
				fmt.Scan(&students[i].matakuliah[subjectIndex].grades.uts)
				fmt.Printf("Input nilai UAS untuk %s: ", students[i].matakuliah[subjectIndex].subject)
				fmt.Scan(&students[i].matakuliah[subjectIndex].grades.uas)
				fmt.Printf("Input nilai Quiz untuk %s: ", students[i].matakuliah[subjectIndex].subject)
				fmt.Scan(&students[i].matakuliah[subjectIndex].grades.quiz)

				students[i].matakuliah[subjectIndex].grades.total = (students[i].matakuliah[subjectIndex].grades.uts + students[i].matakuliah[subjectIndex].grades.uas + students[i].matakuliah[subjectIndex].grades.quiz) / 3
				students[i].matakuliah[subjectIndex].grades.grade = calculateGrade(students[i].matakuliah[subjectIndex].grades.total)

				fmt.Println("Data mata kuliah telah diubah.")
				totalNilai := 0.0
				for _, matkul := range students[i].matakuliah {
					totalNilai += matkul.grades.total
				}
				students[i].nilai = totalNilai / float64(len(students[i].matakuliah))
			} else {
				fmt.Println("Nomor mata kuliah tidak valid.")
			}
			return
		}
	}
	fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan!")
}

func inputData(students *[]Mahasiswa) {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	fmt.Printf("Input jumlah mahasiswa: ")
	fmt.Scan(&n)
	bufio.NewScanner(os.Stdin).Scan()
	for i := 0; i < n; i++ {
		var student Mahasiswa
		var subjectCount int
		fmt.Printf("Input Nama Mahasiswa #%d: ", i+1)
		scanner.Scan()
		student.nama = scanner.Text()
		fmt.Printf("Input NIM: ")
		scanner.Scan()
		student.nim = scanner.Text()
		fmt.Printf("Input jumlah mata kuliah untuk %s: ", student.nama)
		fmt.Scan(&subjectCount)
		student.sks += subjectCount * 2
		bufio.NewScanner(os.Stdin).Scan()
		totalNilai := 0.0
		for j := 0; j < subjectCount; j++ {
			var matkul Matakuliah
			fmt.Printf("Input nama mata kuliah #%d: ", j+1)
			scanner.Scan()
			matkul.subject = scanner.Text()
			fmt.Printf("Input nilai UTS untuk %s: ", matkul.subject)
			fmt.Scan(&matkul.grades.uts)
			bufio.NewScanner(os.Stdin).Scan()
			fmt.Printf("Input nilai UAS untuk %s: ", matkul.subject)
			fmt.Scan(&matkul.grades.uas)
			bufio.NewScanner(os.Stdin).Scan()
			fmt.Printf("Input nilai Quiz untuk %s: ", matkul.subject)
			fmt.Scan(&matkul.grades.quiz)
			bufio.NewScanner(os.Stdin).Scan()
			matkul.grades.total = (matkul.grades.uts + matkul.grades.uas + matkul.grades.quiz) / 3
			matkul.grades.grade = calculateGrade(matkul.grades.total)
			totalNilai += matkul.grades.total
			student.matakuliah = append(student.matakuliah, matkul)
		}
		student.nilai = totalNilai / float64(subjectCount)
		*students = append(*students, student)
	}
}

func displayAll(students []Mahasiswa) {
	if len(students) == 0 {
		fmt.Println("Data Kosong!")
		return
	}
	for i, student := range students {
		fmt.Printf("Mahasiswa #%d:\n", i+1)
		fmt.Println("Nama:", student.nama)
		fmt.Println("NIM:", student.nim)
		fmt.Printf("Total SKS: %d\n", student.sks)
		for _, matkul := range student.matakuliah {
			fmt.Printf("Mata Kuliah: %s\n", matkul.subject)
			fmt.Printf("Nilai UTS: %.2f\n", matkul.grades.uts)
			fmt.Printf("Nilai UAS: %.2f\n", matkul.grades.uas)
			fmt.Printf("Nilai Quiz: %.2f\n", matkul.grades.quiz)
			fmt.Printf("Total nilai: %.2f\n", matkul.grades.total)
			fmt.Printf("Grade: %s\n", matkul.grades.grade)
		}
	}
}

func calculateGrade(total float64) string {
	switch {
	case total >= 85:
		return "A"
	case total >= 70:
		return "B"
	case total >= 55:
		return "C"
	case total >= 40:
		return "D"
	default:
		return "E"
	}
}

func saveToFile(students []Mahasiswa, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(students)
	if err != nil {
		fmt.Println("Error encoding data:", err)
	}
}
