I could have done all this with awk in less than 6 lines smh. Just trying to get conformable with Go's
syntax.

# Grade Tracker

Program to simply report the current grade in a course. Most of similar tools like this online don't have
any way to save the grades.

## Usage

`./proc-tracker-platform-arch <path_to_file>`

The file:

1. Needs to be in this format: `Name, achieved, possible, % of final grade` (csv, see `sample.bighead`)
2. File extension does not matter
3. `#` is treated like a comment and ignored
4. See `result.bighead` for example output
5. Keep 'name' under 15 characters to preserve formatting
6. If there's any issue with the file, execution stops immediately

## If you 100% trust me and want to use the binaries I compiled

1. You don't need to install go
2. Download your version from the `release` section on the repo homepage (if I can figure it out).
    1. Try the amd64 version first
3. There's no gui, so run from command line

## Build and compile yourself because you don't trust me

1. The code was written with go version 1.20.3. So, you need to have that version or higher
    1. I didn't really look much into the versions, so it might work with older versions
2. See Makefile, and note your platform and architecture. Try `amd64` version first
3. run `make <grade-tracker-platform-arch>`
4. or you can simply run `make` and try everything to find which works
