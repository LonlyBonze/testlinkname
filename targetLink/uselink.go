package targetLink

import _ "testlinkname/originLink"

func funcUse() string

//TestFuncUse ...
func TestFuncUse() string {
	return funcUse()
}
