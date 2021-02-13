package SpineGo

import (
	"SpineGo/attachments"
	"SpineGo/utlis"
	"os"
)

const (
	BoneRotate     byte = 0
	BoneTranslate  byte = 1
	BoneScale      byte = 2
	BoneShear      byte = 3
	SlotAttachment byte = 0
	SlotColor      byte = 1
	SlotTwoColor   byte = 2
	PathPosition   byte = 0
	PathSpacing    byte = 1
	PathMix        byte = 2
	CurveStepped   byte = 1
	CurveBezier    byte = 2
)

type SkeletonBinary struct {
	tempColor1       utlis.Color
	tempColor2       utlis.Color
	attachmentLoader attachments.AttachmentLoader
}

func ReadSkeletonData(file *os.File) {

}
