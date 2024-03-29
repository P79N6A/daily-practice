import json
from bitcoin.rpc import RawProxy

p = RawProxy()

blockheight = 234349

blockhash = p.getblockhash(blockheight)

block = p.getblock(blockhash)

transactions = block['tx']

block_value = 0

for txid in transactions:
    tx_value = 0
    raw_tx = p.getrawtransaction(txid)
    decoded_tx = p.decoderawtransaction(raw_tx)

    for output in decoded_tx['vout']:
        tx_value = tx_value + output['value']

    block_value = block_value + tx_value

print("total value in block[{0}]: {1}".format(blockheight, block_value))