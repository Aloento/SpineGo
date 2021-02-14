package SpineGo

import (
	"SpineGo/utils"
)

type slotData struct {
	index          int
	name           string
	boneData       *BoneData
	color          utils.Color
	darkColor      utils.Color
	attachmentName string
	blendMode      utils.BlendMode
}

func NewSlotData(index int, name string, data *BoneData) *slotData {
	d := new(slotData)
	d.index = index
	d.name = name
	d.boneData = data
	d.color = utils.Color{
		R: 1,
		G: 1,
		B: 1,
		A: 1,
	}
	return d
}
