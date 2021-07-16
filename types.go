package fireblockssdkgo

type VaultAccountResponse struct {
	id            string
	name          string
	hiddenOnUI    bool
	assets        AssetResponse // change to array
	customerRefId string        // optional
	autoFuel      bool
}

type AssetResponse struct {
	id    string
	total string
	/**
	 * @deprecated Replaced by "total"
	 */
	balance              string // optional
	lockedAmount         string // optional
	available            string // optional
	pending              string // optional
	selfStakedCPU        string // optional
	selfStakedNetwork    string // optional
	pendingRefundCPU     string // optional
	pendingRefundNetwork string // optional
	totalStakedCPU       string // optional
	totalStakedNetwork   string // optional
}

type UnfreezeTransactionResponse struct {
	success bool
}

type CreateVaultAssetResponse struct {
	id             string
	address        string
	legacyAddress  string
	tag            string
	eosAccountName string
}

//      type WalletAssetResponse extends AssetResponse interface{
// 	status string
// 	address string
// 	tag string
//     }

type WalletContainerResponse struct {
	id   string
	name string
	// assets WalletAssetResponse[]
	customerRefID string // optional
}

type CreateTransactionResponse struct {
	id     string
	status string
}

type EstimateFeeResponse struct {
	low    EstimatedFee
	medium EstimatedFee
	high   EstimatedFee
}

type EstimateTransactionFeeResponse struct {
	low    EstimatedTransactionFee
	medium EstimatedTransactionFee
	high   EstimatedTransactionFee
}

type EstimatedFee struct {
	networkFee string
	gasPrice   string
	feePerByte string // optional
}

type EstimatedTransactionFee struct {
	networkFee string // optional
	gasPrice   string // optional
	gasLimit   string // optional
	feePerByte string // optional
}

type PeerType string

const (
	EXCHANGE_ACCOUNT   PeerType = "EXCHANGE_ACCOUNT"
	INTERNAL_WALLET    PeerType = "INTERNAL_WALLET"
	EXTERNAL_WALLET    PeerType = "EXTERNAL_WALLET"
	UNKNOWN            PeerType = "UNKNOWN"
	NETWORK_CONNECTION PeerType = "NETWORK_CONNECTION"
	FIAT_ACCOUNT       PeerType = "FIAT_ACCOUNT"
	COMPOUND           PeerType = "COMPOUND"
	ONE_TIME_ADDRESS   PeerType = "ONE_TIME_ADDRESS"
)

type TransferPeerPath struct {
	peerType    PeerType // optional
	id          string
	virtualId   string      // optional
	virtualType VirtualType // optional
	address     string      // optional
}

type DestinationTransferPeerPath struct {
	peerType       PeerType
	id             string          // optional
	virtualID      string          // optional
	virtualType    VirtualType     // optional
	oneTimeAddress IOneTimeAddress // optional
}

type IOneTimeAddress struct {
	address string
	tag     string // optional
}

type DepositAddressResponse struct {
	assetID       string
	address       string
	tag           string // optional
	description   string // optional
	Type          string // optional
	customerRefID string // optional
	addressFormat string // optional
	legacyAddress string // optional
}

type GenerateAddressResponse struct {
	address       string // optional
	tag           string // optional
	legacyAddress string // optional
}

type SigningAlgorithm string

const (
	MPC_ECDSA_SECP256K1 SigningAlgorithm = "MPC_ECDSA_SECP256K1"
	MPC_ECDSA_SECP256R1 SigningAlgorithm = "MPC_ECDSA_SECP256R1"
	MPC_EDDSA_ED25519   SigningAlgorithm = "MPC_EDDSA_ED25519"
)

type RawMessageData struct {
	messages  RawMessage       // TODO: array
	algorithm SigningAlgorithm // optional
}

type RawMessage struct {
	content           string
	bip44addressIndex int64 // optional
	bip44change       int64 // optional
	derivationPath    int64 // TODO array
}

type TransactionDestination struct {
	amount      string
	destination DestinationTransferPeerPath
}

