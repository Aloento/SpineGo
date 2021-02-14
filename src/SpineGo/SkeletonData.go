package SpineGo

type SkeletonData struct {
	name    string
	hash    string
	version string
	x       float32
	y       float32
	width   float32
	height  float32
	Nonessential
}

type Nonessential struct {
	fps        float32
	imagesPath string
	audioPath  string
}
