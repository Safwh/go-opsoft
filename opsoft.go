package opsoft

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// OpSoft ...
type OpSoft struct {
	op       *ole.IDispatch
	IUnknown *ole.IUnknown
}

// New return *OpSoft.OpSoft
func New() (op *OpSoft, err error) {
	var com OpSoft
	// 创建对象
	ole.CoInitialize(0)
	com.IUnknown, err = oleutil.CreateObject("op.opsoft")
	if err != nil {
		return nil, err
	}
	// 查询接口
	com.op, err = com.IUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return &com, nil
}

// Release 释放
func (com *OpSoft) Release() {
	com.IUnknown.Release()
	com.op.Release()
	ole.CoUninitialize()
}
