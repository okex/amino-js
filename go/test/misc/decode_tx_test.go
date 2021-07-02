package misc

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/amino-js/go/lib/exchain/ethcmn"
	"github.com/cosmos/amino-js/go/src"
	"math/big"
	"strings"
	"testing"
)

func TestDecodeTx(t *testing.T) {
	str := "1gHwYl3uCkOoo2GaChS536x615s3L5HNHZqLKYPpCK3tiRIUF5FHHqSe5EQ+wA9u2QAy8QEasjAaEQoFdWF0b20SCDExNjU3OTk1EhMKDQoFdWF0b20SBDUwMDAQwJoMGmoKJuta6YchAtQaCqFnshaZQp6rIkvAPyzThvCvXSDO+9AzbxVErqJPEkDWdRwgfQItPT+dDSYFMOtPqQwbbQ1j8+wfs/aBzhulg0YsRiMGZ1Z69dQmi5IT/0D/rRAb1xh6rJN7mQUN4g/FIgoxMTIyNjcyNzU0"
	original, err := base64.StdEncoding.DecodeString(str)
	fmt.Println(original)
	originalDecoded, _ := src.DecodeTx(original, false)
	fmt.Println(originalDecoded)
	json := string(originalDecoded)
	fmt.Println(json)
	encoded, _ := src.EncodeTx(originalDecoded, false)

	encodedDecoded, _ := src.DecodeTx(encoded, false)

	json = string(originalDecoded)
	fmt.Println(json)
	_ = err
	_ = encoded
	_ = json
	_ = encodedDecoded
	return
}

func TestEncodeEthereumTx(t *testing.T) {
	type TmpTxData struct {
		AccountNonce uint64   `json:"nonce"`
		Price        *big.Int `json:"gasPrice"`
		GasLimit     uint64   `json:"gas"`
		Recipient    string   `json:"to" rlp:"nil"` // nil means contract creation
		Amount       *big.Int `json:"value"`
		Payload      []byte   `json:"input"`

		// signature values
		V *big.Int `json:"v"`
		R *big.Int `json:"r"`
		S *big.Int `json:"s"`

		// hash is only used when marshaling to JSON
		Hash *ethcmn.Hash `json:"hash" rlp:"-"`
	}
	amount, res := new(big.Int).SetString("1000000000000000000000", 10)
	if !res {
		panic("invalid r")
	}
	r, res := new(big.Int).SetString("96962823357929674581456147164326828477780712855340773268809156628353093709286", 10)
	if !res {
		panic("invalid r")
	}
	s, res := new(big.Int).SetString("4240042300301810108734890837836473848765664079489718075672737221057986070883", 10)
	if !res {
		panic("invalid r")
	}
	tx := TmpTxData{
		AccountNonce: 0,
		Price: big.NewInt(100000000),
		GasLimit: 21000,
		Recipient: "0xF76b10a1f318825173ad9F83f112e570782bD83E",
		Amount: amount,
		Payload: []byte(""),

		V: big.NewInt(51),
		R: r,
		S: s,
	}
	json_str := src.GetCdc().MustMarshalJSON(tx)
	fmt.Println(string(json_str))
	amino_data, err := src.EncodeEthereumTx(json_str, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(amino_data))
	fmt.Println(strings.ToUpper(hex.EncodeToString(amino_data)))
}