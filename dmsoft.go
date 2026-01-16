//go:build windows

//export GOARCH=386

package dmsoft

import (
	_ "embed"
	"encoding/json"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

//go:embed dll.json
var DllJson []byte

var (
	DllM = make(map[string]string)
	_    = json.Unmarshal(DllJson, &DllM)
)

type Dmsoft struct {
	dm       *ole.IDispatch
	IUnknown *ole.IUnknown
}

func NewDmsoft() *Dmsoft {
	var com Dmsoft
	var err error
	err = ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	if err != nil {
		panic(err.Error() + "___NewDmsoft")
	}
	com.IUnknown, err = oleutil.CreateObject(DllM["Willpwr"])
	if err != nil {
		panic(err)
	}
	com.dm, err = com.IUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		panic(err)
	}
	return &com
}

// Release 释放
func (com *Dmsoft) Release() {
	com.IUnknown.Release()
	com.dm.Release()
	ole.CoUninitialize()
}

// SetDllPathA Ascii
func SetDllPathA(path string, mode int) bool {
	var _p0 *uint16
	dmReg32 := syscall.NewLazyDLL(DllM["DmReg.dll"])
	_SetDllPathA := dmReg32.NewProc(DllM["SetDllPathA"])
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _SetDllPathA.Call(uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	return ret != 0
}

// SetDllPathW Unicode
func SetDllPathW(path string, mode int) bool {
	var _p0 *uint16
	dmReg32 := syscall.NewLazyDLL(DllM["DmReg.dll"])
	_SetDllPathW := dmReg32.NewProc(DllM["SetDllPathW"])
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _SetDllPathW.Call(uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	return ret != 0
}
