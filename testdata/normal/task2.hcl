batch {
    name = "task 2"
    comment = "also invoke http web api"

    task "start" {
        run {
            command = ["echo", "start"]
        }
    }

    task "example" {
        invoke {
            url = "https://example.com"
            // request_headers = {
            //     "Content-Type" = "application/json"
            // }
        }
    }
}
