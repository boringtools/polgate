package polgate

import rego.v1

default allow := false

allow if {
    count(violations) == 0
}