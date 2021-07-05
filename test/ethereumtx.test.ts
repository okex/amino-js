import * as Amino from '../';


const ethereumTx = {
    nonce: 0,
    gasPrice: "100000000",
    gas: 21000,
    to: "0xF76b10a1f318825173ad9F83f112e570782bD83E",
    value: "1000000000000000000000",
    input: new Uint8Array(),

    v: "51",
    r: "96962823357929674581456147164326828477780712855340773268809156628353093709286",
    s: "4240042300301810108734890837836473848765664079489718075672737221057986070883",
}

describe('Ethereum Tx', () => {
    describe('encoding', () => {
            it('encode bytes', () => {
                const bytes = Amino.marshalEthereumTx(ethereumTx);
                console.log(bytes)
            });
    });
});