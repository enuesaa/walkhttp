batch {
    name = "task 1"
    comment = "simple task"

    task "echo1" {
        run {
            command = ["echo", "aaa"]
        }
    }

    task "echo2" {
        run {
            command = ["echo", "bbb"]
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
