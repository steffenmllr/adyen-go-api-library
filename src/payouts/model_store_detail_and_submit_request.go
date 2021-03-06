/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payouts

import (
	"time"
)

// StoreDetailAndSubmitRequest struct for StoreDetailAndSubmitRequest
type StoreDetailAndSubmitRequest struct {
	// This field contains additional data, which may be required for a particular request.
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
	Amount         Amount                  `json:"amount"`
	Bank           *BankAccount            `json:"bank,omitempty"`
	BillingAddress *Address                `json:"billingAddress,omitempty"`
	Card           *Card                   `json:"card,omitempty"`
	// The date of birth. Format: [ISO-8601](https://www.w3.org/TR/NOTE-datetime); example: YYYY-MM-DD For Paysafecard it must be the same as used when registering the Paysafecard account. > This field is mandatory for natural persons.
	DateOfBirth time.Time `json:"dateOfBirth"`
	// The type of the entity the payout is processed for.
	EntityType string `json:"entityType"`
	// An integer value that is added to the normal fraud score. The value can be either positive or negative.
	FraudOffset int32 `json:"fraudOffset,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The shopper's nationality.  A valid value is an ISO 2-character country code (e.g. 'NL').
	Nationality string    `json:"nationality"`
	Recurring   Recurring `json:"recurring"`
	// The merchant reference for this payment. This reference will be used in all communication to the merchant about the status of the payout. Although it is a good idea to make sure it is unique, this is not a requirement.
	Reference string `json:"reference"`
	// The name of the brand to make a payout to.  For Paysafecard it must be set to `paysafecard`.
	SelectedBrand string `json:"selectedBrand,omitempty"`
	// The shopper's email address.
	ShopperEmail string `json:"shopperEmail"`
	ShopperName  *Name  `json:"shopperName,omitempty"`
	// The shopper's reference for the payment transaction.
	ShopperReference string `json:"shopperReference"`
	// The description of this payout. This description is shown on the bank statement of the shopper (if this is supported by the chosen payment method).
	ShopperStatement string `json:"shopperStatement,omitempty"`
	// The shopper's social security number.
	SocialSecurityNumber string `json:"socialSecurityNumber,omitempty"`
	// The shopper's phone number.
	TelephoneNumber string `json:"telephoneNumber,omitempty"`
}
