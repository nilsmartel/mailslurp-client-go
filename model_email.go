/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Overview  #### Inboxes  Inboxes have real email addresses that can send and receive emails. You can create inboxes with specific email addresses (using custom domains). You can also use randomly assigned MailSlurp addresses as unique, disposable test addresses.   See the InboxController or [inbox and email address guide](https://www.mailslurp.com/guides/) for more information.  #### Receive Emails You can receive emails in a number of ways. You can fetch emails and attachments directly from an inbox. Or you can use `waitFor` endpoints to hold a connection open until an email is received that matches given criteria (such as subject or body content). You can also use webhooks to have emails from multiple inboxes forwarded to your server via HTTP POST.  InboxController methods with `waitFor` in the name have a long timeout period and instruct MailSlurp to wait until an expected email is received. You can set conditions on email counts, subject or body matches, and more.  Most receive methods only return an email ID and not the full email (to keep response sizes low). To fetch the full body or attachments for an email use the email's ID with EmailController endpoints.  See the InboxController or [receiving emails guide](https://www.mailslurp.com/guides/) for more information.  #### Send Emails You can send templated HTML emails in several ways. You must first create an inbox to send an email. An inbox can have a specific address or a randomly assigned one. You can send emails from an inbox using `to`, `cc`, and `bcc` recipient lists or with contacts and contact groups.   Emails can contain plain-text or HTML bodies. You can also use email templates that support [moustache](https://mustache.github.io/) template variables. You can send attachments by first posting files to the AttachmentController and then using the returned IDs in the `attachments` field of the send options.  See the InboxController or [sending emails guide](https://www.mailslurp.com/guides/) for more information.  ## Templates MailSlurp emails support templates. You can create templates in the dashboard or API that contain [moustache](https://mustache.github.io/) style variables: for instance `Hello {{name}}`. Then when sending emails you can pass a map of variables names and values to be used. Additionally, when sending emails with contact groups you can use properties of the contact in your templates like `{{firstName}}` and `{{lastName}}``.  ## Explore     
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
import (
	"time"
)
// Email Representation of an email received by an inbox. Use the ID to access more properties of an email using the EmailController endpoints.
type Email struct {
	Analysis EmailAnalysis `json:"analysis,omitempty"`
	// List of IDs of attachments found in the email. Use these IDs with the Inbox and Email Controllers to download attachments and attachment meta data such as filesize, name, extension.
	Attachments []string `json:"attachments,omitempty"`
	// List of `BCC` recipients email was addressed to
	Bcc []string `json:"bcc,omitempty"`
	// The body of the email message
	Body string `json:"body,omitempty"`
	// List of `CC` recipients email was addressed to
	Cc []string `json:"cc,omitempty"`
	// Detected character set of the email body such as UTF-8
	Charset string `json:"charset,omitempty"`
	// When was the email received by MailSlurp
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Who was the email sent from
	From string `json:"from,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	// ID of the email
	Id string `json:"id,omitempty"`
	// ID of the inbox that received the email
	InboxId string `json:"inboxId,omitempty"`
	// Was HTML sent in the email body
	IsHTML bool `json:"isHTML,omitempty"`
	// Has the email been viewed ever
	Read bool `json:"read,omitempty"`
	// The subject line of the email message
	Subject string `json:"subject,omitempty"`
	// List of `To` recipients email was addressed to
	To []string `json:"to,omitempty"`
	// When was the email last updated
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// ID of user that email belongs
	UserId string `json:"userId,omitempty"`
}
