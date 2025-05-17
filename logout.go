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
func logOut(A [NMAX]account, id, pass int){
	var i int
	var keluar bool
	keluar = false 
	for i = 0; i < NMAX; i++{
		if A[i].id == id && A[i].password == pass{
			keluar = true
			i = NMAX
		}
	}
	if keluar == true{
		//Keluar dari akun
	}else{
		fmt.Print("Password atau Id Salah")
	}
}