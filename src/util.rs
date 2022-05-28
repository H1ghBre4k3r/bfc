use std::{fs, io};

/// Read the content of a file.
pub fn read_file(path: &String) -> Result<String, io::Error> {
    match fs::read(path) {
        Err(err) => Err(err),
        Ok(content) => Ok(String::from_utf8(content).unwrap()),
    }
}
