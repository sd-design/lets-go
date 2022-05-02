package main

import "fmt"

func main() {

	var num1, num2 string
	fmt.Scan(&num1, &num2)

	lenght := len(num1)
	lenght2 := len(num2)
	for i := 0; i < lenght; i++ {

		for i2 := 0; i2 < lenght2; i2++ {
			if num1[i] == num2[i2] {
				fmt.Printf("%s ", string(num1[i]))
			}
		}
	}
}

//Decision without string
// func main() {
//     var a,b,an, bn int
//     fmt.Scan(&a,&b)

//     for i:=1000;i>=1 && a>0;i=i/10 {
//         if a/i ==0 {continue}
//         an = a/i%10
//         for i:=1;i<10000 && b>0;i=i*10 {
//             bn = b/i%10
//             if bn == an {fmt.Print(an," "); break}
//         }
//     }
// }

// Recursive decision
// func matchDigits(x, y int){
//     if x > 10{
//         matchDigits(x / 10, y)
//         matchDigits(x % 10, y)
//     } else if y > 10{
//         matchDigits(x, y / 10)
//         matchDigits(x, y % 10)
//     } else if (x == y) {
//         fmt.Print(x, " ")
//     }
// }

// func main(){
//     var a, b int

//     fmt.Scan(&a, &b)
//     matchDigits(a, b)
// }
