

package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
	

	// if an array does not initialized explicitly, then the default value of this array is 0

	var myarr[3]string


	myarr[0]="aa"
	myarr[1]="aa"
	myarr[2]="aa"


	fmt.Println("Elements of array:")
	fmt.Println(myarr[0])
	fmt.Println(myarr[1])
	fmt.Println(myarr[2])



	myarr0:=[...]int{1,2,3}
	for i:=0;i<len(myarr0);i++ {
		fmt.Println(myarr0[i])
		fmt.Printf("%d\n",myarr0[i])
	}

	myarr00:=[...]int{1,2,4}
	// we can directly compare two arrays of same type using == operator

	fmt.Println(myarr0==myarr00)

	myarr1:=[3]string{"abc","de","fg"}
	
	fmt.Println("elements directly define")

	for i:=0;i<3;i++{
		fmt.Println(myarr1[i])
	}


	var myarr2[2][2] int
	myarr2[0][0]=100
	myarr2[0][1]=200
	myarr2[1][0]=300
	myarr2[1][1]=400
	fmt.Println("Elements multidimensional")
	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			fmt.Println(myarr2[i][j])
		}
	}
	
	myarr3:=[2][2]string{{"aa","bb"},{"cc","dd"}}

	fmt.Println("Elements multidimensional directly define")
	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			fmt.Println(myarr3[i][j])
		}
	}


	// var input string
	// fmt.Scanln(&input)
	// fmt.Println(input)

	// func Scanf(format string, a ...interface{}) (n int, err error)
	// fmt.Scanf("%s", &name)
	// fmt.Scanf("%d", &alphabet_count)


	//var size int
	//fmt.Scanln(&size)
	//size:=22
	var myarr4[3] int
	fmt.Println(myarr4)

        var slice = make([]int, 1, 3)
        fmt.Println(slice)
        fmt.Println(len(slice))
        fmt.Println(cap(slice))
	aaa:=&slice
	bbb:=&myarr4
        fmt.Println("address",aaa)
        fmt.Println("address",bbb)

	fmt.Println()
	fmt.Printf("address: %p , %v , %T , %p, %v",aaa,*aaa,aaa,&aaa,&aaa)
	fmt.Println()
	fmt.Printf("address: %p , %v , %T , %p, %v",bbb,*bbb,bbb,&aaa,&aaa)
	fmt.Println()
	
    	i:= int(42)
        fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", i, &i, i)
	fmt.Println()
	p := &i
	fmt.Printf("2. main  -- p %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)

        slice = append(slice, 6, 0, 2, 4, 3, 1)
        fmt.Println(slice)
	



	fmt.Println()
	fmt.Println()
	fmt.Println()
        var slice1 = make([]int, 0,8)
	fmt.Println(slice1)
	s1:=&slice1
	fmt.Printf("type of s1 :  %T  \n ",s1)
	fmt.Printf("%p  %T  %v  %T\n",&slice1,s1,*s1,s1)


	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice1))
	data := (unsafe.Pointer(hdr.Data))

	// reflect.SliceHeader contains a Data field which contains a pointer to the underlying array of a slice.
	fmt.Println("hdr   ",hdr)
	fmt.Println("hdr.Data   ",hdr.Data)
	hdrData:=hdr.Data
	fmt.Printf("hdrData  %T \n",hdrData)
	fmt.Println("&hdrData  ",&hdrData)
	arrayOfSlice:=&hdrData
	fmt.Printf("arrayOfSlice   %T \n",*arrayOfSlice)

	slice1=append(slice1,1,2,3)
	hdr1 := (*reflect.SliceHeader)(unsafe.Pointer(&slice1))
	data1 := (unsafe.Pointer(hdr1.Data))
	hdr1Data:=hdr1.Data
	fmt.Printf("hdrData  %T \n",hdr1Data)
	fmt.Println("&hdrData  ",&hdr1Data)
	arrayOfSlice1:=&hdr1Data
	fmt.Println("hdr1.Data  ",hdr1.Data)
	fmt.Println("&hdr1.Data  ",&hdr1Data)

	fmt.Println(slice1)
	fmt.Println(data)
	fmt.Println(data1)
	fmt.Println(arrayOfSlice1)
	fmt.Printf("type %T \n",data)
	dataUintptr:=uintptr(data)
	fmt.Println(dataUintptr)
	fmt.Printf("type   %T \n",dataUintptr)

	fmt.Println()
	fmt.Println()
	fmt.Println()

        var array1[]int
	a1:=&array1
	fmt.Printf("%p  %T  %v\n",&array1,a1,*a1)

	
	var testar=make([]int,0,4)
	testar=append(testar,22)
	testar=append(testar,33)
	testar=append(testar,33)
	testar=append(testar,33)
	fmt.Println(len(testar))
	fmt.Println(cap(testar))
	fmt.Println(testar)
	fmt.Println(testar[1])
	//xx:=22	
	var arrxx[10]int
	fmt.Println(arrxx)


	for i:=0;i<len(arrxx);i++ {
		fmt.Println(arrxx[i])
	}

	var arrll[]int
	arrll=append(arrll,1,2,3,4)
	fmt.Println(arrll)
	fmt.Println(len(arrll))
	fmt.Println(cap(arrll))

	for i:=0;i<len(arrll);i++ {
		fmt.Println(arrll[i])
	}

}
