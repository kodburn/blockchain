package blockchain

import (
	"log"

	"github.com/boltdb/bolt"
)

// BlockchainIterator iterates over blocks in the blockchain.
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Next returns next block starting from the tip.
func (i *BlockchainIterator) Next() (block *Block, last bool) {
	block = &Block{}
	last = false

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)
		if block.PrevBlockHash == nil {
			last = true
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return
}
