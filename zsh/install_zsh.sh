if [ -f /etc/issue ]; then
  sudo yum install zsh git -y
fi
curl -L https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh | sh
chsh -s /bin/zsh
home=`echo ~`
if [ -f /etc/issue ]; then
  sed -i 's/robbyrussell/af-magic/g' $home/.zshrc
else
  sed -i "" 's/robbyrussell/af-magic/g' $home/.zshrc
fi
source ~/.zshrc
