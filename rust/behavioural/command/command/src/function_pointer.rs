type FnPtr = fn() -> String;

pub struct Command {
    execute: FnPtr,
    rollback: FnPtr,
}

pub struct Schema {
    commands: Vec<Command>,
}

impl Schema {
    pub fn fp_new() -> Self {
        Self { commands: vec![] }
    }
    pub fn fp_add_migration(&mut self, execute: FnPtr, rollback: FnPtr) {
        self.commands.push(Command { execute, rollback });
    }
    pub fn fp_execute(&self) -> Vec<String> {
        self.commands.iter().map(|cmd| (cmd.execute)()).collect()
    }
    pub fn fp_rollback(&self) -> Vec<String> {
        self.commands
            .iter()
            .rev()
            .map(|cmd| (cmd.rollback)())
            .collect()
    }
}

pub fn fp_add_field() -> String {
    "add field".to_string()
}

pub fn fp_remove_field() -> String {
    "remove field".to_string()
}
