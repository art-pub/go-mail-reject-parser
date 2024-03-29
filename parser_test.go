package mailrejectparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const BODY_REJECT = `
...


Reporting-MTA: dns; someserver.somedomain
X-Postfix-Queue-ID: 4V58tN6DfJz33tq
X-Postfix-Sender: rfc822; someone@somewhere.domain
Arrival-Date: Thu, 28 Mar 2024 18:00:28 +0100 (CET)

Final-Recipient: rfc822; someone@somewhere.domain
Original-Recipient: rfc822;someone@somewhere.domain
Action: failed
Status: 5.0.0
Remote-MTA: dns; recipient.server.domain
Diagnostic-Code: smtp; 550 Requested action not taken: mailbox unavailable
...`

const BODY_NORMAL = `
...

Successful message delivery, normal msg, etc.

...`

func TestIsRejected1(t *testing.T) {
	// given
	body := BODY_REJECT

	// when
	r := IsRejected(body)

	// then
	assert.Equal(t, true, r, "expected IsRejected to return true, but returned false!")
}

func TestIsRejected2(t *testing.T) {
	// given
	body := BODY_NORMAL

	// when
	r := IsRejected(body)

	// then
	assert.Equal(t, false, r, "expected IsRejected to return false, but returned true!")
}

func TestGetReason(t *testing.T) {
	// given
	body := BODY_REJECT

	// when
	r := GetReason([]byte(body))

	// then
	eor := "rfc822;someone@somewhere.domain"
	es := "5.0.0"
	edc := "smtp; 550 Requested action not taken: mailbox unavailable\n..."
	ead := "Thu, 28 Mar 2024 18:00:28 +0100 (CET)"
	assert.Equal(t, eor, r.OriginalRecipient, "expected "+eor+", but got "+r.OriginalRecipient+"!")
	assert.Equal(t, es, r.Status, "expected "+eor+", but got "+r.Status+"!")
	assert.Equal(t, edc, r.DiagnosticCode, "expected "+eor+", but got "+r.DiagnosticCode+"!")
	assert.Equal(t, ead, r.ArrivalDate, "expected "+ead+", but got "+r.ArrivalDate+"!")
}
