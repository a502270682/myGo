MyGoTemplate

# migrate
migrate -database "mysql://root:@tcp(127.0.0.1:3306)/test" 
-source file://migrations up