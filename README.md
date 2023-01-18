# proto.boilerplate
proto repository boilerplate

# Template usage
This template is intended to be used in the generation of proto projects, it is configured with 
all the scripts and workflows for nodejs and java proto package generation and publish.

The purpose of a proto repository is to define the domain language and messages for a given 
service or domain model.

# Getting started

1. Update the project with your repository name and information in the following places:
   - `nodejs/package.json`
   - `java/build.gradle` 
   - `java/settings.gradle`
   - `.github/workflows/pr.yml`
   - `src/buf.yaml`

# Further documentation

This project uses [buf](https://docs.buf.build/introduction) for generating the
stubs of the different languages configured.

# Documentation

The documentation is autogenerated with the documentation in the proto files, can
be seen in the `doc` directory