package config

import (
    "fmt"
    "flag"
)

var (
    securityConfig *SecurityConfig
)

type SecurityConfig struct {
    EncryptCost int
    CookieName string
    CookieDomain string
    CookiePath string
}

func (c *SecurityConfig) String() string {
    return fmt.Sprintf("SecurityConfig[EncryptCost=%v CookieName=%v CookieDomain=%v CookiePath=%v]", c.EncryptCost, c.CookieName, c.CookieDomain, c.CookiePath)
}

func BindSecurityConfiguration() *SecurityConfig {
    if securityConfig == nil {
        securityConfig = &SecurityConfig{}
        flag.IntVar(&securityConfig.EncryptCost, "encrypt_cost", 10, "Bcrypt encrypt cost")
        flag.StringVar(&securityConfig.CookieName, "cookie_name", "FIVECOLORS_ID", "Session Cookie name")
        flag.StringVar(&securityConfig.CookieDomain, "cookie_domain", "darkside.e-pedion.com", "Session Cookie Domain")
        flag.StringVar(&securityConfig.CookiePath, "cookie_path", "/", "Session Cookie Path")
    }
    return securityConfig
}

