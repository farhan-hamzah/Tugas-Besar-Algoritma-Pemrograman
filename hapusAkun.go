package main
import "fmt"
const NMAX int = 100
type transaksi struct{
	userId int
	tipeTransaksi string
	quantityFiat int
	quantityKripto int
	tanggal int
	bulan int
	tahun int
}
type crypto struct{
	kriptoId int
	id int
	username string
	tanggalTranksasi transaksi
}
type wallet struct{
	id int
	name string
	pasokan int
	kripto crypto 
}
type account struct{
	id int
	username string
	password int
	moneyFiat float64
	saldoVirtual wallet
}
func main() {
	var data [NMAX]account
	var id, pass int

	// Simulasi data
	data[0] = account{id: 1, username: "andi", password: 123, moneyFiat: 5000}
	data[1] = account{id: 2, username: "budi", password: 456, moneyFiat: 3000, saldoVirtual: wallet{id: 1, name: "budi", pasokan: 20}}
	fmt.Print("Masukan Id: ")
	fmt.Scan(&id)
	fmt.Print("Masukan Password: ")
	fmt.Scan(&pass)
	// Hapus akun dengan id=1 dan password=123
	hapusAccount(&data, id, pass)
}

func hapusAccount(akun *[NMAX]account, id, pass int){
	var i, idx, temp int
	var lakukan bool
	lakukan = false
	temp = NMAX
	i = 0
	for i < NMAX{
		if akun[i].id == id && akun[i].password == pass{
			idx = i
			lakukan = true
			i = NMAX
		}else{
			i++
		}
	}
	if idx >= 0 && idx < temp && lakukan == true{
		for i = idx; i < temp-1; i++{
			akun[i].id = akun[i+1].id
			akun[i].username = akun[i+1].username
			akun[i].password = akun[i+1].password
			akun[i].moneyFiat = akun[i+1].moneyFiat
			akun[i].saldoVirtual.id = akun[i+1].saldoVirtual.id
			akun[i].saldoVirtual.name = akun[i+1].saldoVirtual.name
			akun[i].saldoVirtual.pasokan = akun[i+1].saldoVirtual.pasokan
		}
		temp--	
	}else{
		fmt.Println("Indeks Tidak Valid")
	}

	fmt.Println("\nAkun Yang Tersedia:")
	fmt.Printf("%-5s %-10s %-10s %-10s\n", "ID", "Username", "Saldo", "Pasokan")
	for i = 0; i < NMAX; i++ {
		if akun[i].id != 0{
			fmt.Printf("%-5d %-10s Rp%-10.2f %-10d\n", akun[i].id, akun[i].username, akun[i].moneyFiat, akun[i].saldoVirtual.pasokan)
		}
	}
	fmt.Println()
}