package servicedescriptors

type HttpUrl struct {
	Type              string            `json:"type"`                        // "httpUrl"
	Name              string            `json:"name"`                        // MUST be a descriptive name of what data this ServiceDescriptor is describing
	Path              string            `json:"path"`                        // MUST be the path requested where path is defined by RFC 3986 section 3.3
	Verb              string            `json:"verb"`                        // MUST be the the HTTP Verb used in the request to this path.
	ParamsContentType string            `json:"paramsContentType,omitempty"` // MUST be the format in which the params field is stored. The value of this field MUST be one of the following: application/x-www-form-urlencoded, multipart/form-data, application/json, or application/xml.
	Params            string            `json:"params,omitempty"`            // MUST be any HTTP parameters sent along with this request. The format of the data stored in this field is described by paramsContentType.
	Screenshot        string            `json:"screenshot,omitempty"`        // MUST be the path to an image file that represents the contents of this endpoint
	Code              int               `json:"code,omitempty"`              // MUST be the HTTP Status code returned by either a GET request (by default) or the HTTP Verb described in the verb attribute. This attribute MUST be an integer.
	ContentType       string            `json:"contentType,omitempty"`       // MUST be the content-type returned by either a GET request (by default) or the HTTP Verb described in the verb attribute
	Length            int               `json:"length,omitempty"`            // MUST be the length of the request returned by either a GET request (by default) or the HTTP Verb described in the verb attribute. This attribute MUST be an integer.
	Headers           map[string]string `json:"headers,omitempty"`           // MUST be a map of relevant headers not included above
	File              string            `json:"file,omitempty"`              // MUST be the file requested by the path above
	FileExt           string            `json:"fileExt,omitempty"`           // MUST be the file extention of the file requested by the path above. This MUST not contain the . (dot). For example, a valid entry would be php not .php
	Hash              string            `json:"hash,omitempty"`              // MUST be a map of hashtype to hash of this endpoint's content. This can be used to determine if the contents have changed or for forensic analysis
}
