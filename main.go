package main

import (
//    "crypto/tls"
    "fmt"
    "log"
    "os"
//    "time"
    "ftp/intercomftp"
)

func main() {

    fileName := "./asd1.txt"
    var err error
    var ftp, ftp1 *intercomftp.FTP

    log.Println("1", err)

    if ftp, err = intercomftp.Connect("127.0.0.1:21"); err != nil {
	fmt.Println(err)
    }

    defer ftp.Close()

    if ftp1, err = intercomftp.Connect("127.0.0.1:21"); err != nil {
	fmt.Println(err)
    }

    defer ftp1.Close()

//    config := tls.Config{
//	InsecureSkipVerify: true,
//	ClientAuth:         tls.RequestClientCert,
//    }
//    log.Println("13", err)

//    if err = ftp.AuthTLS(config); err != nil {
//	//      log.Println("1", err)
//    }

    if err = ftp.Login("user", "root1"); err != nil {
	log.Println("2", err)
    }
    //
    if err = ftp.Cwd("/"); err != nil {
	log.Println("3", err)
    }

    if err = ftp1.Login("user", "root1"); err != nil {
	log.Println("2", err)
    }
    //
    if err = ftp1.Cwd("/"); err != nil {
	log.Println("3", err)
    }

    log.Println("11", err)

    var file *os.File
    if file, err = os.Open(fileName); err != nil {
	log.Println("6", err)
    }
    defer file.Close()

    fmt.Println("start")

    log.Println("17", err)

//    go func() {
    fmt.Println("first")
    nmp := ftp.Stor("./asd5.txt", file)
    log.Println("nmp", nmp)
    if nmp != nil {
        log.Println("7", err)
    } else {
        fmt.Println("first is runung")
    }

//}()
//    go func() {
//	fmt.Println("second")
//	for {
//	    log.Println("tesT")
//	    files, nms := ftp1.List("./")
//	    log.Println("nms,files", nms, "F")
//
//	    if nms == nil {
//		fmt.Println(files)
//	    }
//	    time.Sleep(1 * time.Second)
//	}
//    }()

    fmt.Println("Press any key for exit ...")
    var mnmn string
    fmt.Scan(&mnmn)
}


