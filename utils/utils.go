package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"strings"
)

type Util struct {
	Config *ini.File
}
func (u *Util) LoadConfig(conf string) (err error) {
	u.Config, err = ini.Load(conf)
	if err!=nil {
		u.Config = nil
		return
	}
	return
}

func (u *Util) GetValue(sec, key string) (string) {
	section, err := u.Config.GetSection(sec)
	if err != nil {
		fmt.Printf("Section not found:%s (%v)\n",sec,err)
		return ""
	}
	key2, err := section.GetKey(key)
	if err!=nil {
		fmt.Printf("Key not found:%s (%v)\n", key, err)
		return ""
	}
	return strings.TrimSpace(key2.Value())
}

func (u *Util) GetValue2(sec, key string,silence bool) (string) {
	section, err := u.Config.GetSection(sec)
	if err != nil {
		if silence {
			return ""
		}
		fmt.Printf("Section not found:%s (%v)\n",sec,err)
		return ""
	}
	key2, err := section.GetKey(key)
	if err!=nil {
		if silence {
			return ""
		}
		fmt.Printf("Key not found:%s (%v)\n", key, err)
		return ""
	}
	return strings.TrimSpace(key2.Value())
}


func (u *Util) ReadFile(fpath string) ([]byte, error) {
	contents, err := ioutil.ReadFile(fpath)
	return contents,err
}

func (u *Util) WriteFile(fpath string, contents []byte) ( error) {
	err := ioutil.WriteFile(fpath,contents,0644)
	return err
}