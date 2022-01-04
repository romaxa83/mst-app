#### Simple project as API for library (books)
<div id="library-top"></div>

![-----------------------------------------------------](mst-app/storage/img/rainbow.png)
##### 📚 technology stack
<ul>
<li>Api docs -<a href="https://github.com/swaggo/gin-swagger">swagger</a></li>
<li>Framework - <a href="https://github.com/gin-gonic/gin">gin</a></li>
<li>Database - <a href="https://www.postgresql.org/">postgres</a></li>
<li>ORM - <a href="https://gorm.io/index.html">gorm</a></li>
<li>File storage - <a href="https://min.io/">minio</a></li>
<li>Logger - <a href="https://github.com/sirupsen/logrus">logrus</a></li>
</ul>

![-----------------------------------------------------](./mst-app/storage/img/rainbow.png)
##### features

✅&nbsp;&nbsp;api documentation (http://127.0.0.1:8060/swagger/index.html) <br>
✅&nbsp;&nbsp;crud for entities <br>
✅&nbsp;&nbsp;soft/hard delete? restore from archive <br>
✅&nbsp;&nbsp; implementation relations - hasMany, many2many, polymorphic by gorm<br>
✅&nbsp;&nbsp; pagination<br>
✅&nbsp;&nbsp; import data (only author) from file (only csv)<br>
✅&nbsp;&nbsp; export data (only author) from file (only json)<br>
✅&nbsp;&nbsp; upload image for author and send to storage (minio)<br>

![-----------------------------------------------------](docs/rainbow.png)
##### command

```sh
$ cp .env.dist .env # copy env file and fill variables
$ make run # run service
$ make swagger # generate swagger docs
$ make info # show info
```
