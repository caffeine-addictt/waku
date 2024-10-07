# Labels

!!! FAILURE "Not implemented"

    Labels is a planned feature and currently
    serves no purpose.

Labels describe the labels that are generated
post-setup for supported [Git](https://git-scm.com/) hosting providers:

- **:fontawesome-brands-github: [GitHub](https://github.com)**
- **:fontawesome-brands-gitlab: [GitLab](https://gitlab.com)**

## waku config

=== "Yaml"

    ```yaml
    styles:
      my-style:
        labels:
          # The name of the label
          # (required)
          - name: ""

          # The color of the label.
          # Only HEX is allowed; #fff, #ffffff
          # (required)
            color: ""

          # The description of the label
          # (optional)
            description: ""
    ```

=== "Json"

    ```json
    {
      "labels": [
        {
          // The name of the label
          // (required)
          "name": "",

          // The color of the label.
          // Only HEX is allowed; #fff, #ffffff
          // (required)
          "color": "",

          // The description of the label
          // (optional)
          "description": ""
        }
      ]
    }
    ```
