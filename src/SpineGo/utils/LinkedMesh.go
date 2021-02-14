package utils

import . "SpineGo/attachments"

type linkedMesh struct {
	parent        string
	skin          string
	slotIndex     int
	mesh          MeshAttachment
	inheritDeform bool
}

func NewLinkedMesh(mesh MeshAttachment, skin string, slotIndex int, parent string, inheritDeform bool) {
	m := new(linkedMesh)
	m.mesh = mesh
	m.skin = skin
	m.slotIndex = slotIndex
	m.parent = parent
	m.inheritDeform = inheritDeform
}
