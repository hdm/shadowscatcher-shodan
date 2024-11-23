package services

type HTTP struct {
	// Web technologies that the website uses
	Components map[string]HttpComponent `json:"components,omitempty"`

	// Favicon for the website
	Favicon *HttpFavicon `json:"favicon,omitempty"`

	// HTML content of the website
	HTML string `json:"html"`

	// Numeric hash of the http.html property
	HTMLHash int `json:"html_hash"`

	// Hostname used to grab the HTML
	Host string `json:"host"`

	// Location of the final HTML response (after redirects)
	Location string `json:"location"`

	// List of redirects that were followed
	Redirects []HttpRedirect `json:"redirects"`

	// Robots.txt file for the website
	Robots *string `json:"robots"`

	// Numeric hash of the Robots property
	RobotsHash *int `json:"robots_hash"`

	// security.txt file for the website
	Securitytxt *string `json:"securitytxt"`

	// Numeric hash of the SecurityTxt property
	SecuritytxtHash *int `json:"securitytxt_hash"`

	// Server header from the HTTP response
	Server *string `json:"server"`

	// sitemap.xml for the website
	Sitemap *string `json:"sitemap"`

	// Numeric hash of the Sitemap property
	SitemapHash *int `json:"sitemap_hash"`

	// Title of the website
	Title *string `json:"title"`
	WAF   string  `json:"waf,omitempty"`

	// HTTP status
	Status int `json:"status"`

	// Hash of title
	TitleHash int `json:"title_hash"`

	// Hash of domain
	DomHash int `json:"dom_hash"`

	// Hash of headers
	HeadersHash int `json:"headers_hash"`

	// Overall hash
	Hash int `json:"hash"`

	// Hash of server header
	ServerHash int `json:"server_hash"`
}

type HttpComponent struct {
	Categories []string `json:"categories"`
	Versions   []string `json:"versions,omitempty"`
}

type HttpRedirect struct {
	Data     string `json:"data"`
	HTML     string `json:"html,omitempty"`
	Host     string `json:"host"`
	Location string `json:"location"`
}

type HttpFavicon struct {
	// Base64-encoded favicon image
	Data string `json:"data"`

	// Numeric hash of the data property
	Hash int `json:"hash"`

	// URL of the favicon
	Location string `json:"location"`

	// Mime type
	Mime string `json:"mime"`
}

type HttpScreenshot struct {
	// Base64-encoded screenshot image
	Data string `json:"data"`

	// Numeric hash of the data property
	Hash int `json:"hash"`

	// URL of the screenshot
	Location string `json:"location"`

	// Mime type
	Mime string `json:"mime"`

	// Text
	Text string `json:"text"`
}
