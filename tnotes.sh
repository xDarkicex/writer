#You should add this command to your .bash_Profile replace #{username with your username}


function tnote() {
username="Your username here"
echo "Terminal Notes Quick notes Powered by GOlang"
cd $GOPATH/src/github.com/${username}/writer/ &&
go run writer.Go
    
if [ "$1" == "read" ]; then
  cd $GOPATH/src/github.com/${username}/writer/ &&
  go run reader.Go
fi
if [ "$1" == "config" ]; then
    open $GOPATH/src/github.com/${username}/writer/config.Go;
fi
}
 