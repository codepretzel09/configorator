# confi-go-rator

lightweight configuration management tool based on Ansible, Chef, and Puppet

This tool can be used to configure certain config parameters on debian hosts.


# Getting Started


1. run bootstrap.sh / Clone the repo

    bootstrap.sh will install git and go and clone the repo
    
    Alternatively, If you already have git and go installed, you can clone the repo ``` git clone https://github.com/codepretzel09/configorator.git ``` and proceed to the next step

3. update the config.json file and set your user and password variables (Line 66 and 67 in configorator.go)

    The config.json file is where you'll set your config parameters for your debian hosts. You can view the example currently loaded in config.json

     for a more secure connection, use passwordless auth with a key and use this instead 

     ``` simplessh.ConnectWithKeyFile("hostname_to_ssh_to") ```

4. go run configorator.go

    With config parameters set, you can now run: ``` go run configorator.go ``` from within the project root directory and watch the output for feedback and/or possible errors


# Architecture

This tool uses a json configuration file to declare specific parameters. SSH connections are made to target hosts and commands are executed based on values parsed throughout the configuration file. 

# Challenge

The example in config.json is based off of the following challenge

Your configuration must specify a web server capable of running the PHP application below
- Both servers must respond 200 OK and include the string "Hello, world!" in their response to requests from curl -sv "http://ADDRESS";
- For the purposes of this challenge, please do not reboot any of the provided servers.
