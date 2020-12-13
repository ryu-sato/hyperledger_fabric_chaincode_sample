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
rm -rf fabric/crypto-config
cd fabric
cryptogen generate --config crypto-config.yaml --output crypto-config
cd ..

# Genesisブロックを作成する
rm genesis.block
configtxgen -configPath fabric -channelID mychannel -profile Genesis -outputBlock genesis.block

# channelブロックを作成する
rm mychannel.tx
configtxgen -configPath fabric -channelID mychannel -profile Channel --outputCreateChannelTx mychannel.tx
```
