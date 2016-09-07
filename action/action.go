package action

/* This should represent an atomic change to aptly

- State of change
- Resource to be changed
- Response
*/

import (
	"github.com/queeno/aptlify/config"
)

type ActionStruct struct {
	ResourceName	string
	ChangeType   int
	changeReason []string
}

func (a ActionStruct) isEmpty() bool {
	if a.ChangeType == 0 && a.changeReason == nil {
		return true
	}
	return false
}

func (a *ActionStruct) AddReasonToAction(reason string) {
	if reason == "" {
		a.changeReason = nil
	}
	a.changeReason = append(a.changeReason, reason)
}

func CreateActions(config *config.ConfigStruct, state *config.ConfigStruct) []ActionStruct {

	mirrorActions := createMirrorActions(config.Mirrors, state.Mirrors)
	repoActions := createRepoActions(config.Repos, state.Repos)
	gpgActions := createGpgActions(config.Gpg_keys, state.Gpg_keys)

	return append(append(mirrorActions, repoActions...), gpgActions...)
}