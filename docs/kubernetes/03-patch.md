Start kubectl under the debugger:

    gdb kubectl

Find all functions that match ToJSON:

    (gdb) info func ToJSON
    All functions matching regular expression "ToJSON":
    ...
    void k8s.io/kubernetes/pkg/util/yaml.ToJSON(struct []uint8, struct []uint8, error);
    ...

Set a breakpoint at yaml.ToJSON:

    (gdb) b k8s.io/kubernetes/pkg/util/yaml.ToJSON
    Breakpoint 1 at 0x8c4300: file .../src/k8s.io/kubernetes/pkg/util/yaml/ decoder.go

Provide arguments to kubectl, and resume execution:

    (gdb) r patch deployment in-place-demo --patch 'spec:\n replicas: 7'
    Starting program: /usr/local/bin/kubectl patch deployment in-place-demo --patch 'spec:\n replicas: 7'
    ...
    [New LWP 26442]

Hit the breakpoint:

    Thread 1 "kubectl" hit Breakpoint 1, k8s.io/kubernetes/pkg/util/yaml.ToJSON ...

View the arguments passed to ToJSON:

    (gdb) info args
    data = {array = 0xc4205dedc0 "spec:\\n replicas: 7", len = 19, cap = 32}

We can see the problem. The YAML buffer should have a newline character, but
instead it has a backslash character followed by an n character.

Get a more detailed look at the buffer:

    (gdb) x /20bc 0xc4205dedc0
    0xc4205dedc0:	115 's'	112 'p'	101 'e'	99 'c'	58 ':'	92 '\\'	110 'n'	32 ' '
    0xc4205dedc8:	114 'r'	101 'e'	112 'p'	108 'l'	105 'i'	99 'c'	97 'a'	115 's'
    0xc4205dedd0:	58 ':'	32 ' '	55 '7'	0 '\000'

We see the sequence of two characters:

    92 '\\'	110 'n'

Let the program run and finish:

    (gdb)continue
    (gdb)quit

Verify that the patch did not work:

    kubectl get deployment in-place-demo --output=yaml

    spec:
      progressDeadlineSeconds: 600
      replicas: 6

In the outpu, we can see that the number of replicas was not changed to 7.

Start again.

    gdb kubectl

    (gdb) b k8s.io/kubernetes/pkg/util/yaml.ToJSON
    Breakpoint 1 at 0x8c4300: file /go/src/k8s.io/kubernetes/_output/dockerized/go/src/k8s.io/kubernetes/pkg/util/yaml/decoder.go, line 37.

This time, put a $ in front of the patch string:

    (gdb) r patch deployment in-place-demo --patch $'spec:\n replicas: 7'
    Starting program: /usr/local/bin/kubectl patch deployment in-place-demo --patch $'spec:\n replicas: 7'
    ...
    [New LWP 26524]

Hit the breakpoint:

    Thread 1 "kubectl" hit Breakpoint 1, k8s.io/kubernetes/pkg/util/yaml.ToJSON ...

View the arguments passed to ToJSON:

    (gdb) info args
    data = {array = 0xc4205ce2e0 "spec:\n replicas: 7", len = 18, cap = 32}

This time, we see a newline character in the YAML buffer.

Take a closer look at the YAML buffer:

    (gdb) x /20bc 0xc4205ce2e0
    0xc4205ce2e0:	115 's'	112 'p'	101 'e'	99 'c'	58 ':'	10 '\n'	32 ' '	114 'r'
    0xc4205ce2e8:	101 'e'	112 'p'	108 'l'	105 'i'	99 'c'	97 'a'	115 's'	58 ':'
    0xc4205ce2f0:	32 ' '	55 '7'	0 '\000'	0 '\000'

We see the newline character:

    10 '\n'

Let the program finish:

    (gdb)continue
    (gdb)quit

Verify that the patch succeeded:

    kubectl get deployment in-place-demo --output=yaml

    apiVersion: extensions/v1beta1
    kind: Deployment
    metadata:
      annotations:
        deployment.kubernetes.io/revision: "1"
        kubernetes.io/change-cause: |-
          kubectl patch deployment in-place-demo --patch spec:
           replicas: 7
      ...
    spec:
      ...
      replicas: 7











