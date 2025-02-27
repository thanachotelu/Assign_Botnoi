# Programming Test

### set up
### Assign 2

``` bash
# import library
$ go get "github.com/gin-gonic/gin"

# manage dependencies
$ go mod tidy
```

### Assign 3 (ผมรู้ว่า .env ไม่ควรที่จะ push ขึ้น git แต่ผมไม่รู้ว่าจะให้ตัว environment ที่ใช้อย่างไรก็เลย push .env มาด้วย)
### QR linechatbot
![QR_line_chatbot](Assign3/QR_thanachotelu.png)
``` bash
# import library
$ go get "github.com/line/line-bot-sdk-go/v8/linebot"
$ go get "github.com/spf13/viper"

# manage dependencies
$ go mod tidy

# run ngork to use webhook url
$ ngrok http --url=growing-mistakenly-terrier.ngrok-free.app 8081
```
### ฟังก์ชันการทำงาน
* Text (ถ้านอกเหนือคำเหล่านี้จะตอบกลับ "ขอโทษด้วย ผมไม่เข้าใจ")
    * สวัสดี
    * yo
    * hi
    * ทำอะไรอยู่
* button
    * button
* quickreply
    * quickreply
* carosule
    * carosule
 
### Assign4 - API CRUD (เลือกทำข้อ 5)
``` bash
# import library
$ go get "github.com/gin-gonic/gin"
$ go get "go.mongodb.org/mongo-driver"
$ go get "github.com/spf13/viper"
$ go get "golang.org/x/crypto/bcrypt"

# manage dependencies
$ go mod tidy

# docker-compose
$ docker-compose up -d
```
### API CRUD
* Healthcheck : localhost:8080/health
* GET
    * GetAllUsers : localhost:8080/api/v1/users/
    * GetUserByID : localhost:8080/api/v1/users/{:id}
* POST
    * AddUser : localhost:8080/api/v1/users/
* PUT
    * UpdateUser : localhost:8080/api/v1/users/{:id}
* DELETE
    * DeleteUser : localhost:8080/api/v1/users/{:id}
