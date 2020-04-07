/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payout

// BankAccount struct for BankAccount
type BankAccount struct {
	// The bank account number (without separators).
	BankAccountNumber string `json:"bankAccountNumber,omitempty"`
	// The bank city.
	BankCity string `json:"bankCity,omitempty"`
	// The location id of the bank. The field value is `nil` in most cases.
	BankLocationId string `json:"bankLocationId,omitempty"`
	// The name of the bank.
	BankName string `json:"bankName,omitempty"`
	// The [Business Identifier Code](https://en.wikipedia.org/wiki/ISO_9362) (BIC) is the SWIFT address assigned to a bank. The field value is `nil` in most cases.
	Bic string `json:"bic,omitempty"`
	// Country code where the bank is located.  A valid value is an ISO two-character country code (e.g. 'NL').
	CountryCode string `json:"countryCode,omitempty"`
	// The [International Bank Account Number](https://en.wikipedia.org/wiki/International_Bank_Account_Number) (IBAN).
	Iban string `json:"iban,omitempty"`
	// The name of the bank account holder. If you submit a name with non-Latin characters, we automatically replace some of them with corresponding Latin characters to meet the FATF recommendations. For example: * χ12 is converted to ch12. * üA is converted to euA. * Peter Møller is converted to Peter Mller, because banks don't accept 'ø'. After replacement, the ownerName must have at least three alphanumeric characters (A-Z, a-z, 0-9), and at least one of them must be a valid Latin character (A-Z, a-z). For example: * John17 - allowed. * J17 - allowed. * 171 - not allowed. * John-7 - allowed. > If provided details don't match the required format, the response returns the error message: 203 'Invalid bank account holder name'.
	OwnerName string `json:"ownerName,omitempty"`
	// The bank account holder's tax ID.
	TaxId string `json:"taxId,omitempty"`
}
