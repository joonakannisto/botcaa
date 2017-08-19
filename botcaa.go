package main 
import (
	"io/ioutil"
	"strings"
	"fmt"

	"github.com/go-ini/ini"
)

func main() {
	renewdir := "/etc/letsencrypt/renewal/"
	accountsdir := "/etc/letsencrypt/accounts/acme-v01.api.letsencrypt.org/directory/"
	
	domainconfs, err := ioutil.ReadDir(renewdir)
	if err !=nil {panic(err)}
	var data map[string]string
	data = make(map[string]string)
	for _,file := range domainconfs {
		filename := file.Name()
		if strings.HasSuffix(filename,".conf") {
			domainname := strings.TrimSuffix(filename,"conf")
			data["domain"] = domainname
			cfg, err := ini.InsensitiveLoad(strings.Join([]string{renewdir,filename},""))
			if err != nil {panic(err)}
			section,err :=cfg.GetSection("renewalparams")
			if err != nil {panic(err)}
			if !section.HasKey("account") {
			fmt.Println("No account in config file, strange")
			continue	
			}
			accounthash := section.Key("account").String()
								
		}
	}
}
