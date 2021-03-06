// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package showwallet

import (
	"gopkg.in/macaroon-bakery.v2-unstable/httpbakery"
)

var (
	NewWalletAPIClient = &newWalletAPIClient
	NewJujuclientStore = &newJujuclientStore
)

// WalletAPIClientFnc returns a function that returns the provided walletAPIClient
// and can be used to patch the NewWalletAPIClient variable for tests.
func WalletAPIClientFnc(api walletAPIClient) func(string, *httpbakery.Client) (walletAPIClient, error) {
	return func(string, *httpbakery.Client) (walletAPIClient, error) {
		return api, nil
	}
}
