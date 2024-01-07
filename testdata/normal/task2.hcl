task {
    name = "task 2"
    comment = "also invoke http web api"

    run_command "start" {
        command = ["echo", "start"]
    }
    
    invoke "example" {
        url = "https://example.com"
        request_headers = {
            "Content-Type" = "application/json"
        }
    }
}
