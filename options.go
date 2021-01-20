package recaptcha

type Option func(rc *Recaptcha)

func WithServerName(serverName string) Option {
	return func(rc *Recaptcha) {
		rc.serverName = serverName
	}
}
