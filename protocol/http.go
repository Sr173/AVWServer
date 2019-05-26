package protocol

import (
	"bufio"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

type ScanResult struct{
	a string

}

func HttpGetHandler(w http.ResponseWriter, r*http.Request) {
	getPath_ := strings.Split(r.RequestURI,"/")
	println(r.Method,r.RequestURI,getPath_[1])
	remote := r.RemoteAddr
	remote = remote[0:strings.Index(remote,":")]
	fmt.Print(getPath_)
	if (strings.Contains(getPath_[1],"api")){
		if (strings.Contains(getPath_[2],"db")){
			if (strings.Contains(getPath_[3],"get")){
				if (getPath_[4] == "all"){
					Conn_map[remote] <- "GGGGG"
				}
			}
		}
	}
}

func HttpPostHandler(w http.ResponseWriter, r*http.Request){
	remote := r.RemoteAddr
	remote = remote[0:strings.Index(remote,":")]
	data, _ := ioutil.ReadAll(r.Body)
	uuid,_ := uuid.NewV1()
	file_name := string("C:/Users/admin/AppData/Roaming/ClamAv/file/");
	file_name += uuid.String();
	err2 := ioutil.WriteFile(file_name, data, 0666) //写入文件(字节数组)
	if(err2 != nil) {
		return
	}
	cmd := exec.Command("C:/Users/admin/AppData/Roaming/ClamAv/bin/clamscan.exe","-d","C:\\Users\\admin\\AppData\\Roaming\\ClamAv\\db\\main.cvd")
	ppReader, _ := cmd.StdoutPipe()
	var bufReader = bufio.NewReader(ppReader)
	cmd.Start()
	go func() {
		var buffer []byte = make([]byte, 4096)
		for {
			n, err := bufReader.Read(buffer)
			if err != nil {
				if err == io.EOF {

					fmt.Printf("pipi has Closed\n")
					break
				} else {
					fmt.Println("Read content failed")
				}
			}
			fmt.Print(string(buffer[:n]))
		}
	}()
	cmd.Wait()
}