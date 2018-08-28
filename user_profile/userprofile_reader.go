package userprofile

import (
	"fmt"
	fr "github.com/dailyhunt/feed/redis"
	pb "github.com/dailyhunt/feed/user_profile/pb"
	"github.com/go-redis/redis"
	"github.com/gogo/protobuf/proto"
	"sync"
)

type Options struct {
	redisHosts     string
	prefRedisHosts string
}

type ProfileReader struct {
	rawClient  *redis.ClusterClient
	prefClient *redis.ClusterClient
}

func (pr *ProfileReader) GetProfile(cid string, fields ...string) *RawUserProfile {

	profile := &RawUserProfile{
		raw: make(map[string]*pb.CFeaturewise),
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go pr.getRawProfile(&wg, profile, "N4_"+cid, fields...)
	go pr.getPrefProfile(&wg, profile, cid)

	wg.Wait()

	return profile

}

func (pr *ProfileReader) getRawProfile(wg *sync.WaitGroup, profile *RawUserProfile, key string, fields ...string) {
	defer wg.Done()
	fmt.Println("Getting profile for  user ", key, " fields = ", fields)
	values := pr.rawClient.HMGet(key, fields...).Val()
	for i, field := range fields {
		var fw pb.CFeaturewise
		v := values[i].(string)
		// todo : Error handling
		if err := proto.Unmarshal([]byte(v), &fw); err != nil {
			// TODO : Add LOG and increment counter
		}
		profile.raw[field] = &fw
		fmt.Println("Read prof")
	}
}

// TODO: LZ4 and message pack integration
func (pr *ProfileReader) getPrefProfile(wg *sync.WaitGroup, profile *RawUserProfile, cid string) {
	defer wg.Done()
	redisKey := fmt.Sprintf("NRT_V3~%s~pref", cid)
	profile.pref = pr.prefClient.HGetAll(redisKey).Val()

}

func NewUserProfileReader(opts Options) *ProfileReader {
	client := fr.NewRedisClusterClient(fr.NewRedisClusterConfig(opts.redisHosts))
	prefClient := fr.NewRedisClusterClient(fr.NewRedisClusterConfig(opts.prefRedisHosts))
	return &ProfileReader{
		rawClient:  client,
		prefClient: prefClient,
	}
}
