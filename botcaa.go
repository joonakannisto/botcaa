package main 
import (
	"io/ioutil"
	"strings"
	"fmt"
	"encoding/json"

	"github.com/go-ini/ini"
)

func main() {
	issuer := "letsencrypt.org"
	renewdir := "/etc/letsencrypt/renewal/"
	accountsdir := "/etc/letsencrypt/accounts/acme-v01.api.letsencrypt.org/directory/"
	domainconfs, err := ioutil.ReadDir(renewdir)
	if err !=nil {panic(err)}

	for _,file := range domainconfs {
		filename := file.Name()
		if strings.HasSuffix(filename,".conf") {
			domainname := strings.TrimSuffix(filename,"conf")

			cfg, err := ini.InsensitiveLoad(strings.Join([]string{renewdir,filename},""))
			if err != nil {panic(err)}
			section,err :=cfg.GetSection("renewalparams")
			if err != nil {panic(err)}
			if !section.HasKey("account") {
				fmt.Println("No account in config file, strange")
				continue
			}
			accounthash := section.Key("account").String()

			type AccountUri struct {
				Uri string `json:"uri"`
		        }
			var accountdata AccountUri
			accountsfile,err := ioutil.ReadFile(strings.Join([]string{accountsdir,accounthash,"/regr.json"},""))
			if err != nil {panic(err)}
			err = json.Unmarshal(accountsfile,&accountdata)
			if err != nil {panic(err)}
			fmt.Printf("%s\tIN\tCAA\t0\tissue\t\"%s;account-uri=%s\"\n",domainname,issuer,accountdata.Uri)
		}
	}
}
