#### Simple API project as todo list for user
<div id="todo-top"></div>

![-----------------------------------------------------](/storage/img/rainbow.png)
##### 📚 technology stack
<ul>
<li>Golang - 1.16</li>
<li>Api docs - <a href="https://github.com/swaggo/gin-swagger">swagger</a></li>
<li>Framework - <a href="https://github.com/gin-gonic/gin">gin</a></li>
<li>Database - <a href="https://www.postgresql.org/">postgres</a></li>
<li>File storage - <a href="https://github.com/jlaffaye/ftp">ftp</a></li>
<li>Logger - <a href="https://github.com/sirupsen/logrus">logrus</a></li>
</ul>

![-----------------------------------------------------](/storage/img/rainbow.png)
##### features

✅&nbsp;&nbsp;api documentation (http://127.0.0.1:8060/swagger/index.html) <br>
✅&nbsp;&nbsp;crud for entities (todo list, item) <br>
✅&nbsp;&nbsp;create user,crypto password, generate access/refresh token by <a href="https://github.com/dgrijalva/jwt-go">jwt</a> <br>
✅&nbsp;&nbsp;auth middleware <br>
✅&nbsp;&nbsp;send email<br>
✅&nbsp;&nbsp; upload image/file and send to storage (ftp)<br>
✅&nbsp;&nbsp; <a href="https://github.com/golang-migrate/migrate">migration</a><br>

![-----------------------------------------------------](/storage/img/rainbow.png)
##### command

```sh
$ cp .env.dist .env # copy env file and fill variables
$ make run # run service
$ make swagger # generate swagger docs
$ make migrate_up
$ make migrate_down
$ make info # show info
```

![-----------------------------------------------------](/storage/img/rainbow.png)
#####⚠ ️ install migration tool on linux (global)

```sh
$ https://github.com/mattes/migrate/releases/migrate.linux-amd64.tar.gz
$ tar -xvzf migrate.linux-amd64.tar.gz
$ sudo chmod +x migrate.linux-amd64
$ sudo cp migrate.linux-amd64 /usr/local/bin/
$ sudo ln /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate
$ export PATH = $PATH:/usr/local/bin
$ migrate
```