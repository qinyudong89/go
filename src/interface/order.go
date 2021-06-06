package main

import "fmt"

type Clac interface{
  //计算支付价格
  ClacTotalPlice(num int, plice float64) float64
  ClacTotalFee(num int, plice float64, discount float64) float64
}

type Product struct{
  //购买数量
   num int
   //商品单价
   price float64
   //优惠
   discount float64
}

func (p Product) ClacTotalPrice() float64  {
  return float64(p.num)*p.price
}

func (p Product) ClacTotalFee()float64{
  return float64(p.num)*p.price-p.discount
}

func main()  {
  p := Product{num:2,price:88.8,discount:10}
  payPrice := p.ClacTotalPrice()
  fmt.Printf("payPrice: %d \n", payPrice)

  totalFee := p.ClacTotalFee()
  fmt.Printf("totalFee: %d ", totalFee)


}
