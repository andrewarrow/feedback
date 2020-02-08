Start Here
==================
[Start here](https://github.com/andrewarrow/feedbacks/blob/master/README.md)

Read [the feedbacks readme](https://github.com/andrewarrow/feedbacks/blob/master/README.md) first. Below is the readme for just 1 feedback.



FEEDBACKS
==================

Feedback is a rails inspired framework but in golang not ruby.

```
cp conf.toml.dist conf.toml
go build
./feedback

http://localhost:3000/
```

SETUP on MAC
==================

```
brew install mysql
brew services start mysql
mysql -uroot

  CREATE USER 'dev'@'localhost' IDENTIFIED BY 'password'; 
  GRANT ALL ON *.* TO 'dev'@'localhost' WITH GRANT OPTION;

create database feedback;
```
