Particles is a multisignature accelerator p2p network
-----------------------------------------------------

The goal of `particles` is to allow all peers to locally determine
which transactions they should sign, based on configuration or
external events, then broadcast those signatures on the p2p network.

When transactions attain a certain threshold of signatures, thus if
enough peers agree on what should be signed, the full transactions
along with all required signatures is broadcast to the target
blockchain network.

`particles` is currently an EOSIO network sidecar: it crafts and signs
EOSIO transactions, and broadcasts them to EOSIO networks.
