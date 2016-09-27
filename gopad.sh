#You should add this command to your .bash_Profile replace #{username with your username}

function gopad() {
username="Your Username goes here"
run="main.go"
local=${PWD}

if (( $# !=1 )); then
cd $GOPATH/src/github.com/${username}/writer/ &&
go run ${run} &&
cd ${PWD};
fi
}
 