make all

rm -rf ~/.comdex

mkdir ~/.comdex

VAL_MNEMONIC="copper push brief egg scan entry inform record adjust fossil boss egg comic alien upon aspect dry avoid interest fury window hint race symptom"
RELAYER_MNEMONIC="pony glide frown crisp unfold lawn cup loan trial govern usual matrix theory wash fresh address pioneer between meadow visa buffalo keep gallery swear"


comdex init --chain-id test test
echo $VAL_MNEMONIC | comdex keys add test --recover --keyring-backend=test
echo $RELAYER_MNEMONIC | comdex keys add relayer --recover --keyring-backend=test

comdex add-genesis-account test 100000000000000stake --keyring-backend test
comdex add-genesis-account relayer 100000000000000stake --keyring-backend test

comdex gentx test 1000000000stake --chain-id test --keyring-backend test
comdex collect-gentxs
comdex start --grpc.address="0.0.0.0:9090" --p2p.laddr "tcp://0.0.0.0:26659" --rpc.laddr "tcp://127.0.0.1:26658" --grpc-web.enable=false --log_level trace > ./comdex.log 2>&1 &