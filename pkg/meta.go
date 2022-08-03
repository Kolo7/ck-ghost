package pkg

import (
	"ck-ghost/pkg/config"
	"ck-ghost/pkg/db"
	"context"
	"fmt"
	"github.com/dlclark/regexp2"
	"regexp"
	"strings"
	"sync"
)

type Database struct {
	Name string `db:"name"`
}

func Run(c context.Context) error {
	wg := sync.WaitGroup{}
	errC := make(chan error)
	c, cancel := context.WithCancel(c)
	defer cancel()

	// 检测错误，发现错误就通知所有goroutine结束并且关闭errorchan
	go func() {
		select {
		case <-errC:
			cancel()
			close(errC)
		case <-c.Done():
		}
	}()

	tbs, err := db.TableList(c)
	if err != nil {
		return err
	}

	mv, err := db.MvList(c)
	if err != nil {
		return err
	}

	mvmap := existTB(mv)
	tbs = filter(tbs, mvmap)

	cluster, err := db.ClusterList(c)
	if err != nil {
		return err
	}
	for _, table := range tbs {
		wg.Add(1)
		go func(c context.Context, opt Option) {
			err := createMaterialized(c, opt)
			if err != nil {
				select {
				case errC <- err:
				case <-c.Done():
				}
			}
			wg.Done()
		}(c, Option{
			Cluster:  cluster.Cluster,
			Idc:      config.IdcMap[config.Host],
			Instance: config.Host,
			Appid:    table.Database,
			Table:    table.Name,
		})
	}
	wg.Wait()
	return nil
}

func createMaterialized(c context.Context, option Option) error {
	createSql, err := Build(option)
	if err != nil {
		return err
	}
	err = db.Exec(c, createSql)
	fmt.Printf("log table %s.%s count_mv create success\n", option.Appid, option.Table)
	if err != nil {
		return err
	}
	return nil
}

func filter(tbs []*db.Table, mvMap map[string]struct{}) []*db.Table {
	newTbs := make([]*db.Table, 0)
	for _, e := range tbs {
		if !strings.HasPrefix(e.Database, "app_") ||
			!strings.HasPrefix(e.Name, "log_") ||
			!strings.HasSuffix(e.Name, "_local") ||
			e.IsTemporary ||
			strings.ToLower(e.Engine) != "replicatedmergetree" {
			continue
		}
		_, appExit := config.AppidMap[e.Database]
		_, allExit := config.AppidMap["all"]

		if !appExit && !allExit {
			continue
		}
		if _, exist := mvMap[e.Name]; exist {
			continue
		}
		newTbs = append(newTbs, e)
	}

	return newTbs
}

func existTB(tables []*db.Table) map[string]struct{} {
	mvMap := make(map[string]struct{})
	for _, table := range tables {
		reg, err := regexp.Compile("^app_[0-9]{9}_log_\\w+_count_mv$")
		if err != nil {
			panic(err)
		}
		if !reg.MatchString(table.Name) {
			continue
		}
		relTbReg, err := regexp2.Compile("log_((?!_count_mv)\\w)+", 0)
		if err != nil {
			panic(err)
		}
		match, err := relTbReg.FindStringMatch(table.Name)
		if err != nil {
			panic(err)
		}
		relTb := match.String()
		if relTb != "" {
			mvMap[relTb] = struct{}{}
		}
	}
	return mvMap
}
