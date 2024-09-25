// Copyright 2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main implements a very simple plugin that just outputs text files
// with the names of the top-level messages in each file.
//
// Example: if a/b.proto had top-level messages C, D, the file "a/b.proto.txt" would be
// outputted, containing "C\nD\n".
package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/bufbuild/protoplugin"
	"io"
	"log"
	"os"
	"os/exec"
	"protoc-gen-open-models/pkg/generator"
)

const version = "0.0.1"

func main() {
	const debugMode = false

	// osEnv is the os-based Env used in Main.
	var osEnv = protoplugin.Env{
		Args:    os.Args[1:],
		Environ: os.Environ(),
		Stdin:   os.Stdin,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
	}

	if debugMode {
		data, err := os.ReadFile("protocdata")
		if err != nil {
			// debug file does not exist, read from stream and write file
			data, err = io.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
			os.WriteFile("protocdata", data, 0644)
		}
		osEnv = protoplugin.Env{
			Args:    os.Args[1:],
			Environ: os.Environ(),
			Stdin:   bytes.NewReader(data),
			Stdout:  os.Stdout,
			Stderr:  os.Stderr,
		}

	}

	ctx, cancel := context.WithCancel(context.Background())
	if err := protoplugin.Run(ctx, osEnv, protoplugin.HandlerFunc(handle), protoplugin.WithVersion(version)); err != nil {
		exitError := &exec.ExitError{}
		if errors.As(err, &exitError) {
			cancel()
			// Swallow error message - it was printed via os.Stderr redirection.
			os.Exit(exitError.ExitCode())
		}
		if errString := err.Error(); errString != "" {
			_, _ = fmt.Fprintln(os.Stderr, errString)
		}
		cancel()
		os.Exit(1)
	}

	cancel()
}

func handle(
	_ context.Context,
	_ protoplugin.PluginEnv,
	responseWriter protoplugin.ResponseWriter,
	request protoplugin.Request,
) error {
	// Set the flag indicating that we support proto3 optionals. We don't even use them in this
	// plugin, but protoc will error if it encounters a proto3 file with an optional but the
	// plugin has not indicated it will support it.
	responseWriter.SetFeatureProto3Optional()
	generator.GenerateAll(responseWriter, request)

	return nil
}
