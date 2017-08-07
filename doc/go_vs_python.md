# Compare Go/Python Cluster-Loader

## Path

svt-go can be run at any folder because:

* The path of the conf file can be absolute or relative to the current folder.
* The paths of the templates can be absolute or relative to the svt
binary's folder.

## Template and template only
All objects for a project are create via templates, which means that
each template will be processed before being used to create objects.

The <code>NAMESPACE</code> is a parameter of a template like <code>IDENTIFIER</code>
in the python version. This will save the time of loading json and injecting
variables into it.

Quotas are no longer needed since they are templates anyway.

Parameters for the templates are treated as a map instead of a list of
objects because the former fits more in the situation.


## Parallelism/Concurrency

Go uses [goroutines](https://tour.golang.org/concurrency/1) while python
forks sub-processes.


## Compilation and Runtime

With svt-go, we can choose to run it with source code or binaries (with
packaged conf). It has both pros/cons for both ways. Using source code
direclty results in requiring Go environment while binaries might need
to update the pkg for any code change.

Python runs the source code and python environment comes with Linux out
of the box. The dependencies are installed via Ansible using pip. SVT
does not support binary distribution yet.