package main 
import "fmt"

const nmax int = 1000

type akun  struct {
    username string
    komen[nmax]string
    value int 
    konotasi int 
    sentinel int 
    sentimen string 
    id int 
    
}

type komentars [nmax] akun 
func findmax(A komentars, n int) int {
    var idx_max, k int
    idx_max = 0
    k = 1
    for k < n { 
        if A[idx_max].sentinel  < A[k].sentinel  {
            idx_max = k
        }
        k++
    }
    return idx_max
}

func findmin(A komentars, n int) int {
    var idx_min, k int
    idx_min = 0
    k = 1
    for k < n {
        if A[idx_min].sentinel > A[k].sentinel {
            idx_min = k
        }
        k++
    }
    return idx_min
}

func selectionsortsentimen(A *komentars, n int) {
    var pass, idx, i int
    var temp akun

    for pass = 1; pass < n; pass++ {
        idx = pass - 1
        for i = pass; i < n; i++ {
            if A[idx].konotasi < A[i].konotasi {
                idx = i
            }
        }
        temp = A[pass-1]
        A[pass-1] = A[idx]
        A[idx] = temp
    }

    cetakdata(*A, n)
}


func statistik (A komentars , n int){
    var positif , negatif , netral   , i int 
    for i = 0 ; i < n ; i++ {
        if A[i].sentimen == "Positif"{
            positif++
        }else if A[i].sentimen == "Negatif"{
            negatif++ 
        } else {
            netral++
        }
    }
   
    fmt.Println("Banyak komentar : " , n )
    fmt.Println("Komentar terpanjang oleh username  : " , A[findmax(A  , n )].username )
    fmt.Println("Banyak komentar negatif : " ,  negatif)
    fmt.Println("Banyak komentar positif : " , positif  )
    fmt.Println("Banyak komentar netral : " , netral )
    
    
}

