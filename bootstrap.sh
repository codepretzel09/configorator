# bootstrap.sh
# Problems? make sure i'm executable by running:  chmod +x bootstrap.sh

apt-get install go -y
go mod init github.com/codepretzel09/configorator
go get github.com/sfreiberg/simplessh