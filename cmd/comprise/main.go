// The MIT License (MIT)
//
// Copyright (c) 2015 Alexander Thaller
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package main

import (
	"log"
	"net/rpc"
	"os"

	"github.com/AlexanderThaller/comprise"
)

func main() {
	connection, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("can not dial to server:", err)
	}

	name, err := os.Hostname()
	if err != nil {
		log.Fatal("can not get hostname: ", err)
	}

	var client comprise.Client
	err = connection.Call("Server.RegisterClient", &name, &client)
	if err != nil {
		log.Fatal("can not register with server: ", err)
	}

	err = connection.Call("Server.UnRegisterClient", &client, nil)
	if err != nil {
		log.Fatal("can not unregister with server: ", err)
	}
}
