// 基本设置

package opsoft

// GetBasePath 获取注册在系统中的op.dll的路径.
func (com *OpSoft) GetBasePath() string {
	ret, _ := com.op.CallMethod("GetBasePath")
	return ret.ToString()
}

// GetID 返回当前OP对象的ID值，这个值对于每个对象是唯一存在的
func (com *OpSoft) GetID() int {
	ret, _ := com.op.CallMethod("GetID")
	return int(ret.Val)
}

// GetLastError 获取插件命令的最后错误
func (com *OpSoft) GetLastError() int {
	ret, _ := com.op.CallMethod("GetLastError")
	return int(ret.Val)
}

// GetPath 获取全局路径
func (com *OpSoft) GetPath() string {
	ret, _ := com.op.CallMethod("GetPath")
	return ret.ToString()
}

// SetPath 设置全局路径,设置了此路径后,所有接口调用中,相关的文件都相对于此路径. 比如图片,字库等.
func (com *OpSoft) SetPath(path string) int {
	ret, _ := com.op.CallMethod("SetPath", path)
	return int(ret.Val)
}

// SetShowErrorMsg 设置是否弹出错误信息,默认是打开.
func (com *OpSoft) SetShowErrorMsg(show int) int {
	ret, _ := com.op.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}

// Ver 返回当前插件版本号
func (com *OpSoft) Ver() string {
	ver, _ := com.op.CallMethod("Ver")
	return ver.ToString()
}

// EnablePicCache 设置是否开启或者关闭插件内部的图片缓存机制
func (com *OpSoft) EnablePicCache(enable int) int {
	ret, _ := com.op.CallMethod("EnablePicCache", enable)
	return int(ret.Val)
}
