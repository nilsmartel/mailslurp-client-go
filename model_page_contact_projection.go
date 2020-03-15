/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 67c9a1eda264be4cfe0bb2c76151f0aadf0862bc
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
// PageContactProjection struct for PageContactProjection
type PageContactProjection struct {
	Content []ContactProjection `json:"content,omitempty"`
	Empty bool `json:"empty,omitempty"`
	First bool `json:"first,omitempty"`
	Last bool `json:"last,omitempty"`
	Number int32 `json:"number,omitempty"`
	NumberOfElements int32 `json:"numberOfElements,omitempty"`
	Pageable Pageable `json:"pageable,omitempty"`
	Size int32 `json:"size,omitempty"`
	Sort Sort `json:"sort,omitempty"`
	TotalElements int64 `json:"totalElements,omitempty"`
	TotalPages int32 `json:"totalPages,omitempty"`
}
