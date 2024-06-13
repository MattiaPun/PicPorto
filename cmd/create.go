/*
Copyright © 2024 Mattia Pun <mattia@punjwani.be>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var createCmd = &cobra.Command{
	Use:   "create [project]",
	Short: "Create a new PicPorto project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		os.Mkdir(args[0], 0755)
		os.Mkdir(args[0]+"/photos", 0755)
		os.Mkdir(args[0]+"/photos/cover", 0755)
		os.Mkdir(args[0]+"/photos/avatar", 0755)
		os.Mkdir(args[0]+"/photos/pictures", 0755)

		template := map[string]interface{}{
			"style": map[string]string{
				"theme":   "light",
				"content": "left",
			},
			"profile": map[string]string{
				"name": "John Doe",
				"bio":  "",
			},
			"social": map[string]string{
				"twitter":   "",
				"facebook":  "",
				"instagram": "",
				"linkedin":  "",
			},
			"content": map[string]string{
				"title": "My Portfolio",
				"tags":  "Photography,Portfolio",
				"text":  "Welcome to my portfolio!",
			},
		}

		yamlBytes, err := yaml.Marshal(template)
		if err != nil {
			fmt.Println("Error marshalling YAML:", err)
			return
		}

		err = ioutil.WriteFile(args[0]+"/config.yaml", yamlBytes, 0644)
		if err != nil {
			fmt.Println("Error writing to config file:", err)
			return
		}

		fmt.Print("Created project: " + args[0] + "\n")
	},
}
