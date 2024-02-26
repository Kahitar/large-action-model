table "users2" {
    schema = schema.main

    column "id" {
        null = true
        type = text
    }
    column "username" {
        null = false
        type = text
    }

    primary_key {
        columns = [column.id]
    }
}
schema "main" {
}
