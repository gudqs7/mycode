if [ -f /etc/issue ]; then
  curl -L https://raw.githubusercontent.com/gudqs7/mycode/master/zsh/zsh_highlight_centos.sh | sh
else
  curl -L https://raw.githubusercontent.com/gudqs7/mycode/master/zsh/zsh_highlight_osx.sh | sh
fi
