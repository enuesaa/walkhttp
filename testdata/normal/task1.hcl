task {
    name = "task 1"
    comment = "simple task"

    run_command "echo1" {
        command = ["echo", "aaa"]
    }

    run_command "echo2" {
        command = ["echo", "bbb"]
    }
    
    invoke "example" {
        url = "https://example.com"
        // request_headers = {
        //     "Content-Type" = "application/json"
        // }
    }
}
