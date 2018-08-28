package userprofile

import (
	"github.com/dailyhunt/feed/user_profile/pb"
)

type RawUserProfile struct {
	raw  map[string]*userprofile.CFeaturewise
	pref map[string]string
}
