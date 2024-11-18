package ewallet

import (
	"errors"
)

func init() {
	pengujian = false
}

var saldoAkun float64 = 0
var pengujian = true

type Ewallet interface {
	Withdraw(jumlahUang float64) (float64, error)
	Deposit(jumlahUang float64) (float64, error)
	CetakSaldo() float64
}

type EwalletImplementasi struct{}

func (e EwalletImplementasi) Withdraw(jumlahUang float64) (float64, error) {
	if saldoAkun < jumlahUang {
		return 0, errors.New("saldo anda tidak mencukupi")
	}
	saldoAkun = saldoAkun - jumlahUang
	return saldoAkun, nil
}

func (e EwalletImplementasi) Deposit(jumlahUang float64) (float64, error) {
	if jumlahUang <= 0 {
		return 0, errors.New("nilai setor anda tidak dikenali")
	}
	saldoAkun = saldoAkun + jumlahUang
	return saldoAkun, nil
}

func (e EwalletImplementasi) CetakSaldo() float64 {
	return saldoAkun
}

var saldoUji float64 = 0

type EwalletPengujian struct{}

func (e EwalletPengujian) Withdraw(jumlahUang float64) (float64, error) {
	if saldoUji < jumlahUang {
		return 0, errors.New("saldo anda tidak mencukupi")
	}
	saldoUji = saldoUji - jumlahUang
	return saldoUji, nil
}

func (e EwalletPengujian) Deposit(jumlahUang float64) (float64, error) {
	if jumlahUang <= 0 {
		return 0, errors.New("nilai setor anda tidak dikenali")
	}
	saldoUji = saldoUji + jumlahUang
	return saldoUji, nil
}

func (e EwalletPengujian) CetakSaldo() float64 {
	return saldoUji
}

func JalankanAplikasiEwallet() Ewallet {
	if pengujian {
		return EwalletPengujian{}
	}
	return EwalletImplementasi{}
}

func JalankanPerintah(daftarPerintah []string) (float64, error) {
	var err error
	ewalletInstance := JalankanAplikasiEwallet()

	for _, perintah := range daftarPerintah {
		if perintah == "deposit" {
			_, err = ewalletInstance.Deposit(50000)
		} else {
			_, err = ewalletInstance.Withdraw(25000)
		}
	}
	return ewalletInstance.CetakSaldo(), err
}
