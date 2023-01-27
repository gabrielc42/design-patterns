use crate::trait_object::Schema;

mod fn_trait_objects;
mod function_pointer;
mod trait_object;

fn main() {
    // trait objects
    let mut schema = Schema::new();

    let cmd = Box::new(CreateTable);
    schema.add_migration(cmd);
    let cmd = Box::new(AddField);
    schema.add_migration(cmd);

    assert_eq!(vec!["create table", "add field"], schema.execute());
    assert_eq!(vec!["remove field", "drop table"], schema.rollback());

    // function pointers
    let schema_two = Schema::new();
    schema_two.add_migration(|| "create table".to_string(), || "drop table".to_string());
    schema_two.add_migration(add_field, remove_field);
    assert_eq!(vec!["create table", "add field"], schema.execute());
    assert_eq!(vec!["remove field", "drop table"], schema.rollback());

    // Fn trait objects
    let schema_three = Schema::new();
    schema_three.add_migration(|| "create table", || "drop table");
    schema_three.add_migration(add_field, remove_field);
    assert_eq!(vec!["create table", "add field"], schema.execute());
    assert_eq!(vec!["remove field", "drop table"], schema.rollback());
}
