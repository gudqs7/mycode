if [ -f /etc/issue ]; then
  sh ./zsh_highlight_centos.sh
else
  sh ./zsh_highlight_osx.sh
fi
