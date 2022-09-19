package main
import "fmt"

type Cart struct {
    prices []float32
}

func (x Cart) show() {
    var sum float32 = 0.0
    //calculate the sum of all prices in the Cart
    for _, price := range x.prices {
      sum += price
    }
    
    fmt.Println(sum)
}

func main() {
  c := Cart{}
  var n int
  fmt.Scanln(&n)
  
  c.prices = []float32{}
  
  // take n inputs and add them to the Cart
  for i := 0; i < n; i++ {
    var pricedata float32
		fmt.Scanf("%f", &pricedata)
    c.prices = append(c.prices, pricedata)
  }
  
  //call the show() method of the Cart
  c.show()
  
}