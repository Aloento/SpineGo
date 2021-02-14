package SpineGo

import (
	"SpineGo/attachments"
	"SpineGo/utils"
	"os"
	"strings"
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

type skeletonBinary struct {
	tempColor1       utils.Color
	tempColor2       utils.Color
	linkedMeshes     utils.Array
	attachmentLoader attachments.AttachmentLoader
	scale            float32
}

func NewSkeletonBinary(atlas TextureAtlas) *skeletonBinary {
	b := new(skeletonBinary)
	b.tempColor1 = *utils.NewColor()
	b.tempColor2 = *utils.NewColor()
	b.linkedMeshes = *utils.NewArray(0, 0)
	b.attachmentLoader = *attachments.NewAtlasAttachmentLoader(atlas)
	b.scale = 1
	return b
}

func (b *skeletonBinary) readSkeletonData(file *os.File) (skeletonData *SkeletonData) {
	scale := b.scale
	skeletonData = new(SkeletonData)
	skeletonData.name = func() string {
		name := file.Name()
		dot := strings.LastIndex(name, ".")
		if dot == -1 {
			return name
		}
		return name[0:dot]
	}()

}
