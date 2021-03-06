/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type CardHolderAddressUpdateModel struct {
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Address1         string `json:"address_1,omitempty"`
	Address2         string `json:"address_2,omitempty"`
	City             string `json:"city,omitempty"`
	State            string `json:"state,omitempty"`
	Zip              string `json:"zip,omitempty"`
	Country          string `json:"country,omitempty"`
	Phone            string `json:"phone,omitempty"`
	IsDefaultAddress bool   `json:"is_default_address,omitempty"`
	Active           bool   `json:"active,omitempty"`
	PostalCode       string `json:"postal_code,omitempty"`
}
