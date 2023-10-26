package block

import (
	"bytes"
	"encoding/gob"
)

type Block struct {
	TimeStamp     int64  // время создания блока
	Data          []byte //данные
	PrevBlockHash []byte //хэш предыдущего блока
	Hash          []byte //хэш данного блока
	Nonce         int
}

func NewBlock(data string, _PrevBlockHash []byte) (b *Block) {
	block := &Block{Data: []byte(data), PrevBlockHash: _PrevBlockHash, Hash: []byte{}}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		print("ошибка Encode ")
		panic("ошибка Encode ")
	}
	return result.Bytes()
}

func DeserializeBlock(b []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(b))

	err := decoder.Decode(&block)
	if err != nil {
		print("ошибка Decode ")
		panic("ошибка Decode ")
	}
	return &block
}
