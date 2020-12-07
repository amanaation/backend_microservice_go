# backend_microservice_go

## To setup the above project you will need to install the following:
1. Apache Kafka
2. MySQL
3. Golang

### 1. After installing MySQL, login to your MySQL terminal and type the following queries in your terminal : 
### User Table

```SQL
CREATE TABLE `User` (
  `uid` int NOT NULL AUTO_INCREMENT,
  `lname` varchar(15) NOT NULL,
  `name` varchar(15) NOT NULL,
  `credit_limit` int NOT NULL,
  `credit_balance` int NOT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |
```
#### Order_details Table:
```SQL
CREATE TABLE `Order_details` (
  `eventid` int NOT NULL AUTO_INCREMENT,
  `orderid` int NOT NULL,
  `uid` int NOT NULL,
  `order_status` varchar(15) NOT NULL,
  `event_timestamp` varchar(40) NOT NULL,
  `order_amount` int DEFAULT NOT NULL,
  `order_notes` varchar(100) DEFAULT NOT NULL,
  PRIMARY KEY (`eventid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |

```


#### 2. After installing Go, clone this project on your local system. Navigate to the cloned directory and run the following commands:
```go 
go get "github.com/asaskevich/govalidator"
go get "github.com/segmentio/kafka-go"
go get "github.com/segmentio/kafka-go"
go get "database/sql"
go get "github.com/confluentinc/confluent-kafka-go/kafka"
```

#### 3. Start your Kafka and zookeeper server
#### 4. Run main.go
#### 5. Run User/ProcessUser.go
#### 6. Run Order/Process_Order.go

Explanation : 
There are 2 micro-services running in the background : ProcessOrder and ProcessUser. You can access both the micro-services using Flask Api
```
1. localhost:8080/order?para1=val1&para2=val2
2. localhost:8080/user?para1=val1&para2=val2
```
You can publish 4 types of events to UserOperation :
```
1. Create User : It will have following parameters in its request : type = C, lname, name, credit_limit, credit_balance
3. Update User : It will have following parameters in its request : type = U, setcond1=val1, setcon2=val2, col1=val1, col2=val2. Make sure to prefix word "set" for condition columns.
4. Delete User : It will have following parameters in its request : type = D, cond1=val1, con2=val2
```

You can publish order event to ProcessOrder microservice.
```
1. Create Order : It will have following parameters in its request : uid, amount
```
