package SpineGo

import (
	"os"
)

const (
	BONE_ROTATE     byte = 0
	BONE_TRANSLATE  byte = 1
	BONE_SCALE      byte = 2
	BONE_SHEAR      byte = 3
	SLOT_ATTACHMENT byte = 0
	SLOT_COLOR      byte = 1
	SLOT_TWO_COLOR  byte = 2
	PATH_POSITION   byte = 0
	PATH_SPACING    byte = 1
	PATH_MIX        byte = 2
	CURVE_STEPPED   byte = 1
	CURVE_BEZIER    byte = 2
)

func ReadSkeletonData(file *os.File) {

}
