To install jx-convert-jenkinsfile 0.0.2 see the below. To run, see [the project's README](https://github.com/jenkins-x/jx-convert-jenkinsfile/blob/v0.0.2/README.md).

### Linux

```shell
curl -L https://github.com/jenkins-x/jx-convert-jenkinsfile/releases/download/v0.0.2/jx-convert-jenkinsfile-linux-amd64.tar.gz | tar xzv 
sudo mv jx-convert-jenkinsfile /usr/local/bin
```

### macOS

```shell
curl -L https://github.com/jenkins-x/jx-convert-jenkinsfile/releases/download/v0.0.2/jx-convert-jenkinsfile-darwin-amd64.tar.gz | tar xzv 
sudo mv jx-convert-jenkinsfile /usr/local/bin
```
## Changes

### Bug Fixes

* Don't fail for arbitrary credentials in environment
* Handle container with named steps
* Handle multiline strings and double quotes inside single quoted strings

### Chores

* Add link to create issues in jx repo to README
* Fix the README link in the changelog header.
