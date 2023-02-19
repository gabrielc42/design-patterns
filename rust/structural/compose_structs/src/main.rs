#![allow(dead_code, unused_variables)]

fn main() {
    println!("Hello, world!");
    let a: &mut A = &mut A {
        b: B { f2: 1 },
        c: C { f1: 0, f3: 2 },
    };
    baz(a);
}

// where borrow checker foils us in our plan to use a struct

struct A_err {
    f1: u32,
    f2: u32,
    f3: u32,
}

fn foo_err(a: &mut A_err) -> &u32 {
    &a.f2
}
fn bar_err(a: &mut A_err) -> u32 {
    a.f1 + a.f3
}

fn baz_err(a: &mut A_err) {
    // latter usage of x causes a to be borrow for rest of the function
    let x = foo_err(a);

    // borrow checker error:
    // let y = bar(a); // cannot borrow '*a' as mutable more than once at a time
    println!("{}", x);
}

// can apply pattern this way:
struct A {
    b: B,
    c: C,
}
struct B {
    f2: u32,
}
struct C {
    f1: u32,
    f3: u32,
}

// functions take B or C, rather than A
fn foo(b: &mut B) -> &u32 {
    &b.f2
}
fn bar(c: &mut C) -> u32 {
    c.f1 + c.f3
}

fn baz(a: &mut A) {
    let x = foo(&mut a.b);
    // now it is ok! :D
    let y = bar(&mut a.c);
    println!("{}", x);
}
