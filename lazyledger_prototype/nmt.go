package lazyledger_prototype

import (
	"crypto/sha256"

	"gitlab.com/NebulousLabs/merkletree"
)

type Nmt struct {
	flagHasher *flagDigest
	tree *merkletree.Tree
}

func New() *Nmt {
	// These could be replaced with other hashers / flaggers:
	ndf := NewNamespaceDummyFlagger()
	// the flagHasher has some of the tree logic baked in:
	// - it flags depending on if we are hashing leafs or inner nodes
	fh := NewFlagHasher(ndf, sha256.New())
	flDigest := fh.(*flagDigest)

	tree:= merkletree.New(fh)
	return &Nmt{flagHasher: flDigest, tree: tree}
}

func (nmt *Nmt) Push(data []byte) {
	nmt.tree.Push(data)
}

func (nmt *Nmt) Root() []byte  {
	return nmt.tree.Root()
}