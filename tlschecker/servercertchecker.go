package tlschecker

import (
	"github.com/kamilsamaj/cert-expiration-checker/config"
)

//DaysToExpire fetches a certificate from a host and return the number of days to expire
func DaysToExpire(serverDetails config.ServerCert) (int, error) {
	return 0, nil
}
