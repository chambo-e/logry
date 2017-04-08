package logry

// SetEnvironment sets sentry environment tag.
func SetEnvironment(env string) {
	if hook != nil {
		hook.SetEnvironment(env)
	}
}

// SetRelease sets sentry release tag.
func SetRelease(release string) {
	if hook != nil {
		hook.SetRelease(release)
	}
}
