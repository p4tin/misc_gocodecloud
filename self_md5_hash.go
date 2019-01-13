package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"net/mail"
	"net/smtp"
	"strconv"

	"github.com/hpcloud/tail"
	"github.com/labstack/echo"
)

var hostname string

func tailLogFile(filename string) {
	t, _ := tail.TailFile(filename, tail.Config{Follow: true})
	for line := range t.Lines {
		fmt.Println("tail", filename, line.Text)
	}
}

func hash_file_md5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

type FileInfo struct {
	MD5Hash string
	Status  string
	Self    bool
}

var FileHashKeep map[string]FileInfo

func webServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, FileHashKeep)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

type EmailConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func sendMail(filename string) {
	// authentication configuration
	smtpHost := "smtp.gmail.com"    // change to your SMTP provider address
	smtpPort := 587                 // change to your SMTP provider port number
	smtpPass := "iv17wt0c"          // change here
	smtpUser := "pfort69@gmail.com" // change here

	emailConf := &EmailConfig{smtpUser, smtpPass, smtpHost, smtpPort}

	emailauth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.Host)

	sender := "fromwho@gmail.com" // change here

	receivers := []string{
		"paul@cp-soft.com",
	} // change here

	from := mail.Address{"CPIDS", "pfort69@cgmail.com"}
	to := mail.Address{"Paul Fortin", "paul@cp-soft.com"}
	subject := fmt.Sprintf("File %s has been modified.", filename)
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("File %s has been modified.", filename)))

	// send out the email
	err := smtp.SendMail(smtpHost+":"+strconv.Itoa(emailConf.Port), //convert port number from int to string
		emailauth,
		sender,
		receivers,
		[]byte(message),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	hostname, _ = os.Hostname()
	fmt.Println("Strating monitoring on", hostname)
	FileHashKeep = make(map[string]FileInfo)
	listofFiles := []string{os.Args[0], "/etc/hosts", "/etc/passwd"}

	go webServer()
	go tailLogFile("/Users/Paul/tmp/tail.log")

	for {
		for _, file := range listofFiles {
			hash, err := hash_file_md5(file)
			if err == nil {
				if _, ok := FileHashKeep[file]; ok {
					fi := FileHashKeep[file]
					if fi.MD5Hash != hash {
						fi.MD5Hash = hash
						if fi.Status == "OK" {
							fi.Status = "Changed"
							sendMail(file)
						}
						FileHashKeep[file] = fi
						fmt.Printf(
							"\tFile %s does not match md5s, old md5 = %s, new md5 = %s\n",
							file,
							FileHashKeep[file],
							hash)
					} else {
						if fi.Status == "Changed" {
							fi.Status = "OK"
						}
						FileHashKeep[file] = fi
					}
				} else {
					FileHashKeep[file] = FileInfo{MD5Hash: hash, Status: "OK"}
				}
			}
		}
		fmt.Println(FileHashKeep)
		time.Sleep(15 * time.Second)
	}
}
