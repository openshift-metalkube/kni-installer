package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift-metalkube/kni-installer/pkg/asset"
	"github.com/openshift-metalkube/kni-installer/pkg/asset/templates/content"
)

const (
	etcdNamespaceFileName = "etcd-namespace.yaml"
)

var _ asset.WritableAsset = (*EtcdNamespace)(nil)

// EtcdNamespace is an asset for the etcd namespace
type EtcdNamespace struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *EtcdNamespace) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *EtcdNamespace) Name() string {
	return "EtcdNamespace"
}

// Generate generates the actual files by this asset
func (t *EtcdNamespace) Generate(parents asset.Parents) error {
	fileName := etcdNamespaceFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *EtcdNamespace) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *EtcdNamespace) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, etcdNamespaceFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
