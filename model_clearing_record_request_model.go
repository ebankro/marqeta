/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

import (
	"time"
)

type ClearingRecordRequestModel struct {
	NetworkFees                     []NetworkFeeModel `json:"network_fees,omitempty"`
	Webhook                         *Webhook          `json:"webhook,omitempty"`
	Mid                             string            `json:"mid,omitempty"`
	Amount                          float32           `json:"amount"`
	SourceAmount                    float32           `json:"source_amount"`
	ReconciliationAmount            float32           `json:"reconciliation_amount"`
	ReplacementAmount               float32           `json:"replacement_amount,omitempty"`
	CardholderBillingAmount         float32           `json:"cardholder_billing_amount,omitempty"`
	Cashback                        float32           `json:"cashback,omitempty"`
	LocalTransactionAmount          float32           `json:"local_transaction_amount,omitempty"`
	LocalCurrencyCode               string            `json:"local_currency_code,omitempty"`
	CardholderBillingConversionRate float64           `json:"cardholder_billing_conversion_rate,omitempty"`
	CardholderBillingCurrency       string            `json:"cardholder_billing_currency,omitempty"`
	CardToken                       string            `json:"card_token"`
	CardHash                        string            `json:"card_hash"`
	AcquirerReferenceId             string            `json:"acquirer_reference_id,omitempty"`
	Rrn                             string            `json:"rrn,omitempty"`
	Stan                            string            `json:"stan,omitempty"`
	ProcessingCode                  string            `json:"processing_code,omitempty"`
	AcquirerFee                     *MoneyModel       `json:"acquirer_fee,omitempty"`
	IssuerFee                       *MoneyModel       `json:"issuer_fee,omitempty"`
	FunctionCode                    string            `json:"function_code,omitempty"`
	ReasonCode                      string            `json:"reason_code,omitempty"`
	ApprovalCode                    string            `json:"approval_code,omitempty"`
	// yyyy-MM-dd, yyyy-MM-ddThh:mm:ssZ
	TransactionDate *time.Time `json:"transaction_date,omitempty"`
	// yyyy-MM-dd, yyyy-MM-ddThh:mm:ssZ
	LocalTransactionDate *time.Time `json:"local_transaction_date,omitempty"`
	// yyyy-MM-dd, yyyy-MM-ddThh:mm:ssZ
	SettlementDate         *time.Time `json:"settlement_date,omitempty"`
	NetworkReferenceId     string     `json:"network_reference_id,omitempty"`
	FindOriginalWindowDays int32      `json:"find_original_window_days,omitempty"`
	BatchNumber            string     `json:"batch_number,omitempty"`
	BatchFileName          string     `json:"batch_file_name,omitempty"`
	SequenceNumber         string     `json:"sequence_number,omitempty"`
	MultiClearingCount     string     `json:"multi_clearing_count,omitempty"`
	Network                string     `json:"network,omitempty"`
	// Defaults to VISANET if network is VISA
	SubNetwork                         string                `json:"sub_network,omitempty"`
	CardAcceptor                       *CardAcceptorModel    `json:"card_acceptor,omitempty"`
	CurrencyCode                       string                `json:"currency_code"`
	OriginalDataElements               *OriginalDataElements `json:"original_data_elements,omitempty"`
	PrecedingRelatedTransactionToken   string                `json:"preceding_related_transaction_token,omitempty"`
	SendExpirationDate                 bool                  `json:"send_expiration_date,omitempty"`
	SimulateBatchForClearingRecordHash bool                  `json:"simulate_batch_for_clearing_record_hash,omitempty"`
	CrossBorderIndicator               string                `json:"cross_border_indicator,omitempty"`
	TokenPan                           string                `json:"token_pan,omitempty"`
	PaymentChannelIndicator            string                `json:"payment_channel_indicator,omitempty"`
	IsInstallment                      bool                  `json:"is_installment,omitempty"`
	IsRecurring                        bool                  `json:"is_recurring,omitempty"`
	NetworkMetadata                    *NetworkMetadata      `json:"network_metadata,omitempty"`
	AuthorizationSourceCode            string                `json:"authorization_source_code,omitempty"`
	InterchangeRateDescriptor          string                `json:"interchange_rate_descriptor,omitempty"`
	Mti                                string                `json:"mti,omitempty"`
}
