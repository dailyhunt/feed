package userprofile

import (
	"fmt"
	"github.com/dailyhunt/feed/util"
	"testing"
)

func TestGetProfile(t *testing.T) {

	callRedis()

}

func callRedis() {
	cid := "952256114"
	r := NewUserProfileReader(Options{"192.168.25.64:7000", "192.168.10.109:6000"})
	stop := util.StartTimer("Get-Profile")
	profile := r.GetProfile(cid, "news~hi")
	stop()

	fw := profile.raw["news~hi"]
	fmt.Println(fw.GetIg()["hi_x_128_0"])

	fmt.Println(profile.pref["news~top_s"])

}
