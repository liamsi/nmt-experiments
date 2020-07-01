package trillian_based

import (
	"github.com/google/trillian/merkle"
	nmt_experiments "github.com/liamsi/nmt-experiments"
)

var _ nmt_experiments.Nmt1 = &Nmt{}

type Nmt struct {
	*merkle.InMemoryMerkleTree
}

func New() *Nmt {
	treeHasher := NewNmtHasher(8)
	return &Nmt{merkle.NewInMemoryMerkleTree(treeHasher)}
}

func (n Nmt) Push(data []byte) {
	n.AddLeaf(data)
}

func (n Nmt) Root() []byte {
	return n.CurrentRoot().Hash()
}
