package SpineGo

import "SpineGo/utils"

type SkeletonData struct {
	name    string
	hash    string
	version string
	x       float32
	y       float32
	width   float32
	height  float32
	arrays
	nonessential
}

type arrays struct {
	bones                utils.Array
	slots                utils.Array
	skins                utils.Array
	events               utils.Array
	animations           utils.Array
	ikConstraints        utils.Array
	transformConstraints utils.Array
	pathConstraints      utils.Array
}

type nonessential struct {
	fps        float32
	imagesPath string
	audioPath  string
}

func NewSkeletonData() *SkeletonData {
	d := new(SkeletonData)
	d.bones = *utils.NewArray(0, 0)
	d.slots = *utils.NewArray(0, 0)
	d.skins = *utils.NewArray(0, 0)
	d.events = *utils.NewArray(0, 0)
	d.animations = *utils.NewArray(0, 0)
	d.ikConstraints = *utils.NewArray(0, 0)
	d.transformConstraints = *utils.NewArray(0, 0)
	d.pathConstraints = *utils.NewArray(0, 0)
	return d
}
