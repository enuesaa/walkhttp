batch {
    name = "task1"
    comment = "taskfile for test"

    task "start" {
        run {
            command = ["echo", "start"]
        }
    }
    
    task "example" {
        invoke {
            url = "https://example.com"
        }
    }

    task "end" {
        run {
            command = ["echo", "end"]
        }
    }
}
