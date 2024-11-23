package services

type DAV struct {
	PublicOptions  []string `json:"public_options"`
	AllowedMethods []string `json:"allowed_methods"`
	WebdavType     string   `json:"webdav_type"`
}
