### Steps ###

### Step1: Clone the wasm github repo 


git clone https://github.com/CosmWasm/wasmd

cd /wasmd


make install 

git clone https://github.com/cosmwasm/cw-examples


cd /


rustup target add wasm32-unknown-unknown

RUSTFLAGS='-C link-arg=-s' cargo wasm


cd /home/


./init_script

apt install vim -y

vim /root/.wasmd/config/app.toml



TXFLAG="--chain-id $CHAIN_ID --gas-prices 20stake --gas auto --gas-adjustment 1.3"

wasmd tx wasm store target/wasm32-unknown-unknown/release/cw_nameservice.wasm --from main  $TXFLAG -y -b block



INIT='{"purchase_price":{"amount":"100","denom":"stake"},"transfer_price":{"amount":"999","denom":"stake"}}'

wasmd tx wasm instantiate 1 "$INIT" --from main - --label "awesome name service" --no-admin $TXFLAG

CONTRACT=$(wasmd query wasm list-contract-by-code $CODE_ID --output json | jq -r '.contracts[-1]')
wasmd query wasm contract $CONTRACT

# check contract state

### Step2: Build the wasm binary using the Docker env 


### Step3: Clone the Cosmwasm conttract examples


### Step4: Build the smart contract 


### Step5: Deploy the smart contract to the wasm chain 







