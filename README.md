## go-shop
a simple REST-api that calculates products final price  
# overview #  
this api recieves a bunch of products that are available inside our database  
and responses final payable price that user must pay(considering off)  
amd also lists the products in user's basket  
- - - -  
# functionality #  
imagine we have a very simple table in our database:  
![picture alt](/photos/table.jpg)  
we'll have these two requests as below:
### goods ###  
by sending `localhost:8080/goods`,  
if you're using postman send this request:  
![picture alt](/photos/goods.jpg)  
then you'll receive all products inside database in json format:  
![picture alt](/photos/goods_r.jpg)  
### buy ###  
by sending `localhost:8080/buy`,  
if you're using postman send this request(as an example):  
![picture alt](/photos/buy.jpg)  
(the numbers inside the slice are ids of products)  
then you'll receive final price and all selected products:  
![picture alt](/photos/buy_r.jpg)  
- - - -  
# how to use the code #  
after installing the commands, you should run `shoppingServer start` as below to run the server:  
![picture alt](/photos/start.gif)  
then you can send requests mentioned above :)  
you can intrrupt the server by using<kbd>ctrl + C</kbd>  
__NOTE:__ this CLI app is tested inside jetbrains goland
