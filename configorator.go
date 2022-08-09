package main

import (
        "encoding/json"
        "fmt"
        "github.com/sfreiberg/simplessh"
        "io/ioutil"
        "os"
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
        FileRemove  []string `json:"files-remove"`
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

        }

        fmt.Println("")

        var user string = "root"
        var pass string = "password" // Set SSH password here for target hosts

        // For a more secure connection use a key based auth with:
        // simplessh.ConnectWithKeyFile("hostname_to_ssh_to")

        // COPY FILE(s) AND SET PARAMS
        for i := 0; i < len(config.Configs); i++ {
                client, err := simplessh.ConnectWithPassword(config.Configs[i].Address, user, pass)
                if err != nil {
                        panic(err)
                }

                file := fmt.Sprintf(config.Configs[i].File.Name)
                content := fmt.Sprintf(config.Configs[i].File.Content)
                perms := fmt.Sprintf(config.Configs[i].File.Perms)
                owner := fmt.Sprintf(config.Configs[i].File.Owner)
                group := fmt.Sprintf(config.Configs[i].File.Group)
                defer client.Close()

                fmt.Println("Writing File(s) on " + file + " on " + config.Configs[i].Name)

                client.Exec("echo ' " + content + "' > " + file)
                client.Exec("chmod ' " + perms + " " + file)
                client.Exec("chown ' " + owner + " " + file)
                client.Exec("chgrp ' " + group + " " + file)
        }
        // REMOVE FILE(S)
        for i := 0; i < len(config.Configs); i++ {
                client, err := simplessh.ConnectWithPassword(config.Configs[i].Address, user, pass)
                if err != nil {
                        panic(err)
                }

                files := config.Configs[i].FileRemove
                fmt.Println("Removing File(s) on", config.Configs[i].Name, files)
                defer client.Close()

                for _, element := range files {
                        client.Exec("rm " + element)
                }
                files = files[:0] // Clear the slice

        }
        // INSTALL PACKAGES
        for i := 0; i < len(config.Configs); i++ {
                client, err := simplessh.ConnectWithPassword(config.Configs[i].Address, user, pass)
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
        // REMOVE PACKAGES
        for i := 0; i < len(config.Configs); i++ {
                client, err := simplessh.ConnectWithPassword(config.Configs[i].Address, user, pass)
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
        for i := 0; i < len(config.Configs); i++ {
                client, err := simplessh.ConnectWithPassword(config.Configs[i].Address, user, pass)
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