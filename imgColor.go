// 图色

package opsoft

import ole "github.com/go-ole/go-ole"

// Capture ...
func (com *OpSoft) Capture(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.op.CallMethod("Capture", x1, y1, x2, y2, file)
	return int(ret.Val)
}

// CmpColor ...
func (com *OpSoft) CmpColor(x int, y int, color string, sim float32) int {
	ret, _ := com.op.CallMethod("CmpColor", x, y, color, sim)
	return int(ret.Val)
}

// FindColor ...
func (com *OpSoft) FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FindColor", x1, y1, x2, y2, color, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

// FindColorEx ...
func (com *OpSoft) FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.op.CallMethod("FindColorEx", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

// FinopultiColor  ...
func (com *OpSoft) FinopultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FinopultiColor", x1, y1, x2, y2, firstColor, offsetColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

// FinopultiColorEx ...
func (com *OpSoft) FinopultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.op.CallMethod("FinopultiColorEx", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

// FindPic ...
func (com *OpSoft) FindPic(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FindPic", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Value().(int32))
	*intY = int(y.Value().(int32))
	x.Clear()
	y.Clear()
	return int(ret.Value().(int32))
}

// FindPicEx ...
func (com *OpSoft) FindPicEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.op.CallMethod("FindPicEx", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

// GetColor ...
func (com *OpSoft) GetColor(x, y int) string {
	ret, _ := com.op.CallMethod("GetColor", x, y)
	return ret.ToString()
}

// CapturePre ...
func (com *OpSoft) CapturePre(file string) int {
	ret, _ := com.op.CallMethod("CapturePre", file)
	return int(ret.Val)
}

// EnableDisplayDebug ...
func (com *OpSoft) EnableDisplayDebug(enableDebug int) int {
	ret, _ := com.op.CallMethod("EnableDisplayDebug", enableDebug)
	return int(ret.Val)
}

// GetScreenData ...
// GetScreenDataBmp ...
// SetDisplayInput ...
