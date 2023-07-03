// Package activity 生成活动配置
package activity

const Cofing = `package {{.pkg}}

import (
	"casino/module/business/dao/db"
	"casino/module/business/dao/field"
	"casino/module/business/dao/gm"
	"casino/module/business/service/activity/base/static"
	"casino/module/business/service/backend/api"
	serviceCfg "casino/module/business/service/config"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

type Config struct {
	serviceCfg.DefaultConfig
}

func (c *Config) Name() string {
	return serviceCfg.GenerateName(static.Act{{.actID}}, c)
}

func (c *Config) MemObj() interface{} {
	mu.RLock()
	defer mu.RUnlock()
	return cfg
}

var (
	cfg *Config
	mu  sync.RWMutex
)

func (c *Config) Init() error {
	var (
			err error
			cb  gm.CacheBackup //数据库配置
		)
	cfg = new(Config)
	err = cb.GetCfgByKeyword(nil, db.KeyValue{field.EventName: EventName, field.Keyword: static.Act{{.actID}}})
	if errors.Is(err, sql.ErrNoRows) {
		return sql.ErrNoRows
	}
	if err != nil {
		return fmt.Errorf("[{{.pkg}}.Init] GetCfgByKeyword Scan Error: %s", err.Error())
	}
	err = json.Unmarshal([]byte(cb.Config), cfg)
	if err != nil {
		return fmt.Errorf("[{{.pkg}}.Init] Unmarshal  Error: %s", err.Error())
	}
	return nil
}

func (c *Config) GetMemoryConfig(_ *http.Request) (error, string) {
	cfgByte, _ := json.Marshal(c.MemObj())
	return nil, string(cfgByte)
}

func (c *Config) UpdateConfig(r *http.Request) error {
	payload := r.PostForm.Get(api.Payload)
	var cnf *Config
	if err := json.Unmarshal([]byte(payload), &cnf); err != nil {
		return err
	}
	if err := serviceCfg.Inspect(cnf); err != nil {
		return err
	}
	_ = SetCfg(cnf)
	return nil
}

func GetCfg() *Config {
	mu.RLock()
	defer mu.RUnlock()
	return cfg
}

func SetCfg(c *Config) error {
	if c == nil {
		return errors.New("config is nil")
	}
	mu.Lock()
	defer mu.Unlock()
	cfg = c
	return nil
}
`
const TaskConfig = `package {{.pkg}}

import (
	"casino/module/business/dao/db"
	"casino/module/business/dao/field"
	"casino/module/business/dao/gm"
	"casino/module/business/service/activity/base"
	"casino/module/business/service/activity/base/static"
	"casino/module/business/service/backend/api"
	serviceCfg "casino/module/business/service/config"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

type Config struct {
	serviceCfg.DefaultConfig
}

func (c *Config) Name() string {
	return serviceCfg.GenerateName(static.Act{{.actID}}, c)
}

func (c *Config) MemObj() interface{} {
	mu.RLock()
	defer mu.RUnlock()
	return cfg
}

var (
	cfg *Config
	mu  sync.RWMutex
)

func (c *Config) Init() error {
	var (
			err error
			cb  gm.CacheBackup //数据库配置
		)
	base.Register(static.Act{{.actID}}, base.TaskFuncTemplate()) //将任务方法注册到任务总表中
	cfg = new(Config)
	err = cb.GetCfgByKeyword(nil, db.KeyValue{field.EventName: EventName, field.Keyword: static.Act{{.actID}}})
	if errors.Is(err, sql.ErrNoRows) {
		return sql.ErrNoRows
	}
	if err != nil {
		return fmt.Errorf("[{{.pkg}}.Init] 数据库查询错误: %s", err.Error())
	}
	err = json.Unmarshal([]byte(cb.Config), cfg)
	if err != nil {
		return fmt.Errorf("[{{.pkg}}.Init] 反序列化错误: %s", err.Error())
	}
	return nil
}

func (c *Config) GetMemoryConfig(_ *http.Request) (error, string) {
	cfgByte, _ := json.Marshal(c.MemObj())
	return nil, string(cfgByte)
}

func (c *Config) UpdateConfig(r *http.Request) error {
	payload := r.PostForm.Get(api.Payload)
	var cnf *Config
	if err := json.Unmarshal([]byte(payload), &cnf); err != nil {
		return err
	}
	if err := serviceCfg.Inspect(cnf); err != nil {
		return err
	}
	return SetCfg(cnf)
}

func GetCfg() *Config {
	mu.RLock()
	defer mu.RUnlock()
	return cfg
}

func SetCfg(c *Config) error {
	if c == nil {
		return errors.New("配置是空的")
	}
	mu.Lock()
	defer mu.Unlock()
	cfg = c
	return nil
}
`
