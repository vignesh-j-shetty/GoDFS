package metadata

import (
	"encoding/gob"
	"io"
)

type GobSerializer struct{}

func NewGobSerializer() *GobSerializer {
	gob.Register(INode{})
	gob.Register(FileMetaData{})

	return &GobSerializer{}
}

func (g *GobSerializer) Encode(tree *INode, writer io.Writer) error {
	encoder := gob.NewEncoder(writer)
	return encoder.Encode(tree)
}

func (g *GobSerializer) Decode(tree *INode, reader io.Reader) error {
	decoder := gob.NewDecoder(reader)
	return decoder.Decode(tree)
}