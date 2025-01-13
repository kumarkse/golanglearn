# ###### remixToGanache.py 
# ## remix ide -> compiler env -> extrernal 127.0.0.1:7545
# -> create  a contract and deploy , it would be deopoyed on your ganache
# -> using py file , get ganache url, initialize web3 , set default account to first
# -> find ABI and address of deployed contract 
# -> initialize the contract
# -> perform transactions and verify inside ganache

from web3 import Web3
import json

ganache_url = 'http://127.0.0.1:7545'

web3 = Web3(Web3.HTTPProvider(ganache_url))

web3.eth.default_account=web3.eth.accounts[0]


abi  =  json.loads('[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"getval","outputs":[{"internalType":"int256","name":"","type":"int256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"int256","name":"value","type":"int256"}],"name":"setval","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"val","outputs":[{"internalType":"int256","name":"","type":"int256"}],"stateMutability":"view","type":"function"}]')


addr = web3.to_checksum_address('0x3c7b69b84131D7e7E741a294371F00299dcDFda0')

contract = web3.eth.contract(address=addr,abi=abi)
txn_hash = contract.functions.setval(5).transact()
# print(contract.functions.getval().call())

web3.eth.wait_for_transaction_receipt(txn_hash)

print("updates are : ", contract.functions.getval().call())

