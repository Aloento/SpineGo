package SpineGo

import (
	"SpineGo/utils"
)

type BoneData struct {
	index                                                  int
	name                                                   string
	parent                                                 *BoneData
	color                                                  utils.Color
	length, x, y, rotation, scaleX, scaleY, shearX, shearY float32
	transformMode                                          TransformMode
	skinRequired                                           bool
}

func NewBoneData(index int, name string, parent *BoneData) *BoneData {
	d := new(BoneData)
	d.color = utils.Color{
		R: 0.61,
		G: 0.61,
		B: 0.61,
		A: 1,
	}
	d.scaleX, d.scaleY = 1, 1
	d.transformMode = normal

	d.index = index
	d.name = name
	d.parent = parent
	return d
}

type TransformMode int

const (
	normal                 TransformMode = 0
	onlyTranslation        TransformMode = 1
	noRotationOrReflection TransformMode = 2
	noScale                TransformMode = 3
	noScaleOrReflection    TransformMode = 4
)
