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
	b.tempColor1 = utils.Color{}
	b.tempColor2 = utils.Color{}
	b.linkedMeshes = *utils.NewArray(0, 0)
	b.attachmentLoader = *attachments.NewAtlasAttachmentLoader(atlas)
	b.scale = 1
	return b
}

func (b *skeletonBinary) readSkeletonData(file *os.File) (skeletonData *SkeletonData) {
	scale := b.scale
	skeletonData = NewSkeletonData()
	skeletonData.name = func() string {
		name := file.Name()
		dot := strings.LastIndex(name, ".")
		if dot == -1 {
			return name
		}
		return name[0:dot]
	}()

	input := utils.NewSkeletonInput(file)
	skeletonData.hash = input.ReadString()
	skeletonData.version = input.ReadString()
	skeletonData.x = input.ReadFloat()
	skeletonData.y = input.ReadFloat()
	skeletonData.width = input.ReadFloat()
	skeletonData.height = input.ReadFloat()
	nonessential := input.ReadBool()
	if nonessential {
		skeletonData.fps = input.ReadFloat()
		skeletonData.imagesPath = input.ReadString()
		skeletonData.audioPath = input.ReadString()
	}

	n := input.ReadInt(true)
	input.Strings = *utils.NewArray(0, n)
	o := *input.Strings.SetSize(n)
	for i := 0; i < n; i++ {
		o[i] = input.ReadString()
	}
	n = input.ReadInt(true)
	o = *skeletonData.bones.SetSize(n)
	for i := 0; i < n; i++ {
		name := input.ReadString()
		var parent *BoneData
		if i != 0 {
			parent = skeletonData.bones.Get(input.ReadInt(true)).(*BoneData)
		}
		data := NewBoneData(i, name, parent)
		data.rotation = input.ReadFloat()
		data.x = input.ReadFloat() * scale
		data.y = input.ReadFloat() * scale
		data.scaleX = input.ReadFloat()
		data.scaleY = input.ReadFloat()
		data.shearX = input.ReadFloat()
		data.shearY = input.ReadFloat()
		data.length = input.ReadFloat() * scale
		data.transformMode = TransformMode(input.ReadInt(true))
		data.skinRequired = input.ReadBool()
		if nonessential {

		}
	}

}
