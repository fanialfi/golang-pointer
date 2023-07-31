# Pointer

pointer adalah _reference_ atau alamat memori, variabel pointer berarti variabel yang berisi alamat memori suatu nilai.
Sebagai contoh sebuah variabel bertipe `int` berisi 4, maka yang dimaksud dengan pointer adalah **alamat memori dimana nilai 4 disimpan**

variabel variabel yang memiliki reference atau alamat pointer yang sama, saling terhubung satu sama lain dan nilainya pasti sama, ketika ada perubahan nilai, maka akan memberikan dampak kepada variabel yang memiliki alamat memori yang sama.

variabel bertipe pointer ditandai dengan adanya tanda **asterisk** (`*`) tepat sebelum penulisan tipe data suatu variabel.
Nilai default variabel pointer adalah `nil` (kosong), variabel pointer tidak bisa menerima value yang bukan pointer dan juga sebaliknya, variabel biasa tidak bisa menerima nilai pointer.

<details>
  <summary>contoh pembuatan variabel pointer</summary>

```go
var number *int
var name *string
```
</details><br>

walaupun variabel biasa tidak bisa menerima nilai pointer ataupun juga sebaliknya, namun ada 2 hal penting mengenai pointer :
1. variabel biasa bisa diambil nilai pointernya dengan cara menambahkan tanda **ampersand** (`&`) tepat sebelum nama variabel-nya, metode ini disebut dengan **referencing**.

2. dan sebaliknya nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda **asterisk** (`*`) tepat sebelum nama variabelnya, metode ini disebut dengan **dereferencing**.

<details>
  <summary>example</summary>

```go
var numberA int = 4
var numberB *int = &numberA

func main(){
  fmt.Println("numberA (value)", numberA)
  fmt.Println("numberA (address)", &numberA)

  fmt.Println("numberB (value)", *numberB)
  fmt.Println("numberB (address)", numberB)
}
```
</details><br>

dari contoh di atas variabel `numberB` yang merupakan varibel pointer int dimana value-nya mengambil nilai pointer dari variabel `numberA` menggunakan `&numberA`.
Variabel pointer jika di print ke layar akan menghasilkan string alamat memori dengan notasi hexadecimal.

Lalu untuk mengambil nilai asli variabel pointer dengan caara melakukan **dereferencing** seperti pada contoh di atas (`*numberB`).

Namun ketika salah satu variabel diubah nilainya dan ada variabel lain yang memiliki referensi memori yang sama, maka variabel lain nilainya akan ikut berubah.
<details>
  <summary>example</summary>

```go
var numberA int = 4
var numberB *int = &numberA

func main(){
  fmt.Println("number A (value)",numberA)
  fmt.Println("number A (address)", &numberA)

  fmt.Println("number B (value)", *numberB)
  fmt.Println("number B (address)", numberB)

  fmt.Println()
  // ubah nilai variabel numberB
  *numberB = 30
  fmt.Println("number A (value)",numberA)
  fmt.Println("number A (address)", &numberA)

  fmt.Println("number B (value)", *numberB)
  fmt.Println("number B (address)", numberB)
}
```
</details><br>

pada contoh di atas variabel `numberA` dan `numberB` memiliki referensi alamat memori yang sama, kemudian isi dari `numberB` diubah ke _30_, semua variabel yang memiliki referensi alamat yang sama dengan `numberB` akan ikut berubah (dalam hal ini `numberA`).

pointer juga bisa digunakan sebagai parameter di function, cara penerapannya juga sama dengan cara mendeklarasikan parameter sebagai pointer

<details>
  <summary>example</summary>

```go
package main

import "fmt"

func main(){
  var number int = 4
  fmt.Println("before :", number) // 4
  
  change(&number, 10)
  fmt.Println("after :", number) // 10
}

func change(original *int, number){
  *original = number
}
```
</details><br>

fungsi `change()` diatas memiliki 2 parameter, parameter pertama berupa pointer int, dan parameter ke-2 berupa int.
Didalam function tersebut nilai asli parameter `original` diubah oleh parmeter `number`.