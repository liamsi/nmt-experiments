# Namespaced Merkle tree experiments

This repo contains two simple name spaced merkle tree implementations. 

The first one is shamelessly copied from [@musalbas] LazyLedger [prototype] and can be found under [./lazyledger_prototype].
It is based on [NebulousLabs'] merkle tree implementation. It povides a particular [`hash.Hash`] [implementation](https://github.com/liamsi/nmt-experiments/blob/a093f5c6c2106a14cef0f596f42e151922e85538/lazyledger_prototype/flaghasher.go#L15) to the tree that employs a [`Flagger`] to prefix the cryptographic hash function with the min/max- namespace ID as described in the LazyLedger [academic paper].

The second implementation uses [trillian]. This implementation aims to produce the same trees as the first one. Different than the first, it defines a [`LogHasher`] ([here]) and reuses a trillian (in-memory) tree implementation. Here, the `LogHasher` to prepends the crpytographic hash function with the namespace ID bytes.


This repository is just a playground to explore which abstractions are the right ones for a namespaced merkle tree. 

[@musalbas]: https://github.com/musalbas
[prototype]: https://github.com/lazyledger/lazyledger-prototype
[./lazyledger_prototype]: https://github.com/liamsi/nmt-experiments/tree/master/lazyledger_prototype
[NebulousLabs']: https://gitlab.com/NebulousLabs/merkletree
[`hash.Hash`]: https://golang.org/pkg/hash/#Hash
[`Flagger`]: https://github.com/liamsi/nmt-experiments/blob/a093f5c6c2106a14cef0f596f42e151922e85538/lazyledger_prototype/flagger.go#L4-L16
[academic paper]: https://arxiv.org/abs/1905.09274
[trillian]: https://github.com/google/trillian
[`LogHasher`]: https://github.com/google/trillian/blob/7502e99bb92ecf0ec8add958889c751f2cfc7f59/merkle/hashers/tree_hasher.go#L23-L34
[here]: https://github.com/liamsi/nmt-experiments/blob/a093f5c6c2106a14cef0f596f42e151922e85538/trillian_based/hasher.go#L10-L67
