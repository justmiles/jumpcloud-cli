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

## Examples

List JumpCloud users

    jc user list --output table \
      --query "[].{Id: _id, UserName: username, FirstName: firstname, LastName: lastname, Email: email, MAC: attributes[?name == 'hwaddr'] | [0].value}"

List attributes for JumpCloud user

    jc user attributes --user <username>

Check for matching attributes

    jc user attribute-matches --user <username --key <attribute name> --value <attribute match>

List Tables

    jc group list --query "[?type == 'user_group'].{ID: id, Name: name}" --output table

List Users in Group

    jc group list-members --id $(jc group list --query "[?name == 'administrator'].{ID: id}" --output csv) --output table --query '[].{Email: email, Userame: username, FirstName: firstname, LastName: lastname}'

    jc group list-members --id $(jc group list --query "[?name == 'administrator'].{ID: id}") --output json --query "[].{Email: email, MAC: attributes[?name == 'hwaddr'].value}"
