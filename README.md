this project is just for having fun, I really frustrated because need to recheck continuously my last status resi. :'(

this app will be check the latest status of my resi every 10 minutes.

oya, I think it only solve my current problem. not for general problem. :D

for ubuntu users,
you need to install `libasound2-dev`

```sh
sudo apt install libasound2-dev
```

look at `sample.sh`
copy it to `env.sh`
you need to fill the required data in `env.sh`
```sh
source env.sh
```


1. run the server
```sh
make server
```

2. run the app
```sh
make app
```