/**
 * Copyright 2022 Orix Au Yeung
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	log.Print("openFWUsingDomainName: received a request")

	cmd := exec.Command("/bin/sh", "./script.sh")
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(out)
	log.Print("openFWUsingDomainName: fulfilled the request")
}
