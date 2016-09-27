#GO-PAD

##Go pad allows 

For quick notes taken within your terminal.
`gopad` 
will run gopad note taker write in a quick note

###WEB Interface I have now built a simple web interface for Go pad

web interface shows all notes and allows notes to be added

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

load script to bash profile please run the login command
`bash --login`

This will refresh your terminal with this command you will **not need to exit and restart** your terminal for new bash_profile.

###TODO
- [ ] Change Note type from csv too json
- [ ] Create way to search notes
- [ ] Add edit and destroy methods too notes in web view
- [ ] Add login and session data for multi user env
- [ ] Create Desktop app
- [ ] Make site use AJAX live update site when anyone adds note