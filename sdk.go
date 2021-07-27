package fireblockssdk

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
)

/*
CreateVaultAccount Creates a new vault account
@param name A name for the new vault account
@param hiddenOnUI If true, the created account and all related transactions will not be shown on Fireblocks console
@param customerRefId A customer reference ID
*/
func (f FireblocksSDK) CreateVaultAccount(name string, hiddenOnUI bool, customerRefId string, autoFuel bool) (*VaultAccountResponse, error) {
	resp, err := f.request.post("/v1/vault/accounts", map[string]interface{}{
		"name":          name,
		"hiddenOnUI":    hiddenOnUI,
		"customerRefId": customerRefId,
		"autoFuel":      autoFuel,
	}, nil)
	if err != nil {
		return nil, err
	}

	var result VaultAccountResponse
	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
CreateVaultAsset Creates a new asset within an existing vault account
@param vaultAccountId The vault account ID
@param assetId The asset to add
*/
func (f FireblocksSDK) CreateVaultAsset(vaultAccountID, assetID string) (*CreateVaultAssetResponse, error) {
	resp, err := f.request.post(fmt.Sprintf("/v1/vault/accounts/%s/%s", vaultAccountID, assetID), map[string]interface{}{}, nil)
	if err != nil {
		return nil, err
	}

	var result CreateVaultAssetResponse
	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetDepositAddresses Gets deposit addresses for an asset in a vault account
@param vaultAccountId The vault account ID
@param assetId The ID of the asset for which to get the deposit address
*/
func (f FireblocksSDK) GetDepositAddresses(vaultAccountID, assetID string) ([]DepositAddressResponse, error) {
	resp, err := f.request.get(fmt.Sprintf("/v1/vault/accounts/%s/%s/addresses", vaultAccountID, assetID))
	if err != nil {
		return nil, err
	}

	var result []DepositAddressResponse
	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

/*
GenerateNewAddress Generates a new address for an asset in a vault account
@param vaultAccountId The vault account ID
@param assetId The ID of the asset for which to generate the deposit address
@param description A description for the new address
@param customerRefId A customer reference ID
*/
func (f FireblocksSDK) GenerateNewAddress(vaultAccountID, assetID, description, customerRefID string) (*GenerateAddressResponse, error) {
	resp, err := f.request.post(fmt.Sprintf("/v1/vault/accounts/%s/%s/addresses", vaultAccountID, assetID), map[string]interface{}{
		"description":   description,
		"customerRefId": customerRefID,
	}, nil)
	if err != nil {
		return nil, err
	}

	var result GenerateAddressResponse
	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetTransactions Gets a list of transactions matching the given filter
@param filter.before Only gets transactions created before a given timestamp (in milliseconds)
@param filter.after Only gets transactions created after a given timestamp (in milliseconds)
@param filter.status Only gets transactions with the specified status
@param filter.limit Limit the amount of returned results. If not specified, a limit of 200 results will be used
@param filter.orderBy Determines the order of the results
*/
func (f FireblocksSDK) GetTransactions(filter TransactionFilter) ([]TransactionResponse, error) {
	v, _ := query.Values(filter)
	resp, err := f.request.get("/v1/transactions?" + v.Encode())
	if err != nil {
		return nil, err
	}

	var result []TransactionResponse
	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

/*
GetTransactionByID Gets detailed information for a single transaction
@param txId The transaction id to query
*/
func (f FireblocksSDK) GetTransactionByID(txID string) (*TransactionResponse, error) {
	resp, err := f.request.get(fmt.Sprintf("/v1/transactions/%s", txID))
	if err != nil {
		return nil, err
	}

	var result TransactionResponse
	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
CancelTransactionByID Cancels the selected transaction
@param txId The transaction id to cancel
*/
func (f FireblocksSDK) CancelTransactionByID(txID string) (*OperationSuccessResponse, error) {
	resp, err := f.request.post(fmt.Sprintf("/v1/transactions/%s", txID), map[string]interface{}{}, nil)
	if err != nil {
		return nil, err
	}

	var result OperationSuccessResponse
	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
CreateTransaction Creates a new transaction with the specified options
*/
func (f FireblocksSDK) CreateTransaction(transactionArguments TransactionArguments, requestOptions *RequestOptions) (*CreateTransactionResponse, error) {
	resp, err := f.request.post("/v1/transactions", transactionArguments, requestOptions)
	if err != nil {
		return nil, err
	}

	var result CreateTransactionResponse
	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
