package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)
var i=32 // (smallLetter) camelcase to use within package itself
var I=23 // (Caps) pascal scale to provide accessibility outside of package as well

const (
	User string = "Admin" // constant
	Product string ="Product"
)

// Enumerated constant
const (
	m = iota +1  //0+1=1
	_  // skip 
	_  //skip
	n     //4
	o     //5
)

// Student struct
type Student struct{
	name string // make key as pascalCase then only we can use the fields in other package as well
	rollNo int
	subjects []string
}

// Composition relatin(has-a) - Embedding

//Processor struct
type Processor struct{
	processorName string
	cores int
} 

//Memory struct
type Memory struct{
	memoryCapacity int
	memoryType string
}

// Computer struct
type Computer struct{
	Processor
	Memory
	price int
}

// Foo struct
type Foo struct{
	bar int
}

type rectangle struct {
	width, height int
}

func main() {
	var i int
	i=12
	j:=12 //creates a new variable j
	var foo string = string(i)
	var foo2 string = strconv.Itoa(i)
	var bar bool=true

	fmt.Printf("%v, %T\n",i,i)
	fmt.Printf("%v, %T\n",j,j)
	fmt.Printf("%v, %T\n",foo,foo)
	fmt.Printf("%v, %T\n",foo2,foo2)
	fmt.Printf("%v, %T\n",bar,bar)

	var k int8=12
	var k2 int16=int16(k) // Explicit typecasting is required
	fmt.Printf("%v, %T\n",k,k)
	fmt.Printf("%v, %T\n",k2,k2)

	p:= 20
	q:= 3

	fmt.Println(p+q)
	fmt.Println(p-q)
	fmt.Println(p*q)
	fmt.Println(p/q)
	fmt.Println(p%q)

	fmt.Println(p & q)
	fmt.Println(p | q)
	fmt.Println(p ^ q)
	fmt.Println(p &^ q)


	r:= 3.14
	s:= 1.7e12
	t:= 2.3e12
    fmt.Printf("%v %T\n",r,r)
	fmt.Printf("%v %T\n",s,s)
	fmt.Printf("%v %T\n",t,t)

	var u complex64=1+2i // float 32 is used internally
	var v complex128= 2i //float 64 is used
	fmt.Printf("%v, %T\n",u,u)
	fmt.Printf("%v, %T\n",real(u),real(u))
	fmt.Printf("%v, %T\n",imag(u),imag(u))
	fmt.Printf("%v, %T\n",v,v)

	var w complex64 = complex(2,4)
	fmt.Printf("%v, %T\n",real(w),real(w))
	fmt.Printf("%v, %T\n",imag(w),imag(w))

	e:= "This is a string"
	e1:= "This is another String" 
	e2:= []byte(e) // utf 8
	e3:='T' //rune - utf 32

	fmt.Printf("%v,%T\n",e+e1,e)
	fmt.Printf("%v,%T\n",string(e[1]),e) // extracts 1st character. cast it to get in character format rather than byte format
	fmt.Printf("%v,%T\n",e2,e2)
	fmt.Printf("%v,%T\n",e3,e3)

	const d int =12
	const g float32 = 3.14
	const h string = "This is a constant"
	fmt.Println(d)
	fmt.Println(g)
	fmt.Println(h)

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)

	var amounts [3]int = [3]int{10,20,30}
	amt := [3]int{30,40,50}
	amt1 := [...]int{1,2,3,4,5,6}
	fmt.Printf("Amount: %v\n",amounts)
	fmt.Printf("Amount: %v\n",amt)
	fmt.Printf("%v\n",len(amt))
	fmt.Printf("%v\n",len(amt1))
	amounts[0]= 20
	fmt.Printf("Amount: %v\n",amounts)
	fmt.Printf("Amount: %v\n",amounts[1])

	amtCp := &amounts // willn't copy but provide reference of memory location bcz of &
	fmt.Printf("AmountCP: %v\n",amtCp)

	b1 := [...]int{1,2,3,4,5,6,7}
	b2 := b1[:len(b1)-1]
	b3 := b1[2:]
	b4 := b1[:5]
	b5 := b1[2:7]

	fmt.Printf("b1: %v\n",b1)
	fmt.Printf("b2: %v\n",b2)
	fmt.Printf("b3: %v\n",b3)
	fmt.Printf("b4: %v\n",b4)
	fmt.Printf("b5: %v\n",b5)

	// MultiDimension Array
	var identityMatrix [3][3]int =[3][3]int{
		[3]int{1,0,0},
		[3]int{0,1,0},
		[3]int{0,0,1},
	}

	fmt.Println(identityMatrix)
	fmt.Println(identityMatrix[0][0])
	identityMatrix[1][2]=7
	fmt.Println(identityMatrix)

	//Slices
	var d1 []int =[]int{1,2,3}
	var d2 []int = d1  //copy the pointer location
	fmt.Println(d1)
	fmt.Println(d2)
	d1[0]=4 

	fmt.Println(len(d1)) // Length
	fmt.Println(cap(d2)) //Capacity

	var f1 []int = make([]int,3,10) // wit size and capacity specified
	fmt.Println(len(f1)) 
	fmt.Println(cap(f1)) 

	var g1 []int =[]int{1,2,3}
	var g2 []int =append(g1[1:],5)
	var g3 []int =append(g2,g1...) // use spread operator to extract value and assign to slice c

	fmt.Println(g1)
	fmt.Println(g2)
	fmt.Println(g3)

	shoppingCart := map[string]int{
		"Keyboard": 100,
		"Mouse": 50,
		"Laptop": 1000,
	}
	shoppingCart["Laptop"] = 1400
	shoppingCart["Monitor"] = 1200
	fmt.Println(shoppingCart)
	fmt.Println(len(shoppingCart))
	fmt.Println(shoppingCart["Keyboard"])
	_,ok :=shoppingCart["Mobile"]
	fmt.Println(ok)

	sc := shoppingCart  // refers to pointer so change to sc will reflect to shoppingCart as well
	fmt.Println(sc)
	sc["Monitor"]= 1350 
	fmt.Println(sc)
	fmt.Println(shoppingCart)

	delete(shoppingCart,"Monitor")
	fmt.Println(shoppingCart)

	//shoppingCart2= make(map[string]int)

	// struct impl
	student1 := Student{
		name: "Sanjay", // Even if remove the keyname also then it is correct but not recommended
		rollNo: 3,
		subjects: []string{
			"Maths",
			"Physics",
			"Chemistry",
		},
	}
	fmt.Println(student1)
	fmt.Println(student1.name)
	fmt.Println(student1.subjects[0])
	student1.rollNo=6
	fmt.Println(student1)

	student2 := student1  // this is value reference not pointer reference but we can also make it pointerRef using$ as shown:  student2 := &student1 
	student2.name="Manoj"
	fmt.Println(student1)
	fmt.Println(student2)


	computer := Computer{}
	computer.price = 5000
	computer.processorName = "Intel i5"
	computer.cores =6
	computer.memoryCapacity = 8
	computer.memoryType= "DDR4"

	fmt.Println(computer)

	computer2 := Computer{
		Processor: Processor{
			processorName: "Intel i7",
			cores: 8,
		},
		Memory: Memory{
			memoryCapacity: 16,
			memoryType: "DDR4",
		},
		price: 70000,
	}

	fmt.Println(computer2)
	fmt.Println(computer2.Processor)
	fmt.Println(computer2.processorName)

	// IF-ELSE Control statements
	if true {
		fmt.Println("This is simple if statement")
	}

	if i :=2; i==3{
		fmt.Println("This is true")
	}else if i==2{
		fmt.Println("This is simple else-if block")
	}else{
		fmt.Println("This is simple else block")
	}

	i7 :=10
	j7 :=20
	if i7>0 && j7>0{
		fmt.Println("i3 is greater than 0")
	}

	shoppingCart3 := make(map[string]int)
	shoppingCart3 = map[string]int{
		"Keyboard" : 100,
		"Mouse" : 50,
		"Laptop" : 1000,
	}
	shoppingCart["Laptop"] =1500
	shoppingCart3["Monitor"]= 1200

	if _,ok := shoppingCart3["Monitor"]; ok {  // If item exists then ok will be true otherwise it will be false
		fmt.Println("Item exists in the shopping cart")
	}

	// Switch case
	switch i :=2+3; i {
	case 1,3,5,7,9:
		fmt.Println("This is odd")
	case 2,4,6,8: 
		fmt.Println("This is even")
	default:
		fmt.Println("This is default")
		
	}

	i6 := 2+3
	switch {
	case i6>0:
		fmt.Println("This is odd")
	case i6<5: 
		fmt.Println("This is even")
	default:
		fmt.Println("This is default")	
	}

	var i4 interface{} = 5
	switch i4.(type) {
	case int:
		fmt.Println("This is Int type")
		fmt.Println("This is Int type")
		break // No need to keep break but if required we can use
		fmt.Println("This is Int type")
	case float64:
		fmt.Println("This is float64 type")
	case string:
		fmt.Println("This is string type")
	default:
		fmt.Println("This is default")

	}

	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough // Executes the underneath statement as well
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
		
	}

	// For loop
	for i7 := 0; i7 < 5; i7++ {
		fmt.Println(i)
	}
	for i7,j7 := 0,0; i7 < 5; i7,j7= i7+1,j7+1 {
		fmt.Println(i)
	}

	i8,j8 :=0,0
	for ; i8<5;i8,j8=i8+1,j8+1{
		fmt.Println(i8,j8)
	}
	fmt.Println(i8,j8)


	i9,j9 :=0,0
	for i9<10{
		fmt.Println(i9,j9)
		i9,j9=i9+1,j9+1
		if(i9==6 && j9==6){
			break
		}
	}
	fmt.Println(i9,j9)

	i10,j10 :=0,0
	for { // behaves as while or do-while
		fmt.Println(i10,j10)
		i10,j10=i10+1,j10+1
		if(i10==6 && j10==6){
			break // use continue to skip
		}
	}
	fmt.Println(i10,j10)

	// Nested for loop
	for k2 :=0; k2<5; k2++{
		for l2 :=0; l2 <5; l2++{
			if i*j==0{
				continue
			}
			fmt.Println(i*j)
		}
	}

	a3 := []int{1,2,3,4,5,6,7}
	for p2,q2 := range a3{
		fmt.Println(p2,q2)  // Index, value
	}

	for _,q2 := range a3{
		fmt.Println(q2)  
	}

	shoppingCart4 := make(map[string]int)
	shoppingCart4 = map[string]int{
		"Keyboard" : 100,
		"Mouse" : 50,
		"Laptop" : 1000,
	}
	shoppingCart["Laptop"] =1500
	shoppingCart4["Monitor"]= 1200

	for k,v :=range shoppingCart4{
		fmt.Println(k,v) // to get both key,value
	}
	for k :=range shoppingCart4{
		fmt.Println(k) // to get keys only
	}
	for _,v :=range shoppingCart4{
		fmt.Println(v) // to get values only
	}

	// Pointers
	var a int = 12
	var b *int = &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) // to dereference pointer

	// pointers with struct
	var foo1 *Foo
	fmt.Println(foo1)
	foo1= new(Foo)
	fmt.Println(foo1)
	fmt.Println(*foo1)
	fmt.Println((*foo1).bar)
	foo1.bar= 45  // De-Referencing is taken care by go compiler
	fmt.Println(foo1.bar)

	msg := "Hello User"
	writeMessage("Hello User") // value is passed to function not pointer
	fmt.Println(msg)

	//msg := "Hello User"
	writeMessage2(&msg,"hello") // reference is passed
	fmt.Println(msg)

	sum("Sum: ",1,2,3,4)
	sum("hello",1,2,3)
	total := sum2(1,2,3)
	fmt.Println(total)

	value, error := divide(3,0)
	if error!=nil{
		fmt.Println(error)
		return
	}
	fmt.Println(value)

	// Anonymous function
	func(){
		fmt.Println("This is anonymous function")
	}() // Invocation occurs here due to ()

	fun := func(){
		fmt.Println("Hello")
	}
	fun() // should call anonymous function only after declaration

	rect := rectangle{
		width: 20,
		height: 10,
	}

	area := rect.area2()
	fmt.Println("Area of Rectnagle: ",area)

	var rg geometry = rectangle2{
		width: 3,
		height: 4,
	}
	fmt.Println(rg.area())
	fmt.Println(rg.perimeter())

	var ci  geometry = circle{radius: 5}
	fmt.Println(ci.area())
	fmt.Println(ci.perimeter())

	// Go-routine  - execution happens in its own routine independently
	go writeMsg()
	time.Sleep(100 * time.Millisecond)

	msg2 := "Hii"
	wg.Add(1) // Instead of using sleep using time
	go func(msg string){
		fmt.Println(msg2)
		wg.Done()
	}(msg2)    // To avoid race condition we should pass the variable as well
	msg2 = "Hii people"
	fmt.Println(msg2)
	//time.Sleep(100 * time.Millisecond)
	wg.Wait()
	
	// To check for raceCondition we can use this command i.e.  go run -race Main.go

	var wg sync.WaitGroup
	for i :=1 ;i<=5;i++{
		wg.Add(1)
		go worker(i,&wg)
	}
	wg.Wait()

	// Channel creation
	ch := make(chan int,100) // Made it as buffered channel by adding 100
	wg.Add(2)

	go func(ch <-chan int){ // restricitng to only read channel
		i := <-ch  // receiving data 
		fmt.Println(i)
		//ch <- 27  // Bi-direction channel by passing data to channel back
		for i :=range ch{
			fmt.Println(i)
		}
		for {
			if i,ok := <- ch; ok{
				fmt.Println(i)
			}else{
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int){ // restricitng to only write channel
		ch <- 12  // value is passed to channel
		//fmt.Println(<-ch)
		i=27
		ch <- i
		close(ch) // close the channel
		ch <- 45  // will cause panic as channel was already closed
		wg.Done()
	}(ch)
	wg.Wait()


	R1 := make(chan string)
	R2 := make(chan string)

	go portal1(R1)
	go portal2(R2)

	// Select statement
	select {
	case op1 := <- R1:
		fmt.Println(op1)
	case op2 := <- R2:
		fmt.Println(op2)
	}


}

func portal1(channel1 chan string){
	time.Sleep(2*time.Second)
	channel1 <- "Welcome to channel1"
}

func portal2(channel2 chan string){
	time.Sleep(1*time.Second)
	channel2 <- "Welcome to channel1"
}

var wg = sync.WaitGroup{}

func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Worker %d starting\n",id)
	time.Sleep(time.Second)
	fmt.Println("Worker %d done\n",id)
}

// Function
func writeMessage(msg string){
	fmt.Println(msg)
	msg="Hello Youtube"
}

func writeMessage2(msg *string, msg1 string){ // Dereferened
	*msg="Hello Youtube"
	fmt.Println(*msg)
}

// Variadic function
func sum(msg string, values ...int){ // variadic function should be the last variable
	total :=0
	for _,v := range values{
		total +=v
	}
	fmt.Println(total)

}

func sum2(values ...int) int{ // variadic function should be the last variable
	total :=0
	for _,v := range values{
		total +=v
	}
	return total // return value
}

func divide(a,b float64) (float64,error){
	if b==0.0 || b==0{
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a/b, nil
}

// Methods
func(r rectangle) area2() int{
	return r.width*r.height
}

// Interface - collection of behaviours
type geometry interface{
	area() float64
	perimeter() float64
}

type circle struct{
	radius float64
}

type rectangle2 struct{
	height,width float64
}

func(r rectangle2) area() float64{
	return r.width * r.height
}

func(r rectangle2) perimeter() float64{
	return 2*r.width+2*r.height
}

func(c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func(c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}

func writeMsg(){
	fmt.Println("Hello")
}