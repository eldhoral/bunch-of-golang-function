package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	//variabel
	var namaDepan, namaBelakang, departemen string
	namaDepan, namaBelakang, departemen = "Eldho", "Rizcky", "IT"
	r := gin.Default()
	r.GET("/nama", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Klik Dokter": "Halo!!" + " " + namaDepan + " " + namaBelakang + " " + departemen,
		})
	})

	//operator
	volume := hitungVolume(10, 5, 5)
	luas := hitungLuas(10, 10, 5)
	r.GET("/sum", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Volume dari balok adalah": volume,
			"Luas dari balok adalah":   luas,
		})
	})

	//fungsi
	var hari uint32 = 1
	var menit, jam uint32
	menit, jam = konversiWaktu(1)
	r.GET("/fungsi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Dalam berapa hari":         hari,
			"Berapa menit dalam 1 hari": menit,
			"berapa jam dalam 1 hari":   jam,
		})
	})

	//nilai rapot
	var nilaiAndi, nilaiBudi, nilaiRudi, rataRata float64
	nilaiAndi = (math.Floor(nilaiAcak(2, 10)*100) / 100)
	nilaiBudi = (math.Floor(nilaiAcak(2, 10)*100) / 100)
	nilaiRudi = (math.Floor(nilaiAcak(2, 10)*100) / 100)
	rataRata = (math.Floor(rataRataNilai(nilaiAndi, nilaiBudi, nilaiRudi)*100) / 100)
	r.GET("/rapot", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Nilai Rapot Andi":      nilaiAndi,
			"Nilai Rapot Budi":      nilaiBudi,
			"Nilai Rapot Rudi":      nilaiRudi,
			"Rata Rata Nilai Kelas": rataRata,
		})
	})

	//pointer
	var nilaiA int = 5
	var nilaiB *int = &nilaiA
	r.GET("/pointer", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"nilai dari A":  nilaiA,
			"alamat dari A": &nilaiA,
			"nilai dari B setelah ambil nilai pointer dari A":   *nilaiB,
			"alamat dari B setelah ambil alamat dari pointer A": nilaiB,
		})
	})

	//struct
	var m = mahasiswa{}
	m.nama = "Budi Kece"
	m.angkatan = 2020
	m.jurusan = "IT"
	r.GET("/struct", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Nama":        m.nama,
			"Nama Kampus": m.namaKampus,
			"Jurusan":     m.jurusan,
		})
	})

	//routine di golang
	runtime.GOMAXPROCS(7)
	go routine(2, "Hai")
	routine(5, "Hai Juga")
	var input string
	fmt.Scanln(&input)
	r.GET("/routine", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Routine": "Sudah Jalan",
		})
	})

	//method + struct
	var k1 = kampus{}
	k1.namaKampus = "Universitas Budaya Bencana"
	k1.akreditasi = "B"
	//var cekAkreditasi = m.panggilAkreditasi()
	r.GET("/method", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Selamat Datang di ":               k1.namaKampus,
			"Akreditasi Kampus ini adalah  ":   k1.akreditasi,
			"Apakah Akreditasi kampus ini A? ": k1.bukanAkreditasiA("A"),
		})
	})

	//routine pakai channel
	runtime.GOMAXPROCS(7)
	var pesanChannel = make(chan string, 2)
	pesanChannel <- absen("Abdul")
	pesanChannel <- absen("Hani")
	close(pesanChannel)
	for iterasiChannel := range pesanChannel {
		fmt.Println(iterasiChannel)
	}
	r.GET("/channel", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Pesan": "Channel Sudah Jalan",
		})
	})

	//konversi 1 to satu
	//r.GET("/konversi1tosatu", func(c *gin.Context) {
	//c.JSON(200, gin.H{
	//"1 adalah ":  konversi1tosatu(1),
	//"16 adalah ": konversi1tosatu(16),
	//})
	//})

	//konversi satu to 1
	huruf := "satu, dua, tiga, empat, lima, enam, tujuh, delapan, sembilan"
	konversiHuruf := strings.Replace(huruf, huruf, "1, 2, 3, 4, 5, 6, 7, 8, 9", 1)
	r.GET("/konversisatuto1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Sudah di split": konversiHuruf,
		})
	})

	//jalan di localhost:8080
	r.Run() // listen and serve on localhost:8080

	//run dicmd
	//docker-compose -f docker-compose.dev.yml up
}

//fungsi konversi waktu
func konversiWaktu(hari uint32) (uint32, uint32) {
	// hitung menit
	var menit = hari * 1440
	// hitung jam
	var jam = hari * 24
	// kembalikan 2 nilai
	return menit, jam
}

//fungsi nilai acak
func nilaiAcak(min, max float64) float64 {
	var nilai = (rand.Float64() * (max - min)) + min
	return nilai
}

//fungsi rata rata nilai
func rataRataNilai(nilai1 float64, nilai2 float64, nilai3 float64) float64 {
	var mean = (nilai1 + nilai2 + nilai3) / 3
	return mean
}

//struct kampus
type kampus struct {
	namaKampus, akreditasi string
}
type mahasiswa struct {
	nama, jurusan string
	angkatan      int
	kampus
}

//method struct balok
//func (b balok) hitungVolume() float64 {
//return b.panjang*b.lebar*b.tinggi
//}

//func (b balok) hitungLuas() float64 {
//return 2*(b.panjang*b.tinggi)+2*(b.panjang*b.lebar)+2*(b.lebar*b.tinggi)
//}

func hitungVolume(pV, lV, tV float64) float64 {
	var volumebalok = pV * lV * tV
	return volumebalok
	//fmt.Println("Volume Balok = ", volumebalok)
}

func hitungLuas(pL, lL, tL float64) float64 {
	var luasbalok = 2*(pL*tL) + 2*(pL*lL) + 2*(lL*tL)
	return luasbalok
	//fmt.Println("Luas Balok = ", luasbalok)
}

//func hasilBalok(p, l, t int) *balok {
//return &balok{p, l, t}
//}

//fungsi routine
func routine(max int, text string) {
	for i := 0; i < max; i++ {
		fmt.Println((i + 1), text)
	}
}

//method struct kampus
func (k kampus) bukanAkreditasiA(nilaiAkreditasi string) bool {
	return k.akreditasi > nilaiAkreditasi
}

//routine channel
func absen(namaAbsen string) string {
	var data = "Hai, Selamat Datang " + namaAbsen
	return data
}

//konversi 1 to satu
//func konversi1tosatu(i int) string {
//return string('A' - 1 + i)
//}
