use std::path::{Path, PathBuf};

fn main() {
    println!("Hello, world of functional programming!");

    // imperative
    let mut sum = 0;
    for i in 1..11 {
        sum += i;
    }
    println("{}", sum);

    // declarative
    println!("{}", (1..11).fold(0, |a, b| a + b));

    // generics as type classes
    // if users were to make a mistake and use wrong type
    let mut socket = crate::bootp::listen()?;
    while let Some(request) = socket.next_request()? {
        match request.mount_point().as_ref()
            "/secure" => socket.send("Access denied"),
            _ => {} // continue on...
            
    }
    // rest of code here

    // syntax eror, FileDownloadRequest<Bootp> doesn't implement mount_point()
    // only FileDOwnloadRequest<Nfs> does, which is created by NFS module, not BOOTP

    // generics as type classes is used throughout the Standard Library
    // Vec<u8> can be cast from a String, unlike every other type of Vec<T>
    // also can be cast into a binary heap, but only if they contain a type that implements the Ord trait
    // to_string method was specialized for Cow only of type str

    // embedded-half
    // hyper HTTP client library
    // type state pattern

    // Lenses and Prisms
    // lens allows accessing parts of a data type in an abstract, unified way
    // similar to how Rust traits work with type erasure, but a bit more power and flexibility

    //  Lenses example:

    // two types have a customer ID number which corresponds to a person
    // a struct could represent each of these types, and a trait would have a get_customer_id function implemented

    // lenses however allow code supporting customer ID to be moved from the type to the accessor function
    // thanks to lens-rs crate

    // in the example at the end of this file, a simple function like unique_ids_lens operates on a single lens
    // a prism is a function that operates on a family of lenses

    // Prisms
    // same way traits allow "lens-like" design with static polymorphism and dynamic dispatch,
    // prims split problems into multiple associated types to be composed
    // Serde crate: Deserializer trait example in design book
    // this example explains in depth 
    // a deserializing method, pattern to do so, and solution to the above problem
}

// generics as type classes

// advantages:
// 1. allows fields that are common to multiple states to be de-duplicated.
// making non-shared fields generic means they are implemented once
// 2. impl blocks are easier to read, as they are broken down by state
// methods common to all states are typed once in one block
// methods unique to one state are in a separate block

// disadvantages:
// increases size of binary, due to way monomorphization is implemented in the compiler

// if a type seems to need a "split API" due to construction or partial initialization, consider Builder pattern
// if the API between types doesn't change, only the behavior does, then consider Strategy pattern

// enum AuthInfo {
//     Nfs(crate::nfs::AuthInfo),
//     Bootp(crate::bootp::AuthInfo),
// }

// struct FileDownloadRequest {
//     file_name: PathBuf,
//     authentication: AuthInfo,
// }
// requires programmer to do a runtime check

// struct FileDownloadRequest {
//     file_name: PathBuf,
//     authentication: AuthInfo,
//     mount_point: Option<PathBuf>,
// }

// impl FileDownloadRequest {
//     // ... other methods
//     // Gets an NFS mount point if this is an NFS request
//     // Otherwise, return None
//     pub fn mount_point(&self) -> Option<&Path> {
//         self.mount_point.as_ref()
//     }
// }

// adding a generic type to split the API

mod nfs {
    #[derive(Clone)]
    pub(crate) struct AuthInfo(String); // NFS session management omitted
}

mod bootp {
    pub(crate) struct AuthInfo(); // no authentication in bootp
}

// private module, unless outside users invent their own protocol kinds
mod proto_trait {
    use std::path::{ Path, PathBuf };
    user super::{ bootp, nfs };

    pub(crate) trait ProtoKind {
        type AuthInfo;
        fn auth_info(&self) -> Self::AuthInfo;
    }

    pub struct Nfs {
        auth: nfs::AuthInfo,
        mount_point: PathBuf,
    }

