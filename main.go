package main

import "fmt"

func main() {

	// var todos []string
	todos := []string{
		"test",
	}

	for {

		var opsi int

		fmt.Println("========= TODO APP CLI ==========")
		fmt.Println("")
		fmt.Println("1. Lihat Todo")
		fmt.Println("2. Tambah Todo")
		fmt.Println("3. Edit Todo")
		fmt.Println("4. Hapus Todo")
		fmt.Println("5. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&opsi)

		switch opsi {
		case 1:
			for i, todo := range todos {
				fmt.Println(i+1, ". ", todo)
			}

		case 2:
			var todoBaru string
			fmt.Print("Tambahkan todo baru: ")
			fmt.Scanln(&todoBaru)
			todos = append(todos, todoBaru)
			fmt.Println("Todo Berahsil Ditambahkan!")
		case 3:
			var index int

			fmt.Print("Update todo: ")
			fmt.Scanln(&index)

			var todoBaru string
			fmt.Print("Masukan todo: ")
			fmt.Scanln(&todoBaru)

			todos[index-1] = todoBaru
			fmt.Println("Todo Berahsil DiUpdate!")
		case 4:
			var index int
			fmt.Print("Hapus todo: ")
			fmt.Scanln(&index)
			todos = append(todos[:index-1], todos[index:]...)
			fmt.Println("Todo Berahsil Dihapus!")
		case 5:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}

	}
}
