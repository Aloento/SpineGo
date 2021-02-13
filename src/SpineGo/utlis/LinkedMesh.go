package utlis

import . "SpineGo/attachments"

type LinkedMesh struct {
	parent        string
	skin          string
	slotIndex     int
	mesh          MeshAttachment
	inheritDeform bool
}

func (m *LinkedMesh) LinkedMesh(mesh MeshAttachment, skin string, slotIndex int, parent string, inheritDeform bool) {
	m.mesh = mesh
	m.skin = skin
	m.slotIndex = slotIndex
	m.parent = parent
	m.inheritDeform = inheritDeform
}
