# other design patterns:


### Rust
**compose structs:**

**prefer small crates:** Split crates into multiples to allow more parallel built code, more reusable crates, easier to understand
- The ref_slice crate provides functions for converting &T to &[T].
- The url crate provides tools for working with URLs.
- The num_cpus crate provides a function to query the number of CPUs on a machine.
- url crate from Servo has established multiple usage, however keep in mind versions (1.0 and 0.5 url have different types = dependency hell)

**contain unsafe in small modules:** embed unsafe blocks into small modules for larger modules
- examples include: toolshed and std's String
