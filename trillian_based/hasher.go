package trillian_based

import (
	"bytes"
	"crypto"

	"github.com/google/trillian/merkle/hashers"
)

var _ hashers.LogHasher = &NmtHasher{}

// type NidFunc = func(m []byte) []byte

const (
	LeafPrefix = 0
	NodePrefix = 1
)

type NmtHasher struct {
	baseHasher        crypto.Hash
	NamespaceIDLength int
	flagLen           int
}

func NewNmtHasher(NIDLen int) *NmtHasher {
	return &NmtHasher{
		baseHasher:        crypto.SHA256,
		NamespaceIDLength: NIDLen,
		flagLen:           2 * NIDLen,
	}
}

func (n NmtHasher) EmptyRoot() []byte {
	h := n.baseHasher.New()
	return h.Sum(nil)
}

// CONTRACT: leaf contains it's namespace:
// len(leaf) > 2*NamespaceIDLength
// NID = leaf[:NamespaceIDLength]
func (n NmtHasher) HashLeaf(leaf []byte) []byte {
	h := n.baseHasher.New()

	nid := n.extractNIDFromLeaf(leaf)
	res := append(append(make([]byte, 0), nid...), nid...)
	h.Write([]byte{LeafPrefix})
	h.Write(leaf) // XXX: do we want to remove the extra nid bytes from the leaf?
	// res = nsid||nsid||hash(leafPrefix||leaf), where leaf = nid || raw_data
	return h.Sum(res)
}

func (n NmtHasher) HashChildren(l, r []byte) []byte {
	leftMinNs, leftMaxNs := n.extractMinMaxNsFromNode(l)
	rightMinNs, rightMaxNs := n.extractMinMaxNsFromNode(r)
	minNs := min(leftMinNs, rightMinNs)
	maxNs := max(leftMaxNs, rightMaxNs)

	res := append(append(make([]byte, 0), minNs...), maxNs...)
	h := n.baseHasher.New()
	h.Write([]byte{NodePrefix})
	h.Write(l)
	h.Write(r)
	// res = minNs||maxNs||sha256(NodePrefix||left||right)
	return h.Sum(res)
	// TODO: consider combining vals and calling write once instead:
	// https://github.com/google/trillian/pull/1503
}

func max(ns []byte, ns2 []byte) []byte {
	if bytes.Compare(ns, ns2) >= 0 {
		return ns
	}
	return ns2
}

func min(ns []byte, ns2 []byte) []byte {
	if bytes.Compare(ns, ns2) <= 0 {
		return ns
	}
	return ns2
}

func (n NmtHasher) Size() int {
	return n.baseHasher.Size() + n.flagLen
}

func (n NmtHasher) extractNIDFromLeaf(leaf []byte) []byte {
	return leaf[:n.NamespaceIDLength]
}

// extract minNs, maxNs from node
func (n NmtHasher) extractMinMaxNsFromNode(node []byte) ([]byte, []byte) {
	return node[:n.NamespaceIDLength], node[n.NamespaceIDLength:n.flagLen]
}
