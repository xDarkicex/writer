#You should add this command to your .bash_Profile replace #{username with your username}

function note() {
username="Your Username goes here"
version="1.0.0"
write="writer.go"
read="reader.go"
local=${PWD}
echo "
---------------------------------------------------------------------------------------------
#   Golang Terminal Notes                                                               #
#  ----------------------                                                                   #
#  | Version #$version     |                                                                   #
#  ----------------------                                                                   #
#  By: Gentry Rolofson                                                                      #
#  More info: https://bitdev.io                                                             #
#  Github: @xDarkicex                                                                       #
#                                                                                           #
---------------------------------------------------------------------------------------------
"

if [ "$1" == "read" ]; then
    cd $GOPATH/src/github.com/${username}/writer/ &&
    go run ${read} &&
    cd ${PWD}

fi
if [ "$1" == "config" ]; then
    open $GOPATH/src/github.com/${username}/writer/config/config.Go &&
    cd ${PWD}

fi

if (( $# !=1 )); then
cd $GOPATH/src/github.com/${username}/writer/ &&
go run ${write} &&
cd ${PWD};
fi
}
 