    impl Nfs {
        pub(crate) fn mount_point(&self) -> &Path {
            &self.mount_point
        }
    }

    impl ProtoKind for Nfs {
        type AuthInfo = nfs::AuthInfo;
        fn auth_info(&self) -> Self::AuthInfo {
            self.auth.clone()
        }
    }

    pub struct Bootp(); // no additional metadata

    impl ProtoKind for Bootp {
        type AuthInfo = bootp::AuthInfo;
        fn auth_info(&self) -> Self::AuthInfo {
            bootp::AuthInfo()
        }
    }
}

use proto_trait::ProtoKind; // keep internal to prevent impls
pub use proto_trait::{Nfs, Bootp}; // re-export so callers can see them

struct FileDownloadRequest<P: ProtoKind> {
    file_name: PathBuf,
    protocol: P,
}

// all common API parts go into a generic impl block
impl<P: ProtoKind> FileDownloadRequest<P> {
    fn file_path(&self) -> &Path {
        &self.file_name
    }

    fn auth_info(&self) -> P::AuthInfo {
        self.protocol.auth_info()
    }
}

// all protocol-specific impls go into their own block
impl FileDownloadRequest<Nfs> {
    fn mount_point(&self) -> &Path {
        self.protocol.mount_point()
    }
}

// Lenses example
use std::collections::HashSet;

pub struct Account {
    account_id: u32,
    account_type: String,
    // other fields omitted
}

pub trait CustomerId {
    fn get_customer_id(&self) -> u64;
}

pub struct CreditRecord {
    customer_id: u64,
    name: String,
    dob: String,
    // other fields omitted
}

impl CustomerId for CreditRecord {
    fn get_customer_id(&self) -> u64 {
        self.customer_id
    }
}

pub struct AccountRecord {
    customer_id: u64,
    accounts: Vec<Account>,
}

impl CustomerId for AccountRecord {
    fn get_customer_id(&self) -> u64 {
        self.customer_id
    }
}

// static polymorphism: only one type, but each function call can choose it
fn unique_ids_set<R: CustomerId>(records: &[R]) -> HashSet<u64> {
    records.iter().map(|r| r.get_customer_id()).collect()
}

// dynamic dispatch: iterates over any type with a customer ID, collecting all
// values together
fn unique_ids_iter<I>(iterator: I) -> HashSet<u64>
    where I: Iterator<Item=Box<dyn CustomerId>>
{
    iterator.map(|r| r.as_ref().get_customer_id()).collect()
}

// lens-rs crate
use std::collections::HashSet;

use lens_rs::{optics, Lens, LensRef, Optics};

#[derive(Clone, Debug, Lens /* derive to allow lenses to work */)]
pub struct CreditRecord {
    #[optic(ref)] // macro attribute to allow viewing this field
    customer_id: u64,
    name: String,
    dob: String,
    // other fields omitted
}

#[derive(Clone, Debug)]
pub struct Account {
    account_id: u32,
    account_type: String,
    // other fields omitted
}

#[derive(Clone, Debug, Lens)]
pub struct AccountRecord {
    #[optic(ref)]
    customer_id: u64,
    accounts: Vec<Account>,
}

fn unique_ids_lens<T>(iter: impl Iterator<Item = T>) -> HashSet<u64>
where
    T: LensRef<Optics![customer_id], u64>, // any type with this field
{
    iter.map(|r| *r.view_ref(optics!(customer_id))).collect()
}

// Prisms
// deserialize Serde macro

use serde::Deserialize;

#[derive(Deserialize)]
struct IdRecord {
    name: String,
    customer_id: String,
}

// defined as
pub trait Deserialize<'de>: Sized {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: Deserializer<'de>;
}

// the only reason this model works on any format and any type is 
// because the deserializer trait's output type is specified by the implementor of Visitor it is passed,
// rather than being tied to one specific type, which wasn't true in the account example