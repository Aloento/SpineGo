package attachments

import "SpineGo"

type atlasAttachmentLoader struct {
	atlas SpineGo.TextureAtlas
}

func NewAtlasAttachmentLoader(atlas SpineGo.TextureAtlas) *atlasAttachmentLoader {
	a := new(atlasAttachmentLoader)
	a.atlas = atlas
	return a
}

func (a atlasAttachmentLoader) newRegionAttachment() RegionAttachment {
	panic("implement me")
}

func (a atlasAttachmentLoader) newMeshAttachment() MeshAttachment {
	panic("implement me")
}

func (a atlasAttachmentLoader) newBoundingBoxAttachment() BoundingBoxAttachment {
	panic("implement me")
}

func (a atlasAttachmentLoader) newClippingAttachment() ClippingAttachment {
	panic("implement me")
}

func (a atlasAttachmentLoader) newPathAttachment() PathAttachment {
	panic("implement me")
}

func (a atlasAttachmentLoader) newPointAttachment() PointAttachment {
	panic("implement me")
}
