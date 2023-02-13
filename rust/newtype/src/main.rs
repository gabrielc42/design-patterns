// a tuple struct with a single field to make an opaque wrapper
// for a type. This creates a new type,
// rather than an alias to a type (type items).

// abstraction: example -> wrapping f64 to give distinguishable Miles and Kilometres 

// some type, not necessarily in same module or crate, even
struct Foo {
    //...
}

impl Foo {
    // these functions are not present on Bar
}

// new type
pub struct Bar(Foo);

impl Bar {
    // constructor
    pub fn new(//...
    ) -> Self {
        //...
    }
    // ...
}

fn main() {
    let b = Bar::new(...);

    // Foo and Bar are type incompatible, does not type check:
    // let f: Foo = b;
    // let b: Bar = Foo { ... };
}
