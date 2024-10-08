package main

import (
	"fmt"
	"os"
	"strings"
)

type SparePart struct {
	ID       int
	Name     string
	Quantity int
	Harga    float64
}

type Customer struct {
	ID            int
	Name          string
	ServicesDates int
}

type Transaction struct {
	ID          int
	CustomerID  int
	SparePart   string
	Date        string
	ServiceRate int
}

const MaxSpareParts = 100
const MaxCustomers = 100
const MaxTransactions = 100

var (
	spareParts     [MaxSpareParts]SparePart
	customers      [MaxCustomers]Customer
	transactions   [MaxTransactions]Transaction
	nextPartID     int
	nextCustomerID int
	nextTransID    int
)

func tambahSparePart(id int, name string, quantity int) {
	if nextPartID >= MaxSpareParts {
		fmt.Println("Kapasitas Spare Parts sudah mencapai batas maksimum")
	} else {
		sparePart := SparePart{
			ID:       id,
			Name:     name,
			Quantity: quantity,
		}
		spareParts[nextPartID] = sparePart
		nextPartID++
		fmt.Println("Spare-Part berhasil ditambahkan")
	}
}

func tutupApp() {
	fmt.Println("Terima kasih!")
	os.Exit(0)
}

func cariSparePartByID(id int) *SparePart {
	for i := 0; i < nextPartID; i++ {
		if spareParts[i].ID == id {
			return &spareParts[i]
		}
	}
	return nil
}

func editSparePart(id int, name string, quantity int) bool {
	for i := 0; i < nextPartID; i++ {
		if spareParts[i].ID == id {
			spareParts[i].Name = name
			spareParts[i].Quantity = quantity
			return true
		}
	}
	return false
}

func hapusSparePart(id int) bool {
	for i := 0; i < nextPartID; i++ {
		if spareParts[i].ID == id {
			for j := i; j < nextPartID-1; j++ {
				spareParts[j] = spareParts[j+1]
			}
			nextPartID--
			return true
		}
	}
	return false
}

func tambahPelanggan(id int, name string) {
	var servicesdates string

	fmt.Print("Tanggal (dd/mm): ")
	fmt.Scanln(&servicesdates)

	dateParts := strings.Split(servicesdates, "/")
	dayStr := dateParts[0]
	monthStr := dateParts[1]

	var day, month int
	_, err := fmt.Sscanf(dayStr, "%d", &day)
	if err != nil {
		fmt.Println("Format tanggal tidak valid")
		tambahPelanggan(id, name)
	} else if _, err := fmt.Sscanf(monthStr, "%d", &month); err != nil {
		fmt.Println("Format tanggal tidak valid")
		tambahPelanggan(id, name)
	} else {
		servicesDate := (day * 100) + month

		if nextCustomerID >= MaxCustomers {
			fmt.Println("Kapasitas Pelanggan sudah mencapai batas maksimum")
		} else {
			customer := Customer{
				ID:            id,
				Name:          name,
				ServicesDates: servicesDate,
			}
			customers[nextCustomerID] = customer
			nextCustomerID++
			fmt.Println("Pelanggan berhasil ditambahkan")
		}
	}
}

func editPelanggan(id int, name string) bool {
	for i := 0; i < nextCustomerID; i++ {
		if customers[i].ID == id {
			customers[i].Name = name
			return true
		}
	}
	return false
}

func hapusPelanggan(id int) bool {
	for i := 0; i < nextCustomerID; i++ {
		if customers[i].ID == id {
			for j := i; j < nextCustomerID-1; j++ {
				customers[j] = customers[j+1]
			}
			nextCustomerID--
			return true
		}
	}
	return false
}

func tambahTransaksi(id int, customerID int, sparePart string) {
	if nextTransID >= MaxTransactions {
		fmt.Println("Kapasitas Transaksi sudah mencapai batas maksimum")
	} else {
		transaction := Transaction{
			ID:         id,
			CustomerID: customerID,
			SparePart:  sparePart,
		}
		transactions[nextTransID] = transaction
		nextTransID++
		fmt.Println("Transaksi berhasil ditambahkan")
	}
}

func inputDariUser() string {
	var input string
	fmt.Print("Masukkan input tambahan: ")
	fmt.Scanln(&input)
	return input
}

func hitungServiceRateLogic(customerName string, sparePartName string, additionalInput string) int {
	serviceRate := 0

	if customerName == "Azet" {
		serviceRate += 50000
	} else {
		serviceRate += 100000
	}

	if sparePartName == "Oli" {
		serviceRate += 20000
	} else {
		serviceRate += 300000
	}

	additionalFactor := 1
	if additionalInput == "Urgent" {
		additionalFactor = 2
	}

	serviceRate *= additionalFactor

	return serviceRate
}

func hitungTotalServiceRate(serviceRate int) int {
	totalRate := serviceRate + 10000

	return totalRate
}

