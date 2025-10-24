package metadata

import "io"

type INodeTreeSerializer interface {
	Encode(tree *INode, writer io.Writer) error

	Decode(tree *INode, reader io.Reader) error
}