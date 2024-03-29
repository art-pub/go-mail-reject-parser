# go-mail-reject-parser
When a mail has been rejected, the body contains information about the reason. This little helper parses the body.

> [!WARNING]
**Attention:** The use is expressly at your own risk.

> [!CAUTION]
This module has been designed to return the reject values for the first rejected recipient in the mails. You can easily adjust the results of the regexp in the parser if you need to get the results for more than one recipient.


### Requirements

There are currently no requirements. [github.com/emersion/go-imap](https://github.com/emersion/go-imap) is recommended to access your IMAP accounts with go.

## Usage

Install the package:
```bash
bash$ go get github.com/art-pub/go-mail-reject-parser
```

Use the parser in your code:

```go
	...
	
	logger.Info("Reading all messages:")
	for msg := range messages {
        ...

        // fetch the body
		r := msg.GetBody(section)
		if r == nil {

			logger.Info("Server didn't return message body")
			continue
		}
		m, err := mail.ReadMessage(r)
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		body, err := io.ReadAll(m.Body)
		if err != nil {
			logger.Error(err.Error())
			continue
		}

        // check if is rejected
		if mailrejectparser.IsRejected(string(body)) {
			logger.Info("** Rejected")

            // get the reson
			r := mailrejectparser.GetReason(body)

			logger.Info("*** Status: " + r.Status)
			logger.Info("*** OriginalReceipient: " + r.OriginalRecipient)
			logger.Info("*** DiagnosticCode: " + r.DiagnosticCode)
        ...
    }
	...
```
