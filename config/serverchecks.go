package config

//ServerCert defines a host and certificate details
type ServerCert struct {
	Host       string
	DaysToWarn int
}

//ServerCerts provides with a list of all hosts to check
var ServerCerts = []ServerCert{
	{
		Host:       "www.seznam.cz",
		DaysToWarn: 14,
	},
}
