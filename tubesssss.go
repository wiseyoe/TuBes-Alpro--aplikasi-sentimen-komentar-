
package main
import "fmt"
const nmax int = 10000
type akun struct {
    nama , komentar string 
    kata int 
}
type tabstring [nmax] akun 

func bacadata(A *tabstring , n,y *int ){
    var i int 
    var lagi int
    lagi = 1 
    i = 0 
    *y = 0 
    for lagi == 1 {
        fmt.Scan(&A[i].nama)
        fmt.Scan(&A[*y].komentar)
        A[i].kata = *y 
        fmt.Println("maskan komentar lagi ?")
        fmt.Println("Ya                tidak ")
        fmt.Println("1                   2 ")
        fmt.Scan(&lagi)
        i++
        *y++ 
        
    }
    *n = i 

}
func cetakdata(A tabstring , n , y int ) {
    var i int 
    for i = 0 ; i <= n ; i++{
        fmt.Println("username : " , A[i].nama)
        if A[i].kata - A[i+1].kata < 0 {
            for y = 0 ; y <= A[i].kata ; y ++ {
                fmt.Print(A[y].komentar , " ")
            } 
        }else if A[i].kata + 1 > n {
                for y = A[i-1].kata + 1 ; y <= n ; y++{
                     fmt.Print(A[y].komentar , " ")
                }
        }else {
                    for y = A[i-1].kata + 1 ;   y <= A[i].kata ; y++ {
                         fmt.Print(A[y].komentar , " ")
                    }
                }
        }
}

func main() {
    var n , y  int 
    var A tabstring
    bacadata(&A , &n , &y)
    cetakdata(A , n , y  )
    
}
/*
func tampilanmenu(){
  fmt.Println("  Judul tubes      ")
  fmt.Println("-------------------------------")
  fmt.Println("  MENU        ")
  fmt.Println("-------------------------------")
  fmt.Println("  1. masukan komentar     ")
  fmt.Println("  2. mengedit komentar    ")
  fmt.Println("  3. Hapus komentar ")
  fmt.Println("  4. Statistik   ")
  fmt.Println("-------------------------------")
  fmt.Println("Pilih (1/2/3/4)")
}
*/
