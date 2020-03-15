/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 67c9a1eda264be4cfe0bb2c76151f0aadf0862bc
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
// TemplateVariable struct for TemplateVariable
type TemplateVariable struct {
	Name string `json:"name"`
	VariableType string `json:"variableType"`
}
