import sys
from bitcoin.rpc import RawProxy

def get_txs(block_height):
    p = RawProxy()
    blockhash = p.getblockhash(block_height)
    block = p.getblock(blockhash)
    transactions = block['tx']
    
    for txid in transactions:
        print txid

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print "no tx number."
        sys.exit(0)
   
    block_height = sys.argv[1]
    get_txs(block_height)