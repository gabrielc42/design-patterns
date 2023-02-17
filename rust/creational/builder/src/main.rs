#[derive(Debug, PartialEq)]
pub struct Foo {
    // many complicated fields
    bar: String,
}

impl Foo {
    // this helps discovery of builder
    pub fn builder() -> FooBuilder {
        FooBuilder::default()
    }
}

#[derive(Default)]
pub struct FooBuilder {
    // lots of optional fields
    bar: String,
}

impl FooBuilder {
    pub fn new(/* ... */) -> FooBuilder {
        // set minimally required fields of Foo
        FooBuilder {
            bar: String::from("X"),
        }
    }

    pub fn name(mut self, bar: String) -> FooBuilder {
        // set name on builder itself, return builder
        self.bar = bar;
        self
    }

    // if we can get away w/ not consuming the Builder here,
    // FooBuilder can now be a template for constructing many Foos
    pub fn build(self) -> Foo {
        // create a Foo from the FooBuilder,
        // applying all settings in FooBuilder to Foo
        Foo { bar: self.bar }
    }
}

#[test]
fn builder_test() {
    let foo = Foo {
        bar: String::from("Y"),
    };
    let foo_from_builder: Foo = FooBuilder::new().name(String::from("Y")).build();
    assert_eq!(foo, foo_from_builder);
}

fn main() {
    println!("Hello, world builders!");

    // the code above takes and returns builder by value
    // borrow checker naturally references mutably, allowing code like:
    let mut fb = FooBuilder::new();
    fb.a();
    fb.b();
    let f = fb.build();
    FooBuilder::new().a().b().build();
}

// seen in Rust more frequently than other languages because
// Rust lacks overloading
