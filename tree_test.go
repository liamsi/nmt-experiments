package nmt_experiments_test

import (
	"crypto"
	"math/rand"
	"testing"

	"github.com/google/trillian/merkle"
	"github.com/google/trillian/merkle/rfc6962"
	"github.com/stretchr/testify/assert"
	"gitlab.com/NebulousLabs/merkletree"

	nmt_experiments "github.com/liamsi/nmt-experiments"
	"github.com/liamsi/nmt-experiments/lazyledger_prototype"
	"github.com/liamsi/nmt-experiments/trillian_based"
)

var (
	nsID1 = []byte("first_ns")
	nsID2 = []byte("secondns")
	nsID3 = []byte("third_ns")

	leafs = make([][]byte, 6)
)

func init() {
	leafs[0] = append(nsID1, "leaf1"...)
	leafs[1] = append(nsID1, "leaf2"...)

	leafs[2] = append(nsID2, "leaf3"...)
	leafs[3] = append(nsID2, "leaf4"...)
	leafs[4] = append(nsID2, "leaf5"...)

	leafs[5] = append(nsID3, "leaf6"...)
}

func TestTrees(t *testing.T) {
	protoTree := lazyledger_prototype.New()
	trillianTree := trillian_based.New()
	for _, leaf := range leafs {
		protoTree.Push(leaf)
		trillianTree.Push(leaf)
	}

	r1 := protoTree.Root()
	r2 := trillianTree.Root()
	assert.Equal(t, r1, r2)
}

func TestVanillaTrees(t *testing.T) {
	t.Skip("just to ensure that Nebulous' tree actually matches rfc6962")

	nebTree := merkletree.New(crypto.SHA256.New())
	logTree := merkle.NewInMemoryMerkleTree(rfc6962.DefaultHasher)
	for _, leaf := range leafs {
		nebTree.Push(leaf)
		logTree.AddLeaf(leaf)
	}
	r1 := nebTree.Root()
	r2 := logTree.CurrentRoot().Hash()
	assert.Equal(t, r1, r2)
}
func BenchmarkPushLL(b *testing.B) { benchmarkPushData(b, lazyledger_prototype.New()) }
func BenchmarkRootLL(b *testing.B) { benchmarkComputeRoot(b, lazyledger_prototype.New()) }
func BenchmarkPushTrill(b *testing.B) { benchmarkPushData(b, trillian_based.New()) }
func BenchmarkRootTrill(b *testing.B) { benchmarkComputeRoot(b, trillian_based.New()) }

func benchmarkPushData(b *testing.B, nmt nmt_experiments.Nmt1) {
	leafCnt := 1
	leafSize := 64
	randLeafs := randLeafData(leafCnt, leafSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nmt.Push(randLeafs[0])
	}
}

func benchmarkComputeRoot(b *testing.B, nmt nmt_experiments.Nmt1) {
	leafCnt := 1000
	leafSize := 64
	randLeafs := randLeafData(leafCnt, leafSize)
	for _, leaf := range randLeafs {
		nmt.Push(leaf)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nmt.Push(append(randLeafs[0], 1))
		nmt.Root()
	}
}

func randLeafData(leafCnt int, leafSize int) [][]byte {
	randLeafs := make([][]byte, leafCnt)
	for i := 0; i < leafCnt; i++ {
		randLeafs[i] = make([]byte, leafSize)
		rand.Read(randLeafs[i])
	}
	return randLeafs
}