func hitungServiceRate(transactionID int) {
	var customerName string
	var customerID int
	var sparePartName string
	var transactionFound bool

	found := false

	for _, trans := range transactions {
		if !found && trans.ID == transactionID {
			customer := cariPelangganByID(trans.CustomerID)
			if customer != nil {
				customerName = customer.Name
			}

			sparePart := cariSparePartByNama(trans.SparePart)
			if sparePart != nil {
				sparePartName = sparePart.Name
			}

			transactionFound = true

			additionalInput := inputDariUser()
			serviceRate := hitungServiceRateLogic(customerName, sparePartName, additionalInput)

			totalRate := hitungTotalServiceRate(serviceRate)

			fmt.Printf("Transaksi ditemukan\n")
			fmt.Println("ID Pelanggan: ")
			fmt.Scanln(&customerID)
			fmt.Println("Nama Pelanggan: ")
			fmt.Scanln(&customerName)
			fmt.Println("Nama Sparepart: ")
			fmt.Scanln(&sparePartName)

			fmt.Printf("Tarif Layanan: %d\n", totalRate)

			found = true
		}
	}

	if !transactionFound {
		fmt.Println("Transaksi tidak ditemukan")
	}

}

func cariPelangganByDate(startDate string) {
	fmt.Printf("Daftar pelanggan yang melakukan service pada tanggal %s:\n", startDate)

	dateParts := strings.Split(startDate, "/")
	dayStr := dateParts[0]
	monthStr := dateParts[1]

	var day, month int
	_, err := fmt.Sscanf(dayStr, "%d", &day)
	if err != nil {
		fmt.Println("Format tanggal tidak valid")
		cariPelangganByDate(startDate)
	} else if _, err := fmt.Sscanf(monthStr, "%d", &month); err != nil {
		fmt.Println("Format tanggal tidak valid")
		cariPelangganByDate(startDate)
	} else {
		searchDate := (day * 100) + month

		found := false

		for _, customer := range customers {
			if customer.ServicesDates == searchDate {
				fmt.Println(customer.Name)
				found = true
			}
		}

		if !found {
			fmt.Println("Tidak ada pelanggan yang melakukan service pada tanggal tersebut")
		}
	}
}

func cariPelangganBySparePart(sparePart string) {
	fmt.Printf("Daftar pelanggan yang membeli spare-part %s:\n", sparePart)

	for _, trans := range transactions {
		if strings.EqualFold(trans.SparePart, sparePart) {
			customer := cariPelangganByID(trans.CustomerID)
			if customer != nil {
				fmt.Println(customer.Name)
			}
		}
	}
}

func tunjukkanPartsSeringGanti() {
	fmt.Println("Daftar spare-part secara terurut berdasarkan jumlah yang paling sering diganti:")

	sparePartCount := make(map[string]int)

	for _, trans := range transactions {
		sparePartCount[trans.SparePart]++
	}

	const maxSize = 100
	var sortedParts [maxSize]SparePart
	numSortedParts := 0

	for name, count := range sparePartCount {
		part := SparePart{
			Name:     name,
			Quantity: count,
		}
		if numSortedParts < maxSize {
			sortedParts[numSortedParts] = part
			numSortedParts++
		}
	}

	for i := 0; i < numSortedParts-1; i++ {
		for j := i + 1; j < numSortedParts; j++ {
			if sortedParts[i].Quantity < sortedParts[j].Quantity {
				sortedParts[i], sortedParts[j] = sortedParts[j], sortedParts[i]
			}
		}
	}

	for i := 0; i < numSortedParts; i++ {
		part := sortedParts[i]
		fmt.Printf("Spare-Part: %s, Jumlah Penggantian: %d\n", part.Name, part.Quantity)
	}
}

func parseTanggal(dateStr string) int {
	dateParts := strings.Split(dateStr, "/")
	day := dateParts[0]
	dayNumber := 0
	_, err := fmt.Sscanf(day, "%d", &dayNumber)
	if err != nil {
		fmt.Println("Format tanggal tidak valid")
		return 0
	}
	return dayNumber
}

func cariPelangganByID(id int) *Customer {
	for i := 0; i < nextCustomerID; i++ {
		if customers[i].ID == id {
			return &customers[i]
		}
	}
	return nil
}

func cariSparePartByNama(name string) *SparePart {
	for i := 0; i < nextPartID; i++ {
		if strings.EqualFold(spareParts[i].Name, name) {
			return &spareParts[i]
		}
	}
	return nil
}

func selectionSortSparePartsByNamaAsc() {
	for i := 0; i < nextPartID-1; i++ {
		minIndex := i
		for j := i + 1; j < nextPartID; j++ {
			if strings.Compare(spareParts[j].Name, spareParts[minIndex].Name) < 0 {
				minIndex = j
			}
		}
		if minIndex != i {
			spareParts[i], spareParts[minIndex] = spareParts[minIndex], spareParts[i]
		}
	}

	fmt.Println("Spare-Parts terurut berdasarkan nama (ascending):")
	for i := 0; i < nextPartID; i++ {
		fmt.Printf("ID: %d, Nama: %s, Jumlah: %d\n", spareParts[i].ID, spareParts[i].Name, spareParts[i].Quantity)
	}
}

