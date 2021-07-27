package fireblockssdk

type PeerType string

const (
	PEER_TYPE_VAULT_ACCOUNT      PeerType = "VAULT_ACCOUNT"
	PEER_TYPE_EXCHANGE_ACCOUNT   PeerType = "EXCHANGE_ACCOUNT"
	PEER_TYPE_INTERNAL_WALLET    PeerType = "INTERNAL_WALLET"
	PEER_TYPE_EXTERNAL_WALLET    PeerType = "EXTERNAL_WALLET"
	PEER_TYPE_UNKNOWN            PeerType = "UNKNOWN"
	PEER_TYPE_NETWORK_CONNECTION PeerType = "NETWORK_CONNECTION"
	PEER_TYPE_FIAT_ACCOUNT       PeerType = "FIAT_ACCOUNT"
	PEER_TYPE_COMPOUND           PeerType = "COMPOUND"
	PEER_TYPE_ONE_TIME_ADDRESS   PeerType = "ONE_TIME_ADDRESS"
)

type TransactionStatus string

const (
	TRANSACTION_STATUS_SUBMITTED                         TransactionStatus = "SUBMITTED"
	TRANSACTION_STATUS_QUEUED                            TransactionStatus = "QUEUED"
	TRANSACTION_STATUS_PENDING_SIGNATURE                 TransactionStatus = "PENDING_SIGNATURE"
	TRANSACTION_STATUS_PENDING_AUTHORIZATION             TransactionStatus = "PENDING_AUTHORIZATION"
	TRANSACTION_STATUS_PENDING_3RD_PARTY_MANUAL_APPROVAL TransactionStatus = "PENDING_3RD_PARTY_MANUAL_APPROVAL"
	TRANSACTION_STATUS_PENDING                           TransactionStatus = "PENDING"
	TRANSACTION_STATUS_BROADCASTING                      TransactionStatus = "BROADCASTING"
	TRANSACTION_STATUS_CONFIRMING                        TransactionStatus = "CONFIRMING"
	TRANSACTION_STATUS_CONFIRMED                         TransactionStatus = "CONFIRMED"
	TRANSACTION_STATUS_COMPLETED                         TransactionStatus = "COMPLETED"
	TRANSACTION_STATUS_PENDING_AML_SCREENING             TransactionStatus = "PENDING_AML_SCREENING"
	TRANSACTION_STATUS_PARTIALLY_COMPLETED               TransactionStatus = "PARTIALLY_COMPLETED"
	TRANSACTION_STATUS_CANCELLING                        TransactionStatus = "CANCELLING"
	TRANSACTION_STATUS_CANCELLED                         TransactionStatus = "CANCELLED"
	TRANSACTION_STATUS_REJECTED                          TransactionStatus = "REJECTED"
	TRANSACTION_STATUS_FAILED                            TransactionStatus = "FAILED"
	TRANSACTION_STATUS_TIMEOUT                           TransactionStatus = "TIMEOUT"
	TRANSACTION_STATUS_BLOCKED                           TransactionStatus = "BLOCKED"
	TRANSACTION_STATUS_DROPPED                           TransactionStatus = "DROPPED"
)

type TransactionOperation string

const (
	TRANSACTION_OPERATION_TRANSFER             TransactionOperation = "TRANSFER"
	TRANSACTION_OPERATION_MINT                 TransactionOperation = "MINT"
	TRANSACTION_OPERATION_BURN                 TransactionOperation = "BURN"
	TRANSACTION_OPERATION_SUPPLY_TO_COMPOUND   TransactionOperation = "SUPPLY_TO_COMPOUND"
	TRANSACTION_OPERATION_REDEEM_FROM_COMPOUND TransactionOperation = "REDEEM_FROM_COMPOUND"
	TRANSACTION_OPERATION_RAW                  TransactionOperation = "RAW"
	TRANSACTION_OPERATION_CONTRACT_CALL        TransactionOperation = "CONTRACT_CALL"
)

type NetworkStatus string

const (
	NETWORK_STATUS_DROPPED      NetworkStatus = "DROPPED"
	NETWORK_STATUS_BROADCASTING NetworkStatus = "BROADCASTING"
	NETWORK_STATUS_CONFIRMING   NetworkStatus = "CONFIRMING"
	NETWORK_STATUS_FAILED       NetworkStatus = "FAILED"
	NETWORK_STATUS_CONFIRMED    NetworkStatus = "CONFIRMED"
)

type TransactionOrder string

const (
	TRANSACTION_ORDER_CREATED_AT   TransactionOrder = "createdAt"
	TRANSACTION_ORDER_LAST_UPDATED TransactionOrder = "lastUpdated"
)

type FeeLevel string

const (
	FEE_LEVEL_HIGH   FeeLevel = "HIGH"
	FEE_LEVEL_MEDIUM FeeLevel = "MEDIUM"
	FEE_LEVEL_LOW    FeeLevel = "LOW"
)

