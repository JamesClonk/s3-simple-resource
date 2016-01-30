/*
   Copyright 2016 Fabio Berchtold

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/JamesClonk/s3-simple-resource"
)

type Request struct {
	Source s3resource.Source `json:"source"`
	Params Params            `json:"params"`
}

type Params struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type Response struct {
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <destination>\n", os.Args[0])
		os.Exit(1)
	}

	destination := os.Args[1]
	if err := os.MkdirAll(destination, 0755); err != nil {
		log.Fatalf("creating destination directories: %v\n", err)
	}

	var request Request
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		log.Fatalf("reading request from stdin: %v\n", err)
	}

	client, err := s3resource.NewClient(
		request.Source.AccessKey,
		request.Source.SecretKey,
		request.Source.Host,
		request.Source.Insecure,
	)
	if err != nil {
		log.Fatalf("building s3 client: %v\n", err)
	}

	if err := client.DownloadFile(request.Source.Bucket, request.Params.Source, filepath.Join(destination, request.Params.Target)); err != nil {
		log.Fatalf("downloading file: %v\n", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(Response{}); err != nil {
		log.Fatalf("writing response to stdout: %v\n", err)
	}
}
