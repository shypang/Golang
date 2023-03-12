package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   //4바이트
	Score float64 //8바이트
}

func main() {
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user))
	//12바이트로 나올 줄 알았지만 16바이트가 출력된다.
	//이유는 메모리 정렬때문이다.
	//레지스터 크기가 8바이트인 컴퓨터를 64비트
	//레지스터 크기가 4바이트인 컴퓨터를 32비트

	//레지스터 크기의 배수에 맞게 메모리크기를 가져온다.
	//64비트이므로 8의 배수대로 메모리를 가져오려면 16바이트를 가져와야함
	//메모리 정렬을 위해 필드사이에 공간을 띄우는 것을 메모리 패딩이라고 하는데 현재 16바이트가 할당된 경우도 메모리 패딩이 된 경우도 마찬가지다. 
}
