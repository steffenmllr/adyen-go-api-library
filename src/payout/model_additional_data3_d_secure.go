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

// AdditionalData3DSecure struct for AdditionalData3DSecure
type AdditionalData3DSecure struct {
	// This parameter indicates that you are able to process 3D Secure 2 transactions natively on your payment page. Send this field when you are using `/payments` endpoint with any of our [native 3D Secure 2 solutions](https://docs.adyen.com/checkout/3d-secure/native-3ds2), such as Components or Drop-in. Possible values: * **true** - Ready to support native 3D Secure 2 authentication. Setting this to true does not mean always applying 3D Secure 2. Adyen still selects the version of 3D Secure based on configuration to optimize authorisation rates and improve the shopper's experience. * **false** – Not ready to support native 3D Secure 2 authentication. Adyen will not offer 3D Secure 2 to your shopper regardless of your configuration. > This parameter only indicates your readiness to support 3D Secure 2 natively on Drop-in or Components. To specify that you want to perform 3D Secure on a transaction, use Dynamic 3D Secure or send the executeThreeD parameter.
	Allow3DS2 string `json:"allow3DS2,omitempty"`
	// This parameter indicates if you want to perform 3D Secure authentication on a transaction or not. Allowed values: * **true** – Perform 3D Secure authentication. * **false** – Don't perform 3D Secure authentication. > Alternatively, you can also use Dynamic 3D Secure to configure rules for applying 3D Secure.
	ExecuteThreeD string `json:"executeThreeD,omitempty"`
	// In case of Secure+, this field must be set to **CUPSecurePlus**.
	MpiImplementationType string `json:"mpiImplementationType,omitempty"`
	// Indicates the [exemption type](https://docs-admin.is.adyen.com/payments-fundamentals/psd2-sca-compliance-and-implementation-guide#specifypreferenceinyourapirequest) that you want to request for the transaction. Possible values: * **lowValue**  * **secureCorporate**  * **trustedBeneficiary**  * **transactionRiskAnalysis**
	ScaExemption string `json:"scaExemption,omitempty"`
}
