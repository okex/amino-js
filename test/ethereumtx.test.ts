import './setup';
import * as Amino from '../';
import SHA256 from "crypto-js/sha256";
import hexEncoding from "crypto-js/enc-hex"

const ethereumTx = {
    nonce: "0",
    gasPrice: "100000000",
    gas: "21000",
    to: "0xF76b10a1f318825173ad9F83f112e570782bD83E",
    value: "1000000000000000000000",
    input: "",

    v: "51",
    r: "96962823357929674581456147164326828477780712855340773268809156628353093709286",
    s: "4240042300301810108734890837836473848765664079489718075672737221057986070883",
}

describe('Ethereum Tx', () => {
    describe('encoding', () => {
            it('encode bytes', () => {
                const bytes = Amino.marshalEthereumTx(ethereumTx);
                const hexBytes = Buffer.from(bytes).toString("hex").toUpperCase()
                console.log(hexBytes)
                expect(hexBytes).toBe("E50125A6BE540ADE0112093130303030303030301888A4012214F76B10A1F318825173AD9F83F112E570782BD83E2A16313030303030303030303030303030303030303030303A023531424D39363936323832333335373932393637343538313435363134373136343332363832383437373738303731323835353334303737333236383830393135363632383335333039333730393238364A4C34323430303432333030333031383130313038373334383930383337383336343733383438373635363634303739343839373138303735363732373337323231303537393836303730383833")
                const hash = "0x" + SHA256(hexEncoding.parse(hexBytes)).toString()
                expect(hash).toBe("0xbe648799632a6db88b81fdd48830087f2695f6190c63771e732f29a565b82b56")
            });
    });
});