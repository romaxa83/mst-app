#### Simple project as API for library (books)
<div id="library-top"></div>

![-----------------------------------------------------](/storage/img/rainbow.png)
##### 📚 technology stack
<ul>
<li>Api docs - <a href="https://github.com/swaggo/gin-swagger">swagger</a></li>
<li>Framework - <a href="https://github.com/gin-gonic/gin">gin</a></li>
<li>Database - <a href="https://www.postgresql.org/">postgres</a></li>
<li>ORM - <a href="https://gorm.io/index.html">gorm</a></li>
<li>File storage - <a href="https://min.io/">minio</a></li>
<li>Logger - <a href="https://github.com/sirupsen/logrus">logrus</a></li>
</ul>

![-----------------------------------------------------](/storage/img/rainbow.png)
##### TODO
- [x] api documentation (http://127.0.0.1:8060/swagger/index.html)
- [ ] TLS
- [x] crud
    - [x] category
    - [x] author
    - [x] book
- [ ] soft/hard delete, restore from archive
    - [x] category
    - [ ] author
    - [ ] book
- [x] implementation relations - hasMany, many2many, polymorphic by gorm
- [x] pagination, filters
- [ ] import data from file
    - [ ] author
      - [x] csv
      - [ ] xls
      - [ ] json
- [ ] export data
    - [ ] author
        - [ ] csv
        - [ ] xls
        - [x] json
- [ ] upload image and send to storage
    - [x] author
    - [ ] book
- [ ] i18n

![-----------------------------------------------------](/storage/img/rainbow.png)
##### command

```sh
$ cp .env.dist .env # copy env file and fill variables
$ make run # run service
$ make swagger # generate swagger docs
$ make info # show info
```