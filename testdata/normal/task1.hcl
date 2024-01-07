task {
    name = "task 1"
    comment = "simple task to run commands"

    run_command "echo1" {
        command = ["echo", "aaa"]
    }

    run_command "echo2" {
        command = ["echo", "bbb"]
    }
}
