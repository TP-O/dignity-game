// @generated automatically by Diesel CLI.

diesel::table! {
    rooms (id) {
        id -> Int4,
        user_id -> Int4,
    }
}

diesel::table! {
    users (id) {
        id -> Int4,
        username -> Varchar,
        password -> Varchar,
    }
}

diesel::joinable!(rooms -> users (user_id));

diesel::allow_tables_to_appear_in_same_query!(
    rooms,
    users,
);
