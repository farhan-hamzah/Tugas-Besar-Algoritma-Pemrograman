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
func rangkingKripto(A *[NMAX]account ){
	var i, idx, j, temp int
	temp = NMAX+1
	for i = 0; i < NMAX; i++{
		if A[i].id != 0{
			idx = i
			for j = i+1; j < NMAX; j++{
				if A[j].id != 0 && A[i].saldoVirtual.pasokan < A[j].saldoVirtual.pasokan{
					idx = j
				}
			}
			A[temp] = A[i]
			A[i] = A[idx]
			A[idx] = A[temp]
		}else{
			i = NMAX
		}
	}
	fmt.Println("=== Top Tier Kripto ===")
	for i = 0; i < NMAX; i++ {
		if A[i].id != 0 {
			fmt.Printf("ID: %d | Username: %s | Pasokan: %d\n",
				A[i].id, A[i].username, A[i].saldoVirtual.pasokan)
		}
	}
}