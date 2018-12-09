package packr

type Box struct {
	root string
}

func New(root string) *Box {
	return &Box{
		root: root,
	}
}

func (b *Box) Bytes(path string) []byte {
	return []byte(packData[b.root][path])
}

type walkFunk func(path string, data string) error

func (b *Box) Walk(f walkFunk) error {
	for path, data := range packData[b.root] {
		if err := f(path, data); err != nil {
			return err
		}
	}

	return nil
}
