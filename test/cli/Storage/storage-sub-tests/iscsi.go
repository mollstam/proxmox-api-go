package storagesubtests

import (
	"testing"

	"github.com/mollstam/proxmox-api-go/internal/util"
	"github.com/mollstam/proxmox-api-go/proxmox"
)

var IscsiFull = proxmox.ConfigStorage{
	Enable: true,
	Nodes:  []string{"pve"},
	Type:   "iscsi",
	ISCSI: &proxmox.ConfigStorageISCSI{
		Portal: "10.20.1.1",
		Target: "target-volume",
	},
	Content: &proxmox.ConfigStorageContent{
		DiskImage: util.Pointer(true),
	},
}

var IscsiEmpty = proxmox.ConfigStorage{
	Type: "iscsi",
	ISCSI: &proxmox.ConfigStorageISCSI{
		Portal: "10.20.1.1",
		Target: "target-volume",
	},
	Content: &proxmox.ConfigStorageContent{},
}

func IscsiGetFull(name string, t *testing.T) {
	s := CloneJson(IscsiFull)
	s.ID = name
	Get(s, name, t)
}

func IscsiGetEmpty(name string, t *testing.T) {
	s := CloneJson(IscsiEmpty)
	s.ID = name
	s.Content.DiskImage = util.Pointer(false)
	Get(s, name, t)
}
