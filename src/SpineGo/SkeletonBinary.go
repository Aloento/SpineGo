package SpineGo

import (
	"SpineGo/attachments"
	"SpineGo/utils"
	"os"
    "gopkg.in/guregu/null.v3"   //This is a pack used for null function. Not certificated.
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
	return b
}

func ReadSkeletonData(file *os.File) *skeletonData {

}

func readSkeletonData(FileHandle file) *skeletonData{

    if file == null{
        //Not realized the throw
        //panic()

    }

    scale := skeletonBinary.scale
    skeletonData := new(SkeletonData)
    skeletonData.name = file.nameWithoutExtension()

    if Loader.spineVersion > 37{

        if  input := new(SkeletonInput(file)) {
            panic()
        }

        skeletonData.hash = input.readString()
        if skeletonData.hash.isEmpty(){
            skeletonData.hash = null
        }
        if skeletonData.version.isEmpty(){
            skeletonData.version = null
        }

        skeletonData.x = input.readFloat()
        skeletonData.y = input.readFloat()
        skeletonData.width = input.readFloat()
        skeletonData.height = input.readFloat()

        nonessential := input.readBoolean()

        if nonessential{
            skeletonData.fps = input.readFloat()
            skeletonData.imagesPath = input.readString()

            if skeletonData.imagesPath.isEmpty() {
                skeletonData.imagesPath = null
            }
            skeletonData.audioPath = input.readString()

            if skeletonData.imagesPath.isEmpty() {
                skeletonData.audioPath = null
            }

        }

        var n int
        var o




        input.strings := new(Array)




    }
}
