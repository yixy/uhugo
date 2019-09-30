/*
Copyright © 2019 yixy <youzhilane01@gmail.com>

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
	"io/ioutil"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
	"github.com/yixy/fmatter"
	"github.com/yixy/uhugo/util"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the markdown file",
	Long: `Update will rewrite the markdown file in current directory:
1.	update front matter of the file, title will be rewrite 
	by filename, lastmod will be rewrite by mtime. Only support for yaml fmatter.
2.	if there is a .list file in current dir, the markdown filename and title in 
	front matter will be rewrite by the .list file. Do not modify the file between 
	the execution of the list command and the execution of the update command, 
	otherwise the file md5sum will be change. files with same md5sum will be ignored.

Notice: uhugo update command will rewrite the file, please backup of important data before execution`,
	Run: func(cmd *cobra.Command, args []string) {
		var mdList = make(map[string]string, 0)
		b, err := ioutil.ReadFile(".list")
		if err == nil {
			//save filenames into mdList from .list
			content := string(b)
			fileList := strings.Split(content, "\n")
			var ignore []string
			for _, file := range fileList {
				//filename|md5sum
				result := strings.Split(file, "|")
				if result[0] == "" || result[1] == "" || result[1] == util.EMPTY {
					continue
				}
				_, ok := mdList[result[1]]
				if !ok {
					mdList[result[1]] = result[0]
				} else {
					ignore = append(ignore, result[1])
				}
			}
			for _, value := range ignore {
				delete(mdList, value)
			}
		}
		files, err := ioutil.ReadDir(".")
		if err != nil {
			cmd.PrintErr(err.Error())
			return
		}
		for _, file := range files {
			//md filename
			name := file.Name()
			//filename without prefix and suffix
			_, isMd := util.GetMDRealName(name)
			if file.IsDir() || !isMd {
				continue
			}
			sum, err := util.GetFileMd5(name)
			if err != nil {
				cmd.PrintErr(err)
				return
			}
			//filename format from mdList
			target, ok := mdList[sum]
			target = fmt.Sprintf("%s.md", target)
			if !ok {
				target = name
			}
			//file mod time
			modTime := file.ModTime()
			modTimeStr := modTime.Format(time.RFC3339)
			//file front matter
			var fm map[string]interface{}
			content, err := fmatter.ReadFile(name, &fm)
			if err != nil {
				cmd.PrintErr(err.Error())
				return
			}
			if fm == nil {
				fm = make(map[string]interface{})
				fm["lastmod"] = ""
			}
			lastModV := fm["lastmod"]
			lastMod, ok := lastModV.(string)
			if !ok {
				cmd.PrintErr("lastModV is not string")
				return
			}
			if modTimeStr != lastMod || name != target {
				tmpFile := fmt.Sprintf("%s.tmp", target)
				//write fmatter
				title := strings.TrimSuffix(target, ".md")
				fm["title"] = title
				fm["lastmod"] = modTimeStr
				if fm["date"] == nil {
					fm["date"] = modTimeStr
				}
				if *categories != "" {
					fm["categories"] = strings.Split(*categories, ",")
				}
				if *tags != "" {
					fm["tags"] = strings.Split(*tags, ",")
				}
				fmByte, err := yaml.Marshal(fm)
				if err != nil {
					cmd.PrintErr(err.Error())
					return
				}
				md, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
				if err != nil {
					cmd.PrintErr(err.Error())
					return
				}
				_, err = md.Write([]byte("---\n"))
				if err != nil {
					cmd.PrintErr(err.Error())
					return
				}
				_, err = md.Write(fmByte)
				if err != nil {
					cmd.PrintErr(err.Error())
					return
				}
				_, err = md.Write([]byte("---\n"))
				if err != nil {
					cmd.PrintErr(err.Error())
					return
				}
				_, err = md.Write(content)
				if err != nil {
					cmd.PrintErr(err.Error())
					return
				}
				md.Close()
				err = os.Remove(name)
				if err != nil {
					cmd.PrintErr(err.Error())
					return
				}
				err = os.Rename(tmpFile, target)
				if err != nil {
					cmd.PrintErr(err.Error())
					return
				}
			}

		}
		cmd.Println("file is updated.")
	},
}

var categories *string
var tags *string

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	categories = updateCmd.Flags().StringP("categories", "c", "",
		"Specify categories of front matter, separate by comma.")
	tags = updateCmd.Flags().StringP("tags", "t", "",
		"Specify tags of front matter, separate by comma.")
}
