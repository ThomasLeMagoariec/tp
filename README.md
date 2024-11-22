# TP

A simple tool to make it easier creating C# solutions and setting up the git repo.

Made to be used by EPITA students, but could work for other cases

## Usage

EPITA students make sure you git clone the TP before running this tool.

Once inside the directory created by the clone command you can run:

```bash
tp <project name> [FLAGS]
```

```<project name>``` will be the name given to the solution

FLAGS:

    - ng: this won't create a .gitignore file
    - nr: this won't create a README file
    - ngr: won't create neither a .gitignore or a README
    - nrg: same thing as -ngr

If you encouter any errors or have any questions feel free to contact me on discord: thomas.lemago

## Install

### Linux:

Download the `tp` file and move it to `/usr/local/bin`

### Windows:

Download the `tp.exe` file and it to path via the Modify system variables app

### MacOS

Not supported, but you can download the source code and build it yourself