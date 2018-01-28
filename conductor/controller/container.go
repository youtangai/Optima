package controller

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"bytes"
)

const (
	AUTH_PATH                = "/auth/tokens?nocatalog"
	ZUN_HOST                 = "http://10.6.18.162:9517"
	ZUN_PATH                 = "/v1/containers/"
	ZUN_DISABLE_SERVICE_PATH = "/v1/services/disable"
	ZUN_ENABLE_SERVICE_PATH  = "/v1/services/enable"
	ZUN_SERVICE_BINARY       = "zun-compute"
	CHECKPOINT_PORT          = "62072"
	CHECKPOINT_PATH          = "/checkpoint"
	RESTORE_PORT             = "62073"
	RESTORE_PATH             = "/restore"
	DISABLE_REASON           = "optima-leave"
)

func createContainer(imageName string) (string, error) {
	//openstackにコンテナ作成を依頼
	token, err := authKeyStone()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	jsonStr := `
	{
		"image":"` + imageName + `",
		"name":"` + randomString() + `"
	}
	`

	req, err := http.NewRequest(
		"POST",
		ZUN_HOST+ZUN_PATH,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	//コンテントタイプをせってい
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer resp.Body.Close()
	bytebody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	var containerData interface{}
	err = json.Unmarshal(bytebody, &containerData)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	log.Printf("containerData = %+v", containerData)
	uuid := containerData.(map[string]interface{})["uuid"].(string)

	return uuid, nil
}

func checkpointContainer(containerID, hostName string) (string, error) {
	jsonStr := `
	{
		"container_id":"` + containerID + `"
	}
	`

	req, err := http.NewRequest(
		"POST",
		"http://"+hostName+":"+CHECKPOINT_PORT+CHECKPOINT_PATH,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	//コンテントタイプをせってい
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer resp.Body.Close()
	bytebody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	var chkDirPathJSON interface{}
	err = json.Unmarshal(bytebody, &chkDirPathJSON)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	chkDirPath := chkDirPathJSON.(map[string]interface{})["chk_dir_path"].(string)

	return chkDirPath, nil
}

//コンテナ削除
func deleteContainer(uuid string) error {
	//delete container
	token, err := authKeyStone()
	if err != nil {
		log.Fatal(err)
		return err
	}

	req, err := http.NewRequest(
		"DELETE",
		ZUN_HOST+ZUN_PATH+uuid,
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//token 設定
	req.Header.Set("X-Auth-Token", token)

	//クエリパラム設定
	query := req.URL.Query()
	query.Add("force", "True")
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	requestID := resp.Header.Get("X-Openstack-Request-Id")
	log.Println("delete container request id = " + requestID)
	return nil
}

func restoreContainer(containerID, restoreDir, hostName string) error {
	jsonStr := `
	{
		"container_id":"` + containerID + `",
		"restore_dir":"` + restoreDir + `"
	}
	`

	req, err := http.NewRequest(
		"POST",
		"http://"+hostName+":"+RESTORE_PORT+RESTORE_PATH,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//コンテントタイプをせってい
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()
	byteBody, err := ioutil.ReadAll(resp.Body)
	var messageJSON interface{}
	err = json.Unmarshal(byteBody, &messageJSON)
	if err != nil {
		log.Fatal(err)
		return err
	}
	message := messageJSON.(map[string]interface{})["message"].(string)
	log.Println("restore message " + message)

	return nil
}

//ホストを無効化する処理
func disableHost(hostname string) error {
	token, err := authKeyStone()
	if err != nil {
		log.Fatal(err)
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		ZUN_HOST+ZUN_DISABLE_SERVICE_PATH,
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//token 設定
	req.Header.Set("X-Auth-Token", token)

	//クエリパラム設定
	query := req.URL.Query()
	query.Add("binary", ZUN_SERVICE_BINARY)
	query.Add("host", hostname)
	query.Add("disabled_reason", DISABLE_REASON)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	log.Println("zunservice: disable hostname = " + hostname)
	return nil
}

//ホストを有効化する処理
func enableHost(hostname string) error {
	token, err := authKeyStone()
	if err != nil {
		log.Fatal(err)
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		ZUN_HOST+ZUN_ENABLE_SERVICE_PATH,
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//token 設定
	req.Header.Set("X-Auth-Token", token)

	//クエリパラム設定
	query := req.URL.Query()
	query.Add("binary", ZUN_SERVICE_BINARY)
	query.Add("host", hostname)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	log.Println("zunservice: enable hostname = " + hostname)
	return nil
}

func authKeyStone() (string, error) {
	jsonStr := createAuthJSONStr()
	authURL := os.Getenv("OS_AUTH_URL")

	req, err := http.NewRequest(
		"POST",
		authURL+AUTH_PATH,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	//コンテントタイプをせってい
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer resp.Body.Close()
	token := resp.Header.Get("X-Subject-Token")

	return token, nil
}

func createAuthJSONStr() string {
	userDomainName := os.Getenv("OS_USER_DOMAIN_NAME")
	userName := os.Getenv("OS_USERNAME")
	password := os.Getenv("OS_PASSWORD")
	projectDomainName := os.Getenv("OS_PROJECT_DOMAIN_NAME")
	projectName := os.Getenv("OS_PROJECT_NAME")
	jsonStr := `
{ 
    "auth": { 
        "identity": { 
            "methods":[
                "password"
            ],
            "password": {
                "user": {
                    "domain": {
                        "name": "` + userDomainName + `"
                    },
                    "name": "` + userName + `", 
                    "password": "` + password + `"
                } 
            } 
        }, 
        "scope": { 
            "project": { 
                "domain": { 
                    "name": "` + projectDomainName + `" 
                }, 
                "name":  "` + projectName + `" 
            } 
        } 
    }
}`
	return jsonStr
}

func randomString() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func startContainer(uuid string) error {
	//start container
	token, err := authKeyStone()
	if err != nil {
		log.Fatal(err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		ZUN_HOST+ZUN_PATH+uuid+"/start",
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//コンテントタイプをせってい
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}
