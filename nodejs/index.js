const crypto = require('crypto')
const blcokchain = {
    chain: []
}

const block = {
    timestamp: Date.now(),
    data: '{}',
    prevBlocKHash: '',
    hash: ''
}

function setHash (block) {
    const s = block.prevBlockHash + block.timestamp + block.data
    const hash = crypto.createHash('sha256');
    block.hash = hash.update(s).digest('hex');
}

function createGenesisBlock () {
    const newBlock = {
        timestamp: '2018-01-01',
        data: 'Genesis Block',
        prevBlockHash: '0',
    }
    setHash(newBlock)
    return newBlock
}

function newBlockChain () {
    blcokchain.chain.push(createGenesisBlock())
}

function getLatestBlock () {
    return blcokchain.chain[blcokchain.chain.length-1]
}

function addBlock (block) {
    block.prevBlockHash = getLatestBlock().hash
    setHash(block)
    blcokchain.chain.push(block)
}

function main () {
    newBlockChain()
    addBlock({
        timestamp: '2018-02-01',
        data: '{"amount":10}',
    })
    addBlock({
        timestamp: '2018-03-01',
        data: '{"amount":20}',
    })
    addBlock({
        timestamp: '2018-04-01',
        data: '{"amount":30}',
    })

    return blcokchain;
}

module.exports = main