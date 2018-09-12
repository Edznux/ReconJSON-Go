package servicedescriptors

type HttpCSP struct {
	Type                    string `json:"type"`     // MUST be the type of object. In this case, ServiceDescriptor.
	Name                    string `json:"name"`     // MUST be the name of ServiceDescriptor. In this case, httpCsp.
	Location                string `json:"location"` //  MUST be the location where the CSP was found. Must be either http header or meta tag.
	BaseUri                 string `json:"BaseUri,omitempty"`
	BlockAllMixedContent    string `json:"BlockAllMixedContent,omitempty"`
	ChildSrc                string `json:"ChildSrc,omitempty"`
	ConnectSrc              string `json:"connectSrc,omitempty"`
	DefaultSrc              string `json:"defaultSrc,omitempty"`
	FontSrc                 string `json:"fontSrc,omitempty"`
	FormAction              string `json:"formAction,omitempty"`
	FrameAncestors          string `json:"frameAncestors,omitempty"`
	FrameSrc                string `json:"frameSrc,omitempty"`
	ImgSrc                  string `json:"imgSrc,omitempty"`
	ManifestSrc             string `json:"manifestSrc,omitempty"`
	MediaSrc                string `json:"mediaSrc,omitempty"`
	ObjectSrc               string `json:"objectSrc,omitempty"`
	PluginTypes             string `json:"pluginTypes,omitempty"`
	Referrer                string `json:"referrer,omitempty"`
	ReportUri               string `json:"reportUri,omitempty"`
	RequireSriFor           string `json:"requireSriFor,omitempty"`
	Sandbox                 string `json:"sandbox,omitempty"`
	ScriptSrc               string `json:"scriptSrc,omitempty"`
	StyleSrc                string `json:"styleSrc,omitempty"`
	UpgradeInsecureRequests string `json:"upgradeInsecureRequests,omitempty"`
	WorkerSrc               string `json:"workerSrc,omitempty"`
	reportTo                string `json:"reportTo,omitempty"`
}
