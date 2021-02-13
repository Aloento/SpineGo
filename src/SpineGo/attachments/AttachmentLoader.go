package attachments

type AttachmentLoader interface {
	newRegionAttachment() RegionAttachment
	newMeshAttachment() MeshAttachment
	newBoundingBoxAttachment() BoundingBoxAttachment
	newClippingAttachment() ClippingAttachment
	newPathAttachment() PathAttachment
	newPointAttachment() PointAttachment
}
