# 解决冲突

如果有两笔有冲突的支易(Race Condition)，被两组节点以不同的顺序收到，根据 Bitcoin P2P 规则 ：

* 节点总是相信它首先收到的交易
* 冲突发生时，两个交易将处于 Pending 的状态，不写入区块链

要解决这个冲突，取决于两组节点谁先挖到下一个区块，然后谁接受的交易版本将写入区块链，另一个则被视为 Double Spend 而被抛弃掉
