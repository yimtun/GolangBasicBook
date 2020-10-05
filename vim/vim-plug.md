```
yum -y install git
```


```
mkdir -p /.vim/autoload/
```

```
curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
```


```
cat > /root/.vimrc  << EOF
call plug#begin()
Plug 'fatih/vim-go'
call plug#end()
let g:go_version_warning = 0
EOF
```


直接输入 vim 


```
:PlugInstall
```
