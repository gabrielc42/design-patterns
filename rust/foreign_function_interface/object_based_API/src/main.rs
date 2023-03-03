fn main() {
    // 1. all encapsulated types shouled be owned by Rust,
    // managed by the user, and opaque
    // 2. all transactional data types should be owned by the user,
    //  and transparent
    // 3. all library behavior should be
    // functions acting upon Encapsulated types
    // 4. all library behavior should be encapsulated into
    // types not based on structure, but life/provenance

    // 3 goals with foreign API
    // 1. Make it easy to use in the target language
    // 2. avoid API dictating internal unsafety on the Rust side as much as possible
    // 3. keep potential for memory unsafety and Rust undefined behaviour as small as possible

    println!("Hello, world of foreign function interfaces!");
}

struct Dbm { ... }

impl Dbm {
    /* ... */
    pub fn keys<'it>(&'it self) -> DbmKeysIter<'it> { ... }
    /* ... */
}

struct DbmKeysIter<'it> {
    owner: &'it Dbm,
}

impl<'it> Iterator for DbmKeysIter<'it> { ... }
