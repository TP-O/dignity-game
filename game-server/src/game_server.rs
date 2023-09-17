use crate::handle_err::Error;

use crate::lib::establish_connection;
use crate::models::*;
use diesel::prelude::*;
use crate::schema::users::dsl::*;

#[derive(Debug, Clone)]
pub struct Game {

}
impl Game {
    pub async fn new() -> Self {
        Game{}
    }
    
    pub async fn login(self, usern: String, passw: String) -> Result<Vec<User>, Error>{
        let connection = &mut establish_connection();
        match users.filter(username.eq(usern)).filter(password.eq(passw)).select(User::as_select()).load(connection) {
            Ok(user) => Ok(user),
            Err(e) =>  Err(Error::UserNotFound)
        }

    }

}