type VaultAccountResponse struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	HiddenOnUI    bool            `json:"hiddenOnUi"`
	Assets        []AssetResponse `json:"assets"`
	CustomerRefID string          `json:"customerRefId,omitempty"`
}

type VirtualType string

const (
	VIRTUAL_TYPE_OFF_EXCHANGE VirtualType = "OFF_EXCHANGE"
	VIRTUAL_TYPE_DEFAULT      VirtualType = "DEFAULT"
	VIRTUAL_TYPE_OEC_FEE_BANK VirtualType = "OEC_FEE_BANK"
)

type AssetResponse struct {
	ID                   string `json:"id"`
	Total                string `json:"total"`
	Balance              string `json:"balance"`
	LockedAmount         string `json:"lockedAmount,omitempty"`
	Available            string `json:"available"`
	Pending              string `json:"pending"`
	SelfStakedCPU        string `json:"selfStakedCPU,omitempty"`
	SelfStakedNetwork    string `json:"selfStakedNetwork,omitempty"`
	PendingRefundCPU     string `json:"pendingRefundCPU,omitempty"`
	PendingRefundNetwork string `json:"pendingRefundNetwork,omitempty"`
	TotalStakedCPU       string `json:"totalStakedCPU,omitempty"`
	TotalStakedNetwork   string `json:"totalStakedNetwork,omitempty"`
}

type CreateVaultAssetResponse struct {
	ID             string `json:"id"`
	Address        string `json:"address"`
	LegacyAddress  string `json:"legacyAddress"`
	Tag            string `json:"tag"`
	EosAccountName string `json:"eosAccountName"`
}

type DepositAddressResponse struct {
	AssetID           string `json:"assetId"`
	Address           string `json:"address"`
	Tag               string `json:"tag"`
	Description       string `json:"description"`
	Type              string `json:"type"`
	LegacyAddress     string `json:"legacyAddress"`
	EnterpriseAddress string `json:"enterpriseAddress"`
}

type GenerateAddressResponse struct {
	Address       string `json:"address"`
	Tag           string `json:"tag"`
	LegacyAddress string `json:"legacyAddress"`
}

type TransactionResponse struct {
	ID                            string                   `json:"id"`
	AssetID                       string                   `json:"assetId"`
	Source                        TransferPeerPathResponse `json:"source"`
	Destination                   TransferPeerPathResponse `json:"destination"`
	RequestedAmount               float64                  `json:"requestedAmount"`
	AmountInfo                    AmountInfo               `json:"amountInfo"`
	FeeInfo                       FeeInfo                  `json:"feeInfo"`
	Amount                        float64                  `json:"amount"`
	NetAmount                     float64                  `json:"netAmount"`
	AmountUSD                     float64                  `json:"amountUsd"`
	ServiceFee                    float64                  `json:"serviceFee"`
	NetworkFee                    float64                  `json:"networkFee"`
	CreatedAt                     int64                    `json:"createdAt"`
	LastUpdated                   int64                    `json:"lastUpdated"`
	Status                        TransactionStatus        `json:"status"`
	TxHash                        string                   `json:"txHash"`
	SubStatus                     string                   `json:"subStatus"`
	SourceAddress                 string                   `json:"sourceAddress"`
	DestinationAddress            string                   `json:"destinationAddress"`
	DestinationAddressDescription string                   `json:"destinationAddressDescription"`
	DestinationTag                string                   `json:"destinationTag"`
	SignedBy                      []string                 `json:"signedBy"`
	CreatedBy                     string                   `json:"createdBy"`
	RejectedBy                    string                   `json:"rejectedBy"`
	AddressType                   string                   `json:"addressType"`
	Note                          string                   `json:"note"`
	ExchangeTxID                  string                   `json:"exchangeTxId"`
	FeeCurrency                   string                   `json:"feeCurrency"`
	Operation                     TransactionOperation     `json:"operation"`
	AmlScreeningResult            AmlScreeningResult       `json:"amlScreeningResult"`
	CustomerRefId                 string                   `json:"customerRefId"`
	NumOfConfirmations            int64                    `json:"numOfConfirmations"`
	NetworkRecords                []NetworkRecord          `json:"networkRecords"`
	ReplacedTxHash                string                   `json:"replacedTxHash"`
	Destinations                  []DestinationsResponse   `json:"destinations"`
	SignedMessages                []SignedMessageResponse  `json:"signedMessages"`
	ExtraParameters               interface{}              `json:"extraParameters"`
}

type TransferPeerPathResponse struct {
	Type    PeerType `json:"type"`
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	SubType string   `json:"subType"`
}

type TransferPeerPath struct {
	Type PeerType `json:"type"`
	ID   string   `json:"id"`
}

