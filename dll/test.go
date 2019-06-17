package dll

import (
	"fmt"
	"syscall"
)

func OpenDll() {
	lib := syscall.NewLazyDLL("FPModule_SDK.dll")
	//if err != nil {
	//	fmt.Errorf("load dll")
	//}
	test := lib.NewProc("FPModule_OpenDevice")
	test1 := lib.NewProc("FPModule_FPEnroll")

	fmt.Println("dll:", lib.Name)
	fmt.Println("dll:", test)
	fmt.Println("dll:", test1)

}
