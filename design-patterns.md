# design patterns:

## structural

**adapter:** constructor's instance of a different abstract/interface type, translates paramerters to appropriate format and directly calls one or several methods of the wrapped object

**bridge:** separate out responsibilities into different abstract classes
- when parent abstract class defines basic rules, then concrete classes add
rules
- abstract class that references objects, abstract methods that will be defined
in each of the concrete classes

**proxy:** 
- simplified version of complex or heavy object (skeleton object)
- when original object is present in different address space, and we want
to represent it locally
- add a layer of security to underlying object

**composite:** treating individual objects and compositions of objects in the same way

**decorator:** additional things added to an object statically or dynamically, enhanced interface to original object

**facade:** encapsulates complex subsystem behind a simple interface
- can makes changes to existing subsystem and not affect client
- think complexity of car starting -> startEngine() and stopEngine()

**flyweight:** share parts of object state between multiple objects

## creational

**builder:** collects many constructors into one tight constructor

**factory:** class that gathers classes of similar extension/implementations

**abstract factory:** 

**object pool:** pool of objects initialized and kept for retrieval and return
- when cost to create object of the class is high
and number of such objects that will be needed at a particular time is not much
- when pool object is immutable

**prototype:** copy original object to a new object, then modify it accordingly

**singleton:** one object instantiation to a class


## behavioural

**chain of responsibility:** processing objects responsible for a command,
forwards command to next processor, etc.

**null object:** instantiation of a 'do nothing' class to not deal w/ null references

**observer:** allows objects to notify other objects about changes in their state

**state:** object changing behavior when its internal state changes

**strategy:** set of behaviors -> interchangeable objects inside original context object

**command:** encapsulates what method to call, the method's arguments, 
and the object to which the method belongs for performing a given action

**interpreter:** defines the grammar of a particular language in an 
object-oriented way which can be evaluated by the interpreter itself

**visitor:** new operation w/out introducing modifications to existing object structure

**iterator:** iterates through objects

**mediator:** object that takes care of interaction between dependent objects

**memento:** implement undoable actions, saves state of object at instance and restoring if actions performed need to be undone

**template:** skeleton of algorithm in base class and subclasses override steps w/out changing algorithim's structure

**RAII:** *Rust 

**newtype:** *Rust  


## misc


### Rust
**compose structs:**

**prefer small crates:** Split crates into multiples to allow more parallel built code, more reusable crates, easier to understand
- The ref_slice crate provides functions for converting &T to &[T].
- The url crate provides tools for working with URLs.
- The num_cpus crate provides a function to query the number of CPUs on a machine.
- url crate from Servo has established multiple usage, however keep in mind versions (1.0 and 0.5 url have different types = dependency hell)

**contain unsafe in small modules:** embed unsafe blocks into small modules for larger modules
- examples include: toolshed and std's String
