# Hyperledger fabric を開発モードで実行する

```bash
cd misc
docker-compose up -d
```

# genesis.block 更新方法

* 証明書の期限が切れた場合
* channel 名を変えたい場合

```bash
# mspを作成する
rm -rf msp fabric/crypto-config
cd fabric
cryptogen generate --config=./crypto-config.yaml --output crypto-config
cp -r crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp ../
cd ..

# Genesisブロックを作成する
configtxgen -configPath fabric -channelID mychannel -profile OrdererGenesis -outputBlock genesis.block

# channelブロックを作成する
configtxgen -configPath fabric -profile Channel --outputCreateChannelTx mychannel.tx
```
