package model

import "github.com/sedraxnet/sedraxd/domain/consensus/model/externalapi"

// Multiset represents a secp256k1 multiset
type Multiset interface {
	Add(data []byte)
	Remove(data []byte)
	Hash() *externalapi.DomainHash
	Serialize() []byte
	Clone() Multiset
}
