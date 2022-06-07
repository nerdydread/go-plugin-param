package main

import (
	"github.com/TykTechnologies/tyk/log"
	"net/http"
    //"fmt"
	//"github.com/hashicorp/vault/api"
	//"os"
)

var logger = log.Get()
//var vault_addr = os.Getenv("VAULT_ADDR")

//Add Params
func AddParams(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
    query.Add("api_key", "secret")
	r.URL.RawQuery = query.Encode()
    logger.Info("Query Param Added.")
    logger.Info("URL String:", r.URL.String())
   // getSecret()
}

/*func getSecret(){
    config := &api.Config{
		Address: vault_addr,
	}
    client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
    secret, err := client.Logical().Read("secret/tokens/dev-token")
    if err != nil {
		fmt.Println(err)
		return
	}
    logger.Info("Vault Secret:", secret.Data)
}*/

// called once plugin is loaded, this is where we put all initialization work for plugin
// i.e. setting exported functions, setting up connection pool to storage and etc.
func init() {
	logger.Info("Initialising Example Go Plugin")
}

func main() {}
