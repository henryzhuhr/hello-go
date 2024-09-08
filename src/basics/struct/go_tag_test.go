package go_struct

import (
	"fmt"
	"reflect"
	"testing"
)

type Programmer struct {
	Name     string   `type:"string" json:"name"`
	Age      int      `type:"int" yaml:"age" `
	Language []string `properties:"language"`
}

func TestTag(t *testing.T) {
	programmer := &Programmer{"Tom", 20, []string{"Go", "Python"}}
	fmt.Println("Programmer:")
	fmt.Println("  Tag   =", reflect.TypeOf(*programmer).Field(0).Tag)
	fmt.Println("  Value =", programmer.Name)
	fmt.Println("  Tag   =", reflect.TypeOf(*programmer).Field(1).Tag)
	fmt.Println("  Value =", programmer.Age)
}
