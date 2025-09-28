/*
Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"

	"github.com/J-Siu/go-crypto/crypto"
	"github.com/J-Siu/go-helper/v2/err"
	"github.com/spf13/cobra"
)

var key string
var msg string

// sealanonymousCmd represents the sealanonymous command
var sealanonymousCmd = &cobra.Command{
	Use:   "sealanonymous",
	Short: "x/crypto box seal anonymous. Output is base64 encoded",
	Run: func(cmd *cobra.Command, args []string) {
		encrypted_msg, e := crypto.BoxSealAnonymous(&key, &msg)
		if e == nil && encrypted_msg != nil {
			fmt.Println(*encrypted_msg)
		}
		err.Queue(e)
	},
}

func init() {
	boxCmd.AddCommand(sealanonymousCmd)
	sealanonymousCmd.Flags().StringVarP(&key, "key", "k", "", "(required) base64 encoded public key")
	sealanonymousCmd.Flags().StringVarP(&msg, "msg", "m", "", "(required) plain text message")
	sealanonymousCmd.MarkFlagRequired("key")
	sealanonymousCmd.MarkFlagRequired("msg")
}
