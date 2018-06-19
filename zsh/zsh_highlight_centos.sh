git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
sed -i "" 's/  git/git  zsh-syntax-highlighting/g' ~/.zshrc
