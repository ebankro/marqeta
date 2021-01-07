/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type TransactionModel struct {
	Identifier string `json:"identifier"`
	Token string `json:"token,omitempty"`
	UserToken string `json:"user_token,omitempty"`
	BusinessToken string `json:"business_token,omitempty"`
	ActingUserToken string `json:"acting_user_token"`
	CardToken string `json:"card_token,omitempty"`
	Type_ string `json:"type,omitempty"`
	State string `json:"state"`
	Duration int32 `json:"duration,omitempty"`
	CreatedTime time.Time `json:"created_time,omitempty"`
	UserTransactionTime time.Time `json:"user_transaction_time,omitempty"`
	SettlementDate time.Time `json:"settlement_date,omitempty"`
	RequestAmount float32 `json:"request_amount,omitempty"`
	Amount float32 `json:"amount"`
	CashBackAmount float32 `json:"cash_back_amount,omitempty"`
	CurrencyConversion *CurrencyConversion `json:"currency_conversion,omitempty"`
	IssuerInterchangeAmount float64 `json:"issuer_interchange_amount,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
	ApprovalCode string `json:"approval_code,omitempty"`
	Response *Response `json:"response,omitempty"`
	PrecedingRelatedTransactionToken string `json:"preceding_related_transaction_token,omitempty"`
	PrecedingTransaction *PrecedingTransaction `json:"preceding_transaction,omitempty"`
	AmountToBeReleased float32 `json:"amount_to_be_released,omitempty"`
	IncrementalAuthorizationTransactionTokens []string `json:"incremental_authorization_transaction_tokens,omitempty"`
	Merchant *MerchantResponseModel `json:"merchant,omitempty"`
	Store *StoreResponseModel `json:"store,omitempty"`
	CardAcceptor *TransactionCardAcceptor `json:"card_acceptor,omitempty"`
	Gpa *CardholderBalance `json:"gpa,omitempty"`
	Card *CardResponse `json:"card,omitempty"`
	GpaOrderUnload *GpaReturns `json:"gpa_order_unload,omitempty"`
	GpaOrder *GpaResponse `json:"gpa_order,omitempty"`
	ProgramTransfer *ProgramTransferResponse `json:"program_transfer,omitempty"`
	FeeTransfer *FeeTransferResponse `json:"fee_transfer,omitempty"`
	PeerTransfer *PeerTransferResponse `json:"peer_transfer,omitempty"`
	MsaOrders []MsaOrderResponse `json:"msa_orders,omitempty"`
	MsaOrderUnload *MsaReturns `json:"msa_order_unload,omitempty"`
	OfferOrders []OfferOrderResponse `json:"offer_orders,omitempty"`
	AutoReload *AutoReloadModel `json:"auto_reload,omitempty"`
	DirectDeposit *DepositDepositResponse `json:"direct_deposit,omitempty"`
	Polarity string `json:"polarity,omitempty"`
	RealTimeFeeGroup *RealTimeFeeGroup `json:"real_time_fee_group,omitempty"`
	Fee *Fee `json:"fee,omitempty"`
	Chargeback *ChargebackResponse `json:"chargeback,omitempty"`
	Dispute *DisputeModel `json:"dispute,omitempty"`
	Network string `json:"network,omitempty"`
	Subnetwork string `json:"subnetwork,omitempty"`
	NetworkMetadata *NetworkMetadata `json:"network_metadata,omitempty"`
	AcquirerFeeAmount float32 `json:"acquirer_fee_amount,omitempty"`
	Fees []NetworkFeeModel `json:"fees,omitempty"`
	DigitalWalletToken *DigitalWalletToken `json:"digital_wallet_token,omitempty"`
	User *CardholderMetadata `json:"user,omitempty"`
	Business *BusinessMetadata `json:"business,omitempty"`
	Acquirer *Acquirer `json:"acquirer,omitempty"`
	Fraud *FraudView `json:"fraud,omitempty"`
	Pos *Pos `json:"pos,omitempty"`
	AddressVerification *AddressVerificationModel `json:"address_verification,omitempty"`
	CardSecurityCodeVerification *CardSecurityCodeVerification `json:"card_security_code_verification,omitempty"`
	TransactionMetadata *TransactionMetadata `json:"transaction_metadata,omitempty"`
	OriginalCredit *OriginalCredit `json:"original_credit,omitempty"`
	CardHolderModel *UserCardHolderResponse `json:"card_holder_model,omitempty"`
	StandinApprovedBy string `json:"standin_approved_by,omitempty"`
	StandinBy string `json:"standin_by,omitempty"`
	StandinReason string `json:"standin_reason,omitempty"`
	NetworkReferenceId string `json:"network_reference_id,omitempty"`
	AcquirerReferenceId string `json:"acquirer_reference_id,omitempty"`
	CardholderAuthenticationData *CardholderAuthenticationData `json:"cardholder_authentication_data,omitempty"`
	TransactionAttributes map[string]string `json:"transaction_attributes,omitempty"`
	ClearingRecordSequenceNumber string `json:"clearing_record_sequence_number,omitempty"`
	IssuerReceivedTime string `json:"issuer_received_time,omitempty"`
	IssuerPaymentNode string `json:"issuer_payment_node,omitempty"`
	Program *Program `json:"program,omitempty"`
	BatchNumber string `json:"batch_number,omitempty"`
	FromAccount string `json:"from_account,omitempty"`
	MultiClearingSequenceNumber string `json:"multi_clearing_sequence_number,omitempty"`
	MultiClearingSequenceCount string `json:"multi_clearing_sequence_count,omitempty"`
}
