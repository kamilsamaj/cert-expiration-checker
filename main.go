package main

import (
	"github.com/kamilsamaj/cert-expiration-checker/config"
	"github.com/kamilsamaj/cert-expiration-checker/internal/logger"
	"github.com/kamilsamaj/cert-expiration-checker/tlschecker"
)

func checkServerCerts() {
	for _, serverDef := range config.ServerCerts {
		daysToExpire, err := tlschecker.DaysToExpire(serverDef)

		if err != nil {
			logger.Errorf("error while checking host %v: %v", serverDef.Host, err)
		}
		if daysToExpire <= serverDef.DaysToWarn {
			logger.Errorf("certificate on host %v will expire in %d days", serverDef.Host, daysToExpire)
		}
	}
}

func main() {
	checkServerCerts()
}
