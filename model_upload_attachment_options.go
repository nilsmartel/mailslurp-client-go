/*
 * MailSlurp API
 *
 * For documentation see [developer guide](https://www.mailslurp.com/developers). [Create an account](https://app.mailslurp.com) in the MailSlurp Dashboard to [view your API Key](https://app). For all bugs, feature requests, or help please [see support](https://www.mailslurp.com/support/).
 *
 * API version: 0.0.1-alpha
 * Contact: contact@mailslurp.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp

// Options for uploading files for attachments
type UploadAttachmentOptions struct {
	// Base64 encoded string of file contents
	Base64Contents string `json:"base64Contents,omitempty"`
	// Optional contentType for file. For instance application/pdf
	ContentType string `json:"contentType,omitempty"`
	// Optional filename to save upload with
	Filename string `json:"filename,omitempty"`
}