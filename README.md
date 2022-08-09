# confi-go-rator

lightweight CM tool based on Ansible, Chef, and Puppet

This tool can be used to configure certain config parameters on debian hosts.


# Getting Started

1. Ensure you have passwordless key auth to target hosts

    To set up passwordless ssh, run the following command for all target hosts 
   ``` ssh-copy-id root@X.X.X.X ```

2. run bootstrap.sh / Clone the repo

    bootstrap.sh will install git and go and clone the repo
    
    Alternateively, If you aleady have git and go installed, you can clone the repo ``` git clone https://github.com/codepretzel09/configorator.git ``` and proceed to the next step

3. update the config.json file

    The config.json file is where you'll set your config parameters for your debian hosts. You can view the example currently loaded in config.json

4. go run configorator.go

    With config parameters set, you can now run: ``` go run configorator.go ``` and watch the output for feedback and possible errors


# Architecture

This tool uses a json configuration file to declare specific parameters. SSH connections are made to target hosts and commands are executed based on values parsed throughout the configuration file. 

# Challenge

The example in config.json is based off of the following challenge

Your configuration must specify a web server capable of running the PHP application below
- Both servers must respond 200 OK and include the string "Hello, world!" in their response to requests from curl -sv "http://ADDRESS";
- For the purposes of this challenge, please do not reboot any of the provided servers.
