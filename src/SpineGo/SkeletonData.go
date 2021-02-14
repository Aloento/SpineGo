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
