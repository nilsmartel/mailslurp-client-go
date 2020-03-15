/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 67c9a1eda264be4cfe0bb2c76151f0aadf0862bc
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
import (
	"time"
)
// ContactProjection struct for ContactProjection
type ContactProjection struct {
	Company string `json:"company,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	FirstName string `json:"firstName,omitempty"`
	Id string `json:"id"`
	LastName string `json:"lastName,omitempty"`
	OptOut bool `json:"optOut,omitempty"`
}
