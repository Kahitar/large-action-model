table "clients" {
    schema = schema.main

    column "id" {
        null = true
        type = text
    }
    column "secret" {
        null = false
        type = text
    }
    column "domain" {
        null = false
        type = text
    }

    primary_key {
        columns = [column.id]
    }
}

schema "main" {
}

