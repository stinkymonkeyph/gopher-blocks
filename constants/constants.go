package constants

const (
	BLOCKCHAIN_NAME           = "GopherBlocks"
	HEX_PREFIX                = "0x"
	STATUS_SUCCESS            = "success"
	STATUS_FAILED             = "failed"
	STATUS_PENDING            = "pending"
	MINING_DIFFICULTY         = 3
	MINING_REWARD             = 1200 * DECIMAL
	CURRENCY_NAME             = "gopher"
	DECIMAL                   = 100
	BLOCKCHAIN_REWARD_ADDRESS = "Gopher_Reward_Pool"
	DB_PATH                   = "./data"
	DB_KEY                    = "blockchain_data"
	//temporary airdrop distribution for sender wallet since tx will fail because balance is 0 for sender wallets
	AIRDROP_AMOUNT             = 100
	AIRDROP_ROUND              = 1
	BLOCKCHAIN_AIRDROP_ADDRESS = "Gopher_Airdrop_Pool"
)
