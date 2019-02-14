package s2sclient

// APIAccessor is the interface accessor object
type APIAccessor interface {
	SendIdentity(identityRequest IdentityRequest) error
	SendEvent(eventData S2SRequest, settings ...func(s *S2SRequest) error) error
	SendInstall(installData S2SRequest, settings ...func(s *S2SRequest) error) error
}
