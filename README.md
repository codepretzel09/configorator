# configorator
config mgmt w/ go


Getting Started

1. Ensure you have passwordless key auth to target hosts
2. Clone the repo
3. run bootstrap.sh
4. update the config.json file
5. go run configorator.go



1. Ensure you have passwordless key auth to target hosts

    To set up passwordless ssh, run the following command for all target hosts 
    ssh-copy-id root@X.X.X.X 

2. Clone the repo

    git clone https://github.com/codepretzel09/configorator.git

3. Run bootstrap.sh

    chmod +x bootstrap.sh
    ./bootstrap.sh 
