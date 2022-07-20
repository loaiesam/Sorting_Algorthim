package main
import(
  "fmt"
)

func main() {
  array1 := [...]int{5, 7, 9, 2, 4, 6, 3, 1, 8}
  array2 := [...]int{10, 60, 40, 30, 20, 50, 80, 90, 70}
  array3 := [...]int{400, 600, 800, 700, 200, 500, 100, 300, 900}
  arr(array1)
  arr(array2)
  arr(array3)
}

func arr(arr [9]int) {
  array := make([]int, 9, 64)
  array = arr[:]
  size := len(array)- 1
  for i:= 0; i < size;i++ {
    for j:= 0; j < size-i;j++ {
      if array[j] > array[j+1] {
        tmp := array[j]
        array[j] = array[j+1]
        array[j+1] = tmp
      }
    }
  }
  fmt.Printf("\n Capacity: %d\n Length: %d\n Type: %T\n Value: %v\n Binary: %b\n\n", cap(array),len(array), array, array, array)
  for k:= 0; k < 50; k++ {
    fmt.Print("=-")
  }
  fmt.Print("\n")
}
