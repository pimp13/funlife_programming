use bcrypt::{BcryptResult, DEFAULT_COST, hash, verify};

pub fn hash_password(pass: String) -> BcryptResult<String> {
    hash(pass, DEFAULT_COST)
}

pub fn verify_password(pass: &str, hash: String) -> BcryptResult<bool> {
    verify(pass, &hash)
}
