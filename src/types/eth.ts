

export interface EthereumTx {
    nonce: string;
    gasPrice: string;
    gas: string;
    to: string;
    value: string;
    input: string;

    v: string;
    r: string;
    s: string;
}