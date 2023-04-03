const express = require("express");
const bodyParser = require("body-parser");

const app = express(); 
app.use(bodyParser.urlencoded({extended:true})); 


app.get("/" , function(req,res){
    var path = require("path");
    res.sendFile(path.join(__dirname, '..', '/index.html'));
}); 



app.post("/" , function(req,res){  
    console.log(req.body); 
    res.send("<h1>Some action needed</h1>");
    /*
    var num1 = Number(req.body.num1);
    var num2 = Number(req.body.num2); 
    var result = num1 + num2; 
    res.send("the result is: " + result);
    */
}); 

app.listen(8080, function () {
    console.log("server started at port 8000")});

