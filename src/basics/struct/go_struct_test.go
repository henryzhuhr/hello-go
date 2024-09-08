package go_struct

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStruct(t *testing.T) {
	user := NewUser(0, "J", "123")
	user1 := NewUserWithOptions(
		WithId(2),
		WithName("B"),
		WithPhone("345"),
	)
	user2 := NewUserWithOptions(
		WithName("C"),
	)
	fmt.Println("user1=", *user1)
	fmt.Println("user2=", *user2)
	user3 := NewUserWithNameWithOptions("D", WithPhone("000"))
	fmt.Println("user3=", *user3)
	fmt.Println("user id =", user.GetId())
	user.TrySetId(10)
	fmt.Println("user id =", user.GetId())
	user.SetId(20)
	fmt.Println("user id =", user.GetId())
	fmt.Println("Get tag of user:", reflect.TypeOf(user.Phone), reflect.ValueOf(user.Phone))
}
