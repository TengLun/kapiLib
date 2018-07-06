package campaignclient

// CreateaPIA returns an accessor object. If debug flag is true, a aPIA_Fake is
// returned for debugging purposes. Otherwise, an aPIA struct is returned
func createAPIA(debug bool, a AccountAccessor) (aPIAccessor, error) {
	if debug == true {
		var apiaFake aPIA_Fake
		return apiaFake, nil
	}

	var apia aPIA
	apia.appID = a.AppID
	apia.authKey = a.AuthKey
	return apia, nil
}
