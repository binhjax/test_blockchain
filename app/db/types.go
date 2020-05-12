package account

import  {
  "math/big"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
}

// Lengths of hashes and addresses in bytes.
const (
	// HashLength is the expected length of the hash
	HashLength = 32
	// AddressLength is the expected length of the address
	AddressLength = 20
)

struct Account  {

  Account []byte
  Balance big.Int
  TrustedNodes []string
  Key []byte
}


struct Transaction  {
  Nonce big.Int
  ParentHash  Hash
  Hash Hash
  Account []byte
  TxID big.Int
  TxType int
  TxCode []byte
  TxSigned []byte
  PartnerHash []Hash
  ConsensusHash []Hash
}
