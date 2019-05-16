git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
git config --global alias.co checkout
git config --global alias.ci commit
git config --global alias.br branch
git config --global alias.st status
git config --global alias.unstage 'reset HEAD'
git config --global alias.last 'log -1'

git config --global alias.lw "log --pretty=format:'%cd--%s--%an--%h' --date=format:%Y-%m-%d"
