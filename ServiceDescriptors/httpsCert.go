package servicedescriptors

type HttpsCert struct {
	Type         string   `json:"type"`         // "ServiceDescriptor" MUST be the type of object. In this case, ServiceDescriptor.
	Name         string   `json:"name"`         // MUST be the name of ServiceDescriptor. In this case, httpsCert.
	CertType     string   `json:"certType"`     // MUST be the type of X.509 certificate: either cert or precert.
	Data         string   `json:"data"`         // The raw X.509 (pre-)certificate, encoded in base64.
	DnsNames     []string `json:"dnsNames"`     // MUST be a list of strings representing the DNS identifiers for the cert from both the Subject CN and the DNS SANs.
	Issuer       string   `json:"issuer"`       // MUST be the distinguished name of the certificate's issuer.
	Sha256       string   `json:"sha256"`       // MUST be the hex-encoded SHA-256 digest of the raw X.509 (pre-)certificate.
	PubkeySha256 string   `json:"pubkeySha256"` // MUST be the hex-encoded SHA-256 digest of the Subject Public Key Info.
	NotBefore    string   `json:"notBefore"`    // MUST be the not before date, in RFC3339 format (e.g. 2016-06-16T00:00:00-00:00) .
	NotAfter     string   `json:"notAfter"`     // MUST be the not after date, in RFC3339 format (e.g. 2016-06-16T00:00:00-00:00).
	Logs         Log
}

type Log struct {
	Id        string `json:"id"`        // The ID of the Certificate Transparency log, encoded in base64.
	Index     int    `json:"index"`     // The 0-based index of the (pre-)certificate's entry in the log.
	Timestamp string `json:"timestamp"` // The time at which the (pre-)certificate was submitted to this log, in RFC3339 format (e.g. 2017-05-04T13:39:21.071-00:00
}
