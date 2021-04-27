# JumpCloud CLI

A quick and dirty CLI to interact with JumpCloud

I've only populated API methods as I've needed them but will add more over time.

`jc --help`

    cli to interact with JumpCloud

    Usage:
      jc [command]

    Available Commands:
      help        Help about any command
      user        interact with JumpCloud users

    Flags:
      -h, --help      help for jc
          --version   version for jc

    Use "jc [command] --help" for more information about a command.

`jc user --help`

    interact with JumpCloud users

    Usage:
      jc user [command]

    Available Commands:
      attribute-matches exits successfully if a user's attribute key/value pair matches
      attributes        show attributes for a user
      list              list jumpcloud users

    Flags:
      -h, --help   help for user

    Use "jc user [command] --help" for more information about a command.

## Run A Command

    CRUD - create/read/update/delete
    jc command list|create|d  --name --


    jc command execute --system <x> --group <x> --command-name "" --adoc-command "dir"

    > ......

    dir

## Examples

List JumpCloud users

    jc user list --output table --query "[].{_id:_id, UserName: username, FirstName: firstname, LastName: lastname, Email: email}"

List JumpCloud users with a custom attribute

    jc user list --output table --query "[].{_id:_id, UserName: username, FirstName: firstname, LastName: lastname, Email: email, MAC: attributes[?name == 'hwaddr'].value}"

List attributes for JumpCloud user

    jc user attributes --user <username>

Check for matching attributes

    jc user attribute-matches --user <username --key <attribute name> --value <attribute match>

List groups with table output

    jc group list --query "[?type == 'user_group'].{ID: id, Name: name}" --output table

List users in group

    jc group list-members --id $(jc group list --query "[?name == 'retool'].{ID: id}" --output csv) --output table --query '[].{Email: email, Userame: username, FirstName: firstname, LastName: lastname}'

List users in a hroup with custom attributes

    jc group list-members --id <some group id> --output csv) --output table --query "[].{Email: email, MAC: attributes[?name == 'hwaddr'].value}"

List systems

    jc system list --output table --query "reverse(sort_by(@, &lastContact)) | [].{_id:_id, DeviceName: displayName, LastContact: lastContact, OperatingSystem: os, AgentVersion: agentVersion}"

List systems with thei IP addresses

    jc system list --output table --query "reverse(sort_by(@, &lastContact)) | [].{_id:_id, DisplayName: displayName, LastContact: lastContact, OperatingSystem: os, AgentVersion: agentVersion, IPs: join(',',networkInterfaces[?family == 'IPv4'].address)}"

## Roadmap

- [ ] List system groups

## Reference

- [JumpCloud V1 Documentation](https://docs.jumpcloud.com/1.0)
- [JumpCloud V2 Documentation](https://docs.jumpcloud.com/2.0)