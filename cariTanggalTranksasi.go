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
func mencariDataTranksasi(A [NMAX]transaksi, akun [NMAX]account, tanggal, bulan, tahun, id int){
	var i, temp int
	var ada bool
	ada = false
	for i =0; i < NMAX; i++{
		if A[i].tanggal == tanggal && A[i].bulan == bulan && A[i].tahun == tahun && A[i].userId == id{
			temp = i
			ada = true
			i = NMAX
		}else if A[i].userId == 0{
			i = NMAX
		}
	}
	if ada == true{
		fmt.Println("\nAkun Yang Tersedia:")
		fmt.Printf("%-5s %-10s %-10s %-10s\n", "ID", "Username", "Saldo", "Pasokan")
		fmt.Printf("%-5d %-10s Rp%-10.2f %-10d\n", akun[temp].id, akun[temp].username, akun[temp].moneyFiat, akun[temp].saldoVirtual.pasokan)
	}else{
		fmt.Print("Tidak Ada Riwayat Transaksi")
	}
}