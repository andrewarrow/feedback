
import sys
import os
from placeit import placeit

path = sys.argv[1]
name = sys.argv[2]

def feedback():
    template = """\
{
  "footer": "github.com/andrewarrow/feedback",
  "title": "{{name}}",
  "routes": [{"root": "sessions", "paths": [
                     {"verb": "GET", "second": "", "third": ""},
                     {"verb": "GET", "second": "*", "third": ""},
                     {"verb": "POST", "second": "", "third": ""}
             ]},
             {"root": "users", "paths": [
                     {"verb": "GET", "second": "", "third": ""},
                     {"verb": "GET", "second": "*", "third": ""},
                     {"verb": "GET", "second": "thing", "third": "*"},
                     {"verb": "POST", "second": "", "third": ""}
             ]}
  ],
  "models": [
    {
      "name": "user",
      "fields": [
        {
          "name": "username",
          "flavor": "username",
          "index": "unique",
          "regex": "^[\\\\+@\\\\.a-zA-Z0-9_]{2,50}$"
        },
        {
          "name": "email",
          "flavor": "name",
          "index": "unique",
          "regex": "^[a-zA-Z0-9._%\\\\+-]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,}$",
          "null": "yes"
        },
        {
          "name": "slug",
          "flavor": "name",
          "index": "unique",
          "regex": "^[\\\\-a-zA-Z0-9]{2,20}$",
          "null": "yes"
        },
        {
          "name": "password",
          "flavor": "fewWords",
          "index": "",
          "required": "",
          "regex": "^.{8,100}$",
          "null": ""
        }
      ]
    },
    {
      "name": "ip_data",
      "fields": [
        {
          "name": "ip",
          "flavor": "name",
          "index": "unique"
        },
        {
          "name": "content",
          "flavor": "json"
        }
      ]
    },
    {
      "name": "admin",
      "fields": [
        {
          "name": "user_id",
          "flavor": "int",
          "index": "yes"
        }
      ]
    },
    {
      "name": "cookie_token",
      "fields": [
        {
          "name": "guid",
          "flavor": "uuid",
          "index": "yes",
          "required": "",
          "regex": "",
          "null": ""
        },
        {
          "name": "user_id",
          "flavor": "int",
          "index": "yes",
          "required": "",
          "regex": "",
          "null": ""
        }
      ]
    }
  ]
}
    """

    placeit("app/feedback.json", {"name": name}, template)

