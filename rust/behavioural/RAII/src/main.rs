// Resource Acquisition is Initialization

use std::ops::Deref;

struct Foo {}

struct Mutex<T> {
    // reference to data: T
}

struct MutexGuard<'a, T: 'a> {
    data: &'a T,
}

// locking mutex is explicit

impl<T> Mutex<T> {
    fn lock(&self) -> MutexGuard<T> {
        // lock underlying OS mutex

        // MutexGuard keeps a reference to self
        MutexGuard { data: self }
    }
}

// destructor for unlocking mutex
impl<'a, T> Drop for MutexGuard<'a, T> {
    fn drop(&mut self) {
        // unlock underlying OS mutex
    }
}

// implwmenting Deref means we can treat MutexGuard like a pointer to T
impl<'a, T> Deref for MutexGuard<'a, T> {
    type Target = T;

    fn deref(&self) -> &T {
        self.data
    }
}

fn baz(x: Mutex<Foo>) {
    let xx = x.lock();
    xx.foo(); // foo is a method on Foo
              // borrow checker ensures we cannot store a reference to the underlying
              // Foo which will outlive guard xx

    // x is unlocked when we exit this function and xx's destructor is executed
}
