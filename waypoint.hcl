project = "{{.ProjectName}}"

app "{{.ProjectName}}" {
  build {
    use "docker" {
    }
  }

  deploy {
    use "docker" {
    }
  }
}
