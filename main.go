// Copyright 2022 Canvas02 <Canvas02@protonmail.com>.
// SPDX-License-Identifier: Unlicense

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const PASTE_RS_URL = "http://paste.rs/"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf(" usage:\n    get [id]")
	}

	command := strings.ToLower(os.Args[1])
	if command != "get" {
		return fmt.Errorf("only the get command is implemented")
	}

	id := strings.ToLower(os.Args[2])
	url := PASTE_RS_URL + id

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	println("------ ", url, " -----")
	fmt.Printf("\n%s\n", body)

	return nil
}
