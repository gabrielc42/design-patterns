## structural design patterns

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
- can make changes to existing subsystem and not affect client
- think complexity of car starting -> startEngine() and stopEngine()

**flyweight:** share parts of object state between multiple objects
