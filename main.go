package main

import "fmt"

func main() {
	var a,b,c,n,i int
	fmt.Scanf("%d %d %d", &a, &b , &n)
	c = a/b
	for (i:=0; (c%10)!=0; i+){
		c=c*10
		if (n<i){
			break
		}
	}
fmt.Println(i%10)
}
