package main 
import "fmt"

const nmax int = 1000

type akun  struct {
    username string
    komen[nmax]string
    id int 
    konotasi int 
    sentinel int 
    
    
}

type komentars [nmax] akun 
func seqsearch(A komentars, n int) {
    var i, z, x, lagi int  // i,z,x variabel iterasi lagi untuk option 
    var cari string // variabel yang dicari 
    var found, stop bool // varibel buat berhenti loop dan syarat if 
    
    fmt.Println("Tolong masukan keyword yang ingin dicari: ")
    fmt.Scan(&cari)
    
    i = 0  
    for !stop && i < n {  
        z = 0
        // Cari kata dalam komentar
        for z < A[i].sentinel && A[i].komen[z] != cari {
            z++
        }
        
        // Jika ditemukan
        if z < A[i].sentinel && A[i].komen[z] == cari {
            found = true
            fmt.Println("=== Komentar ditemukan ===")
            fmt.Println("Username: %s\n", A[i].username)
            fmt.Println("Komentar dengan kata " , cari)
            
            // print komentar 
            x = 0  r
            for x < A[i].sentinel {  
                fmt.Print(A[i].komen[x], " ")
                x++
            }
            fmt.Println(" ")
            fmt.Println("=========================")
            
            fmt.Println ("Cari komentar selanjutnya dengan kata yang sama? ")
            fmt.Println("Ya                          Tidak")
            fmt.Println("1                              2")
            fmt.Scan(&lagi)
            
            // cek input wajib 1 atau 2 
            for lagi != 1 && lagi != 2 {  
                fmt.Println("Input tidak valid. Silakan pilih 1 untuk Ya atau 2 untuk Tidak")
                fmt.Scan(&lagi)
            }
            
            if lagi == 2 { // berhenti karna milih 2 (tidak)
                stop = true
            }
        }
        i++
    }
    
    if !found { // kondisi ga ketemu kata
        fmt.Println("Tidak ditemukan komentar yang mengandung kata ", cari)
    }
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
func bacadata(A *komentars, n *int) {
    var i, lagi, z, c int // kecuali lagi merupakan varibel iterasi ,lagi untuk milih opsi 
    var konotasiB string // tampung kata buat di cek positif negatif 
    var sama bool // buat berhenti loop 
    i = 0
    
    for {
        fmt.Println("Harap masukkan username (hanya satu kata):")
        fmt.Scan(&A[i].username)
        sama = false
        c = 0 
        for  c < i && !sama   { // seq search untuk cari username yang sama 
            if A[i].username == A[c].username {
                sama = true
            }
            c++ 
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

        fmt.Println("masukan komentar dan akhiri kata terakhir dengan '.'")
        z = 0
        for {
            fmt.Scan(&A[i].komen[z])
           
            if A[i].komen[z] == "." {
                break // berhenti input 
            }
            
            konotasiB = A[i].komen[z] // sdmpen perkata 
            if negatif(konotasiB) {
                A[i].konotasi = A[i].konotasi - 1  // cek negatif 
            } else if positif(konotasiB) {
                A[i].konotasi = A[i].konotasi + 1  // cek positif 
            } 
            if specialkey(konotasiB) {
                A[i].konotasi = A[i].konotasi * -1  // cek ada kata negasi 
            }
            z++  // lanjut kata selanjutnya 
        }

        A[i].sentinel = z // simpen panjang komentar / berapa banyak kata di username nya 
        A[i].id = i // simpen indeks username,komentar , panjang komentar 
        fmt.Println("masih mau menambahkan komentar?")
        fmt.Println("Ya                          Tidak")
        fmt.Println("1                              2")
        fmt.Scan(&lagi) // cek lanjut tambah komen atau ga 
        
        for lagi != 1 && lagi != 2 { // kondisi input selain 1 sama dua 
            fmt.Println("tolong masukan angka sesuai opsi yang tertera")
            fmt.Println("masih mau menambahkan komentar?")
            fmt.Println("Ya                          Tidak")
            fmt.Println("1                              2")
            fmt.Scan(&lagi)
        }
        
        if lagi == 2 {
            break // berhentiin repeat until pertama 
        }
        i++ // lanjut kata 
        *n = i // ngitung berapa banyak akun 
    }
}

func cetakdata(A komentars , n int ){ // cetak array biasa 
    var i,z int 
    for i = 0 ; i <= n ; i++ {
        fmt.Print("username : " , A[i].username)
        fmt.Println(" ")
        fmt.Println("komentar : ")
        fmt.Println(" ")
        for z = 0 ; z < A[i].sentinel ; z++ {
        fmt.Print(A[i].komen[z]," ")
        }
        fmt.Println(" ")
    }
}

    
func main () {
    var A komentars
    var n,pilihan int 
    
    for {
        menu()
        fmt.Scan(&pilihan)
        if pilihan > 5 && pilihan <= 0 {
            for pilihan > 5 && pilihan <= 0 {
                fmt.Println("tolong masukan angka sesuai dengan opsi")
                 fmt.Scan(&pilihan)
            }
        }    
         switch pilihan {
        case 1 :  bacadata(&A , &n)
       // case 2 :  edit(&A , n )
        //case 3 :  hapus(&A , n )
        case 4 : seqsearch(A , n )
        case 5 : cetakdata(A , n )
       }
        if pilihan == 6  {
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
    fmt.Println("2. Mengubah Komentar ")
    fmt.Println("3. Menghapus Komentar")
    fmt.Println("4. Mencari Komentar ")
    fmt.Println("5. Statistik ")
     fmt.Println("6. Exit ")
    fmt.Println("==================================")
    fmt.Println("pilih(1/2/3/4/5) ")
    
}
