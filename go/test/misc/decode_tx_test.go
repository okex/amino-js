package misc

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/amino-js/go/lib/exchain/ethcmn"
	"github.com/cosmos/amino-js/go/src"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/tmhash"
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
		AccountNonce string   `json:"nonce"`
		Price        string `json:"gasPrice"`
		GasLimit     string   `json:"gas"`
		Recipient    string   `json:"to" rlp:"nil"` // nil means contract creation
		Amount       string `json:"value"`
		Payload      string   `json:"input"`

		// signature values
		V string `json:"v"`
		R string `json:"r"`
		S string `json:"s"`

		// hash is only used when marshaling to JSON
		Hash *ethcmn.Hash `json:"hash" rlp:"-"`
	}

	tx := TmpTxData{
		AccountNonce: "0",
		Price: "100000000",
		GasLimit: "21000",
		Recipient: "0xF76b10a1f318825173ad9F83f112e570782bD83E",
		Amount: "1000000000000000000000",
		Payload: hex.EncodeToString([]byte("")),

		V: "51",
		R: "96962823357929674581456147164326828477780712855340773268809156628353093709286",
		S: "4240042300301810108734890837836473848765664079489718075672737221057986070883",
	}
	json_str := src.GetCdc().MustMarshalJSON(tx)
	fmt.Println(string(json_str))
	amino_data, err := src.EncodeEthereumTx(json_str, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(amino_data))
	hexRes := strings.ToUpper(hex.EncodeToString(amino_data))
	fmt.Println(hexRes)
	require.Equal(t, "E50125A6BE540ADE0112093130303030303030301888A4012214F76B10A1F318825173AD9F83F112E570782BD83E2A16313030303030303030303030303030303030303030303A023531424D39363936323832333335373932393637343538313435363134373136343332363832383437373738303731323835353334303737333236383830393135363632383335333039333730393238364A4C34323430303432333030333031383130313038373334383930383337383336343733383438373635363634303739343839373138303735363732373337323231303537393836303730383833", hexRes)
	hash := "0x" + hex.EncodeToString(tmhash.Sum(amino_data))
	fmt.Println(hash)
	require.Equal(t, "0xbe648799632a6db88b81fdd48830087f2695f6190c63771e732f29a565b82b56", hash)
}