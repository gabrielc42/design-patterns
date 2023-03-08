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
}
