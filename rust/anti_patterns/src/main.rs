fn main() {
    println!("Hello, world of anti patterns!");

    // clone to satisfy the borrow checker
    // define any variable
    let mut x = 5;

    // borrow `x` -- but clone it first
    let y = &mut (x.clone());

    // w/out x.clone(), following line would fail on compilation
    println!("{}", x);

    // perform some action on the borrow
    // to prevent rust form optimizing this out of existence
    *y += 1;

    // cargo clippy detects some cases in which .clone() is not necessary

    // ![deny(warnings)]
    // opts out of warnings
    // RUSTFLAGS="-D warnings" cargo build

    // Deref polymorphism
    // misuse Deref trait to emulate inheritance between
    // structs, and thus reuse methods

    let b = Bar { f: Foo {} };
    b.m();

    // no struct inheritance in rust

    // traits implemented by Foo are not automatically implemented for
    // Bar, so this pattern has problems with bounds checking and thus generic programming

    // only supports single inheritance,
    // no notion of interfaces, class-based privacy,
    // or other inheritance-related features

    // depending on circumstances it might be better to
    // re-implement using traits or write out facade methods
    // to dispatch to Foo manually

    // Deref trait is designed for implementation of customer pointer types
    // takes a pointer-to-T to a T, not convert between types
}

// deref polymorphism
// sometimes we want to emulate common patterns from
// OO languages such as Java:
// class Foo {
//     void m() { ... }
// }

// class Bar extends Foo {}

// public static void main(String[] args) {
//     Bar b = new Bar();
//     b.m();
// }

// can use deref polymorphism to do so:

use std::ops::Deref;

struct Foo {}

impl Foo {
    fn m(&self) {
        //...
    }
}

struct Bar {
    f: Foo,
}

impl Deref for Bar {
    type Target = Foo;
    fn deref(&self) -> &Foo {
        &self.f
    }
}
