from web3 import Web3

ganache_url = 'http://127.0.0.1:7545'
web3 = Web3(Web3.HTTPProvider(ganache_url))


acc1 = '0x30E715243397aEE386e2AbCDc48a6dE73FeB8F52'
acc2 = '0xa396CC3bdBB5EcbAf6db00df7Af5B39e2937Df8D'

private_key_acc1 = '0x69f272f372e9a4f5c53af6645d33651ad9deb927486141d8c60836bb12e6ae08'

# """
# Steps:
#     1) get the nonce
#     2) build the transaction
#     3) sign the transaction
#     4) send transaction
#     5) get transaction hash
# """

# #1)
nonce = web3.eth.get_transaction_count(acc1)
tx = {
    'nonce':nonce,
    'to':acc2,
    'value':web3.to_wei(5,'ether'),
    'gas':2000000,
    'gasPrice':web3.to_wei(50,'gwei')
}

signed_tx = web3.eth.account.sign_transaction(tx,private_key_acc1)

tx_hash = web3.eth.send_raw_transaction(signed_tx.raw_transaction)

print(tx_hash)