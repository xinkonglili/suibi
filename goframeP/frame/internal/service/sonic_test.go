package service

import (
	"testing"

	"github.com/bytedance/sonic"
)

type DriverInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
	Ints []int  `json:"ints"`
}

func TestSonic(t *testing.T) {
	ints := make([]int, 10)
	for i := 0; i < 10; i++ {
		ints = append(ints, i)
	}

	driverInfoStr := `{"name":"driver1","age":18,"sex":"male","ints":[0,1,2,3,4,5,6,7,8,9]}`
	var driverInfo DriverInfo
	err := sonic.UnmarshalString(driverInfoStr, &driverInfo)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v: success", driverInfo)

	marshalString, err := sonic.Marshal(driverInfo)
	marshalString1, err := sonic.MarshalString(driverInfo)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v: success too", marshalString)
	t.Logf("%v: success too", marshalString1)
}
