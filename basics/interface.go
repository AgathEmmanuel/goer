package main
  
import "fmt"
  
type shape interface {
  
    Area() float64
    Volume() float64
}
  
type dimension struct {
    length float64
    height float64
    breadth float64
}
  

func (d dimension) Area() float64 {
  
    return 2*(d.height*d.breadth+d.breadth*d.length+d.length*d.height)
}
  
func (d dimension) Volume() float64 {
  
    return d.height * d.breadth * d.length
}
  

func main() {
  

    var s shape
    s = dimension{10, 14, 20}

    fmt.Println("Area of tank :", s.Area())
    fmt.Println("Volume of tank:", s.Volume())
}