func insertionSortCustomersByIDDesc() {
	for i := 1; i < nextCustomerID; i++ {
		key := customers[i]
		j := i - 1
		for j >= 0 && customers[j].ID < key.ID {
			customers[j+1] = customers[j]
			j--
		}
		customers[j+1] = key
	}

	fmt.Println("Pelanggan terurut berdasarkan ID (descending):")
	for i := 0; i < nextCustomerID; i++ {
		fmt.Printf("ID: %d, Nama: %s\n", customers[i].ID, customers[i].Name)
	}
}

func main() {
	var choice int
	for {
		fmt.Println("==================================== Dhafin Rakhaputra Bernadian/1301220476 ======================================")
		fmt.Println("==================================== Muhammad Zidan Ramadhan/1301223314 ======================================")
		fmt.Println()
		fmt.Println("========================================= Layanan Service Motor ==================================================")
		fmt.Println("Selamat datang! Ada yang bisa dibantu? Ketik nomor yang ada di bawah ini untuk berinteraksi ya~")
		fmt.Println("1. Tambah Spare-Part")
		fmt.Println("2. Edit Spare-Part")
		fmt.Println("3. Hapus Spare-Part")
		fmt.Println("4. Tambah Pelanggan")
		fmt.Println("5. Edit Pelanggan")
		fmt.Println("6. Hapus Pelanggan")
		fmt.Println("7. Tambah Transaksi")
		fmt.Println("8. Hitung Tarif Service")
		fmt.Println("9. Cari Pelanggan berdasarkan Tanggal")
		fmt.Println("10. Cari Pelanggan berdasarkan Spare-Part")
		fmt.Println("11. Tampilkan Spare-Part yang Sering Diganti")
		fmt.Println("12. Urutkan Spare Part Berdasarkan Nama (A-Z)")
		fmt.Println("13. Urutkan Pelanggan Berdasarkan ID (Descending)")
		fmt.Println("14. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var id int
			var name string
			var quantity int
			fmt.Print("ID Spare-Part: ")
			fmt.Scanln(&id)
			fmt.Print("Nama Spare-Part: ")
			fmt.Scanln(&name)
			fmt.Print("Jumlah: ")
			fmt.Scanln(&quantity)

			tambahSparePart(id, name, quantity)

		case 2:
			var id int
			var name string
			var quantity int

			fmt.Print("ID Spare-Part: ")
			fmt.Scanln(&id)
			fmt.Print("Nama Spare-Part: ")
			fmt.Scanln(&name)
			fmt.Print("Jumlah: ")
			fmt.Scanln(&quantity)

			if editSparePart(id, name, quantity) {
				fmt.Println("Spare-Part berhasil diubah")
			} else {
				fmt.Println("Spare-Part tidak ditemukan")
			}

		case 3:
			var id int

			fmt.Print("ID Spare-Part: ")
			fmt.Scanln(&id)

			if hapusSparePart(id) {
				fmt.Println("Spare-Part berhasil dihapus")
			} else {
				fmt.Println("Spare-Part tidak ditemukan")
			}

		case 4:
			var id int
			var name string

			fmt.Print("ID Pelanggan: ")
			fmt.Scanln(&id)
			fmt.Print("Nama Pelanggan: ")
			fmt.Scanln(&name)

			tambahPelanggan(id, name)

		case 5:
			var id int
			var name string

			fmt.Print("ID Pelanggan: ")
			fmt.Scanln(&id)
			fmt.Print("Nama Pelanggan: ")
			fmt.Scanln(&name)

			if editPelanggan(id, name) {
				fmt.Println("Pelanggan berhasil diubah")
			} else {
				fmt.Println("Pelanggan tidak ditemukan")
			}

		case 6:
			var id int

			fmt.Print("ID Pelanggan: ")
			fmt.Scanln(&id)

			if hapusPelanggan(id) {
				fmt.Println("Pelanggan berhasil dihapus")
			} else {
				fmt.Println("Pelanggan tidak ditemukan")
			}

		case 7:
			var transactionID int
			var customerID int
			var sparePart string

			fmt.Print("ID Transaksi: ")
			fmt.Scanln(&transactionID)
			fmt.Print("ID Pelanggan: ")
			fmt.Scanln(&customerID)
			fmt.Print("Spare-Part: ")
			fmt.Scanln(&sparePart)

			tambahTransaksi(transactionID, customerID, sparePart)

		case 8:
			var transactionID int

			fmt.Print("ID Transaksi: ")
			fmt.Scanln(&transactionID)
			hitungServiceRate(transactionID)

		case 9:
			var startDate string

			fmt.Print("Tanggal (dd/mm): ")
			fmt.Scanln(&startDate)

			cariPelangganByDate(startDate)

		case 10:
			var sparePart string

			fmt.Print("Spare-Part: ")
			fmt.Scanln(&sparePart)

			cariPelangganBySparePart(sparePart)

		case 11:
			tunjukkanPartsSeringGanti()

		case 12:
			selectionSortSparePartsByNamaAsc()

		case 13:
			insertionSortCustomersByIDDesc()

		case 14:
			tutupApp()

		default:
			fmt.Println("Menu tidak valid")
		}

		fmt.Println()
	}
}
