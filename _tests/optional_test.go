package tests

import (
	"fmt"
	"testing"

	"github.com/nisus/sugargo/lang"
)

func Test1(t *testing.T) {
	v1 := lang.OfNilable(nil)
	fmt.Println(v1.IsPresent())
	fmt.Println(v1.Get())

	v2, err := lang.Of(nil)
	fmt.Println(v2, err)

	v3 := lang.OfNilable(780)
	fmt.Println(v3)
	v3.IfPresent(func(v interface{}) {
		fmt.Printf("哈哈哈, 存在哟. %+v\n", v)
	})

}
