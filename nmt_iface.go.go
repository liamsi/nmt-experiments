package nmt_experiments

type Nmt1 interface {
	// Option 1) NID has to be part of data, or, the implementation
	// needs to know how to compute NID (from data)
	Push(data []byte)
	Root() []byte
}

type Nmt2 interface {
	// Option 2) data does not need to contain NID:
	Push(namespaceID []byte, data []byte)
	Root() []byte
	// TODO: Consider this too:
	//Root() (minNs, maxNS, root []byte)
}
