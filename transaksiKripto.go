package main
import (
	"fmt"
)
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
func jualBeli(tipeJualBeli string, dataJualBeli[NMAX]*transaksi, wallet[NMAX] account, id, pw int, n *int, uangFiat *float64){
	var pasokanKripto, jumlahKriptoYangDijual, jumlahKriptoYangDibeli,  nilaiBeliKripto int
	var i, tanggal, bulan, tahun, idxWallet int
	var nilaiJualKripto float64
	var valid bool
	valid = false
	i = 0
	for i < *n{
		if wallet[i].id == id && wallet[i].password == pw{
			pasokanKripto += wallet[i].saldoVirtual.pasokan
			idxWallet = i
		}
		i++
	}
	for valid == false{
		i = 0
		if tipeJualBeli == "jual"{	
			fmt.Print("Masukan Jumlah Kripto Yang Ingin Dijual: ")
			fmt.Scan(&jumlahKriptoYangDijual)
			if jumlahKriptoYangDijual <= pasokanKripto{
				nilaiJualKripto = float64(jumlahKriptoYangDijual)*21000
				*uangFiat = nilaiJualKripto
				fmt.Print("Tulis Tanggal, Bulan, dan Tahun: ")
				fmt.Scan(&tanggal, &bulan, &tahun)
				for i < NMAX{
					if dataJualBeli[i].userId == 0{
						dataJualBeli[i].userId = id
						dataJualBeli[i].tipeTransaksi = tipeJualBeli
						dataJualBeli[i].quantityFiat = int(nilaiJualKripto) 
						dataJualBeli[i].quantityKripto = -jumlahKriptoYangDijual
						dataJualBeli[i].tanggal = tanggal
						dataJualBeli[i].bulan = bulan
						dataJualBeli[i].tahun = tahun
						*n = *n+1
						i = NMAX
					}else{
						i++
					}
				}
				wallet[idxWallet].saldoVirtual.pasokan = pasokanKripto-jumlahKriptoYangDijual
				valid = true
			}else{
				fmt.Print("Jumlah Kripto Yang Dimiliki Tidak Tercukupi")
			}
		}else if tipeJualBeli == "beli"{
			fmt.Print("Masukan Jumlah Kripto Yang Ingin DIbeli: ")
			fmt.Scan(&jumlahKriptoYangDibeli)
			nilaiBeliKripto = jumlahKriptoYangDibeli*21000
			if *uangFiat >= float64(nilaiBeliKripto){
				*uangFiat -=float64(nilaiBeliKripto)
				fmt.Print("Tulis Tanggal, Bulan, dan Tahun: ")
				fmt.Scan(&tanggal, &bulan, &tahun)
				for i < NMAX{
					if dataJualBeli[i].userId == 0{
						dataJualBeli[i].userId = id
						dataJualBeli[i].tipeTransaksi = tipeJualBeli
						dataJualBeli[i].quantityFiat = -nilaiBeliKripto 
						dataJualBeli[i].quantityKripto = jumlahKriptoYangDibeli
						dataJualBeli[i].tanggal = tanggal
						dataJualBeli[i].bulan = bulan
						dataJualBeli[i].tahun = tahun
						*n =*n+1
						i = NMAX
					}else{
						i++
					}
				}
				wallet[idxWallet].saldoVirtual.pasokan = pasokanKripto+jumlahKriptoYangDibeli
				valid = true
			}else{
				fmt.Println("Uang Tidak Cukup")
			}		
		}else{
			fmt.Println("Tipe transaksi tidak dikenali (harus 'jual' atau 'beli').")
			valid = false
		}
	}
	
}