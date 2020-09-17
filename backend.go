// 后台设置

package opsoft

// BindWindow 绑定指定的窗口,并指定这个窗口的屏幕颜色获取方式,鼠标仿真模式,键盘仿真模式,以及模式设定
func (com *OpSoft) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.op.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	return int(ret.Val)
}

// UnBindWindow 解除绑定窗口,并释放系统资源
func (com *OpSoft) UnBindWindow() int {
	ret, _ := com.op.CallMethod("UnBindWindow")
	return int(ret.Val)
}
