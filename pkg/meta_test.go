package pkg

import (
	"ck-ghost/pkg/common"
	"ck-ghost/pkg/config"
	"ck-ghost/pkg/db"
	"context"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRun(t *testing.T) {
	convey.Convey("TestRun", t, func() {
		err := db.InitDB(common.BuildDsn("120.92.181.100", "9000", "ckuser", "ckuserkingsoft2019", "compress=1&use_client_time_zone=true"))
		config.Appids = []string{
			"app_200001007",
		}
		config.InitConfig()
		convey.So(err, convey.ShouldBeNil)
		err = Run(context.Background())
		convey.So(err, convey.ShouldBeNil)
	})
}
