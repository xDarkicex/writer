#GO-PAD

##Go pad allows 

For quick notes to be taken within your terminal.
`gopad` 
will run gopad note taker write in a quick note

read them anytime with the read command
`gopad read`



I have included a bash script that will allow you to call this **Go-PAD** from any directory once included into your `./Bash_Profile`

to add `gopad.sh` to your .Bash_Profile open terminal 

type:
`open ~/.Bash_Profile`
This will open a text editor apple default will be textedit.

copy the contents of gopad.sh to the bottom of your .Bash_Profile

exchange username="Your Username goes here" to your proper username 

example: 
`username="xDarkicex"`

please note this will only function if you have a standard goworkspace setup with github.com
`$GOPATH/src/github.com/xDarkicex/writer/`

once script is inserted into bash profile please run the login command
`bash --login`

This will refresh your terminal with this command you will **not need to exit and restart** your terminal for new bash_profile.