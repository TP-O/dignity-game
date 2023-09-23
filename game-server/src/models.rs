use diesel::prelude::*;
use serde::Serialize;

use crate::schema::{rooms, users};


#[derive(Queryable, Identifiable, Selectable, Debug, PartialEq,Serialize)]
#[diesel(table_name = users)]
pub struct User {
    pub id: i32,
    pub username: String,
    pub password: String,
}

#[derive(Queryable, Selectable, Identifiable, Associations, Debug, PartialEq)]
#[diesel(belongs_to(User))]
#[diesel(table_name = rooms)]
pub struct Room {
    pub id: i32,
    pub user_id: i32,
}