// when we use things as just wrapped up we call them functions and when these functions go into the classes we call them methods.
package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang, No super or parent.

	madhav := User{"Madhav", "madhav@google.com", true, 18}
	fmt.Println(madhav)

	fmt.Printf("Details of madhav are: %+v \n", madhav)
	fmt.Printf("Name is: %v\n", madhav.Name)

	madhav.GetStatus()
	madhav.NewMail()

	// method ke andar se property change krenge toh original me change nahi hogi, sirf uss method me change hogi
	// which leads us to a conclusion ki method ke andar ek copy pass on hoti hai.
	fmt.Println(madhav.Email)






	// using pointer...

	kritika := User{Name: "kritika", Email: "kbr@gmail.com", Status: false, Age: 16}

	ptrKritika := &kritika

	ptrKritika.UpdateAge()
	fmt.Println(kritika.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int //this oneage is not gonna be exportable because its first letter is not capital
}

// -------------------------making a method for this struct
// kis class ka hai,
//  class ki ek property
//  bhi pass kr skte
//  hain.				// capital to make it exportable
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}


func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}

func (u User) UpdateAge() {

	u.Age = 18
	
}



// -----------======================-===-=---=-==-=-
// so I just discovered this, ki mai "method ko chahe pointer se call kru ya normally it doesn't matter",
// agr maine definition me * lagaya hai toh muje pointer milega(actually, wo original milega jisse bina kisi * ke use kr skte) nahi toh copy

