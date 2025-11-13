package metadata

import (
	"testing"

	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/config"
)

func TestMetatHandlerGetINodeFromPath1(t *testing.T) {
	cnf := config.Config{}
	handler := NewFileMetaDataHandler(cnf)

	handler.CreateFolder("/", "testfolder")

	folder, err := handler.GetINodeFromPath("/testfolder")
	if err != nil {
		t.Fatal("Expected err nil")
	}

	if !folder.IsDir {
		t.Fatal("Expected to be dir")
	}

	if folder.Name != "testfolder" {
		t.Fatal("Expected name to be testfolder")
	}
}

func TestMetatHandlerGetINodeFromPath2(t *testing.T) {
	cnf := config.Config{}
	handler := NewFileMetaDataHandler(cnf)

	handler.CreateFolder("/", "testfolder")

	folder, err := handler.GetINodeFromPath("/testfolder")
	folder.CreateFolder("testfolder1")
	folder.CreateFolder("testfolder2")

	if err != nil {
		t.Fatal("Expected err nil")
	}

	if !folder.IsDir {
		t.Fatal("Expected to be dir")
	}

	if folder.Name != "testfolder" {
		t.Fatal("Expected name to be testfolder")
	}
}