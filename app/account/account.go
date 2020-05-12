package account

import  {
  "math/big"
   "hash"
}

struct Account  {
  Account string
  Balance big.Int
  TrustedNodes []string
  Key string 
}
