use crate::{
    fn_trait_objects::{fn_to_add_field, fn_to_remove_field, Schema as fn_to_Schema},
    function_pointer::{fp_add_field, fp_remove_field, Schema as fp_Schema},
    trait_object::{AddField, CreateTable, Schema},
};

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
    let mut schema_two = fp_Schema::fp_new();
    schema_two.fp_add_migration(|| "create table".to_string(), || "drop table".to_string());
    schema_two.fp_add_migration(fp_add_field, fp_remove_field);
    assert_eq!(vec!["create table", "add field"], schema_two.fp_execute());
    assert_eq!(vec!["remove field", "drop table"], schema_two.fp_rollback());

    // Fn trait objects
    let mut schema_three = fn_to_Schema::fn_to_new();
    schema_three.fn_to_add_migration(|| "create table", || "drop table");
    schema_three.fn_to_add_migration(fn_to_add_field, fn_to_remove_field);
    assert_eq!(
        vec!["create table", "add field"],
        schema_three.fn_to_execute()
    );
    assert_eq!(
        vec!["remove field", "drop table"],
        schema_three.fn_to_rollback()
    );
}
