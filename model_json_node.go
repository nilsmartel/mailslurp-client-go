/*
 * MailSlurp API
 *
 * ## Introduction  [MailSlurp](https://www.mailslurp.com) is an Email API for developers and QA testers. It let's users: - create emails addresses on demand - receive emails and attachments in code - send templated HTML emails  ## About  This page contains the REST API documentation for MailSlurp. All requests require API Key authentication passed as an `x-api-key` header.  Create an account to [get your free API Key](https://app.mailslurp.com/sign-up/).  ## Resources - 🔑 [Get API Key](https://app.mailslurp.com/sign-up/)                    - 🎓 [Developer Portal](https://www.mailslurp.com/docs/)           - 📦 [Library SDKs](https://www.mailslurp.com/docs/) - ✍️ [Code Examples](https://www.mailslurp.com/examples) - ⚠️ [Report an issue](https://drift.me/mailslurp)  ## Explore  
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
// JsonNode struct for JsonNode
type JsonNode struct {
	Array bool `json:"array,omitempty"`
	BigDecimal bool `json:"bigDecimal,omitempty"`
	BigInteger bool `json:"bigInteger,omitempty"`
	Binary bool `json:"binary,omitempty"`
	Boolean bool `json:"boolean,omitempty"`
	ContainerNode bool `json:"containerNode,omitempty"`
	Double bool `json:"double,omitempty"`
	Empty bool `json:"empty,omitempty"`
	Float bool `json:"float,omitempty"`
	FloatingPointNumber bool `json:"floatingPointNumber,omitempty"`
	Int bool `json:"int,omitempty"`
	IntegralNumber bool `json:"integralNumber,omitempty"`
	Long bool `json:"long,omitempty"`
	MissingNode bool `json:"missingNode,omitempty"`
	NodeType string `json:"nodeType,omitempty"`
	Null bool `json:"null,omitempty"`
	Number bool `json:"number,omitempty"`
	Object bool `json:"object,omitempty"`
	Pojo bool `json:"pojo,omitempty"`
	Short bool `json:"short,omitempty"`
	Textual bool `json:"textual,omitempty"`
	ValueNode bool `json:"valueNode,omitempty"`
}