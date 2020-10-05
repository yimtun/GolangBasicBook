```
yum -y install wget
```

```
wget http://mirror.ghettoforge.org/distributions/gf/gf-release-latest.gf.el7.noarch.rpm 
```



```
rpm -ivh gf-release-latest.gf.el7.noarch.rpm
```

```
yum remove vim-* -y
```


```
yum --enablerepo=gf-plus install vim-enhanced  -y
```
