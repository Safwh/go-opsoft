// 键鼠

package opsoft

import ole "github.com/go-ole/go-ole"

// GetCursorPos ...
func (com *OpSoft) GetCursorPos(x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.op.CallMethod("GetCursorPos", &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

// GetKeyState ...
func (com *OpSoft) GetKeyState(vkCode int) int {
	ret, _ := com.op.CallMethod("GetKeyState", vkCode)
	return int(ret.Val)
}

// KeyDown ...
func (com *OpSoft) KeyDown(vkCode int) int {
	ret, _ := com.op.CallMethod("KeyDown", vkCode)
	return int(ret.Val)
}

// KeyDownChar ...
func (com *OpSoft) KeyDownChar(keyStr string) int {
	ret, _ := com.op.CallMethod("KeyDownChar", keyStr)
	return int(ret.Val)
}

// KeyPress ...
func (com *OpSoft) KeyPress(vkCode int) int {
	ret, _ := com.op.CallMethod("KeyPress", vkCode)
	return int(ret.Val)
}

// KeyPressChar ...
func (com *OpSoft) KeyPressChar(keyStr string) int {
	ret, _ := com.op.CallMethod("KeyPressChar", keyStr)
	return int(ret.Val)
}

// KeyUp ...
func (com *OpSoft) KeyUp(vkCode int) int {
	ret, _ := com.op.CallMethod("KeyUp", vkCode)
	return int(ret.Val)
}

// KeyUpChar ...
func (com *OpSoft) KeyUpChar(keyStr string) int {
	ret, _ := com.op.CallMethod("KeyUpChar", keyStr)
	return int(ret.Val)
}

// LeftClick 按下鼠标左键
func (com *OpSoft) LeftClick() int {
	ret, _ := com.op.CallMethod("LeftClick")
	return int(ret.Val)
}

// LeftDoubleClick 双击鼠标左键
func (com *OpSoft) LeftDoubleClick() int {
	ret, _ := com.op.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

// LeftDown 按住鼠标左键
func (com *OpSoft) LeftDown() int {
	ret, _ := com.op.CallMethod("LeftDown")
	return int(ret.Val)
}

// LeftUp ...
func (com *OpSoft) LeftUp() int {
	ret, _ := com.op.CallMethod("LeftUp")
	return int(ret.Val)
}

// MiddleClick ...
func (com *OpSoft) MiddleClick() int {
	ret, _ := com.op.CallMethod("MiddleClick")
	return int(ret.Val)
}

// MiddleDown ...
func (com *OpSoft) MiddleDown() int {
	ret, _ := com.op.CallMethod("MiddleDown")
	return int(ret.Val)
}

// MiddleUp ...
func (com *OpSoft) MiddleUp() int {
	ret, _ := com.op.CallMethod("MiddleUp")
	return int(ret.Val)
}

// MoveR ...
func (com *OpSoft) MoveR(rx, ry int) int {
	ret, _ := com.op.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

// MoveTo ...
func (com *OpSoft) MoveTo(x, y int) int {
	ret, _ := com.op.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

// MoveToEx ...
func (com *OpSoft) MoveToEx(x, y, w, h int) string {
	ret, _ := com.op.CallMethod("MoveToEx", x, y, w, h)
	return ret.ToString()
}

// RightClick ...
func (com *OpSoft) RightClick() int {
	ret, _ := com.op.CallMethod("RightClick")
	return int(ret.Val)
}

// RightDown ...
func (com *OpSoft) RightDown() int {
	ret, _ := com.op.CallMethod("RightDown")
	return int(ret.Val)
}

// RightUp ...
func (com *OpSoft) RightUp() int {
	ret, _ := com.op.CallMethod("RightUp")
	return int(ret.Val)
}

// WaitKey ...
func (com *OpSoft) WaitKey(vkCode, timeOut int) int {
	ret, _ := com.op.CallMethod("SetSimMode", vkCode, timeOut)
	return int(ret.Val)
}

// WheelDown ...
func (com *OpSoft) WheelDown() int {
	ret, _ := com.op.CallMethod("WheelDown")
	return int(ret.Val)
}

// WheelUp ...
func (com *OpSoft) WheelUp() int {
	ret, _ := com.op.CallMethod("WheelUp")
	return int(ret.Val)
}
