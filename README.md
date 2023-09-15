# go-routesC1
go routes without use of INTERFACES
In this the controllers and services are used wihtout the help of the interfaces.
# main 
In the main fuciton we use the initall 
withe the help of the we are initialising the collection and getting its instance in return 
and we are passing it to the services 
# services
In services we are creating two variables one is of collection type and another is of the context.
we are initializing them with the help of the main funciton initall fuction is use for it
so that insertone, insertmay all those function can be invoked inside those particular collection (the collection which we are passing from the main function)
-> all the fuctions in the services are created as methods so we need to call them in CONTROLLERS
# controllers
In controllers we need to call the methods in the services
since we need to call those methods in the service we need to pass those controllers and ctx with the help of initall function in the main


// SO IN THIS PROGRAM WE CAN EXECUTRE THE FUNCTIONS AND METHODS WITHOUT THE HELP OF THE INTERFACES BY DIRECTLY PASSING THE COLLECTIONS AND CTX.