type AmountInfo struct {
	Amount          string `json:"amount"`
	RequestedAmount string `json:"requestedAmount"`
	NetAmount       string `json:"netAmount"`
	AmountUSD       string `json:"amountUSD"`
}

type FeeInfo struct {
	NetworkFee string `json:"networkFee"`
	ServiceFee string `json:"serviceFee"`
}

type AmlScreeningResult struct {
	Provider        string      `json:"provider"`
	Payload         interface{} `json:"payload"`
	ScreeningStatus string      `json:"screeningStatus"`
	BypassReason    string      `json:"bypassReason"`
	Timestamp       int64       `json:"timestamp"`
}

type NetworkRecord struct {
	Source      TransferPeerPathResponse `json:"source"`
	Destination TransferPeerPathResponse `json:"destination"`
	TxHash      string                   `json:"txHash"`
	NetworkFee  float64                  `json:"networkFee"`
	AssetID     string                   `json:"assetId"`
	NetAmount   float64                  `json:"netAmount"`
	Status      NetworkStatus            `json:"status"`
}

type DestinationsResponse struct {
	Amount                        string                   `json:"amount"`
	Destination                   TransferPeerPathResponse `json:"destination"`
	AmountUSD                     float64                  `json:"amountUSD"`
	DestinationAddress            string                   `json:"destinationAddress"`
	DestinationAddressDescription string                   `json:"destinationAddressDescription"`
	AmlScreeningResult            AmlScreeningResult       `json:"amlScreeningResult"`
	CustomerRefId                 string                   `json:"customerRefId"`
}

type SignedMessageResponse struct {
	Content        string    `json:"content"`
	Algorithm      string    `json:"algorithm"`
	DerivationPath string    `json:"derivationPath"`
	Signature      Signature `json:"signature"`
	PublicKey      string    `json:"publicKey"`
}

type Signature struct {
	FullSig string `json:"fullSig"`
	R       string `json:"r"`
	S       string `json:"s"`
	V       int64  `json:"v"`
}

type TransactionFilter struct {
	Before     int64             `url:"before,omitempty"`
	After      int64             `url:"after,omitempty"`
	Status     TransactionStatus `url:"status,omitempty"`
	OrderBy    TransactionOrder  `url:"orderBy,omitempty"`
	Limit      int64             `url:"limit,omitempty"`
	TxHash     string            `url:"txHash,omitempty"`
	Assets     string            `url:"assets,omitempty"`
	SourceType PeerType          `url:"sourceType,omitempty"`
	DestType   PeerType          `url:"destType,omitempty"`
	SourceID   string            `url:"sourceId,omitempty"`
	DestID     string            `url:"destId,omitempty"`
}

type OperationSuccessResponse struct {
	Success bool `json:"success"`
}

type TransactionArguments struct {
	AssetID         string                           `json:"assetId,omitempty"`
	Source          *TransferPeerPath                `json:"source,omitempty"`
	Destination     *DestinationTransferPeerPath     `json:"destination,omitempty"`
	Destinations    *[]TransactionRequestDestination `json:"destinations,omitempty"`
	Amount          string                           `json:"amount,omitempty"`
	Fee             string                           `json:"fee,omitempty"`
	GasPrice        string                           `json:"gasPrice,omitempty"`
	GasLimit        string                           `json:"gasLimit,omitempty"`
	NetworkFee      string                           `json:"networkFee,omitempty"`
	FeeLevel        FeeLevel                         `json:"feeLevel,omitempty"`
	MaxFee          string                           `json:"maxFee,omitempty"`
	FailOnLowFee    bool                             `json:"failOnLowFee,omitempty"`
	Note            string                           `json:"note,omitempty"`
	AutoStaking     bool                             `json:"autoStaking,omitempty"`
	NetworkStaking  string                           `json:"networkStaking,omitempty"`
	CpuStaking      string                           `json:"cpuStaking,omitempty"`
	Operation       *TransactionOperation            `json:"operation,omitempty"`
	CustomerRefID   string                           `json:"customerRefId,omitempty"`
	ReplaceTxByHash string                           `json:"replaceTxByHash,omitempty"`
	ExtraParameters interface{}                      `json:"extraParameters,omitempty"`
}

type TransactionRequestDestination struct {
	Amount      string                      `json:"amount"`
	Destination DestinationTransferPeerPath `json:"destination"`
}

type DestinationTransferPeerPath struct {
	ID             string          `json:"id,omitempty"`
	Type           PeerType        `json:"type"`
	OneTimeAddress *OneTimeAddress `json:"oneTimeAddress"`
}

type OneTimeAddress struct {
	Address string `json:"address"`
	Tag     string `json:"tag,omitempty"`
}

type CreateTransactionResponse struct {
	ID     string            `json:"id"`
	Status TransactionStatus `json:"status"`
}

type RequestOptions struct {
	IdempotencyKey string `json:"idempotencyKey"`
}
