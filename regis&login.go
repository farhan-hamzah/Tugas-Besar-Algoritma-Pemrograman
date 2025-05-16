package main
import (
	"fmt"
)
const NMAX int = 100
type transaksi struct{
	userId int
	tipeTransaksi string
	quantity int
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
func main(){
	var akun[NMAX]account
	var pilihAkun, n int
	fmt.Println("Selamat datang di apk fokus kripto pilih menu sesuai yang anda inginkan")
	fmt.Println("1. Login")
	fmt.Println("2. Buat akun")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihAkun)
	login(akun, &pilihAkun, &n)
}
func login(akun[NMAX] account, pilihAkun *int, n *int){
	var i, id int
	var usrn string
	var uangFiat float64
	var pass int
	var benarLogin, buatAkun, idUnik bool
	for *pilihAkun != 1 && *pilihAkun != 2{
		fmt.Print("Masukkan pilihan (1/2): ")
		fmt.Scan(pilihAkun)
	}
	if *pilihAkun == 1{
		fmt.Print("Masukan id: ")
		fmt.Scan(&id)
		fmt.Print("Masukan Password: ")
		fmt.Scan(&pass)
		i = 0
		benarLogin = false
		for i < *n{
			if id == akun[i].id && pass == akun[i].password && benarLogin == false{
				benarLogin = true
			}
			i++
		}
		if benarLogin{
			fmt.Println("Login berhasil. ")
		}else{
			fmt.Println("ID atau password salah. ")
		}
	}else if *pilihAkun == 2{
		buatAkun = false
		for buatAkun == false {
			fmt.Print("Buat Id (dalam integer kombinasi 0 - 9): ")
			fmt.Scan(&id)
			fmt.Print("Masukan Username: ")
			fmt.Scan(&usrn)
			fmt.Print("Buat Password (dalam integer kombinasi 0 - 9): ")
			fmt.Scan(&pass)
			i = 0
			idUnik = true
			for i < *n{
				if id == akun[i].id{
					idUnik = false
				}
			}
			if idUnik{
				fmt.Print("Masukan nilai uang: ")
				fmt.Scan(&uangFiat)
				akun[*n].id = id
				akun[*n].username = usrn
				akun[*n].password = pass
				*n = *n+1
				buatAkun  = true
				fmt.Println("Akun berhasil dibuat")
			}else{
				fmt.Println("ID sudah digunakan. Silahkan cobalagi")
			}
		}
			
	}
}