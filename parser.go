package mailrejectparser

import (
	"regexp"
	"strings"
)

type RejectReason struct {
	OriginalRecipient string
	Status            string
	DiagnosticCode    string
}

func IsRejected(body string) bool {
	m, _ := regexp.MatchString(`Action: failed`, body)

	return m
}

func GetReason(body []byte) RejectReason {

	var rr RejectReason

	re := regexp.MustCompile(`Original-Recipient: .*`)
	rr.OriginalRecipient = string(re.Find(body))
	if len(rr.OriginalRecipient) > 20 {
		rr.OriginalRecipient = string(rr.OriginalRecipient[20:])
	}
	// Fallback
	if len(rr.OriginalRecipient) == 0 {
		re = regexp.MustCompile(`Final-Recipient: .*`)
		rr.OriginalRecipient = string(re.Find(body))
		if len(rr.OriginalRecipient) > 17 {
			rr.OriginalRecipient = string(rr.OriginalRecipient[17:])
		}
	}
	rr.OriginalRecipient = strings.TrimSuffix(rr.OriginalRecipient, "\n")
	rr.OriginalRecipient = strings.TrimSuffix(rr.OriginalRecipient, "\r")

	re = regexp.MustCompile(`Status: .*`)
	rr.Status = string(re.Find(body))
	if len(rr.Status) > 8 {
		rr.Status = string(rr.Status[8:])
	}
	rr.Status = strings.TrimSuffix(rr.Status, "\n")
	rr.Status = strings.TrimSuffix(rr.Status, "\r")

	re = regexp.MustCompile(`(?s)Diagnostic-Code: .*($|\r?\n\r?\n)`)
	rr.DiagnosticCode = string(re.Find(body))
	if len(rr.DiagnosticCode) > 17 {
		rr.DiagnosticCode = string(rr.DiagnosticCode[17:])
	}
	// regexp does not negative look-arounds:
	cut, _ := regexp.MatchString(`\r?\n\r?\n`, rr.DiagnosticCode)
	if cut {
		re = regexp.MustCompile(`(\r?\n\r?\n).*`)
		l := re.FindIndex([]byte(rr.DiagnosticCode))
		if l != nil {
			rr.DiagnosticCode = rr.DiagnosticCode[:l[0]]
		}
	}
	rr.DiagnosticCode = strings.TrimSuffix(rr.DiagnosticCode, "\n")
	rr.DiagnosticCode = strings.TrimSuffix(rr.DiagnosticCode, "\r")

	return rr
}
