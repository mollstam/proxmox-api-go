package cli_storage_test

import (
	"testing"

	_ "github.com/mollstam/proxmox-api-go/cli/command/commands"
	"github.com/mollstam/proxmox-api-go/internal/util"
	"github.com/mollstam/proxmox-api-go/proxmox"
	storagesubtests "github.com/mollstam/proxmox-api-go/test/cli/Storage/storage-sub-tests"
)

func Test_Storage_ZFSoverISCSI_0_Cleanup(t *testing.T) {
	storagesubtests.Cleanup("zfs-over-iscsi-test-0", t)
}

func Test_Storage_ZFSoverISCSI_0_Create_Full(t *testing.T) {
	s := storagesubtests.CloneJson(storagesubtests.ZFSoverISCSIFull)
	storagesubtests.Create(s, "zfs-over-iscsi-test-0", t)
}

func Test_Storage_ZFSoverISCSI_0_Get_Full(t *testing.T) {
	s := storagesubtests.CloneJson(storagesubtests.ZFSoverISCSIFull)
	s.ID = "zfs-over-iscsi-test-0"
	s.ZFSoverISCSI.Blocksize = util.Pointer("8k")
	storagesubtests.Get(s, s.ID, t)
}

func Test_Storage_ZFSoverISCSI_0_Update_Empty(t *testing.T) {
	s := storagesubtests.CloneJson(storagesubtests.ZFSoverISCSIEmpty)
	storagesubtests.Update(s, "zfs-over-iscsi-test-0", t)
}

func Test_Storage_ZFSoverISCSI_0_Get_Empty(t *testing.T) {
	s := storagesubtests.CloneJson(storagesubtests.ZFSoverISCSIEmpty)
	s.ID = "zfs-over-iscsi-test-0"
	s.ZFSoverISCSI.Blocksize = util.Pointer("8k")
	s.Content = &proxmox.ConfigStorageContent{
		DiskImage: util.Pointer(true),
	}
	storagesubtests.Get(s, s.ID, t)
}

func Test_Storage_ZFSoverISCSI_0_Delete(t *testing.T) {
	storagesubtests.Delete("zfs-over-iscsi-test-0", t)
}
