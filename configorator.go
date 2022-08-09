package main

import (
        "fmt"
        "github.com/sfreiberg/simplessh"
        "os"
        "encoding/json"
        "io/ioutil"
        )

// Config struct which contains
// an array of server configs
type Configs struct {
        Configs []Config `json:"config"`
}

// Config struct
type Config struct {
        Name        string   `json:"name"`
        Address     string   `json:"address"`
        AppsInstall []string `json:"apps-install"`
        AppsRemove  []string `json:"apps-remove"`
        Reload      []string `json:"reload"`
        File        File     `json:"file"`
}

// File struct 
type File struct {
        Name    string `json:"name"`
        Content string `json:"content"`
        Owner   string `json:"owner"`
        Group   string `json:"group"`
        Perms   string `json:"perms"`
}

func main() {
        // Open our jsonFile
        jsonFile, err := os.Open("config.json")
        // if we os.Open returns an error then handle it
        if err != nil {
                fmt.Println(err)
        }

        fmt.Println("Successfully Opened config.json")
        // defer the closing of our jsonFile so that we can parse it later on
        defer jsonFile.Close()

        // read our opened xmlFile as a byte array.
        byteValue, _ := ioutil.ReadAll(jsonFile)

        // we initialize our Config array
        var config Configs

        // we unmarshal our byteArray which contains our
        // jsonFile's content into 'config' which we defined above
        json.Unmarshal(byteValue, &config)

        // we iterate through every user within our config array
        for i := 0; i < len(config.Configs); i++ {
/*              fmt.Println("IP: " + config.Configs[i].Address)
                fmt.Println(config.Configs[i].AppsInstall)
                fmt.Println(config.Configs[i].AppsRemove)
                fmt.Println("Host: " + config.Configs[i].Name)
                fmt.Println("owner: " + config.Configs[i].File.Owner)
                fmt.Println("filename: " + config.Configs[i].File.Name)
                fmt.Println("content: " + config.Configs[i].File.Content)
                fmt.Println("group: " + config.Configs[i].File.Group)
                fmt.Println("perms: " + config.Configs[i].File.Perms)
                fmt.Println(config.Configs[i].Reload) */

        }

        fmt.Println("")

        // COPY FILE(s) AND SET PARAMS
        // Req: abstraction that allows specifying a file's content and metadata
        // Action Items
        // - Improve Idempotency
        // - Improve iteration of more than one file

        for i := 0; i < len(config.Configs); i++ {
                client, err := simplessh.ConnectWithKeyFile(config.Configs[i].Address, "root", "/root/.ssh/id_rsa")
                if err != nil {
                        panic(err)
                }

                file := fmt.Sprintf(config.Configs[i].File.Name)
                content := fmt.Sprintf(config.Configs[i].File.Content)
                perms := fmt.Sprintf(config.Configs[i].File.Perms)
                owner := fmt.Sprintf(config.Configs[i].File.Owner)
                group := fmt.Sprintf(config.Configs[i].File.Group)
                defer client.Close()

                fmt.Println("writing file " + file + " on " + config.Configs[i].Name)

                client.Exec("echo ' " + content + "' > " + file)
                client.Exec("chmod ' " + perms + " " + file)
                client.Exec("chown ' " + owner + " " + file)
                client.Exec("chgrp ' " + group + " " + file)

        }

        // INSTALL AND REMOVE PACKAGES
        // Req. Your tool must provide an abstraction that allows installing and removing Debian packages
        // Action Items
        // - Improve Idempotency

        for i := 0; i < len(config.Configs); i++ {
                client, err := simplessh.ConnectWithKeyFile(config.Configs[i].Address, "root", "/root/.ssh/id_rsa")
                if err != nil {
                        panic(err)
                }

                pkgList := config.Configs[i].AppsInstall
                fmt.Println("Installing packages on", config.Configs[i].Name, pkgList)
                defer client.Close()

                for _, element := range pkgList {
                        client.Exec("apt-get install " + element + " -y")
                }
                pkgList = pkgList[:0] // Clear the slice
        }

        for i := 0; i < len(config.Configs); i++ {
                client, err := simplessh.ConnectWithKeyFile(config.Configs[i].Address, "root", "/root/.ssh/id_rsa")
                if err != nil {
                        panic(err)
                }

                pkgList := config.Configs[i].AppsRemove
                fmt.Println("Removing packages on", config.Configs[i].Name, pkgList)
                defer client.Close()

                for _, element := range pkgList {
                        client.Exec("apt remove " + element + " -y")
                }
                pkgList = pkgList[:0] // Clear the slice
        }

        // RESTART SERVICES
        // Req.  Your tool must provide some mechanism for restarting a service when relevant files or packages are updated
        // Action Items
        // - Build in file/pkg change detection (if this file has changed, restart x service)

        for i := 0; i < len(config.Configs); i++ {
                client, err := simplessh.ConnectWithKeyFile(config.Configs[i].Address, "root", "/root/.ssh/id_rsa")
                if err != nil {
                        panic(err)
                }

                pkgList := config.Configs[i].Reload
                fmt.Println("Restarting packages on", config.Configs[i].Name, pkgList)

                for _, element := range pkgList {
                        client.Exec("systemctl restart " + element)
                }

        }

}