package originLink

import _ "unsafe"

//go:linkname funcImpl testlinkname/targetLink.funcUse
func funcImpl() string {
	return "link funcUse success"
}
