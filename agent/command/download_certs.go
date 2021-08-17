package command

import (
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	URI          = "/v1/download/client/cert"
	DownloadPath = "/opt/cert.tar.gz"
)

//Download certs from url
func Download(address string) {
	url := "http://" + address + URI
	fmt.Println(url)
	if err := downloadCerts(url); err != nil {
		fmt.Println("download error:", err)
		return
	}
	//解压
	targetPath := utils.BlockchainRoot + "/organizations/crypto"
	if err := utils.ExecLocalCommand("mkdir -p " + targetPath); err != nil {
		fmt.Println(err)
	}

	if err := utils.ExecLocalCommand("mv " + DownloadPath + " " + targetPath); err != nil {
		fmt.Println(err)
		return
	}

	if err := utils.ExecLocalCommand("tar -zxvf " + targetPath + "/cert.tar.gz -C " + targetPath); err != nil {
		fmt.Println(err)
		return
	}

}

func downloadCerts(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(DownloadPath)
	defer out.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
