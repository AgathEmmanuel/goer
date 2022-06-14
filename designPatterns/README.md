
# Desing Patterns in go  


## Behavioural Pattern: Observer  

- provides the ability for a subject to notify a set of "obervers" about changes to the subject  
- used to maintain loose coupling between elements  of a system that interact with each other  

## Behavioural Pattern: Iterator  

- describes a way of accessing the elements contained within an object without exposing the  
  underlying implementation of the container  
- useful for objects that exposes collections of elements, for example a library object that keeps  
  track of books  

## Structural Pattern: Adapter  

- allows the interface of an existing subsystem or API to be used as another interface without  
  modifying the code of the existing API  
- enables incompatible objects to work together  without having to make changes to either one  

## Structural Pattern: Facade  
 
- provide a simple, front-facing interface to a more complex system, library or API  
- improve usability of a more complex API, serve as starting point for refactoring  
  and reduce tight coupling beteween parts of a system  

## Creational Pattern: Builder  

- encapsulates an objects construction process along with specifying the various parts of a complex API  
  enable flexible creation of an object that can have many different representations  
- increses code readability of complex types  
- useful when objects have complex APIs, multiple constructor options, and several different  
  possible representations  


## Creational Pattern: Singleton  

- restrics the instantiation of a class to a single instance and provides glabl access  
  allows for lazy initialization of the class 
- useful when you want to ensure there is only one instance of a class 

## Creational Pattern: Factory  

- allows for construction of objects when the types of those objects is not predetermined at runtime  
- produces code that is more readable when there are multiple ways of creating a particular object  
- used when object creation needs to be flexible and cannot be known before hand   