type TransactionArguments struct {
	assetID         string           // optional
	source          TransferPeerPath // optional
	destination     DestinationTransferPeerPath
	amount          string
	operation       TransactionOperation // optional
	fee             string
	feeLevel        FeeLevel // optional
	failOnLowFee    bool     // optional
	maxFee          string   // optional
	gasPrice        string   // optional
	gasLimit        string   // optional
	note            string   // optional
	cpuStaking      int64
	networkStaking  int64
	autoStaking     bool
	customerRefID   string                 // optional
	extraParameters string                 // TODO: object
	destinations    TransactionDestination // optional // TODO: array
	replaceTxByHash string
}

type FeeLevel string

const (
	HIGH   FeeLevel = "HIGH"
	MEDIUM FeeLevel = "MEDIUM"
	LOW    FeeLevel = "LOW"
)

type ExchangeResponse struct {
	id           string
	Type         string
	name         string
	assets       []AssetResponse
	isSubaccount bool
	status       string
}

type FiatAccountResponse struct {
	id      string
	Type    string
	name    string
	address string
	assets  AssetResponse
}

type Source struct {
	id       string
	peerType PeerType
	name     string
	subType  string
}

type Destination struct {
	id       string
	peerType PeerType
	name     string
	subType  string
}

type AmlScreeningResult struct {
	provider        string
	payload         interface{}
	screeningStatus string
	bypassReason    string
	timestamp       int64
}

type TransactionResponse struct {
	id          string
	assetID     string
	source      Source
	destination Destination
	amount      int64
	/**
	 * @deprecated Replaced by "networkFee"
	 */
	fee                           int64
	networkFee                    int64
	amountUSD                     int64
	netAmount                     int64
	createdAt                     int64
	lastUpdated                   int64
	status                        TransactionStatus
	txHash                        string
	numOfConfirmations            int64
	subStatus                     string
	signedBy                      []string
	createdBy                     string
	rejectedBy                    string
	destinationAddress            string
	destinationAddressDescription string
	destinationTag                string
	addressType                   string
	note                          string
	exchangeTxID                  string
	requestedAmount               int64
	serviceFee                    int64
	feeCurrency                   string
	amlScreeningResult            AmlScreeningResult
	signedMessages                SignedMessageResponse // TODO array
}

type Signature struct {
	fullSig string
	r       string
	s       string
	v       int64
}

type SignedMessageResponse struct {
	content        string
	algorithm      string
	derivationPath string
	signature      Signature
	publicKey      string
}

type CancelTransactionResponse struct {
	success bool
}

type OperationSuccessResponse struct {
	success bool
}

type LocalChannel struct {
	networkID string
	name      string
}

type RemoteChannel struct {
	networkID string
	name      string
}

type NetworkConnectionResponse struct {
	id            string
	localChannel  LocalChannel
	remoteChannel RemoteChannel
}

type TransactionFilter struct {
	before     int64
	after      int64
	status     TransactionStatus
	orderBy    TransactionOrder
	limit      int64
	txHash     string
	assets     string
	sourceType PeerType
	destType   PeerType
	sourceID   string
	destID     string
}

type TransactionOrder string

const (
	CREATED_AT   TransactionOrder = "createdAt"
	LAST_UPDATED TransactionOrder = "lastUpdated"
)

type TransactionStatus string

const (
	SUBMITTED                         TransactionStatus = "SUBMITTED"
	QUEUED                            TransactionStatus = "QUEUED"
	PENDING_SIGNATURE                 TransactionStatus = "PENDING_SIGNATURE"
	PENDING_AUTHORIZATION             TransactionStatus = "PENDING_AUTHORIZATION"
	PENDING_3RD_PARTY_MANUAL_APPROVAL TransactionStatus = "PENDING_3RD_PARTY_MANUAL_APPROVAL"
	PENDING_3RD_PARTY                 TransactionStatus = "PENDING_3RD_PARTY"
	/**
	 * @deprecated
	 */
	PENDING      TransactionStatus = "PENDING"
	BROADCASTING TransactionStatus = "BROADCASTING"
	CONFIRMING   TransactionStatus = "CONFIRMING"
	/**
	 * @deprecated Replaced by "COMPLETED"
	 */
	CONFIRMED             TransactionStatus = "CONFIRMED"
	COMPLETED             TransactionStatus = "COMPLETED"
	PENDING_AML_SCREENING TransactionStatus = "PENDING_AML_SCREENING"
	PARTIALLY_COMPLETED   TransactionStatus = "PARTIALLY_COMPLETED"
	CANCELLING            TransactionStatus = "CANCELLING"
	CANCELLED             TransactionStatus = "CANCELLED"
	REJECTED              TransactionStatus = "REJECTED"
	FAILED                TransactionStatus = "FAILED"
	TIMEOUT               TransactionStatus = "TIMEOUT"
	BLOCKED               TransactionStatus = "BLOCKED"
)

