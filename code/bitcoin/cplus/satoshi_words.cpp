#include <iostream>
#include <bitcoin/bitcoin.hpp>

int main() {
    bc::block_type block = bc::genesis_block();
    
    const bc::transaction_type& coinbase_tx = block.transactions[0];

    assert(coinbase_tx.input_size() == 1);
    const bc::transaction_type& coinbase_input = coinbase_tx.inputs[0];

    const bc::data_chunk& raw_message = save_script(coinbase_input.script);

    std::string message;
    message.resize(raw_message.size());
    std::copy(raw_message.begin(), raw_message.end(), message.begin());

    std::cout << message <<std::endl;
    return 0;
}