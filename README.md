# confi***go***rator


lightweight configuration management tool for a timed challenge

This tool can be used to configure parameters on debian hosts


# Getting Started


1. run bootstrap.sh / Clone the repo

    bootstrap.sh will install git, go, and clone the repo
    
    Alternatively, If you already have git and go installed, you can clone the repo ``` git clone https://github.com/codepretzel09/configorator.git ``` and proceed to the next step

3. update the config.json file and set the ```pass``` variable (Line 67 in configorator.go)

    The config.json file is where you'll set your config parameters for your debian hosts. You can view the example currently loaded in config.json

     *for a more secure connection, use passwordless auth with a key and use this instead 

     ``` simplessh.ConnectWithKeyFile("hostname_to_ssh_to") ```


4. go run configorator.go

    With config parameters set, you can now run: ``` go run configorator.go ```


# Architecture

This tool uses a json configuration file to declare specific parameters. SSH connections are made to target hosts and commands are executed based on values parsed throughout the configuration file. 

# Challenge

The example in config.json is based off of the following challenge

    - configure two debian web servers with php code using the custom CM tool
    - both servers must return 200 OK from a curl request 
    - do not reboot servers at any time