type VirtualType string

const (
	OFF_EXCHANGE VirtualType = "OFF_EXCHANGE"
	DEFAULT      VirtualType = "DEFAULT"
)

type TransactionOperation string

const (
	TRANSFER             TransactionOperation = "TRANSFER"
	MINT                 TransactionOperation = "MINT"
	BURN                 TransactionOperation = "BURN"
	SUPPLY_TO_COMPOUND   TransactionOperation = "SUPPLY_TO_COMPOUND"
	REDEEM_FROM_COMPOUND TransactionOperation = "REDEEM_FROM_COMPOUND"
	RAW                  TransactionOperation = "RAW"
	CONTRACT_CALL        TransactionOperation = "CONTRACT_CALL"
)

type AllocateFundsRequest struct {
	allocationID string
	amount       string
}

type DeallocateFundsRequest struct {
	allocationID string
	amount       string
}

type Terms struct {
	networkConnectionID string
	outgoing            bool
	asset               string
	amount              string
	note                string
}
type CreateTransferTicketArgs struct {
	externalTicketID string
	description      string
	terms            []Terms
}

type TransferTicketStatus string

const (
	TransferTicketStatus_OPEN                TransferTicketStatus = "OPEN"
	TransferTicketStatus_PARTIALLY_FULFILLED TransferTicketStatus = "PARTIALLY_FULFILLED"
	TransferTicketStatus_FULFILLED           TransferTicketStatus = "FULFILLED"
	TransferTicketStatus_FAILED              TransferTicketStatus = "FAILED"
	TransferTicketStatus_CANCELED            TransferTicketStatus = "CANCELED"
)

type TransferTicketTermStatus string

const (
	TransferTicketTermStatus_OPEN      TransferTicketTermStatus = "OPEN"
	TransferTicketTermStatus_FULFILLED TransferTicketTermStatus = "FULFILLED"
)

type TransferTicketResponse struct {
	ticketID         string
	externalTicketId string
	description      string
	status           TransferTicketStatus
	terms            []TermResponse // TODO string
}

type TermResponse struct {
	termID              string
	networkConnectionID string
	outgoing            bool
	asset               string
	amount              string
	txIds               []string
	status              TransferTicketTermStatus
	note                string
}

type ExecuteTermArgsSource struct {
	Type string
	id   string
}
type ExecuteTermArgs struct {
	source   ExecuteTermArgsSource
	fee      int64
	gasPrice int64
}

type CreateTransferTicketResponse struct {
	ticketID string
}

type PublicKeyInfoArgs struct {
	algorithm      string
	derivationPath string
	compressed     bool
}

type PublicKeyInfoForVaultAccountArgs struct {
	assetID        string
	vaultAccountID int64
	change         int64
	addressIndex   int64
	compressed     bool
}

type Configuration struct {
	gasThreshold string
	gasCap       string
	maxGasPrice  string
}

type GasStationInfo struct {
	// balance: { [asset: string]: string }; // TODO
	configuration Configuration
}

type PublicKeyResonse struct {
	status         int64
	algorithm      string
	derivationPath []int64
	publicKey      string
}

type MaxSpendableAmountResponse struct {
	maxSpendableAmount string
}

type VaultAccountsFilter struct {
	namePrefix         string
	nameSuffix         string
	minAmountThreshold int64
}

type VaultBalancesFilter struct {
	accountNamePrefix string
	accountNameSuffix string
}
type RequestOptions struct {
	idempotencyKey string
}

type ValidateAddressResponse struct {
	isValid     bool
	isActive    bool
	requiresTag bool
}
