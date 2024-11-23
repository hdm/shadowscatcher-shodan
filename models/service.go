package models

import (
	"fmt"

	"github.com/shadowscatcher/shodan/models/services"
)

// Service from search results
type Service struct {
	HostInfo
	Location Location `json:"location"`

	// Contains the banner information for the service
	Data string `json:"data"`

	// Information about the cloud provider for this IP
	Cloud Cloud `json:"cloud,omitempty"`

	// The IP address of the host as an integer
	IP *int `json:"ip,omitempty"`

	// The IPv6 address of the host as a string. If this is present then the "IP" and "IPstr" fields wont be.
	IPv6 *string `json:"ipv6,omitempty"`

	// The port number that the service is operating on
	Port int `json:"port"`

	// The timestamp for when the banner was fetched from the device in the UTC timezone.
	// Example: "2014-01-15T05:49:56.283713"
	Timestamp string `json:"timestamp"`

	// Numeric hash of the data property
	Hash int `json:"hash"`

	// An array of strings containing the top-level domains for the hostnames of the device. This is a utility property
	// in case you want to filter by TLD instead of subdomain. It is smart enough to handle global TLDs with several
	// dots in the domain (ex. "co.uk")
	Domains []string `json:"domains"`

	// The network link type. Possible values are: "Ethernet or modem", "generic tunnel or VPN", "DSL", "IPIP or SIT",
	// "SLIP", "IPSec or GRE", "VLAN", "jumbo Ethernet", "Google", "GIF", "PPTP", "loopback", "AX.25 radio modem".
	Link *string `json:"link,omitempty"`

	// Contains experimental and supplemental data for the service. This can include the SSL certificate, robots.txt
	// and other raw information that hasn't yet been formalized into the Banner Specification.
	Opts map[string]interface{} `json:"opts"`

	// Uptime of the IP (in minutes)
	Uptime *int `json:"uptime,omitempty"`

	// Either "udp" or "tcp" to indicate which IP transport protocol was used to fetch the information
	Transport string `json:"transport"`

	// Name of the software running the service
	// In rare occasions can be number. Use ProductString() to get string value of Product
	Product interface{} `json:"product,omitempty"`

	// Version of the software
	// In rare occasions can be number. Use VersionString() to get string value of Version
	Version interface{} `json:"version,omitempty"`

	// Common platform enumeration
	// Can be slice of strings or single string. Use CpeList() to get a slice of strings value
	CPE interface{} `json:"cpe,omitempty"`

	// The title of the website as extracted from the HTML source
	Title *string `json:"title,omitempty"`

	// The type of device (webcam, router, etc.)
	DeviceType *string `json:"devicetype,omitempty"`

	// Miscellaneous information that was extracted about the product
	Info *string `json:"info,omitempty"`

	// The _shodan property contains information about how the banner was generated. It doesn’t store any
	// data about the port/service itself.
	// Availability: All banners
	Shodan Shodan `json:"_shodan"`

	// The vulns property contains information about vulnerabilities that may exist in the service represented
	// by the banner. In general, the Shodan crawlers don’t perform vulnerability testing as a result the
	// vulnerabilities stored in vulns are inferred from the banner and haven’t been verified.
	// Availability: Banners where the software/version has been identified and there exist known CVEs for it
	Vulns map[string]Vulnerability `json:"vulns,omitempty"`

	// Common platform enumeration v2.3
	CPE23 []string `json:"cpe23"`

	// MAC address
	MAC map[string]MAC `json:"mac"`

	// Various tags
	Tags []string `json:"tags"`

	// Hostnames
	Hostnames []string `json:"hostnames"`

	// Operating system
	OS string `json:"os"`

	// Organization
	Organization string `json:"org"`

	// ISP
	ISP string `json:"isp"`

	// ASN
	ASN string `json:"asn"`

	// HTML
	HTML string `json:"html"`

	// Banner
	Banner string `json:"banner"`

	// Platform
	Platform string `json:"platform"`

	// Device
	Device string `json:"device"`

	// Screenshot for the website
	Screenshot *services.HttpScreenshot `json:"screenshot,omitempty"`

	// Availability: Services that require SSL (ex. HTTPS) or support upgrading a connection to SSL/TLS
	// (ex. POP3 with STARTTLS)
	SSL *SSL `json:"ssl,omitempty"`

	// Availability: Cassandra database services that allow connections to the client Thrift port (default: 9160/ tcp)
	Cassandra *services.Cassandra `json:"cassandra,omitempty"`

	// Availability: Services running the IBM DB2 DRDA protocol
	DB2 *services.DB2 `json:"ibm_db2,omitempty"`

	// Availability: DNS servers that support either UDP or TCP (typically on port 53)
	DNS *services.DNS `json:"dns,omitempty"`

	// Availability: Docker services that allow remote connections and don’t have authentication enabled
	Docker *services.Docker `json:"docker,omitempty"`

	// Availability: The elastic property is available in banners that are identified as belonging
	// to an Elastic service
	Elastic *services.Elastic `json:"elastic,omitempty"`

	// The etcd service provides a distributed key/value store used by projects such as Kubernetes.
	// Availability: Ports that are running the etcd service
	Etcd *services.Etcd `json:"etcd,omitempty"`

	// Availability: Devices that complete a handshake in either TCP or UDP for the industrial Ethernet/IP protocol
	EthernetIP *services.EthernetIP `json:"ethernetip,omitempty"`

	// Availability: FTP services running on the default port 21/TCP.
	// If the FTP service supports STARTTLS then the starttls tag will be added to the list of tags on the banner
	// and it will also have a top-level ssl property which contains the certificate, SSL testing results and more
	FTP *services.FTP `json:"ftp,omitempty"`

	// Availability: Devices running Apache Hive servers on any port that Shodan crawls
	Hive *services.Hive `json:"hive,omitempty"`

	// Availability: The banner was generated by a HTTP module (http, https, http-simple-new, https-simple-new) and
	// successfully completed a HTTP handshake
	HTTP *services.HTTP `json:"http,omitempty"`

	// Availability: VPN services that use the ISAKMP protocol (such as IKE)
	ISAKMP *services.ISAKMP `json:"isakmp,omitempty"`

	// Availability: Lantronix devices that are running the configuration service
	Lantronix *services.Lantronix `json:"lantronix,omitempty"`

	// Availability: If the Monero RPC service is enabled and accepting remote connections.
	// Most results are on port 18081, but it can also be available on other ports
	Monero *services.Monero `json:"monero,omitempty"`

	// Availability: MongoDB services that support the binary protocol to interact with the database
	MongoDB *services.Mongo `json:"mongodb,omitempty"`

	// Availability: MQTT services that allow remote connections
	MQTT *services.MQTT `json:"mqtt,omitempty"`

	// Availability: Services that run on port 137 and complete a NetBIOS handshake
	Netbios *services.Netbios `json:"netbios,omitempty"`

	// Availability: NTP daemons supporting at least version 1 or version 2
	NTP *services.NTP `json:"ntp,omitempty"`

	// Availability: Redis services running on the default port 6379/TCP
	Redis *services.Redis `json:"redis,omitempty"`

	// Availability: Services on port 520 that successfully respond to a RIP request
	RIP *services.RIP `json:"rip,omitempty"`

	Rsync *services.Rsync `json:"rsync,omitempty"`

	// Availability: Services that run on port 445 and support either SMBv1 or SMBv2
	SMB *services.SMB `json:"smb,omitempty"`

	// Availability: Any banner generated by the snmp module (typically on 161/UDP)
	SNMP *services.SNMP `json:"snmp,omitempty"`

	// Availability: Any service banner where the initial response starts with “SSH” and subsequently completes a SSH
	// handshake
	SSH *services.SSH `json:"ssh,omitempty"`

	// Availability: Devices running the VertX/Edge door controllers
	Vertx *services.Vertx `json:"vertx,omitempty"`

	// Availability: Devices running the Minecraft game server
	Minecraft *services.Minecraft `json:"minecraft"`

	// Availability: Devices running InfluxDB time-series database
	InfluxDb *services.InfluxDb `json:"influx_db"`

	// Availability: Devices running CoAP IoT protocol service
	CoAP *services.CoAP `json:"coap"`

	// Availability: Devices running IPMI services
	IPMI *services.IPMI `json:"ipmi"`

	// Availability: Devices running RDP services
	RDPEncryption *services.RDPEncryption `json:"rdp_encryption"`

	// Availability: Devices running AMQP services
	AMQP *services.AMQP `json:"amqp"`

	VNC               *services.VNC               `json:"vnc"`
	CiscoAnyConnect   *services.AnyConnect        `json:"cisco_anyconnect"`
	Telnet            *services.Telnet            `json:"telnet"`
	MySQL             *services.MySQL             `json:"mysql"`
	MySQLX            *services.MySQLX            `json:"mysqlx"`
	DAV               *services.DAV               `json:"dav"`
	Hikivision        *services.Hikvision         `json:"hikvision"`
	Stun              *services.Stun              `json:"stun"`
	NTLM              *services.NTLM              `json:"ntlm"`
	MikrotikWinbox    *services.MikrotikWinbox    `json:"mikrotik_winbox"`
	MikrotikRouterOS  *services.MikrotikRouterOS  `json:"mikrotik_routeros"`
	Checkpoint        *services.Checkpoint        `json:"checkpoint"`
	QNAP              *services.QNAP              `json:"qnap"`
	Dahua             *services.Dahua             `json:"dahua"`
	DahuaDVRWeb       *services.DahuaDVRWeb       `json:"dahua_dvr_web"`
	NodeExporter      *services.NodeExporter      `json:"node_exporter"`
	EthereumRPC       *services.EthereumRPC       `json:"ethereum_rpc"`
	EthereumP2P       *services.EthereumP2P       `json:"ethereum_p2p"`
	Fortinet          *services.Fortinet          `json:"fortinet"`
	SynologyDSM       *services.SynologyDSM       `json:"synology_dsm"`
	SynologySRM       *services.SynologySRM       `json:"synology_srm"`
	PPTP              *services.PPTP              `json:"pptp"`
	BGP               *services.BGP               `json:"bgp"`
	Plex              *services.Plex              `json:"plex"`
	TilginHomeGateway *services.TilginHomeGateway `json:"tilginAB_home_gateway"`
	MDNS              *services.MDNS              `json:"mdns"`
	IPPCups           *services.IPPCups           `json:"ipp_cups"`
	DockerRegistry    *services.DockerRegistry    `json:"docker_registry"`
	MicrosoftExchange *services.MicrosoftExchange `json:"microsoft_exchange"`
	LDAP              *services.LDAP              `json:"ldap"`
	UPNP              *services.UPNP              `json:"upnp"`
	AFP               *services.AFP               `json:"afp"`
	Airplay           *services.Airplay           `json:"airplay"`
	ClickHouse        *services.ClickHouse        `json:"clickhouse"`
	Consul            *services.Consul            `json:"consul"`
	MSRPC             *services.MSRPC             `json:"msrpc"`
	BACnet            *services.BACnet            `json:"bacnet"`
	CouchDB           *services.CouchDB           `json:"couchdb"`
	DraytekVigor      *services.DraytekVigor      `json:"draytek_vigor"`
	EPMD              *services.EPMD              `json:"epmd"`
	HPiLO             *services.HPiLO             `json:"hp_ilo"`
	HPPrinterWeb      *services.HPPrinterWeb      `json:"hp_printer_embedded_web_server"`
	HueBridge         *services.HueBridge         `json:"philips_hue"`
	RealPort          *services.RealPort          `json:"realport"`
	ADB               *services.ADB               `json:"android_debug_bridge"`
	GLiNET            *services.GLiNET            `json:"gl_inet"`
	HomeAssistant     *services.HomeAssistant     `json:"home_assistant"`
	InfluxDB          *services.InfluxDb          `json:"influxdb"`
	Bitcoin           *services.Bitcoin           `json:"bitcoin"`
	Chromecast        *services.Chromecast        `json:"chromecast"`
	Domoticz          *services.Domoticz          `json:"domoticz"`
	IPCamera          *services.IPCamera          `json:"ip_camera"`
	ISCSI             *services.ISCSI             `json:"iscsi"`
}

// region public methods

// Product can be a string or number (in rare occasions).
func (s *Service) ProductString() string {
	if s.Product == nil {
		return ""
	}
	return fmt.Sprint(s.Product)
}

// CPE can be a string or list of strings.
func (s *Service) CpeList() []string {
	// TODO: Also parse CPE23

	if s.CPE == nil {
		return []string{}
	}

	switch s.CPE.(type) {
	case string:
		return []string{s.CPE.(string)}
	case []string:
		return s.CPE.([]string)
	default:
		return []string{}
	}
}

// Version can be a string or number (in rare occasions)
func (s *Service) VersionString() string {
	if s.Version == nil {
		return ""
	}
	return fmt.Sprint(s.Version)
}

// get string IP, no matter v4 or v6
func (s *Service) IPString() string {
	// if ipv6 is present, IP and IPstr won't be set.
	if s.IP != nil {
		return s.IPstr
	} else {
		return *s.IPv6
	}
}

// IP:port pair as string
func (s *Service) IpAndPort() string {
	return fmt.Sprintf("%s:%d", s.IPString(), s.Port)
}

// endregion
