/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: d1659dc1567a5b62f65d0612107a50aace03e085
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
// EmailAnalysis Analysis result for email. Each verdict property is a string PASS|FAIL|GRAY or dynamic error message
type EmailAnalysis struct {
	// Verdict of DomainKeys Identified Mail analysis
	DkimVerdict string `json:"dkimVerdict,omitempty"`
	// Verdict of Domain-based Message Authentication Reporting and Conformance analysis
	DmarcVerdict string `json:"dmarcVerdict,omitempty"`
	// Verdict of spam ranking analysis
	SpamVerdict string `json:"spamVerdict,omitempty"`
	// Verdict of Send Policy Framework record spoofing analysis
	SpfVerdict string `json:"spfVerdict,omitempty"`
	// Verdict of virus scan analysis
	VirusVerdict string `json:"virusVerdict,omitempty"`
}