func insertionsortkomen (A *komentars, n int)  {
	var i, j int
	var temp akun
	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1
		for j >= 0 && A[j].sentinel > temp.sentinel {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
    cetakdata(*A , n )

}
func edit (A *komentars, n  int, Cari *int){
    var i , z  int
     binarysearch(A, n, Cari)
    i  =  *Cari 
    
    fmt.Println("Masukan komentar baru dan akhiri dengan '.' : ")
    for {
    fmt.Scan(&A[i].komen[z])
    if A[i].komen[z] == "." || z >= nmax {
        break
    }
    z++
}
    A[i].sentinel = z 

}
func hapuskomentar(A *komentars, n *int, Cari *int) {
    var i, z int
    binarysearch(A, *n, Cari)
    i = *Cari

    if i >= 0 && i < *n {
        fmt.Println("Menghapus komentar dari", A[i].username)
        for z = i; z < *n-1; z++ {
            A[z] = A[z+1]
        }
        *n--
        fmt.Println("Komentar berhasil dihapus.")
    } else {
        fmt.Println("Komentar tidak ditemukan atau tidak valid.")
    }
}
func binarysearch(A *komentars, n int , Cari *int) {
    var left, right, mid, i, z int
    var cari, pilihan string
    var targetValue int
    var found bool = false

    IsinilaiHurufPertama(A, n)
    insertionsort(A, n)

    fmt.Print("Masukkan username yang dicari: ")
    fmt.Scan(&cari)
    targetValue = hurufpertama(cari)

    left = 0
    right = n - 1

    // Binary search untuk cari salah satu index yang cocok
    for left <= right && !found{
        mid = (left + right) / 2
        if A[mid].value == targetValue {
            found = true
        } else if targetValue < A[mid].value {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    if found {
        // Cari index awal sebenarnya (mundur ke kiri)
        i = mid
        for i > 0 && A[i-1].value == targetValue && pilihan != "1" {
            i--
        }

        for i < n && A[i].value == targetValue {
            fmt.Println("Username:", A[i].username)
            fmt.Print("Komentar: ")
            for z = 0; z < A[i].sentinel; z++ {
                fmt.Print(A[i].komen[z], " ")
            }
            fmt.Println(" ")
            if A[i].konotasi < 0 {
             A[i].sentimen = "Negatif"
         } else if A[i].konotasi > 0 {
                A[i].sentimen = "Positif"
            } else {
                A[i].sentimen = "Netral"
            }
            fmt.Println("Sentimen:", A[i].sentimen)
            fmt.Println("================================")

            fmt.Println("Bukan komentar yang dicari?")
            fmt.Println("Ketik 1 untuk ke komentar selanjutnya, atau 2 jika sudah benar :")
            fmt.Print("Pilihan: ")
            fmt.Scan(&pilihan)

            if pilihan != "1" {
                *Cari = i 
                
            }

            i++
        }

    } else {
        fmt.Println("Tidak ada username yang cocok ")
    }
}
func insertionsort (A *komentars, n int)  {
	var i, j int
	var temp akun
    IsinilaiHurufPertama(A , n )
	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1
		for j >= 0 && A[j].username > temp.username {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}

}
func IsinilaiHurufPertama(A *komentars , n int ) {
    var i int 
    var huruf string
    for i = 0 ; i < n ; i++{
        huruf = A[i].username
            A[i].value = hurufpertama(huruf)
            
    }
}
func hurufpertama(s string) int {
    var firstChar byte
    firstChar = s[0] 
    switch {

    case firstChar == 'A' || firstChar == 'a':
        return 1
    case firstChar == 'B' || firstChar == 'b':
        return 2
    case firstChar == 'C' || firstChar == 'c':
        return 3
    case firstChar == 'D' || firstChar == 'd':
        return 4
    case firstChar == 'E' || firstChar == 'e':
        return 5
    case firstChar == 'F' || firstChar == 'f':
        return 6
    case firstChar == 'G' || firstChar == 'g':
        return 7
    case firstChar == 'H' || firstChar == 'h':
        return 8
    case firstChar == 'I' || firstChar == 'i':
        return 9
    case firstChar == 'J' || firstChar == 'j':
        return 10
    case firstChar == 'K' || firstChar == 'k':
        return 11
    case firstChar == 'L' || firstChar == 'l':
        return 12
    case firstChar == 'M' || firstChar == 'm':
        return 13
    case firstChar == 'N' || firstChar == 'n':
        return 14
    case firstChar == 'O' || firstChar == 'o':
        return 15
    case firstChar == 'P' || firstChar == 'p':
        return 16
    case firstChar == 'Q' || firstChar == 'q':
        return 17
    case firstChar == 'R' || firstChar == 'r':
        return 18
    case firstChar == 'S' || firstChar == 's':
        return 19
    case firstChar == 'T' || firstChar == 't':
        return 20
    case firstChar == 'U' || firstChar == 'u':
        return 21
    case firstChar == 'V' || firstChar == 'v':
        return 22
    case firstChar == 'W' || firstChar == 'w':
        return 23
    case firstChar == 'X' || firstChar == 'x':
        return 24
    case firstChar == 'Y' || firstChar == 'y':
        return 25
    case firstChar == 'Z' || firstChar == 'z':
        return 26
    default:
        return 0
    }
}
func bacadata(A *komentars, n *int) {
    var i, lagi, z, c int // kecuali lagi merupakan varibel iterasi ,lagi untuk milih opsi 
    var konotasiB string // tampung kata buat di cek positif negatif 
    var sama bool // buat berhenti loop 
    i = *n
    for {
        fmt.Println("Harap masukkan username (gunakan '_' jika menggunakan spasi ):")
        fmt.Scan(&A[i].username)
        sama = false
        c = 0 
        for  c < i && !sama   { // seq search untuk cari username yang sama 
            if A[i].username == A[c].username {
                sama = true
            }
            c++ 
        A[i].id = i     
        }
        for sama { // masuk ke loop kalau ketemu sama  
            fmt.Println("Username sudah digunakan, silakan gunakan username lain:")
            fmt.Scan(&A[i].username)
            
            c = 0 
            sama = false
            for  c < i && !sama   {
            if A[i].username == A[c].username {
                sama = true
            }
            c++ 
            }
        }

        fmt.Println("Masukan komentar dan akhiri kata terakhir dengan '.'")
        z = 0
        A[i].konotasi = 0 // reset konotasi sebelum diisi
        for  {
            if  z == nmax{
                A[i].komen[z] = "."
            }
            fmt.Scan(&A[i].komen[z])
            
            if A[i].komen[z] == "."  {
                break // berhenti input 
            }
            
            konotasiB = A[i].komen[z] // sdmpen perkata 
            if negatif(konotasiB) {
                A[i].konotasi = A[i].konotasi - 1  // cek negatif 
            } else if positif(konotasiB) {
                A[i].konotasi = A[i].konotasi + 1  // cek positif 
            } 
            
            z++  // lanjut kata selanjutnya 
            
        }
        

        A[i].sentinel = z // simpen panjang komentar / berapa banyak kata di username nya 
        A[i].id = i // simpen indeks username,komentar , panjang komentar 
        for z = 0 ; z <= A[i].sentinel ; z++ {
            konotasiB = A[i].komen[z] 
            if specialkey(konotasiB) {
                A[i].konotasi = A[i].konotasi * -1  // cek ada kata negasi 
            }
        }
        fmt.Println("Masih mau menambahkan komentar?")
        fmt.Println("Ya                          Tidak")
        fmt.Println("1                              2")
        fmt.Scan(&lagi) // cek lanjut tambah komen atau ga 
        
        for lagi != 1 && lagi != 2 { // kondisi input selain 1 sama dua 
            fmt.Println("Tolong masukan angka sesuai opsi yang tertera")
            fmt.Println("Masih mau menambahkan komentar?")
            fmt.Println("Ya                          Tidak")
            fmt.Println("1                              2")
            fmt.Scan(&lagi)
        }
         i++ // lanjut kata 
        *n = i // ngitung berapa banyak akun 
        if lagi == 2 {
            break // berhentiin repeat until pertama 
        }
       
    }
    fmt.Println(*n    , "Komentar berhasil ditambahkan")
    fmt.Println(" ")
}
func cetakdata(A komentars , n int ){ // cetak array biasa 
    var i,z int 
    for i = 0 ; i < n ; i++ {
        fmt.Print("Username : " , A[i].username)
        fmt.Println(" ")
        fmt.Println("Komentar : ")
        
        for z = 0 ; z < A[i].sentinel ; z++ {
        fmt.Print(A[i].komen[z]," ")
        }
        if A[i].konotasi < 0 {
            A[i].sentimen = " Negatif "
        } else if A[i].konotasi > 0  {
            A[i].sentimen = " Positif "
            
        } else {
            A[i].sentimen = "Netral"
        }
        fmt.Println(" ")
        fmt.Println("Sentimen : " , A[i].sentimen )  
        fmt.Println(" ")
    }
}   
func main () {
    var A komentars
    var n , pilihan , Cari int 
    for {
        menu()
        fmt.Scan(&pilihan)
        if pilihan > 10 || pilihan <= 0 {
            for pilihan > 10  ||  pilihan <= 0 {
                fmt.Println("Tolong masukan angka sesuai dengan opsi")
                 fmt.Scan(&pilihan)
            }
        }    
         switch pilihan {
        case 1 :  bacadata(&A , &n)
         case 2: hapuskomentar(&A, &n, &Cari)
        case 3 :  edit(&A , n  , &Cari)
        case 4 : seqsearch(A , n )
        case 5 : fmt.Print("Tilakan masukan username yang akan dicari ")
                 binarysearch(&A,n , &Cari)
        case 6 : cetakdata(A , n )
        case 7 : insertionsortkomen(&A,n)
        case 8 : selectionsortsentimen(&A , n )
        case 9 : statistik(A,n)
             
       }
        if pilihan == 10  {
            break
        }
    }         
}
func menu(){
    fmt.Println(" APLIKLASI ANALISIS SENTIMEN      ")
    fmt.Println("==================================")
    fmt.Println("   Silahkan Pilih Aktivitas    ")
    fmt.Println("==================================")
    fmt.Println("1. Masukan Komentar")
    fmt.Println("2. Hapus Komentar ")
    fmt.Println("3. Mengedit Komentar")
    fmt.Println("4. Mencari Komentar ")
    fmt.Println("5. Mencari username  ")
    fmt.Println("6. Cetak berdasarkan urutan input   ")
    fmt.Println("7. Cetak berdasarkan panjang komentar (dari terpendek ke terbesar )")
    fmt.Println("8. Cetak berdasarkan sentimen (positif ke negatif) ")
    fmt.Println("9. Statistik ")
    fmt.Println("10. Exit ")
    fmt.Println("==================================")
    fmt.Print("pilih(1/2/3/4/5/6/7/8/9/10) ")
    
}
func negatif(word string) bool { // cek kata negatid
    return word == "bodoh" || word == "BODOH" || word == "Bodoh" ||
           word == "tolol" || word == "TOLOL" || word == "Tolol" ||
           word == "kecewa" || word == "KECEWA" || word == "Kecewa" ||
           word == "payah" || word == "PAYAH" || word == "Payah" ||
           word == "gagal" || word == "GAGAL" || word == "Gagal" ||
           word == "marah" || word == "MARAH" || word == "Marah" ||
           word == "benci" || word == "BENCI" || word == "Benci" ||
           word == "sampah" || word == "SAMPAH" || word == "Sampah" ||
           word == "rendahan" || word == "RENDAHAN" || word == "Rendahan" ||
           word == "cacat" || word == "CACAT" || word == "Cacat"
}
func positif(word string) bool { // cek kata positif 
    return word == "keren" || word == "KEREN" || word == "Keren" ||
           word == "bagus" || word == "BAGUS" || word == "Bagus" ||
           word == "suka" || word == "SUKA" || word == "Suka" ||
           word == "menarik" || word == "MENARIK" || word == "Menarik" ||
           word == "mantap" || word == "MANTAP" || word == "Mantap" ||
           word == "hebat" || word == "HEBAT" || word == "Hebat" ||
           word == "membantu" || word == "MEMBANTU" || word == "Membantu" ||
           word == "tepat" || word == "TEPAT" || word == "Tepat" ||
           word == "kreatif" || word == "KREATIF" || word == "Kreatif" ||
           word == "cerdas" || word == "CERDAS" || word == "Cerdas"
}
func specialkey(word string)bool { // cek kata negasi 
    return word == "tak" || word == "TAK" || word == "Tak" ||
           word == "tidak" || word == "TIDAK" || word == "Tidak" ||
           word == "gak" || word == "GAK" || word == "Gak" ||
           word == "bukan" || word == "BUKAN" || word == "Bukan"
}
func seqsearch(A komentars, n int) {
    var lagi int
    var cari string
    var found bool
    var i, z  , j int

    for {
        fmt.Print("Masukan Kata yang ingin dicari: ")
        fmt.Scan(&cari)
        found = false
        i = 0

        // cari semua komentar yang mengandung kata
        for i <= n {
            z = 0
            // Mencari kata dalam komentar
            for z < A[i].sentinel && A[i].komen[z] != cari {
                z++
            }

            // Jika ketemu
            if z < A[i].sentinel {
                if !found {
                    fmt.Println("Komentar yang mengandung kata '", cari, "':")
                    fmt.Println("==================================")
                    found = true
                }
                fmt.Println("Username:", A[i].username)
                fmt.Println("Komentar:")
                for j = 0; j < A[i].sentinel; j++ {
                    fmt.Print(A[i].komen[j], " ")
                }
                fmt.Println(" ")
                fmt.Println("==================================")
            }
            i++
        }

        if !found {
            fmt.Println("Komentar dengan kata '", cari, "' tidak ditemukan")
            fmt.Println("==================================")
        }

        // Konfirmasi pencarian lagi
        fmt.Println("Cari kata lain?")
        fmt.Println("1. Ya")
        fmt.Println("2. Tidak")
        fmt.Scan(&lagi)
        
     
        if lagi == 2 {
           break 
        }
    }
}
